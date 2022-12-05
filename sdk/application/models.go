package application



    
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
    
    // Meta ...
    type Meta struct {

        
            Source string  `json:"source"`
         
    }
    
    // Media ...
    type Media struct {

        
            Meta Meta  `json:"meta"`
            Type string  `json:"type"`
            URL string  `json:"url"`
            Alt string  `json:"alt"`
         
    }
    
    // ProductListingActionPage ...
    type ProductListingActionPage struct {

        
            Type string  `json:"type"`
            Query map[string]interface{}  `json:"query"`
            Params map[string]interface{}  `json:"params"`
         
    }
    
    // ProductListingAction ...
    type ProductListingAction struct {

        
            Page ProductListingActionPage  `json:"page"`
            Type string  `json:"type"`
         
    }
    
    // ProductBrand ...
    type ProductBrand struct {

        
            Logo Media  `json:"logo"`
            UID float64  `json:"uid"`
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            Description string  `json:"description"`
         
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

        
            Effective Price  `json:"effective"`
            Marked Price  `json:"marked"`
         
    }
    
    // CustomMetaFields ...
    type CustomMetaFields struct {

        
            Value string  `json:"value"`
            Key string  `json:"key"`
         
    }
    
    // ProductDetail ...
    type ProductDetail struct {

        
            Color string  `json:"color"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            TeaserTag string  `json:"teaser_tag"`
            Description string  `json:"description"`
            ItemCode string  `json:"item_code"`
            ImageNature string  `json:"image_nature"`
            Highlights []string  `json:"highlights"`
            ProductOnlineDate string  `json:"product_online_date"`
            Categories []ProductBrand  `json:"categories"`
            Similars []string  `json:"similars"`
            RatingCount float64  `json:"rating_count"`
            Medias []Media  `json:"medias"`
            Rating float64  `json:"rating"`
            Discount string  `json:"discount"`
            Type string  `json:"type"`
            ShortDescription string  `json:"short_description"`
            Tryouts []string  `json:"tryouts"`
            Price ProductListingPrice  `json:"price"`
            ItemType string  `json:"item_type"`
            Brand ProductBrand  `json:"brand"`
            HasVariant bool  `json:"has_variant"`
            Attributes map[string]interface{}  `json:"attributes"`
            Tags []string  `json:"tags"`
            UID float64  `json:"uid"`
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
         
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
            Col4 ColumnHeader  `json:"col_4"`
            Col1 ColumnHeader  `json:"col_1"`
            Col5 ColumnHeader  `json:"col_5"`
            Col3 ColumnHeader  `json:"col_3"`
            Col6 ColumnHeader  `json:"col_6"`
         
    }
    
    // SizeChartValues ...
    type SizeChartValues struct {

        
            Col2 string  `json:"col_2"`
            Col4 string  `json:"col_4"`
            Col1 string  `json:"col_1"`
            Col5 string  `json:"col_5"`
            Col3 string  `json:"col_3"`
            Col6 string  `json:"col_6"`
         
    }
    
    // SizeChart ...
    type SizeChart struct {

        
            Unit string  `json:"unit"`
            Image string  `json:"image"`
            Title string  `json:"title"`
            Headers ColumnHeaders  `json:"headers"`
            Description string  `json:"description"`
            Sizes []SizeChartValues  `json:"sizes"`
            SizeTip string  `json:"size_tip"`
         
    }
    
    // ProductSize ...
    type ProductSize struct {

        
            SellerIdentifiers []string  `json:"seller_identifiers"`
            Quantity float64  `json:"quantity"`
            Value string  `json:"value"`
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
            SizeChart SizeChart  `json:"size_chart"`
            Sizes []ProductSize  `json:"sizes"`
            Price ProductListingPrice  `json:"price"`
            Stores ProductSizeStores  `json:"stores"`
            Discount string  `json:"discount"`
         
    }
    
    // AttributeDetail ...
    type AttributeDetail struct {

        
            Description string  `json:"description"`
            Display string  `json:"display"`
            Key string  `json:"key"`
            Logo string  `json:"logo"`
         
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

        
            ColorName string  `json:"color_name"`
            Color string  `json:"color"`
            Medias []Media  `json:"medias"`
            UID float64  `json:"uid"`
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
            Value string  `json:"value"`
            IsAvailable bool  `json:"is_available"`
         
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
    
    // CompanyDetail ...
    type CompanyDetail struct {

        
            ID float64  `json:"id"`
            Name string  `json:"name"`
         
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

        
            Count float64  `json:"count"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductStockPrice ...
    type ProductStockPrice struct {

        
            Effective float64  `json:"effective"`
            Currency string  `json:"currency"`
            Marked float64  `json:"marked"`
         
    }
    
    // ProductStockStatusItem ...
    type ProductStockStatusItem struct {

        
            Company CompanyDetail  `json:"company"`
            Store StoreDetail  `json:"store"`
            Seller Seller  `json:"seller"`
            Identifier map[string]interface{}  `json:"identifier"`
            UID string  `json:"uid"`
            Size string  `json:"size"`
            Quantity float64  `json:"quantity"`
            Price ProductStockPrice  `json:"price"`
            ItemID float64  `json:"item_id"`
         
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
    
    // ProductSortOn ...
    type ProductSortOn struct {

        
            Value string  `json:"value"`
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // ProductVariantListingResponse ...
    type ProductVariantListingResponse struct {

        
            Header string  `json:"header"`
            DisplayType string  `json:"display_type"`
            Total float64  `json:"total"`
            Items []ProductVariantItemResponse  `json:"items"`
            Key string  `json:"key"`
         
    }
    
    // ProductListingDetail ...
    type ProductListingDetail struct {

        
            Color string  `json:"color"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            TeaserTag string  `json:"teaser_tag"`
            Description string  `json:"description"`
            Sizes []string  `json:"sizes"`
            ItemCode string  `json:"item_code"`
            ImageNature string  `json:"image_nature"`
            Highlights []string  `json:"highlights"`
            ProductOnlineDate string  `json:"product_online_date"`
            Categories []ProductBrand  `json:"categories"`
            Similars []string  `json:"similars"`
            RatingCount float64  `json:"rating_count"`
            Medias []Media  `json:"medias"`
            Rating float64  `json:"rating"`
            Discount string  `json:"discount"`
            Type string  `json:"type"`
            ShortDescription string  `json:"short_description"`
            Variants []ProductVariantListingResponse  `json:"variants"`
            Sellable bool  `json:"sellable"`
            Tryouts []string  `json:"tryouts"`
            Price ProductListingPrice  `json:"price"`
            ItemType string  `json:"item_type"`
            Brand ProductBrand  `json:"brand"`
            HasVariant bool  `json:"has_variant"`
            Attributes map[string]interface{}  `json:"attributes"`
            Tags []string  `json:"tags"`
            UID float64  `json:"uid"`
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            CustomMeta []CustomMetaFields  `json:"_custom_meta"`
         
    }
    
    // ProductFiltersKey ...
    type ProductFiltersKey struct {

        
            Name string  `json:"name"`
            Logo string  `json:"logo"`
            Kind string  `json:"kind"`
            Display string  `json:"display"`
         
    }
    
    // ProductFiltersValue ...
    type ProductFiltersValue struct {

        
            Min float64  `json:"min"`
            IsSelected bool  `json:"is_selected"`
            Count float64  `json:"count"`
            DisplayFormat string  `json:"display_format"`
            SelectedMin float64  `json:"selected_min"`
            QueryFormat string  `json:"query_format"`
            CurrencySymbol string  `json:"currency_symbol"`
            SelectedMax float64  `json:"selected_max"`
            CurrencyCode string  `json:"currency_code"`
            Value string  `json:"value"`
            Max float64  `json:"max"`
            Display string  `json:"display"`
         
    }
    
    // ProductFilters ...
    type ProductFilters struct {

        
            Key ProductFiltersKey  `json:"key"`
            Values []ProductFiltersValue  `json:"values"`
         
    }
    
    // ProductListingResponse ...
    type ProductListingResponse struct {

        
            Page Page  `json:"page"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Items []ProductListingDetail  `json:"items"`
            Filters []ProductFilters  `json:"filters"`
         
    }
    
    // ImageUrls ...
    type ImageUrls struct {

        
            Portrait Media  `json:"portrait"`
            Landscape Media  `json:"landscape"`
         
    }
    
    // BrandItem ...
    type BrandItem struct {

        
            Logo Media  `json:"logo"`
            Action ProductListingAction  `json:"action"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            Slug string  `json:"slug"`
            Departments []string  `json:"departments"`
            Banners ImageUrls  `json:"banners"`
            Discount string  `json:"discount"`
         
    }
    
    // BrandListingResponse ...
    type BrandListingResponse struct {

        
            Page Page  `json:"page"`
            Items []BrandItem  `json:"items"`
         
    }
    
    // BrandDetailResponse ...
    type BrandDetailResponse struct {

        
            Logo Media  `json:"logo"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // ThirdLevelChild ...
    type ThirdLevelChild struct {

        
            Childs []map[string]interface{}  `json:"childs"`
            UID float64  `json:"uid"`
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // SecondLevelChild ...
    type SecondLevelChild struct {

        
            Childs []ThirdLevelChild  `json:"childs"`
            UID float64  `json:"uid"`
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // Child ...
    type Child struct {

        
            Childs []SecondLevelChild  `json:"childs"`
            UID float64  `json:"uid"`
            Action ProductListingAction  `json:"action"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // CategoryItems ...
    type CategoryItems struct {

        
            Childs []Child  `json:"childs"`
            Action ProductListingAction  `json:"action"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // DepartmentCategoryTree ...
    type DepartmentCategoryTree struct {

        
            Items []CategoryItems  `json:"items"`
            Department string  `json:"department"`
         
    }
    
    // DepartmentIdentifier ...
    type DepartmentIdentifier struct {

        
            UID float64  `json:"uid"`
            Slug string  `json:"slug"`
         
    }
    
    // CategoryListingResponse ...
    type CategoryListingResponse struct {

        
            Data []DepartmentCategoryTree  `json:"data"`
            Departments []DepartmentIdentifier  `json:"departments"`
         
    }
    
    // CategoryMetaResponse ...
    type CategoryMetaResponse struct {

        
            Logo Media  `json:"logo"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Banners ImageUrls  `json:"banners"`
         
    }
    
    // HomeListingResponse ...
    type HomeListingResponse struct {

        
            Page Page  `json:"page"`
            Items []ProductListingDetail  `json:"items"`
            Message string  `json:"message"`
         
    }
    
    // Department ...
    type Department struct {

        
            Logo Media  `json:"logo"`
            UID float64  `json:"uid"`
            PriorityOrder float64  `json:"priority_order"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
         
    }
    
    // DepartmentResponse ...
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // AutocompleteItem ...
    type AutocompleteItem struct {

        
            Type string  `json:"type"`
            Logo Media  `json:"logo"`
            Action ProductListingAction  `json:"action"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Display string  `json:"display"`
         
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

        
            AllowFacets bool  `json:"allow_facets"`
            Logo Media  `json:"logo"`
            Description string  `json:"description"`
            Priority float64  `json:"priority"`
            Query []CollectionQuery  `json:"query"`
            AllowSort bool  `json:"allow_sort"`
            Tag []string  `json:"tag"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            Banners ImageUrls  `json:"banners"`
            Type string  `json:"type"`
            SortOn string  `json:"sort_on"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Action ProductListingAction  `json:"action"`
            UID string  `json:"uid"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Badge map[string]interface{}  `json:"badge"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Cron map[string]interface{}  `json:"cron"`
            Meta map[string]interface{}  `json:"meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
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
    
    // GetCollectionListingResponse ...
    type GetCollectionListingResponse struct {

        
            Page Page  `json:"page"`
            Items []GetCollectionDetailNest  `json:"items"`
            Filters CollectionListingFilter  `json:"filters"`
         
    }
    
    // CollectionDetailResponse ...
    type CollectionDetailResponse struct {

        
            AllowFacets bool  `json:"allow_facets"`
            Logo Media  `json:"logo"`
            Description string  `json:"description"`
            Priority float64  `json:"priority"`
            Query []CollectionQuery  `json:"query"`
            AllowSort bool  `json:"allow_sort"`
            Tag []string  `json:"tag"`
            IsActive bool  `json:"is_active"`
            AppID string  `json:"app_id"`
            Banners ImageUrls  `json:"banners"`
            Type string  `json:"type"`
            SortOn string  `json:"sort_on"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Badge map[string]interface{}  `json:"badge"`
            Schedule map[string]interface{}  `json:"_schedule"`
            Cron map[string]interface{}  `json:"cron"`
            Meta map[string]interface{}  `json:"meta"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
         
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

        
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
            Address string  `json:"address"`
            LatLong LatLong  `json:"lat_long"`
            StoreCode string  `json:"store_code"`
            Country string  `json:"country"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            StoreEmail string  `json:"store_email"`
            State string  `json:"state"`
         
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
            Name string  `json:"name"`
            CompanyType string  `json:"company_type"`
         
    }
    
    // StoreAddressSerializer ...
    type StoreAddressSerializer struct {

        
            City string  `json:"city"`
            Pincode float64  `json:"pincode"`
            Latitude float64  `json:"latitude"`
            Country string  `json:"country"`
            Landmark string  `json:"landmark"`
            Longitude float64  `json:"longitude"`
            State string  `json:"state"`
            Address1 string  `json:"address1"`
            Address2 string  `json:"address2"`
         
    }
    
    // SellerPhoneNumber ...
    type SellerPhoneNumber struct {

        
            CountryCode float64  `json:"country_code"`
            Number string  `json:"number"`
         
    }
    
    // StoreDepartments ...
    type StoreDepartments struct {

        
            Logo string  `json:"logo"`
            UID float64  `json:"uid"`
            PriorityOrder float64  `json:"priority_order"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
         
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
            Address StoreAddressSerializer  `json:"address"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Departments []StoreDepartments  `json:"departments"`
            Manager StoreManagerSerializer  `json:"manager"`
         
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

        
            Open bool  `json:"open"`
            Closing Time  `json:"closing"`
            Opening Time  `json:"opening"`
            Weekday string  `json:"weekday"`
         
    }
    
    // StoreDetails ...
    type StoreDetails struct {

        
            Company CompanyStore  `json:"company"`
            Address StoreAddressSerializer  `json:"address"`
            ContactNumbers []SellerPhoneNumber  `json:"contact_numbers"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
            Timing []StoreTiming  `json:"timing"`
            CustomJson map[string]interface{}  `json:"_custom_json"`
            Departments []StoreDepartments  `json:"departments"`
            Manager StoreManagerSerializer  `json:"manager"`
         
    }
    
    // UserDetail ...
    type UserDetail struct {

        
            Contact string  `json:"contact"`
            UserID string  `json:"user_id"`
            Username string  `json:"username"`
            SuperUser bool  `json:"super_user"`
         
    }
    
    // ProductDetails ...
    type ProductDetails struct {

        
            Media []map[string]interface{}  `json:"media"`
            GroupedAttributes map[string]interface{}  `json:"grouped_attributes"`
            Identifier map[string]interface{}  `json:"identifier"`
            Description interface{}  `json:"description"`
            Images []interface{}  `json:"images"`
            ItemCode interface{}  `json:"item_code"`
            ImageNature interface{}  `json:"image_nature"`
            Highlights []interface{}  `json:"highlights"`
            HsnCode float64  `json:"hsn_code"`
            OutOfStock bool  `json:"out_of_stock"`
            RatingCount float64  `json:"rating_count"`
            BrandUID float64  `json:"brand_uid"`
            Rating float64  `json:"rating"`
            ShortDescription interface{}  `json:"short_description"`
            HasVariant bool  `json:"has_variant"`
            TemplateTag interface{}  `json:"template_tag"`
            Attributes map[string]interface{}  `json:"attributes"`
            Name interface{}  `json:"name"`
            Slug interface{}  `json:"slug"`
            IsSet bool  `json:"is_set"`
            CountryOfOrigin interface{}  `json:"country_of_origin"`
         
    }
    
    // Size ...
    type Size struct {

        
            Value interface{}  `json:"value"`
            Display interface{}  `json:"display"`
            IsAvailable bool  `json:"is_available"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // ProductGroupPrice ...
    type ProductGroupPrice struct {

        
            MinEffective float64  `json:"min_effective"`
            Currency interface{}  `json:"currency"`
            MaxMarked float64  `json:"max_marked"`
            MaxEffective float64  `json:"max_effective"`
            MinMarked float64  `json:"min_marked"`
         
    }
    
    // ProductInGroup ...
    type ProductInGroup struct {

        
            ProductDetails ProductDetails  `json:"product_details"`
            AllowRemove bool  `json:"allow_remove"`
            AutoAddToCart bool  `json:"auto_add_to_cart"`
            MaxQuantity float64  `json:"max_quantity"`
            Sizes []Size  `json:"sizes"`
            ProductUID float64  `json:"product_uid"`
            AutoSelect bool  `json:"auto_select"`
            MinQuantity float64  `json:"min_quantity"`
            Price ProductGroupPrice  `json:"price"`
         
    }
    
    // ProductGroupingModel ...
    type ProductGroupingModel struct {

        
            ModifiedOn string  `json:"modified_on"`
            ID interface{}  `json:"_id"`
            SameStoreAssignment bool  `json:"same_store_assignment"`
            VerifiedOn string  `json:"verified_on"`
            CompanyID float64  `json:"company_id"`
            Logo string  `json:"logo"`
            ModifiedBy UserDetail  `json:"modified_by"`
            IsActive bool  `json:"is_active"`
            VerifiedBy UserDetail  `json:"verified_by"`
            Name interface{}  `json:"name"`
            Slug interface{}  `json:"slug"`
            Choice interface{}  `json:"choice"`
            Meta map[string]interface{}  `json:"meta"`
            CreatedOn string  `json:"created_on"`
            CreatedBy UserDetail  `json:"created_by"`
            Products []ProductInGroup  `json:"products"`
            PageVisibility []interface{}  `json:"page_visibility"`
         
    }
    
    // ProductBundle ...
    type ProductBundle struct {

        
            Items []ProductGroupingModel  `json:"items"`
         
    }
    
    // SellerV2 ...
    type SellerV2 struct {

        
            Count float64  `json:"count"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ArticleAssignmentV2 ...
    type ArticleAssignmentV2 struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // DetailsSchemaV2 ...
    type DetailsSchemaV2 struct {

        
            Value string  `json:"value"`
            Type string  `json:"type"`
            Key string  `json:"key"`
         
    }
    
    // SellerGroupAttributes ...
    type SellerGroupAttributes struct {

        
            Title string  `json:"title"`
            Details []DetailsSchemaV2  `json:"details"`
         
    }
    
    // StrategyWiseListingSchemaV2 ...
    type StrategyWiseListingSchemaV2 struct {

        
            Pincode float64  `json:"pincode"`
            Tat float64  `json:"tat"`
            Distance float64  `json:"distance"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // MarketPlaceSttributesSchemaV2 ...
    type MarketPlaceSttributesSchemaV2 struct {

        
            Title string  `json:"title"`
            Details []DetailsSchemaV2  `json:"details"`
         
    }
    
    // ProductStockPriceV2 ...
    type ProductStockPriceV2 struct {

        
            Effective float64  `json:"effective"`
            Currency string  `json:"currency"`
            Marked float64  `json:"marked"`
         
    }
    
    // ProductStockUnitPriceV2 ...
    type ProductStockUnitPriceV2 struct {

        
            Unit string  `json:"unit"`
            Price float64  `json:"price"`
            CurrencyCode string  `json:"currency_code"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // StoreV2 ...
    type StoreV2 struct {

        
            Count float64  `json:"count"`
            UID float64  `json:"uid"`
            Name string  `json:"name"`
         
    }
    
    // ProductSetDistributionSizeV2 ...
    type ProductSetDistributionSizeV2 struct {

        
            Size string  `json:"size"`
            Pieces float64  `json:"pieces"`
         
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
    
    // ReturnConfigSchemaV2 ...
    type ReturnConfigSchemaV2 struct {

        
            Unit string  `json:"unit"`
            Time float64  `json:"time"`
            Returnable bool  `json:"returnable"`
         
    }
    
    // ProductSizePriceResponseV2 ...
    type ProductSizePriceResponseV2 struct {

        
            LongLat []float64  `json:"long_lat"`
            Pincode float64  `json:"pincode"`
            Seller SellerV2  `json:"seller"`
            ArticleAssignment ArticleAssignmentV2  `json:"article_assignment"`
            GroupedAttributes []SellerGroupAttributes  `json:"grouped_attributes"`
            StrategyWiseListing []StrategyWiseListingSchemaV2  `json:"strategy_wise_listing"`
            SpecialBadge string  `json:"special_badge"`
            MarketplaceAttributes []MarketPlaceSttributesSchemaV2  `json:"marketplace_attributes"`
            PricePerPiece ProductStockPriceV2  `json:"price_per_piece"`
            PricePerUnit ProductStockUnitPriceV2  `json:"price_per_unit"`
            Discount string  `json:"discount"`
            Store StoreV2  `json:"store"`
            SellerCount float64  `json:"seller_count"`
            Quantity float64  `json:"quantity"`
            Price ProductStockPriceV2  `json:"price"`
            Set ProductSetV2  `json:"set"`
            ItemType string  `json:"item_type"`
            ArticleID string  `json:"article_id"`
            ReturnConfig ReturnConfigSchemaV2  `json:"return_config"`
         
    }
    
    // ProductSizeSellerFilterSchemaV2 ...
    type ProductSizeSellerFilterSchemaV2 struct {

        
            Value string  `json:"value"`
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // ProductSizeSellersResponseV2 ...
    type ProductSizeSellersResponseV2 struct {

        
            Page Page  `json:"page"`
            SortOn []ProductSizeSellerFilterSchemaV2  `json:"sort_on"`
            Items []ProductSizePriceResponseV2  `json:"items"`
         
    }
    

    
    // CartCurrency ...
    type CartCurrency struct {

        
            Symbol string  `json:"symbol"`
            Code string  `json:"code"`
         
    }
    
    // RawBreakup ...
    type RawBreakup struct {

        
            GstCharges float64  `json:"gst_charges"`
            YouSaved float64  `json:"you_saved"`
            CodCharge float64  `json:"cod_charge"`
            Subtotal float64  `json:"subtotal"`
            Total float64  `json:"total"`
            Coupon float64  `json:"coupon"`
            Discount float64  `json:"discount"`
            ConvenienceFee float64  `json:"convenience_fee"`
            DeliveryCharge float64  `json:"delivery_charge"`
            MrpTotal float64  `json:"mrp_total"`
            FyndCash float64  `json:"fynd_cash"`
            Vog float64  `json:"vog"`
         
    }
    
    // LoyaltyPoints ...
    type LoyaltyPoints struct {

        
            Description string  `json:"description"`
            IsApplied bool  `json:"is_applied"`
            Total float64  `json:"total"`
            Applicable float64  `json:"applicable"`
         
    }
    
    // DisplayBreakup ...
    type DisplayBreakup struct {

        
            CurrencyCode string  `json:"currency_code"`
            Message []string  `json:"message"`
            Display string  `json:"display"`
            Value float64  `json:"value"`
            Key string  `json:"key"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // CouponBreakup ...
    type CouponBreakup struct {

        
            SubTitle string  `json:"sub_title"`
            CouponValue float64  `json:"coupon_value"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Message string  `json:"message"`
            Value float64  `json:"value"`
            CouponType string  `json:"coupon_type"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Description string  `json:"description"`
            Title string  `json:"title"`
            Type string  `json:"type"`
            IsApplied bool  `json:"is_applied"`
            Code string  `json:"code"`
            UID string  `json:"uid"`
         
    }
    
    // CartBreakup ...
    type CartBreakup struct {

        
            Raw RawBreakup  `json:"raw"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
            Display []DisplayBreakup  `json:"display"`
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
    
    // ProductImage ...
    type ProductImage struct {

        
            AspectRatio string  `json:"aspect_ratio"`
            URL string  `json:"url"`
            SecureURL string  `json:"secure_url"`
         
    }
    
    // CategoryInfo ...
    type CategoryInfo struct {

        
            Name string  `json:"name"`
            UID float64  `json:"uid"`
         
    }
    
    // CartProduct ...
    type CartProduct struct {

        
            Name string  `json:"name"`
            Action ProductAction  `json:"action"`
            Brand BaseInfo  `json:"brand"`
            Slug string  `json:"slug"`
            Type string  `json:"type"`
            Images []ProductImage  `json:"images"`
            Categories []CategoryInfo  `json:"categories"`
            UID float64  `json:"uid"`
         
    }
    
    // ProductPrice ...
    type ProductPrice struct {

        
            CurrencyCode string  `json:"currency_code"`
            Selling float64  `json:"selling"`
            AddOn float64  `json:"add_on"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // ProductPriceInfo ...
    type ProductPriceInfo struct {

        
            Base ProductPrice  `json:"base"`
            Converted ProductPrice  `json:"converted"`
         
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
            Offer map[string]interface{}  `json:"offer"`
            MatchedBuyRules []string  `json:"matched_buy_rules"`
         
    }
    
    // FreeGiftItem ...
    type FreeGiftItem struct {

        
            ItemImagesURL []string  `json:"item_images_url"`
            ItemName string  `json:"item_name"`
            ItemSlug string  `json:"item_slug"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemID float64  `json:"item_id"`
         
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

        
            BuyRules []BuyRules  `json:"buy_rules"`
            DiscountRules []DiscountRulesApp  `json:"discount_rules"`
            PromoID string  `json:"promo_id"`
            ArticleQuantity float64  `json:"article_quantity"`
            Amount float64  `json:"amount"`
            PromotionName string  `json:"promotion_name"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            MrpPromotion bool  `json:"mrp_promotion"`
            PromotionType string  `json:"promotion_type"`
            OfferText string  `json:"offer_text"`
         
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

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            Price ArticlePriceInfo  `json:"price"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Quantity float64  `json:"quantity"`
            Store BaseInfo  `json:"store"`
            Size string  `json:"size"`
            Type string  `json:"type"`
            Seller BaseInfo  `json:"seller"`
            UID string  `json:"uid"`
         
    }
    
    // ProductAvailability ...
    type ProductAvailability struct {

        
            IsValid bool  `json:"is_valid"`
            Sizes []string  `json:"sizes"`
            OtherStoreQuantity float64  `json:"other_store_quantity"`
            Deliverable bool  `json:"deliverable"`
            OutOfStock bool  `json:"out_of_stock"`
         
    }
    
    // CartProductIdentifer ...
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            PromoMeta PromoMeta  `json:"promo_meta"`
            CouponMessage string  `json:"coupon_message"`
            Message string  `json:"message"`
            Product CartProduct  `json:"product"`
            Price ProductPriceInfo  `json:"price"`
            Discount string  `json:"discount"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            IsSet bool  `json:"is_set"`
            Quantity float64  `json:"quantity"`
            PromotionsApplied []AppliedPromotion  `json:"promotions_applied"`
            Article ProductArticle  `json:"article"`
            BulkOffer map[string]interface{}  `json:"bulk_offer"`
            PricePerUnit ProductPriceInfo  `json:"price_per_unit"`
            Availability ProductAvailability  `json:"availability"`
            Key string  `json:"key"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
         
    }
    
    // PaymentSelectionLock ...
    type PaymentSelectionLock struct {

        
            PaymentIdentifier string  `json:"payment_identifier"`
            DefaultOptions string  `json:"default_options"`
            Enabled bool  `json:"enabled"`
         
    }
    
    // CartDetailResponse ...
    type CartDetailResponse struct {

        
            IsValid bool  `json:"is_valid"`
            Currency CartCurrency  `json:"currency"`
            CheckoutMode string  `json:"checkout_mode"`
            CouponText string  `json:"coupon_text"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Gstin string  `json:"gstin"`
            Message string  `json:"message"`
            Comment string  `json:"comment"`
            LastModified string  `json:"last_modified"`
            ID string  `json:"id"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Items []CartProductInfo  `json:"items"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            ArticleAssignment map[string]interface{}  `json:"article_assignment"`
            StoreID float64  `json:"store_id"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ArticleID string  `json:"article_id"`
            ItemSize string  `json:"item_size"`
            Display string  `json:"display"`
            Pos bool  `json:"pos"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            ProductGroupTags []string  `json:"product_group_tags"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
            SellerID float64  `json:"seller_id"`
         
    }
    
    // AddCartRequest ...
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
         
    }
    
    // AddCartDetailResponse ...
    type AddCartDetailResponse struct {

        
            Partial bool  `json:"partial"`
            Cart CartDetailResponse  `json:"cart"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // UpdateProductCart ...
    type UpdateProductCart struct {

        
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            ArticleID string  `json:"article_id"`
            ItemSize string  `json:"item_size"`
            ItemIndex float64  `json:"item_index"`
            ParentItemIdentifiers map[string]interface{}  `json:"parent_item_identifiers"`
            Quantity float64  `json:"quantity"`
            ItemID float64  `json:"item_id"`
            Identifiers CartProductIdentifer  `json:"identifiers"`
         
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
    
    // CartItemCountResponse ...
    type CartItemCountResponse struct {

        
            UserCartItemsCount float64  `json:"user_cart_items_count"`
         
    }
    
    // PageCoupon ...
    type PageCoupon struct {

        
            HasPrevious bool  `json:"has_previous"`
            TotalItemCount float64  `json:"total_item_count"`
            HasNext bool  `json:"has_next"`
            Total float64  `json:"total"`
            Current float64  `json:"current"`
         
    }
    
    // Coupon ...
    type Coupon struct {

        
            SubTitle string  `json:"sub_title"`
            CouponValue float64  `json:"coupon_value"`
            MaxDiscountValue float64  `json:"max_discount_value"`
            Message string  `json:"message"`
            ExpiresOn string  `json:"expires_on"`
            CouponType string  `json:"coupon_type"`
            MinimumCartValue float64  `json:"minimum_cart_value"`
            Description string  `json:"description"`
            CouponCode string  `json:"coupon_code"`
            Title string  `json:"title"`
            IsApplicable bool  `json:"is_applicable"`
            IsApplied bool  `json:"is_applied"`
         
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

        
            CurrencyCode string  `json:"currency_code"`
            BulkEffective float64  `json:"bulk_effective"`
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // OfferItem ...
    type OfferItem struct {

        
            Margin float64  `json:"margin"`
            Price OfferPrice  `json:"price"`
            Type string  `json:"type"`
            Quantity float64  `json:"quantity"`
            Best bool  `json:"best"`
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

        
            Latitude float64  `json:"latitude"`
            Longitude float64  `json:"longitude"`
         
    }
    
    // Address ...
    type Address struct {

        
            AreaCode string  `json:"area_code"`
            CheckoutMode string  `json:"checkout_mode"`
            IsActive bool  `json:"is_active"`
            Address string  `json:"address"`
            Meta map[string]interface{}  `json:"meta"`
            State string  `json:"state"`
            CountryCode string  `json:"country_code"`
            IsDefaultAddress bool  `json:"is_default_address"`
            Phone string  `json:"phone"`
            GoogleMapPoint map[string]interface{}  `json:"google_map_point"`
            Country string  `json:"country"`
            Email string  `json:"email"`
            ID string  `json:"id"`
            AddressType string  `json:"address_type"`
            Name string  `json:"name"`
            GeoLocation GeoLocation  `json:"geo_location"`
            Area string  `json:"area"`
            AreaCodeSlug string  `json:"area_code_slug"`
            City string  `json:"city"`
            UserID string  `json:"user_id"`
            Tags []string  `json:"tags"`
            Landmark string  `json:"landmark"`
         
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
            Success bool  `json:"success"`
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

        
            CartID string  `json:"cart_id"`
            BillingAddressID string  `json:"billing_address_id"`
            ID string  `json:"id"`
         
    }
    
    // UpdateCartPaymentRequest ...
    type UpdateCartPaymentRequest struct {

        
            MerchantCode string  `json:"merchant_code"`
            AggregatorName string  `json:"aggregator_name"`
            PaymentIdentifier string  `json:"payment_identifier"`
            ID string  `json:"id"`
            PaymentMode string  `json:"payment_mode"`
            AddressID string  `json:"address_id"`
         
    }
    
    // CouponValidity ...
    type CouponValidity struct {

        
            Valid bool  `json:"valid"`
            Discount float64  `json:"discount"`
            Title string  `json:"title"`
            DisplayMessageEn string  `json:"display_message_en"`
            Code string  `json:"code"`
         
    }
    
    // PaymentCouponValidate ...
    type PaymentCouponValidate struct {

        
            CouponValidity CouponValidity  `json:"coupon_validity"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // ShipmentResponse ...
    type ShipmentResponse struct {

        
            Shipments float64  `json:"shipments"`
            ShipmentType string  `json:"shipment_type"`
            DpOptions map[string]interface{}  `json:"dp_options"`
            OrderType string  `json:"order_type"`
            FulfillmentType string  `json:"fulfillment_type"`
            FulfillmentID float64  `json:"fulfillment_id"`
            Promise ShipmentPromise  `json:"promise"`
            Items []CartProductInfo  `json:"items"`
            BoxType string  `json:"box_type"`
            DpID string  `json:"dp_id"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            Shipments []ShipmentResponse  `json:"shipments"`
            CartID float64  `json:"cart_id"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            CouponText string  `json:"coupon_text"`
            Message string  `json:"message"`
            Error bool  `json:"error"`
            Comment string  `json:"comment"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            IsValid bool  `json:"is_valid"`
            Currency CartCurrency  `json:"currency"`
            Gstin string  `json:"gstin"`
            LastModified string  `json:"last_modified"`
            ID string  `json:"id"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            UID string  `json:"uid"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
         
    }
    
    // CustomerDetails ...
    type CustomerDetails struct {

        
            Email string  `json:"email"`
            Name string  `json:"name"`
            Mobile string  `json:"mobile"`
         
    }
    
    // StaffCheckout ...
    type StaffCheckout struct {

        
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            User string  `json:"user"`
            ID string  `json:"_id"`
         
    }
    
    // CartCheckoutDetailRequest ...
    type CartCheckoutDetailRequest struct {

        
            MerchantCode string  `json:"merchant_code"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            CustomerDetails CustomerDetails  `json:"customer_details"`
            PaymentIdentifier string  `json:"payment_identifier"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
            OrderingStore float64  `json:"ordering_store"`
            BillingAddressID string  `json:"billing_address_id"`
            CallbackURL string  `json:"callback_url"`
            Meta map[string]interface{}  `json:"meta"`
            Aggregator string  `json:"aggregator"`
            PaymentMode string  `json:"payment_mode"`
            AddressID string  `json:"address_id"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            Staff StaffCheckout  `json:"staff"`
         
    }
    
    // CheckCart ...
    type CheckCart struct {

        
            CartID float64  `json:"cart_id"`
            CheckoutMode string  `json:"checkout_mode"`
            ErrorMessage string  `json:"error_message"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Comment string  `json:"comment"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CodCharges float64  `json:"cod_charges"`
            IsValid bool  `json:"is_valid"`
            UserType string  `json:"user_type"`
            Success bool  `json:"success"`
            ID string  `json:"id"`
            CouponText string  `json:"coupon_text"`
            DeliveryChargeOrderValue float64  `json:"delivery_charge_order_value"`
            Message string  `json:"message"`
            StoreCode string  `json:"store_code"`
            CodMessage string  `json:"cod_message"`
            Items []CartProductInfo  `json:"items"`
            Currency CartCurrency  `json:"currency"`
            OrderID string  `json:"order_id"`
            StoreEmps []map[string]interface{}  `json:"store_emps"`
            Gstin string  `json:"gstin"`
            LastModified string  `json:"last_modified"`
            DeliveryCharges float64  `json:"delivery_charges"`
            CodAvailable bool  `json:"cod_available"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            UID string  `json:"uid"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
         
    }
    
    // CartCheckoutResponse ...
    type CartCheckoutResponse struct {

        
            Data map[string]interface{}  `json:"data"`
            OrderID string  `json:"order_id"`
            AppInterceptURL string  `json:"app_intercept_url"`
            Message string  `json:"message"`
            PaymentConfirmURL string  `json:"payment_confirm_url"`
            CallbackURL string  `json:"callback_url"`
            Success bool  `json:"success"`
            Cart CheckCart  `json:"cart"`
         
    }
    
    // CartMetaRequest ...
    type CartMetaRequest struct {

        
            PickUpCustomerDetails map[string]interface{}  `json:"pick_up_customer_details"`
            Gstin string  `json:"gstin"`
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

        
            Token string  `json:"token"`
            Meta map[string]interface{}  `json:"meta"`
            User map[string]interface{}  `json:"user"`
            Source map[string]interface{}  `json:"source"`
            CreatedOn string  `json:"created_on"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            IsValid bool  `json:"is_valid"`
            CartID float64  `json:"cart_id"`
            CheckoutMode string  `json:"checkout_mode"`
            Currency CartCurrency  `json:"currency"`
            CouponText string  `json:"coupon_text"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Gstin string  `json:"gstin"`
            Message string  `json:"message"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            Comment string  `json:"comment"`
            LastModified string  `json:"last_modified"`
            ID string  `json:"id"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Items []CartProductInfo  `json:"items"`
            UID string  `json:"uid"`
            PaymentSelectionLock PaymentSelectionLock  `json:"payment_selection_lock"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Cart SharedCart  `json:"cart"`
            Error string  `json:"error"`
         
    }
    
    // FreeGiftItems ...
    type FreeGiftItems struct {

        
            ItemImagesURL []string  `json:"item_images_url"`
            ItemName string  `json:"item_name"`
            ItemSlug string  `json:"item_slug"`
            ItemBrandName string  `json:"item_brand_name"`
            ItemPriceDetails map[string]interface{}  `json:"item_price_details"`
            ItemID float64  `json:"item_id"`
         
    }
    
    // PromotionOffer ...
    type PromotionOffer struct {

        
            BuyRules map[string]interface{}  `json:"buy_rules"`
            DiscountRules []map[string]interface{}  `json:"discount_rules"`
            Description string  `json:"description"`
            ID string  `json:"id"`
            PromotionGroup string  `json:"promotion_group"`
            ValidTill string  `json:"valid_till"`
            OfferText string  `json:"offer_text"`
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
            Effective float64  `json:"effective"`
            Marked float64  `json:"marked"`
            CurrencySymbol string  `json:"currency_symbol"`
         
    }
    
    // LadderOfferItem ...
    type LadderOfferItem struct {

        
            Margin float64  `json:"margin"`
            MaxQuantity float64  `json:"max_quantity"`
            Price LadderPrice  `json:"price"`
            MinQuantity float64  `json:"min_quantity"`
            Type string  `json:"type"`
         
    }
    
    // LadderPriceOffer ...
    type LadderPriceOffer struct {

        
            BuyRules map[string]interface{}  `json:"buy_rules"`
            OfferPrices []LadderOfferItem  `json:"offer_prices"`
            DiscountRules []map[string]interface{}  `json:"discount_rules"`
            Description string  `json:"description"`
            ID string  `json:"id"`
            PromotionGroup string  `json:"promotion_group"`
            CalculateOn string  `json:"calculate_on"`
            ValidTill string  `json:"valid_till"`
            OfferText string  `json:"offer_text"`
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
            IsDeleted bool  `json:"is_deleted"`
         
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
            Redirections []RedirectionSchema  `json:"redirections"`
            ID string  `json:"_id"`
            UpdatedAt string  `json:"updated_at"`
            CreatedAt string  `json:"created_at"`
         
    }
    
    // RedirectionSchema ...
    type RedirectionSchema struct {

        
            RedirectFrom string  `json:"redirect_from"`
            RedirectTo string  `json:"redirect_to"`
         
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

        
            UserID string  `json:"user_id"`
            MerchantKey string  `json:"merchant_key"`
            Key string  `json:"key"`
            VerifyAPI string  `json:"verify_api"`
            Sdk bool  `json:"sdk"`
            MerchantID string  `json:"merchant_id"`
            API string  `json:"api"`
            ConfigType string  `json:"config_type"`
            Pin string  `json:"pin"`
            Secret string  `json:"secret"`
         
    }
    
    // AggregatorsConfigDetailResponse ...
    type AggregatorsConfigDetailResponse struct {

        
            Simpl AggregatorConfigDetail  `json:"simpl"`
            Payumoney AggregatorConfigDetail  `json:"payumoney"`
            Razorpay AggregatorConfigDetail  `json:"razorpay"`
            Juspay AggregatorConfigDetail  `json:"juspay"`
            Success bool  `json:"success"`
            Mswipe AggregatorConfigDetail  `json:"mswipe"`
            Stripe AggregatorConfigDetail  `json:"stripe"`
            Env string  `json:"env"`
            Ccavenue AggregatorConfigDetail  `json:"ccavenue"`
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

        
            Nickname string  `json:"nickname"`
            CardID string  `json:"card_id"`
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

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            Cards CardPaymentGateway  `json:"cards"`
         
    }
    
    // Card ...
    type Card struct {

        
            Nickname string  `json:"nickname"`
            CardID string  `json:"card_id"`
            CardToken string  `json:"card_token"`
            CardType string  `json:"card_type"`
            CardBrand string  `json:"card_brand"`
            ExpYear float64  `json:"exp_year"`
            CardNumber string  `json:"card_number"`
            CardFingerprint string  `json:"card_fingerprint"`
            ExpMonth float64  `json:"exp_month"`
            Expired bool  `json:"expired"`
            CardIssuer string  `json:"card_issuer"`
            CardBrandImage string  `json:"card_brand_image"`
            AggregatorName string  `json:"aggregator_name"`
            CardReference string  `json:"card_reference"`
            CardIsin string  `json:"card_isin"`
            CardName string  `json:"card_name"`
         
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

        
            PhoneNumber string  `json:"phone_number"`
            Aggregator string  `json:"aggregator"`
            Payload string  `json:"payload"`
            MerchantParams map[string]interface{}  `json:"merchant_params"`
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

        
            Verified bool  `json:"verified"`
            OrderID string  `json:"order_id"`
            Aggregator string  `json:"aggregator"`
            TransactionToken string  `json:"transaction_token"`
            Amount float64  `json:"amount"`
         
    }
    
    // ChargeCustomerResponse ...
    type ChargeCustomerResponse struct {

        
            DeliveryAddressID string  `json:"delivery_address_id"`
            Status string  `json:"status"`
            OrderID string  `json:"order_id"`
            Aggregator string  `json:"aggregator"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
            CartID string  `json:"cart_id"`
         
    }
    
    // PaymentInitializationRequest ...
    type PaymentInitializationRequest struct {

        
            Email string  `json:"email"`
            Vpa string  `json:"vpa"`
            OrderID string  `json:"order_id"`
            Currency string  `json:"currency"`
            Aggregator string  `json:"aggregator"`
            Contact string  `json:"contact"`
            CustomerID string  `json:"customer_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Method string  `json:"method"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Timeout float64  `json:"timeout"`
            Amount float64  `json:"amount"`
         
    }
    
    // PaymentInitializationResponse ...
    type PaymentInitializationResponse struct {

        
            BqrImage string  `json:"bqr_image"`
            Vpa string  `json:"vpa"`
            Status string  `json:"status"`
            Currency string  `json:"currency"`
            Aggregator string  `json:"aggregator"`
            Success bool  `json:"success"`
            PollingURL string  `json:"polling_url"`
            CustomerID string  `json:"customer_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            UpiPollURL string  `json:"upi_poll_url"`
            Method string  `json:"method"`
            VirtualID string  `json:"virtual_id"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Timeout float64  `json:"timeout"`
            Amount float64  `json:"amount"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
         
    }
    
    // PaymentStatusUpdateRequest ...
    type PaymentStatusUpdateRequest struct {

        
            Email string  `json:"email"`
            Vpa string  `json:"vpa"`
            Status string  `json:"status"`
            OrderID string  `json:"order_id"`
            Currency string  `json:"currency"`
            Aggregator string  `json:"aggregator"`
            Contact string  `json:"contact"`
            CustomerID string  `json:"customer_id"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Method string  `json:"method"`
            Amount float64  `json:"amount"`
         
    }
    
    // PaymentStatusUpdateResponse ...
    type PaymentStatusUpdateResponse struct {

        
            Status string  `json:"status"`
            Retry bool  `json:"retry"`
            AggregatorName string  `json:"aggregator_name"`
         
    }
    
    // AggregatorRoute ...
    type AggregatorRoute struct {

        
            Data map[string]interface{}  `json:"data"`
            APILink string  `json:"api_link"`
            PaymentFlow string  `json:"payment_flow"`
            PaymentFlowData map[string]interface{}  `json:"payment_flow_data"`
         
    }
    
    // PaymentFlow ...
    type PaymentFlow struct {

        
            Simpl AggregatorRoute  `json:"simpl"`
            BqrRazorpay AggregatorRoute  `json:"bqr_razorpay"`
            Razorpay AggregatorRoute  `json:"razorpay"`
            Epaylater AggregatorRoute  `json:"epaylater"`
            Juspay AggregatorRoute  `json:"juspay"`
            Fynd AggregatorRoute  `json:"fynd"`
            UpiRazorpay AggregatorRoute  `json:"upi_razorpay"`
            Mswipe AggregatorRoute  `json:"mswipe"`
            Stripe AggregatorRoute  `json:"stripe"`
            Payubiz AggregatorRoute  `json:"payubiz"`
            Jiopay AggregatorRoute  `json:"jiopay"`
            Ccavenue AggregatorRoute  `json:"ccavenue"`
            Rupifi AggregatorRoute  `json:"rupifi"`
         
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

        
            PackageName string  `json:"package_name"`
            Code string  `json:"code"`
         
    }
    
    // PaymentModeList ...
    type PaymentModeList struct {

        
            CardType string  `json:"card_type"`
            DisplayPriority float64  `json:"display_priority"`
            RetryCount float64  `json:"retry_count"`
            ExpYear float64  `json:"exp_year"`
            CardFingerprint string  `json:"card_fingerprint"`
            Expired bool  `json:"expired"`
            AggregatorName string  `json:"aggregator_name"`
            CardID string  `json:"card_id"`
            CardToken string  `json:"card_token"`
            CardBrand string  `json:"card_brand"`
            Code string  `json:"code"`
            IntentApp []IntentApp  `json:"intent_app"`
            ExpMonth float64  `json:"exp_month"`
            CardIssuer string  `json:"card_issuer"`
            Timeout float64  `json:"timeout"`
            CardIsin string  `json:"card_isin"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            Nickname string  `json:"nickname"`
            IntentFlow bool  `json:"intent_flow"`
            IntentAppErrorDictList []IntentAppErrorList  `json:"intent_app_error_dict_list"`
            Name string  `json:"name"`
            CardNumber string  `json:"card_number"`
            CardBrandImage string  `json:"card_brand_image"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            CardReference string  `json:"card_reference"`
            FyndVpa string  `json:"fynd_vpa"`
            DisplayName string  `json:"display_name"`
            CardName string  `json:"card_name"`
            MerchantCode string  `json:"merchant_code"`
         
    }
    
    // RootPaymentMode ...
    type RootPaymentMode struct {

        
            List []PaymentModeList  `json:"list"`
            DisplayPriority float64  `json:"display_priority"`
            Name string  `json:"name"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            AggregatorName string  `json:"aggregator_name"`
            DisplayName string  `json:"display_name"`
            Logo string  `json:"logo"`
            AddCardEnabled bool  `json:"add_card_enabled"`
         
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

        
            Status string  `json:"status"`
            Message string  `json:"message"`
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
    
    // TransferItemsDetails ...
    type TransferItemsDetails struct {

        
            LogoLarge string  `json:"logo_large"`
            ID float64  `json:"id"`
            Name string  `json:"name"`
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
            BranchName string  `json:"branch_name"`
            DelightsUserName string  `json:"delights_user_name"`
            AccountNo string  `json:"account_no"`
            BankName string  `json:"bank_name"`
            ModifiedOn string  `json:"modified_on"`
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
            IsActive bool  `json:"is_active"`
            Address string  `json:"address"`
            BeneficiaryID string  `json:"beneficiary_id"`
            Comment string  `json:"comment"`
            CreatedOn string  `json:"created_on"`
            AccountHolder string  `json:"account_holder"`
            Email string  `json:"email"`
            IfscCode string  `json:"ifsc_code"`
            Mobile string  `json:"mobile"`
            DisplayName string  `json:"display_name"`
            TransferMode string  `json:"transfer_mode"`
         
    }
    
    // OrderBeneficiaryResponse ...
    type OrderBeneficiaryResponse struct {

        
            Beneficiaries []OrderBeneficiaryDetails  `json:"beneficiaries"`
            ShowBeneficiaryDetails bool  `json:"show_beneficiary_details"`
         
    }
    
    // NotFoundResourceError ...
    type NotFoundResourceError struct {

        
            Code string  `json:"code"`
            Success bool  `json:"success"`
            Description string  `json:"description"`
         
    }
    
    // IfscCodeResponse ...
    type IfscCodeResponse struct {

        
            BranchName string  `json:"branch_name"`
            BankName string  `json:"bank_name"`
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

        
            Wallet string  `json:"wallet"`
            Email string  `json:"email"`
            Vpa string  `json:"vpa"`
            Address string  `json:"address"`
            IfscCode string  `json:"ifsc_code"`
            BranchName string  `json:"branch_name"`
            AccountNo string  `json:"account_no"`
            Comment string  `json:"comment"`
            BankName string  `json:"bank_name"`
            AccountHolder string  `json:"account_holder"`
            Mobile string  `json:"mobile"`
         
    }
    
    // AddBeneficiaryDetailsRequest ...
    type AddBeneficiaryDetailsRequest struct {

        
            OrderID string  `json:"order_id"`
            Delights bool  `json:"delights"`
            RequestID string  `json:"request_id"`
            Details BeneficiaryModeDetails  `json:"details"`
            TransferMode string  `json:"transfer_mode"`
            Otp string  `json:"otp"`
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
            BranchName string  `json:"branch_name"`
            AccountNo string  `json:"account_no"`
            BankName string  `json:"bank_name"`
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

        
            Success bool  `json:"success"`
            IsVerifiedFlag string  `json:"is_verified_flag"`
            RequestID string  `json:"request_id"`
         
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
    
    // BalanceDetails ...
    type BalanceDetails struct {

        
            Currency string  `json:"currency"`
            Value float64  `json:"value"`
            FormattedValue string  `json:"formatted_value"`
         
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
    
    // KYCAddress ...
    type KYCAddress struct {

        
            LandMark string  `json:"land_mark"`
            City string  `json:"city"`
            Addressline1 string  `json:"addressline1"`
            Addressline2 string  `json:"addressline2"`
            OwnershipType string  `json:"ownership_type"`
            State string  `json:"state"`
            Pincode string  `json:"pincode"`
         
    }
    
    // UserPersonalInfoInDetails ...
    type UserPersonalInfoInDetails struct {

        
            Phone string  `json:"phone"`
            Pan string  `json:"pan"`
            Email string  `json:"email"`
            MothersName string  `json:"mothers_name"`
            MiddleName string  `json:"middle_name"`
            VoterID string  `json:"voter_id"`
            MobileVerified bool  `json:"mobile_verified"`
            FirstName string  `json:"first_name"`
            AddressAsPerID KYCAddress  `json:"address_as_per_id"`
            EmailVerified bool  `json:"email_verified"`
            Gender string  `json:"gender"`
            DrivingLicense string  `json:"driving_license"`
            FathersName string  `json:"fathers_name"`
            Passport string  `json:"passport"`
            Dob string  `json:"dob"`
            LastName string  `json:"last_name"`
         
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
            OsVersion string  `json:"os_version"`
            IdentificationNumber string  `json:"identification_number"`
            DeviceModel string  `json:"device_model"`
            DeviceType string  `json:"device_type"`
            Os string  `json:"os"`
            DeviceMake string  `json:"device_make"`
         
    }
    
    // BusinessDetails ...
    type BusinessDetails struct {

        
            BusinessType string  `json:"business_type"`
            Pan string  `json:"pan"`
            BusinessOwnershipType string  `json:"business_ownership_type"`
            Address KYCAddress  `json:"address"`
            Name string  `json:"name"`
            ShopAndEstablishment map[string]interface{}  `json:"shop_and_establishment"`
            Vintage string  `json:"vintage"`
            Fssai string  `json:"fssai"`
            Fda string  `json:"fda"`
            Gstin string  `json:"gstin"`
            EntityType string  `json:"entity_type"`
         
    }
    
    // CustomerOnboardingRequest ...
    type CustomerOnboardingRequest struct {

        
            PersonalInfo UserPersonalInfoInDetails  `json:"personal_info"`
            Mcc string  `json:"mcc"`
            Aggregator string  `json:"aggregator"`
            MarketplaceInfo MarketplaceInfo  `json:"marketplace_info"`
            Device DeviceDetails  `json:"device"`
            BusinessInfo BusinessDetails  `json:"business_info"`
            Source string  `json:"source"`
         
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
    
    // ShipmentPayment ...
    type ShipmentPayment struct {

        
            Mop string  `json:"mop"`
            Mode string  `json:"mode"`
            Logo string  `json:"logo"`
            Status string  `json:"status"`
         
    }
    
    // ShipmentStatus ...
    type ShipmentStatus struct {

        
            Title string  `json:"title"`
            HexCode string  `json:"hex_code"`
         
    }
    
    // FulfillingStore ...
    type FulfillingStore struct {

        
            Name string  `json:"name"`
            Code string  `json:"code"`
            CompanyName string  `json:"company_name"`
            CompanyID float64  `json:"company_id"`
            ID float64  `json:"id"`
         
    }
    
    // ShipmentTotalDetails ...
    type ShipmentTotalDetails struct {

        
            TotalPrice float64  `json:"total_price"`
            Sizes float64  `json:"sizes"`
            Pieces float64  `json:"pieces"`
         
    }
    
    // Prices ...
    type Prices struct {

        
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            RefundAmount float64  `json:"refund_amount"`
            TransferPrice float64  `json:"transfer_price"`
            CodCharges float64  `json:"cod_charges"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            ValueOfGood float64  `json:"value_of_good"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            CouponValue float64  `json:"coupon_value"`
            FyndCredits float64  `json:"fynd_credits"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            AmountPaid float64  `json:"amount_paid"`
            PriceEffective float64  `json:"price_effective"`
            RefundCredit float64  `json:"refund_credit"`
            Discount float64  `json:"discount"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CashbackApplied float64  `json:"cashback_applied"`
            Cashback float64  `json:"cashback"`
            PriceMarked float64  `json:"price_marked"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
         
    }
    
    // Invoice ...
    type Invoice struct {

        
            UpdatedDate string  `json:"updated_date"`
            InvoiceURL string  `json:"invoice_url"`
            LabelURL string  `json:"label_url"`
         
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
    
    // NestedTrackingDetails ...
    type NestedTrackingDetails struct {

        
            Time string  `json:"time"`
            Status string  `json:"status"`
            IsPassed bool  `json:"is_passed"`
            IsCurrent bool  `json:"is_current"`
         
    }
    
    // TrackingDetails ...
    type TrackingDetails struct {

        
            Status string  `json:"status"`
            IsPassed bool  `json:"is_passed"`
            Time string  `json:"time"`
            TrackingDetails []NestedTrackingDetails  `json:"tracking_details"`
            IsCurrent bool  `json:"is_current"`
         
    }
    
    // AppliedPromos ...
    type AppliedPromos struct {

        
            PromoID string  `json:"promo_id"`
            PromotionName string  `json:"promotion_name"`
            Amount float64  `json:"amount"`
            AppliedFreeArticles []AppliedFreeArticles  `json:"applied_free_articles"`
            ArticleQuantity float64  `json:"article_quantity"`
            PromotionType string  `json:"promotion_type"`
            MrpPromotion bool  `json:"mrp_promotion"`
         
    }
    
    // ItemBrand ...
    type ItemBrand struct {

        
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
    }
    
    // Item ...
    type Item struct {

        
            Name string  `json:"name"`
            Size string  `json:"size"`
            Code string  `json:"code"`
            Image []string  `json:"image"`
            SellerIdentifier string  `json:"seller_identifier"`
            Brand ItemBrand  `json:"brand"`
            SlugKey string  `json:"slug_key"`
            ID float64  `json:"id"`
         
    }
    
    // Identifiers ...
    type Identifiers struct {

        
            Ean string  `json:"ean"`
            SkuCode string  `json:"sku_code"`
         
    }
    
    // FinancialBreakup ...
    type FinancialBreakup struct {

        
            Identifiers Identifiers  `json:"identifiers"`
            GstTag string  `json:"gst_tag"`
            AmountPaidRoundoff float64  `json:"amount_paid_roundoff"`
            RefundAmount float64  `json:"refund_amount"`
            GstFee float64  `json:"gst_fee"`
            TransferPrice float64  `json:"transfer_price"`
            CodCharges float64  `json:"cod_charges"`
            CouponEffectiveDiscount float64  `json:"coupon_effective_discount"`
            ValueOfGood float64  `json:"value_of_good"`
            HsnCode string  `json:"hsn_code"`
            CouponValue float64  `json:"coupon_value"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            FyndCredits float64  `json:"fynd_credits"`
            ItemName string  `json:"item_name"`
            AmountPaid float64  `json:"amount_paid"`
            PromotionEffectiveDiscount float64  `json:"promotion_effective_discount"`
            PriceEffective float64  `json:"price_effective"`
            RefundCredit float64  `json:"refund_credit"`
            TotalUnits float64  `json:"total_units"`
            Discount float64  `json:"discount"`
            GstTaxPercentage float64  `json:"gst_tax_percentage"`
            DeliveryCharge float64  `json:"delivery_charge"`
            CashbackApplied float64  `json:"cashback_applied"`
            Size string  `json:"size"`
            Cashback float64  `json:"cashback"`
            PriceMarked float64  `json:"price_marked"`
            BrandCalculatedAmount float64  `json:"brand_calculated_amount"`
         
    }
    
    // CurrentStatus ...
    type CurrentStatus struct {

        
            UpdatedAt string  `json:"updated_at"`
            Name string  `json:"name"`
            Status string  `json:"status"`
            JourneyType string  `json:"journey_type"`
         
    }
    
    // Bags ...
    type Bags struct {

        
            Prices Prices  `json:"prices"`
            AppliedPromos []AppliedPromos  `json:"applied_promos"`
            LineNumber float64  `json:"line_number"`
            DeliveryDate string  `json:"delivery_date"`
            ReturnableDate string  `json:"returnable_date"`
            Item Item  `json:"item"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
            CanReturn bool  `json:"can_return"`
            CanCancel bool  `json:"can_cancel"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            SellerIdentifier string  `json:"seller_identifier"`
            ParentPromoBags map[string]interface{}  `json:"parent_promo_bags"`
            ID float64  `json:"id"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // FulfillingCompany ...
    type FulfillingCompany struct {

        
            Name string  `json:"name"`
            ID float64  `json:"id"`
         
    }
    
    // DeliveryAddress ...
    type DeliveryAddress struct {

        
            Email string  `json:"email"`
            UpdatedAt string  `json:"updated_at"`
            Address2 string  `json:"address2"`
            CreatedAt string  `json:"created_at"`
            City string  `json:"city"`
            AddressCategory string  `json:"address_category"`
            AddressType string  `json:"address_type"`
            Area string  `json:"area"`
            Landmark string  `json:"landmark"`
            Address1 string  `json:"address1"`
            State string  `json:"state"`
            Address string  `json:"address"`
            ContactPerson string  `json:"contact_person"`
            Phone string  `json:"phone"`
            Longitude float64  `json:"longitude"`
            Version string  `json:"version"`
            Name string  `json:"name"`
            Pincode string  `json:"pincode"`
            Country string  `json:"country"`
            Latitude float64  `json:"latitude"`
         
    }
    
    // ShipmentUserInfo ...
    type ShipmentUserInfo struct {

        
            Mobile string  `json:"mobile"`
            Gender string  `json:"gender"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
         
    }
    
    // Shipments ...
    type Shipments struct {

        
            OrderType string  `json:"order_type"`
            Payment ShipmentPayment  `json:"payment"`
            ReturnableDate string  `json:"returnable_date"`
            CustomMeta []map[string]interface{}  `json:"custom_meta"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            TotalBags float64  `json:"total_bags"`
            ShowTrackLink bool  `json:"show_track_link"`
            ShowDownloadInvoice bool  `json:"show_download_invoice"`
            TotalDetails ShipmentTotalDetails  `json:"total_details"`
            TrakingNo string  `json:"traking_no"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            RefundDetails map[string]interface{}  `json:"refund_details"`
            Prices Prices  `json:"prices"`
            Invoice Invoice  `json:"invoice"`
            Promise Promise  `json:"promise"`
            NeedHelpURL string  `json:"need_help_url"`
            ShipmentID string  `json:"shipment_id"`
            BeneficiaryDetails bool  `json:"beneficiary_details"`
            CanBreak map[string]interface{}  `json:"can_break"`
            TrackingDetails []TrackingDetails  `json:"tracking_details"`
            Bags []Bags  `json:"bags"`
            TrackURL string  `json:"track_url"`
            AwbNo string  `json:"awb_no"`
            OrderID string  `json:"order_id"`
            FulfillingCompany FulfillingCompany  `json:"fulfilling_company"`
            DeliveryDate string  `json:"delivery_date"`
            CanReturn bool  `json:"can_return"`
            CanCancel bool  `json:"can_cancel"`
            DpName string  `json:"dp_name"`
            SizeInfo map[string]interface{}  `json:"size_info"`
            DeliveryAddress DeliveryAddress  `json:"delivery_address"`
            UserInfo ShipmentUserInfo  `json:"user_info"`
            Comment string  `json:"comment"`
         
    }
    
    // BagsForReorderArticleAssignment ...
    type BagsForReorderArticleAssignment struct {

        
            Level string  `json:"level"`
            Strategy string  `json:"strategy"`
         
    }
    
    // BagsForReorder ...
    type BagsForReorder struct {

        
            ItemSize string  `json:"item_size"`
            ItemID float64  `json:"item_id"`
            ArticleAssignment BagsForReorderArticleAssignment  `json:"article_assignment"`
            SellerID float64  `json:"seller_id"`
            StoreID float64  `json:"store_id"`
            Quantity float64  `json:"quantity"`
         
    }
    
    // UserInfo ...
    type UserInfo struct {

        
            Mobile string  `json:"mobile"`
            Gender string  `json:"gender"`
            Email string  `json:"email"`
            Name string  `json:"name"`
         
    }
    
    // OrderSchema ...
    type OrderSchema struct {

        
            TotalShipmentsInOrder float64  `json:"total_shipments_in_order"`
            OrderID string  `json:"order_id"`
            OrderCreatedTime string  `json:"order_created_time"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            Shipments []Shipments  `json:"shipments"`
            BagsForReorder []BagsForReorder  `json:"bags_for_reorder"`
            UserInfo UserInfo  `json:"user_info"`
         
    }
    
    // OrderPage ...
    type OrderPage struct {

        
            ItemTotal float64  `json:"item_total"`
            Size float64  `json:"size"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
            Current float64  `json:"current"`
         
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
            PresignedType string  `json:"presigned_type"`
            Success bool  `json:"success"`
            ShipmentID string  `json:"shipment_id"`
         
    }
    
    // Track ...
    type Track struct {

        
            Status string  `json:"status"`
            AccountName string  `json:"account_name"`
            UpdatedAt string  `json:"updated_at"`
            ShipmentType string  `json:"shipment_type"`
            Reason string  `json:"reason"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            UpdatedTime string  `json:"updated_time"`
            Awb string  `json:"awb"`
         
    }
    
    // ShipmentTrack ...
    type ShipmentTrack struct {

        
            Results []Track  `json:"results"`
         
    }
    
    // CustomerDetailsResponse ...
    type CustomerDetailsResponse struct {

        
            Name string  `json:"name"`
            OrderID string  `json:"order_id"`
            ShipmentID string  `json:"shipment_id"`
            Country string  `json:"country"`
            Phone string  `json:"phone"`
         
    }
    
    // SendOtpToCustomerResponse ...
    type SendOtpToCustomerResponse struct {

        
            Message string  `json:"message"`
            RequestID string  `json:"request_id"`
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
    
    // BagReasons ...
    type BagReasons struct {

        
            QuestionSet []QuestionSet  `json:"question_set"`
            DisplayName string  `json:"display_name"`
            QcType []string  `json:"qc_type"`
            ID float64  `json:"id"`
         
    }
    
    // ShipmentBagReasons ...
    type ShipmentBagReasons struct {

        
            Reasons []BagReasons  `json:"reasons"`
            Success bool  `json:"success"`
         
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
    
    // StatusesBody ...
    type StatusesBody struct {

        
            Shipments map[string]interface{}  `json:"shipments"`
            ExcludeBagsNextState string  `json:"exclude_bags_next_state"`
            Status string  `json:"status"`
         
    }
    
    // ShipmentStatusUpdateBody ...
    type ShipmentStatusUpdateBody struct {

        
            LockAfterTransition bool  `json:"lock_after_transition"`
            ForceTransition bool  `json:"force_transition"`
            Statuses []StatusesBody  `json:"statuses"`
            Task bool  `json:"task"`
            UnlockBeforeTransition bool  `json:"unlock_before_transition"`
         
    }
    
    // StatusesBodyResponse ...
    type StatusesBodyResponse struct {

        
            Shipments map[string]interface{}  `json:"shipments"`
         
    }
    
    // ShipmentApplicationStatusResponse ...
    type ShipmentApplicationStatusResponse struct {

        
            Statuses []StatusesBodyResponse  `json:"statuses"`
         
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

        
            OrderType string  `json:"order_type"`
            BillingAddressID string  `json:"billing_address_id"`
            MerchantCode string  `json:"merchant_code"`
            AddressID string  `json:"address_id"`
            DeliveryAddress map[string]interface{}  `json:"delivery_address"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            BillingAddress map[string]interface{}  `json:"billing_address"`
            Meta map[string]interface{}  `json:"meta"`
            PaymentParams map[string]interface{}  `json:"payment_params"`
            PaymentMode string  `json:"payment_mode"`
            Staff StaffCheckout  `json:"staff"`
            Files []Files  `json:"files"`
            CallbackURL string  `json:"callback_url"`
            PickAtStoreUID float64  `json:"pick_at_store_uid"`
            PaymentIdentifier string  `json:"payment_identifier"`
            Pos bool  `json:"pos"`
            OrderingStore float64  `json:"ordering_store"`
            Aggregator string  `json:"aggregator"`
            CustomerDetails CustomerDetails  `json:"customer_details"`
            ExtraMeta map[string]interface{}  `json:"extra_meta"`
         
    }
    
    // CartDeliveryModesResponse ...
    type CartDeliveryModesResponse struct {

        
            PickupStores []float64  `json:"pickup_stores"`
            AvailableModes []string  `json:"available_modes"`
         
    }
    
    // PickupStoreDetail ...
    type PickupStoreDetail struct {

        
            City string  `json:"city"`
            Phone string  `json:"phone"`
            Country string  `json:"country"`
            Pincode float64  `json:"pincode"`
            State string  `json:"state"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Name string  `json:"name"`
            Area string  `json:"area"`
            StoreCode string  `json:"store_code"`
            UID float64  `json:"uid"`
            ID float64  `json:"id"`
            Email string  `json:"email"`
            AreaCode string  `json:"area_code"`
            Address string  `json:"address"`
            AddressType string  `json:"address_type"`
            Landmark string  `json:"landmark"`
         
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
    
