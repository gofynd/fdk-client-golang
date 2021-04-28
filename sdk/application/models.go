package application



    
    // Media ...
    type Media struct {

        
            URL string  `json:"url"`
            Type string  `json:"type"`
         
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
    
    // ProductListingActionPage ...
    type ProductListingActionPage struct {

        
            Query *map[string]interface{}  `json:"query"`
            Params *map[string]interface{}  `json:"params"`
            Type string  `json:"type"`
         
    }
    
    // ProductListingAction ...
    type ProductListingAction struct {

        
            Page ProductListingActionPage  `json:"page"`
            Type string  `json:"type"`
         
    }
    
    // ProductBrand ...
    type ProductBrand struct {

        
            Name string  `json:"name"`
            Action ProductListingAction  `json:"action"`
            UID int  `json:"uid"`
            Logo Media  `json:"logo"`
         
    }
    
    // ProductDetail ...
    type ProductDetail struct {

        
            Slug string  `json:"slug"`
            ItemType string  `json:"item_type"`
            Similars []string  `json:"similars"`
            Medias []Media  `json:"medias"`
            Highlights []string  `json:"highlights"`
            TeaserTag string  `json:"teaser_tag"`
            HasVariant bool  `json:"has_variant"`
            Attributes *map[string]interface{}  `json:"attributes"`
            Name string  `json:"name"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Description string  `json:"description"`
            UID int  `json:"uid"`
            ItemCode string  `json:"item_code"`
            Color string  `json:"color"`
            ProductOnlineDate string  `json:"product_online_date"`
            Tryouts []string  `json:"tryouts"`
            RatingCount int  `json:"rating_count"`
            Categories []ProductBrand  `json:"categories"`
            ImageNature string  `json:"image_nature"`
            Type string  `json:"type"`
            Brand ProductBrand  `json:"brand"`
            ShortDescription string  `json:"short_description"`
            Rating int  `json:"rating"`
         
    }
    
    // ErrorResponse ...
    type ErrorResponse struct {

        
            Error string  `json:"error"`
         
    }
    
    // ProductSizeStores ...
    type ProductSizeStores struct {

        
            Count int  `json:"count"`
         
    }
    
    // ColumnHeader ...
    type ColumnHeader struct {

        
            Convertable bool  `json:"convertable"`
            Value string  `json:"value"`
         
    }
    
    // ColumnHeaders ...
    type ColumnHeaders struct {

        
            Col6 ColumnHeader  `json:"col_6"`
            Col1 ColumnHeader  `json:"col_1"`
            Col3 ColumnHeader  `json:"col_3"`
            Col5 ColumnHeader  `json:"col_5"`
            Col2 ColumnHeader  `json:"col_2"`
            Col4 ColumnHeader  `json:"col_4"`
         
    }
    
    // SizeChartValues ...
    type SizeChartValues struct {

        
            Col6 string  `json:"col_6"`
            Col1 string  `json:"col_1"`
            Col3 string  `json:"col_3"`
            Col5 string  `json:"col_5"`
            Col2 string  `json:"col_2"`
            Col4 string  `json:"col_4"`
         
    }
    
    // SizeChart ...
    type SizeChart struct {

        
            Unit string  `json:"unit"`
            Headers ColumnHeaders  `json:"headers"`
            Image string  `json:"image"`
            Title string  `json:"title"`
            SizeTip string  `json:"size_tip"`
            Description string  `json:"description"`
            Sizes []SizeChartValues  `json:"sizes"`
         
    }
    
    // ProductSize ...
    type ProductSize struct {

        
            Quantity int  `json:"quantity"`
            Display string  `json:"display"`
            Value string  `json:"value"`
            IsAvailable bool  `json:"is_available"`
         
    }
    
    // Price ...
    type Price struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            Max int  `json:"max"`
            Min int  `json:"min"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ProductListingPrice ...
    type ProductListingPrice struct {

        
            Effective Price  `json:"effective"`
            Marked Price  `json:"marked"`
         
    }
    
    // ProductSizes ...
    type ProductSizes struct {

        
            Sellable bool  `json:"sellable"`
            Stores ProductSizeStores  `json:"stores"`
            Discount string  `json:"discount"`
            SizeChart SizeChart  `json:"size_chart"`
            Sizes []ProductSize  `json:"sizes"`
            Price ProductListingPrice  `json:"price"`
         
    }
    
    // ArticleAssignment ...
    type ArticleAssignment struct {

        
            Strategy string  `json:"strategy"`
            Level string  `json:"level"`
         
    }
    
    // Seller ...
    type Seller struct {

        
            Count int  `json:"count"`
            Name string  `json:"name"`
            UID int  `json:"uid"`
         
    }
    
    // Store ...
    type Store struct {

        
            Count int  `json:"count"`
            Name string  `json:"name"`
            UID int  `json:"uid"`
         
    }
    
    // ProductSetDistributionSize ...
    type ProductSetDistributionSize struct {

        
            Size string  `json:"size"`
            Pieces int  `json:"pieces"`
         
    }
    
    // ProductSetDistribution ...
    type ProductSetDistribution struct {

        
            Sizes []ProductSetDistributionSize  `json:"sizes"`
         
    }
    
    // ProductSet ...
    type ProductSet struct {

        
            Quantity int  `json:"quantity"`
            SizeDistribution ProductSetDistribution  `json:"size_distribution"`
         
    }
    
    // ProductStockPrice ...
    type ProductStockPrice struct {

        
            Effective int  `json:"effective"`
            Marked int  `json:"marked"`
            Currency string  `json:"currency"`
         
    }
    
    // ProductSizePriceResponse ...
    type ProductSizePriceResponse struct {

        
            ArticleID string  `json:"article_id"`
            Pincode int  `json:"pincode"`
            SpecialBadge string  `json:"special_badge"`
            Quantity int  `json:"quantity"`
            ArticleAssignment ArticleAssignment  `json:"article_assignment"`
            Seller Seller  `json:"seller"`
            LongLat []int  `json:"long_lat"`
            StrategyWiseListing []*map[string]interface{}  `json:"strategy_wise_listing"`
            ItemType string  `json:"item_type"`
            Store Store  `json:"store"`
            SellerCount int  `json:"seller_count"`
            Discount string  `json:"discount"`
            Set ProductSet  `json:"set"`
            Price ProductStockPrice  `json:"price"`
            PricePerPrice ProductStockPrice  `json:"price_per_price"`
         
    }
    
    // ProductSizeSellerFilter ...
    type ProductSizeSellerFilter struct {

        
            Value string  `json:"value"`
            Name string  `json:"name"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // ProductSizeSellersResponse ...
    type ProductSizeSellersResponse struct {

        
            Page Page  `json:"page"`
            SortOn []ProductSizeSellerFilter  `json:"sort_on"`
            Items []ProductSizePriceResponse  `json:"items"`
         
    }
    
    // AttributeDetail ...
    type AttributeDetail struct {

        
            Key string  `json:"key"`
            Description string  `json:"description"`
            Display string  `json:"display"`
         
    }
    
    // ProductsComparisonResponse ...
    type ProductsComparisonResponse struct {

        
            AttributesMetadata []AttributeDetail  `json:"attributes_metadata"`
            Items []ProductDetail  `json:"items"`
         
    }
    
    // ProductCompareResponse ...
    type ProductCompareResponse struct {

        
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
            AttributesMetadata []AttributeDetail  `json:"attributes_metadata"`
            Items []ProductDetail  `json:"items"`
         
    }
    
    // ProductFrequentlyComparedSimilarResponse ...
    type ProductFrequentlyComparedSimilarResponse struct {

        
            Similars []ProductCompareResponse  `json:"similars"`
         
    }
    
    // ProductSimilarItem ...
    type ProductSimilarItem struct {

        
            Subtitle string  `json:"subtitle"`
            Title string  `json:"title"`
            Items []ProductDetail  `json:"items"`
         
    }
    
    // SimilarProductByTypeResponse ...
    type SimilarProductByTypeResponse struct {

        
            Similars []ProductSimilarItem  `json:"similars"`
         
    }
    
    // ProductVariantItemResponse ...
    type ProductVariantItemResponse struct {

        
            IsAvailable bool  `json:"is_available"`
            Color string  `json:"color"`
            ColorName string  `json:"color_name"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Action ProductListingAction  `json:"action"`
            Medias []Media  `json:"medias"`
            UID int  `json:"uid"`
            Value string  `json:"value"`
         
    }
    
    // ProductVariantResponse ...
    type ProductVariantResponse struct {

        
            DisplayType string  `json:"display_type"`
            Header string  `json:"header"`
            Items []ProductVariantItemResponse  `json:"items"`
         
    }
    
    // ProductVariantsResponse ...
    type ProductVariantsResponse struct {

        
            Variants []ProductVariantResponse  `json:"variants"`
         
    }
    
    // StoreDetail ...
    type StoreDetail struct {

        
            Name string  `json:"name"`
            ID int  `json:"id"`
            City string  `json:"city"`
            Code string  `json:"code"`
         
    }
    
    // CompanyDetail ...
    type CompanyDetail struct {

        
            Name string  `json:"name"`
            ID int  `json:"id"`
         
    }
    
    // ProductStockStatusItem ...
    type ProductStockStatusItem struct {

        
            Seller Seller  `json:"seller"`
            Quantity int  `json:"quantity"`
            Size string  `json:"size"`
            Identifier *map[string]interface{}  `json:"identifier"`
            ItemID int  `json:"item_id"`
            Store StoreDetail  `json:"store"`
            UID string  `json:"uid"`
            Company CompanyDetail  `json:"company"`
            Price ProductStockPrice  `json:"price"`
         
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
    
    // ProductFiltersKey ...
    type ProductFiltersKey struct {

        
            Name string  `json:"name"`
            Display string  `json:"display"`
            Kind string  `json:"kind"`
            Logo string  `json:"logo"`
         
    }
    
    // ProductFiltersValue ...
    type ProductFiltersValue struct {

        
            Min int  `json:"min"`
            SelectedMax int  `json:"selected_max"`
            Count int  `json:"count"`
            SelectedMin int  `json:"selected_min"`
            Max int  `json:"max"`
            CurrencyCode string  `json:"currency_code"`
            Value string  `json:"value"`
            DisplayFormat string  `json:"display_format"`
            CurrencySymbol string  `json:"currency_symbol"`
            QueryFormat string  `json:"query_format"`
            Display string  `json:"display"`
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
            IsSelected bool  `json:"is_selected"`
            Value string  `json:"value"`
         
    }
    
    // ProductListingDetail ...
    type ProductListingDetail struct {

        
            Slug string  `json:"slug"`
            ItemType string  `json:"item_type"`
            Similars []string  `json:"similars"`
            Medias []Media  `json:"medias"`
            Highlights []string  `json:"highlights"`
            Discount string  `json:"discount"`
            TeaserTag string  `json:"teaser_tag"`
            HasVariant bool  `json:"has_variant"`
            Attributes *map[string]interface{}  `json:"attributes"`
            Name string  `json:"name"`
            GroupedAttributes []ProductDetailGroupedAttribute  `json:"grouped_attributes"`
            Description string  `json:"description"`
            UID int  `json:"uid"`
            ItemCode string  `json:"item_code"`
            Color string  `json:"color"`
            ProductOnlineDate string  `json:"product_online_date"`
            Tryouts []string  `json:"tryouts"`
            Sellable bool  `json:"sellable"`
            Price ProductListingPrice  `json:"price"`
            RatingCount int  `json:"rating_count"`
            Categories []ProductBrand  `json:"categories"`
            ImageNature string  `json:"image_nature"`
            Type string  `json:"type"`
            Brand ProductBrand  `json:"brand"`
            ShortDescription string  `json:"short_description"`
            Rating int  `json:"rating"`
         
    }
    
    // ProductListingResponse ...
    type ProductListingResponse struct {

        
            Filters []ProductFilters  `json:"filters"`
            Page Page  `json:"page"`
            SortOn []ProductSortOn  `json:"sort_on"`
            Items []ProductListingDetail  `json:"items"`
         
    }
    
    // ImageUrls ...
    type ImageUrls struct {

        
            Landscape Media  `json:"landscape"`
            Portrait Media  `json:"portrait"`
         
    }
    
    // BrandItem ...
    type BrandItem struct {

        
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            Action ProductListingAction  `json:"action"`
            Slug string  `json:"slug"`
            Discount string  `json:"discount"`
            UID int  `json:"uid"`
            Departments []string  `json:"departments"`
            Logo Media  `json:"logo"`
         
    }
    
    // BrandListingResponse ...
    type BrandListingResponse struct {

        
            Page Page  `json:"page"`
            Items []BrandItem  `json:"items"`
         
    }
    
    // BrandDetailResponse ...
    type BrandDetailResponse struct {

        
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            UID int  `json:"uid"`
            Logo Media  `json:"logo"`
         
    }
    
    // CategoryItems ...
    type CategoryItems struct {

        
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            Action ProductListingAction  `json:"action"`
            Slug string  `json:"slug"`
            Childs []*map[string]interface{}  `json:"childs"`
            UID int  `json:"uid"`
         
    }
    
    // DepartmentCategoryTree ...
    type DepartmentCategoryTree struct {

        
            Department string  `json:"department"`
            Items []CategoryItems  `json:"items"`
         
    }
    
    // DepartmentIdentifier ...
    type DepartmentIdentifier struct {

        
            Slug string  `json:"slug"`
            UID int  `json:"uid"`
         
    }
    
    // CategoryListingResponse ...
    type CategoryListingResponse struct {

        
            Data []DepartmentCategoryTree  `json:"data"`
            Departments []DepartmentIdentifier  `json:"departments"`
         
    }
    
    // CategoryMetaResponse ...
    type CategoryMetaResponse struct {

        
            Banners ImageUrls  `json:"banners"`
            Name string  `json:"name"`
            UID int  `json:"uid"`
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

        
            PriorityOrder int  `json:"priority_order"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            UID int  `json:"uid"`
            Logo Media  `json:"logo"`
         
    }
    
    // DepartmentResponse ...
    type DepartmentResponse struct {

        
            Items []Department  `json:"items"`
         
    }
    
    // AutocompleteItem ...
    type AutocompleteItem struct {

        
            Action ProductListingAction  `json:"action"`
            Display string  `json:"display"`
            Type string  `json:"type"`
            Logo Media  `json:"logo"`
         
    }
    
    // AutoCompleteResponse ...
    type AutoCompleteResponse struct {

        
            Items []AutocompleteItem  `json:"items"`
         
    }
    
    // CollectionListingFilterType ...
    type CollectionListingFilterType struct {

        
            Name string  `json:"name"`
            Display string  `json:"display"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // CollectionListingFilterTag ...
    type CollectionListingFilterTag struct {

        
            Name string  `json:"name"`
            Display string  `json:"display"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // CollectionListingFilter ...
    type CollectionListingFilter struct {

        
            Type []CollectionListingFilterType  `json:"type"`
            Tags []CollectionListingFilterTag  `json:"tags"`
         
    }
    
    // GetCollectionDetailNest ...
    type GetCollectionDetailNest struct {

        
            AllowSort bool  `json:"allow_sort"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Badge *map[string]interface{}  `json:"badge"`
            Banners ImageUrls  `json:"banners"`
            Slug string  `json:"slug"`
            Action ProductListingAction  `json:"action"`
            Query *map[string]interface{}  `json:"query"`
            Schedule *map[string]interface{}  `json:"_schedule"`
            Tag []string  `json:"tag"`
            Name string  `json:"name"`
            Description string  `json:"description"`
            UID string  `json:"uid"`
            Logo Media  `json:"logo"`
            AppID string  `json:"app_id"`
            Cron *map[string]interface{}  `json:"cron"`
            IsActive bool  `json:"is_active"`
            AllowFacets bool  `json:"allow_facets"`
            Meta *map[string]interface{}  `json:"meta"`
            Type string  `json:"type"`
         
    }
    
    // GetCollectionListingResponse ...
    type GetCollectionListingResponse struct {

        
            Filters CollectionListingFilter  `json:"filters"`
            Page Page  `json:"page"`
            Items []GetCollectionDetailNest  `json:"items"`
         
    }
    
    // CollectionDetailResponse ...
    type CollectionDetailResponse struct {

        
            AllowSort bool  `json:"allow_sort"`
            AllowFacets bool  `json:"allow_facets"`
            VisibleFacetsKeys []string  `json:"visible_facets_keys"`
            Badge *map[string]interface{}  `json:"badge"`
            Schedule *map[string]interface{}  `json:"_schedule"`
            Banners ImageUrls  `json:"banners"`
            AppID string  `json:"app_id"`
            Name string  `json:"name"`
            Cron *map[string]interface{}  `json:"cron"`
            Slug string  `json:"slug"`
            Tag []string  `json:"tag"`
            Meta *map[string]interface{}  `json:"meta"`
            Logo Media  `json:"logo"`
            Description string  `json:"description"`
            Type string  `json:"type"`
            Query *map[string]interface{}  `json:"query"`
            IsActive bool  `json:"is_active"`
         
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

        
            Count int  `json:"count"`
         
    }
    
    // FollowIdsData ...
    type FollowIdsData struct {

        
            Brands []int  `json:"brands"`
            Collections []int  `json:"collections"`
            Products []int  `json:"products"`
         
    }
    
    // FollowIdsResponse ...
    type FollowIdsResponse struct {

        
            Data FollowIdsData  `json:"data"`
         
    }
    
    // LatLong ...
    type LatLong struct {

        
            Coordinates []int  `json:"coordinates"`
            Type string  `json:"type"`
         
    }
    
    // Store1 ...
    type Store1 struct {

        
            StoreEmail string  `json:"store_email"`
            Pincode int  `json:"pincode"`
            Country string  `json:"country"`
            Address string  `json:"address"`
            Name string  `json:"name"`
            City string  `json:"city"`
            State string  `json:"state"`
            LatLong LatLong  `json:"lat_long"`
            UID int  `json:"uid"`
            StoreCode string  `json:"store_code"`
         
    }
    
    // StoreListingResponse ...
    type StoreListingResponse struct {

        
            Page Page  `json:"page"`
            Items []Store1  `json:"items"`
         
    }
    

    
    // CartCurrency ...
    type CartCurrency struct {

        
            Code string  `json:"code"`
            Symbol string  `json:"symbol"`
         
    }
    
    // PromiseTimestamp ...
    type PromiseTimestamp struct {

        
            Max int  `json:"max"`
            Min int  `json:"min"`
         
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
    
    // CouponBreakup ...
    type CouponBreakup struct {

        
            IsApplied bool  `json:"is_applied"`
            Value int  `json:"value"`
            Code string  `json:"code"`
            UID string  `json:"uid"`
            Type string  `json:"type"`
            Message string  `json:"message"`
         
    }
    
    // RawBreakup ...
    type RawBreakup struct {

        
            YouSaved int  `json:"you_saved"`
            CodCharge int  `json:"cod_charge"`
            Total int  `json:"total"`
            Coupon int  `json:"coupon"`
            FyndCash int  `json:"fynd_cash"`
            Subtotal int  `json:"subtotal"`
            DeliveryCharge int  `json:"delivery_charge"`
         
    }
    
    // DisplayBreakup ...
    type DisplayBreakup struct {

        
            Value int  `json:"value"`
            CurrencySymbol string  `json:"currency_symbol"`
            Message []string  `json:"message"`
            Key string  `json:"key"`
            Display string  `json:"display"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // LoyaltyPoints ...
    type LoyaltyPoints struct {

        
            Applicable int  `json:"applicable"`
            Total int  `json:"total"`
            IsApplied bool  `json:"is_applied"`
            Description string  `json:"description"`
         
    }
    
    // CartBreakup ...
    type CartBreakup struct {

        
            Coupon CouponBreakup  `json:"coupon"`
            Raw RawBreakup  `json:"raw"`
            Display []DisplayBreakup  `json:"display"`
            LoyaltyPoints LoyaltyPoints  `json:"loyalty_points"`
         
    }
    
    // CartProductIdentifer ...
    type CartProductIdentifer struct {

        
            Identifier string  `json:"identifier"`
         
    }
    
    // ProductAvailability ...
    type ProductAvailability struct {

        
            OtherStoreQuantity int  `json:"other_store_quantity"`
            Sizes []string  `json:"sizes"`
            OutOfStock bool  `json:"out_of_stock"`
            Deliverable bool  `json:"deliverable"`
            IsValid bool  `json:"is_valid"`
         
    }
    
    // ProductPrice ...
    type ProductPrice struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            Selling int  `json:"selling"`
            Effective int  `json:"effective"`
            Marked int  `json:"marked"`
            AddOn int  `json:"add_on"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // ProductPriceInfo ...
    type ProductPriceInfo struct {

        
            Base ProductPrice  `json:"base"`
            Converted ProductPrice  `json:"converted"`
         
    }
    
    // BaseInfo ...
    type BaseInfo struct {

        
            Name string  `json:"name"`
            UID int  `json:"uid"`
         
    }
    
    // BasePrice ...
    type BasePrice struct {

        
            CurrencySymbol string  `json:"currency_symbol"`
            Effective int  `json:"effective"`
            CurrencyCode string  `json:"currency_code"`
            Marked int  `json:"marked"`
         
    }
    
    // ArticlePriceInfo ...
    type ArticlePriceInfo struct {

        
            Base BasePrice  `json:"base"`
            Converted BasePrice  `json:"converted"`
         
    }
    
    // ProductArticle ...
    type ProductArticle struct {

        
            Seller BaseInfo  `json:"seller"`
            Size string  `json:"size"`
            Price ArticlePriceInfo  `json:"price"`
            Store BaseInfo  `json:"store"`
            Quantity int  `json:"quantity"`
            UID string  `json:"uid"`
            Type string  `json:"type"`
         
    }
    
    // Image ...
    type Image struct {

        
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

        
            URL string  `json:"url"`
            Type string  `json:"type"`
            Query ActionQuery  `json:"query"`
         
    }
    
    // CategoryInfo ...
    type CategoryInfo struct {

        
            Name string  `json:"name"`
            UID int  `json:"uid"`
         
    }
    
    // Product ...
    type Product struct {

        
            Slug string  `json:"slug"`
            Images []Image  `json:"images"`
            Action ProductAction  `json:"action"`
            Categories []CategoryInfo  `json:"categories"`
            Name string  `json:"name"`
            Brand BaseInfo  `json:"brand"`
            UID int  `json:"uid"`
            Type string  `json:"type"`
         
    }
    
    // CartProductInfo ...
    type CartProductInfo struct {

        
            Identifiers CartProductIdentifer  `json:"identifiers"`
            Discount string  `json:"discount"`
            Availability ProductAvailability  `json:"availability"`
            BulkOffer *map[string]interface{}  `json:"bulk_offer"`
            IsSet bool  `json:"is_set"`
            CouponMessage string  `json:"coupon_message"`
            Price ProductPriceInfo  `json:"price"`
            Key string  `json:"key"`
            Article ProductArticle  `json:"article"`
            Quantity int  `json:"quantity"`
            Product Product  `json:"product"`
            Message string  `json:"message"`
         
    }
    
    // CartResponse ...
    type CartResponse struct {

        
            CartID int  `json:"cart_id"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Currency CartCurrency  `json:"currency"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            IsValid bool  `json:"is_valid"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            LastModified string  `json:"last_modified"`
            Items []CartProductInfo  `json:"items"`
            UID string  `json:"uid"`
            Message string  `json:"message"`
            CouponText string  `json:"coupon_text"`
         
    }
    
    // AddProductCart ...
    type AddProductCart struct {

        
            ItemID int  `json:"item_id"`
            SellerID int  `json:"seller_id"`
            ArticleID string  `json:"article_id"`
            ArticleAssignment *map[string]interface{}  `json:"article_assignment"`
            Pos bool  `json:"pos"`
            ItemSize string  `json:"item_size"`
            Display string  `json:"display"`
            StoreID int  `json:"store_id"`
            Quantity int  `json:"quantity"`
         
    }
    
    // AddCartRequest ...
    type AddCartRequest struct {

        
            Items []AddProductCart  `json:"items"`
         
    }
    
    // AddCartResponse ...
    type AddCartResponse struct {

        
            Cart CartResponse  `json:"cart"`
            Success bool  `json:"success"`
            Partial bool  `json:"partial"`
            Message string  `json:"message"`
         
    }
    
    // UpdateProductCart ...
    type UpdateProductCart struct {

        
            Identifiers CartProductIdentifer  `json:"identifiers"`
            ArticleID string  `json:"article_id"`
            ItemID int  `json:"item_id"`
            ItemIndex int  `json:"item_index"`
            ItemSize string  `json:"item_size"`
            Quantity int  `json:"quantity"`
         
    }
    
    // UpdateCartRequest ...
    type UpdateCartRequest struct {

        
            Items []UpdateProductCart  `json:"items"`
            Operation string  `json:"operation"`
         
    }
    
    // UpdateCartResponse ...
    type UpdateCartResponse struct {

        
            Cart CartResponse  `json:"cart"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // CartItemCountResponse ...
    type CartItemCountResponse struct {

        
            UserCartItemsCount int  `json:"user_cart_items_count"`
         
    }
    
    // Coupon ...
    type Coupon struct {

        
            IsApplied bool  `json:"is_applied"`
            ExpiresOn string  `json:"expires_on"`
            MaxDiscountValue int  `json:"max_discount_value"`
            SubTitle string  `json:"sub_title"`
            Title string  `json:"title"`
            CouponValue int  `json:"coupon_value"`
            MinimumCartValue int  `json:"minimum_cart_value"`
            IsApplicable bool  `json:"is_applicable"`
            CouponCode string  `json:"coupon_code"`
            Message string  `json:"message"`
         
    }
    
    // PageCoupon ...
    type PageCoupon struct {

        
            HasPrevious bool  `json:"has_previous"`
            Total int  `json:"total"`
            Current int  `json:"current"`
            TotalItemCount int  `json:"total_item_count"`
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

        
            CurrencySymbol string  `json:"currency_symbol"`
            BulkEffective int  `json:"bulk_effective"`
            Effective int  `json:"effective"`
            Marked int  `json:"marked"`
            CurrencyCode string  `json:"currency_code"`
         
    }
    
    // OfferItem ...
    type OfferItem struct {

        
            Margin int  `json:"margin"`
            Total int  `json:"total"`
            Best bool  `json:"best"`
            Price OfferPrice  `json:"price"`
            AutoApplied bool  `json:"auto_applied"`
            Quantity int  `json:"quantity"`
            Type string  `json:"type"`
         
    }
    
    // OfferSeller ...
    type OfferSeller struct {

        
            Name string  `json:"name"`
            UID int  `json:"uid"`
         
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
    
    // GeoLocation ...
    type GeoLocation struct {

        
            Latitude int  `json:"latitude"`
            Longitude int  `json:"longitude"`
         
    }
    
    // Address ...
    type Address struct {

        
            State string  `json:"state"`
            UserID string  `json:"user_id"`
            Name string  `json:"name"`
            Phone string  `json:"phone"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Tags []string  `json:"tags"`
            AreaCode string  `json:"area_code"`
            GoogleMapPoint *map[string]interface{}  `json:"google_map_point"`
            AddressType string  `json:"address_type"`
            GeoLocation GeoLocation  `json:"geo_location"`
            Area string  `json:"area"`
            IsDefaultAddress bool  `json:"is_default_address"`
            Meta *map[string]interface{}  `json:"meta"`
            Email string  `json:"email"`
            City string  `json:"city"`
            CountryCode string  `json:"country_code"`
            CheckoutMode string  `json:"checkout_mode"`
            Country string  `json:"country"`
            IsActive bool  `json:"is_active"`
            Landmark string  `json:"landmark"`
            Address string  `json:"address"`
            UID int  `json:"uid"`
         
    }
    
    // GetAddressesResponse ...
    type GetAddressesResponse struct {

        
            Address []Address  `json:"address"`
         
    }
    
    // SaveAddressResponse ...
    type SaveAddressResponse struct {

        
            Success string  `json:"success"`
            IsDefaultAddress bool  `json:"is_default_address"`
            AddressID int  `json:"address_id"`
         
    }
    
    // UpdateAddressResponse ...
    type UpdateAddressResponse struct {

        
            Success bool  `json:"success"`
            IsDefaultAddress bool  `json:"is_default_address"`
            IsUpdated bool  `json:"is_updated"`
            AddressID int  `json:"address_id"`
         
    }
    
    // DeleteAddressResponse ...
    type DeleteAddressResponse struct {

        
            IsDeleted bool  `json:"is_deleted"`
            AddressID int  `json:"address_id"`
         
    }
    
    // SelectCartAddressRequest ...
    type SelectCartAddressRequest struct {

        
            BillingAddressID int  `json:"billing_address_id"`
            UID string  `json:"uid"`
            AddressID string  `json:"address_id"`
         
    }
    
    // UpdateCartPaymentRequest ...
    type UpdateCartPaymentRequest struct {

        
            PaymentIdentifier string  `json:"payment_identifier"`
            MerchantCode string  `json:"merchant_code"`
            PaymentMode string  `json:"payment_mode"`
            AggregatorName string  `json:"aggregator_name"`
            UID int  `json:"uid"`
            AddressID int  `json:"address_id"`
         
    }
    
    // CouponValidity ...
    type CouponValidity struct {

        
            Discount int  `json:"discount"`
            DisplayMessageEn string  `json:"display_message_en"`
            Title string  `json:"title"`
            Code string  `json:"code"`
            Valid bool  `json:"valid"`
         
    }
    
    // PaymentCouponValidate ...
    type PaymentCouponValidate struct {

        
            Success bool  `json:"success"`
            Message string  `json:"message"`
            CouponValidity CouponValidity  `json:"coupon_validity"`
         
    }
    
    // ShipmentResponse ...
    type ShipmentResponse struct {

        
            FulfillmentType string  `json:"fulfillment_type"`
            FulfillmentID int  `json:"fulfillment_id"`
            ShipmentType string  `json:"shipment_type"`
            DpID int  `json:"dp_id"`
            Promise ShipmentPromise  `json:"promise"`
            OrderType string  `json:"order_type"`
            DpOptions *map[string]interface{}  `json:"dp_options"`
            Items []CartProductInfo  `json:"items"`
            Shipments int  `json:"shipments"`
            BoxType string  `json:"box_type"`
         
    }
    
    // CartShipmentsResponse ...
    type CartShipmentsResponse struct {

        
            CartID int  `json:"cart_id"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Currency CartCurrency  `json:"currency"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            Error bool  `json:"error"`
            IsValid bool  `json:"is_valid"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            LastModified string  `json:"last_modified"`
            Shipments []ShipmentResponse  `json:"shipments"`
            UID string  `json:"uid"`
            Message string  `json:"message"`
            CouponText string  `json:"coupon_text"`
         
    }
    
    // CartCheckoutRequest ...
    type CartCheckoutRequest struct {

        
            PaymentIdentifier string  `json:"payment_identifier"`
            MerchantCode string  `json:"merchant_code"`
            Aggregator string  `json:"aggregator"`
            PaymentParams *map[string]interface{}  `json:"payment_params"`
            AddressID int  `json:"address_id"`
            CallbackURL string  `json:"callback_url"`
            ExtraMeta *map[string]interface{}  `json:"extra_meta"`
            BillingAddress *map[string]interface{}  `json:"billing_address"`
            Staff *map[string]interface{}  `json:"staff"`
            Meta *map[string]interface{}  `json:"meta"`
            OrderingStore int  `json:"ordering_store"`
            DeliveryAddress *map[string]interface{}  `json:"delivery_address"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            BillingAddressID int  `json:"billing_address_id"`
            FyndstoreEmpID string  `json:"fyndstore_emp_id"`
            PaymentMode string  `json:"payment_mode"`
         
    }
    
    // CheckCart ...
    type CheckCart struct {

        
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            CodAvailable bool  `json:"cod_available"`
            StoreCode string  `json:"store_code"`
            UserType string  `json:"user_type"`
            OrderID string  `json:"order_id"`
            Success bool  `json:"success"`
            LastModified string  `json:"last_modified"`
            Items []CartProductInfo  `json:"items"`
            CodMessage string  `json:"cod_message"`
            Currency CartCurrency  `json:"currency"`
            ErrorMessage string  `json:"error_message"`
            CodCharges int  `json:"cod_charges"`
            IsValid bool  `json:"is_valid"`
            DeliveryCharges int  `json:"delivery_charges"`
            CartID int  `json:"cart_id"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            StoreEmps []*map[string]interface{}  `json:"store_emps"`
            DeliveryChargeOrderValue int  `json:"delivery_charge_order_value"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            UID string  `json:"uid"`
            Message string  `json:"message"`
            CouponText string  `json:"coupon_text"`
         
    }
    
    // CartCheckoutResponse ...
    type CartCheckoutResponse struct {

        
            Cart CheckCart  `json:"cart"`
            Data *map[string]interface{}  `json:"data"`
            CallbackURL string  `json:"callback_url"`
            AppInterceptURL string  `json:"app_intercept_url"`
            OrderID string  `json:"order_id"`
            Success bool  `json:"success"`
            Message string  `json:"message"`
         
    }
    
    // CartMetaRequest ...
    type CartMetaRequest struct {

        
            PickUpCustomerDetails *map[string]interface{}  `json:"pick_up_customer_details"`
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

        
            UID int  `json:"uid"`
            Meta *map[string]interface{}  `json:"meta"`
         
    }
    
    // GetShareCartLinkResponse ...
    type GetShareCartLinkResponse struct {

        
            ShareURL string  `json:"share_url"`
            Token string  `json:"token"`
         
    }
    
    // SharedCartDetails ...
    type SharedCartDetails struct {

        
            Source *map[string]interface{}  `json:"source"`
            Meta *map[string]interface{}  `json:"meta"`
            CreatedOn string  `json:"created_on"`
            Token string  `json:"token"`
            User *map[string]interface{}  `json:"user"`
         
    }
    
    // SharedCart ...
    type SharedCart struct {

        
            CartID int  `json:"cart_id"`
            CheckoutMode string  `json:"checkout_mode"`
            DeliveryChargeInfo string  `json:"delivery_charge_info"`
            Currency CartCurrency  `json:"currency"`
            DeliveryPromise ShipmentPromise  `json:"delivery_promise"`
            IsValid bool  `json:"is_valid"`
            RestrictCheckout bool  `json:"restrict_checkout"`
            Comment string  `json:"comment"`
            Gstin string  `json:"gstin"`
            BreakupValues CartBreakup  `json:"breakup_values"`
            LastModified string  `json:"last_modified"`
            Items []CartProductInfo  `json:"items"`
            UID string  `json:"uid"`
            SharedCartDetails SharedCartDetails  `json:"shared_cart_details"`
            Message string  `json:"message"`
            CouponText string  `json:"coupon_text"`
         
    }
    
    // SharedCartResponse ...
    type SharedCartResponse struct {

        
            Cart SharedCart  `json:"cart"`
            Error string  `json:"error"`
         
    }
    

    
    // TicketList ...
    type TicketList struct {

        
            Items []Ticket  `json:"items"`
            Filters Filter  `json:"filters"`
            Page Page  `json:"page"`
         
    }
    
    // Page ...
    type Page struct {

        
            ItemTotal int  `json:"item_total"`
            NextID string  `json:"next_id"`
            HasPrevious bool  `json:"has_previous"`
            HasNext bool  `json:"has_next"`
            Current int  `json:"current"`
            Type string  `json:"type"`
            Size int  `json:"size"`
         
    }
    
    // TicketHistoryList ...
    type TicketHistoryList struct {

        
            Docs []TicketHistory  `json:"docs"`
            Limit int  `json:"limit"`
            Page int  `json:"page"`
            Pages int  `json:"pages"`
            Total int  `json:"total"`
         
    }
    
    // CustomFormList ...
    type CustomFormList struct {

        
            Docs []CustomForm  `json:"docs"`
            Limit int  `json:"limit"`
            Page int  `json:"page"`
            Pages int  `json:"pages"`
            Total int  `json:"total"`
         
    }
    
    // CreateCustomFormPayload ...
    type CreateCustomFormPayload struct {

        
            Slug string  `json:"slug"`
            Title string  `json:"title"`
            Inputs []*map[string]interface{}  `json:"inputs"`
            Description string  `json:"description"`
            HeaderImage string  `json:"header_image"`
            ShouldNotify bool  `json:"should_notify"`
            SuccessMessage string  `json:"success_message"`
            PollForAssignment PollForAssignment  `json:"poll_for_assignment"`
         
    }
    
    // EditCustomFormPayload ...
    type EditCustomFormPayload struct {

        
            Title string  `json:"title"`
            Inputs []*map[string]interface{}  `json:"inputs"`
            Description string  `json:"description"`
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
            Source string  `json:"source"`
            Status string  `json:"status"`
            Priority string  `json:"priority"`
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
            Notify []*map[string]interface{}  `json:"notify"`
         
    }
    
    // Filter ...
    type Filter struct {

        
            Priorities []Priority  `json:"priorities"`
            Categories []TicketCategory  `json:"categories"`
            Statuses []Status  `json:"statuses"`
            Assignees []*map[string]interface{}  `json:"assignees"`
         
    }
    
    // TicketHistoryPayload ...
    type TicketHistoryPayload struct {

        
            Value *map[string]interface{}  `json:"value"`
            Type string  `json:"type"`
         
    }
    
    // CustomFormSubmissionPayload ...
    type CustomFormSubmissionPayload struct {

        
            Response []KeyValue  `json:"response"`
         
    }
    
    // KeyValue ...
    type KeyValue struct {

        
            Key string  `json:"key"`
            Value string  `json:"value"`
         
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
    
    // PhoneNumber ...
    type PhoneNumber struct {

        
            Active bool  `json:"active"`
            Primary bool  `json:"primary"`
            Verified bool  `json:"verified"`
            Phone string  `json:"phone"`
            CountryCode int  `json:"country_code"`
         
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
            Type string  `json:"type"`
         
    }
    
    // TicketContent ...
    type TicketContent struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
            Attachments []TicketAsset  `json:"attachments"`
         
    }
    
    // AddTicketPayload ...
    type AddTicketPayload struct {

        
            Status string  `json:"status"`
            Priority string  `json:"priority"`
            Category string  `json:"category"`
            Content TicketContent  `json:"content"`
         
    }
    
    // Priority ...
    type Priority struct {

        
            Key string  `json:"key"`
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
         
    }
    
    // SubmitButton ...
    type SubmitButton struct {

        
            Title string  `json:"title"`
            TitleColor string  `json:"title_color"`
            BackgroundColor string  `json:"background_color"`
         
    }
    
    // PollForAssignment ...
    type PollForAssignment struct {

        
            Duration int  `json:"duration"`
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
            LoginRequired bool  `json:"login_required"`
            ShouldNotify bool  `json:"should_notify"`
            SuccessMessage string  `json:"success_message"`
            SubmitButton SubmitButton  `json:"submit_button"`
            Inputs []*map[string]interface{}  `json:"inputs"`
            CreatedOn CreatedOn  `json:"created_on"`
            CreatedBy *map[string]interface{}  `json:"created_by"`
            PollForAssignment PollForAssignment  `json:"poll_for_assignment"`
            ID string  `json:"_id"`
         
    }
    
    // TicketHistory ...
    type TicketHistory struct {

        
            Type string  `json:"type"`
            Value *map[string]interface{}  `json:"value"`
            TicketID string  `json:"ticket_id"`
            CreatedOn CreatedOn  `json:"created_on"`
            CreatedBy *map[string]interface{}  `json:"created_by"`
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
            Source string  `json:"source"`
            Status Status  `json:"status"`
            Priority Priority  `json:"priority"`
            CreatedBy *map[string]interface{}  `json:"created_by"`
            AssignedTo *map[string]interface{}  `json:"assigned_to"`
            Tags []string  `json:"tags"`
            CustomJson *map[string]interface{}  `json:"_custom_json"`
            ID string  `json:"_id"`
            UpdatedAt string  `json:"updated_at"`
            CreatedAt string  `json:"created_at"`
         
    }
    

    
    // PaginationSchema ...
    type PaginationSchema struct {

        
            Size int  `json:"size"`
            ItemTotal int  `json:"item_total"`
            HasNext bool  `json:"has_next"`
            Type string  `json:"type"`
            Current int  `json:"current"`
         
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
            AvailablePages AvailablePages  `json:"available_pages"`
            AvailableSections []availableSectionSchema  `json:"available_sections"`
            Sections []sectionSchema  `json:"sections"`
            Constants *map[string]interface{}  `json:"constants"`
            Styles *map[string]interface{}  `json:"styles"`
            Config Config  `json:"config"`
            Settings *map[string]interface{}  `json:"settings"`
            Font Font  `json:"font"`
            ID string  `json:"_id"`
            V int  `json:"__v"`
            Colors Colors  `json:"colors"`
         
    }
    
    // pagesSchema ...
    type pagesSchema struct {

        
            Text string  `json:"text"`
            Path string  `json:"path"`
            Type string  `json:"type"`
            Value string  `json:"value"`
            Sections Sections  `json:"sections"`
         
    }
    
    // availableSectionSchema ...
    type availableSectionSchema struct {

        
            Blocks Blocks  `json:"blocks"`
            Name string  `json:"name"`
            Label string  `json:"label"`
            Props *map[string]interface{}  `json:"props"`
         
    }
    
    // sectionSchema ...
    type sectionSchema struct {

        
            PageKey string  `json:"page_key"`
            PageSections PageSections  `json:"page_sections"`
         
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
         
    }
    
    // CommonJs ...
    type CommonJs struct {

        
            Link string  `json:"link"`
         
    }
    
    // Css ...
    type Css struct {

        
            Link string  `json:"link"`
         
    }
    
    // AvailablePages ...
    type AvailablePages struct {

        
            Path string  `json:"path"`
            Type string  `json:"type"`
            Text string  `json:"text"`
            Value string  `json:"value"`
            Seo Seo  `json:"seo"`
            Props *map[string]interface{}  `json:"props"`
            Sections Sections  `json:"sections"`
         
    }
    
    // Seo ...
    type Seo struct {

        
            Title string  `json:"title"`
            Description string  `json:"description"`
         
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
            List ListSchema  `json:"list"`
         
    }
    
    // Preset ...
    type Preset struct {

        
            Sections sectionSchema  `json:"sections"`
         
    }
    
    // GlobalSchema ...
    type GlobalSchema struct {

        
            Props *map[string]interface{}  `json:"props"`
         
    }
    
    // ListSchema ...
    type ListSchema struct {

        
            Global *map[string]interface{}  `json:"global"`
            Page ConfigPage  `json:"page"`
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

        
            Props *map[string]interface{}  `json:"props"`
         
    }
    
    // ConfigPage ...
    type ConfigPage struct {

        
            Settings *map[string]interface{}  `json:"settings"`
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
            Props BlocksProps  `json:"props"`
         
    }
    
    // BlocksProps ...
    type BlocksProps struct {

        
            ID string  `json:"id"`
            Label string  `json:"label"`
            Type string  `json:"type"`
         
    }
    
    // PageSections ...
    type PageSections struct {

        
            Blocks PageSectionsBlocks  `json:"blocks"`
            Name string  `json:"name"`
            Label string  `json:"label"`
            Props *map[string]interface{}  `json:"props"`
            Preset *map[string]interface{}  `json:"preset"`
            Predicate Predicate  `json:"predicate"`
         
    }
    
    // PageSectionsBlocks ...
    type PageSectionsBlocks struct {

        
            Type string  `json:"type"`
            Name string  `json:"name"`
            Props PageSectionsBlocksProps  `json:"props"`
         
    }
    
    // PageSectionsBlocksProps ...
    type PageSectionsBlocksProps struct {

        
            ID string  `json:"id"`
            Label string  `json:"label"`
            Type string  `json:"type"`
         
    }
    
    // Predicate ...
    type Predicate struct {

        
            Screen Screen  `json:"screen"`
            User PredicateUserSchema  `json:"user"`
            Route Route  `json:"route"`
         
    }
    
    // Screen ...
    type Screen struct {

        
            Mobile bool  `json:"mobile"`
            Desktop bool  `json:"desktop"`
            Tablet bool  `json:"tablet"`
         
    }
    
    // PredicateUserSchema ...
    type PredicateUserSchema struct {

        
            Authenticated bool  `json:"authenticated"`
            Anonymous bool  `json:"anonymous"`
         
    }
    
    // Route ...
    type Route struct {

        
            Selected string  `json:"selected"`
            ExactURL string  `json:"exact_url"`
            Query *map[string]interface{}  `json:"query"`
         
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
            Mobile string  `json:"mobile"`
            Email string  `json:"email"`
            Gender string  `json:"gender"`
            Dob string  `json:"dob"`
            ProfilePicURL string  `json:"profile_pic_url"`
            AndroidHash string  `json:"android_hash"`
            Sender string  `json:"sender"`
            RegisterToken string  `json:"register_token"`
         
    }
    
    // SendEmailOtpRequestSchema ...
    type SendEmailOtpRequestSchema struct {

        
            Email string  `json:"email"`
            Action string  `json:"action"`
            Token string  `json:"token"`
            RegisterToken string  `json:"register_token"`
         
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
    
    // AuthSuccess ...
    type AuthSuccess struct {

        
            RegisterToken string  `json:"register_token"`
            UserExists bool  `json:"user_exists"`
            User UserSchema  `json:"user"`
         
    }
    
    // SendOtpResponse ...
    type SendOtpResponse struct {

        
            ResendTimer int  `json:"resend_timer"`
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
         
    }
    
    // ResetPasswordSuccess ...
    type ResetPasswordSuccess struct {

        
            Status string  `json:"status"`
         
    }
    
    // RegisterFormSuccess ...
    type RegisterFormSuccess struct {

        
            Email string  `json:"email"`
            ResendTimer int  `json:"resend_timer"`
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
    
    // OtpSuccess ...
    type OtpSuccess struct {

        
            ResendTimer int  `json:"resend_timer"`
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

        
            ResendToken string  `json:"resend_token"`
            RegisterToken string  `json:"register_token"`
         
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
            Expiry int  `json:"expiry"`
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
         
    }
    
    // RequiredFields ...
    type RequiredFields struct {

        
            Email Email  `json:"email"`
            Mobile Mobile  `json:"mobile"`
         
    }
    
    // Mobile ...
    type Mobile struct {

        
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
    
    // UserSchema ...
    type UserSchema struct {

        
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            PhoneNumbers []PhoneNumber  `json:"phone_numbers"`
            Emails []Email  `json:"emails"`
            Gender string  `json:"gender"`
            Dob string  `json:"dob"`
            Active bool  `json:"active"`
            ProfilePicURL string  `json:"profile_pic_url"`
            Username string  `json:"username"`
            AccountType string  `json:"account_type"`
            UID string  `json:"uid"`
            Debug Debug  `json:"debug"`
            HasOldPasswordHash bool  `json:"has_old_password_hash"`
            ID string  `json:"_id"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
         
    }
    

    
    // ApplicationLegal ...
    type ApplicationLegal struct {

        
            Application string  `json:"application"`
            Tnc string  `json:"tnc"`
            Policy string  `json:"policy"`
            Shipping string  `json:"shipping"`
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
            CustomMetaTags []*map[string]interface{}  `json:"custom_meta_tags"`
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
    
    // StorefrontAnnouncement ...
    type StorefrontAnnouncement struct {

        
            Announcements AnnouncementSchema  `json:"announcements"`
            RefreshRate int  `json:"refresh_rate"`
            RefreshPages []string  `json:"refresh_pages"`
         
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
            Duration int  `json:"duration"`
            NextSchedule []*map[string]interface{}  `json:"next_schedule"`
         
    }
    
    // NextSchedule ...
    type NextSchedule struct {

        
            Start string  `json:"start"`
            End string  `json:"end"`
         
    }
    
    // AnnouncementSchema ...
    type AnnouncementSchema struct {

        
            PageSlug []*map[string]interface{}  `json:"page_slug"`
         
    }
    
    // announcementSchema ...
    type announcementSchema struct {

        
            Announcement string  `json:"announcement"`
            Schedule scheduleStartSchema  `json:"schedule"`
         
    }
    
    // scheduleStartSchema ...
    type scheduleStartSchema struct {

        
            Start string  `json:"start"`
         
    }
    
    // BlogGetResponse ...
    type BlogGetResponse struct {

        
            Items []*map[string]interface{}  `json:"items"`
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
            CustomJson *map[string]interface{}  `json:"_custom_json"`
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
            CustomJson *map[string]interface{}  `json:"_custom_json"`
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

        
            Params *map[string]interface{}  `json:"params"`
            Query *map[string]interface{}  `json:"query"`
            URL string  `json:"url"`
            Type string  `json:"type"`
         
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
            SortOrder int  `json:"sort_order"`
            SubNavigation *NavigationReference  `json:"sub_navigation"`
         
    }
    
    // LandingPage ...
    type LandingPage struct {

        
            Data LandingPageSchema  `json:"data"`
            Success bool  `json:"success"`
         
    }
    
    // ConfigurationSchema ...
    type ConfigurationSchema struct {

        
            SleepTime int  `json:"sleep_time"`
            StartOnLaunch bool  `json:"start_on_launch"`
            Duration int  `json:"duration"`
            SlideDirection string  `json:"slide_direction"`
         
    }
    
    // SlideshowMedia ...
    type SlideshowMedia struct {

        
            Type string  `json:"type"`
            URL string  `json:"url"`
            BgColor string  `json:"bg_color"`
            Duration int  `json:"duration"`
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

        
            Announcements *map[string]interface{}  `json:"announcements"`
            RefreshRate int  `json:"refresh_rate"`
            RefreshPages []string  `json:"refresh_pages"`
         
    }
    
    // AnnouncementDataSchema ...
    type AnnouncementDataSchema struct {

        
            PageSlug string  `json:"page_slug"`
            Announcement StorefrontAnnouncement  `json:"announcement"`
         
    }
    
    // FaqResponseSchema ...
    type FaqResponseSchema struct {

        
            Faqs []FAQ  `json:"faqs"`
         
    }
    
    // UpdateHandpickedSchema ...
    type UpdateHandpickedSchema struct {

        
            Tag HandpickedTagSchema  `json:"tag"`
         
    }
    
    // HandpickedTagSchema ...
    type HandpickedTagSchema struct {

        
            Position string  `json:"position"`
            Attributes *map[string]interface{}  `json:"attributes"`
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
            Attributes *map[string]interface{}  `json:"attributes"`
            Content string  `json:"content"`
         
    }
    
    // CreateTagRequestSchema ...
    type CreateTagRequestSchema struct {

        
            Tags []CreateTagSchema  `json:"tags"`
         
    }
    
    // APIError ...
    type APIError struct {

        
            Message string  `json:"message"`
            Status int  `json:"status"`
            Code string  `json:"code"`
            Exception string  `json:"exception"`
            Info string  `json:"info"`
            RequestID string  `json:"request_id"`
            StackTrace string  `json:"stack_trace"`
            Meta *map[string]interface{}  `json:"meta"`
         
    }
    
    // CategorySchema ...
    type CategorySchema struct {

        
            Index int  `json:"index"`
            ID string  `json:"_id"`
            Slug string  `json:"slug"`
            Title string  `json:"title"`
            Application string  `json:"application"`
         
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

        
            Index int  `json:"index"`
            Title string  `json:"title"`
            Description string  `json:"description"`
            Children ChildrenSchema  `json:"children"`
            ID string  `json:"_id"`
            Slug string  `json:"slug"`
            Application string  `json:"application"`
            IconURL string  `json:"icon_url"`
            CustomJson *map[string]interface{}  `json:"_custom_json"`
         
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

        
            Faqs []*map[string]interface{}  `json:"faqs"`
         
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

        
            Items []*map[string]interface{}  `json:"items"`
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
            CustomJson *map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // DefaultNavigationResponse ...
    type DefaultNavigationResponse struct {

        
            Items []*map[string]interface{}  `json:"items"`
         
    }
    
    // NavigationGetResponse ...
    type NavigationGetResponse struct {

        
            Items []*map[string]interface{}  `json:"items"`
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
            Version int  `json:"version"`
            Navigation NavigationReference  `json:"navigation"`
         
    }
    
    // NavigationRequest ...
    type NavigationRequest struct {

        
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Platform []string  `json:"platform"`
            Orientation Orientation  `json:"orientation"`
            Navigation NavigationReference  `json:"navigation"`
         
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
            Content []ContentSchema  `json:"content"`
            CreatedBy CreatedBySchema  `json:"created_by"`
            DateMeta DateMeta  `json:"date_meta"`
            Schedule ScheduleSchema  `json:"_schedule"`
         
    }
    
    // ContentSchema ...
    type ContentSchema struct {

        
            Type string  `json:"type"`
            Value string  `json:"value"`
         
    }
    
    // CustomPage ...
    type CustomPage struct {

        
            Data CustomPageSchema  `json:"data"`
         
    }
    
    // CustomBlogSchema ...
    type CustomBlogSchema struct {

        
            ID string  `json:"_id"`
            Title string  `json:"title"`
            Slug string  `json:"slug"`
            ReadingTime string  `json:"reading_time"`
            Application string  `json:"application"`
            Description string  `json:"description"`
            FeatureImage FeatureImage  `json:"feature_image"`
            Published bool  `json:"published"`
            Archived bool  `json:"archived"`
            Tags []string  `json:"tags"`
            Content ContentSchema  `json:"content"`
            Author Author  `json:"author"`
            Schedule ScheduleSchema  `json:"_schedule"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            CustomJson *map[string]interface{}  `json:"_custom_json"`
         
    }
    
    // FeatureImage ...
    type FeatureImage struct {

        
            SecureURL string  `json:"secure_url"`
         
    }
    
    // CustomBlog ...
    type CustomBlog struct {

        
            Data CustomBlogSchema  `json:"data"`
         
    }
    
    // PageGetResponse ...
    type PageGetResponse struct {

        
            Items []*map[string]interface{}  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // PageSpec ...
    type PageSpec struct {

        
            Specifications []*map[string]interface{}  `json:"specifications"`
         
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
            Content []PageContent  `json:"content"`
            CreatedBy CreatedBySchema  `json:"created_by"`
            DateMeta DateMeta  `json:"date_meta"`
            Description string  `json:"description"`
            FeatureImage Asset  `json:"feature_image"`
            PageMeta []PageMeta  `json:"page_meta"`
            Schedule ScheduleSchema  `json:"_schedule"`
            CustomJson *map[string]interface{}  `json:"_custom_json"`
            Orientation string  `json:"orientation"`
            Platform string  `json:"platform"`
            Published bool  `json:"published"`
            Slug string  `json:"slug"`
            Tags []string  `json:"tags"`
            Title string  `json:"title"`
            Type string  `json:"type"`
            Seo SEO  `json:"seo"`
            Visibility *map[string]interface{}  `json:"visibility"`
         
    }
    
    // CreatedBySchema ...
    type CreatedBySchema struct {

        
            ID string  `json:"id"`
         
    }
    
    // PageContent ...
    type PageContent struct {

        
            Type string  `json:"type"`
            Value *map[string]interface{}  `json:"value"`
         
    }
    
    // PageMeta ...
    type PageMeta struct {

        
            Key string  `json:"key"`
            Value *map[string]interface{}  `json:"value"`
         
    }
    
    // PageRequest ...
    type PageRequest struct {

        
            Schedule CronSchedule  `json:"_schedule"`
            Application string  `json:"application"`
            Author Author  `json:"author"`
            CustomJson *map[string]interface{}  `json:"_custom_json"`
            Orientation string  `json:"orientation"`
            Content []ResourceContent  `json:"content"`
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
            Duration int  `json:"duration"`
         
    }
    
    // PagePublishRequest ...
    type PagePublishRequest struct {

        
            Publish bool  `json:"publish"`
         
    }
    
    // PageMetaSchema ...
    type PageMetaSchema struct {

        
            SystemPages []*map[string]interface{}  `json:"system_pages"`
            CustomPages []*map[string]interface{}  `json:"custom_pages"`
            ApplicationID string  `json:"application_id"`
         
    }
    
    // SlideshowGetResponse ...
    type SlideshowGetResponse struct {

        
            Items []*map[string]interface{}  `json:"items"`
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
            CustomJson *map[string]interface{}  `json:"_custom_json"`
         
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
            Attributes *map[string]interface{}  `json:"attributes"`
            Content string  `json:"content"`
         
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
            Count int  `json:"count"`
         
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
            UpdatedAt string  `json:"updated_at"`
            Personalized bool  `json:"personalized"`
            Campaign CampaignShortLink  `json:"campaign"`
            Redirects Redirects  `json:"redirects"`
            Attribution Attribution  `json:"attribution"`
            SocialMediaTags SocialMediaTags  `json:"social_media_tags"`
            Count int  `json:"count"`
         
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

        
            Expiry int  `json:"expiry"`
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
            Size int  `json:"size"`
            Upload Upload  `json:"upload"`
            Cdn CDN  `json:"cdn"`
         
    }
    
    // StartRequest ...
    type StartRequest struct {

        
            FileName string  `json:"file_name"`
            ContentType string  `json:"content_type"`
            Size int  `json:"size"`
            Tags []string  `json:"tags"`
         
    }
    
    // CompleteResponse ...
    type CompleteResponse struct {

        
            ID string  `json:"_id"`
            FileName string  `json:"file_name"`
            FilePath string  `json:"file_path"`
            ContentType string  `json:"content_type"`
            Method string  `json:"method"`
            Namespace string  `json:"namespace"`
            Operation string  `json:"operation"`
            Size int  `json:"size"`
            Upload Upload  `json:"upload"`
            Cdn CDN  `json:"cdn"`
            Success string  `json:"success"`
            Tags []string  `json:"tags"`
            CreatedOn string  `json:"created_on"`
            ModifiedOn string  `json:"modified_on"`
         
    }
    
    // Opts ...
    type Opts struct {

        
            Attempts int  `json:"attempts"`
            Timestamp int  `json:"timestamp"`
            Delay int  `json:"delay"`
         
    }
    
    // CopyFileTask ...
    type CopyFileTask struct {

        
            ID string  `json:"id"`
            Name string  `json:"name"`
            Data BulkRequest  `json:"data"`
            Opts Opts  `json:"opts"`
            Progress int  `json:"progress"`
            Delay int  `json:"delay"`
            Timestamp int  `json:"timestamp"`
            AttemptsMade int  `json:"attempts_made"`
            Stacktrace []string  `json:"stacktrace"`
            FinishedOn int  `json:"finished_on"`
            ProcessedOn int  `json:"processed_on"`
         
    }
    
    // BulkResponse ...
    type BulkResponse struct {

        
            TrackingURL string  `json:"tracking_url"`
            Task CopyFileTask  `json:"task"`
         
    }
    
    // ReqConfiguration ...
    type ReqConfiguration struct {

        
            Concurrency int  `json:"concurrency"`
         
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
            Expiry int  `json:"expiry"`
         
    }
    
    // SignUrlResponse ...
    type SignUrlResponse struct {

        
            Urls []Urls  `json:"urls"`
         
    }
    
    // SignUrlRequest ...
    type SignUrlRequest struct {

        
            Expiry int  `json:"expiry"`
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
            UID int  `json:"uid"`
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

        
            Build int  `json:"build"`
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
    
    // UpdateDialog ...
    type UpdateDialog struct {

        
            Type string  `json:"type"`
            Interval int  `json:"interval"`
         
    }
    
    // OrderingStoreSelectRequest ...
    type OrderingStoreSelectRequest struct {

        
            OrderingStore OrderingStoreSelect  `json:"ordering_store"`
         
    }
    
    // OrderingStoreSelect ...
    type OrderingStoreSelect struct {

        
            UID int  `json:"uid"`
         
    }
    
    // AppStaff ...
    type AppStaff struct {

        
            ID string  `json:"_id"`
            OrderIncent bool  `json:"order_incent"`
            Stores []int  `json:"stores"`
            Application string  `json:"application"`
            Title string  `json:"title"`
            User string  `json:"user"`
            EmployeeCode string  `json:"employee_code"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
            ProfilePicURL string  `json:"profile_pic_url"`
         
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
            Params *map[string]interface{}  `json:"params"`
            Query *map[string]interface{}  `json:"query"`
         
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
            V int  `json:"__v"`
         
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
            ListingPage ListingPageFeature  `json:"listing_page"`
            Currency CurrencyFeature  `json:"currency"`
            RevenueEngine RevenueEngineFeature  `json:"revenue_engine"`
            Feedback FeedbackFeature  `json:"feedback"`
            CompareProducts CompareProductsFeature  `json:"compare_products"`
         
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
    
    // ListingPageFeature ...
    type ListingPageFeature struct {

        
            SortOn string  `json:"sort_on"`
         
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
            DecimalDigits int  `json:"decimal_digits"`
            Symbol string  `json:"symbol"`
         
    }
    
    // Domain ...
    type Domain struct {

        
            Verified bool  `json:"verified"`
            IsPrimary bool  `json:"is_primary"`
            IsDefault bool  `json:"is_default"`
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

        
            From string  `json:"from"`
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
            CacheTtl int  `json:"cache_ttl"`
            IsInternal bool  `json:"is_internal"`
            IsActive bool  `json:"is_active"`
            ID string  `json:"_id"`
            Name string  `json:"name"`
            Owner string  `json:"owner"`
            CompanyID int  `json:"company_id"`
            Token string  `json:"token"`
            Redirections []ApplicationRedirections  `json:"redirections"`
            Meta []ApplicationMeta  `json:"meta"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V int  `json:"__v"`
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
            Brands []int  `json:"brands"`
         
    }
    
    // StoreCriteriaRule ...
    type StoreCriteriaRule struct {

        
            Companies []int  `json:"companies"`
            Brands []int  `json:"brands"`
         
    }
    
    // InventoryStoreRule ...
    type InventoryStoreRule struct {

        
            Criteria string  `json:"criteria"`
            Rules []StoreCriteriaRule  `json:"rules"`
            Stores []int  `json:"stores"`
         
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
            EnforcedStores []int  `json:"enforced_stores"`
            Rules ArticleAssignmentRule  `json:"rules"`
         
    }
    
    // CompanyAboutAddress ...
    type CompanyAboutAddress struct {

        
            Pincode int  `json:"pincode"`
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
            CountryCode int  `json:"country_code"`
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
            V int  `json:"__v"`
         
    }
    
    // InformationAddress ...
    type InformationAddress struct {

        
            Loc string  `json:"loc"`
            AddressLine []string  `json:"address_line"`
            Phone InformationPhone  `json:"phone"`
            City string  `json:"city"`
            Country string  `json:"country"`
            Pincode int  `json:"pincode"`
         
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

        
            Facebook Facebook  `json:"facebook"`
            Instagram Instagram  `json:"instagram"`
            Twitter Twitter  `json:"twitter"`
            Pinterest Pinterest  `json:"pinterest"`
            GooglePlus GooglePlus  `json:"google_plus"`
            Youtube Youtube  `json:"youtube"`
            LinkedIn LinkedIn  `json:"linked_in"`
            Vimeo Vimeo  `json:"vimeo"`
            BlogLink BlogLink  `json:"blog_link"`
         
    }
    
    // Instagram ...
    type Instagram struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // Twitter ...
    type Twitter struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // Pinterest ...
    type Pinterest struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // GooglePlus ...
    type GooglePlus struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // Youtube ...
    type Youtube struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // LinkedIn ...
    type LinkedIn struct {

        
            Title string  `json:"title"`
            Icon string  `json:"icon"`
            Link string  `json:"link"`
         
    }
    
    // Vimeo ...
    type Vimeo struct {

        
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
    
    // StoreLatLong ...
    type StoreLatLong struct {

        
            Type string  `json:"type"`
            Coordinates []int  `json:"coordinates"`
         
    }
    
    // OptedStoreAddress ...
    type OptedStoreAddress struct {

        
            State string  `json:"state"`
            Address1 string  `json:"address1"`
            LatLong StoreLatLong  `json:"lat_long"`
            Address2 string  `json:"address2"`
            Pincode int  `json:"pincode"`
            Country string  `json:"country"`
            City string  `json:"city"`
         
    }
    
    // OrderingStore ...
    type OrderingStore struct {

        
            Address OptedStoreAddress  `json:"address"`
            ID string  `json:"_id"`
            UID int  `json:"uid"`
            Name string  `json:"name"`
            DisplayName string  `json:"display_name"`
            StoreType string  `json:"store_type"`
            StoreCode string  `json:"store_code"`
            Pincode int  `json:"pincode"`
            Code string  `json:"code"`
         
    }
    
    // OrderingStores ...
    type OrderingStores struct {

        
            Page Page  `json:"page"`
            Items []OrderingStore  `json:"items"`
            DeployedStores []int  `json:"deployed_stores"`
            AllStores bool  `json:"all_stores"`
            Enabled bool  `json:"enabled"`
            Type string  `json:"type"`
            ID string  `json:"_id"`
            App string  `json:"app"`
            V int  `json:"__v"`
         
    }
    
    // TokenResponse ...
    type TokenResponse struct {

        
            Tokens Tokens  `json:"tokens"`
            ID string  `json:"_id"`
            Application string  `json:"application"`
            CreatedAt string  `json:"created_at"`
            UpdatedAt string  `json:"updated_at"`
            V int  `json:"__v"`
         
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
            Auth Auth  `json:"auth"`
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
    
    // Auth ...
    type Auth struct {

        
            Google Google  `json:"google"`
            Facebook Facebook  `json:"facebook"`
            Accountkit Accountkit  `json:"accountkit"`
         
    }
    
    // GoogleMap ...
    type GoogleMap struct {

        
            Credentials GoogleMapCredentials  `json:"credentials"`
         
    }
    
    // GoogleMapCredentials ...
    type GoogleMapCredentials struct {

        
            APIKey string  `json:"api_key"`
         
    }
    

    
    // AggregatorConfigDetail ...
    type AggregatorConfigDetail struct {

        
            ConfigType string  `json:"config_type"`
            Key string  `json:"key"`
            MerchantKey string  `json:"merchant_key"`
            Sdk bool  `json:"sdk"`
            Pin string  `json:"pin"`
            Secret string  `json:"secret"`
            API string  `json:"api"`
            MerchantID string  `json:"merchant_id"`
            UserID string  `json:"user_id"`
            VerifyAPI string  `json:"verify_api"`
         
    }
    
    // AggregatorsConfigDetailResponse ...
    type AggregatorsConfigDetailResponse struct {

        
            Rupifi AggregatorConfigDetail  `json:"rupifi"`
            Ccavenue AggregatorConfigDetail  `json:"ccavenue"`
            Success bool  `json:"success"`
            Simpl AggregatorConfigDetail  `json:"simpl"`
            Juspay AggregatorConfigDetail  `json:"juspay"`
            Razorpay AggregatorConfigDetail  `json:"razorpay"`
            Mswipe AggregatorConfigDetail  `json:"mswipe"`
            Stripe AggregatorConfigDetail  `json:"stripe"`
            Payumoney AggregatorConfigDetail  `json:"payumoney"`
            Env string  `json:"env"`
         
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

        
            CardID string  `json:"card_id"`
            Refresh bool  `json:"refresh"`
         
    }
    
    // AttachCardsResponse ...
    type AttachCardsResponse struct {

        
            Data *map[string]interface{}  `json:"data"`
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

        
            Message string  `json:"message"`
            Cards CardPaymentGateway  `json:"cards"`
            Success bool  `json:"success"`
         
    }
    
    // Card ...
    type Card struct {

        
            ExpYear int  `json:"exp_year"`
            ExpMonth int  `json:"exp_month"`
            CardFingerprint string  `json:"card_fingerprint"`
            CardBrandImage string  `json:"card_brand_image"`
            CardType string  `json:"card_type"`
            CardID string  `json:"card_id"`
            Nickname string  `json:"nickname"`
            CardName string  `json:"card_name"`
            CardToken string  `json:"card_token"`
            Expired bool  `json:"expired"`
            CardBrand string  `json:"card_brand"`
            CardNumber string  `json:"card_number"`
            CardIsin string  `json:"card_isin"`
            AggregatorName string  `json:"aggregator_name"`
            CardReference string  `json:"card_reference"`
            CardIssuer string  `json:"card_issuer"`
         
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

        
            TransactionAmountInPaise int  `json:"transaction_amount_in_paise"`
            MerchantParams *map[string]interface{}  `json:"merchant_params"`
            Aggregator string  `json:"aggregator"`
            Payload string  `json:"payload"`
            PhoneNumber string  `json:"phone_number"`
         
    }
    
    // ValidateCustomerResponse ...
    type ValidateCustomerResponse struct {

        
            Data *map[string]interface{}  `json:"data"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // ChargeCustomerRequest ...
    type ChargeCustomerRequest struct {

        
            TransactionToken string  `json:"transaction_token"`
            Aggregator string  `json:"aggregator"`
            Verified bool  `json:"verified"`
            Amount int  `json:"amount"`
            OrderID string  `json:"order_id"`
         
    }
    
    // ChargeCustomerResponse ...
    type ChargeCustomerResponse struct {

        
            Success bool  `json:"success"`
            DeliveryAddressID string  `json:"delivery_address_id"`
            Aggregator string  `json:"aggregator"`
            Message string  `json:"message"`
            CartID string  `json:"cart_id"`
            OrderID string  `json:"order_id"`
            Status string  `json:"status"`
         
    }
    
    // PaymentInitializationRequest ...
    type PaymentInitializationRequest struct {

        
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Aggregator string  `json:"aggregator"`
            MerchantOrderID string  `json:"merchant_order_id"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            Timeout int  `json:"timeout"`
            Method string  `json:"method"`
            VirtualID string  `json:"virtual_id"`
            CustomerID string  `json:"customer_id"`
            PollingURL string  `json:"polling_url"`
         
    }
    
    // PaymentInitializationResponse ...
    type PaymentInitializationResponse struct {

        
            Currency string  `json:"currency"`
            Success bool  `json:"success"`
            CustomerID string  `json:"customer_id"`
            RazorpayPaymentID string  `json:"razorpay_payment_id"`
            Aggregator string  `json:"aggregator"`
            AggregatorOrderID string  `json:"aggregator_order_id"`
            Method string  `json:"method"`
            VirtualID string  `json:"virtual_id"`
            Timeout int  `json:"timeout"`
            BqrImage string  `json:"bqr_image"`
            Vpa string  `json:"vpa"`
            Amount int  `json:"amount"`
            MerchantOrderID string  `json:"merchant_order_id"`
            PollingURL string  `json:"polling_url"`
            Status string  `json:"status"`
         
    }
    
    // PaymentStatusUpdateRequest ...
    type PaymentStatusUpdateRequest struct {

        
            Currency string  `json:"currency"`
            Aggregator string  `json:"aggregator"`
            Contact string  `json:"contact"`
            MerchantOrderID string  `json:"merchant_order_id"`
            Method string  `json:"method"`
            Email string  `json:"email"`
            Vpa string  `json:"vpa"`
            Amount int  `json:"amount"`
            CustomerID string  `json:"customer_id"`
            OrderID string  `json:"order_id"`
            Status string  `json:"status"`
         
    }
    
    // PaymentStatusUpdateResponse ...
    type PaymentStatusUpdateResponse struct {

        
            AggregatorName string  `json:"aggregator_name"`
            Retry bool  `json:"retry"`
            Status string  `json:"status"`
         
    }
    
    // AggregatorRoute ...
    type AggregatorRoute struct {

        
            APILink string  `json:"api_link"`
            Data *map[string]interface{}  `json:"data"`
            PaymentFlow string  `json:"payment_flow"`
         
    }
    
    // PaymentFlow ...
    type PaymentFlow struct {

        
            Rupifi AggregatorRoute  `json:"rupifi"`
            Ccavenue AggregatorRoute  `json:"ccavenue"`
            Simpl AggregatorRoute  `json:"simpl"`
            BqrRazorpay AggregatorRoute  `json:"bqr_razorpay"`
            Razorpay AggregatorRoute  `json:"razorpay"`
            Juspay AggregatorRoute  `json:"juspay"`
            Fynd AggregatorRoute  `json:"fynd"`
            Stripe AggregatorRoute  `json:"stripe"`
            Payubiz AggregatorRoute  `json:"payubiz"`
            UpiRazorpay AggregatorRoute  `json:"upi_razorpay"`
            Msipe AggregatorRoute  `json:"msipe"`
         
    }
    
    // PaymentModeLogo ...
    type PaymentModeLogo struct {

        
            Large string  `json:"large"`
            Small string  `json:"small"`
         
    }
    
    // PaymentModeList ...
    type PaymentModeList struct {

        
            ExpMonth int  `json:"exp_month"`
            IntentAppErrorList []string  `json:"intent_app_error_list"`
            CardType string  `json:"card_type"`
            CardID string  `json:"card_id"`
            IntentFlow string  `json:"intent_flow"`
            CardToken string  `json:"card_token"`
            Expired bool  `json:"expired"`
            Name string  `json:"name"`
            CardBrand string  `json:"card_brand"`
            CardReference string  `json:"card_reference"`
            ExpYear int  `json:"exp_year"`
            FyndVpa string  `json:"fynd_vpa"`
            Code string  `json:"code"`
            Timeout int  `json:"timeout"`
            RetryCount int  `json:"retry_count"`
            LogoURL PaymentModeLogo  `json:"logo_url"`
            CardNumber string  `json:"card_number"`
            MerchantCode string  `json:"merchant_code"`
            DisplayPriority int  `json:"display_priority"`
            CardBrandImage string  `json:"card_brand_image"`
            DisplayName string  `json:"display_name"`
            CardName string  `json:"card_name"`
            CardFingerprint string  `json:"card_fingerprint"`
            CardIsin string  `json:"card_isin"`
            AggregatorName string  `json:"aggregator_name"`
            Nickname string  `json:"nickname"`
            CardIssuer string  `json:"card_issuer"`
         
    }
    
    // RootPaymentMode ...
    type RootPaymentMode struct {

        
            DisplayPriority int  `json:"display_priority"`
            AnonymousEnable bool  `json:"anonymous_enable"`
            Name string  `json:"name"`
            DisplayName string  `json:"display_name"`
            AddCardEnabled bool  `json:"add_card_enabled"`
            List []PaymentModeList  `json:"list"`
            AggregatorName string  `json:"aggregator_name"`
         
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
    
    // TransferItemsDetails ...
    type TransferItemsDetails struct {

        
            LogoSmall string  `json:"logo_small"`
            Name string  `json:"name"`
            DisplayName bool  `json:"display_name"`
            ID string  `json:"id"`
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

        
            TransferMode string  `json:"transfer_mode"`
            Enable bool  `json:"enable"`
         
    }
    
    // UpdateRefundTransferModeResponse ...
    type UpdateRefundTransferModeResponse struct {

        
            Success bool  `json:"success"`
         
    }
    
    // OrderBeneficiaryDetails ...
    type OrderBeneficiaryDetails struct {

        
            BankName string  `json:"bank_name"`
            AccountNo string  `json:"account_no"`
            Title string  `json:"title"`
            Subtitle string  `json:"subtitle"`
            IfscCode string  `json:"ifsc_code"`
            CreatedOn string  `json:"created_on"`
            DelightsUserName string  `json:"delights_user_name"`
            Mobile bool  `json:"mobile"`
            BranchName bool  `json:"branch_name"`
            IsActive bool  `json:"is_active"`
            ModifiedOn string  `json:"modified_on"`
            Comment bool  `json:"comment"`
            Email string  `json:"email"`
            AccountHolder string  `json:"account_holder"`
            Address string  `json:"address"`
            BeneficiaryID string  `json:"beneficiary_id"`
            DisplayName string  `json:"display_name"`
            ID int  `json:"id"`
            TransferMode string  `json:"transfer_mode"`
         
    }
    
    // OrderBeneficiaryResponse ...
    type OrderBeneficiaryResponse struct {

        
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
            BranchName string  `json:"branch_name"`
            Success bool  `json:"success"`
         
    }
    
    // ErrorCodeDescription ...
    type ErrorCodeDescription struct {

        
            Code string  `json:"code"`
            Description string  `json:"description"`
            Success bool  `json:"success"`
         
    }
    
    // AddBeneficiaryViaOtpVerificationRequest ...
    type AddBeneficiaryViaOtpVerificationRequest struct {

        
            RequestID string  `json:"request_id"`
            HashKey string  `json:"hash_key"`
            Otp string  `json:"otp"`
         
    }
    
    // AddBeneficiaryViaOtpVerificationResponse ...
    type AddBeneficiaryViaOtpVerificationResponse struct {

        
            RequestID string  `json:"request_id"`
            HashKey string  `json:"hash_key"`
            Otp string  `json:"otp"`
         
    }
    
    // WrongOtpError ...
    type WrongOtpError struct {

        
            Description string  `json:"description"`
            Success string  `json:"success"`
         
    }
    
    // BankDetails ...
    type BankDetails struct {

        
            Address string  `json:"address"`
            IfscCode string  `json:"ifsc_code"`
            Comment string  `json:"comment"`
            AccountNo string  `json:"account_no"`
            Email string  `json:"email"`
            AccountHolder string  `json:"account_holder"`
            BankName string  `json:"bank_name"`
            Mobile string  `json:"mobile"`
            BranchName string  `json:"branch_name"`
         
    }
    
    // AddBeneficiaryDetailsRequest ...
    type AddBeneficiaryDetailsRequest struct {

        
            Details BankDetails  `json:"details"`
            Delights bool  `json:"delights"`
            TransferMode string  `json:"transfer_mode"`
            ShipmentID string  `json:"shipment_id"`
            OrderID string  `json:"order_id"`
         
    }
    
    // RefundAccountResponse ...
    type RefundAccountResponse struct {

        
            Data *map[string]interface{}  `json:"data"`
            Message string  `json:"message"`
            Success bool  `json:"success"`
         
    }
    
    // WalletOtpRequest ...
    type WalletOtpRequest struct {

        
            CountryCode string  `json:"country_code"`
            Mobile bool  `json:"mobile"`
         
    }
    
    // WalletOtpResponse ...
    type WalletOtpResponse struct {

        
            IsVerifiedFlag string  `json:"is_verified_flag"`
            RequestID string  `json:"request_id"`
            Success bool  `json:"success"`
         
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
    

    
    // OrderById ...
    type OrderById struct {

        
            Order OrderSchema  `json:"order"`
         
    }
    
    // OrderList ...
    type OrderList struct {

        
            Items []OrderSchema  `json:"items"`
            Page OrderPage  `json:"page"`
            Filters OrderFilters  `json:"filters"`
         
    }
    
    // OrderPage ...
    type OrderPage struct {

        
            ItemTotal int  `json:"item_total"`
            Type string  `json:"type"`
            Size int  `json:"size"`
            Current int  `json:"current"`
            HasNext bool  `json:"has_next"`
         
    }
    
    // OrderFilters ...
    type OrderFilters struct {

        
            Statuses []OrderStatuses  `json:"statuses"`
            Months []OrderMonths  `json:"months"`
         
    }
    
    // OrderStatuses ...
    type OrderStatuses struct {

        
            Display string  `json:"display"`
            Value int  `json:"value"`
            IsSelected bool  `json:"is_selected"`
         
    }
    
    // OrderMonths ...
    type OrderMonths struct {

        
            FromDate string  `json:"from_date"`
            ID string  `json:"id"`
            IsSelected bool  `json:"is_selected"`
            Name string  `json:"name"`
            ToDate string  `json:"to_date"`
            Value int  `json:"value"`
         
    }
    
    // ShipmentById ...
    type ShipmentById struct {

        
            Shipment Shipments  `json:"shipment"`
         
    }
    
    // ShipmentReasons ...
    type ShipmentReasons struct {

        
            Reasons []Reasons  `json:"reasons"`
         
    }
    
    // ShipmentStatusUpdateBody ...
    type ShipmentStatusUpdateBody struct {

        
            Statuses []StatusesBody  `json:"statuses"`
            ForceTransition bool  `json:"force_transition"`
         
    }
    
    // StatusesBody ...
    type StatusesBody struct {

        
            Status string  `json:"status"`
            Shipments *map[string]interface{}  `json:"shipments"`
         
    }
    
    // ShipmentStatusUpdate ...
    type ShipmentStatusUpdate struct {

        
            Message []*map[string]interface{}  `json:"message"`
            Status bool  `json:"status"`
         
    }
    
    // ShipmentTrack ...
    type ShipmentTrack struct {

        
            Results []Track  `json:"results"`
         
    }
    
    // OrderSchema ...
    type OrderSchema struct {

        
            OrderID string  `json:"order_id"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            OrderCreatedTime string  `json:"order_created_time"`
            Shipments []Shipments  `json:"shipments"`
            TotalShipmentsInOrder int  `json:"total_shipments_in_order"`
            UserInfo UserInfo  `json:"user_info"`
         
    }
    
    // PosOrderById ...
    type PosOrderById struct {

        
            Order OrderSchema  `json:"order"`
         
    }
    
    // Bags ...
    type Bags struct {

        
            Item Item  `json:"item"`
            Prices Prices  `json:"prices"`
            CurrentStatus CurrentStatus  `json:"current_status"`
            ID int  `json:"id"`
            FinancialBreakup []FinancialBreakup  `json:"financial_breakup"`
         
    }
    
    // Item ...
    type Item struct {

        
            Brand ItemBrand  `json:"brand"`
            Name string  `json:"name"`
            Size string  `json:"size"`
            SlugKey string  `json:"slug_key"`
            Image []string  `json:"image"`
            Code string  `json:"code"`
         
    }
    
    // Prices ...
    type Prices struct {

        
            AmountPaidRoundoff int  `json:"amount_paid_roundoff"`
            FyndCredits int  `json:"fynd_credits"`
            CodCharges int  `json:"cod_charges"`
            Cashback int  `json:"cashback"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            PriceMarked int  `json:"price_marked"`
            TransferPrice int  `json:"transfer_price"`
            CouponValue int  `json:"coupon_value"`
            PriceEffective int  `json:"price_effective"`
            RefundCredit int  `json:"refund_credit"`
            AmountPaid int  `json:"amount_paid"`
            RefundAmount int  `json:"refund_amount"`
            CashbackApplied int  `json:"cashback_applied"`
            GstTaxPercentage int  `json:"gst_tax_percentage"`
            ValueOfGood int  `json:"value_of_good"`
            BrandCalculatedAmount int  `json:"brand_calculated_amount"`
            PromotionEffectiveDiscount int  `json:"promotion_effective_discount"`
            Discount int  `json:"discount"`
            CouponEffectiveDiscount int  `json:"coupon_effective_discount"`
            DeliveryCharge int  `json:"delivery_charge"`
         
    }
    
    // CurrentStatus ...
    type CurrentStatus struct {

        
            UpdatedAt string  `json:"updated_at"`
            Status string  `json:"status"`
            Name string  `json:"name"`
            JourneyType string  `json:"journey_type"`
         
    }
    
    // FinancialBreakup ...
    type FinancialBreakup struct {

        
            BrandCalculatedAmount int  `json:"brand_calculated_amount"`
            CouponValue int  `json:"coupon_value"`
            AmountPaidRoundoff int  `json:"amount_paid_roundoff"`
            GstFee string  `json:"gst_fee"`
            RefundCredit int  `json:"refund_credit"`
            Cashback int  `json:"cashback"`
            RefundAmount int  `json:"refund_amount"`
            ValueOfGood int  `json:"value_of_good"`
            PromotionEffectiveDiscount int  `json:"promotion_effective_discount"`
            Size string  `json:"size"`
            TotalUnits int  `json:"total_units"`
            Discount int  `json:"discount"`
            AmountPaid int  `json:"amount_paid"`
            FyndCredits int  `json:"fynd_credits"`
            AddedToFyndCash bool  `json:"added_to_fynd_cash"`
            DeliveryCharge int  `json:"delivery_charge"`
            HsnCode string  `json:"hsn_code"`
            CouponEffectiveDiscount int  `json:"coupon_effective_discount"`
            TransferPrice int  `json:"transfer_price"`
            Identifiers Identifiers  `json:"identifiers"`
            GstTag string  `json:"gst_tag"`
            PriceMarked int  `json:"price_marked"`
            PriceEffective int  `json:"price_effective"`
            CodCharges int  `json:"cod_charges"`
            ItemName string  `json:"item_name"`
            CashbackApplied int  `json:"cashback_applied"`
            GstTaxPercentage int  `json:"gst_tax_percentage"`
         
    }
    
    // Identifiers ...
    type Identifiers struct {

        
            Ean string  `json:"ean"`
            SkuCode string  `json:"sku_code"`
         
    }
    
    // ItemBrand ...
    type ItemBrand struct {

        
            Name string  `json:"name"`
            Logo string  `json:"logo"`
         
    }
    
    // BreakupValues ...
    type BreakupValues struct {

        
            Display string  `json:"display"`
            Value int  `json:"value"`
            Name string  `json:"name"`
         
    }
    
    // DeliveryAddress ...
    type DeliveryAddress struct {

        
            Pincode string  `json:"pincode"`
            Landmark string  `json:"landmark"`
            ContactPerson string  `json:"contact_person"`
            Phone string  `json:"phone"`
            State string  `json:"state"`
            Version string  `json:"version"`
            Address1 string  `json:"address1"`
            CreatedAt string  `json:"created_at"`
            AddressType string  `json:"address_type"`
            AddressCategory string  `json:"address_category"`
            Area string  `json:"area"`
            City string  `json:"city"`
            Latitude int  `json:"latitude"`
            Longitude int  `json:"longitude"`
            Email string  `json:"email"`
            Country string  `json:"country"`
            Address2 string  `json:"address2"`
            UpdatedAt string  `json:"updated_at"`
            Name string  `json:"name"`
            Address string  `json:"address"`
         
    }
    
    // FulfillingStore ...
    type FulfillingStore struct {

        
            Code string  `json:"code"`
            ID int  `json:"id"`
            Name string  `json:"name"`
            CompanyID int  `json:"company_id"`
         
    }
    
    // Invoice ...
    type Invoice struct {

        
            UpdatedDate string  `json:"updated_date"`
            InvoiceURL string  `json:"invoice_url"`
            LabelURL string  `json:"label_url"`
         
    }
    
    // Promise ...
    type Promise struct {

        
            Timestamp Timestamp  `json:"timestamp"`
         
    }
    
    // Timestamp ...
    type Timestamp struct {

        
            Min string  `json:"min"`
            Max string  `json:"max"`
         
    }
    
    // Reasons ...
    type Reasons struct {

        
            ReasonText string  `json:"reason_text"`
            ShowTextArea bool  `json:"show_text_area"`
            FeedbackType string  `json:"feedback_type"`
            Flow string  `json:"flow"`
            ReasonID int  `json:"reason_id"`
            Priority int  `json:"priority"`
         
    }
    
    // ShipmentStatus ...
    type ShipmentStatus struct {

        
            Title string  `json:"title"`
            HexCode string  `json:"hex_code"`
         
    }
    
    // ShipmentUserInfo ...
    type ShipmentUserInfo struct {

        
            Gender string  `json:"gender"`
            Mobile string  `json:"mobile"`
            FirstName string  `json:"first_name"`
            LastName string  `json:"last_name"`
         
    }
    
    // Shipments ...
    type Shipments struct {

        
            OrderID string  `json:"order_id"`
            BreakupValues []BreakupValues  `json:"breakup_values"`
            TrackURL string  `json:"track_url"`
            TrakingNo string  `json:"traking_no"`
            TrackingDetails []TrackingDetails  `json:"tracking_details"`
            BeneficiaryDetails bool  `json:"beneficiary_details"`
            CanReturn bool  `json:"can_return"`
            Prices Prices  `json:"prices"`
            NeedHelpURL string  `json:"need_help_url"`
            ShipmentID string  `json:"shipment_id"`
            TotalBags int  `json:"total_bags"`
            DeliveryAddress DeliveryAddress  `json:"delivery_address"`
            Invoice Invoice  `json:"invoice"`
            Comment string  `json:"comment"`
            OrderType string  `json:"order_type"`
            Promise Promise  `json:"promise"`
            FulfillingStore FulfillingStore  `json:"fulfilling_store"`
            Bags []Bags  `json:"bags"`
            CanCancel bool  `json:"can_cancel"`
            Payment ShipmentPayment  `json:"payment"`
            ShipmentCreatedAt string  `json:"shipment_created_at"`
            ShipmentStatus ShipmentStatus  `json:"shipment_status"`
            UserInfo ShipmentUserInfo  `json:"user_info"`
            SizeInfo *map[string]interface{}  `json:"size_info"`
            TotalDetails ShipmentTotalDetails  `json:"total_details"`
         
    }
    
    // ShipmentTotalDetails ...
    type ShipmentTotalDetails struct {

        
            TotalPrice int  `json:"total_price"`
            Sizes int  `json:"sizes"`
            Pieces int  `json:"pieces"`
         
    }
    
    // ShipmentPayment ...
    type ShipmentPayment struct {

        
            Logo string  `json:"logo"`
            Mode string  `json:"mode"`
            Status string  `json:"status"`
         
    }
    
    // Track ...
    type Track struct {

        
            Awb string  `json:"awb"`
            UpdatedAt string  `json:"updated_at"`
            LastLocationRecievedAt string  `json:"last_location_recieved_at"`
            Reason string  `json:"reason"`
            ShipmentType string  `json:"shipment_type"`
            Status string  `json:"status"`
            UpdatedTime string  `json:"updated_time"`
            AccountName string  `json:"account_name"`
         
    }
    
    // TrackingDetails ...
    type TrackingDetails struct {

        
            IsCurrent bool  `json:"is_current"`
            Status string  `json:"status"`
            Time string  `json:"time"`
            IsPassed string  `json:"is_passed"`
         
    }
    
    // UserInfo ...
    type UserInfo struct {

        
            Gender string  `json:"gender"`
            Mobile string  `json:"mobile"`
            Name string  `json:"name"`
            Email string  `json:"email"`
         
    }
    
    // ApefaceApiError ...
    type ApefaceApiError struct {

        
            Message string  `json:"message"`
         
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

        
            Absolute int  `json:"absolute"`
            Currency string  `json:"currency"`
            DisplayAbsolute string  `json:"display_absolute"`
            DisplayPercent string  `json:"display_percent"`
            Percent int  `json:"percent"`
         
    }
    
    // Error ...
    type Error struct {

        
            Code int  `json:"code"`
            Exception string  `json:"exception"`
            Info string  `json:"info"`
            Message string  `json:"message"`
         
    }
    
    // Offer ...
    type Offer struct {

        
            Schedule string  `json:"_schedule"`
            Active bool  `json:"active"`
            ApplicationID string  `json:"application_id"`
            BannerImage Asset  `json:"banner_image"`
            CreatedAt string  `json:"created_at"`
            Name string  `json:"name"`
            Rule *map[string]interface{}  `json:"rule"`
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
            OrderAmount int  `json:"order_amount"`
         
    }
    
    // OrderDiscountResponse ...
    type OrderDiscountResponse struct {

        
            AppliedRuleBucket OrderDiscountRuleBucket  `json:"applied_rule_bucket"`
            BaseDiscount DiscountProperties  `json:"base_discount"`
            Discount DiscountProperties  `json:"discount"`
            OrderAmount int  `json:"order_amount"`
            Points int  `json:"points"`
         
    }
    
    // OrderDiscountRuleBucket ...
    type OrderDiscountRuleBucket struct {

        
            High int  `json:"high"`
            Low int  `json:"low"`
            Max int  `json:"max"`
            Value int  `json:"value"`
            ValueType string  `json:"value_type"`
         
    }
    
    // PointsHistory ...
    type PointsHistory struct {

        
            ID string  `json:"_id"`
            ApplicationID string  `json:"application_id"`
            Claimed bool  `json:"claimed"`
            CreatedAt string  `json:"created_at"`
            ExpiresOn string  `json:"expires_on"`
            Meta string  `json:"meta"`
            Points int  `json:"points"`
            RemainingPoints int  `json:"remaining_points"`
            Text1 string  `json:"text_1"`
            Text2 string  `json:"text_2"`
            Text3 string  `json:"text_3"`
            TxnName string  `json:"txn_name"`
            UpdatedAt string  `json:"updated_at"`
            UserID string  `json:"user_id"`
         
    }
    
    // PointsHistoryResponse ...
    type PointsHistoryResponse struct {

        
            History []PointsHistory  `json:"history"`
            Page Page  `json:"page"`
         
    }
    
    // PointsResponse ...
    type PointsResponse struct {

        
            Points int  `json:"points"`
         
    }
    
    // RedeemReferralCodeRequest ...
    type RedeemReferralCodeRequest struct {

        
            DeviceID string  `json:"device_id"`
            ReferralCode string  `json:"referral_code"`
         
    }
    
    // RedeemReferralCodeResponse ...
    type RedeemReferralCodeResponse struct {

        
            Message string  `json:"message"`
            Points int  `json:"points"`
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
            Points int  `json:"points"`
            Redeemed bool  `json:"redeemed"`
            ReferralCode string  `json:"referral_code"`
         
    }
    
    // RewardsArticle ...
    type RewardsArticle struct {

        
            ID string  `json:"id"`
            Points int  `json:"points"`
            Price int  `json:"price"`
         
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
    

    
    // AbuseReport ...
    type AbuseReport struct {

        
            Abused bool  `json:"abused"`
            DateMeta DateMeta  `json:"date_meta"`
            Description string  `json:"description"`
            Entity Entity  `json:"entity"`
            ID string  `json:"id"`
            Name string  `json:"name"`
            State FeedbackState  `json:"state"`
            Tags []TagMeta  `json:"tags"`
         
    }
    
    // Access ...
    type Access struct {

        
            Answer bool  `json:"answer"`
            AskQuestion bool  `json:"ask_question"`
            Comment bool  `json:"comment"`
            Rnr bool  `json:"rnr"`
         
    }
    
    // AddMediaListRequest ...
    type AddMediaListRequest struct {

        
            EntityID string  `json:"entity_id"`
            EntityType string  `json:"entity_type"`
            MediaList []AddMediaRequest  `json:"media_list"`
            RefID string  `json:"ref_id"`
            RefType string  `json:"ref_type"`
         
    }
    
    // AddMediaRequest ...
    type AddMediaRequest struct {

        
            CloudID string  `json:"cloud_id"`
            CloudName string  `json:"cloud_name"`
            CloudProvider string  `json:"cloud_provider"`
            EntityID string  `json:"entity_id"`
            EntityType string  `json:"entity_type"`
            MediaURL string  `json:"media_url"`
            RefID string  `json:"ref_id"`
            RefType string  `json:"ref_type"`
            Tags []string  `json:"tags"`
            ThumbnailURL string  `json:"thumbnail_url"`
            Type string  `json:"type"`
         
    }
    
    // Attribute ...
    type Attribute struct {

        
            DateMeta DateMeta  `json:"date_meta"`
            Description string  `json:"description"`
            ID string  `json:"id"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Tags []TagMeta  `json:"tags"`
         
    }
    
    // AttributeObject ...
    type AttributeObject struct {

        
            Description string  `json:"description"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
            Title string  `json:"title"`
            Type string  `json:"type"`
            Value int  `json:"value"`
         
    }
    
    // AttributeResponse ...
    type AttributeResponse struct {

        
            Items []Attribute  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CheckEligibilityResponse ...
    type CheckEligibilityResponse struct {

        
            Access Access  `json:"access"`
            MediaCloud MediaCloud  `json:"media_cloud"`
         
    }
    
    // Cloud ...
    type Cloud struct {

        
            ID string  `json:"id"`
            Name string  `json:"name"`
            Provider string  `json:"provider"`
         
    }
    
    // Comment ...
    type Comment struct {

        
            Comment []string  `json:"comment"`
            DateMeta DateMeta  `json:"date_meta"`
            Entity Entity  `json:"entity"`
            ID string  `json:"id"`
            Name string  `json:"name"`
            State FeedbackState  `json:"state"`
            Tags []TagMeta  `json:"tags"`
            VoteCount VoteCount  `json:"vote_count"`
         
    }
    
    // CommentGetResponse ...
    type CommentGetResponse struct {

        
            Items []Comment  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // CommentRequest ...
    type CommentRequest struct {

        
            Comment []string  `json:"comment"`
            EntityID string  `json:"entity_id"`
            EntityType string  `json:"entity_type"`
         
    }
    
    // CreateQNARequest ...
    type CreateQNARequest struct {

        
            Choices []string  `json:"choices"`
            EntityID string  `json:"entity_id"`
            EntityType string  `json:"entity_type"`
            MaxLen int  `json:"max_len"`
            SortPriority int  `json:"sort_priority"`
            Tags []string  `json:"tags"`
            Text string  `json:"text"`
            Type string  `json:"type"`
         
    }
    
    // CursorGetResponse ...
    type CursorGetResponse struct {

        
            Items []*map[string]interface{}  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // DeviceMeta ...
    type DeviceMeta struct {

        
            AppVersion string  `json:"app_version"`
            Platform string  `json:"platform"`
         
    }
    
    // Entity ...
    type Entity struct {

        
            ID string  `json:"id"`
            Type string  `json:"type"`
         
    }
    
    // FeedbackError ...
    type FeedbackError struct {

        
            Code string  `json:"code"`
            Exception string  `json:"exception"`
            Info string  `json:"info"`
            Message string  `json:"message"`
            Meta *map[string]interface{}  `json:"meta"`
            RequestID string  `json:"request_id"`
            StackTrace string  `json:"stack_trace"`
            Status int  `json:"status"`
         
    }
    
    // FeedbackMedia ...
    type FeedbackMedia struct {

        
            Cloud Cloud  `json:"cloud"`
            DateMeta DateMeta  `json:"date_meta"`
            Description string  `json:"description"`
            Entity Entity  `json:"entity"`
            ID string  `json:"id"`
            Name string  `json:"name"`
            Reference Entity  `json:"reference"`
            State MediaState  `json:"state"`
            Tags []TagMeta  `json:"tags"`
            Type string  `json:"type"`
            URL Url  `json:"url"`
         
    }
    
    // FeedbackState ...
    type FeedbackState struct {

        
            Active bool  `json:"active"`
            Approve bool  `json:"approve"`
         
    }
    
    // InsertResponse ...
    type InsertResponse struct {

        
            Ids string  `json:"ids"`
         
    }
    
    // MediaCloud ...
    type MediaCloud struct {

        
            Key string  `json:"key"`
            Name string  `json:"name"`
            Namespace string  `json:"namespace"`
            Path string  `json:"path"`
            Provider string  `json:"provider"`
            Secret string  `json:"secret"`
         
    }
    
    // MediaGetResponse ...
    type MediaGetResponse struct {

        
            Items []Comment  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // MediaMeta ...
    type MediaMeta struct {

        
            Cloud Cloud  `json:"cloud"`
            Comment []string  `json:"comment"`
            Description string  `json:"description"`
            ID string  `json:"id"`
            Type string  `json:"type"`
            URL Url  `json:"url"`
         
    }
    
    // MediaState ...
    type MediaState struct {

        
            Approve bool  `json:"approve"`
            Archive bool  `json:"archive"`
         
    }
    
    // NumberGetResponse ...
    type NumberGetResponse struct {

        
            Items []*map[string]interface{}  `json:"items"`
            Page PageNumber  `json:"page"`
         
    }
    
    // PageNumber ...
    type PageNumber struct {

        
            Current int  `json:"current"`
            HasNext bool  `json:"has_next"`
            ItemTotal int  `json:"item_total"`
            Size int  `json:"size"`
            Type string  `json:"type"`
         
    }
    
    // QNA ...
    type QNA struct {

        
            Comments []Comment  `json:"comments"`
            DateMeta DateMeta  `json:"date_meta"`
            Entity Entity  `json:"entity"`
            ID string  `json:"id"`
            Name string  `json:"name"`
            Question Question  `json:"question"`
            State QNAState  `json:"state"`
            Tag []string  `json:"tag"`
            Tags []TagMeta  `json:"tags"`
         
    }
    
    // QNAGetResponse ...
    type QNAGetResponse struct {

        
            Items []QNA  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // QNAState ...
    type QNAState struct {

        
            Active bool  `json:"active"`
            Approve bool  `json:"approve"`
            Modify bool  `json:"modify"`
            Priority int  `json:"priority"`
            Reply bool  `json:"reply"`
            Vote bool  `json:"vote"`
         
    }
    
    // Question ...
    type Question struct {

        
            Choices []string  `json:"choices"`
            MaxLen int  `json:"max_len"`
            Text string  `json:"text"`
            Type string  `json:"type"`
         
    }
    
    // Rating ...
    type Rating struct {

        
            Attributes []Attribute  `json:"attributes"`
            AttributesSlugs []string  `json:"attributes_slugs"`
            UI UI  `json:"ui"`
         
    }
    
    // RatingGetResponse ...
    type RatingGetResponse struct {

        
            Items []Rating  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ReportAbuseGetResponse ...
    type ReportAbuseGetResponse struct {

        
            Items []AbuseReport  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // ReportAbuseRequest ...
    type ReportAbuseRequest struct {

        
            Description string  `json:"description"`
            EntityID string  `json:"entity_id"`
            EntityType string  `json:"entity_type"`
         
    }
    
    // Review ...
    type Review struct {

        
            Description string  `json:"description"`
            Header string  `json:"header"`
            ImageMeta MediaMeta  `json:"image_meta"`
            Title string  `json:"title"`
            VideoMeta MediaMeta  `json:"video_meta"`
            VoteAllowed bool  `json:"vote_allowed"`
         
    }
    
    // ReviewGetResponse ...
    type ReviewGetResponse struct {

        
            Items []Review  `json:"items"`
            Page Page  `json:"page"`
         
    }
    
    // SaveAttributeRequest ...
    type SaveAttributeRequest struct {

        
            Description string  `json:"description"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
         
    }
    
    // TagMeta ...
    type TagMeta struct {

        
            Media []MediaMeta  `json:"media"`
            Name string  `json:"name"`
            Type string  `json:"type"`
         
    }
    
    // UI ...
    type UI struct {

        
            FeedbackQuestion []string  `json:"feedback_question"`
            Icon UIIcon  `json:"icon"`
            Text []string  `json:"text"`
            Type string  `json:"type"`
         
    }
    
    // UIIcon ...
    type UIIcon struct {

        
            Active string  `json:"active"`
            Inactive string  `json:"inactive"`
            Selected []string  `json:"selected"`
         
    }
    
    // UpdateAbuseStatusRequest ...
    type UpdateAbuseStatusRequest struct {

        
            Abusive bool  `json:"abusive"`
            Active bool  `json:"active"`
            Approve bool  `json:"approve"`
            Description string  `json:"description"`
            EntityID string  `json:"entity_id"`
            EntityType string  `json:"entity_type"`
            ID string  `json:"id"`
         
    }
    
    // UpdateAttributeRequest ...
    type UpdateAttributeRequest struct {

        
            Description string  `json:"description"`
            Name string  `json:"name"`
            Slug string  `json:"slug"`
         
    }
    
    // UpdateCommentRequest ...
    type UpdateCommentRequest struct {

        
            Active bool  `json:"active"`
            Approve bool  `json:"approve"`
            Comment []string  `json:"comment"`
            ID string  `json:"id"`
         
    }
    
    // UpdateMediaListRequest ...
    type UpdateMediaListRequest struct {

        
            Approve bool  `json:"approve"`
            Archive bool  `json:"archive"`
            EntityType string  `json:"entity_type"`
            Ids []string  `json:"ids"`
         
    }
    
    // UpdateQNARequest ...
    type UpdateQNARequest struct {

        
            Active bool  `json:"active"`
            Approve bool  `json:"approve"`
            Choices []string  `json:"choices"`
            ID string  `json:"id"`
            Tags []string  `json:"tags"`
         
    }
    
    // UpdateResponse ...
    type UpdateResponse struct {

        
            ID string  `json:"id"`
         
    }
    
    // UpdateReviewRequest ...
    type UpdateReviewRequest struct {

        
            Active bool  `json:"active"`
            Application string  `json:"application"`
            Approve bool  `json:"approve"`
            Archive bool  `json:"archive"`
            AttributesRating []AttributeObject  `json:"attributes_rating"`
            Description string  `json:"description"`
            DeviceMeta DeviceMeta  `json:"device_meta"`
            EntityID string  `json:"entity_id"`
            EntityType string  `json:"entity_type"`
            MediaResource []MediaMeta  `json:"media_resource"`
            Rating int  `json:"rating"`
            ReviewID string  `json:"review_id"`
            TemplateID string  `json:"template_id"`
            Title string  `json:"title"`
         
    }
    
    // UpdateVoteRequest ...
    type UpdateVoteRequest struct {

        
            Action string  `json:"action"`
            Active bool  `json:"active"`
            ID string  `json:"id"`
            RefID string  `json:"ref_id"`
            RefType string  `json:"ref_type"`
         
    }
    
    // Url ...
    type Url struct {

        
            Main string  `json:"main"`
            Thumbnail string  `json:"thumbnail"`
         
    }
    
    // Vote ...
    type Vote struct {

        
            Action string  `json:"action"`
            DateMeta DateMeta  `json:"date_meta"`
            Entity Entity  `json:"entity"`
            ID string  `json:"id"`
            Name string  `json:"name"`
            Reference Entity  `json:"reference"`
            State FeedbackState  `json:"state"`
            Tags []TagMeta  `json:"tags"`
         
    }
    
    // VoteCount ...
    type VoteCount struct {

        
            Downvote int  `json:"downvote"`
            Upvote int  `json:"upvote"`
         
    }
    
    // VoteRequest ...
    type VoteRequest struct {

        
            Action string  `json:"action"`
            EntityID string  `json:"entity_id"`
            EntityType string  `json:"entity_type"`
            RefID string  `json:"ref_id"`
            RefType string  `json:"ref_type"`
         
    }
    
    // VoteResponse ...
    type VoteResponse struct {

        
            Items []Vote  `json:"items"`
            Page Page  `json:"page"`
         
    }
    

    
    // UpdateCartShipmentItem ...
    type UpdateCartShipmentItem struct {

        
            Quantity int  `json:"quantity"`
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
    
    // CartPosCheckoutRequest ...
    type CartPosCheckoutRequest struct {

        
            DeliveryAddress *map[string]interface{}  `json:"delivery_address"`
            PaymentIdentifier string  `json:"payment_identifier"`
            Files []Files  `json:"files"`
            PickAtStoreUID int  `json:"pick_at_store_uid"`
            ExtraMeta *map[string]interface{}  `json:"extra_meta"`
            BillingAddress *map[string]interface{}  `json:"billing_address"`
            Staff *map[string]interface{}  `json:"staff"`
            BillingAddressID int  `json:"billing_address_id"`
            AddressID int  `json:"address_id"`
            PaymentParams *map[string]interface{}  `json:"payment_params"`
            OrderType string  `json:"order_type"`
            Pos bool  `json:"pos"`
            OrderingStore int  `json:"ordering_store"`
            PaymentMode string  `json:"payment_mode"`
            PaymentAutoConfirm bool  `json:"payment_auto_confirm"`
            FyndstoreEmpID string  `json:"fyndstore_emp_id"`
            CallbackURL string  `json:"callback_url"`
            Aggregator string  `json:"aggregator"`
            Meta *map[string]interface{}  `json:"meta"`
            MerchantCode string  `json:"merchant_code"`
         
    }
    
    // CartDeliveryModesResponse ...
    type CartDeliveryModesResponse struct {

        
            AvailableModes []string  `json:"available_modes"`
            PickupStores []int  `json:"pickup_stores"`
         
    }
    
    // PickupStoreDetail ...
    type PickupStoreDetail struct {

        
            Phone string  `json:"phone"`
            Country string  `json:"country"`
            Address string  `json:"address"`
            AddressType string  `json:"address_type"`
            City string  `json:"city"`
            Pincode int  `json:"pincode"`
            State string  `json:"state"`
            StoreCode string  `json:"store_code"`
            AreaCode string  `json:"area_code"`
            UID int  `json:"uid"`
            AreaCodeSlug string  `json:"area_code_slug"`
            Email string  `json:"email"`
            Name string  `json:"name"`
            Landmark string  `json:"landmark"`
            Area string  `json:"area"`
         
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
    
    // GetTatProductReqBody ...
    type GetTatProductReqBody struct {

        
            LocationDetails []LocationDetailsReq  `json:"location_details"`
            ToPincode string  `json:"to_pincode"`
         
    }
    
    // LocationDetailsReq ...
    type LocationDetailsReq struct {

        
            FromPincode string  `json:"from_pincode"`
            Articles []TatReqProductArticles  `json:"articles"`
            FulfillmentID int  `json:"fulfillment_id"`
         
    }
    
    // TatReqProductArticles ...
    type TatReqProductArticles struct {

        
            Category LogisticRequestCategory  `json:"category"`
         
    }
    
    // LogisticRequestCategory ...
    type LogisticRequestCategory struct {

        
            ID int  `json:"id"`
            Level string  `json:"level"`
         
    }
    
    // GetTatProductResponse ...
    type GetTatProductResponse struct {

        
            LocationDetails LocationDetails  `json:"location_details"`
            RequestUUID string  `json:"request_uuid"`
            Error *map[string]interface{}  `json:"error"`
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
            Articles TatProductArticles  `json:"articles"`
            FulfillmentID int  `json:"fulfillment_id"`
         
    }
    
    // TatProductArticles ...
    type TatProductArticles struct {

        
            Error *map[string]interface{}  `json:"error"`
            Category LogisticResponseCategory  `json:"category"`
            Promise LogisticPromise  `json:"promise"`
         
    }
    
    // LogisticResponseCategory ...
    type LogisticResponseCategory struct {

        
            ID int  `json:"id"`
            Level string  `json:"level"`
         
    }
    
    // LogisticPromise ...
    type LogisticPromise struct {

        
            Timestamp LogisticTimestamp  `json:"timestamp"`
            Formatted Formatted  `json:"formatted"`
         
    }
    
    // LogisticTimestamp ...
    type LogisticTimestamp struct {

        
            Min int  `json:"min"`
            Max int  `json:"max"`
         
    }
    
    // Formatted ...
    type Formatted struct {

        
            Min string  `json:"min"`
            Max string  `json:"max"`
         
    }
    
