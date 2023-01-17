package application



    
    // ApplicationItemSEO ...
    type ApplicationItemSEO struct {

        
            Description interface{}  `json:"description"`
            Title interface{}  `json:"title"`
         
    }
    
    // ProductListingActionPage ...
    type ProductListingActionPage struct {

        
            Query map[string]interface{}  `json:"query"`
            Type string  `json:"type"`
            Params map[string]interface{}  `json:"params"`
         
    }
    
    // ProductListingAction ...
    type ProductListingAction struct {

        
            Type string  `json:"type"`
            Page ProductListingActionPage  `json:"page"`
         
    }
    
    // CustomMetaFields ...
    type CustomMetaFields struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
         
    }
    
    // Meta ...
    type Meta struct {

        
            Source string  `json:"source"`
         
    }
    
    // Media ...
    type Media struct {

        
            Alt string  `json:"alt"`
            Type string  `json:"type"`
            URL string  `json:"url"`
            Meta Meta  `json:"meta"`
         
    }
    
    // ProductBrand ...
    type ProductBrand struct {

        
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            Logo Media  `json:"logo"`
            Description string  `json:"description"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductDetailAttribute ...
    type ProductDetailAttribute struct {

        
            Key string  `json:"key"`
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // ProductDetailGroupedAttribute ...
    type ProductDetailGroupedAttribute struct {

        
            Details []ProductDetailAttribute  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // Price ...
    type Price struct {

        
            Min float64  `json:"min"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
            Max float64  `json:"max"`
         
    }
    
    // ProductListingPrice ...
    type ProductListingPrice struct {

        
            Effective Price  `json:"effective"`
            Marked Price  `json:"marked"`
         
    }
    
    // ApplicationItemMOQ ...
    type ApplicationItemMOQ struct {

        
            Minimum float64  `json:"minimum"`
            Maximum float64  `json:"maximum"`
            IncrementUnit float64  `json:"increment_unit"`
         
    }
    
    // ProductDetail ...
    type ProductDetail struct {

        
            Highlights []string  `json:"highlights"`
            Similars []string  `json:"similars"`
            ProductOnlineDate string  `json:"product_online_date"`
            Seo ApplicationItemSEO  `json:"seo"`
            Action ProductListingAction  `json:"action"`
            Type string  `json:"type"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            ItemType string  `json:"item_type"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            Categories []ProductBrand  `json:"categories"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Rating float64  `json:"rating"`
            Medias []Media  `json:"medias"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Tags []string  `json:"tags"`
            HasVariant bool  `json:"has_variant"`
            ShortDescription string  `json:"short_description"`
            Price ProductListingPrice  `json:"price"`
            Name string  `json:"name"`
            Discount string  `json:"discount"`
            Attributes map[string]interface{}  `json:"attributes"`
            IsDependent bool  `json:"is_dependent"`
            Moq ApplicationItemMOQ  `json:"moq"`
            Brand ProductBrand  `json:"brand"`
            ItemCode string  `json:"item_code"`
            Color string  `json:"color"`
            Tryouts []string  `json:"tryouts"`
            ImageNature string  `json:"image_nature"`
            TeaserTag string  `json:"teaser_tag"`
            RatingCount float64  `json:"rating_count"`
         
    }
    
    // ErrorResponse ...
    type ErrorResponse struct {

        
            Error string  `json:"error"`
         
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
            Unit string  `json:"unit"`
            Length float64  `json:"length"`
            IsDefault bool  `json:"is_default"`
            Height float64  `json:"height"`
         
    }
    
    // ProductSize ...
    type ProductSize struct {

        
            Weight Weight  `json:"weight"`
            Quantity float64  `json:"quantity"`
            Dimension Dimension  `json:"dimension"`
            Value string  `json:"value"`
            IsAvailable bool  `json:"is_available"`
            Display string  `json:"display"`
            SellerIdentifiers []string  `json:"seller_identifiers"`
         
    }
    
    // SizeChartValues ...
    type SizeChartValues struct {

        
            Col1 string  `json:"col_1"`
            Col2 string  `json:"col_2"`
            Col4 string  `json:"col_4"`
            Col3 string  `json:"col_3"`
            Col5 string  `json:"col_5"`
            Col6 string  `json:"col_6"`
         
    }
    
    // ColumnHeader ...
    type ColumnHeader struct {

        
            Convertable bool  `json:"convertable"`
            Value string  `json:"value"`
         
    }
    
    // ColumnHeaders ...
    type ColumnHeaders struct {

        
            Col1 ColumnHeader  `json:"col_1"`
            Col2 ColumnHeader  `json:"col_2"`
            Col4 ColumnHeader  `json:"col_4"`
            Col3 ColumnHeader  `json:"col_3"`
            Col5 ColumnHeader  `json:"col_5"`
            Col6 ColumnHeader  `json:"col_6"`
         
    }
    
    // SizeChart ...
    type SizeChart struct {

        
            Image string  `json:"image"`
            Title string  `json:"title"`
            Unit string  `json:"unit"`
            Sizes []SizeChartValues  `json:"sizes"`
            SizeTip string  `json:"size_tip"`
            Headers ColumnHeaders  `json:"headers"`
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
            SizeChart SizeChart  `json:"size_chart"`
            Price ProductListingPrice  `json:"price"`
            Sellable bool  `json:"sellable"`
            Stores ProductSizeStores  `json:"stores"`
            Discount string  `json:"discount"`
         
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

        
            Items []ProductDetail  `json:"items"`
            AttributesMetadata []AttributeMetadata  `json:"attributes_metadata"`
         
    }
    
    // ProductCompareResponse ...
    type ProductCompareResponse struct {

        
            Items []ProductDetail  `json:"items"`
            AttributesMetadata []AttributeMetadata  `json:"attributes_metadata"`
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
         
    }
    
    // ProductFrequentlyComparedSimilarResponse ...
    type ProductFrequentlyComparedSimilarResponse struct {

        
            Similars ProductCompareResponse  `json:"similars"`
         
    }
    
    // ProductVariantItemResponse ...
    type ProductVariantItemResponse struct {

        
            Slug string  `json:"slug"`
            Value string  `json:"value"`
            IsAvailable bool  `json:"is_available"`
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            ColorName string  `json:"color_name"`
            Color string  `json:"color"`
            Medias []Media  `json:"medias"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductVariantResponse ...
    type ProductVariantResponse struct {

        
            Key string  `json:"key"`
            Items []ProductVariantItemResponse  `json:"items"`
            DisplayType string  `json:"display_type"`
            Header string  `json:"header"`
         
    }
    
    // ProductVariantsResponse ...
    type ProductVariantsResponse struct {

        
            Variants []ProductVariantResponse  `json:"variants"`
         
    }
    
    // StoreDetail ...
    type StoreDetail struct {

        
            City string  `json:"city"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            ID float64  `json:"id"`
         
    }
    
    // ProductStockPrice ...
    type ProductStockPrice struct {

        
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            Currency string  `json:"currency"`
         
    }
    
    // CompanyDetail ...
    type CompanyDetail struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // Seller ...
    type Seller struct {

        
            Name string  `json:"name"`
            Count float64  `json:"count"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductStockStatusItem ...
    type ProductStockStatusItem struct {

        
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
            Store StoreDetail  `json:"store"`
            Size string  `json:"size"`
            Price ProductStockPrice  `json:"price"`
            Company CompanyDetail  `json:"company"`
            Identifier map[string]interface{}  `json:"identifier"`
            Seller Seller  `json:"seller"`
            UID string  `json:"uid"`
         
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
    
    // ProductFiltersKey ...
    type ProductFiltersKey struct {

        
            Name string  `json:"name"`
            Display string  `json:"display"`
            Logo string  `json:"logo"`
            Kind string  `json:"kind"`
         
    }
    
    // ProductFiltersValue ...
    type ProductFiltersValue struct {

        
            Max float64  `json:"max"`
            SelectedMin float64  `json:"selected_min"`
            CurrencyCode string  `json:"currency_code"`
            QueryFormat string  `json:"query_format"`
            CurrencySymbol string  `json:"currency_symbol"`
            Value string  `json:"value"`
            Count float64  `json:"count"`
            DisplayFormat string  `json:"display_format"`
            Display string  `json:"display"`
            SelectedMax float64  `json:"selected_max"`
            Min float64  `json:"min"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // ProductFilters ...
    type ProductFilters struct {

        
            Key ProductFiltersKey  `json:"key"`
            Values []ProductFiltersValue  `json:"values"`
         
    }
    
    // ProductSortOn ...
    type ProductSortOn struct {

        
            Name string  `json:"name"`
            Value string  `json:"value"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // ProductVariantListingResponse ...
    type ProductVariantListingResponse struct {

        
            Total float64  `json:"total"`
            Items []ProductVariantItemResponse  `json:"items"`
            DisplayType string  `json:"display_type"`
            Key string  `json:"key"`
            Header string  `json:"header"`
         
    }
    
    // NetQuantity ...
    type NetQuantity struct {

        
            Unit interface{}  `json:"unit"`
            Value float64  `json:"value"`
         
    }
    
    // ProductListingDetail ...
    type ProductListingDetail struct {

        
            Highlights []string  `json:"highlights"`
            Variants []ProductVariantListingResponse  `json:"variants"`
            Similars []string  `json:"similars"`
            Sizes []string  `json:"sizes"`
            ProductOnlineDate string  `json:"product_online_date"`
            Seo ApplicationItemSEO  `json:"seo"`
            Action ProductListingAction  `json:"action"`
            Type string  `json:"type"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            ItemType string  `json:"item_type"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            Categories []ProductBrand  `json:"categories"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Rating float64  `json:"rating"`
            Sellable bool  `json:"sellable"`
            Medias []Media  `json:"medias"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Tags []string  `json:"tags"`
            HasVariant bool  `json:"has_variant"`
            ShortDescription string  `json:"short_description"`
            Price ProductListingPrice  `json:"price"`
            Name string  `json:"name"`
            Discount string  `json:"discount"`
            Attributes map[string]interface{}  `json:"attributes"`
            IsDependent bool  `json:"is_dependent"`
            Moq ApplicationItemMOQ  `json:"moq"`
            Brand ProductBrand  `json:"brand"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            ItemCode string  `json:"item_code"`
            Identifiers []string  `json:"identifiers"`
            Color string  `json:"color"`
            Tryouts []string  `json:"tryouts"`
            ImageNature string  `json:"image_nature"`
            TeaserTag string  `json:"teaser_tag"`
            RatingCount float64  `json:"rating_count"`
         
    }
    
    // ProductListingResponse ...
    type ProductListingResponse struct {

        
            Filters []ProductFilters  `json:"filters"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Items []ProductListingDetail  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ImageUrls ...
    type ImageUrls struct {

        
            Portrait Media  `json:"portrait"`
            Landscape Media  `json:"landscape"`
         
    }
    
    // BrandItem ...
    type BrandItem struct {

        
            Slug string  `json:"slug"`
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            Departments []string  `json:"departments"`
            Logo Media  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            Description string  `json:"description"`
            Discount string  `json:"discount"`
            UID float64  `json:"uid"`
         
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
            Logo Media  `json:"logo"`
            Description string  `json:"description"`
            UID float64  `json:"uid"`
         
    }
    
    // DepartmentIdentifier ...
    type DepartmentIdentifier struct {

        
            Slug string  `json:"slug"`
            UID float64  `json:"uid"`
         
    }
    
    // ThirdLevelChild ...
    type ThirdLevelChild struct {

        
            Slug string  `json:"slug"`
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            Childs []map[string]interface{}  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
         
    }
    
    // SecondLevelChild ...
    type SecondLevelChild struct {

        
            Slug string  `json:"slug"`
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            Childs []ThirdLevelChild  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
         
    }
    
    // Child ...
    type Child struct {

        
            Slug string  `json:"slug"`
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            Childs []SecondLevelChild  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
         
    }
    
    // CategoryItems ...
    type CategoryItems struct {

        
            Slug string  `json:"slug"`
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            Childs []Child  `json:"childs"`
            Banners ImageUrls  `json:"banners"`
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
            Logo Media  `json:"logo"`
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
            Name string  `json:"name"`
            Logo Media  `json:"logo"`
            PriorityOrder float64  `json:"priority_order"`
            UID float64  `json:"uid"`
         
    }
    
    // DepartmentResponse ...
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // AutocompleteItem ...
    type AutocompleteItem struct {

        
            Action ProductListingAction  `json:"action"`
            Type string  `json:"type"`
            Logo Media  `json:"logo"`
            Display string  `json:"display"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
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

        
            Op string  `json:"op"`
            Value []interface{}  `json:"value"`
            Attribute string  `json:"attribute"`
         
    }
    
    // GetCollectionDetailNest ...
    type GetCollectionDetailNest struct {

        
            AllowFacets bool  `json:"allow_facets"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            Action ProductListingAction  `json:"action"`
            Logo Media  `json:"logo"`
            Type string  `json:"type"`
            Meta map[string]interface{}  `json:"meta"`
            Description string  `json:"description"`
            SortOn string  `json:"sort_on"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Badge map[string]interface{}  `json:"badge"`
            Slug string  `json:"slug"`
            Banners ImageUrls  `json:"banners"`
            AllowSort bool  `json:"allow_sort"`
            Tag []string  `json:"tag"`
            UID string  `json:"uid"`
            Cron map[string]interface{}  `json:"cron"`
            Name string  `json:"name"`
            Query []CollectionQuery  `json:"query"`
            Priority float64  `json:"priority"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
         
    }
    
    // GetCollectionListingResponse ...
    type GetCollectionListingResponse struct {

        
            Filters CollectionListingFilter  `json:"filters"`
            Items []GetCollectionDetailNest  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CollectionDetailResponse ...
    type CollectionDetailResponse struct {

        
            AllowFacets bool  `json:"allow_facets"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            Logo Media  `json:"logo"`
            Type string  `json:"type"`
            Meta map[string]interface{}  `json:"meta"`
            Description string  `json:"description"`
            SortOn string  `json:"sort_on"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Badge map[string]interface{}  `json:"badge"`
            Slug string  `json:"slug"`
            Banners ImageUrls  `json:"banners"`
            AllowSort bool  `json:"allow_sort"`
            Tag []string  `json:"tag"`
            Cron map[string]interface{}  `json:"cron"`
            Name string  `json:"name"`
            Query []CollectionQuery  `json:"query"`
            Priority float64  `json:"priority"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
         
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
            Products []float64  `json:"products"`
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

        
            StoreCode string  `json:"store_code"`
            Pincode float64  `json:"pincode"`
            Country string  `json:"country"`
            StoreEmail string  `json:"store_email"`
            State string  `json:"state"`
            Address string  `json:"address"`
            City string  `json:"city"`
            Name string  `json:"name"`
            LatLong LatLong  `json:"lat_long"`
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
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            PriorityOrder float64  `json:"priority_order"`
            UID float64  `json:"uid"`
         
    }
    
    // StoreAddressSerializer ...
    type StoreAddressSerializer struct {

        
            Pincode float64  `json:"pincode"`
            Longitude float64  `json:"longitude"`
            Country string  `json:"country"`
            Address1 string  `json:"address1"`
            Latitude float64  `json:"latitude"`
            City string  `json:"city"`
            State string  `json:"state"`
            Address2 string  `json:"address2"`
            Landmark string  `json:"landmark"`
         
    }
    
    // SellerPhoneNumber ...
    type SellerPhoneNumber struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // StoreManagerSerializer ...
    type StoreManagerSerializer struct {

        
            Email string  `json:"email"`
            Name string  `json:"name"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
         
    }
    
    // CompanyStore ...
    type CompanyStore struct {

        
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
            BusinessType string  `json:"business_type"`
            UID float64  `json:"uid"`
         
    }
    
    // AppStore ...
    type AppStore struct {

        
            Address StoreAddressSerializer  `json:"address"`
            Name string  `json:"name"`
            Manager StoreManagerSerializer  `json:"manager"`
            Company CompanyStore  `json:"company"`
            Departments []StoreDepartments  `json:"departments"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            UID float64  `json:"uid"`
         
    }
    
    // ApplicationStoreListing ...
    type ApplicationStoreListing struct {

        
            Filters []StoreDepartments  `json:"filters"`
            Items []AppStore  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // Time ...
    type Time struct {

        
            Minute float64  `json:"minute"`
            Hour float64  `json:"hour"`
         
    }
    
    // StoreTiming ...
    type StoreTiming struct {

        
            Closing Time  `json:"closing"`
            Open bool  `json:"open"`
            Weekday string  `json:"weekday"`
            Opening Time  `json:"opening"`
         
    }
    
    // StoreDetails ...
    type StoreDetails struct {

        
            Timing []StoreTiming  `json:"timing"`
            Address StoreAddressSerializer  `json:"address"`
            Name string  `json:"name"`
            Manager StoreManagerSerializer  `json:"manager"`
            Company CompanyStore  `json:"company"`
            Departments []StoreDepartments  `json:"departments"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
         
    }
    
    // UserDetail ...
    type UserDetail struct {

        
            Contact string  `json:"contact"`
            SuperUser bool  `json:"super_user"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // Size ...
    type Size struct {

        
            Display interface{}  `json:"display"`
            Value interface{}  `json:"value"`
            IsAvailable bool  `json:"is_available"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // ProductGroupPrice ...
    type ProductGroupPrice struct {

        
            MaxMarked float64  `json:"max_marked"`
            MinMarked float64  `json:"min_marked"`
            MaxEffective float64  `json:"max_effective"`
            MinEffective float64  `json:"min_effective"`
            Currency interface{}  `json:"currency"`
         
    }
    
    // ProductDetails ...
    type ProductDetails struct {

        
            Highlights []interface{}  `json:"highlights"`
            Media []map[string]interface{}  `json:"media"`
            Description interface{}  `json:"description"`
            CountryOfOrigin interface{}  `json:"country_of_origin"`
            Slug interface{}  `json:"slug"`
            GroupedAttributes map[string]interface{}  `json:"grouped_attributes"`
            Images []interface{}  `json:"images"`
            Rating float64  `json:"rating"`
            Identifier map[string]interface{}  `json:"identifier"`
            TemplateTag interface{}  `json:"template_tag"`
            IsSet bool  `json:"is_set"`
            HasVariant bool  `json:"has_variant"`
            ShortDescription interface{}  `json:"short_description"`
            Name interface{}  `json:"name"`
            Attributes map[string]interface{}  `json:"attributes"`
            BrandUID float64  `json:"brand_uid"`
            ItemCode interface{}  `json:"item_code"`
            HsnCode float64  `json:"hsn_code"`
            OutOfStock bool  `json:"out_of_stock"`
            ImageNature interface{}  `json:"image_nature"`
            RatingCount float64  `json:"rating_count"`
         
    }
    
    // ProductInGroup ...
    type ProductInGroup struct {

        
            ProductUID float64  `json:"product_uid"`
            AutoSelect bool  `json:"auto_select"`
            MinQuantity float64  `json:"min_quantity"`
            Sizes []Size  `json:"sizes"`
            Price ProductGroupPrice  `json:"price"`
            MaxQuantity float64  `json:"max_quantity"`
            ProductDetails ProductDetails  `json:"product_details"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            AllowRemove bool  `json:"allow_remove"`
         
    }
    
    // ProductGroupingModel ...
    type ProductGroupingModel struct {

        
            ModifiedBy UserDetail  `json:"modified_by"`
            Choice interface{}  `json:"choice"`
            CompanyID float64  `json:"company_id"`
            ModifiedOn string  `json:"modified_on"`
            Slug interface{}  `json:"slug"`
            IsActive bool  `json:"is_active"`
            CreatedOn string  `json:"created_on"`
            Name interface{}  `json:"name"`
            Logo string  `json:"logo"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            Products []ProductInGroup  `json:"products"`
            Meta map[string]interface{}  `json:"meta"`
            CreatedBy UserDetail  `json:"created_by"`
            ID interface{}  `json:"_id"`
            VerifiedOn string  `json:"verified_on"`
            VerifiedBy UserDetail  `json:"verified_by"`
            PageVisibility []interface{}  `json:"page_visibility"`
         
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

        
            SizeDistribution ProductSetDistributionV2  `json:"size_distribution"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // DetailsSchemaV2 ...
    type DetailsSchemaV2 struct {

        
            Key string  `json:"key"`
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // SellerGroupAttributes ...
    type SellerGroupAttributes struct {

        
            Details []DetailsSchemaV2  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // StrategyWiseListingSchemaV2 ...
    type StrategyWiseListingSchemaV2 struct {

        
            Tat float64  `json:"tat"`
            Pincode float64  `json:"pincode"`
            Quantity float64  `json:"quantity"`
            Distance float64  `json:"distance"`
         
    }
    
    // ArticleAssignmentV2 ...
    type ArticleAssignmentV2 struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // ProductStockPriceV2 ...
    type ProductStockPriceV2 struct {

        
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            Currency string  `json:"currency"`
         
    }
    
    // SellerV2 ...
    type SellerV2 struct {

        
            Name string  `json:"name"`
            Count float64  `json:"count"`
            UID float64  `json:"uid"`
         
    }
    
    // MarketPlaceSttributesSchemaV2 ...
    type MarketPlaceSttributesSchemaV2 struct {

        
            Details []DetailsSchemaV2  `json:"details"`
            Title string  `json:"title"`
         
    }
    
    // ReturnConfigSchemaV2 ...
    type ReturnConfigSchemaV2 struct {

        
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // StoreV2 ...
    type StoreV2 struct {

        
            Name string  `json:"name"`
            Count float64  `json:"count"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductStockUnitPriceV2 ...
    type ProductStockUnitPriceV2 struct {

        
            Price float64  `json:"price"`
            Unit string  `json:"unit"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ProductSizePriceResponseV2 ...
    type ProductSizePriceResponseV2 struct {

        
            IsGift bool  `json:"is_gift"`
            IsCod bool  `json:"is_cod"`
            ItemType string  `json:"item_type"`
            Set ProductSetV2  `json:"set"`
            Pincode float64  `json:"pincode"`
            Quantity float64  `json:"quantity"`
            LongLat []float64  `json:"long_lat"`
            GroupedAttributes []SellerGroupAttributes  `json:"grouped_attributes"`
            StrategyWiseListing []StrategyWiseListingSchemaV2  `json:"strategy_wise_listing"`
            SellerCount float64  `json:"seller_count"`
            SpecialBadge string  `json:"special_badge"`
            ArticleAssignment ArticleAssignmentV2  `json:"article_assignment"`
            Price ProductStockPriceV2  `json:"price"`
            Seller SellerV2  `json:"seller"`
            Discount string  `json:"discount"`
            MarketplaceAttributes []MarketPlaceSttributesSchemaV2  `json:"marketplace_attributes"`
            ReturnConfig ReturnConfigSchemaV2  `json:"return_config"`
            Store StoreV2  `json:"store"`
            ArticleID string  `json:"article_id"`
            PricePerUnit ProductStockUnitPriceV2  `json:"price_per_unit"`
            PricePerPiece ProductStockPriceV2  `json:"price_per_piece"`
         
    }
    
    // ProductSizeSellerFilterSchemaV2 ...
    type ProductSizeSellerFilterSchemaV2 struct {

        
            Name string  `json:"name"`
            Value string  `json:"value"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // ProductSizeSellersResponseV2 ...
    type ProductSizeSellersResponseV2 struct {

        
            SortOn []ProductSizeSellerFilterSchemaV2  `json:"sort_on"`
            Items []ProductSizePriceResponseV2  `json:"items"`
            Page Page  `json:"page"`
         
    }
    

    
    // LoyaltyPoints ...
    type LoyaltyPoints struct {

        
            Applicable float64  `json:"applicable"`
            Total float64  `json:"total"`
            IsApplied bool  `json:"is_applied"`
            Description string  `json:"description"`
         
    }
    
    // CouponBreakup ...
    type CouponBreakup struct {

        
            Value float64  `json:"value"`
            UID string  `json:"uid"`
            Code string  `json:"code"`
            IsApplied bool  `json:"is_applied"`
            Message string  `json:"message"`
            Type string  `json:"type"`
         
    }
    
    // RawBreakup ...
    type RawBreakup struct {

        
            Discount float64  `json:"discount"`
            Coupon float64  `json:"coupon"`
            Total float64  `json:"total"`
            ConvenienceFee float64  `json:"convenience_fee"`
            Subtotal float64  `json:"subtotal"`
            DeliveryCharge float64  `json:"delivery_charge"`
            YouSaved float64  `json:"you_saved"`
            Vog float64  `json:"vog"`
            FyndCash float64  `json:"fynd_cash"`
            MrpTotal float64  `json:"mrp_total"`
            CodCharge float64  `json:"cod_charge"`
            GstCharges float64  `json:"gst_charges"`
         
    }
    
    // DisplayBreakup ...
    type DisplayBreakup struct {

        
            Value float64  `json:"value"`
            Display string  `json:"display"`
            Key string  `json:"key"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Message []string  `json:"message"`
         
    }
    
    // CartBreakup ...
    type CartBreakup struct {

        
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Coupon CouponBreakup  `json:"coupon"`
            Raw RawBreakup  `json:"raw"`
            Display []DisplayBreakup  `json:"display"`
         
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
            DefaultOptions string  `json:"default_options"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
    }
    
    // BaseInfo ...
    type BaseInfo struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
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

        
            Base BasePrice  `json:"base"`
            Converted BasePrice  `json:"converted"`
         
    }
    
    // ProductArticle ...
    type ProductArticle struct {

        
            Seller BaseInfo  `json:"seller"`
            Quantity float64  `json:"quantity"`
            Size string  `json:"size"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Price ArticlePriceInfo  `json:"price"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Store BaseInfo  `json:"store"`
            UID string  `json:"uid"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Type string  `json:"type"`
         
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

        
            Slug string  `json:"slug"`
            Brand BaseInfo  `json:"brand"`
            Name string  `json:"name"`
            Categories []CategoryInfo  `json:"categories"`
            Images []ProductImage  `json:"images"`
            Action ProductAction  `json:"action"`
            UID float64  `json:"uid"`
            Type string  `json:"type"`
         
    }
    
    // CartProductIdentifer ...
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductPrice ...
    type ProductPrice struct {

        
            Effective float64  `json:"effective"`
            AddOn float64  `json:"add_on"`
            Marked float64  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Selling float64  `json:"selling"`
         
    }
    
    // ProductPriceInfo ...
    type ProductPriceInfo struct {

        
            Base ProductPrice  `json:"base"`
            Converted ProductPrice  `json:"converted"`
         
    }
    
    // AppliedPromotion ...
    type AppliedPromotion struct {

        
            ArticleQuantity float64  `json:"article_quantity"`
            PromoID string  `json:"promo_id"`
            PromotionType string  `json:"promotion_type"`
            MrpPromotion bool  `json:"mrp_promotion"`
            OfferText string  `json:"offer_text"`
            Amount float64  `json:"amount"`
         
    }
    
    // PromoMeta ...
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // ProductAvailability ...
    type ProductAvailability struct {

        
            IsValid bool  `json:"is_valid"`
            OutOfStock bool  `json:"out_of_stock"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            Sizes []string  `json:"sizes"`
            Deliverable bool  `json:"deliverable"`
         
    }
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            Article ProductArticle  `json:"article"`
            Discount string  `json:"discount"`
            Product CartProduct  `json:"product"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            CouponMessage string  `json:"coupon_message"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            Message string  `json:"message"`
            Key string  `json:"key"`
            Price ProductPriceInfo  `json:"price"`
            IsSet bool  `json:"is_set"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Availability ProductAvailability  `json:"availability"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // CartCurrency ...
    type CartCurrency struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // CartDetailResponse ...
    type CartDetailResponse struct {

        
            BreakupValues CartBreakup  `json:"breakup_values"`
            IsValid bool  `json:"is_valid"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            LastModified string  `json:"last_modified"`
            Comment string  `json:"comment"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Items []CartProductInfo  `json:"items"`
            Gstin string  `json:"gstin"`
            BuyNow bool  `json:"buy_now"`
            Currency CartCurrency  `json:"currency"`
            CouponText string  `json:"coupon_text"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            ID string  `json:"id"`
            Message string  `json:"message"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            Quantity float64  `json:"quantity"`
            Pos bool  `json:"pos"`
            Display string  `json:"display"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ArticleID string  `json:"article_id"`
            ItemID float64  `json:"item_id"`
            ProductGroupTags []string  `json:"product_group_tags"`
            SellerID float64  `json:"seller_id"`
            ItemSize string  `json:"item_size"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            StoreID float64  `json:"store_id"`
         
    }
    
    // AddCartRequest ...
    type AddCartRequest struct {

        
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

        
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ArticleID string  `json:"article_id"`
            ItemID float64  `json:"item_id"`
            ItemIndex float64  `json:"item_index"`
            ItemSize string  `json:"item_size"`
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
    
    // PageCoupon ...
    type PageCoupon struct {

        
            HasPrevious bool  `json:"has_previous"`
            Total float64  `json:"total"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
            TotalItemCount float64  `json:"total_item_count"`
         
    }
    
    // Coupon ...
    type Coupon struct {

        
            IsApplicable bool  `json:"is_applicable"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            ExpiresOn string  `json:"expires_on"`
            Message string  `json:"message"`
            SubTitle string  `json:"sub_title"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            CouponCode string  `json:"coupon_code"`
            Title string  `json:"title"`
            IsApplied bool  `json:"is_applied"`
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
    
    // OfferSeller ...
    type OfferSeller struct {

        
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // OfferPrice ...
    type OfferPrice struct {

        
            BulkEffective float64  `json:"bulk_effective"`
            Marked float64  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
         
    }
    
    // OfferItem ...
    type OfferItem struct {

        
            Best bool  `json:"best"`
            Total float64  `json:"total"`
            Margin float64  `json:"margin"`
            AutoApplied bool  `json:"auto_applied"`
            Price OfferPrice  `json:"price"`
            Quantity float64  `json:"quantity"`
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

        
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // Address ...
    type Address struct {

        
            State string  `json:"state"`
            Area string  `json:"area"`
            AddressType string  `json:"address_type"`
            Tags []string  `json:"tags"`
            Phone string  `json:"phone"`
            Meta map[string]interface{}  `json:"meta"`
            CheckoutMode string  `json:"checkout_mode"`
            IsDefaultAddress bool  `json:"is_default_address"`
            ID string  `json:"id"`
            Address string  `json:"address"`
            CountryCode string  `json:"country_code"`
            Country string  `json:"country"`
            AreaCode string  `json:"area_code"`
            Email string  `json:"email"`
            City string  `json:"city"`
            GeoLocation GeoLocation  `json:"geo_location"`
            AreaCodeSlug string  `json:"area_code_slug"`
            IsActive bool  `json:"is_active"`
            Name string  `json:"name"`
            Landmark string  `json:"landmark"`
            UserID string  `json:"user_id"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
         
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

        
            IsUpdated bool  `json:"is_updated"`
            IsDefaultAddress bool  `json:"is_default_address"`
            ID string  `json:"id"`
            Success bool  `json:"success"`
         
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

        
            PaymentMode string  `json:"payment_mode"`
            AggregatorName string  `json:"aggregator_name"`
            PaymentIdentifier string  `json:"payment_identifier"`
            AddressID string  `json:"address_id"`
            MerchantCode string  `json:"merchant_code"`
            ID string  `json:"id"`
         
    }
    
    // CouponValidity ...
    type CouponValidity struct {

        
            Discount float64  `json:"discount"`
            Valid bool  `json:"valid"`
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

        
            FulfillmentID float64  `json:"fulfillment_id"`
            Shipments float64  `json:"shipments"`
            OrderType string  `json:"order_type"`
            Items []CartProductInfo  `json:"items"`
            BoxType string  `json:"box_type"`
            ShipmentType string  `json:"shipment_type"`
            DpID string  `json:"dp_id"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            FulfillmentType string  `json:"fulfillment_type"`
            Promise ShipmentPromise  `json:"promise"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Shipments []ShipmentResponse  `json:"shipments"`
            Comment string  `json:"comment"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Gstin string  `json:"gstin"`
            Error bool  `json:"error"`
            BuyNow bool  `json:"buy_now"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Message string  `json:"message"`
            ID string  `json:"id"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            IsValid bool  `json:"is_valid"`
            CheckoutMode string  `json:"checkout_mode"`
            LastModified string  `json:"last_modified"`
            CartID float64  `json:"cart_id"`
            Currency CartCurrency  `json:"currency"`
            CouponText string  `json:"coupon_text"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            UID string  `json:"uid"`
         
    }
    
    // StaffCheckout ...
    type StaffCheckout struct {

        
            ID string  `json:"_id"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            User string  `json:"user"`
         
    }
    
    // CartCheckoutDetailRequest ...
    type CartCheckoutDetailRequest struct {

        
            BillingAddress map[string]interface{}  `json:"billing_address"`
            PaymentMode string  `json:"payment_mode"`
            CallbackURL string  `json:"callback_url"`
            PaymentIdentifier string  `json:"payment_identifier"`
            Staff StaffCheckout  `json:"staff"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            AddressID string  `json:"address_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            BillingAddressID string  `json:"billing_address_id"`
            Aggregator string  `json:"aggregator"`
            MerchantCode string  `json:"merchant_code"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            OrderingStore float64  `json:"ordering_store"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // CheckCart ...
    type CheckCart struct {

        
            UserType string  `json:"user_type"`
            CodMessage string  `json:"cod_message"`
            Success bool  `json:"success"`
            Gstin string  `json:"gstin"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            StoreCode string  `json:"store_code"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CheckoutMode string  `json:"checkout_mode"`
            OrderID string  `json:"order_id"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            CouponText string  `json:"coupon_text"`
            DeliveryCharges float64  `json:"delivery_charges"`
            UID string  `json:"uid"`
            ID string  `json:"id"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Comment string  `json:"comment"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CodCharges float64  `json:"cod_charges"`
            BuyNow bool  `json:"buy_now"`
            Message string  `json:"message"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            IsValid bool  `json:"is_valid"`
            LastModified string  `json:"last_modified"`
            ErrorMessage string  `json:"error_message"`
            Items []CartProductInfo  `json:"items"`
            CartID float64  `json:"cart_id"`
            Currency CartCurrency  `json:"currency"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            CodAvailable bool  `json:"cod_available"`
         
    }
    
    // CartCheckoutResponse ...
    type CartCheckoutResponse struct {

        
            CallbackURL string  `json:"callback_url"`
            Success bool  `json:"success"`
            OrderID string  `json:"order_id"`
            Data map[string]interface{}  `json:"data"`
            Cart CheckCart  `json:"cart"`
            AppInterceptURL string  `json:"app_intercept_url"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Message string  `json:"message"`
         
    }
    
    // CartMetaRequest ...
    type CartMetaRequest struct {

        
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            Comment string  `json:"comment"`
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

        
            Token string  `json:"token"`
            ShareURL string  `json:"share_url"`
         
    }
    
    // SharedCartDetails ...
    type SharedCartDetails struct {

        
            CreatedOn string  `json:"created_on"`
            Source map[string]interface{}  `json:"source"`
            User map[string]interface{}  `json:"user"`
            Token string  `json:"token"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            Gstin string  `json:"gstin"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CheckoutMode string  `json:"checkout_mode"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            CouponText string  `json:"coupon_text"`
            UID string  `json:"uid"`
            ID string  `json:"id"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Comment string  `json:"comment"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            BuyNow bool  `json:"buy_now"`
            Message string  `json:"message"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            IsValid bool  `json:"is_valid"`
            LastModified string  `json:"last_modified"`
            Items []CartProductInfo  `json:"items"`
            CartID float64  `json:"cart_id"`
            Currency CartCurrency  `json:"currency"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Error string  `json:"error"`
            Cart SharedCart  `json:"cart"`
         
    }
    
    // PromotionOffer ...
    type PromotionOffer struct {

        
            Description string  `json:"description"`
            OfferText string  `json:"offer_text"`
            ValidTill string  `json:"valid_till"`
            ID string  `json:"id"`
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
    
    // CurrencyInfo ...
    type CurrencyInfo struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // LadderPrice ...
    type LadderPrice struct {

        
            Marked float64  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            OfferPrice float64  `json:"offer_price"`
            Effective float64  `json:"effective"`
         
    }
    
    // LadderOfferItem ...
    type LadderOfferItem struct {

        
            Margin float64  `json:"margin"`
            MaxQuantity float64  `json:"max_quantity"`
            MinQuantity float64  `json:"min_quantity"`
            Price LadderPrice  `json:"price"`
            Type string  `json:"type"`
         
    }
    
    // LadderPriceOffer ...
    type LadderPriceOffer struct {

        
            Description string  `json:"description"`
            OfferText string  `json:"offer_text"`
            ValidTill string  `json:"valid_till"`
            OfferPrices []LadderOfferItem  `json:"offer_prices"`
            ID string  `json:"id"`
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

        
            API string  `json:"api"`
            VerifyAPI string  `json:"verify_api"`
            ConfigType string  `json:"config_type"`
            MerchantID string  `json:"merchant_id"`
            MerchantKey string  `json:"merchant_key"`
            UserID string  `json:"user_id"`
            Secret string  `json:"secret"`
            Key string  `json:"key"`
            Sdk bool  `json:"sdk"`
            Pin string  `json:"pin"`
         
    }
    
    // AggregatorsConfigDetailResponse ...
    type AggregatorsConfigDetailResponse struct {

        
            Stripe AggregatorConfigDetail  `json:"stripe"`
            Env string  `json:"env"`
            Mswipe AggregatorConfigDetail  `json:"mswipe"`
            Rupifi AggregatorConfigDetail  `json:"rupifi"`
            Simpl AggregatorConfigDetail  `json:"simpl"`
            Ccavenue AggregatorConfigDetail  `json:"ccavenue"`
            Payumoney AggregatorConfigDetail  `json:"payumoney"`
            Success bool  `json:"success"`
            Juspay AggregatorConfigDetail  `json:"juspay"`
            Razorpay AggregatorConfigDetail  `json:"razorpay"`
         
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
            Refresh bool  `json:"refresh"`
            Nickname string  `json:"nickname"`
            CardID string  `json:"card_id"`
         
    }
    
    // AttachCardsResponse ...
    type AttachCardsResponse struct {

        
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // CardPaymentGateway ...
    type CardPaymentGateway struct {

        
            CustomerID string  `json:"customer_id"`
            API string  `json:"api"`
            Aggregator string  `json:"aggregator"`
         
    }
    
    // ActiveCardPaymentGatewayResponse ...
    type ActiveCardPaymentGatewayResponse struct {

        
            Message string  `json:"message"`
            Cards CardPaymentGateway  `json:"cards"`
            Success bool  `json:"success"`
         
    }
    
    // Card ...
    type Card struct {

        
            CardIssuer string  `json:"card_issuer"`
            ExpMonth float64  `json:"exp_month"`
            CardBrand string  `json:"card_brand"`
            Nickname string  `json:"nickname"`
            CardIsin string  `json:"card_isin"`
            Expired bool  `json:"expired"`
            CardNumber string  `json:"card_number"`
            CardToken string  `json:"card_token"`
            CardName string  `json:"card_name"`
            CardFingerprint string  `json:"card_fingerprint"`
            ExpYear float64  `json:"exp_year"`
            CardType string  `json:"card_type"`
            AggregatorName string  `json:"aggregator_name"`
            CardID string  `json:"card_id"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardBrandImage string  `json:"card_brand_image"`
            CardReference string  `json:"card_reference"`
         
    }
    
    // ListCardsResponse ...
    type ListCardsResponse struct {

        
            Message string  `json:"message"`
            Data []Card  `json:"data"`
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

        
            Payload string  `json:"payload"`
            TransactionAmountInPaise float64  `json:"transaction_amount_in_paise"`
            Aggregator string  `json:"aggregator"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            MerchantParams map[string]interface{}  `json:"merchant_params"`
            OrderItems []map[string]interface{}  `json:"order_items"`
            PhoneNumber string  `json:"phone_number"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
         
    }
    
    // ValidateCustomerResponse ...
    type ValidateCustomerResponse struct {

        
            Message string  `json:"message"`
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // ChargeCustomerRequest ...
    type ChargeCustomerRequest struct {

        
            OrderID string  `json:"order_id"`
            Amount float64  `json:"amount"`
            Aggregator string  `json:"aggregator"`
            TransactionToken string  `json:"transaction_token"`
            Verified bool  `json:"verified"`
         
    }
    
    // ChargeCustomerResponse ...
    type ChargeCustomerResponse struct {

        
            OrderID string  `json:"order_id"`
            CartID string  `json:"cart_id"`
            Aggregator string  `json:"aggregator"`
            Message string  `json:"message"`
            Status string  `json:"status"`
            Success bool  `json:"success"`
            DeliveryAddressID string  `json:"delivery_address_id"`
         
    }
    
    // PaymentInitializationRequest ...
    type PaymentInitializationRequest struct {

        
            OrderID string  `json:"order_id"`
            Contact string  `json:"contact"`
            CustomerID string  `json:"customer_id"`
            Amount float64  `json:"amount"`
            Method string  `json:"method"`
            Aggregator string  `json:"aggregator"`
            Currency string  `json:"currency"`
            Email string  `json:"email"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Timeout float64  `json:"timeout"`
            Vpa string  `json:"vpa"`
         
    }
    
    // PaymentInitializationResponse ...
    type PaymentInitializationResponse struct {

        
            CustomerID string  `json:"customer_id"`
            Vpa string  `json:"vpa"`
            Amount float64  `json:"amount"`
            Method string  `json:"method"`
            Aggregator string  `json:"aggregator"`
            Currency string  `json:"currency"`
            UpiPollURL string  `json:"upi_poll_url"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            Status string  `json:"status"`
            BqrImage string  `json:"bqr_image"`
            PollingURL string  `json:"polling_url"`
            MerchantOrderID string  `json:"merchant_order_id"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Timeout float64  `json:"timeout"`
            Success bool  `json:"success"`
            VirtualID string  `json:"virtual_id"`
         
    }
    
    // PaymentStatusUpdateRequest ...
    type PaymentStatusUpdateRequest struct {

        
            OrderID string  `json:"order_id"`
            Contact string  `json:"contact"`
            CustomerID string  `json:"customer_id"`
            Amount float64  `json:"amount"`
            Method string  `json:"method"`
            Aggregator string  `json:"aggregator"`
            Currency string  `json:"currency"`
            Status string  `json:"status"`
            Email string  `json:"email"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Vpa string  `json:"vpa"`
         
    }
    
    // PaymentStatusUpdateResponse ...
    type PaymentStatusUpdateResponse struct {

        
            Retry bool  `json:"retry"`
            RedirectURL string  `json:"redirect_url"`
            Status string  `json:"status"`
            AggregatorName string  `json:"aggregator_name"`
            Success bool  `json:"success"`
         
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

        
            LogoURL PaymentModeLogo  `json:"logo_url"`
            CardNumber string  `json:"card_number"`
            CardToken string  `json:"card_token"`
            DisplayPriority float64  `json:"display_priority"`
            Code string  `json:"code"`
            IntentApp []IntentApp  `json:"intent_app"`
            Expired bool  `json:"expired"`
            CardName string  `json:"card_name"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            CardType string  `json:"card_type"`
            CardIssuer string  `json:"card_issuer"`
            CardReference string  `json:"card_reference"`
            CardIsin string  `json:"card_isin"`
            RetryCount float64  `json:"retry_count"`
            FyndVpa string  `json:"fynd_vpa"`
            ExpYear float64  `json:"exp_year"`
            MerchantCode string  `json:"merchant_code"`
            AggregatorName string  `json:"aggregator_name"`
            IntentFlow bool  `json:"intent_flow"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            Name string  `json:"name"`
            CodLimit float64  `json:"cod_limit"`
            CardBrandImage string  `json:"card_brand_image"`
            Nickname string  `json:"nickname"`
            ExpMonth float64  `json:"exp_month"`
            CardBrand string  `json:"card_brand"`
            CardFingerprint string  `json:"card_fingerprint"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            DisplayName string  `json:"display_name"`
            CardID string  `json:"card_id"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            Timeout float64  `json:"timeout"`
            RemainingLimit float64  `json:"remaining_limit"`
         
    }
    
    // RootPaymentMode ...
    type RootPaymentMode struct {

        
            List []PaymentModeList  `json:"list"`
            SaveCard bool  `json:"save_card"`
            DisplayName string  `json:"display_name"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            DisplayPriority float64  `json:"display_priority"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            AggregatorName string  `json:"aggregator_name"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            Name string  `json:"name"`
         
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

        
            BqrRazorpay AggregatorRoute  `json:"bqr_razorpay"`
            Stripe AggregatorRoute  `json:"stripe"`
            Fynd AggregatorRoute  `json:"fynd"`
            Epaylater AggregatorRoute  `json:"epaylater"`
            UpiRazorpay AggregatorRoute  `json:"upi_razorpay"`
            Jiopay AggregatorRoute  `json:"jiopay"`
            Mswipe AggregatorRoute  `json:"mswipe"`
            Simpl AggregatorRoute  `json:"simpl"`
            Ccavenue AggregatorRoute  `json:"ccavenue"`
            Rupifi AggregatorRoute  `json:"rupifi"`
            Payubiz AggregatorRoute  `json:"payubiz"`
            Juspay AggregatorRoute  `json:"juspay"`
            Razorpay AggregatorRoute  `json:"razorpay"`
         
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

        
            Status string  `json:"status"`
            Display bool  `json:"display"`
            Message string  `json:"message"`
         
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
    
    // TransferItemsDetails ...
    type TransferItemsDetails struct {

        
            LogoSmall string  `json:"logo_small"`
            ID float64  `json:"id"`
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

        
            IsActive bool  `json:"is_active"`
            AccountNo string  `json:"account_no"`
            Comment string  `json:"comment"`
            DelightsUserName string  `json:"delights_user_name"`
            ID float64  `json:"id"`
            CreatedOn string  `json:"created_on"`
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
            BeneficiaryID string  `json:"beneficiary_id"`
            BankName string  `json:"bank_name"`
            ModifiedOn string  `json:"modified_on"`
            IfscCode string  `json:"ifsc_code"`
            TransferMode string  `json:"transfer_mode"`
            Address string  `json:"address"`
            BranchName string  `json:"branch_name"`
            DisplayName string  `json:"display_name"`
            Email string  `json:"email"`
            AccountHolder string  `json:"account_holder"`
            Mobile string  `json:"mobile"`
         
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
            BranchName string  `json:"branch_name"`
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

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // WrongOtpError ...
    type WrongOtpError struct {

        
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Success string  `json:"success"`
            Description string  `json:"description"`
         
    }
    
    // BeneficiaryModeDetails ...
    type BeneficiaryModeDetails struct {

        
            Comment string  `json:"comment"`
            Mobile string  `json:"mobile"`
            BankName string  `json:"bank_name"`
            IfscCode string  `json:"ifsc_code"`
            AccountNo string  `json:"account_no"`
            Wallet string  `json:"wallet"`
            Email string  `json:"email"`
            AccountHolder string  `json:"account_holder"`
            Address string  `json:"address"`
            BranchName string  `json:"branch_name"`
            Vpa string  `json:"vpa"`
         
    }
    
    // AddBeneficiaryDetailsRequest ...
    type AddBeneficiaryDetailsRequest struct {

        
            OrderID string  `json:"order_id"`
            TransferMode string  `json:"transfer_mode"`
            Delights bool  `json:"delights"`
            Details BeneficiaryModeDetails  `json:"details"`
            Otp string  `json:"otp"`
            RequestID string  `json:"request_id"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // RefundAccountResponse ...
    type RefundAccountResponse struct {

        
            Message string  `json:"message"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // BankDetailsForOTP ...
    type BankDetailsForOTP struct {

        
            BankName string  `json:"bank_name"`
            IfscCode string  `json:"ifsc_code"`
            AccountNo string  `json:"account_no"`
            AccountHolder string  `json:"account_holder"`
            BranchName string  `json:"branch_name"`
         
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

        
            IsVerifiedFlag string  `json:"is_verified_flag"`
            Success bool  `json:"success"`
            RequestID string  `json:"request_id"`
         
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

        
            MerchantName string  `json:"merchant_name"`
            Amount float64  `json:"amount"`
            PaymentLinkURL string  `json:"payment_link_url"`
            PaymentLinkCurrentStatus string  `json:"payment_link_current_status"`
            ExternalOrderID string  `json:"external_order_id"`
            Message string  `json:"message"`
            PollingTimeout float64  `json:"polling_timeout"`
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
         
    }
    
    // ErrorDescription ...
    type ErrorDescription struct {

        
            InvalidID bool  `json:"invalid_id"`
            Amount float64  `json:"amount"`
            Expired bool  `json:"expired"`
            Cancelled bool  `json:"cancelled"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Msg string  `json:"msg"`
            PaymentTransactionID string  `json:"payment_transaction_id"`
            MerchantName string  `json:"merchant_name"`
         
    }
    
    // CreatePaymentLinkMeta ...
    type CreatePaymentLinkMeta struct {

        
            CartID string  `json:"cart_id"`
            Amount string  `json:"amount"`
            CheckoutMode string  `json:"checkout_mode"`
            AssignCardID string  `json:"assign_card_id"`
            Pincode string  `json:"pincode"`
         
    }
    
    // CreatePaymentLinkRequest ...
    type CreatePaymentLinkRequest struct {

        
            Amount float64  `json:"amount"`
            Description string  `json:"description"`
            MobileNumber string  `json:"mobile_number"`
            ExternalOrderID string  `json:"external_order_id"`
            Email string  `json:"email"`
            Meta CreatePaymentLinkMeta  `json:"meta"`
         
    }
    
    // CreatePaymentLinkResponse ...
    type CreatePaymentLinkResponse struct {

        
            PaymentLinkURL string  `json:"payment_link_url"`
            PaymentLinkID string  `json:"payment_link_id"`
            Message string  `json:"message"`
            PollingTimeout float64  `json:"polling_timeout"`
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
         
    }
    
    // CancelOrResendPaymentLinkRequest ...
    type CancelOrResendPaymentLinkRequest struct {

        
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // ResendPaymentLinkResponse ...
    type ResendPaymentLinkResponse struct {

        
            Message string  `json:"message"`
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
            PollingTimeout float64  `json:"polling_timeout"`
         
    }
    
    // CancelPaymentLinkResponse ...
    type CancelPaymentLinkResponse struct {

        
            Message string  `json:"message"`
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
         
    }
    
    // PollingPaymentLinkResponse ...
    type PollingPaymentLinkResponse struct {

        
            OrderID string  `json:"order_id"`
            Amount float64  `json:"amount"`
            RedirectURL string  `json:"redirect_url"`
            PaymentLinkID string  `json:"payment_link_id"`
            Status string  `json:"status"`
            Message string  `json:"message"`
            AggregatorName string  `json:"aggregator_name"`
            StatusCode float64  `json:"status_code"`
            Success bool  `json:"success"`
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

        
            Mode string  `json:"mode"`
            Name string  `json:"name"`
            Meta PaymentMethodsMeta  `json:"meta"`
         
    }
    
    // CreateOrderUserRequest ...
    type CreateOrderUserRequest struct {

        
            Currency string  `json:"currency"`
            PaymentLinkID string  `json:"payment_link_id"`
            FailureCallbackURL string  `json:"failure_callback_url"`
            PaymentMethods CreateOrderUserPaymentMethods  `json:"payment_methods"`
            SuccessCallbackURL string  `json:"success_callback_url"`
            Meta map[string]interface{}  `json:"meta"`
         
    }
    
    // CreateOrderUserData ...
    type CreateOrderUserData struct {

        
            OrderID string  `json:"order_id"`
            Contact string  `json:"contact"`
            CustomerID string  `json:"customer_id"`
            Amount float64  `json:"amount"`
            Method string  `json:"method"`
            Aggregator string  `json:"aggregator"`
            Currency string  `json:"currency"`
            Email string  `json:"email"`
            MerchantOrderID string  `json:"merchant_order_id"`
            CallbackURL string  `json:"callback_url"`
         
    }
    
    // CreateOrderUserResponse ...
    type CreateOrderUserResponse struct {

        
            OrderID string  `json:"order_id"`
            Data CreateOrderUserData  `json:"data"`
            Message string  `json:"message"`
            StatusCode float64  `json:"status_code"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            CallbackURL string  `json:"callback_url"`
            Success bool  `json:"success"`
         
    }
    
    // BalanceDetails ...
    type BalanceDetails struct {

        
            FormattedValue string  `json:"formatted_value"`
            Value float64  `json:"value"`
            Currency string  `json:"currency"`
         
    }
    
    // CreditSummary ...
    type CreditSummary struct {

        
            Status string  `json:"status"`
            MerchantCustomerRefID string  `json:"merchant_customer_ref_id"`
            StatusMessage string  `json:"status_message"`
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
            SignupURL string  `json:"signup_url"`
            IsRegistered bool  `json:"is_registered"`
         
    }
    
    // CheckCreditResponse ...
    type CheckCreditResponse struct {

        
            Data CreditDetail  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // DeviceDetails ...
    type DeviceDetails struct {

        
            IdentificationNumber string  `json:"identification_number"`
            IdentifierType string  `json:"identifier_type"`
            OsVersion string  `json:"os_version"`
            Os string  `json:"os"`
            DeviceType string  `json:"device_type"`
            DeviceModel string  `json:"device_model"`
            DeviceMake string  `json:"device_make"`
         
    }
    
    // KYCAddress ...
    type KYCAddress struct {

        
            Addressline2 string  `json:"addressline2"`
            State string  `json:"state"`
            Addressline1 string  `json:"addressline1"`
            OwnershipType string  `json:"ownership_type"`
            LandMark string  `json:"land_mark"`
            City string  `json:"city"`
            Pincode string  `json:"pincode"`
         
    }
    
    // BusinessDetails ...
    type BusinessDetails struct {

        
            ShopAndEstablishment map[string]interface{}  `json:"shop_and_establishment"`
            Vintage string  `json:"vintage"`
            BusinessOwnershipType string  `json:"business_ownership_type"`
            BusinessType string  `json:"business_type"`
            EntityType string  `json:"entity_type"`
            Gstin string  `json:"gstin"`
            Fda string  `json:"fda"`
            Pan string  `json:"pan"`
            Fssai string  `json:"fssai"`
            Name string  `json:"name"`
            Address KYCAddress  `json:"address"`
         
    }
    
    // UserPersonalInfoInDetails ...
    type UserPersonalInfoInDetails struct {

        
            Dob string  `json:"dob"`
            EmailVerified bool  `json:"email_verified"`
            DrivingLicense string  `json:"driving_license"`
            FathersName string  `json:"fathers_name"`
            FirstName string  `json:"first_name"`
            Passport string  `json:"passport"`
            Phone string  `json:"phone"`
            AddressAsPerID KYCAddress  `json:"address_as_per_id"`
            MothersName string  `json:"mothers_name"`
            Gender string  `json:"gender"`
            Email string  `json:"email"`
            LastName string  `json:"last_name"`
            VoterID string  `json:"voter_id"`
            Pan string  `json:"pan"`
            MobileVerified bool  `json:"mobile_verified"`
            MiddleName string  `json:"middle_name"`
         
    }
    
    // MarketplaceInfo ...
    type MarketplaceInfo struct {

        
            MembershipID string  `json:"membership_id"`
            Name string  `json:"name"`
            DateOfJoining string  `json:"date_of_joining"`
         
    }
    
    // CustomerOnboardingRequest ...
    type CustomerOnboardingRequest struct {

        
            Device DeviceDetails  `json:"device"`
            Aggregator string  `json:"aggregator"`
            Source string  `json:"source"`
            Mcc string  `json:"mcc"`
            BusinessInfo BusinessDetails  `json:"business_info"`
            PersonalInfo UserPersonalInfoInDetails  `json:"personal_info"`
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
    

    
    // BreakupValues ...
    type BreakupValues struct {

        
            Value float64  `json:"value"`
            Name string  `json:"name"`
            Display string  `json:"display"`
         
    }
    
    // UserInfo ...
    type UserInfo struct {

        
            Name string  `json:"name"`
            Gender string  `json:"gender"`
            Mobile string  `json:"mobile"`
            Email string  `json:"email"`
         
    }
    
    // ShipmentUserInfo ...
    type ShipmentUserInfo struct {

        
            LastName string  `json:"last_name"`
            FirstName string  `json:"first_name"`
            Gender string  `json:"gender"`
            Mobile string  `json:"mobile"`
         
    }
    
    // FulfillingCompany ...
    type FulfillingCompany struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // NestedTrackingDetails ...
    type NestedTrackingDetails struct {

        
            IsCurrent bool  `json:"is_current"`
            IsPassed bool  `json:"is_passed"`
            Time string  `json:"time"`
            Status string  `json:"status"`
         
    }
    
    // TrackingDetails ...
    type TrackingDetails struct {

        
            IsCurrent bool  `json:"is_current"`
            Status string  `json:"status"`
            Time string  `json:"time"`
            TrackingDetails []NestedTrackingDetails  `json:"tracking_details"`
            IsPassed bool  `json:"is_passed"`
         
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
    
    // FulfillingStore ...
    type FulfillingStore struct {

        
            CompanyName string  `json:"company_name"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            ID float64  `json:"id"`
            Code string  `json:"code"`
         
    }
    
    // ShipmentPayment ...
    type ShipmentPayment struct {

        
            Mode string  `json:"mode"`
            Status string  `json:"status"`
            PaymentMode string  `json:"payment_mode"`
            DisplayName string  `json:"display_name"`
            Mop string  `json:"mop"`
            Logo string  `json:"logo"`
         
    }
    
    // DeliveryAddress ...
    type DeliveryAddress struct {

        
            UpdatedAt string  `json:"updated_at"`
            Landmark string  `json:"landmark"`
            Phone string  `json:"phone"`
            State string  `json:"state"`
            ContactPerson string  `json:"contact_person"`
            Address1 string  `json:"address1"`
            Email string  `json:"email"`
            Name string  `json:"name"`
            City string  `json:"city"`
            CreatedAt string  `json:"created_at"`
            Version string  `json:"version"`
            Address2 string  `json:"address2"`
            Area string  `json:"area"`
            Longitude float64  `json:"longitude"`
            Address string  `json:"address"`
            AddressType string  `json:"address_type"`
            AddressCategory string  `json:"address_category"`
            Country string  `json:"country"`
            Pincode string  `json:"pincode"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // Invoice ...
    type Invoice struct {

        
            UpdatedDate string  `json:"updated_date"`
            InvoiceURL string  `json:"invoice_url"`
            LabelURL string  `json:"label_url"`
         
    }
    
    // Prices ...
    type Prices struct {

        
            Discount float64  `json:"discount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            RefundCredit float64  `json:"refund_credit"`
            ValueOfGood float64  `json:"value_of_good"`
            RefundAmount float64  `json:"refund_amount"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            CodCharges float64  `json:"cod_charges"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            CashbackApplied float64  `json:"cashback_applied"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            CouponValue float64  `json:"coupon_value"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            TransferPrice float64  `json:"transfer_price"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Cashback float64  `json:"cashback"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            AmountPaid float64  `json:"amount_paid"`
            PriceEffective float64  `json:"price_effective"`
            FyndCredits float64  `json:"fynd_credits"`
            PriceMarked float64  `json:"price_marked"`
         
    }
    
    // ShipmentTotalDetails ...
    type ShipmentTotalDetails struct {

        
            TotalPrice float64  `json:"total_price"`
            Pieces float64  `json:"pieces"`
            Sizes float64  `json:"sizes"`
         
    }
    
    // Identifiers ...
    type Identifiers struct {

        
            SkuCode string  `json:"sku_code"`
            Ean string  `json:"ean"`
         
    }
    
    // FinancialBreakup ...
    type FinancialBreakup struct {

        
            Discount float64  `json:"discount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            RefundCredit float64  `json:"refund_credit"`
            ValueOfGood float64  `json:"value_of_good"`
            RefundAmount float64  `json:"refund_amount"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            CodCharges float64  `json:"cod_charges"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            CashbackApplied float64  `json:"cashback_applied"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            Identifiers Identifiers  `json:"identifiers"`
            CouponValue float64  `json:"coupon_value"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            TransferPrice float64  `json:"transfer_price"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Cashback float64  `json:"cashback"`
            TotalUnits float64  `json:"total_units"`
            ItemName string  `json:"item_name"`
            Size string  `json:"size"`
            HsnCode string  `json:"hsn_code"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            GstTag string  `json:"gst_tag"`
            GstFee float64  `json:"gst_fee"`
            AmountPaid float64  `json:"amount_paid"`
            FyndCredits float64  `json:"fynd_credits"`
            PriceEffective float64  `json:"price_effective"`
            PriceMarked float64  `json:"price_marked"`
         
    }
    
    // CurrentStatus ...
    type CurrentStatus struct {

        
            Name string  `json:"name"`
            UpdatedAt string  `json:"updated_at"`
            JourneyType string  `json:"journey_type"`
            Status string  `json:"status"`
         
    }
    
    // ItemBrand ...
    type ItemBrand struct {

        
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
    }
    
    // Item ...
    type Item struct {

        
            SellerIdentifier string  `json:"seller_identifier"`
            Name string  `json:"name"`
            Brand ItemBrand  `json:"brand"`
            Image []string  `json:"image"`
            ID float64  `json:"id"`
            SlugKey string  `json:"slug_key"`
            Code string  `json:"code"`
            Size string  `json:"size"`
         
    }
    
    // AppliedFreeArticles ...
    type AppliedFreeArticles struct {

        
            ParentItemIdentifier string  `json:"parent_item_identifier"`
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
            FreeGiftItemDetails map[string]interface{}  `json:"free_gift_item_details"`
         
    }
    
    // AppliedPromos ...
    type AppliedPromos struct {

        
            PromotionType string  `json:"promotion_type"`
            PromoID string  `json:"promo_id"`
            PromotionName string  `json:"promotion_name"`
            MrpPromotion bool  `json:"mrp_promotion"`
            ArticleQuantity float64  `json:"article_quantity"`
            Amount float64  `json:"amount"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
         
    }
    
    // Bags ...
    type Bags struct {

        
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            DeliveryDate string  `json:"delivery_date"`
            LineNumber float64  `json:"line_number"`
            Prices Prices  `json:"prices"`
            CanCancel bool  `json:"can_cancel"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            CanReturn bool  `json:"can_return"`
            SellerIdentifier string  `json:"seller_identifier"`
            Item Item  `json:"item"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            ReturnableDate string  `json:"returnable_date"`
            ID float64  `json:"id"`
            Quantity float64  `json:"quantity"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
         
    }
    
    // ShipmentStatus ...
    type ShipmentStatus struct {

        
            Title string  `json:"title"`
            HexCode string  `json:"hex_code"`
         
    }
    
    // Shipments ...
    type Shipments struct {

        
            UserInfo ShipmentUserInfo  `json:"user_info"`
            TrakingNo string  `json:"traking_no"`
            AwbNo string  `json:"awb_no"`
            ShipmentID string  `json:"shipment_id"`
            FulfillingCompany FulfillingCompany  `json:"fulfilling_company"`
            BeneficiaryDetails bool  `json:"beneficiary_details"`
            TrackingDetails []TrackingDetails  `json:"tracking_details"`
            Promise Promise  `json:"promise"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            SizeInfo map[string]interface{}  `json:"size_info"`
            ReturnableDate string  `json:"returnable_date"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            Payment ShipmentPayment  `json:"payment"`
            OrderID string  `json:"order_id"`
            DeliveryAddress DeliveryAddress  `json:"delivery_address"`
            ShowDownloadInvoice bool  `json:"show_download_invoice"`
            RefundDetails map[string]interface{}  `json:"refund_details"`
            TotalBags float64  `json:"total_bags"`
            Invoice Invoice  `json:"invoice"`
            TrackURL string  `json:"track_url"`
            NeedHelpURL string  `json:"need_help_url"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            DeliveryDate string  `json:"delivery_date"`
            CanBreak map[string]interface{}  `json:"can_break"`
            Prices Prices  `json:"prices"`
            CanReturn bool  `json:"can_return"`
            TotalDetails ShipmentTotalDetails  `json:"total_details"`
            Bags []Bags  `json:"bags"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            CanCancel bool  `json:"can_cancel"`
            ShowTrackLink bool  `json:"show_track_link"`
            DpName string  `json:"dp_name"`
            OrderType string  `json:"order_type"`
            Comment string  `json:"comment"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
         
    }
    
    // BagsForReorderArticleAssignment ...
    type BagsForReorderArticleAssignment struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // BagsForReorder ...
    type BagsForReorder struct {

        
            StoreID float64  `json:"store_id"`
            SellerID float64  `json:"seller_id"`
            ItemID float64  `json:"item_id"`
            ArticleAssignment BagsForReorderArticleAssignment  `json:"article_assignment"`
            ItemSize string  `json:"item_size"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // OrderSchema ...
    type OrderSchema struct {

        
            OrderID string  `json:"order_id"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            OrderCreatedTime string  `json:"order_created_time"`
            UserInfo UserInfo  `json:"user_info"`
            Shipments []Shipments  `json:"shipments"`
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
            BagsForReorder []BagsForReorder  `json:"bags_for_reorder"`
         
    }
    
    // OrderPage ...
    type OrderPage struct {

        
            ItemTotal float64  `json:"item_total"`
            Type string  `json:"type"`
            Current float64  `json:"current"`
            HasNext bool  `json:"has_next"`
            Size float64  `json:"size"`
         
    }
    
    // OrderStatuses ...
    type OrderStatuses struct {

        
            Value float64  `json:"value"`
            IsSelected bool  `json:"is_selected"`
            Display string  `json:"display"`
         
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
            ShipmentID string  `json:"shipment_id"`
            Success bool  `json:"success"`
            PresignedType string  `json:"presigned_type"`
         
    }
    
    // Track ...
    type Track struct {

        
            UpdatedAt string  `json:"updated_at"`
            Awb string  `json:"awb"`
            Reason string  `json:"reason"`
            Status string  `json:"status"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            UpdatedTime string  `json:"updated_time"`
            AccountName string  `json:"account_name"`
            ShipmentType string  `json:"shipment_type"`
         
    }
    
    // ShipmentTrack ...
    type ShipmentTrack struct {

        
            Results []Track  `json:"results"`
         
    }
    
    // CustomerDetailsResponse ...
    type CustomerDetailsResponse struct {

        
            OrderID string  `json:"order_id"`
            ShipmentID string  `json:"shipment_id"`
            Phone string  `json:"phone"`
            Country string  `json:"country"`
            Name string  `json:"name"`
         
    }
    
    // SendOtpToCustomerResponse ...
    type SendOtpToCustomerResponse struct {

        
            RequestID string  `json:"request_id"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
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

        
            Reasons []BagReasons  `json:"reasons"`
            QcType []string  `json:"qc_type"`
            QuestionSet []QuestionSet  `json:"question_set"`
            DisplayName string  `json:"display_name"`
            ID float64  `json:"id"`
            Meta BagReasonMeta  `json:"meta"`
         
    }
    
    // ShipmentBagReasons ...
    type ShipmentBagReasons struct {

        
            Reasons []BagReasons  `json:"reasons"`
            Success bool  `json:"success"`
         
    }
    
    // ShipmentReason ...
    type ShipmentReason struct {

        
            ReasonText string  `json:"reason_text"`
            ReasonID float64  `json:"reason_id"`
            FeedbackType string  `json:"feedback_type"`
            ShowTextArea bool  `json:"show_text_area"`
            Priority float64  `json:"priority"`
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
    
    // ReasonsData ...
    type ReasonsData struct {

        
            Entities []EntitiesReasons  `json:"entities"`
            Products []ProductsReasons  `json:"products"`
         
    }
    
    // EntitiesDataUpdates ...
    type EntitiesDataUpdates struct {

        
            Filters []map[string]interface{}  `json:"filters"`
            Data map[string]interface{}  `json:"data"`
         
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
    
    // DataUpdates ...
    type DataUpdates struct {

        
            Entities []EntitiesDataUpdates  `json:"entities"`
            Products []ProductsDataUpdates  `json:"products"`
         
    }
    
    // ShipmentsRequest ...
    type ShipmentsRequest struct {

        
            Identifier string  `json:"identifier"`
            Products []Products  `json:"products"`
            Reasons ReasonsData  `json:"reasons"`
            DataUpdates DataUpdates  `json:"data_updates"`
         
    }
    
    // StatuesRequest ...
    type StatuesRequest struct {

        
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
            Shipments []ShipmentsRequest  `json:"shipments"`
            Status string  `json:"status"`
         
    }
    
    // UpdateShipmentStatusRequest ...
    type UpdateShipmentStatusRequest struct {

        
            LockAfterTransition bool  `json:"lock_after_transition"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
            Statuses []StatuesRequest  `json:"statuses"`
            Task bool  `json:"task"`
            ForceTransition bool  `json:"force_transition"`
         
    }
    
    // StatusesBodyResponse ...
    type StatusesBodyResponse struct {

        
            Shipments []map[string]interface{}  `json:"shipments"`
         
    }
    
    // ShipmentApplicationStatusResponse ...
    type ShipmentApplicationStatusResponse struct {

        
            Statuses []StatusesBodyResponse  `json:"statuses"`
         
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

        
            Key string  `json:"key"`
            Values []string  `json:"values"`
         
    }
    
    // CartPosCheckoutDetailRequest ...
    type CartPosCheckoutDetailRequest struct {

        
            OrderingStore float64  `json:"ordering_store"`
            Meta map[string]interface{}  `json:"meta"`
            Aggregator string  `json:"aggregator"`
            OrderType string  `json:"order_type"`
            BillingAddressID string  `json:"billing_address_id"`
            Pos bool  `json:"pos"`
            MerchantCode string  `json:"merchant_code"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            AddressID string  `json:"address_id"`
            Staff StaffCheckout  `json:"staff"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            CallbackURL string  `json:"callback_url"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            Files []Files  `json:"files"`
            PaymentMode string  `json:"payment_mode"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // CartDeliveryModesResponse ...
    type CartDeliveryModesResponse struct {

        
            PickupStores []float64  `json:"pickup_stores"`
            AvailableModes []string  `json:"available_modes"`
         
    }
    
    // PickupStoreDetail ...
    type PickupStoreDetail struct {

        
            Phone string  `json:"phone"`
            AddressType string  `json:"address_type"`
            State string  `json:"state"`
            AreaCode string  `json:"area_code"`
            Name string  `json:"name"`
            Country string  `json:"country"`
            Pincode float64  `json:"pincode"`
            Area string  `json:"area"`
            StoreCode string  `json:"store_code"`
            UID float64  `json:"uid"`
            ID float64  `json:"id"`
            Email string  `json:"email"`
            City string  `json:"city"`
            Address string  `json:"address"`
            Landmark string  `json:"landmark"`
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
    
