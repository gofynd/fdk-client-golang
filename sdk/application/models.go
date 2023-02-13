package application



    
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
    
    // Meta ...
    type Meta struct {

        
            Source string  `json:"source"`
         
    }
    
    // Media ...
    type Media struct {

        
            URL string  `json:"url"`
            Meta Meta  `json:"meta"`
            Alt string  `json:"alt"`
            Type string  `json:"type"`
         
    }
    
    // ProductBrand ...
    type ProductBrand struct {

        
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Description string  `json:"description"`
            Logo Media  `json:"logo"`
         
    }
    
    // CustomMetaFields ...
    type CustomMetaFields struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // ApplicationItemMOQ ...
    type ApplicationItemMOQ struct {

        
            Maximum float64  `json:"maximum"`
            Minimum float64  `json:"minimum"`
            IncrementUnit float64  `json:"increment_unit"`
         
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
    
    // ApplicationItemSEO ...
    type ApplicationItemSEO struct {

        
            Title interface{}  `json:"title"`
            Description interface{}  `json:"description"`
         
    }
    
    // ProductDetail ...
    type ProductDetail struct {

        
            Tryouts []string  `json:"tryouts"`
            Categories []ProductBrand  `json:"categories"`
            Brand ProductBrand  `json:"brand"`
            Similars []string  `json:"similars"`
            ItemType string  `json:"item_type"`
            Color string  `json:"color"`
            IsDependent bool  `json:"is_dependent"`
            ItemCode string  `json:"item_code"`
            Medias []Media  `json:"medias"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Name string  `json:"name"`
            Moq ApplicationItemMOQ  `json:"moq"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            Type string  `json:"type"`
            ImageNature string  `json:"image_nature"`
            Price ProductListingPrice  `json:"price"`
            Action ProductListingAction  `json:"action"`
            Discount string  `json:"discount"`
            RatingCount float64  `json:"rating_count"`
            Rating float64  `json:"rating"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            HasVariant bool  `json:"has_variant"`
            Attributes map[string]interface{}  `json:"attributes"`
            TeaserTag string  `json:"teaser_tag"`
            Seo ApplicationItemSEO  `json:"seo"`
            ProductOnlineDate string  `json:"product_online_date"`
            Highlights []string  `json:"highlights"`
            Description string  `json:"description"`
            Tags []string  `json:"tags"`
            ShortDescription string  `json:"short_description"`
         
    }
    
    // ErrorResponse ...
    type ErrorResponse struct {

        
            Error string  `json:"error"`
         
    }
    
    // ColumnHeader ...
    type ColumnHeader struct {

        
            Convertable bool  `json:"convertable"`
            Value string  `json:"value"`
         
    }
    
    // ColumnHeaders ...
    type ColumnHeaders struct {

        
            Col2 ColumnHeader  `json:"col_2"`
            Col4 ColumnHeader  `json:"col_4"`
            Col6 ColumnHeader  `json:"col_6"`
            Col3 ColumnHeader  `json:"col_3"`
            Col5 ColumnHeader  `json:"col_5"`
            Col1 ColumnHeader  `json:"col_1"`
         
    }
    
    // SizeChartValues ...
    type SizeChartValues struct {

        
            Col2 string  `json:"col_2"`
            Col4 string  `json:"col_4"`
            Col6 string  `json:"col_6"`
            Col3 string  `json:"col_3"`
            Col5 string  `json:"col_5"`
            Col1 string  `json:"col_1"`
         
    }
    
    // SizeChart ...
    type SizeChart struct {

        
            Headers ColumnHeaders  `json:"headers"`
            Image string  `json:"image"`
            Unit string  `json:"unit"`
            Sizes []SizeChartValues  `json:"sizes"`
            SizeTip string  `json:"size_tip"`
            Title string  `json:"title"`
            Description string  `json:"description"`
         
    }
    
    // ProductSizeStores ...
    type ProductSizeStores struct {

        
            Count float64  `json:"count"`
         
    }
    
    // Dimension ...
    type Dimension struct {

        
            IsDefault bool  `json:"is_default"`
            Height float64  `json:"height"`
            Length float64  `json:"length"`
            Width float64  `json:"width"`
            Unit string  `json:"unit"`
         
    }
    
    // Weight ...
    type Weight struct {

        
            Shipping float64  `json:"shipping"`
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
         
    }
    
    // ProductSize ...
    type ProductSize struct {

        
            Dimension Dimension  `json:"dimension"`
            SellerIdentifiers []string  `json:"seller_identifiers"`
            Quantity float64  `json:"quantity"`
            Weight Weight  `json:"weight"`
            IsAvailable bool  `json:"is_available"`
            Value string  `json:"value"`
            Display string  `json:"display"`
         
    }
    
    // ProductSizes ...
    type ProductSizes struct {

        
            SizeChart SizeChart  `json:"size_chart"`
            Price ProductListingPrice  `json:"price"`
            Discount string  `json:"discount"`
            Stores ProductSizeStores  `json:"stores"`
            Sizes []ProductSize  `json:"sizes"`
            Sellable bool  `json:"sellable"`
            MultiSize bool  `json:"multi_size"`
         
    }
    
    // AttributeDetail ...
    type AttributeDetail struct {

        
            Description string  `json:"description"`
            Display string  `json:"display"`
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

        
            Items []ProductDetail  `json:"items"`
            AttributesMetadata []AttributeMetadata  `json:"attributes_metadata"`
         
    }
    
    // ProductCompareResponse ...
    type ProductCompareResponse struct {

        
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
            Items []ProductDetail  `json:"items"`
            AttributesMetadata []AttributeMetadata  `json:"attributes_metadata"`
         
    }
    
    // ProductFrequentlyComparedSimilarResponse ...
    type ProductFrequentlyComparedSimilarResponse struct {

        
            Similars ProductCompareResponse  `json:"similars"`
         
    }
    
    // ProductVariantItemResponse ...
    type ProductVariantItemResponse struct {

        
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            IsAvailable bool  `json:"is_available"`
            Value string  `json:"value"`
            Color string  `json:"color"`
            ColorName string  `json:"color_name"`
            Medias []Media  `json:"medias"`
         
    }
    
    // ProductVariantResponse ...
    type ProductVariantResponse struct {

        
            Header string  `json:"header"`
            Items []ProductVariantItemResponse  `json:"items"`
            DisplayType string  `json:"display_type"`
            Key string  `json:"key"`
         
    }
    
    // ProductVariantsResponse ...
    type ProductVariantsResponse struct {

        
            Variants []ProductVariantResponse  `json:"variants"`
         
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
    
    // StoreDetail ...
    type StoreDetail struct {

        
            ID float64  `json:"id"`
            Code string  `json:"code"`
            City string  `json:"city"`
            Name string  `json:"name"`
         
    }
    
    // Seller ...
    type Seller struct {

        
            UID float64  `json:"uid"`
            Count float64  `json:"count"`
            Name string  `json:"name"`
         
    }
    
    // ProductStockStatusItem ...
    type ProductStockStatusItem struct {

        
            Company CompanyDetail  `json:"company"`
            Price ProductStockPrice  `json:"price"`
            Store StoreDetail  `json:"store"`
            Seller Seller  `json:"seller"`
            UID string  `json:"uid"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
            Identifier map[string]interface{}  `json:"identifier"`
            Size string  `json:"size"`
         
    }
    
    // ProductStockStatusResponse ...
    type ProductStockStatusResponse struct {

        
            Items []ProductStockStatusItem  `json:"items"`
         
    }
    
    // ProductStockPolling ...
    type ProductStockPolling struct {

        
            Page Page  `json:"page"`
            Items []ProductStockStatusItem  `json:"items"`
         
    }
    
    // ProductVariantListingResponse ...
    type ProductVariantListingResponse struct {

        
            DisplayType string  `json:"display_type"`
            Header string  `json:"header"`
            Items []ProductVariantItemResponse  `json:"items"`
            Total float64  `json:"total"`
            Key string  `json:"key"`
         
    }
    
    // NetQuantity ...
    type NetQuantity struct {

        
            Unit interface{}  `json:"unit"`
            Value float64  `json:"value"`
         
    }
    
    // ProductListingDetail ...
    type ProductListingDetail struct {

        
            Tryouts []string  `json:"tryouts"`
            Categories []ProductBrand  `json:"categories"`
            Brand ProductBrand  `json:"brand"`
            Similars []string  `json:"similars"`
            ItemType string  `json:"item_type"`
            Color string  `json:"color"`
            Sellable bool  `json:"sellable"`
            IsDependent bool  `json:"is_dependent"`
            Variants []ProductVariantListingResponse  `json:"variants"`
            ItemCode string  `json:"item_code"`
            Medias []Media  `json:"medias"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Name string  `json:"name"`
            Moq ApplicationItemMOQ  `json:"moq"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            Identifiers []string  `json:"identifiers"`
            Type string  `json:"type"`
            ImageNature string  `json:"image_nature"`
            Price ProductListingPrice  `json:"price"`
            Action ProductListingAction  `json:"action"`
            Discount string  `json:"discount"`
            RatingCount float64  `json:"rating_count"`
            Rating float64  `json:"rating"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Sizes []string  `json:"sizes"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            HasVariant bool  `json:"has_variant"`
            Attributes map[string]interface{}  `json:"attributes"`
            TeaserTag string  `json:"teaser_tag"`
            Seo ApplicationItemSEO  `json:"seo"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            ProductOnlineDate string  `json:"product_online_date"`
            Highlights []string  `json:"highlights"`
            Description string  `json:"description"`
            Tags []string  `json:"tags"`
            ShortDescription string  `json:"short_description"`
         
    }
    
    // ProductSortOn ...
    type ProductSortOn struct {

        
            IsSelected bool  `json:"is_selected"`
            Value string  `json:"value"`
            Name string  `json:"name"`
         
    }
    
    // ProductFiltersValue ...
    type ProductFiltersValue struct {

        
            DisplayFormat string  `json:"display_format"`
            CurrencySymbol string  `json:"currency_symbol"`
            QueryFormat string  `json:"query_format"`
            Min float64  `json:"min"`
            SelectedMin float64  `json:"selected_min"`
            CurrencyCode string  `json:"currency_code"`
            SelectedMax float64  `json:"selected_max"`
            IsSelected bool  `json:"is_selected"`
            Max float64  `json:"max"`
            Count float64  `json:"count"`
            Value string  `json:"value"`
            Display string  `json:"display"`
         
    }
    
    // ProductFiltersKey ...
    type ProductFiltersKey struct {

        
            Kind string  `json:"kind"`
            Display string  `json:"display"`
            Logo string  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // ProductFilters ...
    type ProductFilters struct {

        
            Values []ProductFiltersValue  `json:"values"`
            Key ProductFiltersKey  `json:"key"`
         
    }
    
    // ProductListingResponse ...
    type ProductListingResponse struct {

        
            Page Page  `json:"page"`
            Items []ProductListingDetail  `json:"items"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Filters []ProductFilters  `json:"filters"`
         
    }
    
    // ImageUrls ...
    type ImageUrls struct {

        
            Landscape Media  `json:"landscape"`
            Portrait Media  `json:"portrait"`
         
    }
    
    // BrandItem ...
    type BrandItem struct {

        
            Action ProductListingAction  `json:"action"`
            Discount string  `json:"discount"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            Departments []string  `json:"departments"`
            Logo Media  `json:"logo"`
            Description string  `json:"description"`
         
    }
    
    // BrandListingResponse ...
    type BrandListingResponse struct {

        
            Page Page  `json:"page"`
            Items []BrandItem  `json:"items"`
         
    }
    
    // BrandDetailResponse ...
    type BrandDetailResponse struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Description string  `json:"description"`
            Logo Media  `json:"logo"`
         
    }
    
    // DepartmentIdentifier ...
    type DepartmentIdentifier struct {

        
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
         
    }
    
    // ThirdLevelChild ...
    type ThirdLevelChild struct {

        
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Childs []map[string]interface{}  `json:"childs"`
         
    }
    
    // SecondLevelChild ...
    type SecondLevelChild struct {

        
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Childs []ThirdLevelChild  `json:"childs"`
         
    }
    
    // Child ...
    type Child struct {

        
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Childs []SecondLevelChild  `json:"childs"`
         
    }
    
    // CategoryItems ...
    type CategoryItems struct {

        
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            Childs []Child  `json:"childs"`
         
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

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Logo Media  `json:"logo"`
         
    }
    
    // HomeListingResponse ...
    type HomeListingResponse struct {

        
            Page Page  `json:"page"`
            Message string  `json:"message"`
            Items []ProductListingDetail  `json:"items"`
         
    }
    
    // Department ...
    type Department struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            PriorityOrder float64  `json:"priority_order"`
            Logo Media  `json:"logo"`
         
    }
    
    // DepartmentResponse ...
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // AutocompleteItem ...
    type AutocompleteItem struct {

        
            Action ProductListingAction  `json:"action"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Display string  `json:"display"`
            Type string  `json:"type"`
            Logo Media  `json:"logo"`
         
    }
    
    // AutoCompleteResponse ...
    type AutoCompleteResponse struct {

        
            Items []AutocompleteItem  `json:"items"`
         
    }
    
    // CollectionQuery ...
    type CollectionQuery struct {

        
            Attribute string  `json:"attribute"`
            Op string  `json:"op"`
            Value []interface{}  `json:"value"`
         
    }
    
    // GetCollectionDetailNest ...
    type GetCollectionDetailNest struct {

        
            Banners ImageUrls  `json:"banners"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Meta map[string]interface{}  `json:"meta"`
            Name string  `json:"name"`
            AppID string  `json:"app_id"`
            Slug string  `json:"slug"`
            UID string  `json:"uid"`
            IsActive bool  `json:"is_active"`
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            Logo Media  `json:"logo"`
            Tag []string  `json:"tag"`
            Action ProductListingAction  `json:"action"`
            SortOn string  `json:"sort_on"`
            AllowSort bool  `json:"allow_sort"`
            Priority float64  `json:"priority"`
            AllowFacets bool  `json:"allow_facets"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Cron map[string]interface{}  `json:"cron"`
            Badge map[string]interface{}  `json:"badge"`
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

        
            Page Page  `json:"page"`
            Items []GetCollectionDetailNest  `json:"items"`
            Filters CollectionListingFilter  `json:"filters"`
         
    }
    
    // CollectionDetailResponse ...
    type CollectionDetailResponse struct {

        
            Banners ImageUrls  `json:"banners"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Meta map[string]interface{}  `json:"meta"`
            Name string  `json:"name"`
            AppID string  `json:"app_id"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            Query []CollectionQuery  `json:"query"`
            Type string  `json:"type"`
            Logo Media  `json:"logo"`
            Tag []string  `json:"tag"`
            SortOn string  `json:"sort_on"`
            AllowSort bool  `json:"allow_sort"`
            Priority float64  `json:"priority"`
            AllowFacets bool  `json:"allow_facets"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Cron map[string]interface{}  `json:"cron"`
            Badge map[string]interface{}  `json:"badge"`
            Description string  `json:"description"`
         
    }
    
    // GetFollowListingResponse ...
    type GetFollowListingResponse struct {

        
            Page Page  `json:"page"`
            Items []ProductListingDetail  `json:"items"`
         
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
            Name string  `json:"name"`
            City string  `json:"city"`
            LatLong LatLong  `json:"lat_long"`
            StoreCode string  `json:"store_code"`
            UID float64  `json:"uid"`
            StoreEmail string  `json:"store_email"`
            State string  `json:"state"`
            Address string  `json:"address"`
            Country string  `json:"country"`
         
    }
    
    // StoreListingResponse ...
    type StoreListingResponse struct {

        
            Page Page  `json:"page"`
            Items []Store  `json:"items"`
         
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

        
            Number string  `json:"number"`
            CountryCode float64  `json:"country_code"`
         
    }
    
    // StoreManagerSerializer ...
    type StoreManagerSerializer struct {

        
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
            Email string  `json:"email"`
            Name string  `json:"name"`
         
    }
    
    // StoreAddressSerializer ...
    type StoreAddressSerializer struct {

        
            Address2 string  `json:"address2"`
            Country string  `json:"country"`
            City string  `json:"city"`
            Longitude float64  `json:"longitude"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
            Latitude float64  `json:"latitude"`
            Pincode float64  `json:"pincode"`
            Address1 string  `json:"address1"`
         
    }
    
    // StoreDepartments ...
    type StoreDepartments struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
            PriorityOrder float64  `json:"priority_order"`
            Logo string  `json:"logo"`
         
    }
    
    // AppStore ...
    type AppStore struct {

        
            Company CompanyStore  `json:"company"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Manager StoreManagerSerializer  `json:"manager"`
            Address StoreAddressSerializer  `json:"address"`
            Departments []StoreDepartments  `json:"departments"`
         
    }
    
    // ApplicationStoreListing ...
    type ApplicationStoreListing struct {

        
            Page Page  `json:"page"`
            Items []AppStore  `json:"items"`
            Filters []StoreDepartments  `json:"filters"`
         
    }
    
    // Time ...
    type Time struct {

        
            Hour float64  `json:"hour"`
            Minute float64  `json:"minute"`
         
    }
    
    // StoreTiming ...
    type StoreTiming struct {

        
            Weekday string  `json:"weekday"`
            Closing Time  `json:"closing"`
            Opening Time  `json:"opening"`
            Open bool  `json:"open"`
         
    }
    
    // StoreDetails ...
    type StoreDetails struct {

        
            Company CompanyStore  `json:"company"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Manager StoreManagerSerializer  `json:"manager"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Address StoreAddressSerializer  `json:"address"`
            Departments []StoreDepartments  `json:"departments"`
            Timing []StoreTiming  `json:"timing"`
         
    }
    
    // UserDetail ...
    type UserDetail struct {

        
            SuperUser bool  `json:"super_user"`
            Username string  `json:"username"`
            UserID string  `json:"user_id"`
            Contact string  `json:"contact"`
         
    }
    
    // ProductGroupPrice ...
    type ProductGroupPrice struct {

        
            MaxMarked float64  `json:"max_marked"`
            MinEffective float64  `json:"min_effective"`
            Currency interface{}  `json:"currency"`
            MaxEffective float64  `json:"max_effective"`
            MinMarked float64  `json:"min_marked"`
         
    }
    
    // ProductDetails ...
    type ProductDetails struct {

        
            CountryOfOrigin interface{}  `json:"country_of_origin"`
            ItemCode interface{}  `json:"item_code"`
            TemplateTag interface{}  `json:"template_tag"`
            Name interface{}  `json:"name"`
            Slug interface{}  `json:"slug"`
            OutOfStock bool  `json:"out_of_stock"`
            ImageNature interface{}  `json:"image_nature"`
            HsnCode float64  `json:"hsn_code"`
            RatingCount float64  `json:"rating_count"`
            Rating float64  `json:"rating"`
            Identifier map[string]interface{}  `json:"identifier"`
            GroupedAttributes map[string]interface{}  `json:"grouped_attributes"`
            HasVariant bool  `json:"has_variant"`
            IsSet bool  `json:"is_set"`
            Attributes map[string]interface{}  `json:"attributes"`
            Media []map[string]interface{}  `json:"media"`
            BrandUID float64  `json:"brand_uid"`
            Images []interface{}  `json:"images"`
            Highlights []interface{}  `json:"highlights"`
            Description interface{}  `json:"description"`
            ShortDescription interface{}  `json:"short_description"`
         
    }
    
    // Size ...
    type Size struct {

        
            Quantity float64  `json:"quantity"`
            Display interface{}  `json:"display"`
            IsAvailable bool  `json:"is_available"`
            Value interface{}  `json:"value"`
         
    }
    
    // ProductInGroup ...
    type ProductInGroup struct {

        
            ProductUID float64  `json:"product_uid"`
            Price ProductGroupPrice  `json:"price"`
            ProductDetails ProductDetails  `json:"product_details"`
            MaxQuantity float64  `json:"max_quantity"`
            AllowRemove bool  `json:"allow_remove"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            Sizes []Size  `json:"sizes"`
            AutoSelect bool  `json:"auto_select"`
            MinQuantity float64  `json:"min_quantity"`
         
    }
    
    // ProductGroupingModel ...
    type ProductGroupingModel struct {

        
            CreatedBy UserDetail  `json:"created_by"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            ModifiedBy UserDetail  `json:"modified_by"`
            Name interface{}  `json:"name"`
            PageVisibility []interface{}  `json:"page_visibility"`
            CreatedOn string  `json:"created_on"`
            Slug interface{}  `json:"slug"`
            CompanyID float64  `json:"company_id"`
            ModifiedOn string  `json:"modified_on"`
            Products []ProductInGroup  `json:"products"`
            VerifiedBy UserDetail  `json:"verified_by"`
            IsActive bool  `json:"is_active"`
            Choice interface{}  `json:"choice"`
            VerifiedOn string  `json:"verified_on"`
            Meta map[string]interface{}  `json:"meta"`
            ID interface{}  `json:"_id"`
            Logo string  `json:"logo"`
         
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
    
    // ProductStockPriceV2 ...
    type ProductStockPriceV2 struct {

        
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            Currency string  `json:"currency"`
         
    }
    
    // SellerV2 ...
    type SellerV2 struct {

        
            UID float64  `json:"uid"`
            Count float64  `json:"count"`
            Name string  `json:"name"`
         
    }
    
    // StoreV2 ...
    type StoreV2 struct {

        
            UID float64  `json:"uid"`
            Count float64  `json:"count"`
            Name string  `json:"name"`
         
    }
    
    // ReturnConfigSchemaV2 ...
    type ReturnConfigSchemaV2 struct {

        
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
         
    }
    
    // StrategyWiseListingSchemaV2 ...
    type StrategyWiseListingSchemaV2 struct {

        
            Quantity float64  `json:"quantity"`
            Distance float64  `json:"distance"`
            Tat float64  `json:"tat"`
            Pincode float64  `json:"pincode"`
         
    }
    
    // DetailsSchemaV2 ...
    type DetailsSchemaV2 struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            Key string  `json:"key"`
         
    }
    
    // MarketPlaceSttributesSchemaV2 ...
    type MarketPlaceSttributesSchemaV2 struct {

        
            Title string  `json:"title"`
            Details []DetailsSchemaV2  `json:"details"`
         
    }
    
    // ProductStockUnitPriceV2 ...
    type ProductStockUnitPriceV2 struct {

        
            Price float64  `json:"price"`
            Unit string  `json:"unit"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // SellerGroupAttributes ...
    type SellerGroupAttributes struct {

        
            Title string  `json:"title"`
            Details []DetailsSchemaV2  `json:"details"`
         
    }
    
    // ArticleAssignmentV2 ...
    type ArticleAssignmentV2 struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // ProductSizePriceResponseV2 ...
    type ProductSizePriceResponseV2 struct {

        
            Set ProductSetV2  `json:"set"`
            ItemType string  `json:"item_type"`
            IsGift bool  `json:"is_gift"`
            SellerCount float64  `json:"seller_count"`
            PricePerPiece ProductStockPriceV2  `json:"price_per_piece"`
            IsCod bool  `json:"is_cod"`
            Seller SellerV2  `json:"seller"`
            Store StoreV2  `json:"store"`
            ReturnConfig ReturnConfigSchemaV2  `json:"return_config"`
            Quantity float64  `json:"quantity"`
            Pincode float64  `json:"pincode"`
            StrategyWiseListing []StrategyWiseListingSchemaV2  `json:"strategy_wise_listing"`
            Price ProductStockPriceV2  `json:"price"`
            SpecialBadge string  `json:"special_badge"`
            Discount string  `json:"discount"`
            MarketplaceAttributes []MarketPlaceSttributesSchemaV2  `json:"marketplace_attributes"`
            PricePerUnit ProductStockUnitPriceV2  `json:"price_per_unit"`
            GroupedAttributes []SellerGroupAttributes  `json:"grouped_attributes"`
            ArticleID string  `json:"article_id"`
            ArticleAssignment ArticleAssignmentV2  `json:"article_assignment"`
            LongLat []float64  `json:"long_lat"`
         
    }
    
    // ProductSizeSellerFilterSchemaV2 ...
    type ProductSizeSellerFilterSchemaV2 struct {

        
            IsSelected bool  `json:"is_selected"`
            Value string  `json:"value"`
            Name string  `json:"name"`
         
    }
    
    // ProductSizeSellersResponseV2 ...
    type ProductSizeSellersResponseV2 struct {

        
            Page Page  `json:"page"`
            Items []ProductSizePriceResponseV2  `json:"items"`
            SortOn []ProductSizeSellerFilterSchemaV2  `json:"sort_on"`
         
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
    
    // BaseInfo ...
    type BaseInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
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

        
            Brand BaseInfo  `json:"brand"`
            Type string  `json:"type"`
            Images []ProductImage  `json:"images"`
            Categories []CategoryInfo  `json:"categories"`
            Action ProductAction  `json:"action"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
         
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

        
            Base BasePrice  `json:"base"`
            Converted BasePrice  `json:"converted"`
         
    }
    
    // ProductArticle ...
    type ProductArticle struct {

        
            Seller BaseInfo  `json:"seller"`
            Type string  `json:"type"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Size string  `json:"size"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            SellerIdentifier string  `json:"seller_identifier"`
            UID string  `json:"uid"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Identifier map[string]interface{}  `json:"identifier"`
            Price ArticlePriceInfo  `json:"price"`
            Quantity float64  `json:"quantity"`
            Store BaseInfo  `json:"store"`
         
    }
    
    // BuyRules ...
    type BuyRules struct {

        
            CartConditions map[string]interface{}  `json:"cart_conditions"`
            ItemCriteria map[string]interface{}  `json:"item_criteria"`
         
    }
    
    // FreeGiftItem ...
    type FreeGiftItem struct {

        
            ItemName string  `json:"item_name"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemID float64  `json:"item_id"`
            ItemSlug string  `json:"item_slug"`
            ItemImagesURL []string  `json:"item_images_url"`
         
    }
    
    // AppliedFreeArticles ...
    type AppliedFreeArticles struct {

        
            FreeGiftItemDetails FreeGiftItem  `json:"free_gift_item_details"`
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
         
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

        
            PromotionGroup string  `json:"promotion_group"`
            BuyRules []BuyRules  `json:"buy_rules"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            PromotionType string  `json:"promotion_type"`
            Amount float64  `json:"amount"`
            PromotionName string  `json:"promotion_name"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            PromoID string  `json:"promo_id"`
            MrpPromotion bool  `json:"mrp_promotion"`
            ArticleQuantity float64  `json:"article_quantity"`
            OfferText string  `json:"offer_text"`
         
    }
    
    // PromoMeta ...
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // CartProductIdentifer ...
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductAvailability ...
    type ProductAvailability struct {

        
            IsValid bool  `json:"is_valid"`
            Deliverable bool  `json:"deliverable"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            Sizes []string  `json:"sizes"`
            OutOfStock bool  `json:"out_of_stock"`
         
    }
    
    // ProductPrice ...
    type ProductPrice struct {

        
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            AddOn float64  `json:"add_on"`
            Selling float64  `json:"selling"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ProductPriceInfo ...
    type ProductPriceInfo struct {

        
            Base ProductPrice  `json:"base"`
            Converted ProductPrice  `json:"converted"`
         
    }
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            Product CartProduct  `json:"product"`
            IsSet bool  `json:"is_set"`
            Article ProductArticle  `json:"article"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Availability ProductAvailability  `json:"availability"`
            CouponMessage string  `json:"coupon_message"`
            Price ProductPriceInfo  `json:"price"`
            Message string  `json:"message"`
            Quantity float64  `json:"quantity"`
            Key string  `json:"key"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Discount string  `json:"discount"`
         
    }
    
    // CartCurrency ...
    type CartCurrency struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // PaymentSelectionLock ...
    type PaymentSelectionLock struct {

        
            DefaultOptions string  `json:"default_options"`
            PaymentIdentifier string  `json:"payment_identifier"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // RawBreakup ...
    type RawBreakup struct {

        
            CodCharge float64  `json:"cod_charge"`
            Vog float64  `json:"vog"`
            FyndCash float64  `json:"fynd_cash"`
            ConvenienceFee float64  `json:"convenience_fee"`
            GstCharges float64  `json:"gst_charges"`
            YouSaved float64  `json:"you_saved"`
            MrpTotal float64  `json:"mrp_total"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Total float64  `json:"total"`
            Coupon float64  `json:"coupon"`
            Discount float64  `json:"discount"`
            Subtotal float64  `json:"subtotal"`
         
    }
    
    // CouponBreakup ...
    type CouponBreakup struct {

        
            Type string  `json:"type"`
            Title string  `json:"title"`
            Value float64  `json:"value"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Code string  `json:"code"`
            SubTitle string  `json:"sub_title"`
            CouponType string  `json:"coupon_type"`
            UID string  `json:"uid"`
            IsApplied bool  `json:"is_applied"`
            Description string  `json:"description"`
            Message string  `json:"message"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            CouponValue float64  `json:"coupon_value"`
         
    }
    
    // LoyaltyPoints ...
    type LoyaltyPoints struct {

        
            Total float64  `json:"total"`
            Applicable float64  `json:"applicable"`
            IsApplied bool  `json:"is_applied"`
            Description string  `json:"description"`
         
    }
    
    // DisplayBreakup ...
    type DisplayBreakup struct {

        
            Value float64  `json:"value"`
            Message []string  `json:"message"`
            Key string  `json:"key"`
            CurrencySymbol string  `json:"currency_symbol"`
            Display string  `json:"display"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // CartBreakup ...
    type CartBreakup struct {

        
            Raw RawBreakup  `json:"raw"`
            Coupon CouponBreakup  `json:"coupon"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Display []DisplayBreakup  `json:"display"`
         
    }
    
    // CartDetailResponse ...
    type CartDetailResponse struct {

        
            IsValid bool  `json:"is_valid"`
            Gstin string  `json:"gstin"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Items []CartProductInfo  `json:"items"`
            Currency CartCurrency  `json:"currency"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            LastModified string  `json:"last_modified"`
            ID string  `json:"id"`
            Comment string  `json:"comment"`
            BuyNow bool  `json:"buy_now"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Message string  `json:"message"`
            CouponText string  `json:"coupon_text"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CheckoutMode string  `json:"checkout_mode"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            ProductGroupTags []string  `json:"product_group_tags"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            SellerID float64  `json:"seller_id"`
            StoreID float64  `json:"store_id"`
            ItemID float64  `json:"item_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ArticleID string  `json:"article_id"`
            Pos bool  `json:"pos"`
            Quantity float64  `json:"quantity"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            Display string  `json:"display"`
            ItemSize string  `json:"item_size"`
         
    }
    
    // AddCartRequest ...
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
         
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
            ItemID float64  `json:"item_id"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
            ItemSize string  `json:"item_size"`
         
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
    
    // DeleteCartDetailResponse ...
    type DeleteCartDetailResponse struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // CartItemCountResponse ...
    type CartItemCountResponse struct {

        
            UserCartItemsCount float64  `json:"user_cart_items_count"`
         
    }
    
    // Coupon ...
    type Coupon struct {

        
            CouponCode string  `json:"coupon_code"`
            Title string  `json:"title"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            IsApplicable bool  `json:"is_applicable"`
            SubTitle string  `json:"sub_title"`
            Message string  `json:"message"`
            CouponType string  `json:"coupon_type"`
            ExpiresOn string  `json:"expires_on"`
            IsApplied bool  `json:"is_applied"`
            Description string  `json:"description"`
            CouponValue float64  `json:"coupon_value"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
         
    }
    
    // PageCoupon ...
    type PageCoupon struct {

        
            HasPrevious bool  `json:"has_previous"`
            Current float64  `json:"current"`
            TotalItemCount float64  `json:"total_item_count"`
            Total float64  `json:"total"`
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
    
    // OfferSeller ...
    type OfferSeller struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // OfferPrice ...
    type OfferPrice struct {

        
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            BulkEffective float64  `json:"bulk_effective"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // OfferItem ...
    type OfferItem struct {

        
            Type string  `json:"type"`
            Best bool  `json:"best"`
            Price OfferPrice  `json:"price"`
            Quantity float64  `json:"quantity"`
            AutoApplied bool  `json:"auto_applied"`
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

        
            CountryCode string  `json:"country_code"`
            AreaCode string  `json:"area_code"`
            Area string  `json:"area"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            AreaCodeSlug string  `json:"area_code_slug"`
            UserID string  `json:"user_id"`
            Landmark string  `json:"landmark"`
            GeoLocation GeoLocation  `json:"geo_location"`
            CheckoutMode string  `json:"checkout_mode"`
            Meta map[string]interface{}  `json:"meta"`
            State string  `json:"state"`
            Country string  `json:"country"`
            IsActive bool  `json:"is_active"`
            ID string  `json:"id"`
            IsDefaultAddress bool  `json:"is_default_address"`
            Phone string  `json:"phone"`
            City string  `json:"city"`
            AddressType string  `json:"address_type"`
            Tags []string  `json:"tags"`
            Address string  `json:"address"`
            Name string  `json:"name"`
            Email string  `json:"email"`
         
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
            IsUpdated bool  `json:"is_updated"`
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
            AggregatorName string  `json:"aggregator_name"`
            MerchantCode string  `json:"merchant_code"`
            ID string  `json:"id"`
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CouponValidity ...
    type CouponValidity struct {

        
            DisplayMessageEn string  `json:"display_message_en"`
            Title string  `json:"title"`
            Valid bool  `json:"valid"`
            Code string  `json:"code"`
            Discount float64  `json:"discount"`
         
    }
    
    // PaymentCouponValidate ...
    type PaymentCouponValidate struct {

        
            Success bool  `json:"success"`
            CouponValidity CouponValidity  `json:"coupon_validity"`
            Message string  `json:"message"`
         
    }
    
    // ShipmentResponse ...
    type ShipmentResponse struct {

        
            FulfillmentID float64  `json:"fulfillment_id"`
            ShipmentType string  `json:"shipment_type"`
            Promise ShipmentPromise  `json:"promise"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            Shipments float64  `json:"shipments"`
            Items []CartProductInfo  `json:"items"`
            DpID string  `json:"dp_id"`
            FulfillmentType string  `json:"fulfillment_type"`
            OrderType string  `json:"order_type"`
            BoxType string  `json:"box_type"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            Gstin string  `json:"gstin"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Currency CartCurrency  `json:"currency"`
            LastModified string  `json:"last_modified"`
            ID string  `json:"id"`
            UID string  `json:"uid"`
            BuyNow bool  `json:"buy_now"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            IsValid bool  `json:"is_valid"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Shipments []ShipmentResponse  `json:"shipments"`
            Error bool  `json:"error"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Message string  `json:"message"`
            Comment string  `json:"comment"`
            CouponText string  `json:"coupon_text"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CheckoutMode string  `json:"checkout_mode"`
            CartID float64  `json:"cart_id"`
         
    }
    
    // StaffCheckout ...
    type StaffCheckout struct {

        
            ID string  `json:"_id"`
            FirstName string  `json:"first_name"`
            User string  `json:"user"`
            LastName string  `json:"last_name"`
         
    }
    
    // CartCheckoutDetailRequest ...
    type CartCheckoutDetailRequest struct {

        
            AddressID string  `json:"address_id"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            Meta map[string]interface{}  `json:"meta"`
            Staff StaffCheckout  `json:"staff"`
            MerchantCode string  `json:"merchant_code"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            BillingAddressID string  `json:"billing_address_id"`
            Aggregator string  `json:"aggregator"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            PaymentIdentifier string  `json:"payment_identifier"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            OrderingStore float64  `json:"ordering_store"`
            PaymentMode string  `json:"payment_mode"`
            CallbackURL string  `json:"callback_url"`
         
    }
    
    // CheckCart ...
    type CheckCart struct {

        
            Gstin string  `json:"gstin"`
            Items []CartProductInfo  `json:"items"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            LastModified string  `json:"last_modified"`
            UID string  `json:"uid"`
            StoreCode string  `json:"store_code"`
            BuyNow bool  `json:"buy_now"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Success bool  `json:"success"`
            IsValid bool  `json:"is_valid"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Message string  `json:"message"`
            Comment string  `json:"comment"`
            UserType string  `json:"user_type"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            CheckoutMode string  `json:"checkout_mode"`
            CartID float64  `json:"cart_id"`
            Currency CartCurrency  `json:"currency"`
            ID string  `json:"id"`
            OrderID string  `json:"order_id"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            CodAvailable bool  `json:"cod_available"`
            CodMessage string  `json:"cod_message"`
            CodCharges float64  `json:"cod_charges"`
            CouponText string  `json:"coupon_text"`
            ErrorMessage string  `json:"error_message"`
            DeliveryCharges float64  `json:"delivery_charges"`
         
    }
    
    // CartCheckoutResponse ...
    type CartCheckoutResponse struct {

        
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            AppInterceptURL string  `json:"app_intercept_url"`
            OrderID string  `json:"order_id"`
            Message string  `json:"message"`
            Cart CheckCart  `json:"cart"`
            CallbackURL string  `json:"callback_url"`
         
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

        
            Meta map[string]interface{}  `json:"meta"`
            User map[string]interface{}  `json:"user"`
            CreatedOn string  `json:"created_on"`
            Source map[string]interface{}  `json:"source"`
            Token string  `json:"token"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            Gstin string  `json:"gstin"`
            Items []CartProductInfo  `json:"items"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            LastModified string  `json:"last_modified"`
            UID string  `json:"uid"`
            BuyNow bool  `json:"buy_now"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            IsValid bool  `json:"is_valid"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Message string  `json:"message"`
            Comment string  `json:"comment"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CheckoutMode string  `json:"checkout_mode"`
            CartID float64  `json:"cart_id"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            Currency CartCurrency  `json:"currency"`
            ID string  `json:"id"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            CouponText string  `json:"coupon_text"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Cart SharedCart  `json:"cart"`
            Error string  `json:"error"`
         
    }
    
    // FreeGiftItems ...
    type FreeGiftItems struct {

        
            ItemName string  `json:"item_name"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemID float64  `json:"item_id"`
            ItemSlug string  `json:"item_slug"`
            ItemImagesURL []string  `json:"item_images_url"`
         
    }
    
    // PromotionOffer ...
    type PromotionOffer struct {

        
            PromotionGroup string  `json:"promotion_group"`
            ValidTill string  `json:"valid_till"`
            DiscountRules []map[string]interface{}  `json:"discount_rules"`
            ID string  `json:"id"`
            FreeGiftItems []FreeGiftItems  `json:"free_gift_items"`
            Description string  `json:"description"`
            BuyRules map[string]interface{}  `json:"buy_rules"`
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
    
    // LadderPrice ...
    type LadderPrice struct {

        
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            OfferPrice float64  `json:"offer_price"`
         
    }
    
    // LadderOfferItem ...
    type LadderOfferItem struct {

        
            Type string  `json:"type"`
            MinQuantity float64  `json:"min_quantity"`
            MaxQuantity float64  `json:"max_quantity"`
            Price LadderPrice  `json:"price"`
            Margin float64  `json:"margin"`
         
    }
    
    // LadderPriceOffer ...
    type LadderPriceOffer struct {

        
            PromotionGroup string  `json:"promotion_group"`
            CalculateOn string  `json:"calculate_on"`
            ValidTill string  `json:"valid_till"`
            DiscountRules []map[string]interface{}  `json:"discount_rules"`
            ID string  `json:"id"`
            FreeGiftItems []FreeGiftItems  `json:"free_gift_items"`
            OfferPrices []LadderOfferItem  `json:"offer_prices"`
            Description string  `json:"description"`
            BuyRules map[string]interface{}  `json:"buy_rules"`
            OfferText string  `json:"offer_text"`
         
    }
    
    // CurrencyInfo ...
    type CurrencyInfo struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // LadderPriceOffers ...
    type LadderPriceOffers struct {

        
            AvailableOffers []LadderPriceOffer  `json:"available_offers"`
            Currency CurrencyInfo  `json:"currency"`
         
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
    
    // TicketCategory ...
    type TicketCategory struct {

        
            Key string  `json:"key"`
            Display string  `json:"display"`
            Form CustomForm  `json:"form"`
            SubCategories []TicketSubCategory  `json:"sub_categories"`
            FeedbackForm TicketFeedbackForm  `json:"feedback_form"`
         
    }
    
    // TicketSubCategory ...
    type TicketSubCategory struct {

        
            Key string  `json:"key"`
            Display string  `json:"display"`
         
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
            ShowSupportDris bool  `json:"show_support_dris"`
            Integration map[string]interface{}  `json:"integration"`
         
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

        
            VerifyAPI string  `json:"verify_api"`
            MerchantID string  `json:"merchant_id"`
            Key string  `json:"key"`
            MerchantKey string  `json:"merchant_key"`
            Secret string  `json:"secret"`
            Pin string  `json:"pin"`
            UserID string  `json:"user_id"`
            API string  `json:"api"`
            Sdk bool  `json:"sdk"`
            ConfigType string  `json:"config_type"`
         
    }
    
    // AggregatorsConfigDetailResponse ...
    type AggregatorsConfigDetailResponse struct {

        
            Env string  `json:"env"`
            Juspay AggregatorConfigDetail  `json:"juspay"`
            Success bool  `json:"success"`
            Ccavenue AggregatorConfigDetail  `json:"ccavenue"`
            Razorpay AggregatorConfigDetail  `json:"razorpay"`
            Stripe AggregatorConfigDetail  `json:"stripe"`
            Simpl AggregatorConfigDetail  `json:"simpl"`
            Rupifi AggregatorConfigDetail  `json:"rupifi"`
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

        
            Refresh bool  `json:"refresh"`
            CardID string  `json:"card_id"`
            Nickname string  `json:"nickname"`
            NameOnCard string  `json:"name_on_card"`
         
    }
    
    // AttachCardsResponse ...
    type AttachCardsResponse struct {

        
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
            Message string  `json:"message"`
         
    }
    
    // CardPaymentGateway ...
    type CardPaymentGateway struct {

        
            Aggregator string  `json:"aggregator"`
            CustomerID string  `json:"customer_id"`
            API string  `json:"api"`
         
    }
    
    // ActiveCardPaymentGatewayResponse ...
    type ActiveCardPaymentGatewayResponse struct {

        
            Success bool  `json:"success"`
            Cards CardPaymentGateway  `json:"cards"`
            Message string  `json:"message"`
         
    }
    
    // Card ...
    type Card struct {

        
            CardToken string  `json:"card_token"`
            CardReference string  `json:"card_reference"`
            CardIssuer string  `json:"card_issuer"`
            CardName string  `json:"card_name"`
            Expired bool  `json:"expired"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardID string  `json:"card_id"`
            CardType string  `json:"card_type"`
            ExpYear float64  `json:"exp_year"`
            AggregatorName string  `json:"aggregator_name"`
            CardFingerprint string  `json:"card_fingerprint"`
            CardBrandImage string  `json:"card_brand_image"`
            CardIsin string  `json:"card_isin"`
            CardNumber string  `json:"card_number"`
            Nickname string  `json:"nickname"`
            ExpMonth float64  `json:"exp_month"`
            CardBrand string  `json:"card_brand"`
         
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

        
            BillingAddress map[string]interface{}  `json:"billing_address"`
            TransactionAmountInPaise float64  `json:"transaction_amount_in_paise"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            Aggregator string  `json:"aggregator"`
            Payload string  `json:"payload"`
            OrderItems []map[string]interface{}  `json:"order_items"`
            MerchantParams map[string]interface{}  `json:"merchant_params"`
            PhoneNumber string  `json:"phone_number"`
         
    }
    
    // ValidateCustomerResponse ...
    type ValidateCustomerResponse struct {

        
            Success bool  `json:"success"`
            Data map[string]interface{}  `json:"data"`
            Message string  `json:"message"`
         
    }
    
    // ChargeCustomerRequest ...
    type ChargeCustomerRequest struct {

        
            OrderID string  `json:"order_id"`
            Verified bool  `json:"verified"`
            Amount float64  `json:"amount"`
            TransactionToken string  `json:"transaction_token"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // ChargeCustomerResponse ...
    type ChargeCustomerResponse struct {

        
            OrderID string  `json:"order_id"`
            CartID string  `json:"cart_id"`
            Success bool  `json:"success"`
            Aggregator string  `json:"aggregator"`
            Status string  `json:"status"`
            DeliveryAddressID string  `json:"delivery_address_id"`
            Message string  `json:"message"`
         
    }
    
    // PaymentInitializationRequest ...
    type PaymentInitializationRequest struct {

        
            OrderID string  `json:"order_id"`
            Contact string  `json:"contact"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Amount float64  `json:"amount"`
            Timeout float64  `json:"timeout"`
            Method string  `json:"method"`
            CustomerID string  `json:"customer_id"`
            Aggregator string  `json:"aggregator"`
            Vpa string  `json:"vpa"`
            Email string  `json:"email"`
            Currency string  `json:"currency"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
         
    }
    
    // PaymentInitializationResponse ...
    type PaymentInitializationResponse struct {

        
            MerchantOrderID string  `json:"merchant_order_id"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            Amount float64  `json:"amount"`
            Timeout float64  `json:"timeout"`
            Success bool  `json:"success"`
            Method string  `json:"method"`
            CustomerID string  `json:"customer_id"`
            VirtualID string  `json:"virtual_id"`
            BqrImage string  `json:"bqr_image"`
            Aggregator string  `json:"aggregator"`
            PollingURL string  `json:"polling_url"`
            UpiPollURL string  `json:"upi_poll_url"`
            Vpa string  `json:"vpa"`
            Currency string  `json:"currency"`
            Status string  `json:"status"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
         
    }
    
    // PaymentStatusUpdateRequest ...
    type PaymentStatusUpdateRequest struct {

        
            OrderID string  `json:"order_id"`
            Contact string  `json:"contact"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Amount float64  `json:"amount"`
            Method string  `json:"method"`
            CustomerID string  `json:"customer_id"`
            Aggregator string  `json:"aggregator"`
            Vpa string  `json:"vpa"`
            Email string  `json:"email"`
            Currency string  `json:"currency"`
            Status string  `json:"status"`
         
    }
    
    // PaymentStatusUpdateResponse ...
    type PaymentStatusUpdateResponse struct {

        
            RedirectURL string  `json:"redirect_url"`
            Retry bool  `json:"retry"`
            Success bool  `json:"success"`
            AggregatorName string  `json:"aggregator_name"`
            Status string  `json:"status"`
         
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

        
            Code string  `json:"code"`
            PackageName string  `json:"package_name"`
            Logos PaymentModeLogo  `json:"logos"`
            DisplayName string  `json:"display_name"`
         
    }
    
    // PaymentModeList ...
    type PaymentModeList struct {

        
            MerchantCode string  `json:"merchant_code"`
            CardFingerprint string  `json:"card_fingerprint"`
            Nickname string  `json:"nickname"`
            Name string  `json:"name"`
            IntentFlow bool  `json:"intent_flow"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            CardNumber string  `json:"card_number"`
            FyndVpa string  `json:"fynd_vpa"`
            ExpMonth float64  `json:"exp_month"`
            Expired bool  `json:"expired"`
            CardToken string  `json:"card_token"`
            CardName string  `json:"card_name"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            RetryCount float64  `json:"retry_count"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            Timeout float64  `json:"timeout"`
            CardID string  `json:"card_id"`
            DisplayPriority float64  `json:"display_priority"`
            DisplayName string  `json:"display_name"`
            CardType string  `json:"card_type"`
            ExpYear float64  `json:"exp_year"`
            CardIsin string  `json:"card_isin"`
            IntentApp []IntentApp  `json:"intent_app"`
            CardIssuer string  `json:"card_issuer"`
            CardReference string  `json:"card_reference"`
            Code string  `json:"code"`
            CardBrandImage string  `json:"card_brand_image"`
            AggregatorName string  `json:"aggregator_name"`
            CardBrand string  `json:"card_brand"`
         
    }
    
    // RootPaymentMode ...
    type RootPaymentMode struct {

        
            AnonymousEnable bool  `json:"anonymous_enable"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            DisplayPriority float64  `json:"display_priority"`
            DisplayName string  `json:"display_name"`
            AggregatorName string  `json:"aggregator_name"`
            List []PaymentModeList  `json:"list"`
            Name string  `json:"name"`
            SaveCard bool  `json:"save_card"`
         
    }
    
    // AggregatorRoute ...
    type AggregatorRoute struct {

        
            PaymentFlowData map[string]interface{}  `json:"payment_flow_data"`
            PaymentFlow string  `json:"payment_flow"`
            Data map[string]interface{}  `json:"data"`
            APILink string  `json:"api_link"`
         
    }
    
    // PaymentFlow ...
    type PaymentFlow struct {

        
            BqrRazorpay AggregatorRoute  `json:"bqr_razorpay"`
            Juspay AggregatorRoute  `json:"juspay"`
            Jiopay AggregatorRoute  `json:"jiopay"`
            Epaylater AggregatorRoute  `json:"epaylater"`
            Ccavenue AggregatorRoute  `json:"ccavenue"`
            Razorpay AggregatorRoute  `json:"razorpay"`
            Payubiz AggregatorRoute  `json:"payubiz"`
            Stripe AggregatorRoute  `json:"stripe"`
            UpiRazorpay AggregatorRoute  `json:"upi_razorpay"`
            Simpl AggregatorRoute  `json:"simpl"`
            Rupifi AggregatorRoute  `json:"rupifi"`
            Fynd AggregatorRoute  `json:"fynd"`
            Mswipe AggregatorRoute  `json:"mswipe"`
         
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
            Status string  `json:"status"`
            Message string  `json:"message"`
         
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

        
            Status string  `json:"status"`
            IsValid bool  `json:"is_valid"`
            UpiVpa string  `json:"upi_vpa"`
            CustomerName string  `json:"customer_name"`
         
    }
    
    // ValidateVPAResponse ...
    type ValidateVPAResponse struct {

        
            Success bool  `json:"success"`
            Data ValidateUPI  `json:"data"`
         
    }
    
    // TransferItemsDetails ...
    type TransferItemsDetails struct {

        
            DisplayName string  `json:"display_name"`
            Name string  `json:"name"`
            ID float64  `json:"id"`
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

        
            Address string  `json:"address"`
            ID float64  `json:"id"`
            Email string  `json:"email"`
            DelightsUserName string  `json:"delights_user_name"`
            Subtitle string  `json:"subtitle"`
            AccountNo string  `json:"account_no"`
            TransferMode string  `json:"transfer_mode"`
            AccountHolder string  `json:"account_holder"`
            CreatedOn string  `json:"created_on"`
            Mobile string  `json:"mobile"`
            ModifiedOn string  `json:"modified_on"`
            Title string  `json:"title"`
            BeneficiaryID string  `json:"beneficiary_id"`
            IfscCode string  `json:"ifsc_code"`
            DisplayName string  `json:"display_name"`
            Comment string  `json:"comment"`
            BankName string  `json:"bank_name"`
            IsActive bool  `json:"is_active"`
            BranchName string  `json:"branch_name"`
         
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

        
            Success bool  `json:"success"`
            BranchName string  `json:"branch_name"`
            BankName string  `json:"bank_name"`
         
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

        
            Success string  `json:"success"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Description string  `json:"description"`
         
    }
    
    // BeneficiaryModeDetails ...
    type BeneficiaryModeDetails struct {

        
            Address string  `json:"address"`
            AccountNo string  `json:"account_no"`
            BankName string  `json:"bank_name"`
            AccountHolder string  `json:"account_holder"`
            IfscCode string  `json:"ifsc_code"`
            Wallet string  `json:"wallet"`
            Comment string  `json:"comment"`
            BranchName string  `json:"branch_name"`
            Mobile string  `json:"mobile"`
            Email string  `json:"email"`
            Vpa string  `json:"vpa"`
         
    }
    
    // AddBeneficiaryDetailsRequest ...
    type AddBeneficiaryDetailsRequest struct {

        
            OrderID string  `json:"order_id"`
            TransferMode string  `json:"transfer_mode"`
            RequestID string  `json:"request_id"`
            Details BeneficiaryModeDetails  `json:"details"`
            Delights bool  `json:"delights"`
            Otp string  `json:"otp"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // RefundAccountResponse ...
    type RefundAccountResponse struct {

        
            Success bool  `json:"success"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Data map[string]interface{}  `json:"data"`
            Message string  `json:"message"`
         
    }
    
    // BankDetailsForOTP ...
    type BankDetailsForOTP struct {

        
            BankName string  `json:"bank_name"`
            AccountNo string  `json:"account_no"`
            AccountHolder string  `json:"account_holder"`
            IfscCode string  `json:"ifsc_code"`
            BranchName string  `json:"branch_name"`
         
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
            StatusCode float64  `json:"status_code"`
            PaymentLinkCurrentStatus string  `json:"payment_link_current_status"`
            MerchantName string  `json:"merchant_name"`
            Amount float64  `json:"amount"`
            Success bool  `json:"success"`
            PollingTimeout float64  `json:"polling_timeout"`
            ExternalOrderID string  `json:"external_order_id"`
            Message string  `json:"message"`
         
    }
    
    // ErrorDescription ...
    type ErrorDescription struct {

        
            MerchantOrderID string  `json:"merchant_order_id"`
            MerchantName string  `json:"merchant_name"`
            Amount float64  `json:"amount"`
            InvalidID bool  `json:"invalid_id"`
            Msg string  `json:"msg"`
            PaymentTransactionID string  `json:"payment_transaction_id"`
            Cancelled bool  `json:"cancelled"`
            Expired bool  `json:"expired"`
         
    }
    
    // CreatePaymentLinkMeta ...
    type CreatePaymentLinkMeta struct {

        
            CartID string  `json:"cart_id"`
            Pincode string  `json:"pincode"`
            Amount string  `json:"amount"`
            CheckoutMode string  `json:"checkout_mode"`
            AssignCardID string  `json:"assign_card_id"`
         
    }
    
    // CreatePaymentLinkRequest ...
    type CreatePaymentLinkRequest struct {

        
            MobileNumber string  `json:"mobile_number"`
            Amount float64  `json:"amount"`
            Description string  `json:"description"`
            Meta CreatePaymentLinkMeta  `json:"meta"`
            Email string  `json:"email"`
            ExternalOrderID string  `json:"external_order_id"`
         
    }
    
    // CreatePaymentLinkResponse ...
    type CreatePaymentLinkResponse struct {

        
            PaymentLinkURL string  `json:"payment_link_url"`
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
            PaymentLinkID string  `json:"payment_link_id"`
            PollingTimeout float64  `json:"polling_timeout"`
            Message string  `json:"message"`
         
    }
    
    // CancelOrResendPaymentLinkRequest ...
    type CancelOrResendPaymentLinkRequest struct {

        
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // ResendPaymentLinkResponse ...
    type ResendPaymentLinkResponse struct {

        
            PollingTimeout float64  `json:"polling_timeout"`
            Success bool  `json:"success"`
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

        
            RedirectURL string  `json:"redirect_url"`
            OrderID string  `json:"order_id"`
            StatusCode float64  `json:"status_code"`
            Amount float64  `json:"amount"`
            Success bool  `json:"success"`
            PaymentLinkID string  `json:"payment_link_id"`
            AggregatorName string  `json:"aggregator_name"`
            HttpStatus float64  `json:"http_status"`
            Status string  `json:"status"`
            Message string  `json:"message"`
         
    }
    
    // PaymentMethodsMeta ...
    type PaymentMethodsMeta struct {

        
            PaymentGateway string  `json:"payment_gateway"`
            MerchantCode string  `json:"merchant_code"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // CreateOrderUserPaymentMethods ...
    type CreateOrderUserPaymentMethods struct {

        
            Meta PaymentMethodsMeta  `json:"meta"`
            Mode string  `json:"mode"`
            Name string  `json:"name"`
         
    }
    
    // CreateOrderUserRequest ...
    type CreateOrderUserRequest struct {

        
            SuccessCallbackURL string  `json:"success_callback_url"`
            PaymentMethods CreateOrderUserPaymentMethods  `json:"payment_methods"`
            FailureCallbackURL string  `json:"failure_callback_url"`
            PaymentLinkID string  `json:"payment_link_id"`
            Meta map[string]interface{}  `json:"meta"`
            Currency string  `json:"currency"`
         
    }
    
    // CreateOrderUserData ...
    type CreateOrderUserData struct {

        
            OrderID string  `json:"order_id"`
            Contact string  `json:"contact"`
            CallbackURL string  `json:"callback_url"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Amount float64  `json:"amount"`
            Method string  `json:"method"`
            CustomerID string  `json:"customer_id"`
            Aggregator string  `json:"aggregator"`
            Email string  `json:"email"`
            Currency string  `json:"currency"`
         
    }
    
    // CreateOrderUserResponse ...
    type CreateOrderUserResponse struct {

        
            OrderID string  `json:"order_id"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            CallbackURL string  `json:"callback_url"`
            StatusCode float64  `json:"status_code"`
            Data CreateOrderUserData  `json:"data"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // BalanceDetails ...
    type BalanceDetails struct {

        
            Currency string  `json:"currency"`
            FormattedValue string  `json:"formatted_value"`
            Value float64  `json:"value"`
         
    }
    
    // CreditSummary ...
    type CreditSummary struct {

        
            Status string  `json:"status"`
            Balance BalanceDetails  `json:"balance"`
            MerchantCustomerRefID string  `json:"merchant_customer_ref_id"`
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
            Status bool  `json:"status"`
            IsRegistered bool  `json:"is_registered"`
         
    }
    
    // CheckCreditResponse ...
    type CheckCreditResponse struct {

        
            Success bool  `json:"success"`
            Data CreditDetail  `json:"data"`
         
    }
    
    // DeviceDetails ...
    type DeviceDetails struct {

        
            DeviceMake string  `json:"device_make"`
            DeviceModel string  `json:"device_model"`
            IdentificationNumber string  `json:"identification_number"`
            IdentifierType string  `json:"identifier_type"`
            Os string  `json:"os"`
            DeviceType string  `json:"device_type"`
            OsVersion string  `json:"os_version"`
         
    }
    
    // KYCAddress ...
    type KYCAddress struct {

        
            Addressline2 string  `json:"addressline2"`
            Addressline1 string  `json:"addressline1"`
            Pincode string  `json:"pincode"`
            OwnershipType string  `json:"ownership_type"`
            State string  `json:"state"`
            LandMark string  `json:"land_mark"`
            City string  `json:"city"`
         
    }
    
    // UserPersonalInfoInDetails ...
    type UserPersonalInfoInDetails struct {

        
            MothersName string  `json:"mothers_name"`
            Dob string  `json:"dob"`
            Passport string  `json:"passport"`
            DrivingLicense string  `json:"driving_license"`
            MiddleName string  `json:"middle_name"`
            VoterID string  `json:"voter_id"`
            MobileVerified bool  `json:"mobile_verified"`
            Gender string  `json:"gender"`
            Pan string  `json:"pan"`
            AddressAsPerID KYCAddress  `json:"address_as_per_id"`
            Phone string  `json:"phone"`
            EmailVerified bool  `json:"email_verified"`
            Email string  `json:"email"`
            FirstName string  `json:"first_name"`
            FathersName string  `json:"fathers_name"`
            LastName string  `json:"last_name"`
         
    }
    
    // MarketplaceInfo ...
    type MarketplaceInfo struct {

        
            MembershipID string  `json:"membership_id"`
            Name string  `json:"name"`
            DateOfJoining string  `json:"date_of_joining"`
         
    }
    
    // BusinessDetails ...
    type BusinessDetails struct {

        
            BusinessOwnershipType string  `json:"business_ownership_type"`
            BusinessType string  `json:"business_type"`
            Address KYCAddress  `json:"address"`
            Fssai string  `json:"fssai"`
            ShopAndEstablishment map[string]interface{}  `json:"shop_and_establishment"`
            Vintage string  `json:"vintage"`
            Name string  `json:"name"`
            Gstin string  `json:"gstin"`
            EntityType string  `json:"entity_type"`
            Fda string  `json:"fda"`
            Pan string  `json:"pan"`
         
    }
    
    // CustomerOnboardingRequest ...
    type CustomerOnboardingRequest struct {

        
            Mcc string  `json:"mcc"`
            Device DeviceDetails  `json:"device"`
            PersonalInfo UserPersonalInfoInDetails  `json:"personal_info"`
            Aggregator string  `json:"aggregator"`
            MarketplaceInfo MarketplaceInfo  `json:"marketplace_info"`
            Source string  `json:"source"`
            BusinessInfo BusinessDetails  `json:"business_info"`
         
    }
    
    // OnboardSummary ...
    type OnboardSummary struct {

        
            RedirectURL string  `json:"redirect_url"`
            Session map[string]interface{}  `json:"session"`
            Status bool  `json:"status"`
         
    }
    
    // CustomerOnboardingResponse ...
    type CustomerOnboardingResponse struct {

        
            Success bool  `json:"success"`
            Data OnboardSummary  `json:"data"`
         
    }
    

    
    // OrderStatuses ...
    type OrderStatuses struct {

        
            IsSelected bool  `json:"is_selected"`
            Display string  `json:"display"`
            Value float64  `json:"value"`
         
    }
    
    // OrderFilters ...
    type OrderFilters struct {

        
            Statuses []OrderStatuses  `json:"statuses"`
         
    }
    
    // OrderPage ...
    type OrderPage struct {

        
            ItemTotal float64  `json:"item_total"`
            Current float64  `json:"current"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
         
    }
    
    // UserInfo ...
    type UserInfo struct {

        
            Gender string  `json:"gender"`
            Mobile string  `json:"mobile"`
            Email string  `json:"email"`
            Name string  `json:"name"`
         
    }
    
    // BreakupValues ...
    type BreakupValues struct {

        
            Display string  `json:"display"`
            Name string  `json:"name"`
            Value float64  `json:"value"`
         
    }
    
    // BagsForReorderArticleAssignment ...
    type BagsForReorderArticleAssignment struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // BagsForReorder ...
    type BagsForReorder struct {

        
            ItemSize string  `json:"item_size"`
            ArticleAssignment BagsForReorderArticleAssignment  `json:"article_assignment"`
            ItemID float64  `json:"item_id"`
            Quantity float64  `json:"quantity"`
            StoreID float64  `json:"store_id"`
            SellerID float64  `json:"seller_id"`
         
    }
    
    // ShipmentUserInfo ...
    type ShipmentUserInfo struct {

        
            Gender string  `json:"gender"`
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            Mobile string  `json:"mobile"`
         
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

        
            Status string  `json:"status"`
            TrackingDetails []NestedTrackingDetails  `json:"tracking_details"`
            IsPassed bool  `json:"is_passed"`
            Time string  `json:"time"`
            IsCurrent bool  `json:"is_current"`
         
    }
    
    // ShipmentStatus ...
    type ShipmentStatus struct {

        
            Title string  `json:"title"`
            HexCode string  `json:"hex_code"`
         
    }
    
    // FulfillingStore ...
    type FulfillingStore struct {

        
            ID float64  `json:"id"`
            CompanyName string  `json:"company_name"`
            Code string  `json:"code"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
         
    }
    
    // DeliveryAddress ...
    type DeliveryAddress struct {

        
            Address1 string  `json:"address1"`
            Phone string  `json:"phone"`
            AddressCategory string  `json:"address_category"`
            Version string  `json:"version"`
            Pincode string  `json:"pincode"`
            Landmark string  `json:"landmark"`
            Address string  `json:"address"`
            Area string  `json:"area"`
            CreatedAt string  `json:"created_at"`
            Email string  `json:"email"`
            State string  `json:"state"`
            Name string  `json:"name"`
            City string  `json:"city"`
            AddressType string  `json:"address_type"`
            Address2 string  `json:"address2"`
            Country string  `json:"country"`
            UpdatedAt string  `json:"updated_at"`
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
            ContactPerson string  `json:"contact_person"`
         
    }
    
    // Prices ...
    type Prices struct {

        
            Discount float64  `json:"discount"`
            CashbackApplied float64  `json:"cashback_applied"`
            DeliveryCharge float64  `json:"delivery_charge"`
            ValueOfGood float64  `json:"value_of_good"`
            RefundCredit float64  `json:"refund_credit"`
            PriceMarked float64  `json:"price_marked"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            AmountPaid float64  `json:"amount_paid"`
            CouponValue float64  `json:"coupon_value"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            PriceEffective float64  `json:"price_effective"`
            CodCharges float64  `json:"cod_charges"`
            RefundAmount float64  `json:"refund_amount"`
            TransferPrice float64  `json:"transfer_price"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            FyndCredits float64  `json:"fynd_credits"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            Cashback float64  `json:"cashback"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
         
    }
    
    // Invoice ...
    type Invoice struct {

        
            InvoiceURL string  `json:"invoice_url"`
            UpdatedDate string  `json:"updated_date"`
            LabelURL string  `json:"label_url"`
         
    }
    
    // ItemBrand ...
    type ItemBrand struct {

        
            Logo string  `json:"logo"`
            Name string  `json:"name"`
         
    }
    
    // Item ...
    type Item struct {

        
            SlugKey string  `json:"slug_key"`
            Image []string  `json:"image"`
            SellerIdentifier string  `json:"seller_identifier"`
            ID float64  `json:"id"`
            Code string  `json:"code"`
            Size string  `json:"size"`
            Brand ItemBrand  `json:"brand"`
            Name string  `json:"name"`
         
    }
    
    // AppliedPromos ...
    type AppliedPromos struct {

        
            PromotionName string  `json:"promotion_name"`
            MrpPromotion bool  `json:"mrp_promotion"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromotionType string  `json:"promotion_type"`
            Amount float64  `json:"amount"`
            PromoID string  `json:"promo_id"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
         
    }
    
    // CurrentStatus ...
    type CurrentStatus struct {

        
            Status string  `json:"status"`
            JourneyType string  `json:"journey_type"`
            UpdatedAt string  `json:"updated_at"`
            Name string  `json:"name"`
         
    }
    
    // Identifiers ...
    type Identifiers struct {

        
            SkuCode string  `json:"sku_code"`
            Ean string  `json:"ean"`
         
    }
    
    // FinancialBreakup ...
    type FinancialBreakup struct {

        
            Discount float64  `json:"discount"`
            CashbackApplied float64  `json:"cashback_applied"`
            HsnCode string  `json:"hsn_code"`
            DeliveryCharge float64  `json:"delivery_charge"`
            ValueOfGood float64  `json:"value_of_good"`
            RefundCredit float64  `json:"refund_credit"`
            GstFee float64  `json:"gst_fee"`
            PriceMarked float64  `json:"price_marked"`
            Identifiers Identifiers  `json:"identifiers"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            AmountPaid float64  `json:"amount_paid"`
            CouponValue float64  `json:"coupon_value"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            GstTag string  `json:"gst_tag"`
            Size string  `json:"size"`
            PriceEffective float64  `json:"price_effective"`
            CodCharges float64  `json:"cod_charges"`
            RefundAmount float64  `json:"refund_amount"`
            TotalUnits float64  `json:"total_units"`
            TransferPrice float64  `json:"transfer_price"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            FyndCredits float64  `json:"fynd_credits"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            Cashback float64  `json:"cashback"`
            ItemName string  `json:"item_name"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
         
    }
    
    // Bags ...
    type Bags struct {

        
            LineNumber float64  `json:"line_number"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            CanReturn bool  `json:"can_return"`
            Prices Prices  `json:"prices"`
            CanCancel bool  `json:"can_cancel"`
            ID float64  `json:"id"`
            SellerIdentifier string  `json:"seller_identifier"`
            Quantity float64  `json:"quantity"`
            DeliveryDate string  `json:"delivery_date"`
            Item Item  `json:"item"`
            ReturnableDate string  `json:"returnable_date"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
         
    }
    
    // ShipmentTotalDetails ...
    type ShipmentTotalDetails struct {

        
            Sizes float64  `json:"sizes"`
            Pieces float64  `json:"pieces"`
            TotalPrice float64  `json:"total_price"`
         
    }
    
    // ShipmentPayment ...
    type ShipmentPayment struct {

        
            Mop string  `json:"mop"`
            Logo string  `json:"logo"`
            PaymentMode string  `json:"payment_mode"`
            Status string  `json:"status"`
            Mode string  `json:"mode"`
            DisplayName string  `json:"display_name"`
         
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
    
    // FulfillingCompany ...
    type FulfillingCompany struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
    }
    
    // Shipments ...
    type Shipments struct {

        
            UserInfo ShipmentUserInfo  `json:"user_info"`
            TrackingDetails []TrackingDetails  `json:"tracking_details"`
            NeedHelpURL string  `json:"need_help_url"`
            BeneficiaryDetails bool  `json:"beneficiary_details"`
            DpName string  `json:"dp_name"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            TotalBags float64  `json:"total_bags"`
            DeliveryAddress DeliveryAddress  `json:"delivery_address"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            Prices Prices  `json:"prices"`
            Invoice Invoice  `json:"invoice"`
            ReturnableDate string  `json:"returnable_date"`
            Bags []Bags  `json:"bags"`
            TotalDetails ShipmentTotalDetails  `json:"total_details"`
            DeliveryDate string  `json:"delivery_date"`
            TrakingNo string  `json:"traking_no"`
            CanBreak map[string]interface{}  `json:"can_break"`
            AwbNo string  `json:"awb_no"`
            OrderID string  `json:"order_id"`
            CanReturn bool  `json:"can_return"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            Payment ShipmentPayment  `json:"payment"`
            CanCancel bool  `json:"can_cancel"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            OrderType string  `json:"order_type"`
            SizeInfo map[string]interface{}  `json:"size_info"`
            ShowDownloadInvoice bool  `json:"show_download_invoice"`
            ShowTrackLink bool  `json:"show_track_link"`
            Promise Promise  `json:"promise"`
            TrackURL string  `json:"track_url"`
            ShipmentID string  `json:"shipment_id"`
            Comment string  `json:"comment"`
            FulfillingCompany FulfillingCompany  `json:"fulfilling_company"`
            RefundDetails map[string]interface{}  `json:"refund_details"`
         
    }
    
    // OrderSchema ...
    type OrderSchema struct {

        
            UserInfo UserInfo  `json:"user_info"`
            OrderID string  `json:"order_id"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            OrderCreatedTime string  `json:"order_created_time"`
            BagsForReorder []BagsForReorder  `json:"bags_for_reorder"`
            Shipments []Shipments  `json:"shipments"`
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
         
    }
    
    // OrderList ...
    type OrderList struct {

        
            Filters OrderFilters  `json:"filters"`
            Page OrderPage  `json:"page"`
            Items []OrderSchema  `json:"items"`
         
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

        
            PresignedType string  `json:"presigned_type"`
            Success bool  `json:"success"`
            PresignedURL string  `json:"presigned_url"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // Track ...
    type Track struct {

        
            Status string  `json:"status"`
            Reason string  `json:"reason"`
            UpdatedTime string  `json:"updated_time"`
            ShipmentType string  `json:"shipment_type"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            Awb string  `json:"awb"`
            UpdatedAt string  `json:"updated_at"`
            AccountName string  `json:"account_name"`
         
    }
    
    // ShipmentTrack ...
    type ShipmentTrack struct {

        
            Results []Track  `json:"results"`
         
    }
    
    // CustomerDetailsResponse ...
    type CustomerDetailsResponse struct {

        
            OrderID string  `json:"order_id"`
            Phone string  `json:"phone"`
            ShipmentID string  `json:"shipment_id"`
            Country string  `json:"country"`
            Name string  `json:"name"`
         
    }
    
    // SendOtpToCustomerResponse ...
    type SendOtpToCustomerResponse struct {

        
            RequestID string  `json:"request_id"`
            Success bool  `json:"success"`
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

        
            QuestionSet []QuestionSet  `json:"question_set"`
            Meta BagReasonMeta  `json:"meta"`
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
            Reasons []BagReasons  `json:"reasons"`
            QcType []string  `json:"qc_type"`
         
    }
    
    // ShipmentBagReasons ...
    type ShipmentBagReasons struct {

        
            Success bool  `json:"success"`
            Reasons []BagReasons  `json:"reasons"`
         
    }
    
    // ShipmentReason ...
    type ShipmentReason struct {

        
            Priority float64  `json:"priority"`
            ReasonText string  `json:"reason_text"`
            ShowTextArea bool  `json:"show_text_area"`
            Flow string  `json:"flow"`
            ReasonID float64  `json:"reason_id"`
            FeedbackType string  `json:"feedback_type"`
         
    }
    
    // ShipmentReasons ...
    type ShipmentReasons struct {

        
            Reasons []ShipmentReason  `json:"reasons"`
         
    }
    
    // Products ...
    type Products struct {

        
            LineNumber float64  `json:"line_number"`
            Identifier string  `json:"identifier"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // ProductsReasonsData ...
    type ProductsReasonsData struct {

        
            ReasonText string  `json:"reason_text"`
            ReasonID float64  `json:"reason_id"`
         
    }
    
    // ProductsReasonsFilters ...
    type ProductsReasonsFilters struct {

        
            LineNumber float64  `json:"line_number"`
            Identifier string  `json:"identifier"`
            Quantity float64  `json:"quantity"`
         
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

        
            Filters []map[string]interface{}  `json:"filters"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // DataUpdates ...
    type DataUpdates struct {

        
            Products []ProductsDataUpdates1  `json:"products"`
            Entities []EntitiesDataUpdates  `json:"entities"`
         
    }
    
    // ShipmentsRequest ...
    type ShipmentsRequest struct {

        
            Products []Products  `json:"products"`
            Identifier string  `json:"identifier"`
            Reasons ReasonsData  `json:"reasons"`
            DataUpdates DataUpdates  `json:"data_updates"`
         
    }
    
    // StatuesRequest ...
    type StatuesRequest struct {

        
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
            Status string  `json:"status"`
            Shipments []ShipmentsRequest1  `json:"shipments"`
         
    }
    
    // UpdateShipmentStatusRequest ...
    type UpdateShipmentStatusRequest struct {

        
            LockAfterTransition bool  `json:"lock_after_transition"`
            ForceTransition bool  `json:"force_transition"`
            Task bool  `json:"task"`
            Statuses []StatuesRequest  `json:"statuses"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
         
    }
    
    // StatusesBodyResponse ...
    type StatusesBodyResponse struct {

        
            Shipments []map[string]interface{}  `json:"shipments"`
         
    }
    
    // ShipmentApplicationStatusResponse ...
    type ShipmentApplicationStatusResponse struct {

        
            Statuses []StatusesBodyResponse  `json:"statuses"`
         
    }
    
    // Products1 ...
    type Products1 struct {

        
            Quantity float64  `json:"quantity"`
            LineNumber float64  `json:"line_number"`
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
    
    // ProductsReasonsFilters1 ...
    type ProductsReasonsFilters1 struct {

        
            Quantity float64  `json:"quantity"`
            LineNumber float64  `json:"line_number"`
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductsReasonsData1 ...
    type ProductsReasonsData1 struct {

        
            ReasonID float64  `json:"reason_id"`
            ReasonText string  `json:"reason_text"`
         
    }
    
    // ProductsReasons1 ...
    type ProductsReasons1 struct {

        
            Filters []ProductsReasonsFilters1  `json:"filters"`
            Data ProductsReasonsData1  `json:"data"`
         
    }
    
    // EntityReasonData1 ...
    type EntityReasonData1 struct {

        
            ReasonID float64  `json:"reason_id"`
            ReasonText string  `json:"reason_text"`
         
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
    
    // ShipmentsRequest1 ...
    type ShipmentsRequest1 struct {

        
            Products []Products1  `json:"products"`
            DataUpdates DataUpdates  `json:"data_updates"`
            Identifier string  `json:"identifier"`
            Reasons ReasonsData1  `json:"reasons"`
         
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
            ShipmentType string  `json:"shipment_type"`
            ArticleUID string  `json:"article_uid"`
         
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

        
            BillingAddressID string  `json:"billing_address_id"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            PaymentIdentifier string  `json:"payment_identifier"`
            AddressID string  `json:"address_id"`
            PaymentMode string  `json:"payment_mode"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            Meta map[string]interface{}  `json:"meta"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            OrderingStore float64  `json:"ordering_store"`
            MerchantCode string  `json:"merchant_code"`
            Aggregator string  `json:"aggregator"`
            Files []Files  `json:"files"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            Staff StaffCheckout  `json:"staff"`
            OrderType string  `json:"order_type"`
            CallbackURL string  `json:"callback_url"`
            Pos bool  `json:"pos"`
         
    }
    
    // CartDeliveryModesResponse ...
    type CartDeliveryModesResponse struct {

        
            AvailableModes []string  `json:"available_modes"`
            PickupStores []float64  `json:"pickup_stores"`
         
    }
    
    // PickupStoreDetail ...
    type PickupStoreDetail struct {

        
            Pincode float64  `json:"pincode"`
            Email string  `json:"email"`
            Name string  `json:"name"`
            Area string  `json:"area"`
            AreaCodeSlug string  `json:"area_code_slug"`
            AreaCode string  `json:"area_code"`
            City string  `json:"city"`
            StoreCode string  `json:"store_code"`
            Phone string  `json:"phone"`
            Landmark string  `json:"landmark"`
            State string  `json:"state"`
            UID float64  `json:"uid"`
            ID float64  `json:"id"`
            AddressType string  `json:"address_type"`
            Address string  `json:"address"`
            Country string  `json:"country"`
         
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
    
