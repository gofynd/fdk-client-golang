package application



    
    // Meta ...
    type Meta struct {

        
            Source string  `json:"source"`
         
    }
    
    // Media ...
    type Media struct {

        
            Alt string  `json:"alt"`
            Meta Meta  `json:"meta"`
            Type string  `json:"type"`
            URL string  `json:"url"`
         
    }
    
    // Price ...
    type Price struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // ProductListingPrice ...
    type ProductListingPrice struct {

        
            Marked Price  `json:"marked"`
            Effective Price  `json:"effective"`
         
    }
    
    // NetQuantity ...
    type NetQuantity struct {

        
            Value float64  `json:"value"`
            Unit interface{}  `json:"unit"`
         
    }
    
    // CustomMetaFields ...
    type CustomMetaFields struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
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
    
    // ProductBrand ...
    type ProductBrand struct {

        
            Action ProductListingAction  `json:"action"`
            Logo Media  `json:"logo"`
            Description string  `json:"description"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ApplicationItemMOQ ...
    type ApplicationItemMOQ struct {

        
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
            Minimum float64  `json:"minimum"`
         
    }
    
    // ProductDetailAttribute ...
    type ProductDetailAttribute struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            Key string  `json:"key"`
         
    }
    
    // ProductDetailGroupedAttribute ...
    type ProductDetailGroupedAttribute struct {

        
            Details []ProductDetailAttribute  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // ProductCategoryMap ...
    type ProductCategoryMap struct {

        
            L2 ProductBrand  `json:"l2"`
            L1 ProductBrand  `json:"l1"`
            L3 ProductBrand  `json:"l3"`
         
    }
    
    // ApplicationItemSEO ...
    type ApplicationItemSEO struct {

        
            Description interface{}  `json:"description"`
            Title interface{}  `json:"title"`
         
    }
    
    // ProductDetail ...
    type ProductDetail struct {

        
            Medias []Media  `json:"medias"`
            Slug string  `json:"slug"`
            Highlights []string  `json:"highlights"`
            Tags []string  `json:"tags"`
            ItemType string  `json:"item_type"`
            Color string  `json:"color"`
            Price ProductListingPrice  `json:"price"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Brand ProductBrand  `json:"brand"`
            ItemCode string  `json:"item_code"`
            HasVariant bool  `json:"has_variant"`
            Moq ApplicationItemMOQ  `json:"moq"`
            RatingCount float64  `json:"rating_count"`
            TeaserTag string  `json:"teaser_tag"`
            Rating float64  `json:"rating"`
            IsDependent bool  `json:"is_dependent"`
            Attributes map[string]interface{}  `json:"attributes"`
            Discount string  `json:"discount"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            CategoryMap ProductCategoryMap  `json:"category_map"`
            Similars []string  `json:"similars"`
            Tryouts []string  `json:"tryouts"`
            ProductOnlineDate string  `json:"product_online_date"`
            Action ProductListingAction  `json:"action"`
            ShortDescription string  `json:"short_description"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Description string  `json:"description"`
            Seo ApplicationItemSEO  `json:"seo"`
            ImageNature string  `json:"image_nature"`
            Categories []ProductBrand  `json:"categories"`
            Name string  `json:"name"`
            Type string  `json:"type"`
            UID float64  `json:"uid"`
         
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

        
            Col4 ColumnHeader  `json:"col_4"`
            Col1 ColumnHeader  `json:"col_1"`
            Col6 ColumnHeader  `json:"col_6"`
            Col2 ColumnHeader  `json:"col_2"`
            Col3 ColumnHeader  `json:"col_3"`
            Col5 ColumnHeader  `json:"col_5"`
         
    }
    
    // SizeChartValues ...
    type SizeChartValues struct {

        
            Col4 string  `json:"col_4"`
            Col1 string  `json:"col_1"`
            Col6 string  `json:"col_6"`
            Col2 string  `json:"col_2"`
            Col3 string  `json:"col_3"`
            Col5 string  `json:"col_5"`
         
    }
    
    // SizeChart ...
    type SizeChart struct {

        
            Headers ColumnHeaders  `json:"headers"`
            SizeTip string  `json:"size_tip"`
            Unit string  `json:"unit"`
            Image string  `json:"image"`
            Title string  `json:"title"`
            Sizes []SizeChartValues  `json:"sizes"`
            Description string  `json:"description"`
         
    }
    
    // Dimension ...
    type Dimension struct {

        
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
            Length float64  `json:"length"`
            Height float64  `json:"height"`
            Width float64  `json:"width"`
         
    }
    
    // Weight ...
    type Weight struct {

        
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
            Shipping float64  `json:"shipping"`
         
    }
    
    // ProductSize ...
    type ProductSize struct {

        
            Dimension Dimension  `json:"dimension"`
            Value string  `json:"value"`
            Quantity float64  `json:"quantity"`
            Weight Weight  `json:"weight"`
            SellerIdentifiers []string  `json:"seller_identifiers"`
            Display string  `json:"display"`
            IsAvailable bool  `json:"is_available"`
         
    }
    
    // ProductSizeStores ...
    type ProductSizeStores struct {

        
            Count float64  `json:"count"`
         
    }
    
    // ProductSizes ...
    type ProductSizes struct {

        
            Sellable bool  `json:"sellable"`
            Discount string  `json:"discount"`
            SizeChart SizeChart  `json:"size_chart"`
            MultiSize bool  `json:"multi_size"`
            Sizes []ProductSize  `json:"sizes"`
            Price ProductListingPrice  `json:"price"`
            Stores ProductSizeStores  `json:"stores"`
         
    }
    
    // AttributeDetail ...
    type AttributeDetail struct {

        
            Display string  `json:"display"`
            Logo string  `json:"logo"`
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

        
            Medias []Media  `json:"medias"`
            Slug string  `json:"slug"`
            Action ProductListingAction  `json:"action"`
            Value string  `json:"value"`
            ColorName string  `json:"color_name"`
            Color string  `json:"color"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            IsAvailable bool  `json:"is_available"`
         
    }
    
    // ProductVariantResponse ...
    type ProductVariantResponse struct {

        
            Key string  `json:"key"`
            Header string  `json:"header"`
            Items []ProductVariantItemResponse  `json:"items"`
            DisplayType string  `json:"display_type"`
         
    }
    
    // ProductVariantsResponse ...
    type ProductVariantsResponse struct {

        
            Variants []ProductVariantResponse  `json:"variants"`
         
    }
    
    // StoreDetail ...
    type StoreDetail struct {

        
            Name string  `json:"name"`
            City string  `json:"city"`
            ID float64  `json:"id"`
            Code string  `json:"code"`
         
    }
    
    // Seller ...
    type Seller struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Count float64  `json:"count"`
         
    }
    
    // CompanyDetail ...
    type CompanyDetail struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // ProductStockPrice ...
    type ProductStockPrice struct {

        
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            Currency string  `json:"currency"`
         
    }
    
    // ProductStockStatusItem ...
    type ProductStockStatusItem struct {

        
            Store StoreDetail  `json:"store"`
            Seller Seller  `json:"seller"`
            Quantity float64  `json:"quantity"`
            Company CompanyDetail  `json:"company"`
            ItemID float64  `json:"item_id"`
            Size string  `json:"size"`
            Price ProductStockPrice  `json:"price"`
            UID string  `json:"uid"`
            Identifier map[string]interface{}  `json:"identifier"`
         
    }
    
    // ProductStockStatusResponse ...
    type ProductStockStatusResponse struct {

        
            Items []ProductStockStatusItem  `json:"items"`
         
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
    
    // ProductVariantListingResponse ...
    type ProductVariantListingResponse struct {

        
            Items []ProductVariantItemResponse  `json:"items"`
            Key string  `json:"key"`
            Header string  `json:"header"`
            Total float64  `json:"total"`
            DisplayType string  `json:"display_type"`
         
    }
    
    // ProductListingDetail ...
    type ProductListingDetail struct {

        
            Medias []Media  `json:"medias"`
            Slug string  `json:"slug"`
            Highlights []string  `json:"highlights"`
            Tags []string  `json:"tags"`
            ItemType string  `json:"item_type"`
            Color string  `json:"color"`
            Price ProductListingPrice  `json:"price"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Brand ProductBrand  `json:"brand"`
            Sellable bool  `json:"sellable"`
            ItemCode string  `json:"item_code"`
            HasVariant bool  `json:"has_variant"`
            Moq ApplicationItemMOQ  `json:"moq"`
            RatingCount float64  `json:"rating_count"`
            TeaserTag string  `json:"teaser_tag"`
            Identifiers []string  `json:"identifiers"`
            Rating float64  `json:"rating"`
            Sizes []string  `json:"sizes"`
            Variants []ProductVariantListingResponse  `json:"variants"`
            IsDependent bool  `json:"is_dependent"`
            Attributes map[string]interface{}  `json:"attributes"`
            Discount string  `json:"discount"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            CategoryMap ProductCategoryMap  `json:"category_map"`
            Similars []string  `json:"similars"`
            Tryouts []string  `json:"tryouts"`
            ProductOnlineDate string  `json:"product_online_date"`
            Action ProductListingAction  `json:"action"`
            ShortDescription string  `json:"short_description"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Description string  `json:"description"`
            Seo ApplicationItemSEO  `json:"seo"`
            ImageNature string  `json:"image_nature"`
            Categories []ProductBrand  `json:"categories"`
            Name string  `json:"name"`
            Type string  `json:"type"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductFiltersValue ...
    type ProductFiltersValue struct {

        
            CurrencyCode string  `json:"currency_code"`
            DisplayFormat string  `json:"display_format"`
            IsSelected bool  `json:"is_selected"`
            Min float64  `json:"min"`
            SelectedMax float64  `json:"selected_max"`
            Value string  `json:"value"`
            QueryFormat string  `json:"query_format"`
            Max float64  `json:"max"`
            Count float64  `json:"count"`
            Display string  `json:"display"`
            CurrencySymbol string  `json:"currency_symbol"`
            SelectedMin float64  `json:"selected_min"`
         
    }
    
    // ProductFiltersKey ...
    type ProductFiltersKey struct {

        
            Display string  `json:"display"`
            Kind string  `json:"kind"`
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
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

        
            Discount string  `json:"discount"`
            Departments []string  `json:"departments"`
            Slug string  `json:"slug"`
            Action ProductListingAction  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            Logo Media  `json:"logo"`
            Description string  `json:"description"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // BrandListingResponse ...
    type BrandListingResponse struct {

        
            Items []BrandItem  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // BrandDetailResponse ...
    type BrandDetailResponse struct {

        
            Banners ImageUrls  `json:"banners"`
            Logo Media  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Description string  `json:"description"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // DepartmentIdentifier ...
    type DepartmentIdentifier struct {

        
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
         
    }
    
    // CategoryBanner ...
    type CategoryBanner struct {

        
            Portrait Media  `json:"portrait"`
            Landscape Media  `json:"landscape"`
         
    }
    
    // ThirdLevelChild ...
    type ThirdLevelChild struct {

        
            Slug string  `json:"slug"`
            Action ProductListingAction  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Childs []map[string]interface{}  `json:"childs"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // SecondLevelChild ...
    type SecondLevelChild struct {

        
            Slug string  `json:"slug"`
            Action ProductListingAction  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Childs []ThirdLevelChild  `json:"childs"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // Child ...
    type Child struct {

        
            Slug string  `json:"slug"`
            Action ProductListingAction  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Childs []SecondLevelChild  `json:"childs"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // CategoryItems ...
    type CategoryItems struct {

        
            Slug string  `json:"slug"`
            Action ProductListingAction  `json:"action"`
            Banners CategoryBanner  `json:"banners"`
            Childs []Child  `json:"childs"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
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

        
            Banners ImageUrls  `json:"banners"`
            Logo Media  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
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

        
            Action ProductListingAction  `json:"action"`
            Logo Media  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Display string  `json:"display"`
            Type string  `json:"type"`
         
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

        
            Slug string  `json:"slug"`
            AppID string  `json:"app_id"`
            Query []CollectionQuery  `json:"query"`
            Cron map[string]interface{}  `json:"cron"`
            Tag []string  `json:"tag"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Meta map[string]interface{}  `json:"meta"`
            Banners ImageUrls  `json:"banners"`
            Logo Media  `json:"logo"`
            Badge map[string]interface{}  `json:"badge"`
            IsActive bool  `json:"is_active"`
            SortOn string  `json:"sort_on"`
            AllowSort bool  `json:"allow_sort"`
            Priority float64  `json:"priority"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Action ProductListingAction  `json:"action"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Description string  `json:"description"`
            AllowFacets bool  `json:"allow_facets"`
            UID string  `json:"uid"`
            Name string  `json:"name"`
            Type string  `json:"type"`
         
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

        
            Slug string  `json:"slug"`
            AppID string  `json:"app_id"`
            Query []CollectionQuery  `json:"query"`
            Cron map[string]interface{}  `json:"cron"`
            Tag []string  `json:"tag"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Meta map[string]interface{}  `json:"meta"`
            Banners ImageUrls  `json:"banners"`
            Logo Media  `json:"logo"`
            Badge map[string]interface{}  `json:"badge"`
            IsActive bool  `json:"is_active"`
            SortOn string  `json:"sort_on"`
            AllowSort bool  `json:"allow_sort"`
            Priority float64  `json:"priority"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Description string  `json:"description"`
            AllowFacets bool  `json:"allow_facets"`
            Name string  `json:"name"`
            Type string  `json:"type"`
         
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

        
            Collections []float64  `json:"collections"`
            Brands []float64  `json:"brands"`
            Products []float64  `json:"products"`
         
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

        
            Name string  `json:"name"`
            Address string  `json:"address"`
            LatLong LatLong  `json:"lat_long"`
            State string  `json:"state"`
            Country string  `json:"country"`
            StoreCode string  `json:"store_code"`
            Pincode float64  `json:"pincode"`
            StoreEmail string  `json:"store_email"`
            City string  `json:"city"`
            UID float64  `json:"uid"`
         
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
            Logo string  `json:"logo"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // StoreAddressSerializer ...
    type StoreAddressSerializer struct {

        
            Address1 string  `json:"address1"`
            Latitude float64  `json:"latitude"`
            State string  `json:"state"`
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
            Longitude float64  `json:"longitude"`
            Address2 string  `json:"address2"`
            Pincode float64  `json:"pincode"`
            City string  `json:"city"`
         
    }
    
    // SellerPhoneNumber ...
    type SellerPhoneNumber struct {

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // CompanyStore ...
    type CompanyStore struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            BusinessType string  `json:"business_type"`
            CompanyType string  `json:"company_type"`
         
    }
    
    // StoreManagerSerializer ...
    type StoreManagerSerializer struct {

        
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Name string  `json:"name"`
            Email string  `json:"email"`
         
    }
    
    // AppStore ...
    type AppStore struct {

        
            Departments []StoreDepartments  `json:"departments"`
            Address StoreAddressSerializer  `json:"address"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Company CompanyStore  `json:"company"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
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

        
            Closing Time  `json:"closing"`
            Opening Time  `json:"opening"`
            Weekday string  `json:"weekday"`
            Open bool  `json:"open"`
         
    }
    
    // StoreDetails ...
    type StoreDetails struct {

        
            Departments []StoreDepartments  `json:"departments"`
            Address StoreAddressSerializer  `json:"address"`
            Timing []StoreTiming  `json:"timing"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Company CompanyStore  `json:"company"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Manager StoreManagerSerializer  `json:"manager"`
         
    }
    
    // ProductDetails ...
    type ProductDetails struct {

        
            Slug interface{}  `json:"slug"`
            Highlights []interface{}  `json:"highlights"`
            BrandUID float64  `json:"brand_uid"`
            ImageNature interface{}  `json:"image_nature"`
            IsSet bool  `json:"is_set"`
            Images []interface{}  `json:"images"`
            OutOfStock bool  `json:"out_of_stock"`
            Media []map[string]interface{}  `json:"media"`
            Identifier map[string]interface{}  `json:"identifier"`
            ItemCode interface{}  `json:"item_code"`
            HasVariant bool  `json:"has_variant"`
            RatingCount float64  `json:"rating_count"`
            Rating float64  `json:"rating"`
            CountryOfOrigin interface{}  `json:"country_of_origin"`
            HsnCode float64  `json:"hsn_code"`
            Attributes map[string]interface{}  `json:"attributes"`
            GroupedAttributes map[string]interface{}  `json:"grouped_attributes"`
            ShortDescription interface{}  `json:"short_description"`
            Description interface{}  `json:"description"`
            TemplateTag interface{}  `json:"template_tag"`
            Name interface{}  `json:"name"`
         
    }
    
    // Size ...
    type Size struct {

        
            Display interface{}  `json:"display"`
            Quantity float64  `json:"quantity"`
            Value interface{}  `json:"value"`
            IsAvailable bool  `json:"is_available"`
         
    }
    
    // ProductGroupPrice ...
    type ProductGroupPrice struct {

        
            Currency interface{}  `json:"currency"`
            MinMarked float64  `json:"min_marked"`
            MinEffective float64  `json:"min_effective"`
            MaxMarked float64  `json:"max_marked"`
            MaxEffective float64  `json:"max_effective"`
         
    }
    
    // ProductInGroup ...
    type ProductInGroup struct {

        
            AllowRemove bool  `json:"allow_remove"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            ProductUID float64  `json:"product_uid"`
            MinQuantity float64  `json:"min_quantity"`
            ProductDetails ProductDetails  `json:"product_details"`
            Sizes []Size  `json:"sizes"`
            Price ProductGroupPrice  `json:"price"`
            MaxQuantity float64  `json:"max_quantity"`
            AutoSelect bool  `json:"auto_select"`
         
    }
    
    // UserDetail ...
    type UserDetail struct {

        
            Username string  `json:"username"`
            SuperUser bool  `json:"super_user"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
         
    }
    
    // ProductGroupingModel ...
    type ProductGroupingModel struct {

        
            Slug interface{}  `json:"slug"`
            Products []ProductInGroup  `json:"products"`
            VerifiedOn string  `json:"verified_on"`
            CreatedBy UserDetail  `json:"created_by"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            ModifiedOn string  `json:"modified_on"`
            ModifiedBy UserDetail  `json:"modified_by"`
            CreatedOn string  `json:"created_on"`
            Choice interface{}  `json:"choice"`
            PageVisibility []interface{}  `json:"page_visibility"`
            VerifiedBy UserDetail  `json:"verified_by"`
            ID interface{}  `json:"_id"`
            CompanyID float64  `json:"company_id"`
            Logo string  `json:"logo"`
            IsActive bool  `json:"is_active"`
            Name interface{}  `json:"name"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // ProductBundle ...
    type ProductBundle struct {

        
            Items []ProductGroupingModel  `json:"items"`
         
    }
    
    // ArticleAssignmentV2 ...
    type ArticleAssignmentV2 struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // DetailsSchemaV2 ...
    type DetailsSchemaV2 struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            Key string  `json:"key"`
         
    }
    
    // MarketPlaceSttributesSchemaV2 ...
    type MarketPlaceSttributesSchemaV2 struct {

        
            Details []DetailsSchemaV2  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // SellerV2 ...
    type SellerV2 struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Count float64  `json:"count"`
         
    }
    
    // ReturnConfigSchemaV2 ...
    type ReturnConfigSchemaV2 struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
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

        
            SizeDistribution ProductSetDistributionV2  `json:"size_distribution"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // StoreV2 ...
    type StoreV2 struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Count float64  `json:"count"`
         
    }
    
    // ProductStockPriceV2 ...
    type ProductStockPriceV2 struct {

        
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            Currency string  `json:"currency"`
         
    }
    
    // StrategyWiseListingSchemaV2 ...
    type StrategyWiseListingSchemaV2 struct {

        
            Distance float64  `json:"distance"`
            Pincode float64  `json:"pincode"`
            Quantity float64  `json:"quantity"`
            Tat float64  `json:"tat"`
         
    }
    
    // SellerGroupAttributes ...
    type SellerGroupAttributes struct {

        
            Details []DetailsSchemaV2  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // ProductStockUnitPriceV2 ...
    type ProductStockUnitPriceV2 struct {

        
            CurrencyCode string  `json:"currency_code"`
            Price float64  `json:"price"`
            CurrencySymbol string  `json:"currency_symbol"`
            Unit string  `json:"unit"`
         
    }
    
    // ProductSizePriceResponseV2 ...
    type ProductSizePriceResponseV2 struct {

        
            ArticleAssignment ArticleAssignmentV2  `json:"article_assignment"`
            MarketplaceAttributes []MarketPlaceSttributesSchemaV2  `json:"marketplace_attributes"`
            Seller SellerV2  `json:"seller"`
            ItemType string  `json:"item_type"`
            ReturnConfig ReturnConfigSchemaV2  `json:"return_config"`
            SpecialBadge string  `json:"special_badge"`
            LongLat []float64  `json:"long_lat"`
            IsGift bool  `json:"is_gift"`
            SellerCount float64  `json:"seller_count"`
            Set ProductSetV2  `json:"set"`
            Store StoreV2  `json:"store"`
            IsCod bool  `json:"is_cod"`
            ArticleID string  `json:"article_id"`
            PricePerPiece ProductStockPriceV2  `json:"price_per_piece"`
            Pincode float64  `json:"pincode"`
            StrategyWiseListing []StrategyWiseListingSchemaV2  `json:"strategy_wise_listing"`
            Discount string  `json:"discount"`
            GroupedAttributes []SellerGroupAttributes  `json:"grouped_attributes"`
            PricePerUnit ProductStockUnitPriceV2  `json:"price_per_unit"`
            Quantity float64  `json:"quantity"`
            Price ProductStockPriceV2  `json:"price"`
         
    }
    
    // ProductSizeSellerFilterSchemaV2 ...
    type ProductSizeSellerFilterSchemaV2 struct {

        
            Value string  `json:"value"`
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
         
    }
    
    // ProductSizeSellersResponseV2 ...
    type ProductSizeSellersResponseV2 struct {

        
            SortOn []ProductSizeSellerFilterSchemaV2  `json:"sort_on"`
            Items []ProductSizePriceResponseV2  `json:"items"`
            Page Page  `json:"page"`
         
    }
    

    
    // ProductPrice ...
    type ProductPrice struct {

        
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
            AddOn float64  `json:"add_on"`
            Selling float64  `json:"selling"`
         
    }
    
    // ProductPriceInfo ...
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
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
    
    // BaseInfo ...
    type BaseInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductImage ...
    type ProductImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // Tags ...
    type Tags struct {

        
            Tags map[string]interface{}  `json:"tags"`
         
    }
    
    // CategoryInfo ...
    type CategoryInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // CartProduct ...
    type CartProduct struct {

        
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Action ProductAction  `json:"action"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Brand BaseInfo  `json:"brand"`
            Slug string  `json:"slug"`
            Type string  `json:"type"`
            Tags []string  `json:"tags"`
            Images []ProductImage  `json:"images"`
            TeaserTag Tags  `json:"teaser_tag"`
            Categories []CategoryInfo  `json:"categories"`
         
    }
    
    // CouponDetails ...
    type CouponDetails struct {

        
            Code string  `json:"code"`
            DiscountSingleQuantity float64  `json:"discount_single_quantity"`
            DiscountTotalQuantity float64  `json:"discount_total_quantity"`
         
    }
    
    // DiscountRulesApp ...
    type DiscountRulesApp struct {

        
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            RawOffer map[string]interface{}  `json:"raw_offer"`
            MatchedBuyRules []string  `json:"matched_buy_rules"`
            Offer map[string]interface{}  `json:"offer"`
         
    }
    
    // BuyRules ...
    type BuyRules struct {

        
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
            CartConditions map[string]interface{}  `json:"cart_conditions"`
         
    }
    
    // FreeGiftItem ...
    type FreeGiftItem struct {

        
            ItemBrandName string  `json:"item_brand_name"`
            ItemName string  `json:"item_name"`
            ItemImagesURL []string  `json:"item_images_url"`
            ItemSlug string  `json:"item_slug"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // AppliedFreeArticles ...
    type AppliedFreeArticles struct {

        
            Quantity float64  `json:"quantity"`
            ArticleID string  `json:"article_id"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
         
    }
    
    // AppliedPromotion ...
    type AppliedPromotion struct {

        
            PromotionGroup string  `json:"promotion_group"`
            Amount float64  `json:"amount"`
            MrpPromotion bool  `json:"mrp_promotion"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            PromoID string  `json:"promo_id"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromotionName string  `json:"promotion_name"`
            OfferText string  `json:"offer_text"`
            PromotionType string  `json:"promotion_type"`
            BuyRules []BuyRules  `json:"buy_rules"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
         
    }
    
    // BasePrice ...
    type BasePrice struct {

        
            Marked float64  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
            Effective float64  `json:"effective"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // ArticlePriceInfo ...
    type ArticlePriceInfo struct {

        
            Converted BasePrice  `json:"converted"`
            Base BasePrice  `json:"base"`
         
    }
    
    // ProductArticle ...
    type ProductArticle struct {

        
            Price ArticlePriceInfo  `json:"price"`
            Size string  `json:"size"`
            UID string  `json:"uid"`
            Identifier map[string]interface{}  `json:"identifier"`
            Quantity float64  `json:"quantity"`
            IsGiftVisible bool  `json:"is_gift_visible"`
            Type string  `json:"type"`
            Store BaseInfo  `json:"store"`
            ProductGroupTags []string  `json:"product_group_tags"`
            GiftCard map[string]interface{}  `json:"gift_card"`
            IsGift bool  `json:"is_gift"`
            SellerIdentifier string  `json:"seller_identifier"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Seller BaseInfo  `json:"seller"`
         
    }
    
    // PromoMeta ...
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // ProductAvailability ...
    type ProductAvailability struct {

        
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            Deliverable bool  `json:"deliverable"`
            OutOfStock bool  `json:"out_of_stock"`
            Sizes []string  `json:"sizes"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // CartProductIdentifer ...
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            IsSet bool  `json:"is_set"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Price ProductPriceInfo  `json:"price"`
            Product CartProduct  `json:"product"`
            Quantity float64  `json:"quantity"`
            Discount string  `json:"discount"`
            Coupon CouponDetails  `json:"coupon"`
            Key string  `json:"key"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            CouponMessage string  `json:"coupon_message"`
            Message string  `json:"message"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Article ProductArticle  `json:"article"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Availability ProductAvailability  `json:"availability"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
         
    }
    
    // RawBreakup ...
    type RawBreakup struct {

        
            CodCharge float64  `json:"cod_charge"`
            GstCharges float64  `json:"gst_charges"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Discount float64  `json:"discount"`
            Coupon float64  `json:"coupon"`
            YouSaved float64  `json:"you_saved"`
            MrpTotal float64  `json:"mrp_total"`
            Total float64  `json:"total"`
            ConvenienceFee float64  `json:"convenience_fee"`
            FyndCash float64  `json:"fynd_cash"`
            Vog float64  `json:"vog"`
            Subtotal float64  `json:"subtotal"`
         
    }
    
    // CouponBreakup ...
    type CouponBreakup struct {

        
            UID string  `json:"uid"`
            CouponType string  `json:"coupon_type"`
            Type string  `json:"type"`
            Value float64  `json:"value"`
            CouponValue float64  `json:"coupon_value"`
            Title string  `json:"title"`
            SubTitle string  `json:"sub_title"`
            Message string  `json:"message"`
            Description string  `json:"description"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            IsApplied bool  `json:"is_applied"`
            Code string  `json:"code"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
         
    }
    
    // DisplayBreakup ...
    type DisplayBreakup struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            Display string  `json:"display"`
            CurrencyCode string  `json:"currency_code"`
            Key string  `json:"key"`
            Value float64  `json:"value"`
            Message []string  `json:"message"`
         
    }
    
    // LoyaltyPoints ...
    type LoyaltyPoints struct {

        
            Description string  `json:"description"`
            Applicable float64  `json:"applicable"`
            IsApplied bool  `json:"is_applied"`
            Total float64  `json:"total"`
         
    }
    
    // CartBreakup ...
    type CartBreakup struct {

        
            Raw RawBreakup  `json:"raw"`
            Coupon CouponBreakup  `json:"coupon"`
            Display []DisplayBreakup  `json:"display"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
         
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
    
    // CartDetailResponse ...
    type CartDetailResponse struct {

        
            Items []CartProductInfo  `json:"items"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Comment string  `json:"comment"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CheckoutMode string  `json:"checkout_mode"`
            LastModified string  `json:"last_modified"`
            BuyNow bool  `json:"buy_now"`
            Message string  `json:"message"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CouponText string  `json:"coupon_text"`
            ID string  `json:"id"`
            Gstin string  `json:"gstin"`
            Currency CartCurrency  `json:"currency"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            Display string  `json:"display"`
            StoreID float64  `json:"store_id"`
            Quantity float64  `json:"quantity"`
            Pos bool  `json:"pos"`
            ItemSize string  `json:"item_size"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            ProductGroupTags []string  `json:"product_group_tags"`
            ArticleID string  `json:"article_id"`
            IsGift bool  `json:"is_gift"`
            GiftMessage string  `json:"gift_message"`
            ItemID float64  `json:"item_id"`
            SellerID float64  `json:"seller_id"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
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

        
            ItemIndex float64  `json:"item_index"`
            Quantity float64  `json:"quantity"`
            ItemSize string  `json:"item_size"`
            ArticleID string  `json:"article_id"`
            IsGift bool  `json:"is_gift"`
            GiftMessage string  `json:"gift_message"`
            ItemID float64  `json:"item_id"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
         
    }
    
    // UpdateCartRequest ...
    type UpdateCartRequest struct {

        
            Operation string  `json:"operation"`
            Items []UpdateProductCart  `json:"items"`
         
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

        
            IsApplicable bool  `json:"is_applicable"`
            CouponType string  `json:"coupon_type"`
            ExpiresOn string  `json:"expires_on"`
            Title string  `json:"title"`
            SubTitle string  `json:"sub_title"`
            Message string  `json:"message"`
            Description string  `json:"description"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            IsApplied bool  `json:"is_applied"`
            CouponValue float64  `json:"coupon_value"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            CouponCode string  `json:"coupon_code"`
         
    }
    
    // PageCoupon ...
    type PageCoupon struct {

        
            TotalItemCount float64  `json:"total_item_count"`
            HasNext bool  `json:"has_next"`
            HasPrevious bool  `json:"has_previous"`
            Total float64  `json:"total"`
            Current float64  `json:"current"`
         
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
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
            BulkEffective float64  `json:"bulk_effective"`
         
    }
    
    // OfferItem ...
    type OfferItem struct {

        
            Margin float64  `json:"margin"`
            Price OfferPrice  `json:"price"`
            Quantity float64  `json:"quantity"`
            Best bool  `json:"best"`
            Type string  `json:"type"`
            Total float64  `json:"total"`
            AutoApplied bool  `json:"auto_applied"`
         
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

        
            AddressType string  `json:"address_type"`
            CheckoutMode string  `json:"checkout_mode"`
            Area string  `json:"area"`
            Email string  `json:"email"`
            UserID string  `json:"user_id"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            Phone string  `json:"phone"`
            ID string  `json:"id"`
            GeoLocation GeoLocation  `json:"geo_location"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Address string  `json:"address"`
            Landmark string  `json:"landmark"`
            CountryCode string  `json:"country_code"`
            State string  `json:"state"`
            Tags []string  `json:"tags"`
            IsActive bool  `json:"is_active"`
            IsDefaultAddress bool  `json:"is_default_address"`
            Country string  `json:"country"`
            City string  `json:"city"`
            Name string  `json:"name"`
            AreaCode string  `json:"area_code"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // GetAddressesResponse ...
    type GetAddressesResponse struct {

        
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

        
            Success bool  `json:"success"`
            IsUpdated bool  `json:"is_updated"`
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
            ID string  `json:"id"`
            CartID string  `json:"cart_id"`
         
    }
    
    // UpdateCartPaymentRequest ...
    type UpdateCartPaymentRequest struct {

        
            PaymentMode string  `json:"payment_mode"`
            AddressID string  `json:"address_id"`
            AggregatorName string  `json:"aggregator_name"`
            ID string  `json:"id"`
            MerchantCode string  `json:"merchant_code"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CouponValidity ...
    type CouponValidity struct {

        
            Valid bool  `json:"valid"`
            Discount float64  `json:"discount"`
            DisplayMessageEn string  `json:"display_message_en"`
            Code string  `json:"code"`
            Title string  `json:"title"`
         
    }
    
    // PaymentCouponValidate ...
    type PaymentCouponValidate struct {

        
            CouponValidity CouponValidity  `json:"coupon_validity"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // ShipmentResponse ...
    type ShipmentResponse struct {

        
            DpID string  `json:"dp_id"`
            Shipments float64  `json:"shipments"`
            BoxType string  `json:"box_type"`
            Items []CartProductInfo  `json:"items"`
            OrderType string  `json:"order_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
            ShipmentType string  `json:"shipment_type"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Promise ShipmentPromise  `json:"promise"`
            FulfillmentType string  `json:"fulfillment_type"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            Comment string  `json:"comment"`
            Error bool  `json:"error"`
            UID string  `json:"uid"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CheckoutMode string  `json:"checkout_mode"`
            BuyNow bool  `json:"buy_now"`
            Message string  `json:"message"`
            Currency CartCurrency  `json:"currency"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Shipments []ShipmentResponse  `json:"shipments"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            LastModified string  `json:"last_modified"`
            CartID float64  `json:"cart_id"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CouponText string  `json:"coupon_text"`
            ID string  `json:"id"`
            Gstin string  `json:"gstin"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // StaffCheckout ...
    type StaffCheckout struct {

        
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            User string  `json:"user"`
            ID string  `json:"_id"`
         
    }
    
    // CartCheckoutDetailRequest ...
    type CartCheckoutDetailRequest struct {

        
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            OrderingStore float64  `json:"ordering_store"`
            PaymentMode string  `json:"payment_mode"`
            AddressID string  `json:"address_id"`
            Staff StaffCheckout  `json:"staff"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            Meta map[string]interface{}  `json:"meta"`
            CallbackURL string  `json:"callback_url"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            MerchantCode string  `json:"merchant_code"`
            Aggregator string  `json:"aggregator"`
            BillingAddressID string  `json:"billing_address_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CheckCart ...
    type CheckCart struct {

        
            Success bool  `json:"success"`
            Comment string  `json:"comment"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryCharges float64  `json:"delivery_charges"`
            BuyNow bool  `json:"buy_now"`
            Message string  `json:"message"`
            Items []CartProductInfo  `json:"items"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            CodMessage string  `json:"cod_message"`
            StoreCode string  `json:"store_code"`
            ErrorMessage string  `json:"error_message"`
            ID string  `json:"id"`
            Gstin string  `json:"gstin"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            IsValid bool  `json:"is_valid"`
            UserType string  `json:"user_type"`
            UID string  `json:"uid"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            OrderID string  `json:"order_id"`
            Currency CartCurrency  `json:"currency"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CodAvailable bool  `json:"cod_available"`
            CodCharges float64  `json:"cod_charges"`
            LastModified string  `json:"last_modified"`
            CouponText string  `json:"coupon_text"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CartID float64  `json:"cart_id"`
         
    }
    
    // CartCheckoutResponse ...
    type CartCheckoutResponse struct {

        
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Success bool  `json:"success"`
            AppInterceptURL string  `json:"app_intercept_url"`
            Data map[string]interface{}  `json:"data"`
            CallbackURL string  `json:"callback_url"`
            OrderID string  `json:"order_id"`
            Message string  `json:"message"`
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

        
            Source map[string]interface{}  `json:"source"`
            User map[string]interface{}  `json:"user"`
            Token string  `json:"token"`
            Meta map[string]interface{}  `json:"meta"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            Comment string  `json:"comment"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            CheckoutMode string  `json:"checkout_mode"`
            BuyNow bool  `json:"buy_now"`
            Message string  `json:"message"`
            Items []CartProductInfo  `json:"items"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            ID string  `json:"id"`
            Gstin string  `json:"gstin"`
            IsValid bool  `json:"is_valid"`
            UID string  `json:"uid"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Currency CartCurrency  `json:"currency"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            LastModified string  `json:"last_modified"`
            CouponText string  `json:"coupon_text"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CartID float64  `json:"cart_id"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Error string  `json:"error"`
            Cart SharedCart  `json:"cart"`
         
    }
    
    // FreeGiftItems ...
    type FreeGiftItems struct {

        
            ItemBrandName string  `json:"item_brand_name"`
            ItemName string  `json:"item_name"`
            ItemImagesURL []string  `json:"item_images_url"`
            ItemSlug string  `json:"item_slug"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // PromotionOffer ...
    type PromotionOffer struct {

        
            PromotionGroup string  `json:"promotion_group"`
            DiscountRules []map[string]interface{}  `json:"discount_rules"`
            Description string  `json:"description"`
            ID string  `json:"id"`
            FreeGiftItems []FreeGiftItems  `json:"free_gift_items"`
            ValidTill string  `json:"valid_till"`
            OfferText string  `json:"offer_text"`
            BuyRules map[string]interface{}  `json:"buy_rules"`
         
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

        
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
            OfferPrice float64  `json:"offer_price"`
         
    }
    
    // LadderOfferItem ...
    type LadderOfferItem struct {

        
            MinQuantity float64  `json:"min_quantity"`
            Margin float64  `json:"margin"`
            Price LadderPrice  `json:"price"`
            Type string  `json:"type"`
            MaxQuantity float64  `json:"max_quantity"`
         
    }
    
    // LadderPriceOffer ...
    type LadderPriceOffer struct {

        
            PromotionGroup string  `json:"promotion_group"`
            DiscountRules []map[string]interface{}  `json:"discount_rules"`
            CalculateOn string  `json:"calculate_on"`
            Description string  `json:"description"`
            OfferPrices []LadderOfferItem  `json:"offer_prices"`
            ID string  `json:"id"`
            FreeGiftItems []FreeGiftItems  `json:"free_gift_items"`
            ValidTill string  `json:"valid_till"`
            OfferText string  `json:"offer_text"`
            BuyRules map[string]interface{}  `json:"buy_rules"`
         
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
            Debug Debug  `json:"debug"`
            HasOldPasswordHash bool  `json:"has_old_password_hash"`
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
            SubNavigation []SubNavigationReference  `json:"sub_navigation"`
         
    }
    
    // SubNavigationReference ...
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

        
            MerchantID string  `json:"merchant_id"`
            Secret string  `json:"secret"`
            VerifyAPI string  `json:"verify_api"`
            Key string  `json:"key"`
            UserID string  `json:"user_id"`
            MerchantKey string  `json:"merchant_key"`
            Pin string  `json:"pin"`
            API string  `json:"api"`
            Sdk bool  `json:"sdk"`
            ConfigType string  `json:"config_type"`
         
    }
    
    // AggregatorsConfigDetailResponse ...
    type AggregatorsConfigDetailResponse struct {

        
            Payumoney AggregatorConfigDetail  `json:"payumoney"`
            Ccavenue AggregatorConfigDetail  `json:"ccavenue"`
            Simpl AggregatorConfigDetail  `json:"simpl"`
            Juspay AggregatorConfigDetail  `json:"juspay"`
            Success bool  `json:"success"`
            Mswipe AggregatorConfigDetail  `json:"mswipe"`
            Razorpay AggregatorConfigDetail  `json:"razorpay"`
            Stripe AggregatorConfigDetail  `json:"stripe"`
            Rupifi AggregatorConfigDetail  `json:"rupifi"`
            Env string  `json:"env"`
         
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

        
            NameOnCard string  `json:"name_on_card"`
            Nickname string  `json:"nickname"`
            Refresh bool  `json:"refresh"`
            CardID string  `json:"card_id"`
         
    }
    
    // AttachCardsResponse ...
    type AttachCardsResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // CardPaymentGateway ...
    type CardPaymentGateway struct {

        
            CustomerID string  `json:"customer_id"`
            Aggregator string  `json:"aggregator"`
            API string  `json:"api"`
         
    }
    
    // ActiveCardPaymentGatewayResponse ...
    type ActiveCardPaymentGatewayResponse struct {

        
            Cards CardPaymentGateway  `json:"cards"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // Card ...
    type Card struct {

        
            CardType string  `json:"card_type"`
            ExpYear float64  `json:"exp_year"`
            AggregatorName string  `json:"aggregator_name"`
            CardIssuer string  `json:"card_issuer"`
            CardName string  `json:"card_name"`
            ExpMonth float64  `json:"exp_month"`
            CardFingerprint string  `json:"card_fingerprint"`
            CardBrand string  `json:"card_brand"`
            CardNumber string  `json:"card_number"`
            CardID string  `json:"card_id"`
            Nickname string  `json:"nickname"`
            CardReference string  `json:"card_reference"`
            Expired bool  `json:"expired"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardIsin string  `json:"card_isin"`
            CardToken string  `json:"card_token"`
            CardBrandImage string  `json:"card_brand_image"`
         
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

        
            BillingAddress map[string]interface{}  `json:"billing_address"`
            Aggregator string  `json:"aggregator"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            PhoneNumber string  `json:"phone_number"`
            OrderItems []map[string]interface{}  `json:"order_items"`
            MerchantParams map[string]interface{}  `json:"merchant_params"`
            Payload string  `json:"payload"`
            TransactionAmountInPaise float64  `json:"transaction_amount_in_paise"`
         
    }
    
    // ValidateCustomerResponse ...
    type ValidateCustomerResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // ChargeCustomerRequest ...
    type ChargeCustomerRequest struct {

        
            Aggregator string  `json:"aggregator"`
            Verified bool  `json:"verified"`
            TransactionToken string  `json:"transaction_token"`
            Amount float64  `json:"amount"`
            OrderID string  `json:"order_id"`
         
    }
    
    // ChargeCustomerResponse ...
    type ChargeCustomerResponse struct {

        
            Aggregator string  `json:"aggregator"`
            CartID string  `json:"cart_id"`
            Success bool  `json:"success"`
            Status string  `json:"status"`
            Message string  `json:"message"`
            DeliveryAddressID string  `json:"delivery_address_id"`
            OrderID string  `json:"order_id"`
         
    }
    
    // PaymentInitializationRequest ...
    type PaymentInitializationRequest struct {

        
            Aggregator string  `json:"aggregator"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Currency string  `json:"currency"`
            CustomerID string  `json:"customer_id"`
            Timeout float64  `json:"timeout"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Method string  `json:"method"`
            Vpa string  `json:"vpa"`
            Email string  `json:"email"`
            Amount float64  `json:"amount"`
            Contact string  `json:"contact"`
            OrderID string  `json:"order_id"`
         
    }
    
    // PaymentInitializationResponse ...
    type PaymentInitializationResponse struct {

        
            UpiPollURL string  `json:"upi_poll_url"`
            Aggregator string  `json:"aggregator"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Currency string  `json:"currency"`
            CustomerID string  `json:"customer_id"`
            PollingURL string  `json:"polling_url"`
            Timeout float64  `json:"timeout"`
            BqrImage string  `json:"bqr_image"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Success bool  `json:"success"`
            Status string  `json:"status"`
            VirtualID string  `json:"virtual_id"`
            Method string  `json:"method"`
            Vpa string  `json:"vpa"`
            Amount float64  `json:"amount"`
         
    }
    
    // PaymentStatusUpdateRequest ...
    type PaymentStatusUpdateRequest struct {

        
            Aggregator string  `json:"aggregator"`
            Currency string  `json:"currency"`
            CustomerID string  `json:"customer_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Status string  `json:"status"`
            Contact string  `json:"contact"`
            Method string  `json:"method"`
            Vpa string  `json:"vpa"`
            Email string  `json:"email"`
            Amount float64  `json:"amount"`
            OrderID string  `json:"order_id"`
         
    }
    
    // PaymentStatusUpdateResponse ...
    type PaymentStatusUpdateResponse struct {

        
            AggregatorName string  `json:"aggregator_name"`
            RedirectURL string  `json:"redirect_url"`
            Success bool  `json:"success"`
            Status string  `json:"status"`
            Retry bool  `json:"retry"`
         
    }
    
    // PaymentModeLogo ...
    type PaymentModeLogo struct {

        
            Large string  `json:"large"`
            Small string  `json:"small"`
         
    }
    
    // IntentAppErrorList ...
    type IntentAppErrorList struct {

        
            PackageName string  `json:"package_name"`
            Code string  `json:"code"`
         
    }
    
    // IntentApp ...
    type IntentApp struct {

        
            Logos PaymentModeLogo  `json:"logos"`
            DisplayName string  `json:"display_name"`
            PackageName string  `json:"package_name"`
            Code string  `json:"code"`
         
    }
    
    // PaymentModeList ...
    type PaymentModeList struct {

        
            CardType string  `json:"card_type"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            AggregatorName string  `json:"aggregator_name"`
            Nickname string  `json:"nickname"`
            CardIsin string  `json:"card_isin"`
            CardToken string  `json:"card_token"`
            MerchantCode string  `json:"merchant_code"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            CardBrand string  `json:"card_brand"`
            CardNumber string  `json:"card_number"`
            CardID string  `json:"card_id"`
            IntentFlow bool  `json:"intent_flow"`
            Expired bool  `json:"expired"`
            FyndVpa string  `json:"fynd_vpa"`
            CardIssuer string  `json:"card_issuer"`
            Timeout float64  `json:"timeout"`
            CardFingerprint string  `json:"card_fingerprint"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            DisplayPriority float64  `json:"display_priority"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            RemainingLimit float64  `json:"remaining_limit"`
            CodLimit float64  `json:"cod_limit"`
            RetryCount float64  `json:"retry_count"`
            ExpYear float64  `json:"exp_year"`
            CardBrandImage string  `json:"card_brand_image"`
            DisplayName string  `json:"display_name"`
            CardName string  `json:"card_name"`
            ExpMonth float64  `json:"exp_month"`
            CardReference string  `json:"card_reference"`
            Code string  `json:"code"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            Name string  `json:"name"`
            IntentApp []IntentApp  `json:"intent_app"`
         
    }
    
    // RootPaymentMode ...
    type RootPaymentMode struct {

        
            DisplayName string  `json:"display_name"`
            AggregatorName string  `json:"aggregator_name"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            SaveCard bool  `json:"save_card"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            DisplayPriority float64  `json:"display_priority"`
            List []PaymentModeList  `json:"list"`
            Name string  `json:"name"`
            AddCardEnabled bool  `json:"add_card_enabled"`
         
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

        
            Epaylater AggregatorRoute  `json:"epaylater"`
            Ccavenue AggregatorRoute  `json:"ccavenue"`
            BqrRazorpay AggregatorRoute  `json:"bqr_razorpay"`
            Simpl AggregatorRoute  `json:"simpl"`
            Juspay AggregatorRoute  `json:"juspay"`
            Rupifi AggregatorRoute  `json:"rupifi"`
            Razorpay AggregatorRoute  `json:"razorpay"`
            Mswipe AggregatorRoute  `json:"mswipe"`
            Stripe AggregatorRoute  `json:"stripe"`
            UpiRazorpay AggregatorRoute  `json:"upi_razorpay"`
            Jiopay AggregatorRoute  `json:"jiopay"`
            Payubiz AggregatorRoute  `json:"payubiz"`
            Fynd AggregatorRoute  `json:"fynd"`
         
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

        
            DisplayName string  `json:"display_name"`
            LogoLarge string  `json:"logo_large"`
            ID float64  `json:"id"`
            Name string  `json:"name"`
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

        
            DelightsUserName string  `json:"delights_user_name"`
            IfscCode string  `json:"ifsc_code"`
            BeneficiaryID string  `json:"beneficiary_id"`
            BranchName string  `json:"branch_name"`
            CreatedOn string  `json:"created_on"`
            Subtitle string  `json:"subtitle"`
            ID float64  `json:"id"`
            TransferMode string  `json:"transfer_mode"`
            Email string  `json:"email"`
            Comment string  `json:"comment"`
            Address string  `json:"address"`
            AccountNo string  `json:"account_no"`
            ModifiedOn string  `json:"modified_on"`
            IsActive bool  `json:"is_active"`
            DisplayName string  `json:"display_name"`
            Mobile string  `json:"mobile"`
            BankName string  `json:"bank_name"`
            AccountHolder string  `json:"account_holder"`
            Title string  `json:"title"`
         
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

        
            BranchName string  `json:"branch_name"`
            BankName string  `json:"bank_name"`
            Success bool  `json:"success"`
         
    }
    
    // ErrorCodeDescription ...
    type ErrorCodeDescription struct {

        
            Description string  `json:"description"`
            Success bool  `json:"success"`
            Code string  `json:"code"`
         
    }
    
    // AddBeneficiaryViaOtpVerificationRequest ...
    type AddBeneficiaryViaOtpVerificationRequest struct {

        
            HashKey string  `json:"hash_key"`
            Otp string  `json:"otp"`
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
            BankName string  `json:"bank_name"`
            Address string  `json:"address"`
            AccountHolder string  `json:"account_holder"`
            IfscCode string  `json:"ifsc_code"`
            Wallet string  `json:"wallet"`
            Vpa string  `json:"vpa"`
            AccountNo string  `json:"account_no"`
            Email string  `json:"email"`
            Comment string  `json:"comment"`
            BranchName string  `json:"branch_name"`
         
    }
    
    // AddBeneficiaryDetailsRequest ...
    type AddBeneficiaryDetailsRequest struct {

        
            Delights bool  `json:"delights"`
            Details BeneficiaryModeDetails  `json:"details"`
            RequestID string  `json:"request_id"`
            ShipmentID string  `json:"shipment_id"`
            TransferMode string  `json:"transfer_mode"`
            Otp string  `json:"otp"`
            OrderID string  `json:"order_id"`
         
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
            AccountHolder string  `json:"account_holder"`
            IfscCode string  `json:"ifsc_code"`
            AccountNo string  `json:"account_no"`
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

        
            IsVerifiedFlag string  `json:"is_verified_flag"`
            Success bool  `json:"success"`
            RequestID string  `json:"request_id"`
         
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

        
            ExternalOrderID string  `json:"external_order_id"`
            PollingTimeout float64  `json:"polling_timeout"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            MerchantName string  `json:"merchant_name"`
            PaymentLinkURL string  `json:"payment_link_url"`
            Message string  `json:"message"`
            PaymentLinkCurrentStatus string  `json:"payment_link_current_status"`
            Amount float64  `json:"amount"`
         
    }
    
    // ErrorDescription ...
    type ErrorDescription struct {

        
            PaymentTransactionID string  `json:"payment_transaction_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            MerchantName string  `json:"merchant_name"`
            Expired bool  `json:"expired"`
            Cancelled bool  `json:"cancelled"`
            InvalidID bool  `json:"invalid_id"`
            Msg string  `json:"msg"`
            Amount float64  `json:"amount"`
         
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

        
            Description string  `json:"description"`
            ExternalOrderID string  `json:"external_order_id"`
            MobileNumber string  `json:"mobile_number"`
            Meta CreatePaymentLinkMeta  `json:"meta"`
            Email string  `json:"email"`
            Amount float64  `json:"amount"`
         
    }
    
    // CreatePaymentLinkResponse ...
    type CreatePaymentLinkResponse struct {

        
            PollingTimeout float64  `json:"polling_timeout"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            PaymentLinkURL string  `json:"payment_link_url"`
            Message string  `json:"message"`
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // CancelOrResendPaymentLinkRequest ...
    type CancelOrResendPaymentLinkRequest struct {

        
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // ResendPaymentLinkResponse ...
    type ResendPaymentLinkResponse struct {

        
            PollingTimeout float64  `json:"polling_timeout"`
            Message string  `json:"message"`
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

        
            AggregatorName string  `json:"aggregator_name"`
            RedirectURL string  `json:"redirect_url"`
            HttpStatus float64  `json:"http_status"`
            Success bool  `json:"success"`
            Status string  `json:"status"`
            StatusCode float64  `json:"status_code"`
            Message string  `json:"message"`
            Amount float64  `json:"amount"`
            PaymentLinkID string  `json:"payment_link_id"`
            OrderID string  `json:"order_id"`
         
    }
    
    // PaymentMethodsMeta ...
    type PaymentMethodsMeta struct {

        
            MerchantCode string  `json:"merchant_code"`
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentGateway string  `json:"payment_gateway"`
         
    }
    
    // CreateOrderUserPaymentMethods ...
    type CreateOrderUserPaymentMethods struct {

        
            Name string  `json:"name"`
            Mode string  `json:"mode"`
            Meta PaymentMethodsMeta  `json:"meta"`
         
    }
    
    // CreateOrderUserRequest ...
    type CreateOrderUserRequest struct {

        
            Currency string  `json:"currency"`
            SuccessCallbackURL string  `json:"success_callback_url"`
            PaymentMethods CreateOrderUserPaymentMethods  `json:"payment_methods"`
            Meta map[string]interface{}  `json:"meta"`
            PaymentLinkID string  `json:"payment_link_id"`
            FailureCallbackURL string  `json:"failure_callback_url"`
         
    }
    
    // CreateOrderUserData ...
    type CreateOrderUserData struct {

        
            Aggregator string  `json:"aggregator"`
            Currency string  `json:"currency"`
            CustomerID string  `json:"customer_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            CallbackURL string  `json:"callback_url"`
            Method string  `json:"method"`
            Email string  `json:"email"`
            Amount float64  `json:"amount"`
            Contact string  `json:"contact"`
            OrderID string  `json:"order_id"`
         
    }
    
    // CreateOrderUserResponse ...
    type CreateOrderUserResponse struct {

        
            Data CreateOrderUserData  `json:"data"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            CallbackURL string  `json:"callback_url"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Message string  `json:"message"`
            OrderID string  `json:"order_id"`
         
    }
    
    // BalanceDetails ...
    type BalanceDetails struct {

        
            Value float64  `json:"value"`
            FormattedValue string  `json:"formatted_value"`
            Currency string  `json:"currency"`
         
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

        
            Data CreditSummary  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // RedirectURL ...
    type RedirectURL struct {

        
            Status bool  `json:"status"`
            SignupURL string  `json:"signup_url"`
         
    }
    
    // RedirectToAggregatorResponse ...
    type RedirectToAggregatorResponse struct {

        
            Data RedirectURL  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // CreditDetail ...
    type CreditDetail struct {

        
            Status bool  `json:"status"`
            IsRegistered bool  `json:"is_registered"`
            SignupURL string  `json:"signup_url"`
         
    }
    
    // CheckCreditResponse ...
    type CheckCreditResponse struct {

        
            Data CreditDetail  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // KYCAddress ...
    type KYCAddress struct {

        
            Addressline1 string  `json:"addressline1"`
            City string  `json:"city"`
            Addressline2 string  `json:"addressline2"`
            OwnershipType string  `json:"ownership_type"`
            State string  `json:"state"`
            Pincode string  `json:"pincode"`
            LandMark string  `json:"land_mark"`
         
    }
    
    // UserPersonalInfoInDetails ...
    type UserPersonalInfoInDetails struct {

        
            MiddleName string  `json:"middle_name"`
            VoterID string  `json:"voter_id"`
            Dob string  `json:"dob"`
            Phone string  `json:"phone"`
            FathersName string  `json:"fathers_name"`
            Passport string  `json:"passport"`
            EmailVerified bool  `json:"email_verified"`
            MothersName string  `json:"mothers_name"`
            FirstName string  `json:"first_name"`
            Pan string  `json:"pan"`
            DrivingLicense string  `json:"driving_license"`
            LastName string  `json:"last_name"`
            AddressAsPerID KYCAddress  `json:"address_as_per_id"`
            Gender string  `json:"gender"`
            Email string  `json:"email"`
            MobileVerified bool  `json:"mobile_verified"`
         
    }
    
    // MarketplaceInfo ...
    type MarketplaceInfo struct {

        
            Name string  `json:"name"`
            MembershipID string  `json:"membership_id"`
            DateOfJoining string  `json:"date_of_joining"`
         
    }
    
    // DeviceDetails ...
    type DeviceDetails struct {

        
            DeviceModel string  `json:"device_model"`
            Os string  `json:"os"`
            DeviceMake string  `json:"device_make"`
            IdentifierType string  `json:"identifier_type"`
            DeviceType string  `json:"device_type"`
            OsVersion string  `json:"os_version"`
            IdentificationNumber string  `json:"identification_number"`
         
    }
    
    // BusinessDetails ...
    type BusinessDetails struct {

        
            BusinessOwnershipType string  `json:"business_ownership_type"`
            EntityType string  `json:"entity_type"`
            Vintage string  `json:"vintage"`
            ShopAndEstablishment map[string]interface{}  `json:"shop_and_establishment"`
            Fda string  `json:"fda"`
            Address KYCAddress  `json:"address"`
            Pan string  `json:"pan"`
            BusinessType string  `json:"business_type"`
            Gstin string  `json:"gstin"`
            Name string  `json:"name"`
            Fssai string  `json:"fssai"`
         
    }
    
    // CustomerOnboardingRequest ...
    type CustomerOnboardingRequest struct {

        
            Aggregator string  `json:"aggregator"`
            Mcc string  `json:"mcc"`
            PersonalInfo UserPersonalInfoInDetails  `json:"personal_info"`
            Source string  `json:"source"`
            MarketplaceInfo MarketplaceInfo  `json:"marketplace_info"`
            Device DeviceDetails  `json:"device"`
            BusinessInfo BusinessDetails  `json:"business_info"`
         
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
    
    // BagsForReorderArticleAssignment ...
    type BagsForReorderArticleAssignment struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // BagsForReorder ...
    type BagsForReorder struct {

        
            ArticleAssignment BagsForReorderArticleAssignment  `json:"article_assignment"`
            StoreID float64  `json:"store_id"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
            ItemSize string  `json:"item_size"`
            SellerID float64  `json:"seller_id"`
         
    }
    
    // UserInfo ...
    type UserInfo struct {

        
            Name string  `json:"name"`
            Email string  `json:"email"`
            Mobile string  `json:"mobile"`
            Gender string  `json:"gender"`
         
    }
    
    // ShipmentPayment ...
    type ShipmentPayment struct {

        
            Status string  `json:"status"`
            DisplayName string  `json:"display_name"`
            Logo string  `json:"logo"`
            Mop string  `json:"mop"`
            Mode string  `json:"mode"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // Invoice ...
    type Invoice struct {

        
            LabelURL string  `json:"label_url"`
            InvoiceURL string  `json:"invoice_url"`
            UpdatedDate string  `json:"updated_date"`
         
    }
    
    // FulfillingCompany ...
    type FulfillingCompany struct {

        
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

        
            Max string  `json:"max"`
            Min string  `json:"min"`
         
    }
    
    // Promise ...
    type Promise struct {

        
            ShowPromise bool  `json:"show_promise"`
            Timestamp TimeStampData  `json:"timestamp"`
         
    }
    
    // ShipmentTotalDetails ...
    type ShipmentTotalDetails struct {

        
            TotalPrice float64  `json:"total_price"`
            Pieces float64  `json:"pieces"`
            Sizes float64  `json:"sizes"`
         
    }
    
    // ShipmentUserInfo ...
    type ShipmentUserInfo struct {

        
            FirstName string  `json:"first_name"`
            Mobile string  `json:"mobile"`
            LastName string  `json:"last_name"`
            Gender string  `json:"gender"`
         
    }
    
    // CurrentStatus ...
    type CurrentStatus struct {

        
            Status string  `json:"status"`
            Name string  `json:"name"`
            JourneyType string  `json:"journey_type"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    
    // ItemBrand ...
    type ItemBrand struct {

        
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
    }
    
    // Item ...
    type Item struct {

        
            Name string  `json:"name"`
            Brand ItemBrand  `json:"brand"`
            Image []string  `json:"image"`
            ID float64  `json:"id"`
            Code string  `json:"code"`
            SlugKey string  `json:"slug_key"`
            Size string  `json:"size"`
            SellerIdentifier string  `json:"seller_identifier"`
         
    }
    
    // Prices ...
    type Prices struct {

        
            Cashback float64  `json:"cashback"`
            CashbackApplied float64  `json:"cashback_applied"`
            Discount float64  `json:"discount"`
            RefundCredit float64  `json:"refund_credit"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            PriceEffective float64  `json:"price_effective"`
            PriceMarked float64  `json:"price_marked"`
            DeliveryCharge float64  `json:"delivery_charge"`
            FyndCredits float64  `json:"fynd_credits"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            AmountPaid float64  `json:"amount_paid"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            CodCharges float64  `json:"cod_charges"`
            RefundAmount float64  `json:"refund_amount"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            ValueOfGood float64  `json:"value_of_good"`
            CouponValue float64  `json:"coupon_value"`
            TransferPrice float64  `json:"transfer_price"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
         
    }
    
    // Identifiers ...
    type Identifiers struct {

        
            SkuCode string  `json:"sku_code"`
            Ean string  `json:"ean"`
         
    }
    
    // FinancialBreakup ...
    type FinancialBreakup struct {

        
            Cashback float64  `json:"cashback"`
            CashbackApplied float64  `json:"cashback_applied"`
            GstFee float64  `json:"gst_fee"`
            Discount float64  `json:"discount"`
            RefundCredit float64  `json:"refund_credit"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            Size string  `json:"size"`
            PriceEffective float64  `json:"price_effective"`
            PriceMarked float64  `json:"price_marked"`
            DeliveryCharge float64  `json:"delivery_charge"`
            HsnCode string  `json:"hsn_code"`
            FyndCredits float64  `json:"fynd_credits"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            AmountPaid float64  `json:"amount_paid"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            CodCharges float64  `json:"cod_charges"`
            Identifiers Identifiers  `json:"identifiers"`
            RefundAmount float64  `json:"refund_amount"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            TotalUnits float64  `json:"total_units"`
            ValueOfGood float64  `json:"value_of_good"`
            GstTag string  `json:"gst_tag"`
            CouponValue float64  `json:"coupon_value"`
            TransferPrice float64  `json:"transfer_price"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            ItemName string  `json:"item_name"`
         
    }
    
    // AppliedPromos ...
    type AppliedPromos struct {

        
            PromotionName string  `json:"promotion_name"`
            ArticleQuantity float64  `json:"article_quantity"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            PromoID string  `json:"promo_id"`
            PromotionType string  `json:"promotion_type"`
            MrpPromotion bool  `json:"mrp_promotion"`
            Amount float64  `json:"amount"`
         
    }
    
    // Bags ...
    type Bags struct {

        
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            ReturnableDate string  `json:"returnable_date"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            Item Item  `json:"item"`
            CanCancel bool  `json:"can_cancel"`
            Quantity float64  `json:"quantity"`
            DeliveryDate string  `json:"delivery_date"`
            ID float64  `json:"id"`
            Prices Prices  `json:"prices"`
            LineNumber float64  `json:"line_number"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            CanReturn bool  `json:"can_return"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            SellerIdentifier string  `json:"seller_identifier"`
         
    }
    
    // BreakupValues ...
    type BreakupValues struct {

        
            Value float64  `json:"value"`
            Name string  `json:"name"`
            Display string  `json:"display"`
         
    }
    
    // FulfillingStore ...
    type FulfillingStore struct {

        
            Name string  `json:"name"`
            CompanyName string  `json:"company_name"`
            ID float64  `json:"id"`
            Code string  `json:"code"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // NestedTrackingDetails ...
    type NestedTrackingDetails struct {

        
            Status string  `json:"status"`
            IsCurrent bool  `json:"is_current"`
            Time string  `json:"time"`
            IsPassed bool  `json:"is_passed"`
         
    }
    
    // TrackingDetails ...
    type TrackingDetails struct {

        
            Status string  `json:"status"`
            Time string  `json:"time"`
            TrackingDetails []NestedTrackingDetails  `json:"tracking_details"`
            IsCurrent bool  `json:"is_current"`
            IsPassed bool  `json:"is_passed"`
         
    }
    
    // DeliveryAddress ...
    type DeliveryAddress struct {

        
            Address2 string  `json:"address2"`
            ContactPerson string  `json:"contact_person"`
            Version string  `json:"version"`
            State string  `json:"state"`
            CreatedAt string  `json:"created_at"`
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
            Phone string  `json:"phone"`
            City string  `json:"city"`
            Address string  `json:"address"`
            UpdatedAt string  `json:"updated_at"`
            Address1 string  `json:"address1"`
            AddressType string  `json:"address_type"`
            Pincode string  `json:"pincode"`
            AddressCategory string  `json:"address_category"`
            Name string  `json:"name"`
            Email string  `json:"email"`
            Area string  `json:"area"`
         
    }
    
    // Shipments ...
    type Shipments struct {

        
            ReturnableDate string  `json:"returnable_date"`
            Payment ShipmentPayment  `json:"payment"`
            SizeInfo map[string]interface{}  `json:"size_info"`
            Invoice Invoice  `json:"invoice"`
            OrderID string  `json:"order_id"`
            TrackURL string  `json:"track_url"`
            ShowDownloadInvoice bool  `json:"show_download_invoice"`
            TotalBags float64  `json:"total_bags"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            FulfillingCompany FulfillingCompany  `json:"fulfilling_company"`
            TrakingNo string  `json:"traking_no"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            BeneficiaryDetails bool  `json:"beneficiary_details"`
            Comment string  `json:"comment"`
            AwbNo string  `json:"awb_no"`
            RefundDetails map[string]interface{}  `json:"refund_details"`
            Promise Promise  `json:"promise"`
            DpName string  `json:"dp_name"`
            CanCancel bool  `json:"can_cancel"`
            TotalDetails ShipmentTotalDetails  `json:"total_details"`
            UserInfo ShipmentUserInfo  `json:"user_info"`
            Bags []Bags  `json:"bags"`
            CanReturn bool  `json:"can_return"`
            ShipmentID string  `json:"shipment_id"`
            DeliveryDate string  `json:"delivery_date"`
            NeedHelpURL string  `json:"need_help_url"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            ShowTrackLink bool  `json:"show_track_link"`
            CanBreak map[string]interface{}  `json:"can_break"`
            OrderType string  `json:"order_type"`
            TrackingDetails []TrackingDetails  `json:"tracking_details"`
            Prices Prices  `json:"prices"`
            DeliveryAddress DeliveryAddress  `json:"delivery_address"`
         
    }
    
    // OrderSchema ...
    type OrderSchema struct {

        
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
            BagsForReorder []BagsForReorder  `json:"bags_for_reorder"`
            UserInfo UserInfo  `json:"user_info"`
            OrderCreatedTime string  `json:"order_created_time"`
            Shipments []Shipments  `json:"shipments"`
            OrderID string  `json:"order_id"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
         
    }
    
    // OrderPage ...
    type OrderPage struct {

        
            Type string  `json:"type"`
            ItemTotal float64  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Size float64  `json:"size"`
            Current float64  `json:"current"`
         
    }
    
    // OrderList ...
    type OrderList struct {

        
            Filters OrderFilters  `json:"filters"`
            Items []OrderSchema  `json:"items"`
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
            PresignedURL string  `json:"presigned_url"`
            PresignedType string  `json:"presigned_type"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // Track ...
    type Track struct {

        
            Status string  `json:"status"`
            AccountName string  `json:"account_name"`
            UpdatedAt string  `json:"updated_at"`
            Awb string  `json:"awb"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            Reason string  `json:"reason"`
            ShipmentType string  `json:"shipment_type"`
            UpdatedTime string  `json:"updated_time"`
         
    }
    
    // ShipmentTrack ...
    type ShipmentTrack struct {

        
            Results []Track  `json:"results"`
         
    }
    
    // CustomerDetailsResponse ...
    type CustomerDetailsResponse struct {

        
            Name string  `json:"name"`
            ShipmentID string  `json:"shipment_id"`
            OrderID string  `json:"order_id"`
            Country string  `json:"country"`
            Phone string  `json:"phone"`
         
    }
    
    // SendOtpToCustomerResponse ...
    type SendOtpToCustomerResponse struct {

        
            Success bool  `json:"success"`
            RequestID string  `json:"request_id"`
            ResendTimer float64  `json:"resend_timer"`
            Message string  `json:"message"`
         
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
            QcType []string  `json:"qc_type"`
            Reasons []BagReasons  `json:"reasons"`
            ID float64  `json:"id"`
            Meta BagReasonMeta  `json:"meta"`
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
            ReasonID float64  `json:"reason_id"`
            Flow string  `json:"flow"`
            ShowTextArea bool  `json:"show_text_area"`
            ReasonText string  `json:"reason_text"`
            FeedbackType string  `json:"feedback_type"`
         
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

        
            Products []ProductsDataUpdates1  `json:"products"`
            Entities []EntitiesDataUpdates  `json:"entities"`
         
    }
    
    // ProductsReasonsFilters ...
    type ProductsReasonsFilters struct {

        
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
         
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

        
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // ShipmentsRequest ...
    type ShipmentsRequest struct {

        
            DataUpdates DataUpdates  `json:"data_updates"`
            Identifier string  `json:"identifier"`
            Reasons ReasonsData  `json:"reasons"`
            Products []Products  `json:"products"`
         
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
            Task bool  `json:"task"`
            LockAfterTransition bool  `json:"lock_after_transition"`
            ForceTransition bool  `json:"force_transition"`
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
    
    // ProductsReasonsFilters1 ...
    type ProductsReasonsFilters1 struct {

        
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductsReasonsData1 ...
    type ProductsReasonsData1 struct {

        
            ReasonText string  `json:"reason_text"`
            ReasonID float64  `json:"reason_id"`
         
    }
    
    // ProductsReasons1 ...
    type ProductsReasons1 struct {

        
            Filters []ProductsReasonsFilters1  `json:"filters"`
            Data ProductsReasonsData1  `json:"data"`
         
    }
    
    // EntityReasonData1 ...
    type EntityReasonData1 struct {

        
            ReasonText string  `json:"reason_text"`
            ReasonID float64  `json:"reason_id"`
         
    }
    
    // EntitiesReasons1 ...
    type EntitiesReasons1 struct {

        
            Filters []map[string]interface{}  `json:"filters"`
            Data EntityReasonData1  `json:"data"`
         
    }
    
    // ReasonsData1 ...
    type ReasonsData1 struct {

        
            Products []ProductsReasons1  `json:"products"`
            Entities []EntitiesReasons1  `json:"entities"`
         
    }
    
    // Products1 ...
    type Products1 struct {

        
            LineNumber float64  `json:"line_number"`
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductsDataUpdatesFilters1 ...
    type ProductsDataUpdatesFilters1 struct {

        
            LineNumber float64  `json:"line_number"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductsDataUpdates1 ...
    type ProductsDataUpdates1 struct {

        
            Filters []ProductsDataUpdatesFilters1  `json:"filters"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // ShipmentsRequest1 ...
    type ShipmentsRequest1 struct {

        
            Reasons ReasonsData1  `json:"reasons"`
            Products []Products1  `json:"products"`
            DataUpdates DataUpdates  `json:"data_updates"`
            Identifier string  `json:"identifier"`
         
    }
    
    // creditNoteParameter ...
    type creditNoteParameter struct {

        
            ExpiresIn float64  `json:"expires_in"`
         
    }
    
    // invoiceParameter ...
    type invoiceParameter struct {

        
            DocumentType string  `json:"document_type"`
            ExpiresIn float64  `json:"expires_in"`
         
    }
    
    // ResponseGetInvoiceShipment1 ...
    type ResponseGetInvoiceShipment1 struct {

        
            Success bool  `json:"success"`
            PresignedType string  `json:"presigned_type"`
            ShipmentID string  `json:"shipment_id"`
            PresignedURL string  `json:"presigned_url"`
         
    }
    
    // getInvoiceByShipmentId400Response ...
    type getInvoiceByShipmentId400Response struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // getInvoiceByShipmentId500Response ...
    type getInvoiceByShipmentId500Response struct {

        
            Success bool  `json:"success"`
            PresignedType string  `json:"presigned_type"`
         
    }
    
    // ReqBodyPresignedPOST ...
    type ReqBodyPresignedPOST struct {

        
            Event string  `json:"event"`
            MediaType []interface{}  `json:"media_type"`
            ExpiresIn float64  `json:"expires_in"`
         
    }
    
    // ResponsePresignedGETURL ...
    type ResponsePresignedGETURL struct {

        
            Success bool  `json:"success"`
            PresignedType string  `json:"presigned_type"`
            ShipmentID string  `json:"shipment_id"`
            PresignedURL string  `json:"presigned_url"`
         
    }
    
    // ErrorResponsePresignedUrl ...
    type ErrorResponsePresignedUrl struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            RequestID string  `json:"request_id"`
            Exception string  `json:"exception"`
            StackTrace string  `json:"stack_trace"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    

    
    // ActionPageParams ...
    type ActionPageParams struct {

        
            Slug []string  `json:"slug"`
         
    }
    
    // CatalogueOrderRequest ...
    type CatalogueOrderRequest struct {

        
            Articles []RewardsArticle  `json:"articles"`
         
    }
    
    // CatalogueOrderResponse ...
    type CatalogueOrderResponse struct {

        
            Articles []RewardsArticle  `json:"articles"`
         
    }
    
    // DiscountProperties ...
    type DiscountProperties struct {

        
            Absolute float64  `json:"absolute"`
            Currency string  `json:"currency"`
            DisplayAbsolute string  `json:"display_absolute"`
            DisplayPercent string  `json:"display_percent"`
            Percent float64  `json:"percent"`
         
    }
    
    // Error ...
    type Error struct {

        
            Code float64  `json:"code"`
            Exception string  `json:"exception"`
            Info string  `json:"info"`
            Message string  `json:"message"`
         
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
    
    // OrderDiscountRequest ...
    type OrderDiscountRequest struct {

        
            Currency string  `json:"currency"`
            OrderAmount float64  `json:"order_amount"`
         
    }
    
    // OrderDiscountResponse ...
    type OrderDiscountResponse struct {

        
            AppliedRuleBucket OrderDiscountRuleBucket  `json:"applied_rule_bucket"`
            BaseDiscount DiscountProperties  `json:"base_discount"`
            Discount DiscountProperties  `json:"discount"`
            OrderAmount float64  `json:"order_amount"`
            Points float64  `json:"points"`
         
    }
    
    // OrderDiscountRuleBucket ...
    type OrderDiscountRuleBucket struct {

        
            High float64  `json:"high"`
            Low float64  `json:"low"`
            Max float64  `json:"max"`
            Value float64  `json:"value"`
            ValueType string  `json:"value_type"`
         
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
    
    // PointsHistoryResponse ...
    type PointsHistoryResponse struct {

        
            Items []PointsHistory  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PointsResponse ...
    type PointsResponse struct {

        
            Points float64  `json:"points"`
         
    }
    
    // RedeemReferralCodeRequest ...
    type RedeemReferralCodeRequest struct {

        
            DeviceID string  `json:"device_id"`
            ReferralCode string  `json:"referral_code"`
         
    }
    
    // RedeemReferralCodeResponse ...
    type RedeemReferralCodeResponse struct {

        
            Message string  `json:"message"`
            Points float64  `json:"points"`
            Redeemed bool  `json:"redeemed"`
            ReferrerID string  `json:"referrer_id"`
            ReferrerInfo string  `json:"referrer_info"`
         
    }
    
    // ReferralDetailsResponse ...
    type ReferralDetailsResponse struct {

        
            Referral Offer  `json:"referral"`
            ReferrerInfo string  `json:"referrer_info"`
            Share ShareMessages  `json:"share"`
            User ReferralDetailsUser  `json:"user"`
         
    }
    
    // ReferralDetailsUser ...
    type ReferralDetailsUser struct {

        
            Blocked bool  `json:"blocked"`
            Points float64  `json:"points"`
            Redeemed bool  `json:"redeemed"`
            ReferralCode string  `json:"referral_code"`
         
    }
    
    // RewardsArticle ...
    type RewardsArticle struct {

        
            ID string  `json:"id"`
            Points float64  `json:"points"`
            Price float64  `json:"price"`
         
    }
    
    // Schedule ...
    type Schedule struct {

        
            Cron string  `json:"cron"`
            Duration float64  `json:"duration"`
            End string  `json:"end"`
            Start string  `json:"start"`
         
    }
    
    // ShareMessages ...
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
    

    
    // UpdateCartShipmentItem ...
    type UpdateCartShipmentItem struct {

        
            Quantity float64  `json:"quantity"`
            ArticleUID string  `json:"article_uid"`
            ShipmentType string  `json:"shipment_type"`
         
    }
    
    // UpdateCartShipmentRequest ...
    type UpdateCartShipmentRequest struct {

        
            Shipments []UpdateCartShipmentItem  `json:"shipments"`
         
    }
    
    // Files ...
    type Files struct {

        
            Key string  `json:"key"`
            Values []string  `json:"values"`
         
    }
    
    // CartPosCheckoutDetailRequest ...
    type CartPosCheckoutDetailRequest struct {

        
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            Meta map[string]interface{}  `json:"meta"`
            Staff StaffCheckout  `json:"staff"`
            OrderingStore float64  `json:"ordering_store"`
            Pos bool  `json:"pos"`
            PaymentMode string  `json:"payment_mode"`
            MerchantCode string  `json:"merchant_code"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            OrderType string  `json:"order_type"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            PaymentIdentifier string  `json:"payment_identifier"`
            Aggregator string  `json:"aggregator"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            Files []Files  `json:"files"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            BillingAddressID string  `json:"billing_address_id"`
            CallbackURL string  `json:"callback_url"`
            AddressID string  `json:"address_id"`
         
    }
    
    // CartDeliveryModesResponse ...
    type CartDeliveryModesResponse struct {

        
            AvailableModes []string  `json:"available_modes"`
            PickupStores []float64  `json:"pickup_stores"`
         
    }
    
    // PickupStoreDetail ...
    type PickupStoreDetail struct {

        
            Pincode float64  `json:"pincode"`
            ID float64  `json:"id"`
            AddressType string  `json:"address_type"`
            Landmark string  `json:"landmark"`
            Name string  `json:"name"`
            Email string  `json:"email"`
            City string  `json:"city"`
            UID float64  `json:"uid"`
            StoreCode string  `json:"store_code"`
            Area string  `json:"area"`
            State string  `json:"state"`
            Country string  `json:"country"`
            Phone string  `json:"phone"`
            AreaCode string  `json:"area_code"`
            Address string  `json:"address"`
            AreaCodeSlug string  `json:"area_code_slug"`
         
    }
    
    // StoreDetailsResponse ...
    type StoreDetailsResponse struct {

        
            Items []PickupStoreDetail  `json:"items"`
         
    }
    

    
    // GetPincodeCityResponse ...
    type GetPincodeCityResponse struct {

        
            RequestUUID string  `json:"request_uuid"`
            StormbreakerUUID string  `json:"stormbreaker_uuid"`
            Success bool  `json:"success"`
            Data []LogisticPincodeData  `json:"data"`
         
    }
    
    // LogisticPincodeData ...
    type LogisticPincodeData struct {

        
            Meta LogisticMeta  `json:"meta"`
            Parents []LogisticParents  `json:"parents"`
            SubType string  `json:"sub_type"`
            Name string  `json:"name"`
            Error LogisticError  `json:"error"`
            UID string  `json:"uid"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // LogisticMeta ...
    type LogisticMeta struct {

        
            Zone string  `json:"zone"`
            Deliverables []interface{}  `json:"deliverables"`
         
    }
    
    // LogisticParents ...
    type LogisticParents struct {

        
            SubType string  `json:"sub_type"`
            Name string  `json:"name"`
            DisplayName string  `json:"display_name"`
            UID string  `json:"uid"`
         
    }
    
    // LogisticError ...
    type LogisticError struct {

        
            Type string  `json:"type"`
            Value string  `json:"value"`
            Message string  `json:"message"`
         
    }
    
    // GetPincodeZonesReqBody ...
    type GetPincodeZonesReqBody struct {

        
            Country string  `json:"country"`
            Pincode string  `json:"pincode"`
         
    }
    
    // GetPincodeZonesResponse ...
    type GetPincodeZonesResponse struct {

        
            Zones []interface{}  `json:"zones"`
            ServiceabilityType string  `json:"serviceability_type"`
         
    }
    
    // GetTatProductReqBody ...
    type GetTatProductReqBody struct {

        
            LocationDetails []LocationDetailsReq  `json:"location_details"`
            ToPincode string  `json:"to_pincode"`
            Action string  `json:"action"`
         
    }
    
    // LocationDetailsReq ...
    type LocationDetailsReq struct {

        
            FromPincode string  `json:"from_pincode"`
            Articles []TatReqProductArticles  `json:"articles"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // TatReqProductArticles ...
    type TatReqProductArticles struct {

        
            ManufacturingTime float64  `json:"manufacturing_time"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            Category LogisticRequestCategory  `json:"category"`
         
    }
    
    // LogisticRequestCategory ...
    type LogisticRequestCategory struct {

        
            ID float64  `json:"id"`
            Level string  `json:"level"`
         
    }
    
    // GetTatProductResponse ...
    type GetTatProductResponse struct {

        
            LocationDetails []LocationDetails  `json:"location_details"`
            RequestUUID string  `json:"request_uuid"`
            Error map[string]interface{}  `json:"error"`
            ToCity string  `json:"to_city"`
            Source string  `json:"source"`
            ToPincode string  `json:"to_pincode"`
            Action string  `json:"action"`
            StormbreakerUUID string  `json:"stormbreaker_uuid"`
            Success bool  `json:"success"`
            Identifier string  `json:"identifier"`
            Journey string  `json:"journey"`
         
    }
    
    // LocationDetails ...
    type LocationDetails struct {

        
            FromPincode string  `json:"from_pincode"`
            Articles []TatProductArticles  `json:"articles"`
            FulfillmentID float64  `json:"fulfillment_id"`
         
    }
    
    // TatProductArticles ...
    type TatProductArticles struct {

        
            Error map[string]interface{}  `json:"error"`
            Category LogisticResponseCategory  `json:"category"`
            Promise LogisticPromise  `json:"promise"`
         
    }
    
    // LogisticResponseCategory ...
    type LogisticResponseCategory struct {

        
            ID float64  `json:"id"`
            Level string  `json:"level"`
         
    }
    
    // LogisticPromise ...
    type LogisticPromise struct {

        
            Timestamp LogisticTimestamp  `json:"timestamp"`
            Formatted Formatted  `json:"formatted"`
         
    }
    
    // LogisticTimestamp ...
    type LogisticTimestamp struct {

        
            Min float64  `json:"min"`
            Max float64  `json:"max"`
         
    }
    
    // Formatted ...
    type Formatted struct {

        
            Min string  `json:"min"`
            Max string  `json:"max"`
         
    }
    
