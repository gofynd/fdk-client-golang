package application



import (
	"encoding/json"
	"fmt"
	"github.com/gofynd/fdk-client-golang/sdk/common"
)
// Client holds all the services instance 
type Client struct {
	
		Catalog  *Catalog
	
		Cart  *Cart
	
		Lead  *Lead
	
		Theme  *Theme
	
		User  *User
	
		Content  *Content
	
		Communication  *Communication
	
		Share  *Share
	
		FileStorage  *FileStorage
	
		Configuration  *Configuration
	
		Payment  *Payment
	
		Order  *Order
	
		Rewards  *Rewards
	
		Feedback  *Feedback
	
		PosCart  *PosCart
	
		Logistic  *Logistic
	
}

// NewAppClient provides application client
func NewAppClient(config *AppConfig) *Client {
		return &Client{
			
				Catalog:  NewCatalog(config),
			
				Cart:  NewCart(config),
			
				Lead:  NewLead(config),
			
				Theme:  NewTheme(config),
			
				User:  NewUser(config),
			
				Content:  NewContent(config),
			
				Communication:  NewCommunication(config),
			
				Share:  NewShare(config),
			
				FileStorage:  NewFileStorage(config),
			
				Configuration:  NewConfiguration(config),
			
				Payment:  NewPayment(config),
			
				Order:  NewOrder(config),
			
				Rewards:  NewRewards(config),
			
				Feedback:  NewFeedback(config),
			
				PosCart:  NewPosCart(config),
			
				Logistic:  NewLogistic(config),
			
		}
}


    // Catalog ...
    type Catalog struct {
        config *AppConfig
    }
    // NewCatalog ...
    func NewCatalog(config *AppConfig) *Catalog {
        return &Catalog{config}
    }
    
  
    
    // GetProductDetailBySlug Get a product
    func (ca *Catalog)  GetProductDetailBySlug(Slug string) (*ProductDetail, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getProductDetailBySlugResponse *ProductDetail
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/catalog/v1.0/products/%s/",Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ProductDetail{}, err
            
	    }
        err = json.Unmarshal(response, &getProductDetailBySlugResponse)
        if err != nil {
           
             return &ProductDetail{}, common.NewFDKError(err.Error())
            
        }
        return getProductDetailBySlugResponse, nil
    }
          
    
  
    
    // GetProductSizesBySlug Get the sizes of a product
    func (ca *Catalog)  GetProductSizesBySlug(Slug string, StoreID string) (*ProductSizes, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getProductSizesBySlugResponse *ProductSizes
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          StoreID string  `url:"store_id"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           StoreID : StoreID,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/catalog/v1.0/products/%s/sizes/",Slug),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ProductSizes{}, err
            
	    }
        err = json.Unmarshal(response, &getProductSizesBySlugResponse)
        if err != nil {
           
             return &ProductSizes{}, common.NewFDKError(err.Error())
            
        }
        return getProductSizesBySlugResponse, nil
    }
          
    
  
    
    // GetProductPriceBySlug Get price a product size
    func (ca *Catalog)  GetProductPriceBySlug(Slug string, Size string, Pincode string, StoreID string) (*ProductSizePriceResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getProductPriceBySlugResponse *ProductSizePriceResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          StoreID string  `url:"store_id"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           StoreID : StoreID,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/catalog/v1.0/products/%s/sizes/%s/pincode/%s/price/",Slug,Size,Pincode),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ProductSizePriceResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getProductPriceBySlugResponse)
        if err != nil {
           
             return &ProductSizePriceResponse{}, common.NewFDKError(err.Error())
            
        }
        return getProductPriceBySlugResponse, nil
    }
          
    
  
    
    // GetProductSellersBySlug List sellers of a product
    func (ca *Catalog)  GetProductSellersBySlug(Slug string, Size string, Pincode string, PageNo int, PageSize int) (*ProductSizeSellersResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getProductSellersBySlugResponse *ProductSizeSellersResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          PageNo int  `url:"page_no"` 
          PageSize int  `url:"page_size"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           PageNo : PageNo,
           PageSize : PageSize,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/catalog/v1.0/products/%s/sizes/%s/pincode/%s/sellers/",Slug,Size,Pincode),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ProductSizeSellersResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getProductSellersBySlugResponse)
        if err != nil {
           
             return &ProductSizeSellersResponse{}, common.NewFDKError(err.Error())
            
        }
        return getProductSellersBySlugResponse, nil
    }
          
            
            
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                    
                
                    
                        
                    
                    
                
            
            // GetProductSellersBySlugPaginator List sellers of a product  
            func (ca *Catalog)  GetProductSellersBySlugPaginator(Slug string , Size string , Pincode string , PageSize int ) *common.Paginator {
                paginator := common.NewPaginator("number")
                paginator.Next = func() (interface{}, error) {
                    response, err := ca.GetProductSellersBySlug(Slug, Size, Pincode, paginator.PageNo, PageSize)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    
  
    
    // GetProductComparisonBySlugs Compare products
    func (ca *Catalog)  GetProductComparisonBySlugs(Slug []string) (*ProductsComparisonResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getProductComparisonBySlugsResponse *ProductsComparisonResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Slug []string  `url:"slug"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Slug : Slug,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/catalog/v1.0/products/compare/",),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ProductsComparisonResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getProductComparisonBySlugsResponse)
        if err != nil {
           
             return &ProductsComparisonResponse{}, common.NewFDKError(err.Error())
            
        }
        return getProductComparisonBySlugsResponse, nil
    }
          
    
  
    
    // GetSimilarComparisonProductBySlug Get comparison between similar products
    func (ca *Catalog)  GetSimilarComparisonProductBySlug(Slug string) (*ProductCompareResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getSimilarComparisonProductBySlugResponse *ProductCompareResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/catalog/v1.0/products/%s/similar/compare/",Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ProductCompareResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getSimilarComparisonProductBySlugResponse)
        if err != nil {
           
             return &ProductCompareResponse{}, common.NewFDKError(err.Error())
            
        }
        return getSimilarComparisonProductBySlugResponse, nil
    }
          
    
  
    
    // GetComparedFrequentlyProductBySlug Get comparison between frequently compared products with the given product
    func (ca *Catalog)  GetComparedFrequentlyProductBySlug(Slug string) (*ProductFrequentlyComparedSimilarResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getComparedFrequentlyProductBySlugResponse *ProductFrequentlyComparedSimilarResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/catalog/v1.0/products/%s/similar/compared-frequently/",Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ProductFrequentlyComparedSimilarResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getComparedFrequentlyProductBySlugResponse)
        if err != nil {
           
             return &ProductFrequentlyComparedSimilarResponse{}, common.NewFDKError(err.Error())
            
        }
        return getComparedFrequentlyProductBySlugResponse, nil
    }
          
    
  
    
    // GetProductSimilarByIdentifier Get similar products
    func (ca *Catalog)  GetProductSimilarByIdentifier(Slug string, SimilarType string) (*SimilarProductByTypeResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getProductSimilarByIdentifierResponse *SimilarProductByTypeResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/catalog/v1.0/products/%s/similar/%s/",Slug,SimilarType),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &SimilarProductByTypeResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getProductSimilarByIdentifierResponse)
        if err != nil {
           
             return &SimilarProductByTypeResponse{}, common.NewFDKError(err.Error())
            
        }
        return getProductSimilarByIdentifierResponse, nil
    }
          
    
  
    
    // GetProductVariantsBySlug Get variant of a particular product
    func (ca *Catalog)  GetProductVariantsBySlug(Slug string) (*ProductVariantsResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getProductVariantsBySlugResponse *ProductVariantsResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/catalog/v1.0/products/%s/variants/",Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ProductVariantsResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getProductVariantsBySlugResponse)
        if err != nil {
           
             return &ProductVariantsResponse{}, common.NewFDKError(err.Error())
            
        }
        return getProductVariantsBySlugResponse, nil
    }
          
    
  
    
    // GetProductStockByIds Get the stock of a product
    func (ca *Catalog)  GetProductStockByIds(ItemID string, Alu string, SkuCode string, Ean string, Upc string) (*ProductStockStatusResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getProductStockByIdsResponse *ProductStockStatusResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          ItemID string  `url:"item_id"` 
          Alu string  `url:"alu"` 
          SkuCode string  `url:"sku_code"` 
          Ean string  `url:"ean"` 
          Upc string  `url:"upc"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           ItemID : ItemID,
           Alu : Alu,
           SkuCode : SkuCode,
           Ean : Ean,
           Upc : Upc,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            "/service/application/catalog/v1.0/products/stock-status/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ProductStockStatusResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getProductStockByIdsResponse)
        if err != nil {
           
             return &ProductStockStatusResponse{}, common.NewFDKError(err.Error())
            
        }
        return getProductStockByIdsResponse, nil
    }
          
    
  
    
    // GetProductStockForTimeByIds Get the stock of a product
    func (ca *Catalog)  GetProductStockForTimeByIds(Timestamp string, PageSize int, PageID string) (*ProductStockPolling, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getProductStockForTimeByIdsResponse *ProductStockPolling
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Timestamp string  `url:"timestamp"` 
          PageSize int  `url:"page_size"` 
          PageID string  `url:"page_id"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Timestamp : Timestamp,
           PageSize : PageSize,
           PageID : PageID,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/catalog/v1.0/products/stock-status/poll/",),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ProductStockPolling{}, err
            
	    }
        err = json.Unmarshal(response, &getProductStockForTimeByIdsResponse)
        if err != nil {
           
             return &ProductStockPolling{}, common.NewFDKError(err.Error())
            
        }
        return getProductStockForTimeByIdsResponse, nil
    }
          
            
            
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                    
                        
                    
                
            
            // GetProductStockForTimeByIdsPaginator Get the stock of a product  
            func (ca *Catalog)  GetProductStockForTimeByIdsPaginator(Timestamp string , PageSize int ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                paginator.Next = func() (interface{}, error) {
                    response, err := ca.GetProductStockForTimeByIds(Timestamp, PageSize, paginator.NextID)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    
  
    
    // GetProducts List the products
    func (ca *Catalog)  GetProducts(Q string, F string, Filters bool, SortOn string, PageID string, PageSize int, PageNo int, PageType string) (*ProductListingResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getProductsResponse *ProductListingResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Q string  `url:"q"` 
          F string  `url:"f"` 
          Filters bool  `url:"filters"` 
          SortOn string  `url:"sort_on"` 
          PageID string  `url:"page_id"` 
          PageSize int  `url:"page_size"` 
          PageNo int  `url:"page_no"` 
          PageType string  `url:"page_type"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Q : Q,
           F : F,
           Filters : Filters,
           SortOn : SortOn,
           PageID : PageID,
           PageSize : PageSize,
           PageNo : PageNo,
           PageType : PageType,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            "/service/application/catalog/v1.0/products/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ProductListingResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getProductsResponse)
        if err != nil {
           
             return &ProductListingResponse{}, common.NewFDKError(err.Error())
            
        }
        return getProductsResponse, nil
    }
          
            
            
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                    
                        
                    
                
                    
                        
                    
                    
                
                    
                    
                
                    
                    
                
            
            // GetProductsPaginator List the products  
            func (ca *Catalog)  GetProductsPaginator(Q string , F string , Filters bool , SortOn string , PageSize int ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                paginator.Next = func() (interface{}, error) {
                    response, err := ca.GetProducts(Q, F, Filters, SortOn, paginator.NextID, PageSize, paginator.PageNo, "cursor")
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    
  
    
    // GetBrands List all the brands
    func (ca *Catalog)  GetBrands(Department string, PageNo int, PageSize int) (*BrandListingResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getBrandsResponse *BrandListingResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Department string  `url:"department"` 
          PageNo int  `url:"page_no"` 
          PageSize int  `url:"page_size"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Department : Department,
           PageNo : PageNo,
           PageSize : PageSize,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            "/service/application/catalog/v1.0/brands/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &BrandListingResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getBrandsResponse)
        if err != nil {
           
             return &BrandListingResponse{}, common.NewFDKError(err.Error())
            
        }
        return getBrandsResponse, nil
    }
          
            
            
                
                    
                        
                    
                    
                
                    
                    
                
                    
                        
                    
                    
                
            
            // GetBrandsPaginator List all the brands  
            func (ca *Catalog)  GetBrandsPaginator(Department string , PageSize int ) *common.Paginator {
                paginator := common.NewPaginator("number")
                paginator.Next = func() (interface{}, error) {
                    response, err := ca.GetBrands(Department, paginator.PageNo, PageSize)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    
  
    
    // GetBrandDetailBySlug Get metadata of a brand
    func (ca *Catalog)  GetBrandDetailBySlug(Slug string) (*BrandDetailResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getBrandDetailBySlugResponse *BrandDetailResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/catalog/v1.0/brands/%s/",Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &BrandDetailResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getBrandDetailBySlugResponse)
        if err != nil {
           
             return &BrandDetailResponse{}, common.NewFDKError(err.Error())
            
        }
        return getBrandDetailBySlugResponse, nil
    }
          
    
  
    
    // GetCategories List all the categories
    func (ca *Catalog)  GetCategories(Department string) (*CategoryListingResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getCategoriesResponse *CategoryListingResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Department string  `url:"department"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Department : Department,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            "/service/application/catalog/v1.0/categories/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CategoryListingResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getCategoriesResponse)
        if err != nil {
           
             return &CategoryListingResponse{}, common.NewFDKError(err.Error())
            
        }
        return getCategoriesResponse, nil
    }
          
    
  
    
    // GetCategoryDetailBySlug Get metadata of a category
    func (ca *Catalog)  GetCategoryDetailBySlug(Slug string) (*CategoryMetaResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getCategoryDetailBySlugResponse *CategoryMetaResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/catalog/v1.0/categories/%s/",Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CategoryMetaResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getCategoryDetailBySlugResponse)
        if err != nil {
           
             return &CategoryMetaResponse{}, common.NewFDKError(err.Error())
            
        }
        return getCategoryDetailBySlugResponse, nil
    }
          
    
  
    
    // GetHomeProducts List the products
    func (ca *Catalog)  GetHomeProducts(SortOn string, PageID string, PageSize int) (*HomeListingResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getHomeProductsResponse *HomeListingResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          SortOn string  `url:"sort_on"` 
          PageID string  `url:"page_id"` 
          PageSize int  `url:"page_size"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           SortOn : SortOn,
           PageID : PageID,
           PageSize : PageSize,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            "/service/application/catalog/v1.0/home/listing/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &HomeListingResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getHomeProductsResponse)
        if err != nil {
           
             return &HomeListingResponse{}, common.NewFDKError(err.Error())
            
        }
        return getHomeProductsResponse, nil
    }
          
            
            
                
                    
                        
                    
                    
                
                    
                    
                        
                    
                
                    
                        
                    
                    
                
            
            // GetHomeProductsPaginator List the products  
            func (ca *Catalog)  GetHomeProductsPaginator(SortOn string , PageSize int ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                paginator.Next = func() (interface{}, error) {
                    response, err := ca.GetHomeProducts(SortOn, paginator.NextID, PageSize)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    
  
    
    // GetDepartments List all the departments
    func (ca *Catalog)  GetDepartments() (*DepartmentResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getDepartmentsResponse *DepartmentResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            "/service/application/catalog/v1.0/departments/",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &DepartmentResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getDepartmentsResponse)
        if err != nil {
           
             return &DepartmentResponse{}, common.NewFDKError(err.Error())
            
        }
        return getDepartmentsResponse, nil
    }
          
    
  
    
    // GetSearchResults Get relevant suggestions for a search query
    func (ca *Catalog)  GetSearchResults(Q string) (*AutoCompleteResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getSearchResultsResponse *AutoCompleteResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Q string  `url:"q"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Q : Q,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/catalog/v1.0/auto-complete/",),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &AutoCompleteResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getSearchResultsResponse)
        if err != nil {
           
             return &AutoCompleteResponse{}, common.NewFDKError(err.Error())
            
        }
        return getSearchResultsResponse, nil
    }
          
    
  
    
    // GetCollections List all the collections
    func (ca *Catalog)  GetCollections(PageNo int, PageSize int) (*GetCollectionListingResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getCollectionsResponse *GetCollectionListingResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          PageNo int  `url:"page_no"` 
          PageSize int  `url:"page_size"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           PageNo : PageNo,
           PageSize : PageSize,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            "/service/application/catalog/v1.0/collections/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &GetCollectionListingResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getCollectionsResponse)
        if err != nil {
           
             return &GetCollectionListingResponse{}, common.NewFDKError(err.Error())
            
        }
        return getCollectionsResponse, nil
    }
          
            
            
                
                    
                    
                
                    
                        
                    
                    
                
            
            // GetCollectionsPaginator List all the collections  
            func (ca *Catalog)  GetCollectionsPaginator(PageSize int ) *common.Paginator {
                paginator := common.NewPaginator("number")
                paginator.Next = func() (interface{}, error) {
                    response, err := ca.GetCollections(paginator.PageNo, PageSize)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    
  
    
    // GetCollectionItemsBySlug Get the items in a collection
    func (ca *Catalog)  GetCollectionItemsBySlug(Slug string, F string, Filters bool, SortOn string, PageID string, PageSize int) (*ProductListingResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getCollectionItemsBySlugResponse *ProductListingResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          F string  `url:"f"` 
          Filters bool  `url:"filters"` 
          SortOn string  `url:"sort_on"` 
          PageID string  `url:"page_id"` 
          PageSize int  `url:"page_size"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           F : F,
           Filters : Filters,
           SortOn : SortOn,
           PageID : PageID,
           PageSize : PageSize,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/catalog/v1.0/collections/%s/items/",Slug),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ProductListingResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getCollectionItemsBySlugResponse)
        if err != nil {
           
             return &ProductListingResponse{}, common.NewFDKError(err.Error())
            
        }
        return getCollectionItemsBySlugResponse, nil
    }
          
            
            
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                    
                        
                    
                
                    
                        
                    
                    
                
            
            // GetCollectionItemsBySlugPaginator Get the items in a collection  
            func (ca *Catalog)  GetCollectionItemsBySlugPaginator(Slug string , F string , Filters bool , SortOn string , PageSize int ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                paginator.Next = func() (interface{}, error) {
                    response, err := ca.GetCollectionItemsBySlug(Slug, F, Filters, SortOn, paginator.NextID, PageSize)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    
  
    
    // GetCollectionDetailBySlug Get a particular collection
    func (ca *Catalog)  GetCollectionDetailBySlug(Slug string) (*CollectionDetailResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getCollectionDetailBySlugResponse *CollectionDetailResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/catalog/v1.0/collections/%s/",Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CollectionDetailResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getCollectionDetailBySlugResponse)
        if err != nil {
           
             return &CollectionDetailResponse{}, common.NewFDKError(err.Error())
            
        }
        return getCollectionDetailBySlugResponse, nil
    }
          
    
  
    
    // GetFollowedListing Get a list of followed Products, Brands, Collections
    func (ca *Catalog)  GetFollowedListing(CollectionType string, PageID string, PageSize int) (*GetFollowListingResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getFollowedListingResponse *GetFollowListingResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          PageID string  `url:"page_id"` 
          PageSize int  `url:"page_size"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           PageID : PageID,
           PageSize : PageSize,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/catalog/v1.0/follow/%s/",CollectionType),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &GetFollowListingResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getFollowedListingResponse)
        if err != nil {
           
             return &GetFollowListingResponse{}, common.NewFDKError(err.Error())
            
        }
        return getFollowedListingResponse, nil
    }
          
            
            
                
                    
                        
                    
                    
                
                    
                    
                        
                    
                
                    
                        
                    
                    
                
            
            // GetFollowedListingPaginator Get a list of followed Products, Brands, Collections  
            func (ca *Catalog)  GetFollowedListingPaginator(CollectionType string , PageSize int ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                paginator.Next = func() (interface{}, error) {
                    response, err := ca.GetFollowedListing(CollectionType, paginator.NextID, PageSize)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    
  
    
    // FollowById Follow a particular Product
    func (ca *Catalog)  FollowById(CollectionType string, CollectionID string) (*FollowPostResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            followByIdResponse *FollowPostResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "post",
            fmt.Sprintf("/service/application/catalog/v1.0/follow/%s/%s/",CollectionType,CollectionID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &FollowPostResponse{}, err
            
	    }
        err = json.Unmarshal(response, &followByIdResponse)
        if err != nil {
           
             return &FollowPostResponse{}, common.NewFDKError(err.Error())
            
        }
        return followByIdResponse, nil
    }
          
    
  
    
    // UnfollowById UnFollow a Product
    func (ca *Catalog)  UnfollowById(CollectionType string, CollectionID string) (*FollowPostResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            unfollowByIdResponse *FollowPostResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "delete",
            fmt.Sprintf("/service/application/catalog/v1.0/follow/%s/%s/",CollectionType,CollectionID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &FollowPostResponse{}, err
            
	    }
        err = json.Unmarshal(response, &unfollowByIdResponse)
        if err != nil {
           
             return &FollowPostResponse{}, common.NewFDKError(err.Error())
            
        }
        return unfollowByIdResponse, nil
    }
          
    
  
    
    // GetFollowerCountById Get Follow Count
    func (ca *Catalog)  GetFollowerCountById(CollectionType string, CollectionID string) (*FollowerCountResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getFollowerCountByIdResponse *FollowerCountResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/catalog/v1.0/follow/%s/%s/count/",CollectionType,CollectionID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &FollowerCountResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getFollowerCountByIdResponse)
        if err != nil {
           
             return &FollowerCountResponse{}, common.NewFDKError(err.Error())
            
        }
        return getFollowerCountByIdResponse, nil
    }
          
    
  
    
    // GetFollowIds Get the Ids of followed product, brand and collection.
    func (ca *Catalog)  GetFollowIds(CollectionType string) (*FollowIdsResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getFollowIdsResponse *FollowIdsResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          CollectionType string  `url:"collection_type"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           CollectionType : CollectionType,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            "/service/application/catalog/v1.0/follow/ids/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &FollowIdsResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getFollowIdsResponse)
        if err != nil {
           
             return &FollowIdsResponse{}, common.NewFDKError(err.Error())
            
        }
        return getFollowIdsResponse, nil
    }
          
    
  
    
    // GetStores List store meta information.
    func (ca *Catalog)  GetStores(PageNo int, PageSize int, Q string, Range int, Latitude int, Longitude int) (*StoreListingResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getStoresResponse *StoreListingResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          PageNo int  `url:"page_no"` 
          PageSize int  `url:"page_size"` 
          Q string  `url:"q"` 
          Range int  `url:"range"` 
          Latitude int  `url:"latitude"` 
          Longitude int  `url:"longitude"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           PageNo : PageNo,
           PageSize : PageSize,
           Q : Q,
           Range : Range,
           Latitude : Latitude,
           Longitude : Longitude,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            "/service/application/catalog/v1.0/locations/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &StoreListingResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getStoresResponse)
        if err != nil {
           
             return &StoreListingResponse{}, common.NewFDKError(err.Error())
            
        }
        return getStoresResponse, nil
    }
          
            
            
                
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
            
            // GetStoresPaginator List store meta information.  
            func (ca *Catalog)  GetStoresPaginator(PageSize int , Q string , Range int , Latitude int , Longitude int ) *common.Paginator {
                paginator := common.NewPaginator("number")
                paginator.Next = func() (interface{}, error) {
                    response, err := ca.GetStores(paginator.PageNo, PageSize, Q, Range, Latitude, Longitude)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    

    // Cart ...
    type Cart struct {
        config *AppConfig
    }
    // NewCart ...
    func NewCart(config *AppConfig) *Cart {
        return &Cart{config}
    }
    
  
    
    // GetCart Fetch all Items Added to  Cart
    func (ca *Cart)  GetCart(UID int, I bool, B bool, AssignCardID int) (*CartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getCartResponse *CartResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID int  `url:"uid"` 
          I bool  `url:"i"` 
          B bool  `url:"b"` 
          AssignCardID int  `url:"assign_card_id"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
           I : I,
           B : B,
           AssignCardID : AssignCardID,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            "/service/application/cart/v1.0/detail",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CartResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getCartResponse)
        if err != nil {
           
             return &CartResponse{}, common.NewFDKError(err.Error())
            
        }
        return getCartResponse, nil
    }
          
    
  
    
    // GetCartLastModified Fetch Last-Modified timestamp
    func (ca *Cart)  GetCartLastModified(UID int) (interface{}, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
             getCartLastModifiedResponse interface{}
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID int  `url:"uid"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "head",
            "/service/application/cart/v1.0/detail",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
            return getCartLastModifiedResponse, err
            
	    }
        err = json.Unmarshal(response, &getCartLastModifiedResponse)
        if err != nil {
           
            return getCartLastModifiedResponse, common.NewFDKError(err.Error())
            
        }
        return getCartLastModifiedResponse, nil
    }
          
    
  
    
    // AddItems Add Items to Cart
    func (ca *Cart)  AddItems(I bool, B bool, body  AddCartRequest) (*AddCartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            addItemsResponse *AddCartResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          I bool  `url:"i"` 
          B bool  `url:"b"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           I : I,
           B : B,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &AddCartResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &AddCartResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "post",
            "/service/application/cart/v1.0/detail",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &AddCartResponse{}, err
            
	    }
        err = json.Unmarshal(response, &addItemsResponse)
        if err != nil {
           
             return &AddCartResponse{}, common.NewFDKError(err.Error())
            
        }
        return addItemsResponse, nil
    }
          
    
  
    
    // UpdateCart Update Items already added to Cart
    func (ca *Cart)  UpdateCart(UID int, I bool, B bool, body  UpdateCartRequest) (*UpdateCartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateCartResponse *UpdateCartResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID int  `url:"uid"` 
          I bool  `url:"i"` 
          B bool  `url:"b"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
           I : I,
           B : B,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &UpdateCartResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &UpdateCartResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "put",
            "/service/application/cart/v1.0/detail",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &UpdateCartResponse{}, err
            
	    }
        err = json.Unmarshal(response, &updateCartResponse)
        if err != nil {
           
             return &UpdateCartResponse{}, common.NewFDKError(err.Error())
            
        }
        return updateCartResponse, nil
    }
          
    
  
    
    // GetItemCount Cart item count
    func (ca *Cart)  GetItemCount(UID int) (*CartItemCountResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getItemCountResponse *CartItemCountResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID int  `url:"uid"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            "/service/application/cart/v1.0/basic",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CartItemCountResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getItemCountResponse)
        if err != nil {
           
             return &CartItemCountResponse{}, common.NewFDKError(err.Error())
            
        }
        return getItemCountResponse, nil
    }
          
    
  
    
    // GetCoupons Fetch Coupon
    func (ca *Cart)  GetCoupons(UID int) (*GetCouponResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getCouponsResponse *GetCouponResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID int  `url:"uid"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            "/service/application/cart/v1.0/coupon",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &GetCouponResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getCouponsResponse)
        if err != nil {
           
             return &GetCouponResponse{}, common.NewFDKError(err.Error())
            
        }
        return getCouponsResponse, nil
    }
          
    
  
    
    // ApplyCoupon Apply Coupon
    func (ca *Cart)  ApplyCoupon(I bool, B bool, P bool, UID int, body  ApplyCouponRequest) (*CartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            applyCouponResponse *CartResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          I bool  `url:"i"` 
          B bool  `url:"b"` 
          P bool  `url:"p"` 
          UID int  `url:"uid"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           I : I,
           B : B,
           P : P,
           UID : UID,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &CartResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &CartResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "post",
            "/service/application/cart/v1.0/coupon",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CartResponse{}, err
            
	    }
        err = json.Unmarshal(response, &applyCouponResponse)
        if err != nil {
           
             return &CartResponse{}, common.NewFDKError(err.Error())
            
        }
        return applyCouponResponse, nil
    }
          
    
  
    
    // RemoveCoupon Remove Coupon Applied
    func (ca *Cart)  RemoveCoupon(UID int) (*CartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            removeCouponResponse *CartResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID int  `url:"uid"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "delete",
            "/service/application/cart/v1.0/coupon",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CartResponse{}, err
            
	    }
        err = json.Unmarshal(response, &removeCouponResponse)
        if err != nil {
           
             return &CartResponse{}, common.NewFDKError(err.Error())
            
        }
        return removeCouponResponse, nil
    }
          
    
  
    
    // GetBulkDiscountOffers Get discount offers based on quantity
    func (ca *Cart)  GetBulkDiscountOffers(ItemID int, ArticleID string, UID int, Slug string) (*BulkPriceResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getBulkDiscountOffersResponse *BulkPriceResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          ItemID int  `url:"item_id"` 
          ArticleID string  `url:"article_id"` 
          UID int  `url:"uid"` 
          Slug string  `url:"slug"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           ItemID : ItemID,
           ArticleID : ArticleID,
           UID : UID,
           Slug : Slug,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            "/service/application/cart/v1.0/bulk-price",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &BulkPriceResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getBulkDiscountOffersResponse)
        if err != nil {
           
             return &BulkPriceResponse{}, common.NewFDKError(err.Error())
            
        }
        return getBulkDiscountOffersResponse, nil
    }
          
    
  
    
    // GetAddresses Fetch Address
    func (ca *Cart)  GetAddresses(UID int, MobileNo string, CheckoutMode string, Tags string, IsDefault bool) (*GetAddressesResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getAddressesResponse *GetAddressesResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID int  `url:"uid"` 
          MobileNo string  `url:"mobile_no"` 
          CheckoutMode string  `url:"checkout_mode"` 
          Tags string  `url:"tags"` 
          IsDefault bool  `url:"is_default"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
           MobileNo : MobileNo,
           CheckoutMode : CheckoutMode,
           Tags : Tags,
           IsDefault : IsDefault,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            "/service/application/cart/v1.0/address",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &GetAddressesResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getAddressesResponse)
        if err != nil {
           
             return &GetAddressesResponse{}, common.NewFDKError(err.Error())
            
        }
        return getAddressesResponse, nil
    }
          
    
  
    
    // AddAddress Add Address to the account
    func (ca *Cart)  AddAddress(body  Address) (*SaveAddressResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            addAddressResponse *SaveAddressResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &SaveAddressResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &SaveAddressResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "post",
            "/service/application/cart/v1.0/address",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &SaveAddressResponse{}, err
            
	    }
        err = json.Unmarshal(response, &addAddressResponse)
        if err != nil {
           
             return &SaveAddressResponse{}, common.NewFDKError(err.Error())
            
        }
        return addAddressResponse, nil
    }
          
    
  
    
    // GetAddressById Fetch Single Address
    func (ca *Cart)  GetAddressById(ID int, UID int, MobileNo string, CheckoutMode string, Tags string, IsDefault bool) (*Address, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getAddressByIdResponse *Address
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID int  `url:"uid"` 
          MobileNo string  `url:"mobile_no"` 
          CheckoutMode string  `url:"checkout_mode"` 
          Tags string  `url:"tags"` 
          IsDefault bool  `url:"is_default"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
           MobileNo : MobileNo,
           CheckoutMode : CheckoutMode,
           Tags : Tags,
           IsDefault : IsDefault,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/cart/v1.0/address/%d",ID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &Address{}, err
            
	    }
        err = json.Unmarshal(response, &getAddressByIdResponse)
        if err != nil {
           
             return &Address{}, common.NewFDKError(err.Error())
            
        }
        return getAddressByIdResponse, nil
    }
          
    
  
    
    // UpdateAddress Update Address alreay added to account
    func (ca *Cart)  UpdateAddress(ID int, body  Address) (*UpdateAddressResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateAddressResponse *UpdateAddressResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &UpdateAddressResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &UpdateAddressResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "put",
            fmt.Sprintf("/service/application/cart/v1.0/address/%d",ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &UpdateAddressResponse{}, err
            
	    }
        err = json.Unmarshal(response, &updateAddressResponse)
        if err != nil {
           
             return &UpdateAddressResponse{}, common.NewFDKError(err.Error())
            
        }
        return updateAddressResponse, nil
    }
          
    
  
    
    // RemoveAddress Remove Address Associated to the account
    func (ca *Cart)  RemoveAddress(ID int) (*DeleteAddressResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            removeAddressResponse *DeleteAddressResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "delete",
            fmt.Sprintf("/service/application/cart/v1.0/address/%d",ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &DeleteAddressResponse{}, err
            
	    }
        err = json.Unmarshal(response, &removeAddressResponse)
        if err != nil {
           
             return &DeleteAddressResponse{}, common.NewFDKError(err.Error())
            
        }
        return removeAddressResponse, nil
    }
          
    
  
    
    // SelectAddress Select Address from All Addresses
    func (ca *Cart)  SelectAddress(UID int, I bool, B bool, body  SelectCartAddressRequest) (*CartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            selectAddressResponse *CartResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID int  `url:"uid"` 
          I bool  `url:"i"` 
          B bool  `url:"b"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
           I : I,
           B : B,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &CartResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &CartResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "post",
            "/service/application/cart/v1.0/select-address",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CartResponse{}, err
            
	    }
        err = json.Unmarshal(response, &selectAddressResponse)
        if err != nil {
           
             return &CartResponse{}, common.NewFDKError(err.Error())
            
        }
        return selectAddressResponse, nil
    }
          
    
  
    
    // SelectPaymentMode Update Cart Payment
    func (ca *Cart)  SelectPaymentMode(UID string, body  UpdateCartPaymentRequest) (*CartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            selectPaymentModeResponse *CartResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID string  `url:"uid"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &CartResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &CartResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "put",
            "/service/application/cart/v1.0/payment",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CartResponse{}, err
            
	    }
        err = json.Unmarshal(response, &selectPaymentModeResponse)
        if err != nil {
           
             return &CartResponse{}, common.NewFDKError(err.Error())
            
        }
        return selectPaymentModeResponse, nil
    }
          
    
  
    
    // ValidateCouponForPayment Get Cart Payment for valid coupon
    func (ca *Cart)  ValidateCouponForPayment(UID string, AddressID string, PaymentMode string, PaymentIdentifier string, AggregatorName string, MerchantCode string) (*PaymentCouponValidate, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            validateCouponForPaymentResponse *PaymentCouponValidate
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID string  `url:"uid"` 
          AddressID string  `url:"address_id"` 
          PaymentMode string  `url:"payment_mode"` 
          PaymentIdentifier string  `url:"payment_identifier"` 
          AggregatorName string  `url:"aggregator_name"` 
          MerchantCode string  `url:"merchant_code"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
           AddressID : AddressID,
           PaymentMode : PaymentMode,
           PaymentIdentifier : PaymentIdentifier,
           AggregatorName : AggregatorName,
           MerchantCode : MerchantCode,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            "/service/application/cart/v1.0/payment/validate/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &PaymentCouponValidate{}, err
            
	    }
        err = json.Unmarshal(response, &validateCouponForPaymentResponse)
        if err != nil {
           
             return &PaymentCouponValidate{}, common.NewFDKError(err.Error())
            
        }
        return validateCouponForPaymentResponse, nil
    }
          
    
  
    
    // GetShipments Get delivery date and options before checkout
    func (ca *Cart)  GetShipments(P bool, UID int, AddressID int, AreaCode string) (*CartShipmentsResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getShipmentsResponse *CartShipmentsResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          P bool  `url:"p"` 
          UID int  `url:"uid"` 
          AddressID int  `url:"address_id"` 
          AreaCode string  `url:"area_code"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           P : P,
           UID : UID,
           AddressID : AddressID,
           AreaCode : AreaCode,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            "/service/application/cart/v1.0/shipment",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CartShipmentsResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getShipmentsResponse)
        if err != nil {
           
             return &CartShipmentsResponse{}, common.NewFDKError(err.Error())
            
        }
        return getShipmentsResponse, nil
    }
          
    
  
    
    // CheckoutCart Checkout Cart
    func (ca *Cart)  CheckoutCart(body  CartCheckoutRequest) (*CartCheckoutResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            checkoutCartResponse *CartCheckoutResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &CartCheckoutResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &CartCheckoutResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "post",
            "/service/application/cart/v1.0/checkout",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CartCheckoutResponse{}, err
            
	    }
        err = json.Unmarshal(response, &checkoutCartResponse)
        if err != nil {
           
             return &CartCheckoutResponse{}, common.NewFDKError(err.Error())
            
        }
        return checkoutCartResponse, nil
    }
          
    
  
    
    // UpdateCartMeta Update Cart Meta
    func (ca *Cart)  UpdateCartMeta(UID int, body  CartMetaRequest) (*CartMetaResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateCartMetaResponse *CartMetaResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID int  `url:"uid"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &CartMetaResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &CartMetaResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "put",
            "/service/application/cart/v1.0/meta",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CartMetaResponse{}, err
            
	    }
        err = json.Unmarshal(response, &updateCartMetaResponse)
        if err != nil {
           
             return &CartMetaResponse{}, common.NewFDKError(err.Error())
            
        }
        return updateCartMetaResponse, nil
    }
          
    
  
    
    // GetCartShareLink Generate Cart sharing link token
    func (ca *Cart)  GetCartShareLink(body  GetShareCartLinkRequest) (*GetShareCartLinkResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getCartShareLinkResponse *GetShareCartLinkResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &GetShareCartLinkResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &GetShareCartLinkResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "post",
            "/service/application/cart/v1.0/share-cart",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &GetShareCartLinkResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getCartShareLinkResponse)
        if err != nil {
           
             return &GetShareCartLinkResponse{}, common.NewFDKError(err.Error())
            
        }
        return getCartShareLinkResponse, nil
    }
          
    
  
    
    // GetCartSharedItems Get shared cart snapshot and cart response
    func (ca *Cart)  GetCartSharedItems(Token string) (*SharedCartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getCartSharedItemsResponse *SharedCartResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/cart/v1.0/share-cart/%s",Token),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &SharedCartResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getCartSharedItemsResponse)
        if err != nil {
           
             return &SharedCartResponse{}, common.NewFDKError(err.Error())
            
        }
        return getCartSharedItemsResponse, nil
    }
          
    
  
    
    // UpdateCartWithSharedItems Merge or Replace existing cart
    func (ca *Cart)  UpdateCartWithSharedItems(Token string, Action string) (*SharedCartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateCartWithSharedItemsResponse *SharedCartResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "post",
            fmt.Sprintf("/service/application/cart/v1.0/share-cart/%s/%s",Token,Action),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &SharedCartResponse{}, err
            
	    }
        err = json.Unmarshal(response, &updateCartWithSharedItemsResponse)
        if err != nil {
           
             return &SharedCartResponse{}, common.NewFDKError(err.Error())
            
        }
        return updateCartWithSharedItemsResponse, nil
    }
          
    

    // Lead ...
    type Lead struct {
        config *AppConfig
    }
    // NewLead ...
    func NewLead(config *AppConfig) *Lead {
        return &Lead{config}
    }
    
  
    
    // GetTicket Get Ticket with the specific id
    func (le *Lead)  GetTicket(ID string) (*Ticket, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getTicketResponse *Ticket
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            le.config,
            "get",
            fmt.Sprintf("/service/application/lead/v1.0/ticket/%s",ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &Ticket{}, err
            
	    }
        err = json.Unmarshal(response, &getTicketResponse)
        if err != nil {
           
             return &Ticket{}, common.NewFDKError(err.Error())
            
        }
        return getTicketResponse, nil
    }
          
    
  
    
    // CreateHistory Create history for specific Ticket
    func (le *Lead)  CreateHistory(TicketID string, body  TicketHistoryPayload) (*TicketHistory, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            createHistoryResponse *TicketHistory
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &TicketHistory{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &TicketHistory{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            le.config,
            "post",
            fmt.Sprintf("/service/application/lead/v1.0/ticket/%s/history",TicketID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &TicketHistory{}, err
            
	    }
        err = json.Unmarshal(response, &createHistoryResponse)
        if err != nil {
           
             return &TicketHistory{}, common.NewFDKError(err.Error())
            
        }
        return createHistoryResponse, nil
    }
          
    
  
    
    // CreateTicket Create Ticket
    func (le *Lead)  CreateTicket(body  AddTicketPayload) (*Ticket, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            createTicketResponse *Ticket
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &Ticket{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &Ticket{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            le.config,
            "post",
            "/service/application/lead/v1.0/ticket/",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &Ticket{}, err
            
	    }
        err = json.Unmarshal(response, &createTicketResponse)
        if err != nil {
           
             return &Ticket{}, common.NewFDKError(err.Error())
            
        }
        return createTicketResponse, nil
    }
          
    
  
    
    // GetCustomForm Get specific Custom Form using it's slug
    func (le *Lead)  GetCustomForm(Slug string) (*CustomForm, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getCustomFormResponse *CustomForm
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            le.config,
            "get",
            fmt.Sprintf("/service/application/lead/v1.0/form/%s",Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CustomForm{}, err
            
	    }
        err = json.Unmarshal(response, &getCustomFormResponse)
        if err != nil {
           
             return &CustomForm{}, common.NewFDKError(err.Error())
            
        }
        return getCustomFormResponse, nil
    }
          
    
  
    
    // SubmitCustomForm Submit Response for a specific Custom Form using it's slug
    func (le *Lead)  SubmitCustomForm(Slug string, body  CustomFormSubmissionPayload) (*SubmitCustomFormResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            submitCustomFormResponse *SubmitCustomFormResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &SubmitCustomFormResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &SubmitCustomFormResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            le.config,
            "post",
            fmt.Sprintf("/service/application/lead/v1.0/form/%s/submit",Slug),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &SubmitCustomFormResponse{}, err
            
	    }
        err = json.Unmarshal(response, &submitCustomFormResponse)
        if err != nil {
           
             return &SubmitCustomFormResponse{}, common.NewFDKError(err.Error())
            
        }
        return submitCustomFormResponse, nil
    }
          
    
  
    
    // GetParticipantsInsideVideoRoom Get participants of a specific Video Room using it's unique name
    func (le *Lead)  GetParticipantsInsideVideoRoom(UniqueName string) (*GetParticipantsInsideVideoRoomResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getParticipantsInsideVideoRoomResponse *GetParticipantsInsideVideoRoomResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            le.config,
            "get",
            fmt.Sprintf("/service/application/lead/v1.0/video/room/%s/participants",UniqueName),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &GetParticipantsInsideVideoRoomResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getParticipantsInsideVideoRoomResponse)
        if err != nil {
           
             return &GetParticipantsInsideVideoRoomResponse{}, common.NewFDKError(err.Error())
            
        }
        return getParticipantsInsideVideoRoomResponse, nil
    }
          
    
  
    
    // GetTokenForVideoRoom Get Token to join a specific Video Room using it's unqiue name
    func (le *Lead)  GetTokenForVideoRoom(UniqueName string) (*GetTokenForVideoRoomResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getTokenForVideoRoomResponse *GetTokenForVideoRoomResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            le.config,
            "get",
            fmt.Sprintf("/service/application/lead/v1.0/video/room/%s/token",UniqueName),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &GetTokenForVideoRoomResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getTokenForVideoRoomResponse)
        if err != nil {
           
             return &GetTokenForVideoRoomResponse{}, common.NewFDKError(err.Error())
            
        }
        return getTokenForVideoRoomResponse, nil
    }
          
    

    // Theme ...
    type Theme struct {
        config *AppConfig
    }
    // NewTheme ...
    func NewTheme(config *AppConfig) *Theme {
        return &Theme{config}
    }
    
  
    
    // GetAppliedTheme Get applied theme for an application
    func (th *Theme)  GetAppliedTheme() (*ThemesSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getAppliedThemeResponse *ThemesSchema
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "get",
            "/service/application/theme/v1.0/applied-theme",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ThemesSchema{}, err
            
	    }
        err = json.Unmarshal(response, &getAppliedThemeResponse)
        if err != nil {
           
             return &ThemesSchema{}, common.NewFDKError(err.Error())
            
        }
        return getAppliedThemeResponse, nil
    }
          
    
  
    
    // GetThemeForPreview Get theme for preview
    func (th *Theme)  GetThemeForPreview(ThemeID string) (*ThemesSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getThemeForPreviewResponse *ThemesSchema
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "get",
            fmt.Sprintf("/service/application/theme/v1.0/%s/preview",ThemeID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ThemesSchema{}, err
            
	    }
        err = json.Unmarshal(response, &getThemeForPreviewResponse)
        if err != nil {
           
             return &ThemesSchema{}, common.NewFDKError(err.Error())
            
        }
        return getThemeForPreviewResponse, nil
    }
          
    

    // User ...
    type User struct {
        config *AppConfig
    }
    // NewUser ...
    func NewUser(config *AppConfig) *User {
        return &User{config}
    }
    
  
    
    // LoginWithFacebook Login/Register with Facebook
    func (us *User)  LoginWithFacebook(body  OAuthRequestSchema) (*AuthSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            loginWithFacebookResponse *AuthSuccess
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &AuthSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &AuthSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/login/facebook-token",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &AuthSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &loginWithFacebookResponse)
        if err != nil {
           
             return &AuthSuccess{}, common.NewFDKError(err.Error())
            
        }
        return loginWithFacebookResponse, nil
    }
          
    
  
    
    // LoginWithGoogle Login/Register with Google
    func (us *User)  LoginWithGoogle(body  OAuthRequestSchema) (*AuthSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            loginWithGoogleResponse *AuthSuccess
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &AuthSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &AuthSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/login/google-token",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &AuthSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &loginWithGoogleResponse)
        if err != nil {
           
             return &AuthSuccess{}, common.NewFDKError(err.Error())
            
        }
        return loginWithGoogleResponse, nil
    }
          
    
  
    
    // LoginWithGoogleAndroid Login/Register with Google for android
    func (us *User)  LoginWithGoogleAndroid(body  OAuthRequestSchema) (*AuthSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            loginWithGoogleAndroidResponse *AuthSuccess
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &AuthSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &AuthSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/login/google-android",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &AuthSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &loginWithGoogleAndroidResponse)
        if err != nil {
           
             return &AuthSuccess{}, common.NewFDKError(err.Error())
            
        }
        return loginWithGoogleAndroidResponse, nil
    }
          
    
  
    
    // LoginWithGoogleIOS Login/Register with Google for ios
    func (us *User)  LoginWithGoogleIOS(body  OAuthRequestSchema) (*AuthSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            loginWithGoogleIOSResponse *AuthSuccess
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &AuthSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &AuthSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/login/google-ios",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &AuthSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &loginWithGoogleIOSResponse)
        if err != nil {
           
             return &AuthSuccess{}, common.NewFDKError(err.Error())
            
        }
        return loginWithGoogleIOSResponse, nil
    }
          
    
  
    
    // LoginWithOTP Login/Register with OTP
    func (us *User)  LoginWithOTP(Platform string, body  SendOtpRequestSchema) (*SendOtpResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            loginWithOTPResponse *SendOtpResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Platform string  `url:"platform"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Platform : Platform,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &SendOtpResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &SendOtpResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/login/otp",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &SendOtpResponse{}, err
            
	    }
        err = json.Unmarshal(response, &loginWithOTPResponse)
        if err != nil {
           
             return &SendOtpResponse{}, common.NewFDKError(err.Error())
            
        }
        return loginWithOTPResponse, nil
    }
          
    
  
    
    // LoginWithEmailAndPassword Login/Register with password
    func (us *User)  LoginWithEmailAndPassword(body  PasswordLoginRequestSchema) (*LoginSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            loginWithEmailAndPasswordResponse *LoginSuccess
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &LoginSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &LoginSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/login/password",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &LoginSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &loginWithEmailAndPasswordResponse)
        if err != nil {
           
             return &LoginSuccess{}, common.NewFDKError(err.Error())
            
        }
        return loginWithEmailAndPasswordResponse, nil
    }
          
    
  
    
    // SendResetPasswordEmail Reset Password
    func (us *User)  SendResetPasswordEmail(Platform string, body  SendResetPasswordEmailRequestSchema) (*ResetPasswordSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            sendResetPasswordEmailResponse *ResetPasswordSuccess
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Platform string  `url:"platform"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Platform : Platform,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &ResetPasswordSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &ResetPasswordSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/login/password/reset",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ResetPasswordSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &sendResetPasswordEmailResponse)
        if err != nil {
           
             return &ResetPasswordSuccess{}, common.NewFDKError(err.Error())
            
        }
        return sendResetPasswordEmailResponse, nil
    }
          
    
  
    
    // ForgotPassword 
    func (us *User)  ForgotPassword(body  ForgotPasswordRequestSchema) (*LoginSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            forgotPasswordResponse *LoginSuccess
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &LoginSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &LoginSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/login/password/reset/forgot",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &LoginSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &forgotPasswordResponse)
        if err != nil {
           
             return &LoginSuccess{}, common.NewFDKError(err.Error())
            
        }
        return forgotPasswordResponse, nil
    }
          
    
  
    
    // SendResetToken 
    func (us *User)  SendResetToken(body  CodeRequestBodySchema) (*ResetPasswordSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            sendResetTokenResponse *ResetPasswordSuccess
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &ResetPasswordSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &ResetPasswordSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/login/password/reset/token",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ResetPasswordSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &sendResetTokenResponse)
        if err != nil {
           
             return &ResetPasswordSuccess{}, common.NewFDKError(err.Error())
            
        }
        return sendResetTokenResponse, nil
    }
          
    
  
    
    // LoginWithToken Login/Register with token
    func (us *User)  LoginWithToken(body  TokenRequestBodySchema) (*LoginSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            loginWithTokenResponse *LoginSuccess
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &LoginSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &LoginSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/login/token",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &LoginSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &loginWithTokenResponse)
        if err != nil {
           
             return &LoginSuccess{}, common.NewFDKError(err.Error())
            
        }
        return loginWithTokenResponse, nil
    }
          
    
  
    
    // RegisterWithForm Registration Form
    func (us *User)  RegisterWithForm(Platform string, body  FormRegisterRequestSchema) (*RegisterFormSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            registerWithFormResponse *RegisterFormSuccess
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Platform string  `url:"platform"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Platform : Platform,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &RegisterFormSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &RegisterFormSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/register/form",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &RegisterFormSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &registerWithFormResponse)
        if err != nil {
           
             return &RegisterFormSuccess{}, common.NewFDKError(err.Error())
            
        }
        return registerWithFormResponse, nil
    }
          
    
  
    
    // VerifyEmail Verify email
    func (us *User)  VerifyEmail(body  CodeRequestBodySchema) (*VerifyEmailSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            verifyEmailResponse *VerifyEmailSuccess
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &VerifyEmailSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &VerifyEmailSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/verify/email",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &VerifyEmailSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &verifyEmailResponse)
        if err != nil {
           
             return &VerifyEmailSuccess{}, common.NewFDKError(err.Error())
            
        }
        return verifyEmailResponse, nil
    }
          
    
  
    
    // VerifyMobile Verify mobile
    func (us *User)  VerifyMobile(body  CodeRequestBodySchema) (*VerifyEmailSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            verifyMobileResponse *VerifyEmailSuccess
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &VerifyEmailSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &VerifyEmailSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/verify/mobile",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &VerifyEmailSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &verifyMobileResponse)
        if err != nil {
           
             return &VerifyEmailSuccess{}, common.NewFDKError(err.Error())
            
        }
        return verifyMobileResponse, nil
    }
          
    
  
    
    // HasPassword Check if user has password
    func (us *User)  HasPassword() (*HasPasswordSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            hasPasswordResponse *HasPasswordSuccess
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "get",
            "/service/application/user/authentication/v1.0/has-password",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &HasPasswordSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &hasPasswordResponse)
        if err != nil {
           
             return &HasPasswordSuccess{}, common.NewFDKError(err.Error())
            
        }
        return hasPasswordResponse, nil
    }
          
    
  
    
    // UpdatePassword Update user password
    func (us *User)  UpdatePassword(body  UpdatePasswordRequestSchema) (*VerifyEmailSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updatePasswordResponse *VerifyEmailSuccess
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &VerifyEmailSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &VerifyEmailSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/password",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &VerifyEmailSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &updatePasswordResponse)
        if err != nil {
           
             return &VerifyEmailSuccess{}, common.NewFDKError(err.Error())
            
        }
        return updatePasswordResponse, nil
    }
          
    
  
    
    // Logout Logout user
    func (us *User)  Logout() (*LogoutSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            logoutResponse *LogoutSuccess
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "get",
            "/service/application/user/authentication/v1.0/logout",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &LogoutSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &logoutResponse)
        if err != nil {
           
             return &LogoutSuccess{}, common.NewFDKError(err.Error())
            
        }
        return logoutResponse, nil
    }
          
    
  
    
    // SendOTPOnMobile Send OTP on mobile
    func (us *User)  SendOTPOnMobile(Platform string, body  SendMobileOtpRequestSchema) (*OtpSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            sendOTPOnMobileResponse *OtpSuccess
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Platform string  `url:"platform"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Platform : Platform,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &OtpSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &OtpSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/otp/mobile/send",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &OtpSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &sendOTPOnMobileResponse)
        if err != nil {
           
             return &OtpSuccess{}, common.NewFDKError(err.Error())
            
        }
        return sendOTPOnMobileResponse, nil
    }
          
    
  
    
    // VerifyMobileOTP Verify OTP on mobile
    func (us *User)  VerifyMobileOTP(Platform string, body  VerifyOtpRequestSchema) (*VerifyOtpSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            verifyMobileOTPResponse *VerifyOtpSuccess
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Platform string  `url:"platform"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Platform : Platform,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &VerifyOtpSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &VerifyOtpSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/otp/mobile/verify",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &VerifyOtpSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &verifyMobileOTPResponse)
        if err != nil {
           
             return &VerifyOtpSuccess{}, common.NewFDKError(err.Error())
            
        }
        return verifyMobileOTPResponse, nil
    }
          
    
  
    
    // SendOTPOnEmail Send OTP on email
    func (us *User)  SendOTPOnEmail(Platform string, body  SendEmailOtpRequestSchema) (*EmailOtpSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            sendOTPOnEmailResponse *EmailOtpSuccess
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Platform string  `url:"platform"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Platform : Platform,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &EmailOtpSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &EmailOtpSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/otp/email/send",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &EmailOtpSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &sendOTPOnEmailResponse)
        if err != nil {
           
             return &EmailOtpSuccess{}, common.NewFDKError(err.Error())
            
        }
        return sendOTPOnEmailResponse, nil
    }
          
    
  
    
    // VerifyEmailOTP Verify OTP on email
    func (us *User)  VerifyEmailOTP(Platform string, body  VerifyOtpRequestSchema) (*VerifyOtpSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            verifyEmailOTPResponse *VerifyOtpSuccess
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Platform string  `url:"platform"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Platform : Platform,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &VerifyOtpSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &VerifyOtpSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/otp/email/verify",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &VerifyOtpSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &verifyEmailOTPResponse)
        if err != nil {
           
             return &VerifyOtpSuccess{}, common.NewFDKError(err.Error())
            
        }
        return verifyEmailOTPResponse, nil
    }
          
    
  
    
    // GetLoggedInUser Get logged in user
    func (us *User)  GetLoggedInUser() (*UserSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getLoggedInUserResponse *UserSchema
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "get",
            "/service/application/user/authentication/v1.0/session",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &UserSchema{}, err
            
	    }
        err = json.Unmarshal(response, &getLoggedInUserResponse)
        if err != nil {
           
             return &UserSchema{}, common.NewFDKError(err.Error())
            
        }
        return getLoggedInUserResponse, nil
    }
          
    
  
    
    // GetListOfActiveSessions Get list of sessions
    func (us *User)  GetListOfActiveSessions() (*SessionListSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getListOfActiveSessionsResponse *SessionListSuccess
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "get",
            "/service/application/user/authentication/v1.0/sessions",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &SessionListSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &getListOfActiveSessionsResponse)
        if err != nil {
           
             return &SessionListSuccess{}, common.NewFDKError(err.Error())
            
        }
        return getListOfActiveSessionsResponse, nil
    }
          
    
  
    
    // GetPlatformConfig Get platform config
    func (us *User)  GetPlatformConfig(Name string) (*PlatformSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getPlatformConfigResponse *PlatformSchema
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Name string  `url:"name"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Name : Name,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "get",
            "/service/application/user/platform/v1.0/config",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &PlatformSchema{}, err
            
	    }
        err = json.Unmarshal(response, &getPlatformConfigResponse)
        if err != nil {
           
             return &PlatformSchema{}, common.NewFDKError(err.Error())
            
        }
        return getPlatformConfigResponse, nil
    }
          
    
  
    
    // UpdateProfile Edit Profile Details
    func (us *User)  UpdateProfile(Platform string, body  EditProfileRequestSchema) (*LoginSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateProfileResponse *LoginSuccess
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Platform string  `url:"platform"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Platform : Platform,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &LoginSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &LoginSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/profile/v1.0/detail",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &LoginSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &updateProfileResponse)
        if err != nil {
           
             return &LoginSuccess{}, common.NewFDKError(err.Error())
            
        }
        return updateProfileResponse, nil
    }
          
    
  
    
    // AddMobileNumber Add mobile number to profile
    func (us *User)  AddMobileNumber(Platform string, body  EditMobileRequestSchema) (*VerifyMobileOTPSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            addMobileNumberResponse *VerifyMobileOTPSuccess
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Platform string  `url:"platform"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Platform : Platform,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &VerifyMobileOTPSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &VerifyMobileOTPSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "put",
            "/service/application/user/profile/v1.0/mobile",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &VerifyMobileOTPSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &addMobileNumberResponse)
        if err != nil {
           
             return &VerifyMobileOTPSuccess{}, common.NewFDKError(err.Error())
            
        }
        return addMobileNumberResponse, nil
    }
          
    
  
    
    // DeleteMobileNumber Delete mobile number from profile
    func (us *User)  DeleteMobileNumber(Platform string, Active bool, Primary bool, Verified bool, CountryCode string, Phone string) (*LoginSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            deleteMobileNumberResponse *LoginSuccess
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Platform string  `url:"platform"` 
          Active bool  `url:"active"` 
          Primary bool  `url:"primary"` 
          Verified bool  `url:"verified"` 
          CountryCode string  `url:"country_code"` 
          Phone string  `url:"phone"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Platform : Platform,
           Active : Active,
           Primary : Primary,
           Verified : Verified,
           CountryCode : CountryCode,
           Phone : Phone,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "delete",
            fmt.Sprintf("/service/application/user/profile/v1.0/mobile",),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &LoginSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &deleteMobileNumberResponse)
        if err != nil {
           
             return &LoginSuccess{}, common.NewFDKError(err.Error())
            
        }
        return deleteMobileNumberResponse, nil
    }
          
    
  
    
    // SetMobileNumberAsPrimary Set mobile as primary
    func (us *User)  SetMobileNumberAsPrimary(body  SendVerificationLinkMobileRequestSchema) (*LoginSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            setMobileNumberAsPrimaryResponse *LoginSuccess
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &LoginSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &LoginSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/profile/v1.0/mobile/primary",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &LoginSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &setMobileNumberAsPrimaryResponse)
        if err != nil {
           
             return &LoginSuccess{}, common.NewFDKError(err.Error())
            
        }
        return setMobileNumberAsPrimaryResponse, nil
    }
          
    
  
    
    // SendVerificationLinkToMobile Send verification link to mobile
    func (us *User)  SendVerificationLinkToMobile(Platform string, body  SendVerificationLinkMobileRequestSchema) (*SendMobileVerifyLinkSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            sendVerificationLinkToMobileResponse *SendMobileVerifyLinkSuccess
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Platform string  `url:"platform"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Platform : Platform,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &SendMobileVerifyLinkSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &SendMobileVerifyLinkSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/profile/v1.0/mobile/link/send",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &SendMobileVerifyLinkSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &sendVerificationLinkToMobileResponse)
        if err != nil {
           
             return &SendMobileVerifyLinkSuccess{}, common.NewFDKError(err.Error())
            
        }
        return sendVerificationLinkToMobileResponse, nil
    }
          
    
  
    
    // AddEmail Add email to profile
    func (us *User)  AddEmail(Platform string, body  EditEmailRequestSchema) (*VerifyEmailOTPSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            addEmailResponse *VerifyEmailOTPSuccess
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Platform string  `url:"platform"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Platform : Platform,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &VerifyEmailOTPSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &VerifyEmailOTPSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "put",
            "/service/application/user/profile/v1.0/email",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &VerifyEmailOTPSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &addEmailResponse)
        if err != nil {
           
             return &VerifyEmailOTPSuccess{}, common.NewFDKError(err.Error())
            
        }
        return addEmailResponse, nil
    }
          
    
  
    
    // DeleteEmail Delete email from profile
    func (us *User)  DeleteEmail(Platform string, Active bool, Primary bool, Verified bool, Email string) (*LoginSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            deleteEmailResponse *LoginSuccess
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Platform string  `url:"platform"` 
          Active bool  `url:"active"` 
          Primary bool  `url:"primary"` 
          Verified bool  `url:"verified"` 
          Email string  `url:"email"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Platform : Platform,
           Active : Active,
           Primary : Primary,
           Verified : Verified,
           Email : Email,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "delete",
            fmt.Sprintf("/service/application/user/profile/v1.0/email",),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &LoginSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &deleteEmailResponse)
        if err != nil {
           
             return &LoginSuccess{}, common.NewFDKError(err.Error())
            
        }
        return deleteEmailResponse, nil
    }
          
    
  
    
    // SetEmailAsPrimary Set email as primary
    func (us *User)  SetEmailAsPrimary(body  EditEmailRequestSchema) (*LoginSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            setEmailAsPrimaryResponse *LoginSuccess
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &LoginSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &LoginSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/profile/v1.0/email/primary",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &LoginSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &setEmailAsPrimaryResponse)
        if err != nil {
           
             return &LoginSuccess{}, common.NewFDKError(err.Error())
            
        }
        return setEmailAsPrimaryResponse, nil
    }
          
    
  
    
    // SendVerificationLinkToEmail Send verification link to email
    func (us *User)  SendVerificationLinkToEmail(Platform string, body  EditEmailRequestSchema) (*SendEmailVerifyLinkSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            sendVerificationLinkToEmailResponse *SendEmailVerifyLinkSuccess
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Platform string  `url:"platform"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Platform : Platform,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &SendEmailVerifyLinkSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &SendEmailVerifyLinkSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/profile/v1.0/email/link/send",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &SendEmailVerifyLinkSuccess{}, err
            
	    }
        err = json.Unmarshal(response, &sendVerificationLinkToEmailResponse)
        if err != nil {
           
             return &SendEmailVerifyLinkSuccess{}, common.NewFDKError(err.Error())
            
        }
        return sendVerificationLinkToEmailResponse, nil
    }
          
    

    // Content ...
    type Content struct {
        config *AppConfig
    }
    // NewContent ...
    func NewContent(config *AppConfig) *Content {
        return &Content{config}
    }
    
  
    
    // GetAnnouncements Get live announcements
    func (co *Content)  GetAnnouncements() (*AnnouncementsResponseSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getAnnouncementsResponse *AnnouncementsResponseSchema
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/content/v1.0/announcements",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &AnnouncementsResponseSchema{}, err
            
	    }
        err = json.Unmarshal(response, &getAnnouncementsResponse)
        if err != nil {
           
             return &AnnouncementsResponseSchema{}, common.NewFDKError(err.Error())
            
        }
        return getAnnouncementsResponse, nil
    }
          
    
  
    
    // GetBlog Get Blog by slug
    func (co *Content)  GetBlog(Slug string) (*CustomBlogSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getBlogResponse *CustomBlogSchema
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/application/content/v1.0/blogs/%s",Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CustomBlogSchema{}, err
            
	    }
        err = json.Unmarshal(response, &getBlogResponse)
        if err != nil {
           
             return &CustomBlogSchema{}, common.NewFDKError(err.Error())
            
        }
        return getBlogResponse, nil
    }
          
    
  
    
    // GetBlogs Get blogs
    func (co *Content)  GetBlogs(PageNo int, PageSize int) (*BlogGetResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getBlogsResponse *BlogGetResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          PageNo int  `url:"page_no"` 
          PageSize int  `url:"page_size"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           PageNo : PageNo,
           PageSize : PageSize,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/content/v1.0/blogs/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &BlogGetResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getBlogsResponse)
        if err != nil {
           
             return &BlogGetResponse{}, common.NewFDKError(err.Error())
            
        }
        return getBlogsResponse, nil
    }
          
            
            
                
                    
                    
                
                    
                        
                    
                    
                
            
            // GetBlogsPaginator Get blogs  
            func (co *Content)  GetBlogsPaginator(PageSize int ) *common.Paginator {
                paginator := common.NewPaginator("number")
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetBlogs(paginator.PageNo, PageSize)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    
  
    
    // GetFaqs Get frequently asked questions
    func (co *Content)  GetFaqs() (*FaqResponseSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getFaqsResponse *FaqResponseSchema
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/content/v1.0/faq",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &FaqResponseSchema{}, err
            
	    }
        err = json.Unmarshal(response, &getFaqsResponse)
        if err != nil {
           
             return &FaqResponseSchema{}, common.NewFDKError(err.Error())
            
        }
        return getFaqsResponse, nil
    }
          
    
  
    
    // GetFaqCategories Get FAQ categories list
    func (co *Content)  GetFaqCategories() (*GetFaqCategoriesSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getFaqCategoriesResponse *GetFaqCategoriesSchema
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/content/v1.0/faq/categories",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &GetFaqCategoriesSchema{}, err
            
	    }
        err = json.Unmarshal(response, &getFaqCategoriesResponse)
        if err != nil {
           
             return &GetFaqCategoriesSchema{}, common.NewFDKError(err.Error())
            
        }
        return getFaqCategoriesResponse, nil
    }
          
    
  
    
    // GetFaqBySlug Get frequently asked question
    func (co *Content)  GetFaqBySlug(Slug string) (*FaqSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getFaqBySlugResponse *FaqSchema
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/application/content/v1.0/faq/%s",Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &FaqSchema{}, err
            
	    }
        err = json.Unmarshal(response, &getFaqBySlugResponse)
        if err != nil {
           
             return &FaqSchema{}, common.NewFDKError(err.Error())
            
        }
        return getFaqBySlugResponse, nil
    }
          
    
  
    
    // GetFaqCategoryBySlug Get FAQ category by slug
    func (co *Content)  GetFaqCategoryBySlug(Slug string) (*GetFaqCategoryBySlugSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getFaqCategoryBySlugResponse *GetFaqCategoryBySlugSchema
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/application/content/v1.0/faq/category/%s",Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &GetFaqCategoryBySlugSchema{}, err
            
	    }
        err = json.Unmarshal(response, &getFaqCategoryBySlugResponse)
        if err != nil {
           
             return &GetFaqCategoryBySlugSchema{}, common.NewFDKError(err.Error())
            
        }
        return getFaqCategoryBySlugResponse, nil
    }
          
    
  
    
    // GetFaqsByCategorySlug Get FAQs of a Faq Category slug
    func (co *Content)  GetFaqsByCategorySlug(Slug string) (*GetFaqSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getFaqsByCategorySlugResponse *GetFaqSchema
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/application/content/v1.0/faq/category/%s/faqs",Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &GetFaqSchema{}, err
            
	    }
        err = json.Unmarshal(response, &getFaqsByCategorySlugResponse)
        if err != nil {
           
             return &GetFaqSchema{}, common.NewFDKError(err.Error())
            
        }
        return getFaqsByCategorySlugResponse, nil
    }
          
    
  
    
    // GetLandingPage Get landing page
    func (co *Content)  GetLandingPage() (*LandingPageSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getLandingPageResponse *LandingPageSchema
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/content/v1.0/landing-page",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &LandingPageSchema{}, err
            
	    }
        err = json.Unmarshal(response, &getLandingPageResponse)
        if err != nil {
           
             return &LandingPageSchema{}, common.NewFDKError(err.Error())
            
        }
        return getLandingPageResponse, nil
    }
          
    
  
    
    // GetLegalInformation Get legal information
    func (co *Content)  GetLegalInformation() (*ApplicationLegal, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getLegalInformationResponse *ApplicationLegal
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/content/v1.0/legal",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ApplicationLegal{}, err
            
	    }
        err = json.Unmarshal(response, &getLegalInformationResponse)
        if err != nil {
           
             return &ApplicationLegal{}, common.NewFDKError(err.Error())
            
        }
        return getLegalInformationResponse, nil
    }
          
    
  
    
    // GetNavigations Get navigation
    func (co *Content)  GetNavigations(PageNo int, PageSize int) (*NavigationGetResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getNavigationsResponse *NavigationGetResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          PageNo int  `url:"page_no"` 
          PageSize int  `url:"page_size"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           PageNo : PageNo,
           PageSize : PageSize,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/content/v1.0/navigations/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &NavigationGetResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getNavigationsResponse)
        if err != nil {
           
             return &NavigationGetResponse{}, common.NewFDKError(err.Error())
            
        }
        return getNavigationsResponse, nil
    }
          
            
            
                
                    
                    
                
                    
                        
                    
                    
                
            
            // GetNavigationsPaginator Get navigation  
            func (co *Content)  GetNavigationsPaginator(PageSize int ) *common.Paginator {
                paginator := common.NewPaginator("number")
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetNavigations(paginator.PageNo, PageSize)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    
  
    
    // GetPage Get Page by slug
    func (co *Content)  GetPage(Slug string) (*CustomPageSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getPageResponse *CustomPageSchema
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/application/content/v1.0/pages/%s",Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CustomPageSchema{}, err
            
	    }
        err = json.Unmarshal(response, &getPageResponse)
        if err != nil {
           
             return &CustomPageSchema{}, common.NewFDKError(err.Error())
            
        }
        return getPageResponse, nil
    }
          
    
  
    
    // GetPages Get pages
    func (co *Content)  GetPages(PageNo int, PageSize int) (*PageGetResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getPagesResponse *PageGetResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          PageNo int  `url:"page_no"` 
          PageSize int  `url:"page_size"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           PageNo : PageNo,
           PageSize : PageSize,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/content/v1.0/pages/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &PageGetResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getPagesResponse)
        if err != nil {
           
             return &PageGetResponse{}, common.NewFDKError(err.Error())
            
        }
        return getPagesResponse, nil
    }
          
            
            
                
                    
                    
                
                    
                        
                    
                    
                
            
            // GetPagesPaginator Get pages  
            func (co *Content)  GetPagesPaginator(PageSize int ) *common.Paginator {
                paginator := common.NewPaginator("number")
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetPages(paginator.PageNo, PageSize)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    
  
    
    // GetSEOConfiguration Get seo of application
    func (co *Content)  GetSEOConfiguration() (*SeoComponent, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getSEOConfigurationResponse *SeoComponent
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/content/v1.0/seo",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &SeoComponent{}, err
            
	    }
        err = json.Unmarshal(response, &getSEOConfigurationResponse)
        if err != nil {
           
             return &SeoComponent{}, common.NewFDKError(err.Error())
            
        }
        return getSEOConfigurationResponse, nil
    }
          
    
  
    
    // GetSlideshows Get slideshows
    func (co *Content)  GetSlideshows(PageNo int, PageSize int) (*SlideshowGetResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getSlideshowsResponse *SlideshowGetResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          PageNo int  `url:"page_no"` 
          PageSize int  `url:"page_size"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           PageNo : PageNo,
           PageSize : PageSize,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/content/v1.0/slideshow/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &SlideshowGetResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getSlideshowsResponse)
        if err != nil {
           
             return &SlideshowGetResponse{}, common.NewFDKError(err.Error())
            
        }
        return getSlideshowsResponse, nil
    }
          
            
            
                
                    
                    
                
                    
                        
                    
                    
                
            
            // GetSlideshowsPaginator Get slideshows  
            func (co *Content)  GetSlideshowsPaginator(PageSize int ) *common.Paginator {
                paginator := common.NewPaginator("number")
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetSlideshows(paginator.PageNo, PageSize)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    
  
    
    // GetSlideshow Get slideshow by slug
    func (co *Content)  GetSlideshow(Slug string) (*SlideshowSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getSlideshowResponse *SlideshowSchema
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/application/content/v1.0/slideshow/%s",Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &SlideshowSchema{}, err
            
	    }
        err = json.Unmarshal(response, &getSlideshowResponse)
        if err != nil {
           
             return &SlideshowSchema{}, common.NewFDKError(err.Error())
            
        }
        return getSlideshowResponse, nil
    }
          
    
  
    
    // GetSupportInformation Get support information
    func (co *Content)  GetSupportInformation() (*Support, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getSupportInformationResponse *Support
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/content/v1.0/support",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &Support{}, err
            
	    }
        err = json.Unmarshal(response, &getSupportInformationResponse)
        if err != nil {
           
             return &Support{}, common.NewFDKError(err.Error())
            
        }
        return getSupportInformationResponse, nil
    }
          
    
  
    
    // GetTags Get Tags for application
    func (co *Content)  GetTags() (*TagsSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getTagsResponse *TagsSchema
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/content/v1.0/tags",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &TagsSchema{}, err
            
	    }
        err = json.Unmarshal(response, &getTagsResponse)
        if err != nil {
           
             return &TagsSchema{}, common.NewFDKError(err.Error())
            
        }
        return getTagsResponse, nil
    }
          
    

    // Communication ...
    type Communication struct {
        config *AppConfig
    }
    // NewCommunication ...
    func NewCommunication(config *AppConfig) *Communication {
        return &Communication{config}
    }
    
  
    
    // GetCommunicationConsent Get communication consent
    func (co *Communication)  GetCommunicationConsent() (*CommunicationConsent, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getCommunicationConsentResponse *CommunicationConsent
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/communication/v1.0/consent",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CommunicationConsent{}, err
            
	    }
        err = json.Unmarshal(response, &getCommunicationConsentResponse)
        if err != nil {
           
             return &CommunicationConsent{}, common.NewFDKError(err.Error())
            
        }
        return getCommunicationConsentResponse, nil
    }
          
    
  
    
    // UpsertCommunicationConsent Upsert communication consent
    func (co *Communication)  UpsertCommunicationConsent(body  CommunicationConsentReq) (*CommunicationConsentRes, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            upsertCommunicationConsentResponse *CommunicationConsentRes
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &CommunicationConsentRes{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &CommunicationConsentRes{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            "/service/application/communication/v1.0/consent",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CommunicationConsentRes{}, err
            
	    }
        err = json.Unmarshal(response, &upsertCommunicationConsentResponse)
        if err != nil {
           
             return &CommunicationConsentRes{}, common.NewFDKError(err.Error())
            
        }
        return upsertCommunicationConsentResponse, nil
    }
          
    
  
    
    // UpsertAppPushtoken Upsert push token of a user
    func (co *Communication)  UpsertAppPushtoken(body  PushtokenReq) (*PushtokenRes, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            upsertAppPushtokenResponse *PushtokenRes
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &PushtokenRes{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &PushtokenRes{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            "/service/application/communication/v1.0/pn-token",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &PushtokenRes{}, err
            
	    }
        err = json.Unmarshal(response, &upsertAppPushtokenResponse)
        if err != nil {
           
             return &PushtokenRes{}, common.NewFDKError(err.Error())
            
        }
        return upsertAppPushtokenResponse, nil
    }
          
    

    // Share ...
    type Share struct {
        config *AppConfig
    }
    // NewShare ...
    func NewShare(config *AppConfig) *Share {
        return &Share{config}
    }
    
  
    
    // GetApplicationQRCode Create application QR Code
    func (sh *Share)  GetApplicationQRCode() (*QRCodeResp, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getApplicationQRCodeResponse *QRCodeResp
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            sh.config,
            "post",
            "/service/application/share/v1.0/qr/",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &QRCodeResp{}, err
            
	    }
        err = json.Unmarshal(response, &getApplicationQRCodeResponse)
        if err != nil {
           
             return &QRCodeResp{}, common.NewFDKError(err.Error())
            
        }
        return getApplicationQRCodeResponse, nil
    }
          
    
  
    
    // GetProductQRCodeBySlug Create product QR Code
    func (sh *Share)  GetProductQRCodeBySlug(Slug string) (*QRCodeResp, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getProductQRCodeBySlugResponse *QRCodeResp
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            sh.config,
            "post",
            fmt.Sprintf("/service/application/share/v1.0/qr/products/%s/",Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &QRCodeResp{}, err
            
	    }
        err = json.Unmarshal(response, &getProductQRCodeBySlugResponse)
        if err != nil {
           
             return &QRCodeResp{}, common.NewFDKError(err.Error())
            
        }
        return getProductQRCodeBySlugResponse, nil
    }
          
    
  
    
    // GetCollectionQRCodeBySlug Create collection QR Code
    func (sh *Share)  GetCollectionQRCodeBySlug(Slug string) (*QRCodeResp, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getCollectionQRCodeBySlugResponse *QRCodeResp
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            sh.config,
            "post",
            fmt.Sprintf("/service/application/share/v1.0/qr/collection/%s/",Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &QRCodeResp{}, err
            
	    }
        err = json.Unmarshal(response, &getCollectionQRCodeBySlugResponse)
        if err != nil {
           
             return &QRCodeResp{}, common.NewFDKError(err.Error())
            
        }
        return getCollectionQRCodeBySlugResponse, nil
    }
          
    
  
    
    // GetUrlQRCode Create url QR Code
    func (sh *Share)  GetUrlQRCode(URL string) (*QRCodeResp, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getUrlQRCodeResponse *QRCodeResp
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          URL string  `url:"url"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           URL : URL,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            sh.config,
            "post",
            fmt.Sprintf("/service/application/share/v1.0/qr/url/",),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &QRCodeResp{}, err
            
	    }
        err = json.Unmarshal(response, &getUrlQRCodeResponse)
        if err != nil {
           
             return &QRCodeResp{}, common.NewFDKError(err.Error())
            
        }
        return getUrlQRCodeResponse, nil
    }
          
    
  
    
    // CreateShortLink Create short link
    func (sh *Share)  CreateShortLink(body  ShortLinkReq) (*ShortLinkRes, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            createShortLinkResponse *ShortLinkRes
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &ShortLinkRes{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &ShortLinkRes{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            sh.config,
            "post",
            "/service/application/share/v1.0/links/short-link/",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ShortLinkRes{}, err
            
	    }
        err = json.Unmarshal(response, &createShortLinkResponse)
        if err != nil {
           
             return &ShortLinkRes{}, common.NewFDKError(err.Error())
            
        }
        return createShortLinkResponse, nil
    }
          
    
  
    
    // GetShortLinkByHash Get short link by hash
    func (sh *Share)  GetShortLinkByHash(Hash string) (*ShortLinkRes, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getShortLinkByHashResponse *ShortLinkRes
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            sh.config,
            "get",
            fmt.Sprintf("/service/application/share/v1.0/links/short-link/%s/",Hash),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ShortLinkRes{}, err
            
	    }
        err = json.Unmarshal(response, &getShortLinkByHashResponse)
        if err != nil {
           
             return &ShortLinkRes{}, common.NewFDKError(err.Error())
            
        }
        return getShortLinkByHashResponse, nil
    }
          
    
  
    
    // GetOriginalShortLinkByHash Get original link by hash
    func (sh *Share)  GetOriginalShortLinkByHash(Hash string) (*ShortLinkRes, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getOriginalShortLinkByHashResponse *ShortLinkRes
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            sh.config,
            "get",
            fmt.Sprintf("/service/application/share/v1.0/links/short-link/%s/original/",Hash),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ShortLinkRes{}, err
            
	    }
        err = json.Unmarshal(response, &getOriginalShortLinkByHashResponse)
        if err != nil {
           
             return &ShortLinkRes{}, common.NewFDKError(err.Error())
            
        }
        return getOriginalShortLinkByHashResponse, nil
    }
          
    

    // FileStorage ...
    type FileStorage struct {
        config *AppConfig
    }
    // NewFileStorage ...
    func NewFileStorage(config *AppConfig) *FileStorage {
        return &FileStorage{config}
    }
    
  
    
    // StartUpload This operation initiates upload and returns storage link which is valid for 30 Minutes. You can use that storage link to make subsequent upload request with file buffer or blob.
    func (fi *FileStorage)  StartUpload(Namespace string, body  StartRequest) (*StartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            startUploadResponse *StartResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &StartResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &StartResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            fi.config,
            "post",
            fmt.Sprintf("/service/application/assets/v1.0/namespaces/%s/upload/start/",Namespace),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &StartResponse{}, err
            
	    }
        err = json.Unmarshal(response, &startUploadResponse)
        if err != nil {
           
             return &StartResponse{}, common.NewFDKError(err.Error())
            
        }
        return startUploadResponse, nil
    }
          
    
  
    
    // CompleteUpload This will complete the upload process. After successfully uploading file, you can call this operation to complete the upload process.
    func (fi *FileStorage)  CompleteUpload(Namespace string, body  StartResponse) (*CompleteResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            completeUploadResponse *CompleteResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &CompleteResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &CompleteResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            fi.config,
            "post",
            fmt.Sprintf("/service/application/assets/v1.0/namespaces/%s/upload/complete/",Namespace),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CompleteResponse{}, err
            
	    }
        err = json.Unmarshal(response, &completeUploadResponse)
        if err != nil {
           
             return &CompleteResponse{}, common.NewFDKError(err.Error())
            
        }
        return completeUploadResponse, nil
    }
          
    

    // Configuration ...
    type Configuration struct {
        config *AppConfig
    }
    // NewConfiguration ...
    func NewConfiguration(config *AppConfig) *Configuration {
        return &Configuration{config}
    }
    
  
    
    // GetApplication Get current application details
    func (co *Configuration)  GetApplication() (*Application, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getApplicationResponse *Application
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/configuration/v1.0/application",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &Application{}, err
            
	    }
        err = json.Unmarshal(response, &getApplicationResponse)
        if err != nil {
           
             return &Application{}, common.NewFDKError(err.Error())
            
        }
        return getApplicationResponse, nil
    }
          
    
  
    
    // GetOwnerInfo Get application, owner and seller information
    func (co *Configuration)  GetOwnerInfo() (*ApplicationAboutResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getOwnerInfoResponse *ApplicationAboutResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/configuration/v1.0/about",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ApplicationAboutResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getOwnerInfoResponse)
        if err != nil {
           
             return &ApplicationAboutResponse{}, common.NewFDKError(err.Error())
            
        }
        return getOwnerInfoResponse, nil
    }
          
    
  
    
    // GetBasicDetails Get basic application details
    func (co *Configuration)  GetBasicDetails() (*ApplicationDetail, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getBasicDetailsResponse *ApplicationDetail
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/configuration/v1.0/detail",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ApplicationDetail{}, err
            
	    }
        err = json.Unmarshal(response, &getBasicDetailsResponse)
        if err != nil {
           
             return &ApplicationDetail{}, common.NewFDKError(err.Error())
            
        }
        return getBasicDetailsResponse, nil
    }
          
    
  
    
    // GetIntegrationTokens Get integration tokens
    func (co *Configuration)  GetIntegrationTokens() (*TokenResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getIntegrationTokensResponse *TokenResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/configuration/v1.0/token",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &TokenResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getIntegrationTokensResponse)
        if err != nil {
           
             return &TokenResponse{}, common.NewFDKError(err.Error())
            
        }
        return getIntegrationTokensResponse, nil
    }
          
    
  
    
    // GetOrderingStores Get deployment meta stores
    func (co *Configuration)  GetOrderingStores(PageNo int, PageSize int, Q string) (*OrderingStores, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getOrderingStoresResponse *OrderingStores
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          PageNo int  `url:"page_no"` 
          PageSize int  `url:"page_size"` 
          Q string  `url:"q"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           PageNo : PageNo,
           PageSize : PageSize,
           Q : Q,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/configuration/v1.0/ordering-store/stores",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &OrderingStores{}, err
            
	    }
        err = json.Unmarshal(response, &getOrderingStoresResponse)
        if err != nil {
           
             return &OrderingStores{}, common.NewFDKError(err.Error())
            
        }
        return getOrderingStoresResponse, nil
    }
          
            
            
                
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
            
            // GetOrderingStoresPaginator Get deployment meta stores  
            func (co *Configuration)  GetOrderingStoresPaginator(PageSize int , Q string ) *common.Paginator {
                paginator := common.NewPaginator("number")
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetOrderingStores(paginator.PageNo, PageSize, Q)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    
  
    
    // GetFeatures Get features of application
    func (co *Configuration)  GetFeatures() (*AppFeatureResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getFeaturesResponse *AppFeatureResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/configuration/v1.0/feature",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &AppFeatureResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getFeaturesResponse)
        if err != nil {
           
             return &AppFeatureResponse{}, common.NewFDKError(err.Error())
            
        }
        return getFeaturesResponse, nil
    }
          
    
  
    
    // GetContactInfo Get application information
    func (co *Configuration)  GetContactInfo() (*ApplicationInformation, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getContactInfoResponse *ApplicationInformation
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/configuration/v1.0/information",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ApplicationInformation{}, err
            
	    }
        err = json.Unmarshal(response, &getContactInfoResponse)
        if err != nil {
           
             return &ApplicationInformation{}, common.NewFDKError(err.Error())
            
        }
        return getContactInfoResponse, nil
    }
          
    
  
    
    // GetCurrencies Get application enabled currencies
    func (co *Configuration)  GetCurrencies() (*CurrenciesResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getCurrenciesResponse *CurrenciesResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/configuration/v1.0/currencies",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CurrenciesResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getCurrenciesResponse)
        if err != nil {
           
             return &CurrenciesResponse{}, common.NewFDKError(err.Error())
            
        }
        return getCurrenciesResponse, nil
    }
          
    
  
    
    // GetCurrencyById Get currency by id
    func (co *Configuration)  GetCurrencyById(ID string) (*Currency, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getCurrencyByIdResponse *Currency
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/application/configuration/v1.0/currency/%s",ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &Currency{}, err
            
	    }
        err = json.Unmarshal(response, &getCurrencyByIdResponse)
        if err != nil {
           
             return &Currency{}, common.NewFDKError(err.Error())
            
        }
        return getCurrencyByIdResponse, nil
    }
          
    
  
    
    // GetLanguages Get list of languages
    func (co *Configuration)  GetLanguages() (*LanguageResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getLanguagesResponse *LanguageResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/configuration/v1.0/languages",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &LanguageResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getLanguagesResponse)
        if err != nil {
           
             return &LanguageResponse{}, common.NewFDKError(err.Error())
            
        }
        return getLanguagesResponse, nil
    }
          
    
  
    
    // GetOrderingStoreCookie Get ordering store signed cookie on selection of ordering store. This will be used by cart service to verify coupon against selected ordering store in cart.
    func (co *Configuration)  GetOrderingStoreCookie(body  OrderingStoreSelectRequest) (*SuccessMessageResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getOrderingStoreCookieResponse *SuccessMessageResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &SuccessMessageResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &SuccessMessageResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "post",
            "/service/application/configuration/v1.0/ordering-store/select",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &SuccessMessageResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getOrderingStoreCookieResponse)
        if err != nil {
           
             return &SuccessMessageResponse{}, common.NewFDKError(err.Error())
            
        }
        return getOrderingStoreCookieResponse, nil
    }
          
    
  
    
    // RemoveOrderingStoreCookie Unset ordering store signed cookie on change of sales channel selection via domain in universal fynd store app.
    func (co *Configuration)  RemoveOrderingStoreCookie() (*SuccessMessageResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            removeOrderingStoreCookieResponse *SuccessMessageResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "delete",
            "/service/application/configuration/v1.0/ordering-store/select",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &SuccessMessageResponse{}, err
            
	    }
        err = json.Unmarshal(response, &removeOrderingStoreCookieResponse)
        if err != nil {
           
             return &SuccessMessageResponse{}, common.NewFDKError(err.Error())
            
        }
        return removeOrderingStoreCookieResponse, nil
    }
          
    
  
    
    // GetAppStaffs Get Staff List.
    func (co *Configuration)  GetAppStaffs(OrderIncent bool, OrderingStore int, User string) (*AppStaffResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getAppStaffsResponse *AppStaffResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          OrderIncent bool  `url:"order_incent"` 
          OrderingStore int  `url:"ordering_store"` 
          User string  `url:"user"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           OrderIncent : OrderIncent,
           OrderingStore : OrderingStore,
           User : User,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/configuration/v1.0/staff",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &AppStaffResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getAppStaffsResponse)
        if err != nil {
           
             return &AppStaffResponse{}, common.NewFDKError(err.Error())
            
        }
        return getAppStaffsResponse, nil
    }
          
    

    // Payment ...
    type Payment struct {
        config *AppConfig
    }
    // NewPayment ...
    func NewPayment(config *AppConfig) *Payment {
        return &Payment{config}
    }
    
  
    
    // GetAggregatorsConfig Get payment gateway keys
    func (pa *Payment)  GetAggregatorsConfig(XAPIToken string, Refresh bool) (*AggregatorsConfigDetailResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getAggregatorsConfigResponse *AggregatorsConfigDetailResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Refresh bool  `url:"refresh"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Refresh : Refresh,
        }
        
        
        //Adding extra headers
        var xHeaders = make(map[string]string) 
        
        xHeaders["x-api-token"]=  XAPIToken;
        
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "get",
            fmt.Sprintf("/service/application/payment/v1.0/config/aggregators/key",),
            xHeaders,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &AggregatorsConfigDetailResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getAggregatorsConfigResponse)
        if err != nil {
           
             return &AggregatorsConfigDetailResponse{}, common.NewFDKError(err.Error())
            
        }
        return getAggregatorsConfigResponse, nil
    }
          
    
  
    
    // AttachCardToCustomer Attach a saved card to customer.
    func (pa *Payment)  AttachCardToCustomer(body  AttachCardRequest) (*AttachCardsResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            attachCardToCustomerResponse *AttachCardsResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &AttachCardsResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &AttachCardsResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "post",
            "/service/application/payment/v1.0/card/attach",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &AttachCardsResponse{}, err
            
	    }
        err = json.Unmarshal(response, &attachCardToCustomerResponse)
        if err != nil {
           
             return &AttachCardsResponse{}, common.NewFDKError(err.Error())
            
        }
        return attachCardToCustomerResponse, nil
    }
          
    
  
    
    // GetActiveCardAggregator Fetch active payment gateway for card
    func (pa *Payment)  GetActiveCardAggregator(Refresh bool) (*ActiveCardPaymentGatewayResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getActiveCardAggregatorResponse *ActiveCardPaymentGatewayResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Refresh bool  `url:"refresh"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Refresh : Refresh,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "get",
            "/service/application/payment/v1.0/card/aggregator",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ActiveCardPaymentGatewayResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getActiveCardAggregatorResponse)
        if err != nil {
           
             return &ActiveCardPaymentGatewayResponse{}, common.NewFDKError(err.Error())
            
        }
        return getActiveCardAggregatorResponse, nil
    }
          
    
  
    
    // GetActiveUserCards Fetch the list of saved cards of user.
    func (pa *Payment)  GetActiveUserCards(ForceRefresh bool) (*ListCardsResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getActiveUserCardsResponse *ListCardsResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          ForceRefresh bool  `url:"force_refresh"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           ForceRefresh : ForceRefresh,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "get",
            "/service/application/payment/v1.0/cards",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ListCardsResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getActiveUserCardsResponse)
        if err != nil {
           
             return &ListCardsResponse{}, common.NewFDKError(err.Error())
            
        }
        return getActiveUserCardsResponse, nil
    }
          
    
  
    
    // DeleteUserCard Delete an user card.
    func (pa *Payment)  DeleteUserCard(body  DeletehCardRequest) (*DeleteCardsResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            deleteUserCardResponse *DeleteCardsResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &DeleteCardsResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &DeleteCardsResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "post",
            "/service/application/payment/v1.0/card/remove",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &DeleteCardsResponse{}, err
            
	    }
        err = json.Unmarshal(response, &deleteUserCardResponse)
        if err != nil {
           
             return &DeleteCardsResponse{}, common.NewFDKError(err.Error())
            
        }
        return deleteUserCardResponse, nil
    }
          
    
  
    
    // VerifyCustomerForPayment Validate customer for payment.
    func (pa *Payment)  VerifyCustomerForPayment(body  ValidateCustomerRequest) (*ValidateCustomerResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            verifyCustomerForPaymentResponse *ValidateCustomerResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &ValidateCustomerResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &ValidateCustomerResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "post",
            "/service/application/payment/v1.0/payment/customer/validation",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ValidateCustomerResponse{}, err
            
	    }
        err = json.Unmarshal(response, &verifyCustomerForPaymentResponse)
        if err != nil {
           
             return &ValidateCustomerResponse{}, common.NewFDKError(err.Error())
            
        }
        return verifyCustomerForPaymentResponse, nil
    }
          
    
  
    
    // VerifyAndChargePayment Verify and charge payment
    func (pa *Payment)  VerifyAndChargePayment(body  ChargeCustomerRequest) (*ChargeCustomerResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            verifyAndChargePaymentResponse *ChargeCustomerResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &ChargeCustomerResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &ChargeCustomerResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "post",
            "/service/application/payment/v1.0/payment/confirm/charge",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ChargeCustomerResponse{}, err
            
	    }
        err = json.Unmarshal(response, &verifyAndChargePaymentResponse)
        if err != nil {
           
             return &ChargeCustomerResponse{}, common.NewFDKError(err.Error())
            
        }
        return verifyAndChargePaymentResponse, nil
    }
          
    
  
    
    // InitialisePayment Payment Initialisation server to server for UPI and BharatQR.
    func (pa *Payment)  InitialisePayment(body  PaymentInitializationRequest) (*PaymentInitializationResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            initialisePaymentResponse *PaymentInitializationResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &PaymentInitializationResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &PaymentInitializationResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "post",
            "/service/application/payment/v1.0/payment/request",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &PaymentInitializationResponse{}, err
            
	    }
        err = json.Unmarshal(response, &initialisePaymentResponse)
        if err != nil {
           
             return &PaymentInitializationResponse{}, common.NewFDKError(err.Error())
            
        }
        return initialisePaymentResponse, nil
    }
          
    
  
    
    // CheckAndUpdatePaymentStatus Continous polling to check status of payment on server.
    func (pa *Payment)  CheckAndUpdatePaymentStatus(body  PaymentStatusUpdateRequest) (*PaymentStatusUpdateResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            checkAndUpdatePaymentStatusResponse *PaymentStatusUpdateResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &PaymentStatusUpdateResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &PaymentStatusUpdateResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "post",
            "/service/application/payment/v1.0/payment/confirm/polling",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &PaymentStatusUpdateResponse{}, err
            
	    }
        err = json.Unmarshal(response, &checkAndUpdatePaymentStatusResponse)
        if err != nil {
           
             return &PaymentStatusUpdateResponse{}, common.NewFDKError(err.Error())
            
        }
        return checkAndUpdatePaymentStatusResponse, nil
    }
          
    
  
    
    // GetPaymentModeRoutes Get All Valid Payment Options
    func (pa *Payment)  GetPaymentModeRoutes(Amount int, CartID string, Pincode string, CheckoutMode string, Refresh bool, AssignCardID string, UserDetails string) (*PaymentModeRouteResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getPaymentModeRoutesResponse *PaymentModeRouteResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Amount int  `url:"amount"` 
          CartID string  `url:"cart_id"` 
          Pincode string  `url:"pincode"` 
          CheckoutMode string  `url:"checkout_mode"` 
          Refresh bool  `url:"refresh"` 
          AssignCardID string  `url:"assign_card_id"` 
          UserDetails string  `url:"user_details"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Amount : Amount,
           CartID : CartID,
           Pincode : Pincode,
           CheckoutMode : CheckoutMode,
           Refresh : Refresh,
           AssignCardID : AssignCardID,
           UserDetails : UserDetails,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "get",
            fmt.Sprintf("/service/application/payment/v1.0/payment/options",),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &PaymentModeRouteResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getPaymentModeRoutesResponse)
        if err != nil {
           
             return &PaymentModeRouteResponse{}, common.NewFDKError(err.Error())
            
        }
        return getPaymentModeRoutesResponse, nil
    }
          
    
  
    
    // GetPosPaymentModeRoutes Get All Valid Payment Options for POS
    func (pa *Payment)  GetPosPaymentModeRoutes(Amount int, CartID string, Pincode string, CheckoutMode string, Refresh bool, AssignCardID string, OrderType string, UserDetails string) (*PaymentModeRouteResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getPosPaymentModeRoutesResponse *PaymentModeRouteResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          Amount int  `url:"amount"` 
          CartID string  `url:"cart_id"` 
          Pincode string  `url:"pincode"` 
          CheckoutMode string  `url:"checkout_mode"` 
          Refresh bool  `url:"refresh"` 
          AssignCardID string  `url:"assign_card_id"` 
          OrderType string  `url:"order_type"` 
          UserDetails string  `url:"user_details"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           Amount : Amount,
           CartID : CartID,
           Pincode : Pincode,
           CheckoutMode : CheckoutMode,
           Refresh : Refresh,
           AssignCardID : AssignCardID,
           OrderType : OrderType,
           UserDetails : UserDetails,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "get",
            fmt.Sprintf("/service/application/payment/v1.0/payment/options/pos",),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &PaymentModeRouteResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getPosPaymentModeRoutesResponse)
        if err != nil {
           
             return &PaymentModeRouteResponse{}, common.NewFDKError(err.Error())
            
        }
        return getPosPaymentModeRoutesResponse, nil
    }
          
    
  
    
    // GetActiveRefundTransferModes List Refund Transfer Mode
    func (pa *Payment)  GetActiveRefundTransferModes() (*TransferModeResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getActiveRefundTransferModesResponse *TransferModeResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "get",
            "/service/application/payment/v1.0/refund/transfer-mode",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &TransferModeResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getActiveRefundTransferModesResponse)
        if err != nil {
           
             return &TransferModeResponse{}, common.NewFDKError(err.Error())
            
        }
        return getActiveRefundTransferModesResponse, nil
    }
          
    
  
    
    // EnableOrDisableRefundTransferMode Enable/Disable Refund Transfer Mode
    func (pa *Payment)  EnableOrDisableRefundTransferMode(body  UpdateRefundTransferModeRequest) (*UpdateRefundTransferModeResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            enableOrDisableRefundTransferModeResponse *UpdateRefundTransferModeResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &UpdateRefundTransferModeResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &UpdateRefundTransferModeResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "put",
            "/service/application/payment/v1.0/refund/transfer-mode",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &UpdateRefundTransferModeResponse{}, err
            
	    }
        err = json.Unmarshal(response, &enableOrDisableRefundTransferModeResponse)
        if err != nil {
           
             return &UpdateRefundTransferModeResponse{}, common.NewFDKError(err.Error())
            
        }
        return enableOrDisableRefundTransferModeResponse, nil
    }
          
    
  
    
    // GetUserBeneficiariesDetail List User Beneficiary
    func (pa *Payment)  GetUserBeneficiariesDetail(OrderID string) (*OrderBeneficiaryResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getUserBeneficiariesDetailResponse *OrderBeneficiaryResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          OrderID string  `url:"order_id"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           OrderID : OrderID,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "get",
            fmt.Sprintf("/service/application/payment/v1.0/refund/user/beneficiary",),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &OrderBeneficiaryResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getUserBeneficiariesDetailResponse)
        if err != nil {
           
             return &OrderBeneficiaryResponse{}, common.NewFDKError(err.Error())
            
        }
        return getUserBeneficiariesDetailResponse, nil
    }
          
    
  
    
    // VerifyIfscCode Ifsc Code Verification
    func (pa *Payment)  VerifyIfscCode(IfscCode string) (*IfscCodeResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            verifyIfscCodeResponse *IfscCodeResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          IfscCode string  `url:"ifsc_code"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           IfscCode : IfscCode,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "get",
            "/service/application/payment/v1.0/ifsc-code/verify",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &IfscCodeResponse{}, err
            
	    }
        err = json.Unmarshal(response, &verifyIfscCodeResponse)
        if err != nil {
           
             return &IfscCodeResponse{}, common.NewFDKError(err.Error())
            
        }
        return verifyIfscCodeResponse, nil
    }
          
    
  
    
    // GetOrderBeneficiariesDetail List Order Beneficiary
    func (pa *Payment)  GetOrderBeneficiariesDetail(OrderID string) (*OrderBeneficiaryResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getOrderBeneficiariesDetailResponse *OrderBeneficiaryResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          OrderID string  `url:"order_id"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           OrderID : OrderID,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "get",
            fmt.Sprintf("/service/application/payment/v1.0/refund/order/beneficiaries",),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &OrderBeneficiaryResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getOrderBeneficiariesDetailResponse)
        if err != nil {
           
             return &OrderBeneficiaryResponse{}, common.NewFDKError(err.Error())
            
        }
        return getOrderBeneficiariesDetailResponse, nil
    }
          
    
  
    
    // VerifyOtpAndAddBeneficiaryForBank Save Beneficiary details on otp validation.
    func (pa *Payment)  VerifyOtpAndAddBeneficiaryForBank(body  AddBeneficiaryViaOtpVerificationRequest) (*AddBeneficiaryViaOtpVerificationResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            verifyOtpAndAddBeneficiaryForBankResponse *AddBeneficiaryViaOtpVerificationResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &AddBeneficiaryViaOtpVerificationResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &AddBeneficiaryViaOtpVerificationResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "post",
            "/service/application/payment/v1.0/refund/verification/bank",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &AddBeneficiaryViaOtpVerificationResponse{}, err
            
	    }
        err = json.Unmarshal(response, &verifyOtpAndAddBeneficiaryForBankResponse)
        if err != nil {
           
             return &AddBeneficiaryViaOtpVerificationResponse{}, common.NewFDKError(err.Error())
            
        }
        return verifyOtpAndAddBeneficiaryForBankResponse, nil
    }
          
    
  
    
    // AddBeneficiaryDetails Save bank details for cancelled/returned order
    func (pa *Payment)  AddBeneficiaryDetails(body  AddBeneficiaryDetailsRequest) (*RefundAccountResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            addBeneficiaryDetailsResponse *RefundAccountResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &RefundAccountResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &RefundAccountResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "post",
            "/service/application/payment/v1.0/refund/account",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &RefundAccountResponse{}, err
            
	    }
        err = json.Unmarshal(response, &addBeneficiaryDetailsResponse)
        if err != nil {
           
             return &RefundAccountResponse{}, common.NewFDKError(err.Error())
            
        }
        return addBeneficiaryDetailsResponse, nil
    }
          
    
  
    
    // VerifyOtpAndAddBeneficiaryForWallet Send Otp on Adding wallet beneficiary
    func (pa *Payment)  VerifyOtpAndAddBeneficiaryForWallet(body  WalletOtpRequest) (*WalletOtpResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            verifyOtpAndAddBeneficiaryForWalletResponse *WalletOtpResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &WalletOtpResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &WalletOtpResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "post",
            "/service/application/payment/v1.0/refund/verification/wallet",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &WalletOtpResponse{}, err
            
	    }
        err = json.Unmarshal(response, &verifyOtpAndAddBeneficiaryForWalletResponse)
        if err != nil {
           
             return &WalletOtpResponse{}, common.NewFDKError(err.Error())
            
        }
        return verifyOtpAndAddBeneficiaryForWalletResponse, nil
    }
          
    
  
    
    // UpdateDefaultBeneficiary Mark Default Beneficiary For Refund
    func (pa *Payment)  UpdateDefaultBeneficiary(body  SetDefaultBeneficiaryRequest) (*SetDefaultBeneficiaryResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateDefaultBeneficiaryResponse *SetDefaultBeneficiaryResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &SetDefaultBeneficiaryResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &SetDefaultBeneficiaryResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "post",
            "/service/application/payment/v1.0/refund/beneficiary/default",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &SetDefaultBeneficiaryResponse{}, err
            
	    }
        err = json.Unmarshal(response, &updateDefaultBeneficiaryResponse)
        if err != nil {
           
             return &SetDefaultBeneficiaryResponse{}, common.NewFDKError(err.Error())
            
        }
        return updateDefaultBeneficiaryResponse, nil
    }
          
    

    // Order ...
    type Order struct {
        config *AppConfig
    }
    // NewOrder ...
    func NewOrder(config *AppConfig) *Order {
        return &Order{config}
    }
    
  
    
    // GetOrders Get Orders for application based on application Id
    func (or *Order)  GetOrders(PageNo int, PageSize int, FromDate string, ToDate string, OrderStatus int) (*OrderList, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getOrdersResponse *OrderList
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          PageNo int  `url:"page_no"` 
          PageSize int  `url:"page_size"` 
          FromDate string  `url:"from_date"` 
          ToDate string  `url:"to_date"` 
          OrderStatus int  `url:"order_status"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           PageNo : PageNo,
           PageSize : PageSize,
           FromDate : FromDate,
           ToDate : ToDate,
           OrderStatus : OrderStatus,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            or.config,
            "get",
            "/service/application/order/v1.0/orders",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &OrderList{}, err
            
	    }
        err = json.Unmarshal(response, &getOrdersResponse)
        if err != nil {
           
             return &OrderList{}, common.NewFDKError(err.Error())
            
        }
        return getOrdersResponse, nil
    }
          
    
  
    
    // GetOrderById Get Order by order id for application based on application Id
    func (or *Order)  GetOrderById(OrderID string) (*OrderById, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getOrderByIdResponse *OrderById
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            or.config,
            "get",
            fmt.Sprintf("/service/application/order/v1.0/orders/%s",OrderID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &OrderById{}, err
            
	    }
        err = json.Unmarshal(response, &getOrderByIdResponse)
        if err != nil {
           
             return &OrderById{}, common.NewFDKError(err.Error())
            
        }
        return getOrderByIdResponse, nil
    }
          
    
  
    
    // GetShipmentById Get Shipment by shipment id and order id for application based on application Id
    func (or *Order)  GetShipmentById(ShipmentID string) (*ShipmentById, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getShipmentByIdResponse *ShipmentById
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            or.config,
            "get",
            fmt.Sprintf("/service/application/order/v1.0/orders/shipments/%s",ShipmentID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ShipmentById{}, err
            
	    }
        err = json.Unmarshal(response, &getShipmentByIdResponse)
        if err != nil {
           
             return &ShipmentById{}, common.NewFDKError(err.Error())
            
        }
        return getShipmentByIdResponse, nil
    }
          
    
  
    
    // GetShipmentReasons Get Shipment reasons by shipment id and order id for application based on application Id
    func (or *Order)  GetShipmentReasons(ShipmentID string) (*ShipmentReasons, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getShipmentReasonsResponse *ShipmentReasons
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            or.config,
            "get",
            fmt.Sprintf("/service/application/order/v1.0/orders/shipments/%s/reasons",ShipmentID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ShipmentReasons{}, err
            
	    }
        err = json.Unmarshal(response, &getShipmentReasonsResponse)
        if err != nil {
           
             return &ShipmentReasons{}, common.NewFDKError(err.Error())
            
        }
        return getShipmentReasonsResponse, nil
    }
          
    
  
    
    // UpdateShipmentStatus Update Shipment status by shipment id and order id for application based on application Id
    func (or *Order)  UpdateShipmentStatus(ShipmentID string, body  ShipmentStatusUpdateBody) (*ShipmentStatusUpdate, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateShipmentStatusResponse *ShipmentStatusUpdate
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &ShipmentStatusUpdate{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &ShipmentStatusUpdate{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            or.config,
            "put",
            fmt.Sprintf("/service/application/order/v1.0/orders/shipments/%s/status",ShipmentID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ShipmentStatusUpdate{}, err
            
	    }
        err = json.Unmarshal(response, &updateShipmentStatusResponse)
        if err != nil {
           
             return &ShipmentStatusUpdate{}, common.NewFDKError(err.Error())
            
        }
        return updateShipmentStatusResponse, nil
    }
          
    
  
    
    // TrackShipment Track Shipment by shipment id and order id for application based on application Id
    func (or *Order)  TrackShipment(ShipmentID string) (*ShipmentTrack, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            trackShipmentResponse *ShipmentTrack
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            or.config,
            "get",
            fmt.Sprintf("/service/application/order/v1.0/orders/shipments/%s/track",ShipmentID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ShipmentTrack{}, err
            
	    }
        err = json.Unmarshal(response, &trackShipmentResponse)
        if err != nil {
           
             return &ShipmentTrack{}, common.NewFDKError(err.Error())
            
        }
        return trackShipmentResponse, nil
    }
          
    
  
    
    // GetPosOrderById Get POS Order by order id for application based on application Id
    func (or *Order)  GetPosOrderById(OrderID string) (*PosOrderById, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getPosOrderByIdResponse *PosOrderById
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            or.config,
            "get",
            fmt.Sprintf("/service/application/order/v1.0/orders/pos-order/%s",OrderID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &PosOrderById{}, err
            
	    }
        err = json.Unmarshal(response, &getPosOrderByIdResponse)
        if err != nil {
           
             return &PosOrderById{}, common.NewFDKError(err.Error())
            
        }
        return getPosOrderByIdResponse, nil
    }
          
    

    // Rewards ...
    type Rewards struct {
        config *AppConfig
    }
    // NewRewards ...
    func NewRewards(config *AppConfig) *Rewards {
        return &Rewards{config}
    }
    
  
    
    // GetPointsOnProduct Get reward points that could be earned on any catalogue product.
    func (re *Rewards)  GetPointsOnProduct(body  CatalogueOrderRequest) (*CatalogueOrderResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getPointsOnProductResponse *CatalogueOrderResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &CatalogueOrderResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &CatalogueOrderResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            re.config,
            "post",
            "/service/application/rewards/v1.0/catalogue/offer/order/",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CatalogueOrderResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getPointsOnProductResponse)
        if err != nil {
           
             return &CatalogueOrderResponse{}, common.NewFDKError(err.Error())
            
        }
        return getPointsOnProductResponse, nil
    }
          
    
  
    
    // GetOrderDiscount Calculates the discount on order-amount based on amount ranges configured in order_discount reward.
    func (re *Rewards)  GetOrderDiscount(body  OrderDiscountRequest) (*OrderDiscountResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getOrderDiscountResponse *OrderDiscountResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &OrderDiscountResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &OrderDiscountResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            re.config,
            "post",
            "/service/application/rewards/v1.0/user/offers/order-discount/",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &OrderDiscountResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getOrderDiscountResponse)
        if err != nil {
           
             return &OrderDiscountResponse{}, common.NewFDKError(err.Error())
            
        }
        return getOrderDiscountResponse, nil
    }
          
    
  
    
    // GetUserPoints Total available points of a user for current application
    func (re *Rewards)  GetUserPoints() (*PointsResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getUserPointsResponse *PointsResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            re.config,
            "get",
            "/service/application/rewards/v1.0/user/points",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &PointsResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getUserPointsResponse)
        if err != nil {
           
             return &PointsResponse{}, common.NewFDKError(err.Error())
            
        }
        return getUserPointsResponse, nil
    }
          
    
  
    
    // GetUserPointsHistory Get list of points transactions.
    func (re *Rewards)  GetUserPointsHistory(PageID string, PageSize int) (*PointsHistoryResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getUserPointsHistoryResponse *PointsHistoryResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          PageID string  `url:"page_id"` 
          PageSize int  `url:"page_size"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           PageID : PageID,
           PageSize : PageSize,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            re.config,
            "get",
            "/service/application/rewards/v1.0/user/points/history/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &PointsHistoryResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getUserPointsHistoryResponse)
        if err != nil {
           
             return &PointsHistoryResponse{}, common.NewFDKError(err.Error())
            
        }
        return getUserPointsHistoryResponse, nil
    }
          
            
            
                
                    
                    
                        
                    
                
                    
                        
                    
                    
                
            
            // GetUserPointsHistoryPaginator Get list of points transactions.  
            func (re *Rewards)  GetUserPointsHistoryPaginator(PageSize int ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                paginator.Next = func() (interface{}, error) {
                    response, err := re.GetUserPointsHistory(paginator.NextID, PageSize)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    
  
    
    // GetUserReferralDetails User's referral details.
    func (re *Rewards)  GetUserReferralDetails() (*ReferralDetailsResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getUserReferralDetailsResponse *ReferralDetailsResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            re.config,
            "get",
            "/service/application/rewards/v1.0/user/referral/",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ReferralDetailsResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getUserReferralDetailsResponse)
        if err != nil {
           
             return &ReferralDetailsResponse{}, common.NewFDKError(err.Error())
            
        }
        return getUserReferralDetailsResponse, nil
    }
          
    
  
    
    // RedeemReferralCode Redeems referral code and credits points to users points account.
    func (re *Rewards)  RedeemReferralCode(body  RedeemReferralCodeRequest) (*RedeemReferralCodeResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            redeemReferralCodeResponse *RedeemReferralCodeResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &RedeemReferralCodeResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &RedeemReferralCodeResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            re.config,
            "post",
            "/service/application/rewards/v1.0/user/referral/redeem/",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &RedeemReferralCodeResponse{}, err
            
	    }
        err = json.Unmarshal(response, &redeemReferralCodeResponse)
        if err != nil {
           
             return &RedeemReferralCodeResponse{}, common.NewFDKError(err.Error())
            
        }
        return redeemReferralCodeResponse, nil
    }
          
    

    // Feedback ...
    type Feedback struct {
        config *AppConfig
    }
    // NewFeedback ...
    func NewFeedback(config *AppConfig) *Feedback {
        return &Feedback{config}
    }
    
  
    
    // CreateAbuseReport post a new abuse request
    func (fe *Feedback)  CreateAbuseReport(body  ReportAbuseRequest) (*InsertResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            createAbuseReportResponse *InsertResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &InsertResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &InsertResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "post",
            "/service/application/feedback/v1.0/abuse",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &InsertResponse{}, err
            
	    }
        err = json.Unmarshal(response, &createAbuseReportResponse)
        if err != nil {
           
             return &InsertResponse{}, common.NewFDKError(err.Error())
            
        }
        return createAbuseReportResponse, nil
    }
          
    
  
    
    // UpdateAbuseReport Update abuse details
    func (fe *Feedback)  UpdateAbuseReport(body  UpdateAbuseStatusRequest) (*UpdateResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateAbuseReportResponse *UpdateResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &UpdateResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &UpdateResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "put",
            "/service/application/feedback/v1.0/abuse",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &UpdateResponse{}, err
            
	    }
        err = json.Unmarshal(response, &updateAbuseReportResponse)
        if err != nil {
           
             return &UpdateResponse{}, common.NewFDKError(err.Error())
            
        }
        return updateAbuseReportResponse, nil
    }
          
    
  
    
    // GetAbuseReports Get list of abuse data
    func (fe *Feedback)  GetAbuseReports(EntityID string, EntityType string, ID string, PageID string, PageSize int) (*ReportAbuseGetResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getAbuseReportsResponse *ReportAbuseGetResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          ID string  `url:"id"` 
          PageID string  `url:"page_id"` 
          PageSize int  `url:"page_size"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           ID : ID,
           PageID : PageID,
           PageSize : PageSize,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "get",
            fmt.Sprintf("/service/application/feedback/v1.0/abuse/entity/%s/entity-id/%s",EntityID,EntityType),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ReportAbuseGetResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getAbuseReportsResponse)
        if err != nil {
           
             return &ReportAbuseGetResponse{}, common.NewFDKError(err.Error())
            
        }
        return getAbuseReportsResponse, nil
    }
          
            
            
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                    
                        
                    
                
                    
                        
                    
                    
                
            
            // GetAbuseReportsPaginator Get list of abuse data  
            func (fe *Feedback)  GetAbuseReportsPaginator(EntityID string , EntityType string , ID string , PageSize int ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                paginator.Next = func() (interface{}, error) {
                    response, err := fe.GetAbuseReports(EntityID, EntityType, ID, paginator.NextID, PageSize)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    
  
    
    // GetAttributes Get list of attribute data
    func (fe *Feedback)  GetAttributes(PageNo int, PageSize int) (*AttributeResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getAttributesResponse *AttributeResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          PageNo int  `url:"page_no"` 
          PageSize int  `url:"page_size"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           PageNo : PageNo,
           PageSize : PageSize,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "get",
            "/service/application/feedback/v1.0/attributes",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &AttributeResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getAttributesResponse)
        if err != nil {
           
             return &AttributeResponse{}, common.NewFDKError(err.Error())
            
        }
        return getAttributesResponse, nil
    }
          
            
            
                
                    
                    
                
                    
                        
                    
                    
                
            
            // GetAttributesPaginator Get list of attribute data  
            func (fe *Feedback)  GetAttributesPaginator(PageSize int ) *common.Paginator {
                paginator := common.NewPaginator("number")
                paginator.Next = func() (interface{}, error) {
                    response, err := fe.GetAttributes(paginator.PageNo, PageSize)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    
  
    
    // CreateAttribute Add a new attribute request
    func (fe *Feedback)  CreateAttribute(body  SaveAttributeRequest) (*InsertResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            createAttributeResponse *InsertResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &InsertResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &InsertResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "post",
            "/service/application/feedback/v1.0/attributes",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &InsertResponse{}, err
            
	    }
        err = json.Unmarshal(response, &createAttributeResponse)
        if err != nil {
           
             return &InsertResponse{}, common.NewFDKError(err.Error())
            
        }
        return createAttributeResponse, nil
    }
          
    
  
    
    // GetAttribute Get single attribute data
    func (fe *Feedback)  GetAttribute(Slug string) (*Attribute, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getAttributeResponse *Attribute
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "get",
            fmt.Sprintf("/service/application/feedback/v1.0/attributes/%s",Slug),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &Attribute{}, err
            
	    }
        err = json.Unmarshal(response, &getAttributeResponse)
        if err != nil {
           
             return &Attribute{}, common.NewFDKError(err.Error())
            
        }
        return getAttributeResponse, nil
    }
          
    
  
    
    // UpdateAttribute Update attribute details
    func (fe *Feedback)  UpdateAttribute(Slug string, body  UpdateAttributeRequest) (*UpdateResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateAttributeResponse *UpdateResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &UpdateResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &UpdateResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "put",
            fmt.Sprintf("/service/application/feedback/v1.0/attributes/%s",Slug),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &UpdateResponse{}, err
            
	    }
        err = json.Unmarshal(response, &updateAttributeResponse)
        if err != nil {
           
             return &UpdateResponse{}, common.NewFDKError(err.Error())
            
        }
        return updateAttributeResponse, nil
    }
          
    
  
    
    // CreateComment post a new comment
    func (fe *Feedback)  CreateComment(body  CommentRequest) (*InsertResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            createCommentResponse *InsertResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &InsertResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &InsertResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "post",
            "/service/application/feedback/v1.0/comment",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &InsertResponse{}, err
            
	    }
        err = json.Unmarshal(response, &createCommentResponse)
        if err != nil {
           
             return &InsertResponse{}, common.NewFDKError(err.Error())
            
        }
        return createCommentResponse, nil
    }
          
    
  
    
    // UpdateComment Update comment status
    func (fe *Feedback)  UpdateComment(body  UpdateCommentRequest) (*UpdateResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateCommentResponse *UpdateResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &UpdateResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &UpdateResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "put",
            "/service/application/feedback/v1.0/comment",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &UpdateResponse{}, err
            
	    }
        err = json.Unmarshal(response, &updateCommentResponse)
        if err != nil {
           
             return &UpdateResponse{}, common.NewFDKError(err.Error())
            
        }
        return updateCommentResponse, nil
    }
          
    
  
    
    // GetComments Get list of comments
    func (fe *Feedback)  GetComments(EntityType string, ID string, EntityID string, UserID string, PageID string, PageSize int) (*CommentGetResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getCommentsResponse *CommentGetResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          ID string  `url:"id"` 
          EntityID string  `url:"entity_id"` 
          UserID string  `url:"user_id"` 
          PageID string  `url:"page_id"` 
          PageSize int  `url:"page_size"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           ID : ID,
           EntityID : EntityID,
           UserID : UserID,
           PageID : PageID,
           PageSize : PageSize,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "get",
            fmt.Sprintf("/service/application/feedback/v1.0/comment/entity/%s",EntityType),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CommentGetResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getCommentsResponse)
        if err != nil {
           
             return &CommentGetResponse{}, common.NewFDKError(err.Error())
            
        }
        return getCommentsResponse, nil
    }
          
            
            
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                    
                        
                    
                
                    
                        
                    
                    
                
            
            // GetCommentsPaginator Get list of comments  
            func (fe *Feedback)  GetCommentsPaginator(EntityType string , ID string , EntityID string , UserID string , PageSize int ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                paginator.Next = func() (interface{}, error) {
                    response, err := fe.GetComments(EntityType, ID, EntityID, UserID, paginator.NextID, PageSize)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    
  
    
    // CheckEligibility Checks eligibility and cloud media config
    func (fe *Feedback)  CheckEligibility(EntityType string, EntityID string) (*CheckEligibilityResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            checkEligibilityResponse *CheckEligibilityResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "get",
            fmt.Sprintf("/service/application/feedback/v1.0/config/entity/%s/entity-id/%s",EntityType,EntityID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CheckEligibilityResponse{}, err
            
	    }
        err = json.Unmarshal(response, &checkEligibilityResponse)
        if err != nil {
           
             return &CheckEligibilityResponse{}, common.NewFDKError(err.Error())
            
        }
        return checkEligibilityResponse, nil
    }
          
    
  
    
    // DeleteMedia Delete Media
    func (fe *Feedback)  DeleteMedia() (*UpdateResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            deleteMediaResponse *UpdateResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "delete",
            "/service/application/feedback/v1.0/media/",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &UpdateResponse{}, err
            
	    }
        err = json.Unmarshal(response, &deleteMediaResponse)
        if err != nil {
           
             return &UpdateResponse{}, common.NewFDKError(err.Error())
            
        }
        return deleteMediaResponse, nil
    }
          
    
  
    
    // CreateMedia Add Media
    func (fe *Feedback)  CreateMedia(body  AddMediaListRequest) (*InsertResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            createMediaResponse *InsertResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &InsertResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &InsertResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "post",
            "/service/application/feedback/v1.0/media/",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &InsertResponse{}, err
            
	    }
        err = json.Unmarshal(response, &createMediaResponse)
        if err != nil {
           
             return &InsertResponse{}, common.NewFDKError(err.Error())
            
        }
        return createMediaResponse, nil
    }
          
    
  
    
    // UpdateMedia Update Media
    func (fe *Feedback)  UpdateMedia(body  UpdateMediaListRequest) (*UpdateResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateMediaResponse *UpdateResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &UpdateResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &UpdateResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "put",
            "/service/application/feedback/v1.0/media/",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &UpdateResponse{}, err
            
	    }
        err = json.Unmarshal(response, &updateMediaResponse)
        if err != nil {
           
             return &UpdateResponse{}, common.NewFDKError(err.Error())
            
        }
        return updateMediaResponse, nil
    }
          
    
  
    
    // GetMedias Get Media
    func (fe *Feedback)  GetMedias(EntityType string, EntityID string, ID string, PageID string, PageSize int) (*MediaGetResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getMediasResponse *MediaGetResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          ID string  `url:"id"` 
          PageID string  `url:"page_id"` 
          PageSize int  `url:"page_size"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           ID : ID,
           PageID : PageID,
           PageSize : PageSize,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "get",
            fmt.Sprintf("/service/application/feedback/v1.0/media/entity/%s/entity-id/%s",EntityType,EntityID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &MediaGetResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getMediasResponse)
        if err != nil {
           
             return &MediaGetResponse{}, common.NewFDKError(err.Error())
            
        }
        return getMediasResponse, nil
    }
          
            
            
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                    
                        
                    
                
                    
                        
                    
                    
                
            
            // GetMediasPaginator Get Media  
            func (fe *Feedback)  GetMediasPaginator(EntityType string , EntityID string , ID string , PageSize int ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                paginator.Next = func() (interface{}, error) {
                    response, err := fe.GetMedias(EntityType, EntityID, ID, paginator.NextID, PageSize)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    
  
    
    // GetReviewSummaries Get a review summary
    func (fe *Feedback)  GetReviewSummaries(EntityType string, EntityID string, ID string, PageID string, PageSize int) (*RatingGetResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getReviewSummariesResponse *RatingGetResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          ID string  `url:"id"` 
          PageID string  `url:"page_id"` 
          PageSize int  `url:"page_size"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           ID : ID,
           PageID : PageID,
           PageSize : PageSize,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "get",
            fmt.Sprintf("/service/application/feedback/v1.0/rating/summary/entity/%s/entity-id/%s",EntityType,EntityID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &RatingGetResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getReviewSummariesResponse)
        if err != nil {
           
             return &RatingGetResponse{}, common.NewFDKError(err.Error())
            
        }
        return getReviewSummariesResponse, nil
    }
          
            
            
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                    
                        
                    
                
                    
                        
                    
                    
                
            
            // GetReviewSummariesPaginator Get a review summary  
            func (fe *Feedback)  GetReviewSummariesPaginator(EntityType string , EntityID string , ID string , PageSize int ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                paginator.Next = func() (interface{}, error) {
                    response, err := fe.GetReviewSummaries(EntityType, EntityID, ID, paginator.NextID, PageSize)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    
  
    
    // CreateReview Add customer reviews
    func (fe *Feedback)  CreateReview(body  UpdateReviewRequest) (*UpdateResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            createReviewResponse *UpdateResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &UpdateResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &UpdateResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "post",
            "/service/application/feedback/v1.0/review/",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &UpdateResponse{}, err
            
	    }
        err = json.Unmarshal(response, &createReviewResponse)
        if err != nil {
           
             return &UpdateResponse{}, common.NewFDKError(err.Error())
            
        }
        return createReviewResponse, nil
    }
          
    
  
    
    // UpdateReview Update customer reviews
    func (fe *Feedback)  UpdateReview(body  UpdateReviewRequest) (*UpdateResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateReviewResponse *UpdateResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &UpdateResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &UpdateResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "put",
            "/service/application/feedback/v1.0/review/",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &UpdateResponse{}, err
            
	    }
        err = json.Unmarshal(response, &updateReviewResponse)
        if err != nil {
           
             return &UpdateResponse{}, common.NewFDKError(err.Error())
            
        }
        return updateReviewResponse, nil
    }
          
    
  
    
    // GetReviews Get list of customer reviews
    func (fe *Feedback)  GetReviews(EntityType string, EntityID string, ID string, UserID string, Media string, Rating []int, AttributeRating []string, Facets bool, Sort string, PageID string, PageSize int) (*ReviewGetResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getReviewsResponse *ReviewGetResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          ID string  `url:"id"` 
          UserID string  `url:"user_id"` 
          Media string  `url:"media"` 
          Rating []int  `url:"rating"` 
          AttributeRating []string  `url:"attribute_rating"` 
          Facets bool  `url:"facets"` 
          Sort string  `url:"sort"` 
          PageID string  `url:"page_id"` 
          PageSize int  `url:"page_size"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           ID : ID,
           UserID : UserID,
           Media : Media,
           Rating : Rating,
           AttributeRating : AttributeRating,
           Facets : Facets,
           Sort : Sort,
           PageID : PageID,
           PageSize : PageSize,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "get",
            fmt.Sprintf("/service/application/feedback/v1.0/review/entity/%s/entity-id/%s",EntityType,EntityID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &ReviewGetResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getReviewsResponse)
        if err != nil {
           
             return &ReviewGetResponse{}, common.NewFDKError(err.Error())
            
        }
        return getReviewsResponse, nil
    }
          
            
            
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                    
                        
                    
                
                    
                        
                    
                    
                
            
            // GetReviewsPaginator Get list of customer reviews  
            func (fe *Feedback)  GetReviewsPaginator(EntityType string , EntityID string , ID string , UserID string , Media string , Rating []int , AttributeRating []string , Facets bool , Sort string , PageSize int ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                paginator.Next = func() (interface{}, error) {
                    response, err := fe.GetReviews(EntityType, EntityID, ID, UserID, Media, Rating, AttributeRating, Facets, Sort, paginator.NextID, PageSize)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    
  
    
    // GetTemplates Get the templates for product or l3 type
    func (fe *Feedback)  GetTemplates(TemplateID string, EntityID string, EntityType string) (*CursorGetResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getTemplatesResponse *CursorGetResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          TemplateID string  `url:"template_id"` 
          EntityID string  `url:"entity_id"` 
          EntityType string  `url:"entity_type"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           TemplateID : TemplateID,
           EntityID : EntityID,
           EntityType : EntityType,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "get",
            "/service/application/feedback/v1.0/template/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CursorGetResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getTemplatesResponse)
        if err != nil {
           
             return &CursorGetResponse{}, common.NewFDKError(err.Error())
            
        }
        return getTemplatesResponse, nil
    }
          
    
  
    
    // CreateQuestion Create a new question
    func (fe *Feedback)  CreateQuestion(body  CreateQNARequest) (*InsertResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            createQuestionResponse *InsertResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &InsertResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &InsertResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "post",
            "/service/application/feedback/v1.0/template/qna/",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &InsertResponse{}, err
            
	    }
        err = json.Unmarshal(response, &createQuestionResponse)
        if err != nil {
           
             return &InsertResponse{}, common.NewFDKError(err.Error())
            
        }
        return createQuestionResponse, nil
    }
          
    
  
    
    // UpdateQuestion Update question
    func (fe *Feedback)  UpdateQuestion(body  UpdateQNARequest) (*UpdateResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateQuestionResponse *UpdateResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &UpdateResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &UpdateResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "put",
            "/service/application/feedback/v1.0/template/qna/",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &UpdateResponse{}, err
            
	    }
        err = json.Unmarshal(response, &updateQuestionResponse)
        if err != nil {
           
             return &UpdateResponse{}, common.NewFDKError(err.Error())
            
        }
        return updateQuestionResponse, nil
    }
          
    
  
    
    // GetQuestionAndAnswers Get a list of QnA
    func (fe *Feedback)  GetQuestionAndAnswers(EntityType string, EntityID string, ID string, ShowAnswer bool, PageID string, PageSize int) (*QNAGetResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getQuestionAndAnswersResponse *QNAGetResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          ID string  `url:"id"` 
          ShowAnswer bool  `url:"show_answer"` 
          PageID string  `url:"page_id"` 
          PageSize int  `url:"page_size"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           ID : ID,
           ShowAnswer : ShowAnswer,
           PageID : PageID,
           PageSize : PageSize,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "get",
            fmt.Sprintf("/service/application/feedback/v1.0/template/qna/entity/%s/entity-id/%s",EntityType,EntityID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &QNAGetResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getQuestionAndAnswersResponse)
        if err != nil {
           
             return &QNAGetResponse{}, common.NewFDKError(err.Error())
            
        }
        return getQuestionAndAnswersResponse, nil
    }
          
            
            
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                    
                        
                    
                
                    
                        
                    
                    
                
            
            // GetQuestionAndAnswersPaginator Get a list of QnA  
            func (fe *Feedback)  GetQuestionAndAnswersPaginator(EntityType string , EntityID string , ID string , ShowAnswer bool , PageSize int ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                paginator.Next = func() (interface{}, error) {
                    response, err := fe.GetQuestionAndAnswers(EntityType, EntityID, ID, ShowAnswer, paginator.NextID, PageSize)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    
  
    
    // GetVotes Get list of votes
    func (fe *Feedback)  GetVotes(ID string, RefType string, PageNo int, PageSize int) (*VoteResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getVotesResponse *VoteResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          ID string  `url:"id"` 
          RefType string  `url:"ref_type"` 
          PageNo int  `url:"page_no"` 
          PageSize int  `url:"page_size"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           ID : ID,
           RefType : RefType,
           PageNo : PageNo,
           PageSize : PageSize,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "get",
            "/service/application/feedback/v1.0/vote/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &VoteResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getVotesResponse)
        if err != nil {
           
             return &VoteResponse{}, common.NewFDKError(err.Error())
            
        }
        return getVotesResponse, nil
    }
          
            
            
                
                    
                        
                    
                    
                
                    
                        
                    
                    
                
                    
                    
                
                    
                        
                    
                    
                
            
            // GetVotesPaginator Get list of votes  
            func (fe *Feedback)  GetVotesPaginator(ID string , RefType string , PageSize int ) *common.Paginator {
                paginator := common.NewPaginator("number")
                paginator.Next = func() (interface{}, error) {
                    response, err := fe.GetVotes(ID, RefType, paginator.PageNo, PageSize)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, response.Page.Current+1, response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }

       
    
  
    
    // CreateVote Create a new vote
    func (fe *Feedback)  CreateVote(body  VoteRequest) (*InsertResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            createVoteResponse *InsertResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &InsertResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &InsertResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "post",
            "/service/application/feedback/v1.0/vote/",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &InsertResponse{}, err
            
	    }
        err = json.Unmarshal(response, &createVoteResponse)
        if err != nil {
           
             return &InsertResponse{}, common.NewFDKError(err.Error())
            
        }
        return createVoteResponse, nil
    }
          
    
  
    
    // UpdateVote Update vote
    func (fe *Feedback)  UpdateVote(body  UpdateVoteRequest) (*UpdateResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateVoteResponse *UpdateResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &UpdateResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &UpdateResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            fe.config,
            "put",
            "/service/application/feedback/v1.0/vote/",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &UpdateResponse{}, err
            
	    }
        err = json.Unmarshal(response, &updateVoteResponse)
        if err != nil {
           
             return &UpdateResponse{}, common.NewFDKError(err.Error())
            
        }
        return updateVoteResponse, nil
    }
          
    

    // PosCart ...
    type PosCart struct {
        config *AppConfig
    }
    // NewPosCart ...
    func NewPosCart(config *AppConfig) *PosCart {
        return &PosCart{config}
    }
    
  
    
    // GetCart Fetch all Items Added to  Cart
    func (po *PosCart)  GetCart(UID int, I bool, B bool, AssignCardID int) (*CartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getCartResponse *CartResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID int  `url:"uid"` 
          I bool  `url:"i"` 
          B bool  `url:"b"` 
          AssignCardID int  `url:"assign_card_id"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
           I : I,
           B : B,
           AssignCardID : AssignCardID,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "get",
            "/service/application/pos/cart/v1.0/detail",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CartResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getCartResponse)
        if err != nil {
           
             return &CartResponse{}, common.NewFDKError(err.Error())
            
        }
        return getCartResponse, nil
    }
          
    
  
    
    // GetCartLastModified Fetch Last-Modified timestamp
    func (po *PosCart)  GetCartLastModified(UID int) (interface{}, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
             getCartLastModifiedResponse interface{}
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID int  `url:"uid"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "head",
            "/service/application/pos/cart/v1.0/detail",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
            return getCartLastModifiedResponse, err
            
	    }
        err = json.Unmarshal(response, &getCartLastModifiedResponse)
        if err != nil {
           
            return getCartLastModifiedResponse, common.NewFDKError(err.Error())
            
        }
        return getCartLastModifiedResponse, nil
    }
          
    
  
    
    // AddItems Add Items to Cart
    func (po *PosCart)  AddItems(I bool, B bool, body  AddCartRequest) (*AddCartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            addItemsResponse *AddCartResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          I bool  `url:"i"` 
          B bool  `url:"b"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           I : I,
           B : B,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &AddCartResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &AddCartResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "post",
            "/service/application/pos/cart/v1.0/detail",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &AddCartResponse{}, err
            
	    }
        err = json.Unmarshal(response, &addItemsResponse)
        if err != nil {
           
             return &AddCartResponse{}, common.NewFDKError(err.Error())
            
        }
        return addItemsResponse, nil
    }
          
    
  
    
    // UpdateCart Update Items already added to Cart
    func (po *PosCart)  UpdateCart(UID int, I bool, B bool, body  UpdateCartRequest) (*UpdateCartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateCartResponse *UpdateCartResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID int  `url:"uid"` 
          I bool  `url:"i"` 
          B bool  `url:"b"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
           I : I,
           B : B,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &UpdateCartResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &UpdateCartResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "put",
            "/service/application/pos/cart/v1.0/detail",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &UpdateCartResponse{}, err
            
	    }
        err = json.Unmarshal(response, &updateCartResponse)
        if err != nil {
           
             return &UpdateCartResponse{}, common.NewFDKError(err.Error())
            
        }
        return updateCartResponse, nil
    }
          
    
  
    
    // GetItemCount Cart item count
    func (po *PosCart)  GetItemCount(UID int) (*CartItemCountResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getItemCountResponse *CartItemCountResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID int  `url:"uid"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "get",
            "/service/application/pos/cart/v1.0/basic",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CartItemCountResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getItemCountResponse)
        if err != nil {
           
             return &CartItemCountResponse{}, common.NewFDKError(err.Error())
            
        }
        return getItemCountResponse, nil
    }
          
    
  
    
    // GetCoupons Fetch Coupon
    func (po *PosCart)  GetCoupons(UID int) (*GetCouponResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getCouponsResponse *GetCouponResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID int  `url:"uid"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "get",
            "/service/application/pos/cart/v1.0/coupon",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &GetCouponResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getCouponsResponse)
        if err != nil {
           
             return &GetCouponResponse{}, common.NewFDKError(err.Error())
            
        }
        return getCouponsResponse, nil
    }
          
    
  
    
    // ApplyCoupon Apply Coupon
    func (po *PosCart)  ApplyCoupon(I bool, B bool, P bool, UID int, body  ApplyCouponRequest) (*CartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            applyCouponResponse *CartResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          I bool  `url:"i"` 
          B bool  `url:"b"` 
          P bool  `url:"p"` 
          UID int  `url:"uid"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           I : I,
           B : B,
           P : P,
           UID : UID,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &CartResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &CartResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "post",
            "/service/application/pos/cart/v1.0/coupon",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CartResponse{}, err
            
	    }
        err = json.Unmarshal(response, &applyCouponResponse)
        if err != nil {
           
             return &CartResponse{}, common.NewFDKError(err.Error())
            
        }
        return applyCouponResponse, nil
    }
          
    
  
    
    // RemoveCoupon Remove Coupon Applied
    func (po *PosCart)  RemoveCoupon(UID int) (*CartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            removeCouponResponse *CartResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID int  `url:"uid"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "delete",
            "/service/application/pos/cart/v1.0/coupon",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CartResponse{}, err
            
	    }
        err = json.Unmarshal(response, &removeCouponResponse)
        if err != nil {
           
             return &CartResponse{}, common.NewFDKError(err.Error())
            
        }
        return removeCouponResponse, nil
    }
          
    
  
    
    // GetBulkDiscountOffers Get discount offers based on quantity
    func (po *PosCart)  GetBulkDiscountOffers(ItemID int, ArticleID string, UID int, Slug string) (*BulkPriceResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getBulkDiscountOffersResponse *BulkPriceResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          ItemID int  `url:"item_id"` 
          ArticleID string  `url:"article_id"` 
          UID int  `url:"uid"` 
          Slug string  `url:"slug"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           ItemID : ItemID,
           ArticleID : ArticleID,
           UID : UID,
           Slug : Slug,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "get",
            "/service/application/pos/cart/v1.0/bulk-price",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &BulkPriceResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getBulkDiscountOffersResponse)
        if err != nil {
           
             return &BulkPriceResponse{}, common.NewFDKError(err.Error())
            
        }
        return getBulkDiscountOffersResponse, nil
    }
          
    
  
    
    // GetAddresses Fetch Address
    func (po *PosCart)  GetAddresses(UID int, MobileNo string, CheckoutMode string, Tags string, IsDefault bool) (*GetAddressesResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getAddressesResponse *GetAddressesResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID int  `url:"uid"` 
          MobileNo string  `url:"mobile_no"` 
          CheckoutMode string  `url:"checkout_mode"` 
          Tags string  `url:"tags"` 
          IsDefault bool  `url:"is_default"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
           MobileNo : MobileNo,
           CheckoutMode : CheckoutMode,
           Tags : Tags,
           IsDefault : IsDefault,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "get",
            "/service/application/pos/cart/v1.0/address",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &GetAddressesResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getAddressesResponse)
        if err != nil {
           
             return &GetAddressesResponse{}, common.NewFDKError(err.Error())
            
        }
        return getAddressesResponse, nil
    }
          
    
  
    
    // AddAddress Add Address to the account
    func (po *PosCart)  AddAddress(body  Address) (*SaveAddressResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            addAddressResponse *SaveAddressResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &SaveAddressResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &SaveAddressResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "post",
            "/service/application/pos/cart/v1.0/address",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &SaveAddressResponse{}, err
            
	    }
        err = json.Unmarshal(response, &addAddressResponse)
        if err != nil {
           
             return &SaveAddressResponse{}, common.NewFDKError(err.Error())
            
        }
        return addAddressResponse, nil
    }
          
    
  
    
    // GetAddressById Fetch Single Address
    func (po *PosCart)  GetAddressById(ID int, UID int, MobileNo string, CheckoutMode string, Tags string, IsDefault bool) (*Address, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getAddressByIdResponse *Address
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID int  `url:"uid"` 
          MobileNo string  `url:"mobile_no"` 
          CheckoutMode string  `url:"checkout_mode"` 
          Tags string  `url:"tags"` 
          IsDefault bool  `url:"is_default"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
           MobileNo : MobileNo,
           CheckoutMode : CheckoutMode,
           Tags : Tags,
           IsDefault : IsDefault,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "get",
            fmt.Sprintf("/service/application/pos/cart/v1.0/address/%d",ID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &Address{}, err
            
	    }
        err = json.Unmarshal(response, &getAddressByIdResponse)
        if err != nil {
           
             return &Address{}, common.NewFDKError(err.Error())
            
        }
        return getAddressByIdResponse, nil
    }
          
    
  
    
    // UpdateAddress Update Address alreay added to account
    func (po *PosCart)  UpdateAddress(ID int, body  Address) (*UpdateAddressResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateAddressResponse *UpdateAddressResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &UpdateAddressResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &UpdateAddressResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "put",
            fmt.Sprintf("/service/application/pos/cart/v1.0/address/%d",ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &UpdateAddressResponse{}, err
            
	    }
        err = json.Unmarshal(response, &updateAddressResponse)
        if err != nil {
           
             return &UpdateAddressResponse{}, common.NewFDKError(err.Error())
            
        }
        return updateAddressResponse, nil
    }
          
    
  
    
    // RemoveAddress Remove Address Associated to the account
    func (po *PosCart)  RemoveAddress(ID int) (*DeleteAddressResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            removeAddressResponse *DeleteAddressResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "delete",
            fmt.Sprintf("/service/application/pos/cart/v1.0/address/%d",ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &DeleteAddressResponse{}, err
            
	    }
        err = json.Unmarshal(response, &removeAddressResponse)
        if err != nil {
           
             return &DeleteAddressResponse{}, common.NewFDKError(err.Error())
            
        }
        return removeAddressResponse, nil
    }
          
    
  
    
    // SelectAddress Select Address from All Addresses
    func (po *PosCart)  SelectAddress(UID int, I bool, B bool, body  SelectCartAddressRequest) (*CartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            selectAddressResponse *CartResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID int  `url:"uid"` 
          I bool  `url:"i"` 
          B bool  `url:"b"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
           I : I,
           B : B,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &CartResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &CartResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "post",
            "/service/application/pos/cart/v1.0/select-address",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CartResponse{}, err
            
	    }
        err = json.Unmarshal(response, &selectAddressResponse)
        if err != nil {
           
             return &CartResponse{}, common.NewFDKError(err.Error())
            
        }
        return selectAddressResponse, nil
    }
          
    
  
    
    // SelectPaymentMode Update Cart Payment
    func (po *PosCart)  SelectPaymentMode(UID string, body  UpdateCartPaymentRequest) (*CartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            selectPaymentModeResponse *CartResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID string  `url:"uid"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &CartResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &CartResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "put",
            "/service/application/pos/cart/v1.0/payment",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CartResponse{}, err
            
	    }
        err = json.Unmarshal(response, &selectPaymentModeResponse)
        if err != nil {
           
             return &CartResponse{}, common.NewFDKError(err.Error())
            
        }
        return selectPaymentModeResponse, nil
    }
          
    
  
    
    // ValidateCouponForPayment Get Cart Payment for valid coupon
    func (po *PosCart)  ValidateCouponForPayment(UID string, AddressID string, PaymentMode string, PaymentIdentifier string, AggregatorName string, MerchantCode string) (*PaymentCouponValidate, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            validateCouponForPaymentResponse *PaymentCouponValidate
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID string  `url:"uid"` 
          AddressID string  `url:"address_id"` 
          PaymentMode string  `url:"payment_mode"` 
          PaymentIdentifier string  `url:"payment_identifier"` 
          AggregatorName string  `url:"aggregator_name"` 
          MerchantCode string  `url:"merchant_code"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
           AddressID : AddressID,
           PaymentMode : PaymentMode,
           PaymentIdentifier : PaymentIdentifier,
           AggregatorName : AggregatorName,
           MerchantCode : MerchantCode,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "get",
            "/service/application/pos/cart/v1.0/payment/validate/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &PaymentCouponValidate{}, err
            
	    }
        err = json.Unmarshal(response, &validateCouponForPaymentResponse)
        if err != nil {
           
             return &PaymentCouponValidate{}, common.NewFDKError(err.Error())
            
        }
        return validateCouponForPaymentResponse, nil
    }
          
    
  
    
    // GetShipments Get delivery date and options before checkout
    func (po *PosCart)  GetShipments(PickAtStoreUID int, OrderingStoreID int, P bool, UID int, AddressID int, AreaCode string, OrderType string) (*CartShipmentsResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getShipmentsResponse *CartShipmentsResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          PickAtStoreUID int  `url:"pick_at_store_uid"` 
          OrderingStoreID int  `url:"ordering_store_id"` 
          P bool  `url:"p"` 
          UID int  `url:"uid"` 
          AddressID int  `url:"address_id"` 
          AreaCode string  `url:"area_code"` 
          OrderType string  `url:"order_type"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           PickAtStoreUID : PickAtStoreUID,
           OrderingStoreID : OrderingStoreID,
           P : P,
           UID : UID,
           AddressID : AddressID,
           AreaCode : AreaCode,
           OrderType : OrderType,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "get",
            "/service/application/pos/cart/v1.0/shipment",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CartShipmentsResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getShipmentsResponse)
        if err != nil {
           
             return &CartShipmentsResponse{}, common.NewFDKError(err.Error())
            
        }
        return getShipmentsResponse, nil
    }
          
    
  
    
    // UpdateShipments Update shipment delivery type and quantity before checkout
    func (po *PosCart)  UpdateShipments(I bool, P bool, UID int, AddressID int, OrderType string, body  UpdateCartShipmentRequest) (*CartShipmentsResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateShipmentsResponse *CartShipmentsResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          I bool  `url:"i"` 
          P bool  `url:"p"` 
          UID int  `url:"uid"` 
          AddressID int  `url:"address_id"` 
          OrderType string  `url:"order_type"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           I : I,
           P : P,
           UID : UID,
           AddressID : AddressID,
           OrderType : OrderType,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &CartShipmentsResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &CartShipmentsResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "put",
            "/service/application/pos/cart/v1.0/shipment",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CartShipmentsResponse{}, err
            
	    }
        err = json.Unmarshal(response, &updateShipmentsResponse)
        if err != nil {
           
             return &CartShipmentsResponse{}, common.NewFDKError(err.Error())
            
        }
        return updateShipmentsResponse, nil
    }
          
    
  
    
    // CheckoutCart Checkout Cart
    func (po *PosCart)  CheckoutCart(UID int, body  CartPosCheckoutRequest) (*CartCheckoutResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            checkoutCartResponse *CartCheckoutResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID int  `url:"uid"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &CartCheckoutResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &CartCheckoutResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "post",
            "/service/application/pos/cart/v1.0/checkout",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CartCheckoutResponse{}, err
            
	    }
        err = json.Unmarshal(response, &checkoutCartResponse)
        if err != nil {
           
             return &CartCheckoutResponse{}, common.NewFDKError(err.Error())
            
        }
        return checkoutCartResponse, nil
    }
          
    
  
    
    // UpdateCartMeta Update Cart Meta
    func (po *PosCart)  UpdateCartMeta(UID int, body  CartMetaRequest) (*CartMetaResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateCartMetaResponse *CartMetaResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          UID int  `url:"uid"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           UID : UID,
        }
        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &CartMetaResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &CartMetaResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "put",
            "/service/application/pos/cart/v1.0/meta",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CartMetaResponse{}, err
            
	    }
        err = json.Unmarshal(response, &updateCartMetaResponse)
        if err != nil {
           
             return &CartMetaResponse{}, common.NewFDKError(err.Error())
            
        }
        return updateCartMetaResponse, nil
    }
          
    
  
    
    // GetAvailableDeliveryModes Get available delivery modes for cart
    func (po *PosCart)  GetAvailableDeliveryModes(AreaCode string, UID int) (*CartDeliveryModesResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getAvailableDeliveryModesResponse *CartDeliveryModesResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          AreaCode string  `url:"area_code"` 
          UID int  `url:"uid"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           AreaCode : AreaCode,
           UID : UID,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "get",
            fmt.Sprintf("/service/application/pos/cart/v1.0/available-delivery-mode",),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &CartDeliveryModesResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getAvailableDeliveryModesResponse)
        if err != nil {
           
             return &CartDeliveryModesResponse{}, common.NewFDKError(err.Error())
            
        }
        return getAvailableDeliveryModesResponse, nil
    }
          
    
  
    
    // GetStoreAddressByUid Get list of stores for give uids
    func (po *PosCart)  GetStoreAddressByUid(StoreUID int) (*StoreDetailsResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getStoreAddressByUidResponse *StoreDetailsResponse
            
	    )

        
        //Query holds query params
        type XQuery struct { 
          StoreUID int  `url:"store_uid"`  
        }
        //Query params populating to struct
        xQuery := XQuery{  
           StoreUID : StoreUID,
        }
        
         
        
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "get",
            fmt.Sprintf("/service/application/pos/cart/v1.0/store-address",),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &StoreDetailsResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getStoreAddressByUidResponse)
        if err != nil {
           
             return &StoreDetailsResponse{}, common.NewFDKError(err.Error())
            
        }
        return getStoreAddressByUidResponse, nil
    }
          
    
  
    
    // GetCartShareLink Generate Cart sharing link token
    func (po *PosCart)  GetCartShareLink(body  GetShareCartLinkRequest) (*GetShareCartLinkResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getCartShareLinkResponse *GetShareCartLinkResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &GetShareCartLinkResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &GetShareCartLinkResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "post",
            "/service/application/pos/cart/v1.0/share-cart",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &GetShareCartLinkResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getCartShareLinkResponse)
        if err != nil {
           
             return &GetShareCartLinkResponse{}, common.NewFDKError(err.Error())
            
        }
        return getCartShareLinkResponse, nil
    }
          
    
  
    
    // GetCartSharedItems Get shared cart snapshot and cart response
    func (po *PosCart)  GetCartSharedItems(Token string) (*SharedCartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getCartSharedItemsResponse *SharedCartResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "get",
            fmt.Sprintf("/service/application/pos/cart/v1.0/share-cart/%s",Token),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &SharedCartResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getCartSharedItemsResponse)
        if err != nil {
           
             return &SharedCartResponse{}, common.NewFDKError(err.Error())
            
        }
        return getCartSharedItemsResponse, nil
    }
          
    
  
    
    // UpdateCartWithSharedItems Merge or Replace existing cart
    func (po *PosCart)  UpdateCartWithSharedItems(Token string, Action string) (*SharedCartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            updateCartWithSharedItemsResponse *SharedCartResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "post",
            fmt.Sprintf("/service/application/pos/cart/v1.0/share-cart/%s/%s",Token,Action),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &SharedCartResponse{}, err
            
	    }
        err = json.Unmarshal(response, &updateCartWithSharedItemsResponse)
        if err != nil {
           
             return &SharedCartResponse{}, common.NewFDKError(err.Error())
            
        }
        return updateCartWithSharedItemsResponse, nil
    }
          
    

    // Logistic ...
    type Logistic struct {
        config *AppConfig
    }
    // NewLogistic ...
    func NewLogistic(config *AppConfig) *Logistic {
        return &Logistic{config}
    }
    
  
    
    // GetTatProduct Get Tat Product
    func (lo *Logistic)  GetTatProduct(body  GetTatProductReqBody) (*GetTatProductResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getTatProductResponse *GetTatProductResponse
            
	    )

        
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
           return &GetTatProductResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             return &GetTatProductResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            lo.config,
            "post",
            "/service/application/logistics/v1.0",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &GetTatProductResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getTatProductResponse)
        if err != nil {
           
             return &GetTatProductResponse{}, common.NewFDKError(err.Error())
            
        }
        return getTatProductResponse, nil
    }
          
    
  
    
    // GetPincodeCity Get City from Pincode
    func (lo *Logistic)  GetPincodeCity(Pincode string) (*GetPincodeCityResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
            
            getPincodeCityResponse *GetPincodeCityResponse
            
	    )

        
         
        
        
        //API call
        rawRequest = NewRequest(
            lo.config,
            "get",
            fmt.Sprintf("/service/application/logistics/v1.0/pincode/%s",Pincode),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
	 	   
             return &GetPincodeCityResponse{}, err
            
	    }
        err = json.Unmarshal(response, &getPincodeCityResponse)
        if err != nil {
           
             return &GetPincodeCityResponse{}, common.NewFDKError(err.Error())
            
        }
        return getPincodeCityResponse, nil
    }
          
    

