package application



    
    // NetQuantity ...
    type NetQuantity struct {

        
            Value float64  `json:"value"`
            Unit interface{}  `json:"unit"`
         
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

        
            Type string  `json:"type"`
            URL string  `json:"url"`
            Meta Meta  `json:"meta"`
            Alt string  `json:"alt"`
         
    }
    
    // ProductListingActionPage ...
    type ProductListingActionPage struct {

        
            Query map[string]interface{}  `json:"query"`
            Params map[string]interface{}  `json:"params"`
            Type string  `json:"type"`
         
    }
    
    // ProductListingAction ...
    type ProductListingAction struct {

        
            Type string  `json:"type"`
            Page ProductListingActionPage  `json:"page"`
         
    }
    
    // ProductBrand ...
    type ProductBrand struct {

        
            Name string  `json:"name"`
            Logo Media  `json:"logo"`
            Action ProductListingAction  `json:"action"`
            Description string  `json:"description"`
            UID float64  `json:"uid"`
         
    }
    
    // ApplicationItemMOQ ...
    type ApplicationItemMOQ struct {

        
            IncrementUnit float64  `json:"increment_unit"`
            Maximum float64  `json:"maximum"`
            Minimum float64  `json:"minimum"`
         
    }
    
    // ApplicationItemSEO ...
    type ApplicationItemSEO struct {

        
            Title interface{}  `json:"title"`
            Description interface{}  `json:"description"`
         
    }
    
    // ProductCategoryMap ...
    type ProductCategoryMap struct {

        
            L2 ProductBrand  `json:"l2"`
            L3 ProductBrand  `json:"l3"`
            L1 ProductBrand  `json:"l1"`
         
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

        
            Marked Price  `json:"marked"`
            Effective Price  `json:"effective"`
         
    }
    
    // ProductDetail ...
    type ProductDetail struct {

        
            Name string  `json:"name"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Similars []string  `json:"similars"`
            Tags []string  `json:"tags"`
            ShortDescription string  `json:"short_description"`
            ProductOnlineDate string  `json:"product_online_date"`
            ImageNature string  `json:"image_nature"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            ItemType string  `json:"item_type"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Brand ProductBrand  `json:"brand"`
            Moq ApplicationItemMOQ  `json:"moq"`
            IsDependent bool  `json:"is_dependent"`
            Medias []Media  `json:"medias"`
            TeaserTag string  `json:"teaser_tag"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            Categories []ProductBrand  `json:"categories"`
            Attributes map[string]interface{}  `json:"attributes"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Description string  `json:"description"`
            RatingCount float64  `json:"rating_count"`
            Highlights []string  `json:"highlights"`
            Seo ApplicationItemSEO  `json:"seo"`
            ItemCode string  `json:"item_code"`
            Action ProductListingAction  `json:"action"`
            HasVariant bool  `json:"has_variant"`
            Color string  `json:"color"`
            CategoryMap ProductCategoryMap  `json:"category_map"`
            Rating float64  `json:"rating"`
            Discount string  `json:"discount"`
            Tryouts []string  `json:"tryouts"`
            UID float64  `json:"uid"`
            Price ProductListingPrice  `json:"price"`
         
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

        
            Col2 ColumnHeader  `json:"col_2"`
            Col3 ColumnHeader  `json:"col_3"`
            Col4 ColumnHeader  `json:"col_4"`
            Col1 ColumnHeader  `json:"col_1"`
            Col6 ColumnHeader  `json:"col_6"`
            Col5 ColumnHeader  `json:"col_5"`
         
    }
    
    // SizeChartValues ...
    type SizeChartValues struct {

        
            Col2 string  `json:"col_2"`
            Col3 string  `json:"col_3"`
            Col4 string  `json:"col_4"`
            Col1 string  `json:"col_1"`
            Col6 string  `json:"col_6"`
            Col5 string  `json:"col_5"`
         
    }
    
    // SizeChart ...
    type SizeChart struct {

        
            SizeTip string  `json:"size_tip"`
            Headers ColumnHeaders  `json:"headers"`
            Unit string  `json:"unit"`
            Sizes []SizeChartValues  `json:"sizes"`
            Description string  `json:"description"`
            Image string  `json:"image"`
            Title string  `json:"title"`
         
    }
    
    // ProductSizeStores ...
    type ProductSizeStores struct {

        
            Count float64  `json:"count"`
         
    }
    
    // Weight ...
    type Weight struct {

        
            IsDefault bool  `json:"is_default"`
            Shipping float64  `json:"shipping"`
            Unit string  `json:"unit"`
         
    }
    
    // Dimension ...
    type Dimension struct {

        
            IsDefault bool  `json:"is_default"`
            Unit string  `json:"unit"`
            Length float64  `json:"length"`
            Width float64  `json:"width"`
            Height float64  `json:"height"`
         
    }
    
    // ProductSize ...
    type ProductSize struct {

        
            IsAvailable bool  `json:"is_available"`
            Display string  `json:"display"`
            Value string  `json:"value"`
            Weight Weight  `json:"weight"`
            SellerIdentifiers []string  `json:"seller_identifiers"`
            Quantity float64  `json:"quantity"`
            Dimension Dimension  `json:"dimension"`
         
    }
    
    // ProductSizes ...
    type ProductSizes struct {

        
            MultiSize bool  `json:"multi_size"`
            SizeChart SizeChart  `json:"size_chart"`
            Stores ProductSizeStores  `json:"stores"`
            Discount string  `json:"discount"`
            Sizes []ProductSize  `json:"sizes"`
            Price ProductListingPrice  `json:"price"`
            Sellable bool  `json:"sellable"`
         
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
            Items []ProductDetail  `json:"items"`
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
         
    }
    
    // ProductFrequentlyComparedSimilarResponse ...
    type ProductFrequentlyComparedSimilarResponse struct {

        
            Similars ProductCompareResponse  `json:"similars"`
         
    }
    
    // ProductVariantItemResponse ...
    type ProductVariantItemResponse struct {

        
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Medias []Media  `json:"medias"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            IsAvailable bool  `json:"is_available"`
            Value string  `json:"value"`
            Action ProductListingAction  `json:"action"`
            Color string  `json:"color"`
            UID float64  `json:"uid"`
            ColorName string  `json:"color_name"`
         
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
    
    // Seller ...
    type Seller struct {

        
            Count float64  `json:"count"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductStockPrice ...
    type ProductStockPrice struct {

        
            Currency string  `json:"currency"`
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
         
    }
    
    // StoreDetail ...
    type StoreDetail struct {

        
            Code string  `json:"code"`
            Name string  `json:"name"`
            City string  `json:"city"`
            ID float64  `json:"id"`
         
    }
    
    // CompanyDetail ...
    type CompanyDetail struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // ProductStockStatusItem ...
    type ProductStockStatusItem struct {

        
            Seller Seller  `json:"seller"`
            Size string  `json:"size"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
            UID string  `json:"uid"`
            Price ProductStockPrice  `json:"price"`
            Identifier map[string]interface{}  `json:"identifier"`
            Store StoreDetail  `json:"store"`
            Company CompanyDetail  `json:"company"`
         
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

        
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
            Value string  `json:"value"`
         
    }
    
    // ProductVariantListingResponse ...
    type ProductVariantListingResponse struct {

        
            Total float64  `json:"total"`
            Items []ProductVariantItemResponse  `json:"items"`
            DisplayType string  `json:"display_type"`
            Key string  `json:"key"`
            Header string  `json:"header"`
         
    }
    
    // ProductListingDetail ...
    type ProductListingDetail struct {

        
            Name string  `json:"name"`
            NetQuantity NetQuantity  `json:"net_quantity"`
            Similars []string  `json:"similars"`
            Tags []string  `json:"tags"`
            ShortDescription string  `json:"short_description"`
            Identifiers []string  `json:"identifiers"`
            ProductOnlineDate string  `json:"product_online_date"`
            Sizes []string  `json:"sizes"`
            ImageNature string  `json:"image_nature"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            ItemType string  `json:"item_type"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Brand ProductBrand  `json:"brand"`
            Moq ApplicationItemMOQ  `json:"moq"`
            IsDependent bool  `json:"is_dependent"`
            Medias []Media  `json:"medias"`
            Sellable bool  `json:"sellable"`
            TeaserTag string  `json:"teaser_tag"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            Categories []ProductBrand  `json:"categories"`
            Attributes map[string]interface{}  `json:"attributes"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Description string  `json:"description"`
            RatingCount float64  `json:"rating_count"`
            Highlights []string  `json:"highlights"`
            Seo ApplicationItemSEO  `json:"seo"`
            ItemCode string  `json:"item_code"`
            Action ProductListingAction  `json:"action"`
            HasVariant bool  `json:"has_variant"`
            Color string  `json:"color"`
            CategoryMap ProductCategoryMap  `json:"category_map"`
            Rating float64  `json:"rating"`
            Discount string  `json:"discount"`
            Tryouts []string  `json:"tryouts"`
            UID float64  `json:"uid"`
            Price ProductListingPrice  `json:"price"`
            Variants []ProductVariantListingResponse  `json:"variants"`
         
    }
    
    // ProductFiltersKey ...
    type ProductFiltersKey struct {

        
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            Display string  `json:"display"`
            Kind string  `json:"kind"`
         
    }
    
    // ProductFiltersValue ...
    type ProductFiltersValue struct {

        
            Min float64  `json:"min"`
            IsSelected bool  `json:"is_selected"`
            Max float64  `json:"max"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            DisplayFormat string  `json:"display_format"`
            Display string  `json:"display"`
            Value string  `json:"value"`
            QueryFormat string  `json:"query_format"`
            SelectedMax float64  `json:"selected_max"`
            SelectedMin float64  `json:"selected_min"`
            Count float64  `json:"count"`
         
    }
    
    // ProductFilters ...
    type ProductFilters struct {

        
            Key ProductFiltersKey  `json:"key"`
            Values []ProductFiltersValue  `json:"values"`
         
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

        
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Departments []string  `json:"departments"`
            Logo Media  `json:"logo"`
            Action ProductListingAction  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            Discount string  `json:"discount"`
            Description string  `json:"description"`
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
            Logo Media  `json:"logo"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Description string  `json:"description"`
            UID float64  `json:"uid"`
         
    }
    
    // CategoryBanner ...
    type CategoryBanner struct {

        
            Landscape Media  `json:"landscape"`
            Portrait Media  `json:"portrait"`
         
    }
    
    // ThirdLevelChild ...
    type ThirdLevelChild struct {

        
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Action ProductListingAction  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Childs []map[string]interface{}  `json:"childs"`
         
    }
    
    // SecondLevelChild ...
    type SecondLevelChild struct {

        
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Action ProductListingAction  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Childs []ThirdLevelChild  `json:"childs"`
         
    }
    
    // Child ...
    type Child struct {

        
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Action ProductListingAction  `json:"action"`
            Banners ImageUrls  `json:"banners"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            UID float64  `json:"uid"`
            Childs []SecondLevelChild  `json:"childs"`
         
    }
    
    // CategoryItems ...
    type CategoryItems struct {

        
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Action ProductListingAction  `json:"action"`
            Banners CategoryBanner  `json:"banners"`
            UID float64  `json:"uid"`
            Childs []Child  `json:"childs"`
         
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

        
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            PriorityOrder float64  `json:"priority_order"`
            Logo Media  `json:"logo"`
            UID float64  `json:"uid"`
         
    }
    
    // DepartmentResponse ...
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // AutocompleteItem ...
    type AutocompleteItem struct {

        
            Type string  `json:"type"`
            Logo Media  `json:"logo"`
            Display string  `json:"display"`
            Action ProductListingAction  `json:"action"`
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

        
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            SortOn string  `json:"sort_on"`
            Meta map[string]interface{}  `json:"meta"`
            Banners ImageUrls  `json:"banners"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Badge map[string]interface{}  `json:"badge"`
            Tag []string  `json:"tag"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            AllowFacets bool  `json:"allow_facets"`
            Logo Media  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Description string  `json:"description"`
            Cron map[string]interface{}  `json:"cron"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Query []CollectionQuery  `json:"query"`
            Action ProductListingAction  `json:"action"`
            AllowSort bool  `json:"allow_sort"`
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

        
            Items []GetCollectionDetailNest  `json:"items"`
            Page Page  `json:"page"`
            Filters CollectionListingFilter  `json:"filters"`
         
    }
    
    // CollectionDetailResponse ...
    type CollectionDetailResponse struct {

        
            Name string  `json:"name"`
            Priority float64  `json:"priority"`
            SortOn string  `json:"sort_on"`
            Meta map[string]interface{}  `json:"meta"`
            Banners ImageUrls  `json:"banners"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Badge map[string]interface{}  `json:"badge"`
            Tag []string  `json:"tag"`
            Type string  `json:"type"`
            Slug string  `json:"slug"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            AllowFacets bool  `json:"allow_facets"`
            Logo Media  `json:"logo"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Description string  `json:"description"`
            Cron map[string]interface{}  `json:"cron"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Query []CollectionQuery  `json:"query"`
            AllowSort bool  `json:"allow_sort"`
         
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

        
            Type string  `json:"type"`
            Coordinates []float64  `json:"coordinates"`
         
    }
    
    // Store ...
    type Store struct {

        
            State string  `json:"state"`
            Address string  `json:"address"`
            Name string  `json:"name"`
            City string  `json:"city"`
            LatLong LatLong  `json:"lat_long"`
            Pincode float64  `json:"pincode"`
            StoreCode string  `json:"store_code"`
            UID float64  `json:"uid"`
            Country string  `json:"country"`
            StoreEmail string  `json:"store_email"`
         
    }
    
    // StoreListingResponse ...
    type StoreListingResponse struct {

        
            Items []Store  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // StoreAddressSerializer ...
    type StoreAddressSerializer struct {

        
            State string  `json:"state"`
            Longitude float64  `json:"longitude"`
            Address2 string  `json:"address2"`
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
            Latitude float64  `json:"latitude"`
            Address1 string  `json:"address1"`
            Landmark string  `json:"landmark"`
            Country string  `json:"country"`
         
    }
    
    // StoreDepartments ...
    type StoreDepartments struct {

        
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            PriorityOrder float64  `json:"priority_order"`
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

        
            Name string  `json:"name"`
            Email string  `json:"email"`
            MobileNo SellerPhoneNumber  `json:"mobile_no"`
         
    }
    
    // CompanyStore ...
    type CompanyStore struct {

        
            CompanyType string  `json:"company_type"`
            Name string  `json:"name"`
            BusinessType string  `json:"business_type"`
            UID float64  `json:"uid"`
         
    }
    
    // AppStore ...
    type AppStore struct {

        
            Address StoreAddressSerializer  `json:"address"`
            Name string  `json:"name"`
            Departments []StoreDepartments  `json:"departments"`
            UID float64  `json:"uid"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            Manager StoreManagerSerializer  `json:"manager"`
            Company CompanyStore  `json:"company"`
         
    }
    
    // ApplicationStoreListing ...
    type ApplicationStoreListing struct {

        
            Items []AppStore  `json:"items"`
            Page Page  `json:"page"`
            Filters []StoreDepartments  `json:"filters"`
         
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

        
            Address StoreAddressSerializer  `json:"address"`
            Name string  `json:"name"`
            Departments []StoreDepartments  `json:"departments"`
            Timing []StoreTiming  `json:"timing"`
            UID float64  `json:"uid"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Manager StoreManagerSerializer  `json:"manager"`
            Company CompanyStore  `json:"company"`
         
    }
    
    // UserDetail ...
    type UserDetail struct {

        
            Contact string  `json:"contact"`
            SuperUser bool  `json:"super_user"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
         
    }
    
    // ProductDetails ...
    type ProductDetails struct {

        
            Name interface{}  `json:"name"`
            ShortDescription interface{}  `json:"short_description"`
            OutOfStock bool  `json:"out_of_stock"`
            ImageNature interface{}  `json:"image_nature"`
            GroupedAttributes map[string]interface{}  `json:"grouped_attributes"`
            TemplateTag interface{}  `json:"template_tag"`
            Media []map[string]interface{}  `json:"media"`
            IsSet bool  `json:"is_set"`
            BrandUID float64  `json:"brand_uid"`
            CountryOfOrigin interface{}  `json:"country_of_origin"`
            Slug interface{}  `json:"slug"`
            Attributes map[string]interface{}  `json:"attributes"`
            Description interface{}  `json:"description"`
            RatingCount float64  `json:"rating_count"`
            Highlights []interface{}  `json:"highlights"`
            Identifier map[string]interface{}  `json:"identifier"`
            HsnCode float64  `json:"hsn_code"`
            ItemCode interface{}  `json:"item_code"`
            Images []interface{}  `json:"images"`
            HasVariant bool  `json:"has_variant"`
            Rating float64  `json:"rating"`
         
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
            MaxMarked float64  `json:"max_marked"`
            MaxEffective float64  `json:"max_effective"`
         
    }
    
    // ProductInGroup ...
    type ProductInGroup struct {

        
            MaxQuantity float64  `json:"max_quantity"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            AllowRemove bool  `json:"allow_remove"`
            ProductDetails ProductDetails  `json:"product_details"`
            AutoSelect bool  `json:"auto_select"`
            Sizes []Size  `json:"sizes"`
            Price ProductGroupPrice  `json:"price"`
            ProductUID float64  `json:"product_uid"`
            MinQuantity float64  `json:"min_quantity"`
         
    }
    
    // ProductGroupingModel ...
    type ProductGroupingModel struct {

        
            ID interface{}  `json:"_id"`
            CompanyID float64  `json:"company_id"`
            Name interface{}  `json:"name"`
            ModifiedBy UserDetail  `json:"modified_by"`
            VerifiedBy UserDetail  `json:"verified_by"`
            Choice interface{}  `json:"choice"`
            Slug interface{}  `json:"slug"`
            Products []ProductInGroup  `json:"products"`
            IsActive bool  `json:"is_active"`
            PageVisibility []interface{}  `json:"page_visibility"`
            CreatedOn string  `json:"created_on"`
            Logo string  `json:"logo"`
            ModifiedOn string  `json:"modified_on"`
            VerifiedOn string  `json:"verified_on"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            CreatedBy UserDetail  `json:"created_by"`
            Meta map[string]interface{}  `json:"meta"`
         
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
    
    // SellerV3 ...
    type SellerV3 struct {

        
            Count float64  `json:"count"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // ReturnConfigSchemaV3 ...
    type ReturnConfigSchemaV3 struct {

        
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
            Unit string  `json:"unit"`
         
    }
    
    // ProductStockPriceV3 ...
    type ProductStockPriceV3 struct {

        
            Currency string  `json:"currency"`
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
         
    }
    
    // StoreV3 ...
    type StoreV3 struct {

        
            Count float64  `json:"count"`
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // MarketPlaceSttributesSchemaV3 ...
    type MarketPlaceSttributesSchemaV3 struct {

        
            Details []DetailsSchemaV3  `json:"details"`
            Title string  `json:"title"`
         
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
    
    // StrategyWiseListingSchemaV3 ...
    type StrategyWiseListingSchemaV3 struct {

        
            Distance float64  `json:"distance"`
            Pincode float64  `json:"pincode"`
            Tat float64  `json:"tat"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // ArticleAssignmentV3 ...
    type ArticleAssignmentV3 struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // ProductStockUnitPriceV3 ...
    type ProductStockUnitPriceV3 struct {

        
            Price float64  `json:"price"`
            Unit string  `json:"unit"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ProductSizePriceResponseV3 ...
    type ProductSizePriceResponseV3 struct {

        
            GroupedAttributes []SellerGroupAttributes  `json:"grouped_attributes"`
            ItemType string  `json:"item_type"`
            Seller SellerV3  `json:"seller"`
            ReturnConfig ReturnConfigSchemaV3  `json:"return_config"`
            IsGift bool  `json:"is_gift"`
            Pincode float64  `json:"pincode"`
            SellerCount float64  `json:"seller_count"`
            PricePerPiece ProductStockPriceV3  `json:"price_per_piece"`
            Store StoreV3  `json:"store"`
            ArticleID string  `json:"article_id"`
            MarketplaceAttributes []MarketPlaceSttributesSchemaV3  `json:"marketplace_attributes"`
            Quantity float64  `json:"quantity"`
            SpecialBadge string  `json:"special_badge"`
            IsCod bool  `json:"is_cod"`
            LongLat []float64  `json:"long_lat"`
            Set ProductSetV3  `json:"set"`
            StrategyWiseListing []StrategyWiseListingSchemaV3  `json:"strategy_wise_listing"`
            ArticleAssignment ArticleAssignmentV3  `json:"article_assignment"`
            Discount string  `json:"discount"`
            Price ProductStockPriceV3  `json:"price"`
            PricePerUnit ProductStockUnitPriceV3  `json:"price_per_unit"`
         
    }
    
    // ProductSizeSellerFilterSchemaV3 ...
    type ProductSizeSellerFilterSchemaV3 struct {

        
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
            Value string  `json:"value"`
         
    }
    
    // ProductSizeSellersResponseV3 ...
    type ProductSizeSellersResponseV3 struct {

        
            Items []ProductSizePriceResponseV3  `json:"items"`
            Page Page  `json:"page"`
            SortOn []ProductSizeSellerFilterSchemaV3  `json:"sort_on"`
         
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
            Type string  `json:"type"`
            Query ActionQuery  `json:"query"`
         
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
    
    // CartProduct ...
    type CartProduct struct {

        
            Images []ProductImage  `json:"images"`
            UID float64  `json:"uid"`
            Action ProductAction  `json:"action"`
            Type string  `json:"type"`
            Name string  `json:"name"`
            Brand BaseInfo  `json:"brand"`
            Slug string  `json:"slug"`
            Categories []CategoryInfo  `json:"categories"`
         
    }
    
    // ProductPrice ...
    type ProductPrice struct {

        
            CurrencyCode string  `json:"currency_code"`
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            Selling float64  `json:"selling"`
            Effective float64  `json:"effective"`
            AddOn float64  `json:"add_on"`
         
    }
    
    // ProductPriceInfo ...
    type ProductPriceInfo struct {

        
            Converted ProductPrice  `json:"converted"`
            Base ProductPrice  `json:"base"`
         
    }
    
    // ProductAvailability ...
    type ProductAvailability struct {

        
            IsValid bool  `json:"is_valid"`
            Sizes []string  `json:"sizes"`
            Deliverable bool  `json:"deliverable"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            OutOfStock bool  `json:"out_of_stock"`
         
    }
    
    // BasePrice ...
    type BasePrice struct {

        
            Marked float64  `json:"marked"`
            Effective float64  `json:"effective"`
            CurrencySymbol string  `json:"currency_symbol"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ArticlePriceInfo ...
    type ArticlePriceInfo struct {

        
            Converted BasePrice  `json:"converted"`
            Base BasePrice  `json:"base"`
         
    }
    
    // ProductArticle ...
    type ProductArticle struct {

        
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Price ArticlePriceInfo  `json:"price"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            UID string  `json:"uid"`
            Size string  `json:"size"`
            Type string  `json:"type"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Quantity float64  `json:"quantity"`
            Seller BaseInfo  `json:"seller"`
            Store BaseInfo  `json:"store"`
         
    }
    
    // CartProductIdentifer ...
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // PromoMeta ...
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // Ownership ...
    type Ownership struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // AppliedPromotion ...
    type AppliedPromotion struct {

        
            OfferText string  `json:"offer_text"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromotionType string  `json:"promotion_type"`
            Ownership Ownership  `json:"ownership"`
            MrpPromotion bool  `json:"mrp_promotion"`
            Amount float64  `json:"amount"`
            PromoID string  `json:"promo_id"`
         
    }
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            IsSet bool  `json:"is_set"`
            Product CartProduct  `json:"product"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Price ProductPriceInfo  `json:"price"`
            Key string  `json:"key"`
            Message string  `json:"message"`
            CouponMessage string  `json:"coupon_message"`
            Availability ProductAvailability  `json:"availability"`
            Discount string  `json:"discount"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Article ProductArticle  `json:"article"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            PromoMeta PromoMeta  `json:"promo_meta"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            Quantity float64  `json:"quantity"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
         
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
    
    // RawBreakup ...
    type RawBreakup struct {

        
            ConvenienceFee float64  `json:"convenience_fee"`
            FyndCash float64  `json:"fynd_cash"`
            Coupon float64  `json:"coupon"`
            Subtotal float64  `json:"subtotal"`
            GstCharges float64  `json:"gst_charges"`
            Vog float64  `json:"vog"`
            CodCharge float64  `json:"cod_charge"`
            YouSaved float64  `json:"you_saved"`
            Discount float64  `json:"discount"`
            DeliveryCharge float64  `json:"delivery_charge"`
            Total float64  `json:"total"`
            MrpTotal float64  `json:"mrp_total"`
         
    }
    
    // LoyaltyPoints ...
    type LoyaltyPoints struct {

        
            IsApplied bool  `json:"is_applied"`
            Description string  `json:"description"`
            Applicable float64  `json:"applicable"`
            Total float64  `json:"total"`
         
    }
    
    // DisplayBreakup ...
    type DisplayBreakup struct {

        
            CurrencyCode string  `json:"currency_code"`
            Message []string  `json:"message"`
            Key string  `json:"key"`
            CurrencySymbol string  `json:"currency_symbol"`
            Value float64  `json:"value"`
            Display string  `json:"display"`
         
    }
    
    // CouponBreakup ...
    type CouponBreakup struct {

        
            IsApplied bool  `json:"is_applied"`
            Message string  `json:"message"`
            UID string  `json:"uid"`
            Value float64  `json:"value"`
            Code string  `json:"code"`
            Type string  `json:"type"`
         
    }
    
    // CartBreakup ...
    type CartBreakup struct {

        
            Raw RawBreakup  `json:"raw"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Display []DisplayBreakup  `json:"display"`
            Coupon CouponBreakup  `json:"coupon"`
         
    }
    
    // CartDetailResponse ...
    type CartDetailResponse struct {

        
            IsValid bool  `json:"is_valid"`
            Message string  `json:"message"`
            ID string  `json:"id"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            LastModified string  `json:"last_modified"`
            Currency CartCurrency  `json:"currency"`
            Items []CartProductInfo  `json:"items"`
            Comment string  `json:"comment"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            CheckoutMode string  `json:"checkout_mode"`
            BuyNow bool  `json:"buy_now"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Gstin string  `json:"gstin"`
            CouponText string  `json:"coupon_text"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            StoreID float64  `json:"store_id"`
            ItemSize string  `json:"item_size"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            ItemID float64  `json:"item_id"`
            SellerID float64  `json:"seller_id"`
            Display string  `json:"display"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ProductGroupTags []string  `json:"product_group_tags"`
            ArticleID string  `json:"article_id"`
            Quantity float64  `json:"quantity"`
            Pos bool  `json:"pos"`
         
    }
    
    // AddCartRequest ...
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
         
    }
    
    // AddCartDetailResponse ...
    type AddCartDetailResponse struct {

        
            Success bool  `json:"success"`
            Partial bool  `json:"partial"`
            Message string  `json:"message"`
            Cart CartDetailResponse  `json:"cart"`
         
    }
    
    // UpdateProductCart ...
    type UpdateProductCart struct {

        
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ItemSize string  `json:"item_size"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ItemID float64  `json:"item_id"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ItemIndex float64  `json:"item_index"`
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

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Cart CartDetailResponse  `json:"cart"`
         
    }
    
    // CartItemCountResponse ...
    type CartItemCountResponse struct {

        
            UserCartItemsCount float64  `json:"user_cart_items_count"`
         
    }
    
    // PageCoupon ...
    type PageCoupon struct {

        
            HasPrevious bool  `json:"has_previous"`
            Current float64  `json:"current"`
            Total float64  `json:"total"`
            TotalItemCount float64  `json:"total_item_count"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // Coupon ...
    type Coupon struct {

        
            IsApplied bool  `json:"is_applied"`
            Message string  `json:"message"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            SubTitle string  `json:"sub_title"`
            IsApplicable bool  `json:"is_applicable"`
            CouponCode string  `json:"coupon_code"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            CouponValue float64  `json:"coupon_value"`
            ExpiresOn string  `json:"expires_on"`
            Title string  `json:"title"`
         
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
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
            BulkEffective float64  `json:"bulk_effective"`
         
    }
    
    // OfferItem ...
    type OfferItem struct {

        
            Price OfferPrice  `json:"price"`
            Type string  `json:"type"`
            Margin float64  `json:"margin"`
            Best bool  `json:"best"`
            Total float64  `json:"total"`
            Quantity float64  `json:"quantity"`
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

        
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // Address ...
    type Address struct {

        
            CountryPhoneCode string  `json:"country_phone_code"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            Area string  `json:"area"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Address string  `json:"address"`
            City string  `json:"city"`
            ID string  `json:"id"`
            IsActive bool  `json:"is_active"`
            CountryIsoCode string  `json:"country_iso_code"`
            AreaCode string  `json:"area_code"`
            UserID string  `json:"user_id"`
            Name string  `json:"name"`
            Landmark string  `json:"landmark"`
            GeoLocation GeoLocation  `json:"geo_location"`
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
            Email string  `json:"email"`
            State string  `json:"state"`
            Meta map[string]interface{}  `json:"meta"`
            IsDefaultAddress bool  `json:"is_default_address"`
            Phone string  `json:"phone"`
            CheckoutMode string  `json:"checkout_mode"`
            Tags []string  `json:"tags"`
            AddressType string  `json:"address_type"`
         
    }
    
    // GetAddressesResponse ...
    type GetAddressesResponse struct {

        
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
            IsUpdated bool  `json:"is_updated"`
            IsDefaultAddress bool  `json:"is_default_address"`
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
            ID string  `json:"id"`
            CartID string  `json:"cart_id"`
         
    }
    
    // UpdateCartPaymentRequest ...
    type UpdateCartPaymentRequest struct {

        
            ID string  `json:"id"`
            AggregatorName string  `json:"aggregator_name"`
            MerchantCode string  `json:"merchant_code"`
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentMode string  `json:"payment_mode"`
            AddressID string  `json:"address_id"`
         
    }
    
    // CouponValidity ...
    type CouponValidity struct {

        
            DisplayMessageEn string  `json:"display_message_en"`
            Discount float64  `json:"discount"`
            Valid bool  `json:"valid"`
            Code string  `json:"code"`
            Title string  `json:"title"`
         
    }
    
    // PaymentCouponValidate ...
    type PaymentCouponValidate struct {

        
            CouponValidity CouponValidity  `json:"coupon_validity"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // ShipmentResponse ...
    type ShipmentResponse struct {

        
            BoxType string  `json:"box_type"`
            Promise ShipmentPromise  `json:"promise"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            DpID string  `json:"dp_id"`
            Items []CartProductInfo  `json:"items"`
            ShipmentType string  `json:"shipment_type"`
            FulfillmentType string  `json:"fulfillment_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
            OrderType string  `json:"order_type"`
            Shipments float64  `json:"shipments"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            IsValid bool  `json:"is_valid"`
            Message string  `json:"message"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            UID string  `json:"uid"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            LastModified string  `json:"last_modified"`
            Comment string  `json:"comment"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            BuyNow bool  `json:"buy_now"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Gstin string  `json:"gstin"`
            ID string  `json:"id"`
            CartID float64  `json:"cart_id"`
            Currency CartCurrency  `json:"currency"`
            Error bool  `json:"error"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CheckoutMode string  `json:"checkout_mode"`
            Shipments []ShipmentResponse  `json:"shipments"`
            CouponText string  `json:"coupon_text"`
         
    }
    
    // CartCheckoutCustomMeta ...
    type CartCheckoutCustomMeta struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // StaffCheckout ...
    type StaffCheckout struct {

        
            ID string  `json:"_id"`
            FirstName string  `json:"first_name"`
            EmployeeCode string  `json:"employee_code"`
            LastName string  `json:"last_name"`
            User string  `json:"user"`
         
    }
    
    // CartCheckoutDetailRequest ...
    type CartCheckoutDetailRequest struct {

        
            OrderingStore float64  `json:"ordering_store"`
            CustomMeta []CartCheckoutCustomMeta  `json:"custom_meta"`
            Aggregator string  `json:"aggregator"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Meta map[string]interface{}  `json:"meta"`
            BillingAddressID string  `json:"billing_address_id"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            CallbackURL string  `json:"callback_url"`
            MerchantCode string  `json:"merchant_code"`
            PaymentIdentifier string  `json:"payment_identifier"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            PaymentMode string  `json:"payment_mode"`
            Staff StaffCheckout  `json:"staff"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            AddressID string  `json:"address_id"`
         
    }
    
    // CheckCart ...
    type CheckCart struct {

        
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            UID string  `json:"uid"`
            Items []CartProductInfo  `json:"items"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            CodCharges float64  `json:"cod_charges"`
            ID string  `json:"id"`
            CartID float64  `json:"cart_id"`
            CodMessage string  `json:"cod_message"`
            ErrorMessage string  `json:"error_message"`
            Currency CartCurrency  `json:"currency"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            DeliveryCharges float64  `json:"delivery_charges"`
            IsValid bool  `json:"is_valid"`
            Message string  `json:"message"`
            CouponText string  `json:"coupon_text"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            LastModified string  `json:"last_modified"`
            Comment string  `json:"comment"`
            Success bool  `json:"success"`
            BuyNow bool  `json:"buy_now"`
            Gstin string  `json:"gstin"`
            OrderID string  `json:"order_id"`
            CodAvailable bool  `json:"cod_available"`
            CheckoutMode string  `json:"checkout_mode"`
            UserType string  `json:"user_type"`
            StoreCode string  `json:"store_code"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
         
    }
    
    // CartCheckoutResponse ...
    type CartCheckoutResponse struct {

        
            Message string  `json:"message"`
            OrderID string  `json:"order_id"`
            Data map[string]interface{}  `json:"data"`
            CallbackURL string  `json:"callback_url"`
            AppInterceptURL string  `json:"app_intercept_url"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Success bool  `json:"success"`
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
            Source map[string]interface{}  `json:"source"`
            CreatedOn string  `json:"created_on"`
            User map[string]interface{}  `json:"user"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            UID string  `json:"uid"`
            Items []CartProductInfo  `json:"items"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            ID string  `json:"id"`
            CartID float64  `json:"cart_id"`
            Currency CartCurrency  `json:"currency"`
            IsValid bool  `json:"is_valid"`
            Message string  `json:"message"`
            CouponText string  `json:"coupon_text"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            LastModified string  `json:"last_modified"`
            Comment string  `json:"comment"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            BuyNow bool  `json:"buy_now"`
            Gstin string  `json:"gstin"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Error string  `json:"error"`
            Cart SharedCart  `json:"cart"`
         
    }
    
    // PromotionOffer ...
    type PromotionOffer struct {

        
            OfferText string  `json:"offer_text"`
            PromotionGroup string  `json:"promotion_group"`
            ID string  `json:"id"`
            Description string  `json:"description"`
            ValidTill string  `json:"valid_till"`
         
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

        
            CurrencyCode string  `json:"currency_code"`
            OfferPrice float64  `json:"offer_price"`
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
         
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

        
            OfferText string  `json:"offer_text"`
            PromotionGroup string  `json:"promotion_group"`
            ID string  `json:"id"`
            OfferPrices []LadderOfferItem  `json:"offer_prices"`
            Description string  `json:"description"`
            ValidTill string  `json:"valid_till"`
         
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

        
            ConfigType string  `json:"config_type"`
            Key string  `json:"key"`
            VerifyAPI string  `json:"verify_api"`
            Pin string  `json:"pin"`
            Secret string  `json:"secret"`
            MerchantID string  `json:"merchant_id"`
            UserID string  `json:"user_id"`
            Sdk bool  `json:"sdk"`
            API string  `json:"api"`
            MerchantKey string  `json:"merchant_key"`
         
    }
    
    // AggregatorsConfigDetailResponse ...
    type AggregatorsConfigDetailResponse struct {

        
            Env string  `json:"env"`
            Ccavenue AggregatorConfigDetail  `json:"ccavenue"`
            Juspay AggregatorConfigDetail  `json:"juspay"`
            Mswipe AggregatorConfigDetail  `json:"mswipe"`
            Razorpay AggregatorConfigDetail  `json:"razorpay"`
            Simpl AggregatorConfigDetail  `json:"simpl"`
            Payumoney AggregatorConfigDetail  `json:"payumoney"`
            Stripe AggregatorConfigDetail  `json:"stripe"`
            Success bool  `json:"success"`
            Rupifi AggregatorConfigDetail  `json:"rupifi"`
         
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
            Nickname string  `json:"nickname"`
            NameOnCard string  `json:"name_on_card"`
            Refresh bool  `json:"refresh"`
         
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
            Aggregator string  `json:"aggregator"`
            API string  `json:"api"`
         
    }
    
    // ActiveCardPaymentGatewayResponse ...
    type ActiveCardPaymentGatewayResponse struct {

        
            Cards CardPaymentGateway  `json:"cards"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // Card ...
    type Card struct {

        
            CardReference string  `json:"card_reference"`
            CardFingerprint string  `json:"card_fingerprint"`
            CardIsin string  `json:"card_isin"`
            CardBrand string  `json:"card_brand"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            ExpYear float64  `json:"exp_year"`
            AggregatorName string  `json:"aggregator_name"`
            CardIssuer string  `json:"card_issuer"`
            Nickname string  `json:"nickname"`
            CardToken string  `json:"card_token"`
            CardName string  `json:"card_name"`
            Expired bool  `json:"expired"`
            ExpMonth float64  `json:"exp_month"`
            CardBrandImage string  `json:"card_brand_image"`
            CardNumber string  `json:"card_number"`
            CardID string  `json:"card_id"`
            CardType string  `json:"card_type"`
         
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

        
            MerchantParams map[string]interface{}  `json:"merchant_params"`
            TransactionAmountInPaise float64  `json:"transaction_amount_in_paise"`
            PhoneNumber string  `json:"phone_number"`
            Aggregator string  `json:"aggregator"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            Payload string  `json:"payload"`
            OrderItems []map[string]interface{}  `json:"order_items"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
         
    }
    
    // ValidateCustomerResponse ...
    type ValidateCustomerResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // ChargeCustomerRequest ...
    type ChargeCustomerRequest struct {

        
            TransactionToken string  `json:"transaction_token"`
            Aggregator string  `json:"aggregator"`
            OrderID string  `json:"order_id"`
            Verified bool  `json:"verified"`
            Amount float64  `json:"amount"`
         
    }
    
    // ChargeCustomerResponse ...
    type ChargeCustomerResponse struct {

        
            Aggregator string  `json:"aggregator"`
            OrderID string  `json:"order_id"`
            Message string  `json:"message"`
            CartID string  `json:"cart_id"`
            DeliveryAddressID string  `json:"delivery_address_id"`
            Success bool  `json:"success"`
            Status string  `json:"status"`
         
    }
    
    // PaymentInitializationRequest ...
    type PaymentInitializationRequest struct {

        
            Email string  `json:"email"`
            Vpa string  `json:"vpa"`
            Aggregator string  `json:"aggregator"`
            MerchantOrderID string  `json:"merchant_order_id"`
            OrderID string  `json:"order_id"`
            Currency string  `json:"currency"`
            Timeout float64  `json:"timeout"`
            Contact string  `json:"contact"`
            Amount float64  `json:"amount"`
            CustomerID string  `json:"customer_id"`
            Method string  `json:"method"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
         
    }
    
    // PaymentInitializationResponse ...
    type PaymentInitializationResponse struct {

        
            Vpa string  `json:"vpa"`
            PollingURL string  `json:"polling_url"`
            UpiPollURL string  `json:"upi_poll_url"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Aggregator string  `json:"aggregator"`
            VirtualID string  `json:"virtual_id"`
            Currency string  `json:"currency"`
            Timeout float64  `json:"timeout"`
            Amount float64  `json:"amount"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            CustomerID string  `json:"customer_id"`
            Success bool  `json:"success"`
            BqrImage string  `json:"bqr_image"`
            Status string  `json:"status"`
            Method string  `json:"method"`
         
    }
    
    // PaymentStatusUpdateRequest ...
    type PaymentStatusUpdateRequest struct {

        
            Email string  `json:"email"`
            Vpa string  `json:"vpa"`
            Aggregator string  `json:"aggregator"`
            MerchantOrderID string  `json:"merchant_order_id"`
            OrderID string  `json:"order_id"`
            Currency string  `json:"currency"`
            Contact string  `json:"contact"`
            Amount float64  `json:"amount"`
            CustomerID string  `json:"customer_id"`
            Status string  `json:"status"`
            Method string  `json:"method"`
         
    }
    
    // PaymentStatusUpdateResponse ...
    type PaymentStatusUpdateResponse struct {

        
            Retry bool  `json:"retry"`
            AggregatorName string  `json:"aggregator_name"`
            RedirectURL string  `json:"redirect_url"`
            Success bool  `json:"success"`
            Status string  `json:"status"`
         
    }
    
    // AggregatorRoute ...
    type AggregatorRoute struct {

        
            PaymentFlow string  `json:"payment_flow"`
            Data map[string]interface{}  `json:"data"`
            APILink string  `json:"api_link"`
            PaymentFlowData string  `json:"payment_flow_data"`
         
    }
    
    // PaymentFlow ...
    type PaymentFlow struct {

        
            Fynd AggregatorRoute  `json:"fynd"`
            BqrRazorpay AggregatorRoute  `json:"bqr_razorpay"`
            Juspay AggregatorRoute  `json:"juspay"`
            Ccavenue AggregatorRoute  `json:"ccavenue"`
            Jiopay AggregatorRoute  `json:"jiopay"`
            Mswipe AggregatorRoute  `json:"mswipe"`
            UpiRazorpay AggregatorRoute  `json:"upi_razorpay"`
            Razorpay AggregatorRoute  `json:"razorpay"`
            Simpl AggregatorRoute  `json:"simpl"`
            Epaylater AggregatorRoute  `json:"epaylater"`
            Stripe AggregatorRoute  `json:"stripe"`
            Rupifi AggregatorRoute  `json:"rupifi"`
            Payubiz AggregatorRoute  `json:"payubiz"`
         
    }
    
    // PaymentModeLogo ...
    type PaymentModeLogo struct {

        
            Large string  `json:"large"`
            Small string  `json:"small"`
         
    }
    
    // IntentApp ...
    type IntentApp struct {

        
            Code string  `json:"code"`
            DisplayName string  `json:"display_name"`
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

        
            CardIssuer string  `json:"card_issuer"`
            AggregatorName string  `json:"aggregator_name"`
            CardName string  `json:"card_name"`
            Expired bool  `json:"expired"`
            Name string  `json:"name"`
            ExpMonth float64  `json:"exp_month"`
            Timeout float64  `json:"timeout"`
            CardBrand string  `json:"card_brand"`
            CardID string  `json:"card_id"`
            IntentApp []IntentApp  `json:"intent_app"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            RetryCount float64  `json:"retry_count"`
            ExpYear float64  `json:"exp_year"`
            DisplayPriority float64  `json:"display_priority"`
            CardToken string  `json:"card_token"`
            FyndVpa string  `json:"fynd_vpa"`
            CodLimitPerOrder float64  `json:"cod_limit_per_order"`
            CodLimit float64  `json:"cod_limit"`
            CardType string  `json:"card_type"`
            DisplayName string  `json:"display_name"`
            CardReference string  `json:"card_reference"`
            CardIsin string  `json:"card_isin"`
            CompliantWithTokenisationGuidelines bool  `json:"compliant_with_tokenisation_guidelines"`
            CardBrandImage string  `json:"card_brand_image"`
            Code string  `json:"code"`
            CardFingerprint string  `json:"card_fingerprint"`
            IntentFlow bool  `json:"intent_flow"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            Nickname string  `json:"nickname"`
            MerchantCode string  `json:"merchant_code"`
            RemainingLimit float64  `json:"remaining_limit"`
            CardNumber string  `json:"card_number"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
         
    }
    
    // RootPaymentMode ...
    type RootPaymentMode struct {

        
            DisplayName string  `json:"display_name"`
            SaveCard bool  `json:"save_card"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            AggregatorName string  `json:"aggregator_name"`
            DisplayPriority float64  `json:"display_priority"`
            List []PaymentModeList  `json:"list"`
            Name string  `json:"name"`
            IsPayByCardPl bool  `json:"is_pay_by_card_pl"`
            AnonymousEnable bool  `json:"anonymous_enable"`
         
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
         
    }
    
    // LinkStatus ...
    type LinkStatus struct {

        
            Status bool  `json:"status"`
            Message string  `json:"message"`
         
    }
    
    // ResendOrCancelPaymentResponse ...
    type ResendOrCancelPaymentResponse struct {

        
            Data LinkStatus  `json:"data"`
            Success bool  `json:"success"`
         
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

        
            UpiVpa string  `json:"upi_vpa"`
            CustomerName string  `json:"customer_name"`
            Status string  `json:"status"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // ValidateVPAResponse ...
    type ValidateVPAResponse struct {

        
            Data ValidateUPI  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // TransferItemsDetails ...
    type TransferItemsDetails struct {

        
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
            Name string  `json:"name"`
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

        
            AccountHolder string  `json:"account_holder"`
            BankName string  `json:"bank_name"`
            ModifiedOn string  `json:"modified_on"`
            Address string  `json:"address"`
            IfscCode string  `json:"ifsc_code"`
            AccountNo string  `json:"account_no"`
            BranchName string  `json:"branch_name"`
            Email string  `json:"email"`
            Subtitle string  `json:"subtitle"`
            CreatedOn string  `json:"created_on"`
            TransferMode string  `json:"transfer_mode"`
            ID float64  `json:"id"`
            DisplayName string  `json:"display_name"`
            DelightsUserName string  `json:"delights_user_name"`
            IsActive bool  `json:"is_active"`
            BeneficiaryID string  `json:"beneficiary_id"`
            Title string  `json:"title"`
            Mobile string  `json:"mobile"`
            Comment string  `json:"comment"`
         
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

        
            Email string  `json:"email"`
            Mobile string  `json:"mobile"`
            IfscCode string  `json:"ifsc_code"`
            AccountHolder string  `json:"account_holder"`
            BankName string  `json:"bank_name"`
            AccountNo string  `json:"account_no"`
            Wallet string  `json:"wallet"`
            Address string  `json:"address"`
            Comment string  `json:"comment"`
            Vpa string  `json:"vpa"`
            BranchName string  `json:"branch_name"`
         
    }
    
    // AddBeneficiaryDetailsRequest ...
    type AddBeneficiaryDetailsRequest struct {

        
            ShipmentID string  `json:"shipment_id"`
            OrderID string  `json:"order_id"`
            TransferMode string  `json:"transfer_mode"`
            Otp string  `json:"otp"`
            RequestID string  `json:"request_id"`
            Delights bool  `json:"delights"`
            Details BeneficiaryModeDetails  `json:"details"`
         
    }
    
    // RefundAccountResponse ...
    type RefundAccountResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            IsVerifiedFlag bool  `json:"is_verified_flag"`
         
    }
    
    // BankDetailsForOTP ...
    type BankDetailsForOTP struct {

        
            AccountHolder string  `json:"account_holder"`
            BankName string  `json:"bank_name"`
            AccountNo string  `json:"account_no"`
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

        
            IsBeneficiarySet bool  `json:"is_beneficiary_set"`
            Success bool  `json:"success"`
         
    }
    
    // GetPaymentLinkResponse ...
    type GetPaymentLinkResponse struct {

        
            PaymentLinkURL string  `json:"payment_link_url"`
            PollingTimeout float64  `json:"polling_timeout"`
            ExternalOrderID string  `json:"external_order_id"`
            PaymentLinkCurrentStatus string  `json:"payment_link_current_status"`
            StatusCode float64  `json:"status_code"`
            Amount float64  `json:"amount"`
            MerchantName string  `json:"merchant_name"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // ErrorDescription ...
    type ErrorDescription struct {

        
            MerchantOrderID string  `json:"merchant_order_id"`
            PaymentTransactionID string  `json:"payment_transaction_id"`
            Expired bool  `json:"expired"`
            MerchantName string  `json:"merchant_name"`
            Amount float64  `json:"amount"`
            Msg string  `json:"msg"`
            Cancelled bool  `json:"cancelled"`
            InvalidID bool  `json:"invalid_id"`
         
    }
    
    // ErrorResponse ...
    type ErrorResponse struct {

        
            Error ErrorDescription  `json:"error"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            Message string  `json:"message"`
         
    }
    
    // CreatePaymentLinkMeta ...
    type CreatePaymentLinkMeta struct {

        
            CheckoutMode string  `json:"checkout_mode"`
            CartID string  `json:"cart_id"`
            AssignCardID string  `json:"assign_card_id"`
            Amount string  `json:"amount"`
            Pincode string  `json:"pincode"`
         
    }
    
    // CreatePaymentLinkRequest ...
    type CreatePaymentLinkRequest struct {

        
            Email string  `json:"email"`
            ExternalOrderID string  `json:"external_order_id"`
            Meta CreatePaymentLinkMeta  `json:"meta"`
            Description string  `json:"description"`
            Amount float64  `json:"amount"`
            MobileNumber string  `json:"mobile_number"`
         
    }
    
    // CreatePaymentLinkResponse ...
    type CreatePaymentLinkResponse struct {

        
            PaymentLinkURL string  `json:"payment_link_url"`
            PollingTimeout float64  `json:"polling_timeout"`
            StatusCode float64  `json:"status_code"`
            PaymentLinkID string  `json:"payment_link_id"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
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

        
            OrderID string  `json:"order_id"`
            HttpStatus float64  `json:"http_status"`
            AggregatorName string  `json:"aggregator_name"`
            Success bool  `json:"success"`
            StatusCode float64  `json:"status_code"`
            Amount float64  `json:"amount"`
            RedirectURL string  `json:"redirect_url"`
            PaymentLinkID string  `json:"payment_link_id"`
            Message string  `json:"message"`
            Status string  `json:"status"`
         
    }
    
    // PaymentMethodsMeta ...
    type PaymentMethodsMeta struct {

        
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentGateway string  `json:"payment_gateway"`
            MerchantCode string  `json:"merchant_code"`
         
    }
    
    // CreateOrderUserPaymentMethods ...
    type CreateOrderUserPaymentMethods struct {

        
            Meta PaymentMethodsMeta  `json:"meta"`
            Name string  `json:"name"`
            Mode string  `json:"mode"`
         
    }
    
    // CreateOrderUserRequest ...
    type CreateOrderUserRequest struct {

        
            PaymentMethods CreateOrderUserPaymentMethods  `json:"payment_methods"`
            SuccessCallbackURL string  `json:"success_callback_url"`
            FailureCallbackURL string  `json:"failure_callback_url"`
            Meta map[string]interface{}  `json:"meta"`
            Currency string  `json:"currency"`
            PaymentLinkID string  `json:"payment_link_id"`
         
    }
    
    // CreateOrderUserData ...
    type CreateOrderUserData struct {

        
            Email string  `json:"email"`
            Aggregator string  `json:"aggregator"`
            MerchantOrderID string  `json:"merchant_order_id"`
            OrderID string  `json:"order_id"`
            Currency string  `json:"currency"`
            Contact string  `json:"contact"`
            Amount float64  `json:"amount"`
            CustomerID string  `json:"customer_id"`
            CallbackURL string  `json:"callback_url"`
            Method string  `json:"method"`
         
    }
    
    // CreateOrderUserResponse ...
    type CreateOrderUserResponse struct {

        
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            Message string  `json:"message"`
            OrderID string  `json:"order_id"`
            StatusCode float64  `json:"status_code"`
            Data CreateOrderUserData  `json:"data"`
            Success bool  `json:"success"`
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

        
            Balance BalanceDetails  `json:"balance"`
            StatusMessage string  `json:"status_message"`
            Status string  `json:"status"`
            MerchantCustomerRefID string  `json:"merchant_customer_ref_id"`
         
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
            IsRegistered bool  `json:"is_registered"`
            Status bool  `json:"status"`
         
    }
    
    // CheckCreditResponse ...
    type CheckCreditResponse struct {

        
            Data CreditDetail  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // KYCAddress ...
    type KYCAddress struct {

        
            City string  `json:"city"`
            OwnershipType string  `json:"ownership_type"`
            State string  `json:"state"`
            LandMark string  `json:"land_mark"`
            Addressline2 string  `json:"addressline2"`
            Addressline1 string  `json:"addressline1"`
            Pincode string  `json:"pincode"`
         
    }
    
    // UserPersonalInfoInDetails ...
    type UserPersonalInfoInDetails struct {

        
            Email string  `json:"email"`
            Phone string  `json:"phone"`
            MiddleName string  `json:"middle_name"`
            FathersName string  `json:"fathers_name"`
            Passport string  `json:"passport"`
            DrivingLicense string  `json:"driving_license"`
            FirstName string  `json:"first_name"`
            VoterID string  `json:"voter_id"`
            Dob string  `json:"dob"`
            Gender string  `json:"gender"`
            LastName string  `json:"last_name"`
            Pan string  `json:"pan"`
            MothersName string  `json:"mothers_name"`
            EmailVerified bool  `json:"email_verified"`
            MobileVerified bool  `json:"mobile_verified"`
            AddressAsPerID KYCAddress  `json:"address_as_per_id"`
         
    }
    
    // DeviceDetails ...
    type DeviceDetails struct {

        
            DeviceMake string  `json:"device_make"`
            DeviceType string  `json:"device_type"`
            Os string  `json:"os"`
            OsVersion string  `json:"os_version"`
            DeviceModel string  `json:"device_model"`
            IdentifierType string  `json:"identifier_type"`
            IdentificationNumber string  `json:"identification_number"`
         
    }
    
    // BusinessDetails ...
    type BusinessDetails struct {

        
            Fssai string  `json:"fssai"`
            Gstin string  `json:"gstin"`
            Fda string  `json:"fda"`
            BusinessOwnershipType string  `json:"business_ownership_type"`
            Vintage string  `json:"vintage"`
            Name string  `json:"name"`
            BusinessType string  `json:"business_type"`
            EntityType string  `json:"entity_type"`
            Pan string  `json:"pan"`
            Address KYCAddress  `json:"address"`
            ShopAndEstablishment map[string]interface{}  `json:"shop_and_establishment"`
         
    }
    
    // MarketplaceInfo ...
    type MarketplaceInfo struct {

        
            Name string  `json:"name"`
            DateOfJoining string  `json:"date_of_joining"`
            MembershipID string  `json:"membership_id"`
         
    }
    
    // CustomerOnboardingRequest ...
    type CustomerOnboardingRequest struct {

        
            Aggregator string  `json:"aggregator"`
            PersonalInfo UserPersonalInfoInDetails  `json:"personal_info"`
            Mcc string  `json:"mcc"`
            Device DeviceDetails  `json:"device"`
            BusinessInfo BusinessDetails  `json:"business_info"`
            Source string  `json:"source"`
            MarketplaceInfo MarketplaceInfo  `json:"marketplace_info"`
         
    }
    
    // OnboardSummary ...
    type OnboardSummary struct {

        
            RedirectURL string  `json:"redirect_url"`
            Session map[string]interface{}  `json:"session"`
            Status bool  `json:"status"`
         
    }
    
    // CustomerOnboardingResponse ...
    type CustomerOnboardingResponse struct {

        
            Data OnboardSummary  `json:"data"`
            Success bool  `json:"success"`
         
    }
    

    
    // ShipmentUserInfo ...
    type ShipmentUserInfo struct {

        
            Gender string  `json:"gender"`
            FirstName string  `json:"first_name"`
            Mobile string  `json:"mobile"`
            LastName string  `json:"last_name"`
         
    }
    
    // ShipmentTotalDetails ...
    type ShipmentTotalDetails struct {

        
            TotalPrice float64  `json:"total_price"`
            Pieces float64  `json:"pieces"`
            Sizes float64  `json:"sizes"`
         
    }
    
    // ShipmentPayment ...
    type ShipmentPayment struct {

        
            DisplayName string  `json:"display_name"`
            PaymentMode string  `json:"payment_mode"`
            Status string  `json:"status"`
            Mop string  `json:"mop"`
            Logo string  `json:"logo"`
            Mode string  `json:"mode"`
         
    }
    
    // FulfillingStore ...
    type FulfillingStore struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
            CompanyID float64  `json:"company_id"`
            Code string  `json:"code"`
            CompanyName string  `json:"company_name"`
         
    }
    
    // Invoice ...
    type Invoice struct {

        
            UpdatedDate string  `json:"updated_date"`
            InvoiceURL string  `json:"invoice_url"`
            LabelURL string  `json:"label_url"`
         
    }
    
    // Identifiers ...
    type Identifiers struct {

        
            Ean string  `json:"ean"`
            SkuCode string  `json:"sku_code"`
         
    }
    
    // FinancialBreakup ...
    type FinancialBreakup struct {

        
            CouponValue float64  `json:"coupon_value"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            AmountPaid float64  `json:"amount_paid"`
            ItemName string  `json:"item_name"`
            Size string  `json:"size"`
            GstTag string  `json:"gst_tag"`
            Discount float64  `json:"discount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            GstFee float64  `json:"gst_fee"`
            PriceMarked float64  `json:"price_marked"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            RefundAmount float64  `json:"refund_amount"`
            PriceEffective float64  `json:"price_effective"`
            DeliveryCharge float64  `json:"delivery_charge"`
            TotalUnits float64  `json:"total_units"`
            CodCharges float64  `json:"cod_charges"`
            CashbackApplied float64  `json:"cashback_applied"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            TransferPrice float64  `json:"transfer_price"`
            Cashback float64  `json:"cashback"`
            HsnCode string  `json:"hsn_code"`
            Identifiers Identifiers  `json:"identifiers"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            RefundCredit float64  `json:"refund_credit"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            FyndCredits float64  `json:"fynd_credits"`
            ValueOfGood float64  `json:"value_of_good"`
         
    }
    
    // ItemBrand ...
    type ItemBrand struct {

        
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
    }
    
    // Item ...
    type Item struct {

        
            Image []string  `json:"image"`
            ID float64  `json:"id"`
            Size string  `json:"size"`
            Name string  `json:"name"`
            Code string  `json:"code"`
            Brand ItemBrand  `json:"brand"`
            SellerIdentifier string  `json:"seller_identifier"`
            SlugKey string  `json:"slug_key"`
         
    }
    
    // Prices ...
    type Prices struct {

        
            CouponValue float64  `json:"coupon_value"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
            AmountPaid float64  `json:"amount_paid"`
            FyndCredits float64  `json:"fynd_credits"`
            Discount float64  `json:"discount"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            PriceMarked float64  `json:"price_marked"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            PriceEffective float64  `json:"price_effective"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CodCharges float64  `json:"cod_charges"`
            CashbackApplied float64  `json:"cashback_applied"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            TransferPrice float64  `json:"transfer_price"`
            Cashback float64  `json:"cashback"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            RefundCredit float64  `json:"refund_credit"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            RefundAmount float64  `json:"refund_amount"`
            ValueOfGood float64  `json:"value_of_good"`
         
    }
    
    // CurrentStatus ...
    type CurrentStatus struct {

        
            Name string  `json:"name"`
            Status string  `json:"status"`
            UpdatedAt string  `json:"updated_at"`
            JourneyType string  `json:"journey_type"`
         
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

        
            PromoID string  `json:"promo_id"`
            Amount float64  `json:"amount"`
            PromotionType string  `json:"promotion_type"`
            PromotionName string  `json:"promotion_name"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            MrpPromotion bool  `json:"mrp_promotion"`
            ArticleQuantity float64  `json:"article_quantity"`
         
    }
    
    // Bags ...
    type Bags struct {

        
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            CanCancel bool  `json:"can_cancel"`
            ID float64  `json:"id"`
            Item Item  `json:"item"`
            LineNumber float64  `json:"line_number"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            Prices Prices  `json:"prices"`
            CanReturn bool  `json:"can_return"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            ReturnableDate string  `json:"returnable_date"`
            DeliveryDate string  `json:"delivery_date"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            Quantity float64  `json:"quantity"`
            SellerIdentifier string  `json:"seller_identifier"`
         
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

        
            IsCurrent bool  `json:"is_current"`
            Time string  `json:"time"`
            Status string  `json:"status"`
            TrackingDetails []NestedTrackingDetails  `json:"tracking_details"`
            IsPassed bool  `json:"is_passed"`
         
    }
    
    // BreakupValues ...
    type BreakupValues struct {

        
            Display string  `json:"display"`
            Name string  `json:"name"`
            Value float64  `json:"value"`
         
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

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // ShipmentStatus ...
    type ShipmentStatus struct {

        
            HexCode string  `json:"hex_code"`
            Title string  `json:"title"`
         
    }
    
    // DeliveryAddress ...
    type DeliveryAddress struct {

        
            ContactPerson string  `json:"contact_person"`
            State string  `json:"state"`
            Email string  `json:"email"`
            City string  `json:"city"`
            CreatedAt string  `json:"created_at"`
            Latitude float64  `json:"latitude"`
            AddressCategory string  `json:"address_category"`
            Landmark string  `json:"landmark"`
            Name string  `json:"name"`
            Longitude float64  `json:"longitude"`
            Pincode string  `json:"pincode"`
            Address string  `json:"address"`
            AddressType string  `json:"address_type"`
            Address1 string  `json:"address1"`
            Country string  `json:"country"`
            Address2 string  `json:"address2"`
            Version string  `json:"version"`
            UpdatedAt string  `json:"updated_at"`
            Area string  `json:"area"`
            Phone string  `json:"phone"`
         
    }
    
    // Shipments ...
    type Shipments struct {

        
            TrackURL string  `json:"track_url"`
            DpName string  `json:"dp_name"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            UserInfo ShipmentUserInfo  `json:"user_info"`
            SizeInfo map[string]interface{}  `json:"size_info"`
            TotalDetails ShipmentTotalDetails  `json:"total_details"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            Payment ShipmentPayment  `json:"payment"`
            NeedHelpURL string  `json:"need_help_url"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            ReturnableDate string  `json:"returnable_date"`
            Invoice Invoice  `json:"invoice"`
            Bags []Bags  `json:"bags"`
            CanCancel bool  `json:"can_cancel"`
            ShipmentID string  `json:"shipment_id"`
            CanBreak map[string]interface{}  `json:"can_break"`
            ShowDownloadInvoice bool  `json:"show_download_invoice"`
            OrderType string  `json:"order_type"`
            CanReturn bool  `json:"can_return"`
            TrackingDetails []TrackingDetails  `json:"tracking_details"`
            OrderID string  `json:"order_id"`
            DeliveryDate string  `json:"delivery_date"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            Prices Prices  `json:"prices"`
            TotalBags float64  `json:"total_bags"`
            Promise Promise  `json:"promise"`
            BeneficiaryDetails bool  `json:"beneficiary_details"`
            TrakingNo string  `json:"traking_no"`
            Comment string  `json:"comment"`
            AwbNo string  `json:"awb_no"`
            ShowTrackLink bool  `json:"show_track_link"`
            RefundDetails map[string]interface{}  `json:"refund_details"`
            FulfillingCompany FulfillingCompany  `json:"fulfilling_company"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            DeliveryAddress DeliveryAddress  `json:"delivery_address"`
         
    }
    
    // BagsForReorderArticleAssignment ...
    type BagsForReorderArticleAssignment struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // BagsForReorder ...
    type BagsForReorder struct {

        
            ItemSize string  `json:"item_size"`
            ItemID float64  `json:"item_id"`
            StoreID float64  `json:"store_id"`
            SellerID float64  `json:"seller_id"`
            ArticleAssignment BagsForReorderArticleAssignment  `json:"article_assignment"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // UserInfo ...
    type UserInfo struct {

        
            Gender string  `json:"gender"`
            Name string  `json:"name"`
            Mobile string  `json:"mobile"`
            Email string  `json:"email"`
         
    }
    
    // OrderSchema ...
    type OrderSchema struct {

        
            Shipments []Shipments  `json:"shipments"`
            BagsForReorder []BagsForReorder  `json:"bags_for_reorder"`
            OrderID string  `json:"order_id"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
            OrderCreatedTime string  `json:"order_created_time"`
            UserInfo UserInfo  `json:"user_info"`
         
    }
    
    // OrderPage ...
    type OrderPage struct {

        
            Size float64  `json:"size"`
            Type string  `json:"type"`
            HasNext bool  `json:"has_next"`
            Current float64  `json:"current"`
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
            Success bool  `json:"success"`
            PresignedType string  `json:"presigned_type"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // Track ...
    type Track struct {

        
            UpdatedTime string  `json:"updated_time"`
            ShipmentType string  `json:"shipment_type"`
            Status string  `json:"status"`
            UpdatedAt string  `json:"updated_at"`
            Reason string  `json:"reason"`
            Awb string  `json:"awb"`
            AccountName string  `json:"account_name"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
         
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

        
            Message string  `json:"message"`
            Success bool  `json:"success"`
            ResendTimer float64  `json:"resend_timer"`
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

        
            QcType []string  `json:"qc_type"`
            DisplayName string  `json:"display_name"`
            QuestionSet []QuestionSet  `json:"question_set"`
            ID float64  `json:"id"`
            Reasons []BagReasons  `json:"reasons"`
            Meta BagReasonMeta  `json:"meta"`
         
    }
    
    // ShipmentBagReasons ...
    type ShipmentBagReasons struct {

        
            Success bool  `json:"success"`
            Reasons []BagReasons  `json:"reasons"`
         
    }
    
    // ShipmentReason ...
    type ShipmentReason struct {

        
            ReasonText string  `json:"reason_text"`
            ReasonID float64  `json:"reason_id"`
            Priority float64  `json:"priority"`
            ShowTextArea bool  `json:"show_text_area"`
            Flow string  `json:"flow"`
            FeedbackType string  `json:"feedback_type"`
         
    }
    
    // ShipmentReasons ...
    type ShipmentReasons struct {

        
            Reasons []ShipmentReason  `json:"reasons"`
         
    }
    
    // Products ...
    type Products struct {

        
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
         
    }
    
    // EntitiesDataUpdates ...
    type EntitiesDataUpdates struct {

        
            Filters []map[string]interface{}  `json:"filters"`
            Data map[string]interface{}  `json:"data"`
         
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
    
    // DataUpdates ...
    type DataUpdates struct {

        
            Entities []EntitiesDataUpdates  `json:"entities"`
            Products []ProductsDataUpdates  `json:"products"`
         
    }
    
    // EntityReasonData ...
    type EntityReasonData struct {

        
            ReasonID float64  `json:"reason_id"`
            ReasonText string  `json:"reason_text"`
         
    }
    
    // EntitiesReasons ...
    type EntitiesReasons struct {

        
            Filters []map[string]interface{}  `json:"filters"`
            Data EntityReasonData  `json:"data"`
         
    }
    
    // ProductsReasonsFilters ...
    type ProductsReasonsFilters struct {

        
            Quantity float64  `json:"quantity"`
            Identifier string  `json:"identifier"`
            LineNumber float64  `json:"line_number"`
         
    }
    
    // ProductsReasonsData ...
    type ProductsReasonsData struct {

        
            ReasonID float64  `json:"reason_id"`
            ReasonText string  `json:"reason_text"`
         
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
    
    // ShipmentsRequest ...
    type ShipmentsRequest struct {

        
            Products []Products  `json:"products"`
            Identifier string  `json:"identifier"`
            DataUpdates DataUpdates  `json:"data_updates"`
            Reasons ReasonsData  `json:"reasons"`
         
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
            Task bool  `json:"task"`
            Statuses []StatuesRequest  `json:"statuses"`
            LockAfterTransition bool  `json:"lock_after_transition"`
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
    
    // ErrorResponse ...
    type ErrorResponse struct {

        
            Message string  `json:"message"`
            Status float64  `json:"status"`
            Exception string  `json:"exception"`
            Code string  `json:"code"`
            StackTrace string  `json:"stack_trace"`
         
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
    

    
    // CouponBreakup ...
    type CouponBreakup struct {

        
            Code string  `json:"code"`
            Type string  `json:"type"`
            Message string  `json:"message"`
            UID string  `json:"uid"`
            Value float64  `json:"value"`
            IsApplied bool  `json:"is_applied"`
         
    }
    
    // DisplayBreakup ...
    type DisplayBreakup struct {

        
            CurrencyCode string  `json:"currency_code"`
            Message []string  `json:"message"`
            Key string  `json:"key"`
            CurrencySymbol string  `json:"currency_symbol"`
            Value float64  `json:"value"`
            Display string  `json:"display"`
         
    }
    
    // LoyaltyPoints ...
    type LoyaltyPoints struct {

        
            IsApplied bool  `json:"is_applied"`
            Applicable float64  `json:"applicable"`
            Total float64  `json:"total"`
            Description string  `json:"description"`
         
    }
    
    // RawBreakup ...
    type RawBreakup struct {

        
            Subtotal float64  `json:"subtotal"`
            GstCharges float64  `json:"gst_charges"`
            Coupon float64  `json:"coupon"`
            YouSaved float64  `json:"you_saved"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CodCharge float64  `json:"cod_charge"`
            Vog float64  `json:"vog"`
            FyndCash float64  `json:"fynd_cash"`
            MrpTotal float64  `json:"mrp_total"`
            Total float64  `json:"total"`
            ConvenienceFee float64  `json:"convenience_fee"`
            Discount float64  `json:"discount"`
         
    }
    
    // CartBreakup ...
    type CartBreakup struct {

        
            Coupon CouponBreakup  `json:"coupon"`
            Display []DisplayBreakup  `json:"display"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Raw RawBreakup  `json:"raw"`
         
    }
    
    // PaymentSelectionLock ...
    type PaymentSelectionLock struct {

        
            Enabled bool  `json:"enabled"`
            DefaultOptions string  `json:"default_options"`
            PaymentIdentifier string  `json:"payment_identifier"`
         
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
    
    // ProductAvailability ...
    type ProductAvailability struct {

        
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            OutOfStock bool  `json:"out_of_stock"`
            Sizes []string  `json:"sizes"`
            IsValid bool  `json:"is_valid"`
            Deliverable bool  `json:"deliverable"`
         
    }
    
    // ProductPrice ...
    type ProductPrice struct {

        
            Marked float64  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
            AddOn float64  `json:"add_on"`
            CurrencySymbol string  `json:"currency_symbol"`
            Selling float64  `json:"selling"`
            Effective float64  `json:"effective"`
         
    }
    
    // ProductPriceInfo ...
    type ProductPriceInfo struct {

        
            Base ProductPrice  `json:"base"`
            Converted ProductPrice  `json:"converted"`
         
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

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // CategoryInfo ...
    type CategoryInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductImage ...
    type ProductImage struct {

        
            SecureURL string  `json:"secure_url"`
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
         
    }
    
    // CartProduct ...
    type CartProduct struct {

        
            Slug string  `json:"slug"`
            Action ProductAction  `json:"action"`
            Brand BaseInfo  `json:"brand"`
            Name string  `json:"name"`
            Type string  `json:"type"`
            Categories []CategoryInfo  `json:"categories"`
            UID float64  `json:"uid"`
            Images []ProductImage  `json:"images"`
         
    }
    
    // CartProductIdentifer ...
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // Ownership ...
    type Ownership struct {

        
            PayableBy string  `json:"payable_by"`
            PayableCategory string  `json:"payable_category"`
         
    }
    
    // AppliedPromotion ...
    type AppliedPromotion struct {

        
            OfferText string  `json:"offer_text"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromotionType string  `json:"promotion_type"`
            Amount float64  `json:"amount"`
            PromoID string  `json:"promo_id"`
            Ownership Ownership  `json:"ownership"`
            MrpPromotion bool  `json:"mrp_promotion"`
         
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

        
            Type string  `json:"type"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Seller BaseInfo  `json:"seller"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Store BaseInfo  `json:"store"`
            Price ArticlePriceInfo  `json:"price"`
            UID string  `json:"uid"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Size string  `json:"size"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // PromoMeta ...
    type PromoMeta struct {

        
            Message string  `json:"message"`
         
    }
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            CouponMessage string  `json:"coupon_message"`
            IsSet bool  `json:"is_set"`
            Availability ProductAvailability  `json:"availability"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Message string  `json:"message"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Key string  `json:"key"`
            Product CartProduct  `json:"product"`
            Discount string  `json:"discount"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Price ProductPriceInfo  `json:"price"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Article ProductArticle  `json:"article"`
            PromoMeta PromoMeta  `json:"promo_meta"`
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
            Message string  `json:"message"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            ID string  `json:"id"`
            CouponText string  `json:"coupon_text"`
            Gstin string  `json:"gstin"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            BuyNow bool  `json:"buy_now"`
            IsValid bool  `json:"is_valid"`
            LastModified string  `json:"last_modified"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Items []CartProductInfo  `json:"items"`
            Currency CartCurrency  `json:"currency"`
            Comment string  `json:"comment"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            SellerID float64  `json:"seller_id"`
            Pos bool  `json:"pos"`
            StoreID float64  `json:"store_id"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ItemID float64  `json:"item_id"`
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            ProductGroupTags []string  `json:"product_group_tags"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            ArticleID string  `json:"article_id"`
            ItemSize string  `json:"item_size"`
            Display string  `json:"display"`
            Quantity float64  `json:"quantity"`
         
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

        
            ItemIndex float64  `json:"item_index"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ItemID float64  `json:"item_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ArticleID string  `json:"article_id"`
            ItemSize string  `json:"item_size"`
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
            Cart CartDetailResponse  `json:"cart"`
            Message string  `json:"message"`
         
    }
    
    // CartItemCountResponse ...
    type CartItemCountResponse struct {

        
            UserCartItemsCount float64  `json:"user_cart_items_count"`
         
    }
    
    // Coupon ...
    type Coupon struct {

        
            Title string  `json:"title"`
            CouponValue float64  `json:"coupon_value"`
            Message string  `json:"message"`
            ExpiresOn string  `json:"expires_on"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            IsApplicable bool  `json:"is_applicable"`
            SubTitle string  `json:"sub_title"`
            CouponCode string  `json:"coupon_code"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            IsApplied bool  `json:"is_applied"`
         
    }
    
    // PageCoupon ...
    type PageCoupon struct {

        
            HasNext bool  `json:"has_next"`
            HasPrevious bool  `json:"has_previous"`
            Total float64  `json:"total"`
            TotalItemCount float64  `json:"total_item_count"`
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

        
            BulkEffective float64  `json:"bulk_effective"`
            Marked float64  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
            Effective float64  `json:"effective"`
         
    }
    
    // OfferItem ...
    type OfferItem struct {

        
            Margin float64  `json:"margin"`
            Type string  `json:"type"`
            AutoApplied bool  `json:"auto_applied"`
            Price OfferPrice  `json:"price"`
            Total float64  `json:"total"`
            Best bool  `json:"best"`
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

        
            GeoLocation GeoLocation  `json:"geo_location"`
            AreaCode string  `json:"area_code"`
            CountryIsoCode string  `json:"country_iso_code"`
            Landmark string  `json:"landmark"`
            ID string  `json:"id"`
            UserID string  `json:"user_id"`
            Meta map[string]interface{}  `json:"meta"`
            AreaCodeSlug string  `json:"area_code_slug"`
            City string  `json:"city"`
            Email string  `json:"email"`
            IsDefaultAddress bool  `json:"is_default_address"`
            Phone string  `json:"phone"`
            Area string  `json:"area"`
            CountryPhoneCode string  `json:"country_phone_code"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            CheckoutMode string  `json:"checkout_mode"`
            AddressType string  `json:"address_type"`
            Tags []string  `json:"tags"`
            Address string  `json:"address"`
            State string  `json:"state"`
            Name string  `json:"name"`
            IsActive bool  `json:"is_active"`
            Country string  `json:"country"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // GetAddressesResponse ...
    type GetAddressesResponse struct {

        
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
            CartID string  `json:"cart_id"`
            BillingAddressID string  `json:"billing_address_id"`
         
    }
    
    // UpdateCartPaymentRequest ...
    type UpdateCartPaymentRequest struct {

        
            AggregatorName string  `json:"aggregator_name"`
            ID string  `json:"id"`
            AddressID string  `json:"address_id"`
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentMode string  `json:"payment_mode"`
            MerchantCode string  `json:"merchant_code"`
         
    }
    
    // CouponValidity ...
    type CouponValidity struct {

        
            Code string  `json:"code"`
            Title string  `json:"title"`
            Valid bool  `json:"valid"`
            DisplayMessageEn string  `json:"display_message_en"`
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

        
            DpOptions map[string]interface{}  `json:"dp_options"`
            Shipments float64  `json:"shipments"`
            OrderType string  `json:"order_type"`
            DpID string  `json:"dp_id"`
            FulfillmentType string  `json:"fulfillment_type"`
            Promise ShipmentPromise  `json:"promise"`
            BoxType string  `json:"box_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
            ShipmentType string  `json:"shipment_type"`
            Items []CartProductInfo  `json:"items"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            Message string  `json:"message"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            ID string  `json:"id"`
            CheckoutMode string  `json:"checkout_mode"`
            UID string  `json:"uid"`
            Currency CartCurrency  `json:"currency"`
            Comment string  `json:"comment"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Shipments []ShipmentResponse  `json:"shipments"`
            Error bool  `json:"error"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CouponText string  `json:"coupon_text"`
            Gstin string  `json:"gstin"`
            BuyNow bool  `json:"buy_now"`
            IsValid bool  `json:"is_valid"`
            CartID float64  `json:"cart_id"`
            LastModified string  `json:"last_modified"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
         
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

        
            User string  `json:"user"`
            LastName string  `json:"last_name"`
            EmployeeCode string  `json:"employee_code"`
            ID string  `json:"_id"`
            FirstName string  `json:"first_name"`
         
    }
    
    // CartCheckoutCustomMeta ...
    type CartCheckoutCustomMeta struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // Files ...
    type Files struct {

        
            Key string  `json:"key"`
            Values []string  `json:"values"`
         
    }
    
    // CartPosCheckoutDetailRequest ...
    type CartPosCheckoutDetailRequest struct {

        
            BillingAddressID string  `json:"billing_address_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            Meta map[string]interface{}  `json:"meta"`
            AddressID string  `json:"address_id"`
            Staff StaffCheckout  `json:"staff"`
            PaymentMode string  `json:"payment_mode"`
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            Pos bool  `json:"pos"`
            Aggregator string  `json:"aggregator"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            CustomMeta []CartCheckoutCustomMeta  `json:"custom_meta"`
            OrderType string  `json:"order_type"`
            OrderingStore float64  `json:"ordering_store"`
            Files []Files  `json:"files"`
            MerchantCode string  `json:"merchant_code"`
            CallbackURL string  `json:"callback_url"`
         
    }
    
    // CheckCart ...
    type CheckCart struct {

        
            DeliveryCharges float64  `json:"delivery_charges"`
            Message string  `json:"message"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            ID string  `json:"id"`
            ErrorMessage string  `json:"error_message"`
            Items []CartProductInfo  `json:"items"`
            Comment string  `json:"comment"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Gstin string  `json:"gstin"`
            BuyNow bool  `json:"buy_now"`
            IsValid bool  `json:"is_valid"`
            LastModified string  `json:"last_modified"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            OrderID string  `json:"order_id"`
            CheckoutMode string  `json:"checkout_mode"`
            UID string  `json:"uid"`
            CodMessage string  `json:"cod_message"`
            StoreCode string  `json:"store_code"`
            Currency CartCurrency  `json:"currency"`
            UserType string  `json:"user_type"`
            CodAvailable bool  `json:"cod_available"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            Success bool  `json:"success"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            CodCharges float64  `json:"cod_charges"`
            CouponText string  `json:"coupon_text"`
            CartID float64  `json:"cart_id"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
         
    }
    
    // CartCheckoutResponse ...
    type CartCheckoutResponse struct {

        
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            OrderID string  `json:"order_id"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
            Cart CheckCart  `json:"cart"`
            AppInterceptURL string  `json:"app_intercept_url"`
            CallbackURL string  `json:"callback_url"`
            Data map[string]interface{}  `json:"data"`
         
    }
    
    // CartMetaRequest ...
    type CartMetaRequest struct {

        
            CheckoutMode string  `json:"checkout_mode"`
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            Gstin string  `json:"gstin"`
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
    
    // CartDeliveryModesResponse ...
    type CartDeliveryModesResponse struct {

        
            PickupStores []float64  `json:"pickup_stores"`
            AvailableModes []string  `json:"available_modes"`
         
    }
    
    // PickupStoreDetail ...
    type PickupStoreDetail struct {

        
            Area string  `json:"area"`
            Pincode float64  `json:"pincode"`
            AreaCode string  `json:"area_code"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Name string  `json:"name"`
            City string  `json:"city"`
            Email string  `json:"email"`
            Landmark string  `json:"landmark"`
            ID float64  `json:"id"`
            StoreCode string  `json:"store_code"`
            Phone string  `json:"phone"`
            AddressType string  `json:"address_type"`
            Country string  `json:"country"`
            UID float64  `json:"uid"`
            State string  `json:"state"`
            Address string  `json:"address"`
         
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

        
            User map[string]interface{}  `json:"user"`
            Source map[string]interface{}  `json:"source"`
            CreatedOn string  `json:"created_on"`
            Meta map[string]interface{}  `json:"meta"`
            Token string  `json:"token"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            Message string  `json:"message"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            ID string  `json:"id"`
            Items []CartProductInfo  `json:"items"`
            Comment string  `json:"comment"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            Gstin string  `json:"gstin"`
            BuyNow bool  `json:"buy_now"`
            IsValid bool  `json:"is_valid"`
            LastModified string  `json:"last_modified"`
            CheckoutMode string  `json:"checkout_mode"`
            UID string  `json:"uid"`
            Currency CartCurrency  `json:"currency"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
            CouponText string  `json:"coupon_text"`
            CartID float64  `json:"cart_id"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Error string  `json:"error"`
            Cart SharedCart  `json:"cart"`
         
    }
    

    
    // CountryMetaResponse ...
    type CountryMetaResponse struct {

        
            IsdCode string  `json:"isd_code"`
            CountryCode string  `json:"country_code"`
         
    }
    
    // PincodeErrorSchemaResponse ...
    type PincodeErrorSchemaResponse struct {

        
            Type string  `json:"type"`
            Message string  `json:"message"`
            Value string  `json:"value"`
         
    }
    
    // PincodeMetaResponse ...
    type PincodeMetaResponse struct {

        
            Zone string  `json:"zone"`
            InternalZoneID float64  `json:"internal_zone_id"`
         
    }
    
    // PincodeParentsResponse ...
    type PincodeParentsResponse struct {

        
            SubType string  `json:"sub_type"`
            DisplayName string  `json:"display_name"`
            UID string  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // PincodeDataResponse ...
    type PincodeDataResponse struct {

        
            MetaCode CountryMetaResponse  `json:"meta_code"`
            SubType string  `json:"sub_type"`
            Error PincodeErrorSchemaResponse  `json:"error"`
            DisplayName string  `json:"display_name"`
            UID string  `json:"uid"`
            Meta PincodeMetaResponse  `json:"meta"`
            Parents []PincodeParentsResponse  `json:"parents"`
            Name string  `json:"name"`
         
    }
    
    // PincodeApiResponse ...
    type PincodeApiResponse struct {

        
            Data []PincodeDataResponse  `json:"data"`
            Error PincodeErrorSchemaResponse  `json:"error"`
            Success bool  `json:"success"`
         
    }
    
    // TATCategoryRequest ...
    type TATCategoryRequest struct {

        
            Level string  `json:"level"`
            ID float64  `json:"id"`
         
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
            Articles []TATArticlesRequest  `json:"articles"`
            FromPincode string  `json:"from_pincode"`
         
    }
    
    // TATViewRequest ...
    type TATViewRequest struct {

        
            Journey string  `json:"journey"`
            ToPincode string  `json:"to_pincode"`
            Action string  `json:"action"`
            Identifier string  `json:"identifier"`
            Source string  `json:"source"`
            LocationDetails []TATLocationDetailsRequest  `json:"location_details"`
         
    }
    
    // TATErrorSchemaResponse ...
    type TATErrorSchemaResponse struct {

        
            Type string  `json:"type"`
            Message string  `json:"message"`
            Value string  `json:"value"`
         
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

        
            Category TATCategoryRequest  `json:"category"`
            Error TATErrorSchemaResponse  `json:"error"`
            Promise TATPromiseResponse  `json:"promise"`
            ManufacturingTime float64  `json:"manufacturing_time"`
            ManufacturingTimeUnit string  `json:"manufacturing_time_unit"`
            ManufacturingTimeSeconds float64  `json:"_manufacturing_time_seconds"`
            IsCodAvailable bool  `json:"is_cod_available"`
         
    }
    
    // TATLocationDetailsResponse ...
    type TATLocationDetailsResponse struct {

        
            FulfillmentID float64  `json:"fulfillment_id"`
            Articles []TATArticlesResponse  `json:"articles"`
            FromPincode string  `json:"from_pincode"`
         
    }
    
    // TATViewResponse ...
    type TATViewResponse struct {

        
            Journey string  `json:"journey"`
            ToPincode string  `json:"to_pincode"`
            PaymentMode string  `json:"payment_mode"`
            Success bool  `json:"success"`
            StormbreakerUUID string  `json:"stormbreaker_uuid"`
            Error TATErrorSchemaResponse  `json:"error"`
            Action string  `json:"action"`
            Identifier string  `json:"identifier"`
            ToCity string  `json:"to_city"`
            Source string  `json:"source"`
            RequestUUID string  `json:"request_uuid"`
            LocationDetails []TATLocationDetailsResponse  `json:"location_details"`
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
            SubType string  `json:"sub_type"`
            DisplayName string  `json:"display_name"`
            UID string  `json:"uid"`
            Logistics LogisticsResponse  `json:"logistics"`
            Type string  `json:"type"`
            Meta CountryMetaResponse  `json:"meta"`
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

        
            Zones []string  `json:"zones"`
            ServiceabilityType string  `json:"serviceability_type"`
         
    }
    
