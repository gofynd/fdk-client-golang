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
	
		Common  *Common
	
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
	
		PosCart  *PosCart
	
		Logistic  *Logistic
	
}

// NewAppClient provides application client
func NewAppClient(config *AppConfig) *Client {
		return &Client{
			
				Catalog:  NewCatalog(config),
			
				Cart:  NewCart(config),
			
				Common:  NewCommon(config),
			
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
    func (ca *Catalog)  GetProductDetailBySlug(Slug string) (ProductDetail, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getProductDetailBySlugResponse ProductDetail
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
            return ProductDetail{}, err
	    }
        
        err = json.Unmarshal(response, &getProductDetailBySlugResponse)
        if err != nil {
            return ProductDetail{}, common.NewFDKError(err.Error())
        }
         return getProductDetailBySlugResponse, nil
        
    }
          
    
    
    
  
    
    
    //CatalogGetProductSizesBySlugXQuery holds query params
    type CatalogGetProductSizesBySlugXQuery struct { 
        StoreID float64  `url:"store_id,omitempty"`  
    }
    
    // GetProductSizesBySlug Get the sizes of a product
    func (ca *Catalog)  GetProductSizesBySlug(Slug string, xQuery CatalogGetProductSizesBySlugXQuery) (ProductSizes, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getProductSizesBySlugResponse ProductSizes
	    )

        

        
            
                
            
        

        
        
        
    
         
        
        
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
            return ProductSizes{}, err
	    }
        
        err = json.Unmarshal(response, &getProductSizesBySlugResponse)
        if err != nil {
            return ProductSizes{}, common.NewFDKError(err.Error())
        }
         return getProductSizesBySlugResponse, nil
        
    }
          
    
    
    
  
    
    
    //CatalogGetProductComparisonBySlugsXQuery holds query params
    type CatalogGetProductComparisonBySlugsXQuery struct { 
        Slug []string  `url:"slug,omitempty"`  
    }
    
    // GetProductComparisonBySlugs Compare products
    func (ca *Catalog)  GetProductComparisonBySlugs(xQuery CatalogGetProductComparisonBySlugsXQuery) (ProductsComparisonResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getProductComparisonBySlugsResponse ProductsComparisonResponse
	    )

        

        
            
                
            
        

        
    
         
        
        
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
            return ProductsComparisonResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getProductComparisonBySlugsResponse)
        if err != nil {
            return ProductsComparisonResponse{}, common.NewFDKError(err.Error())
        }
         return getProductComparisonBySlugsResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetSimilarComparisonProductBySlug Get comparison between similar products
    func (ca *Catalog)  GetSimilarComparisonProductBySlug(Slug string) (ProductCompareResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getSimilarComparisonProductBySlugResponse ProductCompareResponse
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
            return ProductCompareResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getSimilarComparisonProductBySlugResponse)
        if err != nil {
            return ProductCompareResponse{}, common.NewFDKError(err.Error())
        }
         return getSimilarComparisonProductBySlugResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetComparedFrequentlyProductBySlug Get comparison between frequently compared products with the given product
    func (ca *Catalog)  GetComparedFrequentlyProductBySlug(Slug string) (ProductFrequentlyComparedSimilarResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getComparedFrequentlyProductBySlugResponse ProductFrequentlyComparedSimilarResponse
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
            return ProductFrequentlyComparedSimilarResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getComparedFrequentlyProductBySlugResponse)
        if err != nil {
            return ProductFrequentlyComparedSimilarResponse{}, common.NewFDKError(err.Error())
        }
         return getComparedFrequentlyProductBySlugResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetProductSimilarByIdentifier Get similar products
    func (ca *Catalog)  GetProductSimilarByIdentifier(Slug string, SimilarType string) (SimilarProductByTypeResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getProductSimilarByIdentifierResponse SimilarProductByTypeResponse
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
            return SimilarProductByTypeResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getProductSimilarByIdentifierResponse)
        if err != nil {
            return SimilarProductByTypeResponse{}, common.NewFDKError(err.Error())
        }
         return getProductSimilarByIdentifierResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetProductVariantsBySlug Get variant of a particular product
    func (ca *Catalog)  GetProductVariantsBySlug(Slug string) (ProductVariantsResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getProductVariantsBySlugResponse ProductVariantsResponse
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
            return ProductVariantsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getProductVariantsBySlugResponse)
        if err != nil {
            return ProductVariantsResponse{}, common.NewFDKError(err.Error())
        }
         return getProductVariantsBySlugResponse, nil
        
    }
          
    
    
    
  
    
    
    //CatalogGetProductStockByIdsXQuery holds query params
    type CatalogGetProductStockByIdsXQuery struct { 
        ItemID string  `url:"item_id,omitempty"` 
        Alu string  `url:"alu,omitempty"` 
        SkuCode string  `url:"sku_code,omitempty"` 
        Ean string  `url:"ean,omitempty"` 
        Upc string  `url:"upc,omitempty"`  
    }
    
    // GetProductStockByIds Get the stock of a product
    func (ca *Catalog)  GetProductStockByIds(xQuery CatalogGetProductStockByIdsXQuery) (ProductStockStatusResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getProductStockByIdsResponse ProductStockStatusResponse
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return ProductStockStatusResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getProductStockByIdsResponse)
        if err != nil {
            return ProductStockStatusResponse{}, common.NewFDKError(err.Error())
        }
         return getProductStockByIdsResponse, nil
        
    }
          
    
    
    
  
    
    
    //CatalogGetProductStockForTimeByIdsXQuery holds query params
    type CatalogGetProductStockForTimeByIdsXQuery struct { 
        Timestamp string  `url:"timestamp,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        PageID string  `url:"page_id,omitempty"`  
    }
    
    // GetProductStockForTimeByIds Get the stock of a product
    func (ca *Catalog)  GetProductStockForTimeByIds(xQuery CatalogGetProductStockForTimeByIdsXQuery) (ProductStockPolling, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getProductStockForTimeByIdsResponse ProductStockPolling
	    )

        

        
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return ProductStockPolling{}, err
	    }
        
        err = json.Unmarshal(response, &getProductStockForTimeByIdsResponse)
        if err != nil {
            return ProductStockPolling{}, common.NewFDKError(err.Error())
        }
         return getProductStockForTimeByIdsResponse, nil
        
    }
          
            
            
            
            
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                    
                        
                    
                
            
            // GetProductStockForTimeByIdsPaginator Get the stock of a product  
            func (ca *Catalog)  GetProductStockForTimeByIdsPaginator( xQuery CatalogGetProductStockForTimeByIdsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 xQuery.PageID = paginator.NextID
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := ca.GetProductStockForTimeByIds(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
       
    
    
    
  
    
    
    //CatalogGetProductsXQuery holds query params
    type CatalogGetProductsXQuery struct { 
        Q string  `url:"q,omitempty"` 
        F string  `url:"f,omitempty"` 
        Filters bool  `url:"filters,omitempty"` 
        SortOn string  `url:"sort_on,omitempty"` 
        PageID string  `url:"page_id,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"` 
        PageType string  `url:"page_type,omitempty"`  
    }
    
    // GetProducts Get all the products
    func (ca *Catalog)  GetProducts(xQuery CatalogGetProductsXQuery) (ProductListingResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getProductsResponse ProductListingResponse
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return ProductListingResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getProductsResponse)
        if err != nil {
            return ProductListingResponse{}, common.NewFDKError(err.Error())
        }
         return getProductsResponse, nil
        
    }
          
            
            
            
            
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                    
                        
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                    
                
                    
                    
                    
                    
                
            
            // GetProductsPaginator Get all the products  
            func (ca *Catalog)  GetProductsPaginator( xQuery CatalogGetProductsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 xQuery.PageID = paginator.NextID
                 
                 
                 
                 
                 
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 xQuery.PageType = "cursor"
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := ca.GetProducts(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
       
    
    
    
  
    
    
    //CatalogGetBrandsXQuery holds query params
    type CatalogGetBrandsXQuery struct { 
        Department string  `url:"department,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetBrands Get all the brands
    func (ca *Catalog)  GetBrands(xQuery CatalogGetBrandsXQuery) (BrandListingResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getBrandsResponse BrandListingResponse
	    )

        

        
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return BrandListingResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getBrandsResponse)
        if err != nil {
            return BrandListingResponse{}, common.NewFDKError(err.Error())
        }
         return getBrandsResponse, nil
        
    }
          
            
            
            
            
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                    
                
                    
                    
                    
                        
                    
                    
                
            
            // GetBrandsPaginator Get all the brands  
            func (ca *Catalog)  GetBrandsPaginator( xQuery CatalogGetBrandsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                 
                 
                 
                 
                 
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := ca.GetBrands(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
       
    
    
    
  
    
    
    // GetBrandDetailBySlug Get metadata of a brand
    func (ca *Catalog)  GetBrandDetailBySlug(Slug string) (BrandDetailResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getBrandDetailBySlugResponse BrandDetailResponse
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
            return BrandDetailResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getBrandDetailBySlugResponse)
        if err != nil {
            return BrandDetailResponse{}, common.NewFDKError(err.Error())
        }
         return getBrandDetailBySlugResponse, nil
        
    }
          
    
    
    
  
    
    
    //CatalogGetCategoriesXQuery holds query params
    type CatalogGetCategoriesXQuery struct { 
        Department string  `url:"department,omitempty"`  
    }
    
    // GetCategories List all the categories
    func (ca *Catalog)  GetCategories(xQuery CatalogGetCategoriesXQuery) (CategoryListingResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getCategoriesResponse CategoryListingResponse
	    )

        

        
            
                
            
        

        
    
         
        
        
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
            return CategoryListingResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCategoriesResponse)
        if err != nil {
            return CategoryListingResponse{}, common.NewFDKError(err.Error())
        }
         return getCategoriesResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetCategoryDetailBySlug Get metadata of a category
    func (ca *Catalog)  GetCategoryDetailBySlug(Slug string) (CategoryMetaResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getCategoryDetailBySlugResponse CategoryMetaResponse
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
            return CategoryMetaResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCategoryDetailBySlugResponse)
        if err != nil {
            return CategoryMetaResponse{}, common.NewFDKError(err.Error())
        }
         return getCategoryDetailBySlugResponse, nil
        
    }
          
    
    
    
  
    
    
    //CatalogGetHomeProductsXQuery holds query params
    type CatalogGetHomeProductsXQuery struct { 
        SortOn string  `url:"sort_on,omitempty"` 
        PageID string  `url:"page_id,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetHomeProducts List the products
    func (ca *Catalog)  GetHomeProducts(xQuery CatalogGetHomeProductsXQuery) (HomeListingResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getHomeProductsResponse HomeListingResponse
	    )

        

        
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return HomeListingResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getHomeProductsResponse)
        if err != nil {
            return HomeListingResponse{}, common.NewFDKError(err.Error())
        }
         return getHomeProductsResponse, nil
        
    }
          
            
            
            
            
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                    
                        
                    
                
                    
                    
                    
                        
                    
                    
                
            
            // GetHomeProductsPaginator List the products  
            func (ca *Catalog)  GetHomeProductsPaginator( xQuery CatalogGetHomeProductsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                 
                 
                 
                 
                 
                 
                 
                 xQuery.PageID = paginator.NextID
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := ca.GetHomeProducts(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
       
    
    
    
  
    
    
    // GetDepartments List all the departments
    func (ca *Catalog)  GetDepartments() (DepartmentResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getDepartmentsResponse DepartmentResponse
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
            return DepartmentResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getDepartmentsResponse)
        if err != nil {
            return DepartmentResponse{}, common.NewFDKError(err.Error())
        }
         return getDepartmentsResponse, nil
        
    }
          
    
    
    
  
    
    
    //CatalogGetSearchResultsXQuery holds query params
    type CatalogGetSearchResultsXQuery struct { 
        Q string  `url:"q,omitempty"`  
    }
    
    // GetSearchResults Get relevant suggestions for a search query
    func (ca *Catalog)  GetSearchResults(xQuery CatalogGetSearchResultsXQuery) (AutoCompleteResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getSearchResultsResponse AutoCompleteResponse
	    )

        

        
            
                
            
        

        
    
         
        
        
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
            return AutoCompleteResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getSearchResultsResponse)
        if err != nil {
            return AutoCompleteResponse{}, common.NewFDKError(err.Error())
        }
         return getSearchResultsResponse, nil
        
    }
          
    
    
    
  
    
    
    //CatalogGetCollectionsXQuery holds query params
    type CatalogGetCollectionsXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Tag []string  `url:"tag,omitempty"`  
    }
    
    // GetCollections List all the collections
    func (ca *Catalog)  GetCollections(xQuery CatalogGetCollectionsXQuery) (GetCollectionListingResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getCollectionsResponse GetCollectionListingResponse
	    )

        

        
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return GetCollectionListingResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCollectionsResponse)
        if err != nil {
            return GetCollectionListingResponse{}, common.NewFDKError(err.Error())
        }
         return getCollectionsResponse, nil
        
    }
          
            
            
            
            
                
                    
                    
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
            
            // GetCollectionsPaginator List all the collections  
            func (ca *Catalog)  GetCollectionsPaginator( xQuery CatalogGetCollectionsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                 
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := ca.GetCollections(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
       
    
    
    
  
    
    
    //CatalogGetCollectionItemsBySlugXQuery holds query params
    type CatalogGetCollectionItemsBySlugXQuery struct { 
        F string  `url:"f,omitempty"` 
        Filters bool  `url:"filters,omitempty"` 
        SortOn string  `url:"sort_on,omitempty"` 
        PageID string  `url:"page_id,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetCollectionItemsBySlug Get the items in a collection
    func (ca *Catalog)  GetCollectionItemsBySlug(Slug string, xQuery CatalogGetCollectionItemsBySlugXQuery) (ProductListingResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getCollectionItemsBySlugResponse ProductListingResponse
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
        

        
        
        
    
         
        
        
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
            return ProductListingResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCollectionItemsBySlugResponse)
        if err != nil {
            return ProductListingResponse{}, common.NewFDKError(err.Error())
        }
         return getCollectionItemsBySlugResponse, nil
        
    }
          
            
            
            
            
                
                    
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                    
                        
                    
                
                    
                    
                    
                        
                    
                    
                
            
            // GetCollectionItemsBySlugPaginator Get the items in a collection  
            func (ca *Catalog)  GetCollectionItemsBySlugPaginator(Slug string  ,  xQuery CatalogGetCollectionItemsBySlugXQuery ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 xQuery.PageID = paginator.NextID
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := ca.GetCollectionItemsBySlug(Slug, xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
       
    
    
    
  
    
    
    // GetCollectionDetailBySlug Get a particular collection
    func (ca *Catalog)  GetCollectionDetailBySlug(Slug string) (CollectionDetailResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getCollectionDetailBySlugResponse CollectionDetailResponse
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
            return CollectionDetailResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCollectionDetailBySlugResponse)
        if err != nil {
            return CollectionDetailResponse{}, common.NewFDKError(err.Error())
        }
         return getCollectionDetailBySlugResponse, nil
        
    }
          
    
    
    
  
    
    
    //CatalogGetFollowedListingXQuery holds query params
    type CatalogGetFollowedListingXQuery struct { 
        PageID string  `url:"page_id,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetFollowedListing Get a list of followed Products, Brands, Collections
    func (ca *Catalog)  GetFollowedListing(CollectionType string, xQuery CatalogGetFollowedListingXQuery) (GetFollowListingResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getFollowedListingResponse GetFollowListingResponse
	    )

        

        
            
                
            
                
            
        

        
        
        
    
         
        
        
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
            return GetFollowListingResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getFollowedListingResponse)
        if err != nil {
            return GetFollowListingResponse{}, common.NewFDKError(err.Error())
        }
         return getFollowedListingResponse, nil
        
    }
          
            
            
            
            
                
                    
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                    
                        
                    
                
                    
                    
                    
                        
                    
                    
                
            
            // GetFollowedListingPaginator Get a list of followed Products, Brands, Collections  
            func (ca *Catalog)  GetFollowedListingPaginator(CollectionType string  ,  xQuery CatalogGetFollowedListingXQuery ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                 
                 
                 
                 xQuery.PageID = paginator.NextID
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := ca.GetFollowedListing(CollectionType, xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
       
    
    
    
  
    
    
    // UnfollowById Unfollow an entity (product/brand/collection)
    func (ca *Catalog)  UnfollowById(CollectionType string, CollectionID string) (FollowPostResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             unfollowByIdResponse FollowPostResponse
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
            return FollowPostResponse{}, err
	    }
        
        err = json.Unmarshal(response, &unfollowByIdResponse)
        if err != nil {
            return FollowPostResponse{}, common.NewFDKError(err.Error())
        }
         return unfollowByIdResponse, nil
        
    }
          
    
    
    
  
    
    
    // FollowById Follow an entity (product/brand/collection)
    func (ca *Catalog)  FollowById(CollectionType string, CollectionID string) (FollowPostResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             followByIdResponse FollowPostResponse
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
            return FollowPostResponse{}, err
	    }
        
        err = json.Unmarshal(response, &followByIdResponse)
        if err != nil {
            return FollowPostResponse{}, common.NewFDKError(err.Error())
        }
         return followByIdResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetFollowerCountById Get Follow Count
    func (ca *Catalog)  GetFollowerCountById(CollectionType string, CollectionID string) (FollowerCountResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getFollowerCountByIdResponse FollowerCountResponse
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
            return FollowerCountResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getFollowerCountByIdResponse)
        if err != nil {
            return FollowerCountResponse{}, common.NewFDKError(err.Error())
        }
         return getFollowerCountByIdResponse, nil
        
    }
          
    
    
    
  
    
    
    //CatalogGetFollowIdsXQuery holds query params
    type CatalogGetFollowIdsXQuery struct { 
        CollectionType string  `url:"collection_type,omitempty"`  
    }
    
    // GetFollowIds Get the IDs of followed products, brands and collections.
    func (ca *Catalog)  GetFollowIds(xQuery CatalogGetFollowIdsXQuery) (FollowIdsResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getFollowIdsResponse FollowIdsResponse
	    )

        

        
            
                
            
        

        
    
         
        
        
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
            return FollowIdsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getFollowIdsResponse)
        if err != nil {
            return FollowIdsResponse{}, common.NewFDKError(err.Error())
        }
         return getFollowIdsResponse, nil
        
    }
          
    
    
    
  
    
    
    //CatalogGetStoresXQuery holds query params
    type CatalogGetStoresXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Q string  `url:"q,omitempty"` 
        City string  `url:"city,omitempty"` 
        Range float64  `url:"range,omitempty"` 
        Latitude float64  `url:"latitude,omitempty"` 
        Longitude float64  `url:"longitude,omitempty"`  
    }
    
    // GetStores Get store meta information.
    func (ca *Catalog)  GetStores(xQuery CatalogGetStoresXQuery) (StoreListingResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getStoresResponse StoreListingResponse
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return StoreListingResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getStoresResponse)
        if err != nil {
            return StoreListingResponse{}, common.NewFDKError(err.Error())
        }
         return getStoresResponse, nil
        
    }
          
            
            
            
            
                
                    
                    
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
            
            // GetStoresPaginator Get store meta information.  
            func (ca *Catalog)  GetStoresPaginator( xQuery CatalogGetStoresXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                 
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := ca.GetStores(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
       
    
    
    
  
    
    
    //CatalogGetInStockLocationsXQuery holds query params
    type CatalogGetInStockLocationsXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Q string  `url:"q,omitempty"` 
        City string  `url:"city,omitempty"` 
        Range float64  `url:"range,omitempty"` 
        Latitude float64  `url:"latitude,omitempty"` 
        Longitude float64  `url:"longitude,omitempty"`  
    }
    
    // GetInStockLocations Get store meta information.
    func (ca *Catalog)  GetInStockLocations(xQuery CatalogGetInStockLocationsXQuery) (ApplicationStoreListing, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getInStockLocationsResponse ApplicationStoreListing
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            "/service/application/catalog/v1.0/in-stock/locations/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ApplicationStoreListing{}, err
	    }
        
        err = json.Unmarshal(response, &getInStockLocationsResponse)
        if err != nil {
            return ApplicationStoreListing{}, common.NewFDKError(err.Error())
        }
         return getInStockLocationsResponse, nil
        
    }
          
            
            
            
            
                
                    
                    
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
            
            // GetInStockLocationsPaginator Get store meta information.  
            func (ca *Catalog)  GetInStockLocationsPaginator( xQuery CatalogGetInStockLocationsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                 
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := ca.GetInStockLocations(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
       
    
    
    
  
    
    
    // GetLocationDetailsById Get store meta information.
    func (ca *Catalog)  GetLocationDetailsById(LocationID float64) (StoreDetails, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getLocationDetailsByIdResponse StoreDetails
	    )

        

        

        
        
        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/catalog/v1.0/locations/undefined/",LocationID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return StoreDetails{}, err
	    }
        
        err = json.Unmarshal(response, &getLocationDetailsByIdResponse)
        if err != nil {
            return StoreDetails{}, common.NewFDKError(err.Error())
        }
         return getLocationDetailsByIdResponse, nil
        
    }
          
    
    
    
  
    
    
    //CatalogGetProductBundlesBySlugXQuery holds query params
    type CatalogGetProductBundlesBySlugXQuery struct { 
        Slug string  `url:"slug,omitempty"` 
        ID string  `url:"id,omitempty"`  
    }
    
    // GetProductBundlesBySlug Get product bundles
    func (ca *Catalog)  GetProductBundlesBySlug(xQuery CatalogGetProductBundlesBySlugXQuery) (ProductBundle, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getProductBundlesBySlugResponse ProductBundle
	    )

        

        
            
                
            
                
            
        

        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            "/service/application/catalog/v1.0/product-grouping/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ProductBundle{}, err
	    }
        
        err = json.Unmarshal(response, &getProductBundlesBySlugResponse)
        if err != nil {
            return ProductBundle{}, common.NewFDKError(err.Error())
        }
         return getProductBundlesBySlugResponse, nil
        
    }
          
    
    
    
  
    
    
    //CatalogGetProductPriceBySlugXQuery holds query params
    type CatalogGetProductPriceBySlugXQuery struct { 
        StoreID float64  `url:"store_id,omitempty"` 
        Pincode string  `url:"pincode,omitempty"` 
        Depth string  `url:"depth,omitempty"` 
        Strategy string  `url:"strategy,omitempty"`  
    }
    
    // GetProductPriceBySlug Get the price of a product size at a PIN Code
    func (ca *Catalog)  GetProductPriceBySlug(Slug string, Size string, xQuery CatalogGetProductPriceBySlugXQuery) (ProductSizePriceResponseV2, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getProductPriceBySlugResponse ProductSizePriceResponseV2
	    )

        

        
            
                
            
                
            
                
            
                
            
        

        
        
        
        
        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/catalog/v2.0/products/%s/sizes/%s/price/",Slug,Size),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ProductSizePriceResponseV2{}, err
	    }
        
        err = json.Unmarshal(response, &getProductPriceBySlugResponse)
        if err != nil {
            return ProductSizePriceResponseV2{}, common.NewFDKError(err.Error())
        }
         return getProductPriceBySlugResponse, nil
        
    }
          
    
    
    
  
    
    
    //CatalogGetProductSellersBySlugXQuery holds query params
    type CatalogGetProductSellersBySlugXQuery struct { 
        Pincode string  `url:"pincode,omitempty"` 
        Strategy string  `url:"strategy,omitempty"` 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Depth string  `url:"depth,omitempty"`  
    }
    
    // GetProductSellersBySlug Get the sellers of a product size at a PIN Code
    func (ca *Catalog)  GetProductSellersBySlug(Slug string, Size string, xQuery CatalogGetProductSellersBySlugXQuery) (ProductSizeSellersResponseV2, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getProductSellersBySlugResponse ProductSizeSellersResponseV2
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
        

        
        
        
        
        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/catalog/v2.0/products/%s/sizes/%s/sellers/",Slug,Size),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ProductSizeSellersResponseV2{}, err
	    }
        
        err = json.Unmarshal(response, &getProductSellersBySlugResponse)
        if err != nil {
            return ProductSizeSellersResponseV2{}, common.NewFDKError(err.Error())
        }
         return getProductSellersBySlugResponse, nil
        
    }
          
            
            
            
            
                
                    
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
            
            // GetProductSellersBySlugPaginator Get the sellers of a product size at a PIN Code  
            func (ca *Catalog)  GetProductSellersBySlugPaginator(Slug string  , Size string  ,  xQuery CatalogGetProductSellersBySlugXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := ca.GetProductSellersBySlug(Slug, Size, xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
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
    
    
    
  
    
    
    //CartGetCartXQuery holds query params
    type CartGetCartXQuery struct { 
        ID string  `url:"id,omitempty"` 
        I bool  `url:"i,omitempty"` 
        B bool  `url:"b,omitempty"` 
        AssignCardID float64  `url:"assign_card_id,omitempty"` 
        AreaCode string  `url:"area_code,omitempty"`  
    }
    
    // GetCart Fetch all items added to the cart
    func (ca *Cart)  GetCart(xQuery CartGetCartXQuery) (CartDetailResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getCartResponse CartDetailResponse
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return CartDetailResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCartResponse)
        if err != nil {
            return CartDetailResponse{}, common.NewFDKError(err.Error())
        }
         return getCartResponse, nil
        
    }
          
    
    
    
  
    
    
    //CartGetCartLastModifiedXQuery holds query params
    type CartGetCartLastModifiedXQuery struct { 
        ID string  `url:"id,omitempty"`  
    }
    
    // GetCartLastModified Fetch last-modified timestamp
    func (ca *Cart)  GetCartLastModified(xQuery CartGetCartLastModifiedXQuery) (interface{}, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             
	    )

        

        
            
                
            
        

        
    
         
        
        
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
            return []byte{}, err
	    }
        
        return response, nil
        
    }
          
    
    
    
  
    
    
    //CartAddItemsXQuery holds query params
    type CartAddItemsXQuery struct { 
        I bool  `url:"i,omitempty"` 
        B bool  `url:"b,omitempty"` 
        AreaCode string  `url:"area_code,omitempty"`  
    }
    
    // AddItems Add items to cart
    func (ca *Cart)  AddItems(xQuery CartAddItemsXQuery, body  AddCartRequest) (AddCartDetailResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             addItemsResponse AddCartDetailResponse
	    )

        
            
        

        
            
                
            
                
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return AddCartDetailResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return AddCartDetailResponse{}, common.NewFDKError(err.Error())
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
            return AddCartDetailResponse{}, err
	    }
        
        err = json.Unmarshal(response, &addItemsResponse)
        if err != nil {
            return AddCartDetailResponse{}, common.NewFDKError(err.Error())
        }
         return addItemsResponse, nil
        
    }
          
    
    
    
  
    
    
    //CartUpdateCartXQuery holds query params
    type CartUpdateCartXQuery struct { 
        ID string  `url:"id,omitempty"` 
        I bool  `url:"i,omitempty"` 
        B bool  `url:"b,omitempty"` 
        AreaCode string  `url:"area_code,omitempty"`  
    }
    
    // UpdateCart Update items in the cart
    func (ca *Cart)  UpdateCart(xQuery CartUpdateCartXQuery, body  UpdateCartRequest) (UpdateCartDetailResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             updateCartResponse UpdateCartDetailResponse
	    )

        
            
        
            
        

        
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return UpdateCartDetailResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return UpdateCartDetailResponse{}, common.NewFDKError(err.Error())
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
            return UpdateCartDetailResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateCartResponse)
        if err != nil {
            return UpdateCartDetailResponse{}, common.NewFDKError(err.Error())
        }
         return updateCartResponse, nil
        
    }
          
    
    
    
  
    
    
    //CartGetItemCountXQuery holds query params
    type CartGetItemCountXQuery struct { 
        ID string  `url:"id,omitempty"`  
    }
    
    // GetItemCount Count items in the cart
    func (ca *Cart)  GetItemCount(xQuery CartGetItemCountXQuery) (CartItemCountResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getItemCountResponse CartItemCountResponse
	    )

        

        
            
                
            
        

        
    
         
        
        
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
            return CartItemCountResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getItemCountResponse)
        if err != nil {
            return CartItemCountResponse{}, common.NewFDKError(err.Error())
        }
         return getItemCountResponse, nil
        
    }
          
    
    
    
  
    
    
    //CartGetCouponsXQuery holds query params
    type CartGetCouponsXQuery struct { 
        ID string  `url:"id,omitempty"`  
    }
    
    // GetCoupons Fetch Coupon
    func (ca *Cart)  GetCoupons(xQuery CartGetCouponsXQuery) (GetCouponResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getCouponsResponse GetCouponResponse
	    )

        

        
            
                
            
        

        
    
         
        
        
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
            return GetCouponResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCouponsResponse)
        if err != nil {
            return GetCouponResponse{}, common.NewFDKError(err.Error())
        }
         return getCouponsResponse, nil
        
    }
          
    
    
    
  
    
    
    //CartApplyCouponXQuery holds query params
    type CartApplyCouponXQuery struct { 
        I bool  `url:"i,omitempty"` 
        B bool  `url:"b,omitempty"` 
        P bool  `url:"p,omitempty"` 
        ID string  `url:"id,omitempty"`  
    }
    
    // ApplyCoupon Apply Coupon
    func (ca *Cart)  ApplyCoupon(xQuery CartApplyCouponXQuery, body  ApplyCouponRequest) (CartDetailResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             applyCouponResponse CartDetailResponse
	    )

        
            
        

        
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return CartDetailResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return CartDetailResponse{}, common.NewFDKError(err.Error())
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
            return CartDetailResponse{}, err
	    }
        
        err = json.Unmarshal(response, &applyCouponResponse)
        if err != nil {
            return CartDetailResponse{}, common.NewFDKError(err.Error())
        }
         return applyCouponResponse, nil
        
    }
          
    
    
    
  
    
    
    //CartRemoveCouponXQuery holds query params
    type CartRemoveCouponXQuery struct { 
        ID string  `url:"id,omitempty"`  
    }
    
    // RemoveCoupon Remove Coupon Applied
    func (ca *Cart)  RemoveCoupon(xQuery CartRemoveCouponXQuery) (CartDetailResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             removeCouponResponse CartDetailResponse
	    )

        

        
            
                
            
        

        
    
         
        
        
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
            return CartDetailResponse{}, err
	    }
        
        err = json.Unmarshal(response, &removeCouponResponse)
        if err != nil {
            return CartDetailResponse{}, common.NewFDKError(err.Error())
        }
         return removeCouponResponse, nil
        
    }
          
    
    
    
  
    
    
    //CartGetBulkDiscountOffersXQuery holds query params
    type CartGetBulkDiscountOffersXQuery struct { 
        ItemID float64  `url:"item_id,omitempty"` 
        ArticleID string  `url:"article_id,omitempty"` 
        UID float64  `url:"uid,omitempty"` 
        Slug string  `url:"slug,omitempty"`  
    }
    
    // GetBulkDiscountOffers Get discount offers based on quantity
    func (ca *Cart)  GetBulkDiscountOffers(xQuery CartGetBulkDiscountOffersXQuery) (BulkPriceResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getBulkDiscountOffersResponse BulkPriceResponse
	    )

        

        
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return BulkPriceResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getBulkDiscountOffersResponse)
        if err != nil {
            return BulkPriceResponse{}, common.NewFDKError(err.Error())
        }
         return getBulkDiscountOffersResponse, nil
        
    }
          
    
    
    
  
    
    
    //CartApplyRewardPointsXQuery holds query params
    type CartApplyRewardPointsXQuery struct { 
        ID string  `url:"id,omitempty"` 
        I bool  `url:"i,omitempty"` 
        B bool  `url:"b,omitempty"`  
    }
    
    // ApplyRewardPoints Apply reward points at cart
    func (ca *Cart)  ApplyRewardPoints(xQuery CartApplyRewardPointsXQuery, body  RewardPointRequest) (CartDetailResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             applyRewardPointsResponse CartDetailResponse
	    )

        
            
        

        
            
                
            
                
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return CartDetailResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return CartDetailResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "post",
            "/service/application/cart/v1.0/redeem/points/",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return CartDetailResponse{}, err
	    }
        
        err = json.Unmarshal(response, &applyRewardPointsResponse)
        if err != nil {
            return CartDetailResponse{}, common.NewFDKError(err.Error())
        }
         return applyRewardPointsResponse, nil
        
    }
          
    
    
    
  
    
    
    //CartGetAddressesXQuery holds query params
    type CartGetAddressesXQuery struct { 
        CartID string  `url:"cart_id,omitempty"` 
        MobileNo string  `url:"mobile_no,omitempty"` 
        CheckoutMode string  `url:"checkout_mode,omitempty"` 
        Tags string  `url:"tags,omitempty"` 
        IsDefault bool  `url:"is_default,omitempty"`  
    }
    
    // GetAddresses Fetch address
    func (ca *Cart)  GetAddresses(xQuery CartGetAddressesXQuery) (GetAddressesResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getAddressesResponse GetAddressesResponse
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return GetAddressesResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getAddressesResponse)
        if err != nil {
            return GetAddressesResponse{}, common.NewFDKError(err.Error())
        }
         return getAddressesResponse, nil
        
    }
          
    
    
    
  
    
    
    // AddAddress Add address to an account
    func (ca *Cart)  AddAddress(body  Address) (SaveAddressResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             addAddressResponse SaveAddressResponse
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return SaveAddressResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return SaveAddressResponse{}, common.NewFDKError(err.Error())
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
            return SaveAddressResponse{}, err
	    }
        
        err = json.Unmarshal(response, &addAddressResponse)
        if err != nil {
            return SaveAddressResponse{}, common.NewFDKError(err.Error())
        }
         return addAddressResponse, nil
        
    }
          
    
    
    
  
    
    
    //CartGetAddressByIdXQuery holds query params
    type CartGetAddressByIdXQuery struct { 
        CartID string  `url:"cart_id,omitempty"` 
        MobileNo string  `url:"mobile_no,omitempty"` 
        CheckoutMode string  `url:"checkout_mode,omitempty"` 
        Tags string  `url:"tags,omitempty"` 
        IsDefault bool  `url:"is_default,omitempty"`  
    }
    
    // GetAddressById Fetch a single address by its ID
    func (ca *Cart)  GetAddressById(ID string, xQuery CartGetAddressByIdXQuery) (Address, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getAddressByIdResponse Address
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
        

        
        
        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/cart/v1.0/address/%s",ID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return Address{}, err
	    }
        
        err = json.Unmarshal(response, &getAddressByIdResponse)
        if err != nil {
            return Address{}, common.NewFDKError(err.Error())
        }
         return getAddressByIdResponse, nil
        
    }
          
    
    
    
  
    
    
    // UpdateAddress Update address added to an account
    func (ca *Cart)  UpdateAddress(ID string, body  Address) (UpdateAddressResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             updateAddressResponse UpdateAddressResponse
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

        

        
        
        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return UpdateAddressResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return UpdateAddressResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "put",
            fmt.Sprintf("/service/application/cart/v1.0/address/%s",ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return UpdateAddressResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateAddressResponse)
        if err != nil {
            return UpdateAddressResponse{}, common.NewFDKError(err.Error())
        }
         return updateAddressResponse, nil
        
    }
          
    
    
    
  
    
    
    // RemoveAddress Remove address associated with an account
    func (ca *Cart)  RemoveAddress(ID string) (DeleteAddressResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             removeAddressResponse DeleteAddressResponse
	    )

        

        

        
        
        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "delete",
            fmt.Sprintf("/service/application/cart/v1.0/address/%s",ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return DeleteAddressResponse{}, err
	    }
        
        err = json.Unmarshal(response, &removeAddressResponse)
        if err != nil {
            return DeleteAddressResponse{}, common.NewFDKError(err.Error())
        }
         return removeAddressResponse, nil
        
    }
          
    
    
    
  
    
    
    //CartSelectAddressXQuery holds query params
    type CartSelectAddressXQuery struct { 
        CartID string  `url:"cart_id,omitempty"` 
        I bool  `url:"i,omitempty"` 
        B bool  `url:"b,omitempty"`  
    }
    
    // SelectAddress Select an address from available addresses
    func (ca *Cart)  SelectAddress(xQuery CartSelectAddressXQuery, body  SelectCartAddressRequest) (CartDetailResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             selectAddressResponse CartDetailResponse
	    )

        
            
        
            
        
            
        

        
            
                
            
                
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return CartDetailResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return CartDetailResponse{}, common.NewFDKError(err.Error())
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
            return CartDetailResponse{}, err
	    }
        
        err = json.Unmarshal(response, &selectAddressResponse)
        if err != nil {
            return CartDetailResponse{}, common.NewFDKError(err.Error())
        }
         return selectAddressResponse, nil
        
    }
          
    
    
    
  
    
    
    //CartSelectPaymentModeXQuery holds query params
    type CartSelectPaymentModeXQuery struct { 
        ID string  `url:"id,omitempty"`  
    }
    
    // SelectPaymentMode Update cart payment
    func (ca *Cart)  SelectPaymentMode(xQuery CartSelectPaymentModeXQuery, body  UpdateCartPaymentRequest) (CartDetailResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             selectPaymentModeResponse CartDetailResponse
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        

        
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return CartDetailResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return CartDetailResponse{}, common.NewFDKError(err.Error())
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
            return CartDetailResponse{}, err
	    }
        
        err = json.Unmarshal(response, &selectPaymentModeResponse)
        if err != nil {
            return CartDetailResponse{}, common.NewFDKError(err.Error())
        }
         return selectPaymentModeResponse, nil
        
    }
          
    
    
    
  
    
    
    //CartValidateCouponForPaymentXQuery holds query params
    type CartValidateCouponForPaymentXQuery struct { 
        ID string  `url:"id,omitempty"` 
        AddressID string  `url:"address_id,omitempty"` 
        PaymentMode string  `url:"payment_mode,omitempty"` 
        PaymentIdentifier string  `url:"payment_identifier,omitempty"` 
        AggregatorName string  `url:"aggregator_name,omitempty"` 
        MerchantCode string  `url:"merchant_code,omitempty"`  
    }
    
    // ValidateCouponForPayment Verify the coupon eligibility against the payment mode
    func (ca *Cart)  ValidateCouponForPayment(xQuery CartValidateCouponForPaymentXQuery) (PaymentCouponValidate, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             validateCouponForPaymentResponse PaymentCouponValidate
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return PaymentCouponValidate{}, err
	    }
        
        err = json.Unmarshal(response, &validateCouponForPaymentResponse)
        if err != nil {
            return PaymentCouponValidate{}, common.NewFDKError(err.Error())
        }
         return validateCouponForPaymentResponse, nil
        
    }
          
    
    
    
  
    
    
    //CartGetShipmentsXQuery holds query params
    type CartGetShipmentsXQuery struct { 
        P bool  `url:"p,omitempty"` 
        ID string  `url:"id,omitempty"` 
        AddressID string  `url:"address_id,omitempty"` 
        AreaCode string  `url:"area_code,omitempty"`  
    }
    
    // GetShipments Get delivery date and options before checkout
    func (ca *Cart)  GetShipments(xQuery CartGetShipmentsXQuery) (CartShipmentsResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getShipmentsResponse CartShipmentsResponse
	    )

        

        
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return CartShipmentsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getShipmentsResponse)
        if err != nil {
            return CartShipmentsResponse{}, common.NewFDKError(err.Error())
        }
         return getShipmentsResponse, nil
        
    }
          
    
    
    
  
    
    
    // CheckoutCart Checkout all items in the cart
    func (ca *Cart)  CheckoutCart(body  CartCheckoutDetailRequest) (CartCheckoutResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             checkoutCartResponse CartCheckoutResponse
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return CartCheckoutResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return CartCheckoutResponse{}, common.NewFDKError(err.Error())
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
            return CartCheckoutResponse{}, err
	    }
        
        err = json.Unmarshal(response, &checkoutCartResponse)
        if err != nil {
            return CartCheckoutResponse{}, common.NewFDKError(err.Error())
        }
         return checkoutCartResponse, nil
        
    }
          
    
    
    
  
    
    
    //CartUpdateCartMetaXQuery holds query params
    type CartUpdateCartMetaXQuery struct { 
        ID string  `url:"id,omitempty"`  
    }
    
    // UpdateCartMeta Update the cart meta
    func (ca *Cart)  UpdateCartMeta(xQuery CartUpdateCartMetaXQuery, body  CartMetaRequest) (CartMetaResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             updateCartMetaResponse CartMetaResponse
	    )

        
            
        
            
        
            
        
            
        

        
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return CartMetaResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return CartMetaResponse{}, common.NewFDKError(err.Error())
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
            return CartMetaResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateCartMetaResponse)
        if err != nil {
            return CartMetaResponse{}, common.NewFDKError(err.Error())
        }
         return updateCartMetaResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetCartShareLink Generate token for sharing the cart
    func (ca *Cart)  GetCartShareLink(body  GetShareCartLinkRequest) (GetShareCartLinkResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getCartShareLinkResponse GetShareCartLinkResponse
	    )

        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return GetShareCartLinkResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return GetShareCartLinkResponse{}, common.NewFDKError(err.Error())
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
            return GetShareCartLinkResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCartShareLinkResponse)
        if err != nil {
            return GetShareCartLinkResponse{}, common.NewFDKError(err.Error())
        }
         return getCartShareLinkResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetCartSharedItems Get details of a shared cart
    func (ca *Cart)  GetCartSharedItems(Token string) (SharedCartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getCartSharedItemsResponse SharedCartResponse
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
            return SharedCartResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCartSharedItemsResponse)
        if err != nil {
            return SharedCartResponse{}, common.NewFDKError(err.Error())
        }
         return getCartSharedItemsResponse, nil
        
    }
          
    
    
    
  
    
    
    // UpdateCartWithSharedItems Merge or replace existing cart
    func (ca *Cart)  UpdateCartWithSharedItems(Token string, Action string) (SharedCartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             updateCartWithSharedItemsResponse SharedCartResponse
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
            return SharedCartResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateCartWithSharedItemsResponse)
        if err != nil {
            return SharedCartResponse{}, common.NewFDKError(err.Error())
        }
         return updateCartWithSharedItemsResponse, nil
        
    }
          
    
    
    
  
    
    
    //CartGetPromotionOffersXQuery holds query params
    type CartGetPromotionOffersXQuery struct { 
        Slug string  `url:"slug,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        PromotionGroup string  `url:"promotion_group,omitempty"`  
    }
    
    // GetPromotionOffers Fetch available promotions
    func (ca *Cart)  GetPromotionOffers(xQuery CartGetPromotionOffersXQuery) (PromotionOffersResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getPromotionOffersResponse PromotionOffersResponse
	    )

        

        
            
                
            
                
            
                
            
        

        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            "/service/application/cart/v1.0/available-promotions",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return PromotionOffersResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getPromotionOffersResponse)
        if err != nil {
            return PromotionOffersResponse{}, common.NewFDKError(err.Error())
        }
         return getPromotionOffersResponse, nil
        
    }
          
    
    
    
  
    
    
    //CartGetLadderOffersXQuery holds query params
    type CartGetLadderOffersXQuery struct { 
        Slug string  `url:"slug,omitempty"` 
        StoreID string  `url:"store_id,omitempty"` 
        PromotionID string  `url:"promotion_id,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetLadderOffers Fetch ladder price promotion
    func (ca *Cart)  GetLadderOffers(xQuery CartGetLadderOffersXQuery) (LadderPriceOffers, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getLadderOffersResponse LadderPriceOffers
	    )

        

        
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            ca.config,
            "get",
            fmt.Sprintf("/service/application/cart/v1.0/available-ladder-prices",),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return LadderPriceOffers{}, err
	    }
        
        err = json.Unmarshal(response, &getLadderOffersResponse)
        if err != nil {
            return LadderPriceOffers{}, common.NewFDKError(err.Error())
        }
         return getLadderOffersResponse, nil
        
    }
          
    

    // Common ...
    type Common struct {
        config *AppConfig
    }
    // NewCommon ...
    func NewCommon(config *AppConfig) *Common {
        return &Common{config}
    }
    
    
    
  
    
    
    //CommonSearchApplicationXQuery holds query params
    type CommonSearchApplicationXQuery struct { 
        Query string  `url:"query,omitempty"`  
    }
    
    // SearchApplication Search Application
    func (co *Common)  SearchApplication(Authorization string, xQuery CommonSearchApplicationXQuery) (ApplicationResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             searchApplicationResponse ApplicationResponse
	    )

        

        
            
                
            
        

        
    
        
        //Adding extra headers
        var xHeaders = make(map[string]string) 
        
         
         xHeaders["authorization"] =  Authorization
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/common/configuration/v1.0/application/search-application",
            xHeaders,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ApplicationResponse{}, err
	    }
        
        err = json.Unmarshal(response, &searchApplicationResponse)
        if err != nil {
            return ApplicationResponse{}, common.NewFDKError(err.Error())
        }
         return searchApplicationResponse, nil
        
    }
          
    
    
    
  
    
    
    //CommonGetLocationsXQuery holds query params
    type CommonGetLocationsXQuery struct { 
        LocationType string  `url:"location_type,omitempty"` 
        ID string  `url:"id,omitempty"`  
    }
    
    // GetLocations Get countries, states, cities
    func (co *Common)  GetLocations(xQuery CommonGetLocationsXQuery) (Locations, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getLocationsResponse Locations
	    )

        

        
            
                
            
                
            
        

        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/common/configuration/v1.0/location",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return Locations{}, err
	    }
        
        err = json.Unmarshal(response, &getLocationsResponse)
        if err != nil {
            return Locations{}, common.NewFDKError(err.Error())
        }
         return getLocationsResponse, nil
        
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
    func (le *Lead)  GetTicket(ID string) (Ticket, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getTicketResponse Ticket
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
            return Ticket{}, err
	    }
        
        err = json.Unmarshal(response, &getTicketResponse)
        if err != nil {
            return Ticket{}, common.NewFDKError(err.Error())
        }
         return getTicketResponse, nil
        
    }
          
    
    
    
  
    
    
    // CreateHistory Create history for specific Ticket
    func (le *Lead)  CreateHistory(ID string, body  TicketHistoryPayload) (TicketHistory, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             createHistoryResponse TicketHistory
	    )

        
            
        
            
                //enum validation inside request body
                err = body.Type.IsValid()
                if err != nil {
                
                    return TicketHistory{}, common.NewFDKError(err.Error())
                }
            
        

        

        
        
        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return TicketHistory{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return TicketHistory{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            le.config,
            "post",
            fmt.Sprintf("/service/application/lead/v1.0/ticket/%s/history",ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return TicketHistory{}, err
	    }
        
        err = json.Unmarshal(response, &createHistoryResponse)
        if err != nil {
            return TicketHistory{}, common.NewFDKError(err.Error())
        }
         return createHistoryResponse, nil
        
    }
          
    
    
    
  
    
    
    // CreateTicket Create Ticket
    func (le *Lead)  CreateTicket(body  AddTicketPayload) (Ticket, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             createTicketResponse Ticket
	    )

        
            
        
            
        
            
                //enum validation inside request body
                err = body.Priority.IsValid()
                if err != nil {
                
                    return Ticket{}, common.NewFDKError(err.Error())
                }
            
        
            
        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return Ticket{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return Ticket{}, common.NewFDKError(err.Error())
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
            return Ticket{}, err
	    }
        
        err = json.Unmarshal(response, &createTicketResponse)
        if err != nil {
            return Ticket{}, common.NewFDKError(err.Error())
        }
         return createTicketResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetCustomForm Get specific Custom Form using it's slug
    func (le *Lead)  GetCustomForm(Slug string) (CustomForm, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getCustomFormResponse CustomForm
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
            return CustomForm{}, err
	    }
        
        err = json.Unmarshal(response, &getCustomFormResponse)
        if err != nil {
            return CustomForm{}, common.NewFDKError(err.Error())
        }
         return getCustomFormResponse, nil
        
    }
          
    
    
    
  
    
    
    // SubmitCustomForm Submit Response for a specific Custom Form using it's slug
    func (le *Lead)  SubmitCustomForm(Slug string, body  CustomFormSubmissionPayload) (SubmitCustomFormResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             submitCustomFormResponse SubmitCustomFormResponse
	    )

        
            
        
            
        

        

        
        
        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return SubmitCustomFormResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return SubmitCustomFormResponse{}, common.NewFDKError(err.Error())
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
            return SubmitCustomFormResponse{}, err
	    }
        
        err = json.Unmarshal(response, &submitCustomFormResponse)
        if err != nil {
            return SubmitCustomFormResponse{}, common.NewFDKError(err.Error())
        }
         return submitCustomFormResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetParticipantsInsideVideoRoom Get participants of a specific Video Room using it's unique name
    func (le *Lead)  GetParticipantsInsideVideoRoom(UniqueName string) (GetParticipantsInsideVideoRoomResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getParticipantsInsideVideoRoomResponse GetParticipantsInsideVideoRoomResponse
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
            return GetParticipantsInsideVideoRoomResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getParticipantsInsideVideoRoomResponse)
        if err != nil {
            return GetParticipantsInsideVideoRoomResponse{}, common.NewFDKError(err.Error())
        }
         return getParticipantsInsideVideoRoomResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetTokenForVideoRoom Get Token to join a specific Video Room using it's unqiue name
    func (le *Lead)  GetTokenForVideoRoom(UniqueName string) (GetTokenForVideoRoomResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getTokenForVideoRoomResponse GetTokenForVideoRoomResponse
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
            return GetTokenForVideoRoomResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getTokenForVideoRoomResponse)
        if err != nil {
            return GetTokenForVideoRoomResponse{}, common.NewFDKError(err.Error())
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
    
    
    
  
    
    
    // GetAllPages Get all pages of a theme
    func (th *Theme)  GetAllPages(ThemeID string) (AllAvailablePageSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getAllPagesResponse AllAvailablePageSchema
	    )

        

        

        
        
        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "get",
            fmt.Sprintf("/service/application/theme/v1.0/%s/page",ThemeID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return AllAvailablePageSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getAllPagesResponse)
        if err != nil {
            return AllAvailablePageSchema{}, common.NewFDKError(err.Error())
        }
         return getAllPagesResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetPage Get page of a theme
    func (th *Theme)  GetPage(ThemeID string, PageValue string) (AvailablePageSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getPageResponse AvailablePageSchema
	    )

        

        

        
        
        
        
        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            th.config,
            "get",
            fmt.Sprintf("/service/application/theme/v1.0/%s/%s",ThemeID,PageValue),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return AvailablePageSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getPageResponse)
        if err != nil {
            return AvailablePageSchema{}, common.NewFDKError(err.Error())
        }
         return getPageResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetAppliedTheme Get the theme currently applied to an application
    func (th *Theme)  GetAppliedTheme() (ThemesSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getAppliedThemeResponse ThemesSchema
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
            return ThemesSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getAppliedThemeResponse)
        if err != nil {
            return ThemesSchema{}, common.NewFDKError(err.Error())
        }
         return getAppliedThemeResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetThemeForPreview Get a theme for a preview
    func (th *Theme)  GetThemeForPreview(ThemeID string) (ThemesSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getThemeForPreviewResponse ThemesSchema
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
            return ThemesSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getThemeForPreviewResponse)
        if err != nil {
            return ThemesSchema{}, common.NewFDKError(err.Error())
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
    
    
    
  
    
    
    //UserLoginWithFacebookXQuery holds query params
    type UserLoginWithFacebookXQuery struct { 
        Platform string  `url:"platform,omitempty"`  
    }
    
    // LoginWithFacebook Login or Register using Facebook
    func (us *User)  LoginWithFacebook(xQuery UserLoginWithFacebookXQuery, body  OAuthRequestSchema) (AuthSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             loginWithFacebookResponse AuthSuccess
	    )

        
            
        
            
        
            
        

        
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return AuthSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return AuthSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/login/facebook-token",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return AuthSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &loginWithFacebookResponse)
        if err != nil {
            return AuthSuccess{}, common.NewFDKError(err.Error())
        }
         return loginWithFacebookResponse, nil
        
    }
          
    
    
    
  
    
    
    //UserLoginWithGoogleXQuery holds query params
    type UserLoginWithGoogleXQuery struct { 
        Platform string  `url:"platform,omitempty"`  
    }
    
    // LoginWithGoogle Login or Register using Google
    func (us *User)  LoginWithGoogle(xQuery UserLoginWithGoogleXQuery, body  OAuthRequestSchema) (AuthSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             loginWithGoogleResponse AuthSuccess
	    )

        
            
        
            
        
            
        

        
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return AuthSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return AuthSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/login/google-token",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return AuthSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &loginWithGoogleResponse)
        if err != nil {
            return AuthSuccess{}, common.NewFDKError(err.Error())
        }
         return loginWithGoogleResponse, nil
        
    }
          
    
    
    
  
    
    
    //UserLoginWithGoogleAndroidXQuery holds query params
    type UserLoginWithGoogleAndroidXQuery struct { 
        Platform string  `url:"platform,omitempty"`  
    }
    
    // LoginWithGoogleAndroid Login or Register using Google on Android
    func (us *User)  LoginWithGoogleAndroid(xQuery UserLoginWithGoogleAndroidXQuery, body  OAuthRequestSchema) (AuthSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             loginWithGoogleAndroidResponse AuthSuccess
	    )

        
            
        
            
        
            
        

        
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return AuthSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return AuthSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/login/google-android",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return AuthSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &loginWithGoogleAndroidResponse)
        if err != nil {
            return AuthSuccess{}, common.NewFDKError(err.Error())
        }
         return loginWithGoogleAndroidResponse, nil
        
    }
          
    
    
    
  
    
    
    //UserLoginWithGoogleIOSXQuery holds query params
    type UserLoginWithGoogleIOSXQuery struct { 
        Platform string  `url:"platform,omitempty"`  
    }
    
    // LoginWithGoogleIOS Login or Register using Google on iOS
    func (us *User)  LoginWithGoogleIOS(xQuery UserLoginWithGoogleIOSXQuery, body  OAuthRequestSchema) (AuthSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             loginWithGoogleIOSResponse AuthSuccess
	    )

        
            
        
            
        
            
        

        
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return AuthSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return AuthSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/login/google-ios",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return AuthSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &loginWithGoogleIOSResponse)
        if err != nil {
            return AuthSuccess{}, common.NewFDKError(err.Error())
        }
         return loginWithGoogleIOSResponse, nil
        
    }
          
    
    
    
  
    
    
    //UserLoginWithAppleIOSXQuery holds query params
    type UserLoginWithAppleIOSXQuery struct { 
        Platform string  `url:"platform,omitempty"`  
    }
    
    // LoginWithAppleIOS Login or Register using Apple on iOS
    func (us *User)  LoginWithAppleIOS(xQuery UserLoginWithAppleIOSXQuery, body  OAuthRequestAppleSchema) (AuthSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             loginWithAppleIOSResponse AuthSuccess
	    )

        
            
        
            
        
            
        

        
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return AuthSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return AuthSuccess{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            us.config,
            "post",
            "/service/application/user/authentication/v1.0/login/apple-ios",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return AuthSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &loginWithAppleIOSResponse)
        if err != nil {
            return AuthSuccess{}, common.NewFDKError(err.Error())
        }
         return loginWithAppleIOSResponse, nil
        
    }
          
    
    
    
  
    
    
    //UserLoginWithOTPXQuery holds query params
    type UserLoginWithOTPXQuery struct { 
        Platform string  `url:"platform,omitempty"`  
    }
    
    // LoginWithOTP Login or Register with OTP
    func (us *User)  LoginWithOTP(xQuery UserLoginWithOTPXQuery, body  SendOtpRequestSchema) (SendOtpResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             loginWithOTPResponse SendOtpResponse
	    )

        
            
        
            
        
            
        

        
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return SendOtpResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return SendOtpResponse{}, common.NewFDKError(err.Error())
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
            return SendOtpResponse{}, err
	    }
        
        err = json.Unmarshal(response, &loginWithOTPResponse)
        if err != nil {
            return SendOtpResponse{}, common.NewFDKError(err.Error())
        }
         return loginWithOTPResponse, nil
        
    }
          
    
    
    
  
    
    
    // LoginWithEmailAndPassword Login or Register with password
    func (us *User)  LoginWithEmailAndPassword(body  PasswordLoginRequestSchema) (LoginSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             loginWithEmailAndPasswordResponse LoginSuccess
	    )

        
            
        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return LoginSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return LoginSuccess{}, common.NewFDKError(err.Error())
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
            return LoginSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &loginWithEmailAndPasswordResponse)
        if err != nil {
            return LoginSuccess{}, common.NewFDKError(err.Error())
        }
         return loginWithEmailAndPasswordResponse, nil
        
    }
          
    
    
    
  
    
    
    //UserSendResetPasswordEmailXQuery holds query params
    type UserSendResetPasswordEmailXQuery struct { 
        Platform string  `url:"platform,omitempty"`  
    }
    
    // SendResetPasswordEmail Reset Password
    func (us *User)  SendResetPasswordEmail(xQuery UserSendResetPasswordEmailXQuery, body  SendResetPasswordEmailRequestSchema) (ResetPasswordSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             sendResetPasswordEmailResponse ResetPasswordSuccess
	    )

        
            
        
            
        

        
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return ResetPasswordSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return ResetPasswordSuccess{}, common.NewFDKError(err.Error())
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
            return ResetPasswordSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &sendResetPasswordEmailResponse)
        if err != nil {
            return ResetPasswordSuccess{}, common.NewFDKError(err.Error())
        }
         return sendResetPasswordEmailResponse, nil
        
    }
          
    
    
    
  
    
    
    // ForgotPassword Forgot Password
    func (us *User)  ForgotPassword(body  ForgotPasswordRequestSchema) (LoginSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             forgotPasswordResponse LoginSuccess
	    )

        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return LoginSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return LoginSuccess{}, common.NewFDKError(err.Error())
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
            return LoginSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &forgotPasswordResponse)
        if err != nil {
            return LoginSuccess{}, common.NewFDKError(err.Error())
        }
         return forgotPasswordResponse, nil
        
    }
          
    
    
    
  
    
    
    // SendResetToken Reset Password using token
    func (us *User)  SendResetToken(body  CodeRequestBodySchema) (ResetPasswordSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             sendResetTokenResponse ResetPasswordSuccess
	    )

        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return ResetPasswordSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return ResetPasswordSuccess{}, common.NewFDKError(err.Error())
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
            return ResetPasswordSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &sendResetTokenResponse)
        if err != nil {
            return ResetPasswordSuccess{}, common.NewFDKError(err.Error())
        }
         return sendResetTokenResponse, nil
        
    }
          
    
    
    
  
    
    
    // LoginWithToken Login or Register with token
    func (us *User)  LoginWithToken(body  TokenRequestBodySchema) (LoginSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             loginWithTokenResponse LoginSuccess
	    )

        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return LoginSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return LoginSuccess{}, common.NewFDKError(err.Error())
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
            return LoginSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &loginWithTokenResponse)
        if err != nil {
            return LoginSuccess{}, common.NewFDKError(err.Error())
        }
         return loginWithTokenResponse, nil
        
    }
          
    
    
    
  
    
    
    //UserRegisterWithFormXQuery holds query params
    type UserRegisterWithFormXQuery struct { 
        Platform string  `url:"platform,omitempty"`  
    }
    
    // RegisterWithForm Registration using a form
    func (us *User)  RegisterWithForm(xQuery UserRegisterWithFormXQuery, body  FormRegisterRequestSchema) (RegisterFormSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             registerWithFormResponse RegisterFormSuccess
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

        
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return RegisterFormSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return RegisterFormSuccess{}, common.NewFDKError(err.Error())
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
            return RegisterFormSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &registerWithFormResponse)
        if err != nil {
            return RegisterFormSuccess{}, common.NewFDKError(err.Error())
        }
         return registerWithFormResponse, nil
        
    }
          
    
    
    
  
    
    
    // VerifyEmail Verify email
    func (us *User)  VerifyEmail(body  CodeRequestBodySchema) (VerifyEmailSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             verifyEmailResponse VerifyEmailSuccess
	    )

        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return VerifyEmailSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return VerifyEmailSuccess{}, common.NewFDKError(err.Error())
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
            return VerifyEmailSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &verifyEmailResponse)
        if err != nil {
            return VerifyEmailSuccess{}, common.NewFDKError(err.Error())
        }
         return verifyEmailResponse, nil
        
    }
          
    
    
    
  
    
    
    // VerifyMobile Verify mobile
    func (us *User)  VerifyMobile(body  CodeRequestBodySchema) (VerifyEmailSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             verifyMobileResponse VerifyEmailSuccess
	    )

        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return VerifyEmailSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return VerifyEmailSuccess{}, common.NewFDKError(err.Error())
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
            return VerifyEmailSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &verifyMobileResponse)
        if err != nil {
            return VerifyEmailSuccess{}, common.NewFDKError(err.Error())
        }
         return verifyMobileResponse, nil
        
    }
          
    
    
    
  
    
    
    // HasPassword Check password
    func (us *User)  HasPassword() (HasPasswordSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             hasPasswordResponse HasPasswordSuccess
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
            return HasPasswordSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &hasPasswordResponse)
        if err != nil {
            return HasPasswordSuccess{}, common.NewFDKError(err.Error())
        }
         return hasPasswordResponse, nil
        
    }
          
    
    
    
  
    
    
    // UpdatePassword Update user password
    func (us *User)  UpdatePassword(body  UpdatePasswordRequestSchema) (VerifyEmailSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             updatePasswordResponse VerifyEmailSuccess
	    )

        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return VerifyEmailSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return VerifyEmailSuccess{}, common.NewFDKError(err.Error())
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
            return VerifyEmailSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &updatePasswordResponse)
        if err != nil {
            return VerifyEmailSuccess{}, common.NewFDKError(err.Error())
        }
         return updatePasswordResponse, nil
        
    }
          
    
    
    
  
    
    
    // Logout Logs out currently logged in user
    func (us *User)  Logout() (LogoutSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             logoutResponse LogoutSuccess
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
            return LogoutSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &logoutResponse)
        if err != nil {
            return LogoutSuccess{}, common.NewFDKError(err.Error())
        }
         return logoutResponse, nil
        
    }
          
    
    
    
  
    
    
    //UserSendOTPOnMobileXQuery holds query params
    type UserSendOTPOnMobileXQuery struct { 
        Platform string  `url:"platform,omitempty"`  
    }
    
    // SendOTPOnMobile Send OTP on mobile
    func (us *User)  SendOTPOnMobile(xQuery UserSendOTPOnMobileXQuery, body  SendMobileOtpRequestSchema) (OtpSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             sendOTPOnMobileResponse OtpSuccess
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

        
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return OtpSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return OtpSuccess{}, common.NewFDKError(err.Error())
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
            return OtpSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &sendOTPOnMobileResponse)
        if err != nil {
            return OtpSuccess{}, common.NewFDKError(err.Error())
        }
         return sendOTPOnMobileResponse, nil
        
    }
          
    
    
    
  
    
    
    //UserVerifyMobileOTPXQuery holds query params
    type UserVerifyMobileOTPXQuery struct { 
        Platform string  `url:"platform,omitempty"`  
    }
    
    // VerifyMobileOTP Verify OTP on mobile
    func (us *User)  VerifyMobileOTP(xQuery UserVerifyMobileOTPXQuery, body  VerifyOtpRequestSchema) (VerifyOtpSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             verifyMobileOTPResponse VerifyOtpSuccess
	    )

        
            
        
            
        
            
        

        
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return VerifyOtpSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return VerifyOtpSuccess{}, common.NewFDKError(err.Error())
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
            return VerifyOtpSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &verifyMobileOTPResponse)
        if err != nil {
            return VerifyOtpSuccess{}, common.NewFDKError(err.Error())
        }
         return verifyMobileOTPResponse, nil
        
    }
          
    
    
    
  
    
    
    //UserSendOTPOnEmailXQuery holds query params
    type UserSendOTPOnEmailXQuery struct { 
        Platform string  `url:"platform,omitempty"`  
    }
    
    // SendOTPOnEmail Send OTP on email
    func (us *User)  SendOTPOnEmail(xQuery UserSendOTPOnEmailXQuery, body  SendEmailOtpRequestSchema) (EmailOtpSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             sendOTPOnEmailResponse EmailOtpSuccess
	    )

        
            
        
            
        
            
        
            
        

        
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return EmailOtpSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return EmailOtpSuccess{}, common.NewFDKError(err.Error())
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
            return EmailOtpSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &sendOTPOnEmailResponse)
        if err != nil {
            return EmailOtpSuccess{}, common.NewFDKError(err.Error())
        }
         return sendOTPOnEmailResponse, nil
        
    }
          
    
    
    
  
    
    
    //UserVerifyEmailOTPXQuery holds query params
    type UserVerifyEmailOTPXQuery struct { 
        Platform string  `url:"platform,omitempty"`  
    }
    
    // VerifyEmailOTP Verify OTP on email
    func (us *User)  VerifyEmailOTP(xQuery UserVerifyEmailOTPXQuery, body  VerifyEmailOtpRequestSchema) (VerifyOtpSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             verifyEmailOTPResponse VerifyOtpSuccess
	    )

        
            
        
            
        
            
        
            
        

        
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return VerifyOtpSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return VerifyOtpSuccess{}, common.NewFDKError(err.Error())
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
            return VerifyOtpSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &verifyEmailOTPResponse)
        if err != nil {
            return VerifyOtpSuccess{}, common.NewFDKError(err.Error())
        }
         return verifyEmailOTPResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetLoggedInUser Get logged in user
    func (us *User)  GetLoggedInUser() (UserObjectSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getLoggedInUserResponse UserObjectSchema
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
            return UserObjectSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getLoggedInUserResponse)
        if err != nil {
            return UserObjectSchema{}, common.NewFDKError(err.Error())
        }
         return getLoggedInUserResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetListOfActiveSessions Get list of sessions
    func (us *User)  GetListOfActiveSessions() (SessionListSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getListOfActiveSessionsResponse SessionListSuccess
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
            return SessionListSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &getListOfActiveSessionsResponse)
        if err != nil {
            return SessionListSuccess{}, common.NewFDKError(err.Error())
        }
         return getListOfActiveSessionsResponse, nil
        
    }
          
    
    
    
  
    
    
    //UserGetPlatformConfigXQuery holds query params
    type UserGetPlatformConfigXQuery struct { 
        Name string  `url:"name,omitempty"`  
    }
    
    // GetPlatformConfig Get platform configurations
    func (us *User)  GetPlatformConfig(xQuery UserGetPlatformConfigXQuery) (PlatformSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getPlatformConfigResponse PlatformSchema
	    )

        

        
            
                
            
        

        
    
         
        
        
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
            return PlatformSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getPlatformConfigResponse)
        if err != nil {
            return PlatformSchema{}, common.NewFDKError(err.Error())
        }
         return getPlatformConfigResponse, nil
        
    }
          
    
    
    
  
    
    
    //UserUpdateProfileXQuery holds query params
    type UserUpdateProfileXQuery struct { 
        Platform string  `url:"platform,omitempty"`  
    }
    
    // UpdateProfile Edit Profile Details
    func (us *User)  UpdateProfile(xQuery UserUpdateProfileXQuery, body  EditProfileRequestSchema) (ProfileEditSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             updateProfileResponse ProfileEditSuccess
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

        
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return ProfileEditSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return ProfileEditSuccess{}, common.NewFDKError(err.Error())
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
            return ProfileEditSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &updateProfileResponse)
        if err != nil {
            return ProfileEditSuccess{}, common.NewFDKError(err.Error())
        }
         return updateProfileResponse, nil
        
    }
          
    
    
    
  
    
    
    //UserAddMobileNumberXQuery holds query params
    type UserAddMobileNumberXQuery struct { 
        Platform string  `url:"platform,omitempty"`  
    }
    
    // AddMobileNumber Add mobile number to profile
    func (us *User)  AddMobileNumber(xQuery UserAddMobileNumberXQuery, body  EditMobileRequestSchema) (VerifyMobileOTPSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             addMobileNumberResponse VerifyMobileOTPSuccess
	    )

        
            
        
            
        

        
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return VerifyMobileOTPSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return VerifyMobileOTPSuccess{}, common.NewFDKError(err.Error())
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
            return VerifyMobileOTPSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &addMobileNumberResponse)
        if err != nil {
            return VerifyMobileOTPSuccess{}, common.NewFDKError(err.Error())
        }
         return addMobileNumberResponse, nil
        
    }
          
    
    
    
  
    
    
    //UserDeleteMobileNumberXQuery holds query params
    type UserDeleteMobileNumberXQuery struct { 
        Platform string  `url:"platform,omitempty"` 
        Active bool  `url:"active,omitempty"` 
        Primary bool  `url:"primary,omitempty"` 
        Verified bool  `url:"verified,omitempty"` 
        CountryCode string  `url:"country_code,omitempty"` 
        Phone string  `url:"phone,omitempty"`  
    }
    
    // DeleteMobileNumber Delete mobile number from profile
    func (us *User)  DeleteMobileNumber(xQuery UserDeleteMobileNumberXQuery) (LoginSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             deleteMobileNumberResponse LoginSuccess
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return LoginSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &deleteMobileNumberResponse)
        if err != nil {
            return LoginSuccess{}, common.NewFDKError(err.Error())
        }
         return deleteMobileNumberResponse, nil
        
    }
          
    
    
    
  
    
    
    // SetMobileNumberAsPrimary Set mobile as primary
    func (us *User)  SetMobileNumberAsPrimary(body  SendVerificationLinkMobileRequestSchema) (LoginSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             setMobileNumberAsPrimaryResponse LoginSuccess
	    )

        
            
        
            
        
            
        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return LoginSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return LoginSuccess{}, common.NewFDKError(err.Error())
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
            return LoginSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &setMobileNumberAsPrimaryResponse)
        if err != nil {
            return LoginSuccess{}, common.NewFDKError(err.Error())
        }
         return setMobileNumberAsPrimaryResponse, nil
        
    }
          
    
    
    
  
    
    
    //UserSendVerificationLinkToMobileXQuery holds query params
    type UserSendVerificationLinkToMobileXQuery struct { 
        Platform string  `url:"platform,omitempty"`  
    }
    
    // SendVerificationLinkToMobile Send verification link to mobile
    func (us *User)  SendVerificationLinkToMobile(xQuery UserSendVerificationLinkToMobileXQuery, body  SendVerificationLinkMobileRequestSchema) (SendMobileVerifyLinkSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             sendVerificationLinkToMobileResponse SendMobileVerifyLinkSuccess
	    )

        
            
        
            
        
            
        
            
        
            
        

        
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return SendMobileVerifyLinkSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return SendMobileVerifyLinkSuccess{}, common.NewFDKError(err.Error())
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
            return SendMobileVerifyLinkSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &sendVerificationLinkToMobileResponse)
        if err != nil {
            return SendMobileVerifyLinkSuccess{}, common.NewFDKError(err.Error())
        }
         return sendVerificationLinkToMobileResponse, nil
        
    }
          
    
    
    
  
    
    
    //UserAddEmailXQuery holds query params
    type UserAddEmailXQuery struct { 
        Platform string  `url:"platform,omitempty"`  
    }
    
    // AddEmail Add email to profile
    func (us *User)  AddEmail(xQuery UserAddEmailXQuery, body  EditEmailRequestSchema) (VerifyEmailOTPSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             addEmailResponse VerifyEmailOTPSuccess
	    )

        
            
        

        
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return VerifyEmailOTPSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return VerifyEmailOTPSuccess{}, common.NewFDKError(err.Error())
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
            return VerifyEmailOTPSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &addEmailResponse)
        if err != nil {
            return VerifyEmailOTPSuccess{}, common.NewFDKError(err.Error())
        }
         return addEmailResponse, nil
        
    }
          
    
    
    
  
    
    
    //UserDeleteEmailXQuery holds query params
    type UserDeleteEmailXQuery struct { 
        Platform string  `url:"platform,omitempty"` 
        Active bool  `url:"active,omitempty"` 
        Primary bool  `url:"primary,omitempty"` 
        Verified bool  `url:"verified,omitempty"` 
        Email string  `url:"email,omitempty"`  
    }
    
    // DeleteEmail Delete email from profile
    func (us *User)  DeleteEmail(xQuery UserDeleteEmailXQuery) (LoginSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             deleteEmailResponse LoginSuccess
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return LoginSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &deleteEmailResponse)
        if err != nil {
            return LoginSuccess{}, common.NewFDKError(err.Error())
        }
         return deleteEmailResponse, nil
        
    }
          
    
    
    
  
    
    
    // SetEmailAsPrimary Set email as primary
    func (us *User)  SetEmailAsPrimary(body  EditEmailRequestSchema) (LoginSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             setEmailAsPrimaryResponse LoginSuccess
	    )

        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return LoginSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return LoginSuccess{}, common.NewFDKError(err.Error())
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
            return LoginSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &setEmailAsPrimaryResponse)
        if err != nil {
            return LoginSuccess{}, common.NewFDKError(err.Error())
        }
         return setEmailAsPrimaryResponse, nil
        
    }
          
    
    
    
  
    
    
    //UserSendVerificationLinkToEmailXQuery holds query params
    type UserSendVerificationLinkToEmailXQuery struct { 
        Platform string  `url:"platform,omitempty"`  
    }
    
    // SendVerificationLinkToEmail Send verification link to email
    func (us *User)  SendVerificationLinkToEmail(xQuery UserSendVerificationLinkToEmailXQuery, body  EditEmailRequestSchema) (SendEmailVerifyLinkSuccess, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             sendVerificationLinkToEmailResponse SendEmailVerifyLinkSuccess
	    )

        
            
        

        
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return SendEmailVerifyLinkSuccess{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return SendEmailVerifyLinkSuccess{}, common.NewFDKError(err.Error())
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
            return SendEmailVerifyLinkSuccess{}, err
	    }
        
        err = json.Unmarshal(response, &sendVerificationLinkToEmailResponse)
        if err != nil {
            return SendEmailVerifyLinkSuccess{}, common.NewFDKError(err.Error())
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
    func (co *Content)  GetAnnouncements() (AnnouncementsResponseSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getAnnouncementsResponse AnnouncementsResponseSchema
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
            return AnnouncementsResponseSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getAnnouncementsResponse)
        if err != nil {
            return AnnouncementsResponseSchema{}, common.NewFDKError(err.Error())
        }
         return getAnnouncementsResponse, nil
        
    }
          
    
    
    
  
    
    
    //ContentGetBlogXQuery holds query params
    type ContentGetBlogXQuery struct { 
        RootID string  `url:"root_id,omitempty"`  
    }
    
    // GetBlog Get a blog
    func (co *Content)  GetBlog(Slug string, xQuery ContentGetBlogXQuery) (BlogSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getBlogResponse BlogSchema
	    )

        

        
            
                
            
        

        
        
        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/application/content/v1.0/blogs/%s",Slug),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return BlogSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getBlogResponse)
        if err != nil {
            return BlogSchema{}, common.NewFDKError(err.Error())
        }
         return getBlogResponse, nil
        
    }
          
    
    
    
  
    
    
    //ContentGetBlogsXQuery holds query params
    type ContentGetBlogsXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetBlogs Get a list of blogs
    func (co *Content)  GetBlogs(xQuery ContentGetBlogsXQuery) (BlogGetResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getBlogsResponse BlogGetResponse
	    )

        

        
            
                
            
                
            
        

        
    
         
        
        
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
            return BlogGetResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getBlogsResponse)
        if err != nil {
            return BlogGetResponse{}, common.NewFDKError(err.Error())
        }
         return getBlogsResponse, nil
        
    }
          
            
            
            
            
                
                    
                    
                    
                    
                
                    
                    
                    
                        
                    
                    
                
            
            // GetBlogsPaginator Get a list of blogs  
            func (co *Content)  GetBlogsPaginator( xQuery ContentGetBlogsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                 
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetBlogs(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
       
    
    
    
  
    
    
    // GetDataLoaders Get the data loaders associated with an application
    func (co *Content)  GetDataLoaders() (DataLoadersSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getDataLoadersResponse DataLoadersSchema
	    )

        

        

        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/content/v1.0/data-loader",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return DataLoadersSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getDataLoadersResponse)
        if err != nil {
            return DataLoadersSchema{}, common.NewFDKError(err.Error())
        }
         return getDataLoadersResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetFaqs Get a list of FAQs
    func (co *Content)  GetFaqs() (FaqResponseSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getFaqsResponse FaqResponseSchema
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
            return FaqResponseSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getFaqsResponse)
        if err != nil {
            return FaqResponseSchema{}, common.NewFDKError(err.Error())
        }
         return getFaqsResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetFaqCategories Get a list of FAQ categories
    func (co *Content)  GetFaqCategories() (GetFaqCategoriesSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getFaqCategoriesResponse GetFaqCategoriesSchema
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
            return GetFaqCategoriesSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getFaqCategoriesResponse)
        if err != nil {
            return GetFaqCategoriesSchema{}, common.NewFDKError(err.Error())
        }
         return getFaqCategoriesResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetFaqBySlug Get an FAQ
    func (co *Content)  GetFaqBySlug(Slug string) (FaqSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getFaqBySlugResponse FaqSchema
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
            return FaqSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getFaqBySlugResponse)
        if err != nil {
            return FaqSchema{}, common.NewFDKError(err.Error())
        }
         return getFaqBySlugResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetFaqCategoryBySlug Get the FAQ category
    func (co *Content)  GetFaqCategoryBySlug(Slug string) (GetFaqCategoryBySlugSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getFaqCategoryBySlugResponse GetFaqCategoryBySlugSchema
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
            return GetFaqCategoryBySlugSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getFaqCategoryBySlugResponse)
        if err != nil {
            return GetFaqCategoryBySlugSchema{}, common.NewFDKError(err.Error())
        }
         return getFaqCategoryBySlugResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetFaqsByCategorySlug Get FAQs using the slug of FAQ category
    func (co *Content)  GetFaqsByCategorySlug(Slug string) (GetFaqSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getFaqsByCategorySlugResponse GetFaqSchema
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
            return GetFaqSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getFaqsByCategorySlugResponse)
        if err != nil {
            return GetFaqSchema{}, common.NewFDKError(err.Error())
        }
         return getFaqsByCategorySlugResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetLandingPage Get the landing page
    func (co *Content)  GetLandingPage() (LandingPageSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getLandingPageResponse LandingPageSchema
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
            return LandingPageSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getLandingPageResponse)
        if err != nil {
            return LandingPageSchema{}, common.NewFDKError(err.Error())
        }
         return getLandingPageResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetLegalInformation Get legal information
    func (co *Content)  GetLegalInformation() (ApplicationLegal, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getLegalInformationResponse ApplicationLegal
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
            return ApplicationLegal{}, err
	    }
        
        err = json.Unmarshal(response, &getLegalInformationResponse)
        if err != nil {
            return ApplicationLegal{}, common.NewFDKError(err.Error())
        }
         return getLegalInformationResponse, nil
        
    }
          
    
    
    
  
    
    
    //ContentGetNavigationsXQuery holds query params
    type ContentGetNavigationsXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetNavigations Get the navigation
    func (co *Content)  GetNavigations(xQuery ContentGetNavigationsXQuery) (NavigationGetResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getNavigationsResponse NavigationGetResponse
	    )

        

        
            
                
            
                
            
        

        
    
         
        
        
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
            return NavigationGetResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getNavigationsResponse)
        if err != nil {
            return NavigationGetResponse{}, common.NewFDKError(err.Error())
        }
         return getNavigationsResponse, nil
        
    }
          
            
            
            
            
                
                    
                    
                    
                    
                
                    
                    
                    
                        
                    
                    
                
            
            // GetNavigationsPaginator Get the navigation  
            func (co *Content)  GetNavigationsPaginator( xQuery ContentGetNavigationsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                 
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetNavigations(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
       
    
    
    
  
    
    
    // GetSEOConfiguration Get the SEO of an application
    func (co *Content)  GetSEOConfiguration() (SeoComponent, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getSEOConfigurationResponse SeoComponent
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
            return SeoComponent{}, err
	    }
        
        err = json.Unmarshal(response, &getSEOConfigurationResponse)
        if err != nil {
            return SeoComponent{}, common.NewFDKError(err.Error())
        }
         return getSEOConfigurationResponse, nil
        
    }
          
    
    
    
  
    
    
    //ContentGetSlideshowsXQuery holds query params
    type ContentGetSlideshowsXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetSlideshows Get the slideshows
    func (co *Content)  GetSlideshows(xQuery ContentGetSlideshowsXQuery) (SlideshowGetResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getSlideshowsResponse SlideshowGetResponse
	    )

        

        
            
                
            
                
            
        

        
    
         
        
        
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
            return SlideshowGetResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getSlideshowsResponse)
        if err != nil {
            return SlideshowGetResponse{}, common.NewFDKError(err.Error())
        }
         return getSlideshowsResponse, nil
        
    }
          
            
            
            
            
                
                    
                    
                    
                    
                
                    
                    
                    
                        
                    
                    
                
            
            // GetSlideshowsPaginator Get the slideshows  
            func (co *Content)  GetSlideshowsPaginator( xQuery ContentGetSlideshowsXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                 
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetSlideshows(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
       
    
    
    
  
    
    
    // GetSlideshow Get a slideshow
    func (co *Content)  GetSlideshow(Slug string) (SlideshowSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getSlideshowResponse SlideshowSchema
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
            return SlideshowSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getSlideshowResponse)
        if err != nil {
            return SlideshowSchema{}, common.NewFDKError(err.Error())
        }
         return getSlideshowResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetSupportInformation Get the support information
    func (co *Content)  GetSupportInformation() (Support, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getSupportInformationResponse Support
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
            return Support{}, err
	    }
        
        err = json.Unmarshal(response, &getSupportInformationResponse)
        if err != nil {
            return Support{}, common.NewFDKError(err.Error())
        }
         return getSupportInformationResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetTags Get the tags associated with an application
    func (co *Content)  GetTags() (TagsSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getTagsResponse TagsSchema
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
            return TagsSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getTagsResponse)
        if err != nil {
            return TagsSchema{}, common.NewFDKError(err.Error())
        }
         return getTagsResponse, nil
        
    }
          
    
    
    
  
    
    
    //ContentGetPageXQuery holds query params
    type ContentGetPageXQuery struct { 
        RootID string  `url:"root_id,omitempty"`  
    }
    
    // GetPage Get a page
    func (co *Content)  GetPage(Slug string, xQuery ContentGetPageXQuery) (PageSchema, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getPageResponse PageSchema
	    )

        

        
            
                
            
        

        
        
        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/application/content/v2.0/pages/%s",Slug),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return PageSchema{}, err
	    }
        
        err = json.Unmarshal(response, &getPageResponse)
        if err != nil {
            return PageSchema{}, common.NewFDKError(err.Error())
        }
         return getPageResponse, nil
        
    }
          
    
    
    
  
    
    
    //ContentGetPagesXQuery holds query params
    type ContentGetPagesXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetPages Get all pages
    func (co *Content)  GetPages(xQuery ContentGetPagesXQuery) (PageGetResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getPagesResponse PageGetResponse
	    )

        

        
            
                
            
                
            
        

        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/content/v2.0/pages/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return PageGetResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getPagesResponse)
        if err != nil {
            return PageGetResponse{}, common.NewFDKError(err.Error())
        }
         return getPagesResponse, nil
        
    }
          
            
            
            
            
                
                    
                    
                    
                    
                
                    
                    
                    
                        
                    
                    
                
            
            // GetPagesPaginator Get all pages  
            func (co *Content)  GetPagesPaginator( xQuery ContentGetPagesXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                 
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetPages(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
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
    func (co *Communication)  GetCommunicationConsent() (CommunicationConsent, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getCommunicationConsentResponse CommunicationConsent
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
            return CommunicationConsent{}, err
	    }
        
        err = json.Unmarshal(response, &getCommunicationConsentResponse)
        if err != nil {
            return CommunicationConsent{}, common.NewFDKError(err.Error())
        }
         return getCommunicationConsentResponse, nil
        
    }
          
    
    
    
  
    
    
    // UpsertCommunicationConsent Upsert communication consent
    func (co *Communication)  UpsertCommunicationConsent(body  CommunicationConsentReq) (CommunicationConsentRes, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             upsertCommunicationConsentResponse CommunicationConsentRes
	    )

        
            
        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return CommunicationConsentRes{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return CommunicationConsentRes{}, common.NewFDKError(err.Error())
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
            return CommunicationConsentRes{}, err
	    }
        
        err = json.Unmarshal(response, &upsertCommunicationConsentResponse)
        if err != nil {
            return CommunicationConsentRes{}, common.NewFDKError(err.Error())
        }
         return upsertCommunicationConsentResponse, nil
        
    }
          
    
    
    
  
    
    
    // UpsertAppPushtoken Upsert push token of a user
    func (co *Communication)  UpsertAppPushtoken(body  PushtokenReq) (PushtokenRes, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             upsertAppPushtokenResponse PushtokenRes
	    )

        
            
        
            
        
            
        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return PushtokenRes{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return PushtokenRes{}, common.NewFDKError(err.Error())
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
            return PushtokenRes{}, err
	    }
        
        err = json.Unmarshal(response, &upsertAppPushtokenResponse)
        if err != nil {
            return PushtokenRes{}, common.NewFDKError(err.Error())
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
    
    
    
  
    
    
    // GetApplicationQRCode Create QR Code of an app
    func (sh *Share)  GetApplicationQRCode() (QRCodeResp, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getApplicationQRCodeResponse QRCodeResp
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
            return QRCodeResp{}, err
	    }
        
        err = json.Unmarshal(response, &getApplicationQRCodeResponse)
        if err != nil {
            return QRCodeResp{}, common.NewFDKError(err.Error())
        }
         return getApplicationQRCodeResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetProductQRCodeBySlug Create QR Code of a product
    func (sh *Share)  GetProductQRCodeBySlug(Slug string) (QRCodeResp, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getProductQRCodeBySlugResponse QRCodeResp
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
            return QRCodeResp{}, err
	    }
        
        err = json.Unmarshal(response, &getProductQRCodeBySlugResponse)
        if err != nil {
            return QRCodeResp{}, common.NewFDKError(err.Error())
        }
         return getProductQRCodeBySlugResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetCollectionQRCodeBySlug Create QR Code of a collection
    func (sh *Share)  GetCollectionQRCodeBySlug(Slug string) (QRCodeResp, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getCollectionQRCodeBySlugResponse QRCodeResp
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
            return QRCodeResp{}, err
	    }
        
        err = json.Unmarshal(response, &getCollectionQRCodeBySlugResponse)
        if err != nil {
            return QRCodeResp{}, common.NewFDKError(err.Error())
        }
         return getCollectionQRCodeBySlugResponse, nil
        
    }
          
    
    
    
  
    
    
    //ShareGetUrlQRCodeXQuery holds query params
    type ShareGetUrlQRCodeXQuery struct { 
        URL string  `url:"url,omitempty"`  
    }
    
    // GetUrlQRCode Create QR Code of a URL
    func (sh *Share)  GetUrlQRCode(xQuery ShareGetUrlQRCodeXQuery) (QRCodeResp, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getUrlQRCodeResponse QRCodeResp
	    )

        

        
            
                
            
        

        
    
         
        
        
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
            return QRCodeResp{}, err
	    }
        
        err = json.Unmarshal(response, &getUrlQRCodeResponse)
        if err != nil {
            return QRCodeResp{}, common.NewFDKError(err.Error())
        }
         return getUrlQRCodeResponse, nil
        
    }
          
    
    
    
  
    
    
    // CreateShortLink Create a short link
    func (sh *Share)  CreateShortLink(body  ShortLinkReq) (ShortLinkRes, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             createShortLinkResponse ShortLinkRes
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return ShortLinkRes{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return ShortLinkRes{}, common.NewFDKError(err.Error())
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
            return ShortLinkRes{}, err
	    }
        
        err = json.Unmarshal(response, &createShortLinkResponse)
        if err != nil {
            return ShortLinkRes{}, common.NewFDKError(err.Error())
        }
         return createShortLinkResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetShortLinkByHash Get short link by hash
    func (sh *Share)  GetShortLinkByHash(Hash string) (ShortLinkRes, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getShortLinkByHashResponse ShortLinkRes
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
            return ShortLinkRes{}, err
	    }
        
        err = json.Unmarshal(response, &getShortLinkByHashResponse)
        if err != nil {
            return ShortLinkRes{}, common.NewFDKError(err.Error())
        }
         return getShortLinkByHashResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetOriginalShortLinkByHash Get original link by hash
    func (sh *Share)  GetOriginalShortLinkByHash(Hash string) (ShortLinkRes, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getOriginalShortLinkByHashResponse ShortLinkRes
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
            return ShortLinkRes{}, err
	    }
        
        err = json.Unmarshal(response, &getOriginalShortLinkByHashResponse)
        if err != nil {
            return ShortLinkRes{}, common.NewFDKError(err.Error())
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
    
    
    
  
    
    
    // StartUpload Initiates an upload and returns a storage link that is valid for 30 minutes. You can use the storage link to make subsequent upload request with file buffer or blob.
    func (fi *FileStorage)  StartUpload(Namespace string, body  StartRequest) (StartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             startUploadResponse StartResponse
	    )

        
            
        
            
        
            
        
            
        
            
        

        

        
        
        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return StartResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return StartResponse{}, common.NewFDKError(err.Error())
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
            return StartResponse{}, err
	    }
        
        err = json.Unmarshal(response, &startUploadResponse)
        if err != nil {
            return StartResponse{}, common.NewFDKError(err.Error())
        }
         return startUploadResponse, nil
        
    }
          
    
    
    
  
    
    
    // CompleteUpload Completes the upload process. After successfully uploading a file, call this API to finish the upload process.
    func (fi *FileStorage)  CompleteUpload(Namespace string, body  StartResponse) (CompleteResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             completeUploadResponse CompleteResponse
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

        

        
        
        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return CompleteResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return CompleteResponse{}, common.NewFDKError(err.Error())
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
            return CompleteResponse{}, err
	    }
        
        err = json.Unmarshal(response, &completeUploadResponse)
        if err != nil {
            return CompleteResponse{}, common.NewFDKError(err.Error())
        }
         return completeUploadResponse, nil
        
    }
          
    
    
    
  
    
    
    // SignUrls Explain here
    func (fi *FileStorage)  SignUrls(body  SignUrlRequest) (SignUrlResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             signUrlsResponse SignUrlResponse
	    )

        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return SignUrlResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return SignUrlResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            fi.config,
            "post",
            "/service/application/assets/v1.0/sign-urls/",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return SignUrlResponse{}, err
	    }
        
        err = json.Unmarshal(response, &signUrlsResponse)
        if err != nil {
            return SignUrlResponse{}, common.NewFDKError(err.Error())
        }
         return signUrlsResponse, nil
        
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
    func (co *Configuration)  GetApplication() (Application, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getApplicationResponse Application
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
            return Application{}, err
	    }
        
        err = json.Unmarshal(response, &getApplicationResponse)
        if err != nil {
            return Application{}, common.NewFDKError(err.Error())
        }
         return getApplicationResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetOwnerInfo Get application, owner and seller information
    func (co *Configuration)  GetOwnerInfo() (ApplicationAboutResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getOwnerInfoResponse ApplicationAboutResponse
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
            return ApplicationAboutResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getOwnerInfoResponse)
        if err != nil {
            return ApplicationAboutResponse{}, common.NewFDKError(err.Error())
        }
         return getOwnerInfoResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetBasicDetails Get basic application details
    func (co *Configuration)  GetBasicDetails() (ApplicationDetail, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getBasicDetailsResponse ApplicationDetail
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
            return ApplicationDetail{}, err
	    }
        
        err = json.Unmarshal(response, &getBasicDetailsResponse)
        if err != nil {
            return ApplicationDetail{}, common.NewFDKError(err.Error())
        }
         return getBasicDetailsResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetIntegrationTokens Get integration tokens
    func (co *Configuration)  GetIntegrationTokens() (AppTokenResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getIntegrationTokensResponse AppTokenResponse
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
            return AppTokenResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getIntegrationTokensResponse)
        if err != nil {
            return AppTokenResponse{}, common.NewFDKError(err.Error())
        }
         return getIntegrationTokensResponse, nil
        
    }
          
    
    
    
  
    
    
    //ConfigurationGetOrderingStoresXQuery holds query params
    type ConfigurationGetOrderingStoresXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        Q string  `url:"q,omitempty"`  
    }
    
    // GetOrderingStores Get deployment stores
    func (co *Configuration)  GetOrderingStores(xQuery ConfigurationGetOrderingStoresXQuery) (OrderingStores, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getOrderingStoresResponse OrderingStores
	    )

        

        
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return OrderingStores{}, err
	    }
        
        err = json.Unmarshal(response, &getOrderingStoresResponse)
        if err != nil {
            return OrderingStores{}, common.NewFDKError(err.Error())
        }
         return getOrderingStoresResponse, nil
        
    }
          
            
            
            
            
                
                    
                    
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
            
            // GetOrderingStoresPaginator Get deployment stores  
            func (co *Configuration)  GetOrderingStoresPaginator( xQuery ConfigurationGetOrderingStoresXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                 
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetOrderingStores(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
       
    
    
    
  
    
    
    // GetStoreDetailById Get ordering store details
    func (co *Configuration)  GetStoreDetailById(StoreID float64) (OrderingStore, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getStoreDetailByIdResponse OrderingStore
	    )

        

        

        
        
        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            fmt.Sprintf("/service/application/configuration/v1.0/ordering-store/stores/undefined",StoreID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return OrderingStore{}, err
	    }
        
        err = json.Unmarshal(response, &getStoreDetailByIdResponse)
        if err != nil {
            return OrderingStore{}, common.NewFDKError(err.Error())
        }
         return getStoreDetailByIdResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetFeatures Get features of application
    func (co *Configuration)  GetFeatures() (AppFeatureResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getFeaturesResponse AppFeatureResponse
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
            return AppFeatureResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getFeaturesResponse)
        if err != nil {
            return AppFeatureResponse{}, common.NewFDKError(err.Error())
        }
         return getFeaturesResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetContactInfo Get application information
    func (co *Configuration)  GetContactInfo() (ApplicationInformation, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getContactInfoResponse ApplicationInformation
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
            return ApplicationInformation{}, err
	    }
        
        err = json.Unmarshal(response, &getContactInfoResponse)
        if err != nil {
            return ApplicationInformation{}, common.NewFDKError(err.Error())
        }
         return getContactInfoResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetCurrencies Get all currencies list
    func (co *Configuration)  GetCurrencies() (CurrenciesResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getCurrenciesResponse CurrenciesResponse
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
            return CurrenciesResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCurrenciesResponse)
        if err != nil {
            return CurrenciesResponse{}, common.NewFDKError(err.Error())
        }
         return getCurrenciesResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetCurrencyById Get currency by its ID
    func (co *Configuration)  GetCurrencyById(ID string) (Currency, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getCurrencyByIdResponse Currency
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
            return Currency{}, err
	    }
        
        err = json.Unmarshal(response, &getCurrencyByIdResponse)
        if err != nil {
            return Currency{}, common.NewFDKError(err.Error())
        }
         return getCurrencyByIdResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetAppCurrencies Get currencies enabled in the application
    func (co *Configuration)  GetAppCurrencies() (AppCurrencyResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getAppCurrenciesResponse AppCurrencyResponse
	    )

        

        

        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/configuration/v1.0/currency",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return AppCurrencyResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getAppCurrenciesResponse)
        if err != nil {
            return AppCurrencyResponse{}, common.NewFDKError(err.Error())
        }
         return getAppCurrenciesResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetLanguages Get list of languages
    func (co *Configuration)  GetLanguages() (LanguageResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getLanguagesResponse LanguageResponse
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
            return LanguageResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getLanguagesResponse)
        if err != nil {
            return LanguageResponse{}, common.NewFDKError(err.Error())
        }
         return getLanguagesResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetOrderingStoreCookie Get an Ordering Store signed cookie on selection of ordering store.
    func (co *Configuration)  GetOrderingStoreCookie(body  OrderingStoreSelectRequest) (SuccessMessageResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getOrderingStoreCookieResponse SuccessMessageResponse
	    )

        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return SuccessMessageResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return SuccessMessageResponse{}, common.NewFDKError(err.Error())
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
            return SuccessMessageResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getOrderingStoreCookieResponse)
        if err != nil {
            return SuccessMessageResponse{}, common.NewFDKError(err.Error())
        }
         return getOrderingStoreCookieResponse, nil
        
    }
          
    
    
    
  
    
    
    // RemoveOrderingStoreCookie Unset the Ordering Store signed cookie.
    func (co *Configuration)  RemoveOrderingStoreCookie() (SuccessMessageResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             removeOrderingStoreCookieResponse SuccessMessageResponse
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
            return SuccessMessageResponse{}, err
	    }
        
        err = json.Unmarshal(response, &removeOrderingStoreCookieResponse)
        if err != nil {
            return SuccessMessageResponse{}, common.NewFDKError(err.Error())
        }
         return removeOrderingStoreCookieResponse, nil
        
    }
          
    
    
    
  
    
    
    //ConfigurationGetAppStaffListXQuery holds query params
    type ConfigurationGetAppStaffListXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        OrderIncent bool  `url:"order_incent,omitempty"` 
        OrderingStore float64  `url:"ordering_store,omitempty"` 
        User string  `url:"user,omitempty"`  
    }
    
    // GetAppStaffList Get a list of staff.
    func (co *Configuration)  GetAppStaffList(xQuery ConfigurationGetAppStaffListXQuery) (AppStaffListResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getAppStaffListResponse AppStaffListResponse
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            co.config,
            "get",
            "/service/application/configuration/v1.0/staff/list",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return AppStaffListResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getAppStaffListResponse)
        if err != nil {
            return AppStaffListResponse{}, common.NewFDKError(err.Error())
        }
         return getAppStaffListResponse, nil
        
    }
          
            
            
            
            
                
                    
                    
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
                    
                    
                    
                        
                    
                    
                
            
            // GetAppStaffListPaginator Get a list of staff.  
            func (co *Configuration)  GetAppStaffListPaginator( xQuery ConfigurationGetAppStaffListXQuery ) *common.Paginator {
                paginator := common.NewPaginator("number")
                 
                 
                 xQuery.PageNo  = paginator.PageNo
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := co.GetAppStaffList(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
       
    
    
    
  
    
    
    //ConfigurationGetAppStaffsXQuery holds query params
    type ConfigurationGetAppStaffsXQuery struct { 
        OrderIncent bool  `url:"order_incent,omitempty"` 
        OrderingStore float64  `url:"ordering_store,omitempty"` 
        User string  `url:"user,omitempty"`  
    }
    
    // GetAppStaffs Get a list of staff.
    func (co *Configuration)  GetAppStaffs(xQuery ConfigurationGetAppStaffsXQuery) (AppStaffResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getAppStaffsResponse AppStaffResponse
	    )

        

        
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return AppStaffResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getAppStaffsResponse)
        if err != nil {
            return AppStaffResponse{}, common.NewFDKError(err.Error())
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
    
    
    
  
    
    
    //PaymentGetAggregatorsConfigXQuery holds query params
    type PaymentGetAggregatorsConfigXQuery struct { 
        Refresh bool  `url:"refresh,omitempty"`  
    }
    
    // GetAggregatorsConfig Get payment gateway keys
    func (pa *Payment)  GetAggregatorsConfig(XAPIToken string, xQuery PaymentGetAggregatorsConfigXQuery) (AggregatorsConfigDetailResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getAggregatorsConfigResponse AggregatorsConfigDetailResponse
	    )

        

        
            
                
            
        

        
    
        
        //Adding extra headers
        var xHeaders = make(map[string]string) 
        
         
         xHeaders["x-api-token"] =  XAPIToken
         
        
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "get",
            "/service/application/payment/v1.0/config/aggregators/key",
            xHeaders,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return AggregatorsConfigDetailResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getAggregatorsConfigResponse)
        if err != nil {
            return AggregatorsConfigDetailResponse{}, common.NewFDKError(err.Error())
        }
         return getAggregatorsConfigResponse, nil
        
    }
          
    
    
    
  
    
    
    // AttachCardToCustomer Attach a saved card to customer.
    func (pa *Payment)  AttachCardToCustomer(body  AttachCardRequest) (AttachCardsResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             attachCardToCustomerResponse AttachCardsResponse
	    )

        
            
        
            
        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return AttachCardsResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return AttachCardsResponse{}, common.NewFDKError(err.Error())
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
            return AttachCardsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &attachCardToCustomerResponse)
        if err != nil {
            return AttachCardsResponse{}, common.NewFDKError(err.Error())
        }
         return attachCardToCustomerResponse, nil
        
    }
          
    
    
    
  
    
    
    //PaymentGetActiveCardAggregatorXQuery holds query params
    type PaymentGetActiveCardAggregatorXQuery struct { 
        Refresh bool  `url:"refresh,omitempty"`  
    }
    
    // GetActiveCardAggregator Fetch active payment gateway for card payments
    func (pa *Payment)  GetActiveCardAggregator(xQuery PaymentGetActiveCardAggregatorXQuery) (ActiveCardPaymentGatewayResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getActiveCardAggregatorResponse ActiveCardPaymentGatewayResponse
	    )

        

        
            
                
            
        

        
    
         
        
        
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
            return ActiveCardPaymentGatewayResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getActiveCardAggregatorResponse)
        if err != nil {
            return ActiveCardPaymentGatewayResponse{}, common.NewFDKError(err.Error())
        }
         return getActiveCardAggregatorResponse, nil
        
    }
          
    
    
    
  
    
    
    //PaymentGetActiveUserCardsXQuery holds query params
    type PaymentGetActiveUserCardsXQuery struct { 
        ForceRefresh bool  `url:"force_refresh,omitempty"`  
    }
    
    // GetActiveUserCards Fetch the list of cards saved by the user
    func (pa *Payment)  GetActiveUserCards(xQuery PaymentGetActiveUserCardsXQuery) (ListCardsResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getActiveUserCardsResponse ListCardsResponse
	    )

        

        
            
                
            
        

        
    
         
        
        
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
            return ListCardsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getActiveUserCardsResponse)
        if err != nil {
            return ListCardsResponse{}, common.NewFDKError(err.Error())
        }
         return getActiveUserCardsResponse, nil
        
    }
          
    
    
    
  
    
    
    // DeleteUserCard Delete a card
    func (pa *Payment)  DeleteUserCard(body  DeletehCardRequest) (DeleteCardsResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             deleteUserCardResponse DeleteCardsResponse
	    )

        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return DeleteCardsResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return DeleteCardsResponse{}, common.NewFDKError(err.Error())
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
            return DeleteCardsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &deleteUserCardResponse)
        if err != nil {
            return DeleteCardsResponse{}, common.NewFDKError(err.Error())
        }
         return deleteUserCardResponse, nil
        
    }
          
    
    
    
  
    
    
    // VerifyCustomerForPayment Validate customer for payment
    func (pa *Payment)  VerifyCustomerForPayment(body  ValidateCustomerRequest) (ValidateCustomerResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             verifyCustomerForPaymentResponse ValidateCustomerResponse
	    )

        
            
        
            
        
            
        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return ValidateCustomerResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return ValidateCustomerResponse{}, common.NewFDKError(err.Error())
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
            return ValidateCustomerResponse{}, err
	    }
        
        err = json.Unmarshal(response, &verifyCustomerForPaymentResponse)
        if err != nil {
            return ValidateCustomerResponse{}, common.NewFDKError(err.Error())
        }
         return verifyCustomerForPaymentResponse, nil
        
    }
          
    
    
    
  
    
    
    // VerifyAndChargePayment Verify and charge payment
    func (pa *Payment)  VerifyAndChargePayment(body  ChargeCustomerRequest) (ChargeCustomerResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             verifyAndChargePaymentResponse ChargeCustomerResponse
	    )

        
            
        
            
        
            
        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return ChargeCustomerResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return ChargeCustomerResponse{}, common.NewFDKError(err.Error())
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
            return ChargeCustomerResponse{}, err
	    }
        
        err = json.Unmarshal(response, &verifyAndChargePaymentResponse)
        if err != nil {
            return ChargeCustomerResponse{}, common.NewFDKError(err.Error())
        }
         return verifyAndChargePaymentResponse, nil
        
    }
          
    
    
    
  
    
    
    // InitialisePayment Initialize a payment (server-to-server) for UPI and BharatQR
    func (pa *Payment)  InitialisePayment(body  PaymentInitializationRequest) (PaymentInitializationResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             initialisePaymentResponse PaymentInitializationResponse
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return PaymentInitializationResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return PaymentInitializationResponse{}, common.NewFDKError(err.Error())
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
            return PaymentInitializationResponse{}, err
	    }
        
        err = json.Unmarshal(response, &initialisePaymentResponse)
        if err != nil {
            return PaymentInitializationResponse{}, common.NewFDKError(err.Error())
        }
         return initialisePaymentResponse, nil
        
    }
          
    
    
    
  
    
    
    // CheckAndUpdatePaymentStatus Performs continuous polling to check status of payment on the server
    func (pa *Payment)  CheckAndUpdatePaymentStatus(body  PaymentStatusUpdateRequest) (PaymentStatusUpdateResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             checkAndUpdatePaymentStatusResponse PaymentStatusUpdateResponse
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return PaymentStatusUpdateResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return PaymentStatusUpdateResponse{}, common.NewFDKError(err.Error())
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
            return PaymentStatusUpdateResponse{}, err
	    }
        
        err = json.Unmarshal(response, &checkAndUpdatePaymentStatusResponse)
        if err != nil {
            return PaymentStatusUpdateResponse{}, common.NewFDKError(err.Error())
        }
         return checkAndUpdatePaymentStatusResponse, nil
        
    }
          
    
    
    
  
    
    
    //PaymentGetPaymentModeRoutesXQuery holds query params
    type PaymentGetPaymentModeRoutesXQuery struct { 
        Amount float64  `url:"amount,omitempty"` 
        CartID string  `url:"cart_id,omitempty"` 
        Pincode string  `url:"pincode,omitempty"` 
        CheckoutMode string  `url:"checkout_mode,omitempty"` 
        Refresh bool  `url:"refresh,omitempty"` 
        CardReference string  `url:"card_reference,omitempty"` 
        UserDetails string  `url:"user_details,omitempty"`  
    }
    
    // GetPaymentModeRoutes Get applicable payment options
    func (pa *Payment)  GetPaymentModeRoutes(xQuery PaymentGetPaymentModeRoutesXQuery) (PaymentModeRouteResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getPaymentModeRoutesResponse PaymentModeRouteResponse
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return PaymentModeRouteResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getPaymentModeRoutesResponse)
        if err != nil {
            return PaymentModeRouteResponse{}, common.NewFDKError(err.Error())
        }
         return getPaymentModeRoutesResponse, nil
        
    }
          
    
    
    
  
    
    
    //PaymentGetPosPaymentModeRoutesXQuery holds query params
    type PaymentGetPosPaymentModeRoutesXQuery struct { 
        Amount float64  `url:"amount,omitempty"` 
        CartID string  `url:"cart_id,omitempty"` 
        Pincode string  `url:"pincode,omitempty"` 
        CheckoutMode string  `url:"checkout_mode,omitempty"` 
        Refresh bool  `url:"refresh,omitempty"` 
        CardReference string  `url:"card_reference,omitempty"` 
        OrderType string  `url:"order_type,omitempty"` 
        UserDetails string  `url:"user_details,omitempty"`  
    }
    
    // GetPosPaymentModeRoutes Get applicable payment options for Point-of-Sale (POS)
    func (pa *Payment)  GetPosPaymentModeRoutes(xQuery PaymentGetPosPaymentModeRoutesXQuery) (PaymentModeRouteResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getPosPaymentModeRoutesResponse PaymentModeRouteResponse
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return PaymentModeRouteResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getPosPaymentModeRoutesResponse)
        if err != nil {
            return PaymentModeRouteResponse{}, common.NewFDKError(err.Error())
        }
         return getPosPaymentModeRoutesResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetRupifiBannerDetails Get CreditLine Offer
    func (pa *Payment)  GetRupifiBannerDetails() (RupifiBannerResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getRupifiBannerDetailsResponse RupifiBannerResponse
	    )

        

        

        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "get",
            "/service/application/payment/v1.0/rupifi/banner",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return RupifiBannerResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getRupifiBannerDetailsResponse)
        if err != nil {
            return RupifiBannerResponse{}, common.NewFDKError(err.Error())
        }
         return getRupifiBannerDetailsResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetEpaylaterBannerDetails Get Epaylater Enabled
    func (pa *Payment)  GetEpaylaterBannerDetails() (EpaylaterBannerResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getEpaylaterBannerDetailsResponse EpaylaterBannerResponse
	    )

        

        

        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "get",
            "/service/application/payment/v1.0/epaylater/banner",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return EpaylaterBannerResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getEpaylaterBannerDetailsResponse)
        if err != nil {
            return EpaylaterBannerResponse{}, common.NewFDKError(err.Error())
        }
         return getEpaylaterBannerDetailsResponse, nil
        
    }
          
    
    
    
  
    
    
    // ResendOrCancelPayment API to resend and cancel a payment link which was already generated.
    func (pa *Payment)  ResendOrCancelPayment(body  ResendOrCancelPaymentRequest) (ResendOrCancelPaymentResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             resendOrCancelPaymentResponse ResendOrCancelPaymentResponse
	    )

        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return ResendOrCancelPaymentResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return ResendOrCancelPaymentResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "post",
            "/service/application/payment/v1.0/payment/resend_or_cancel",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return ResendOrCancelPaymentResponse{}, err
	    }
        
        err = json.Unmarshal(response, &resendOrCancelPaymentResponse)
        if err != nil {
            return ResendOrCancelPaymentResponse{}, common.NewFDKError(err.Error())
        }
         return resendOrCancelPaymentResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetActiveRefundTransferModes Lists the mode of refund
    func (pa *Payment)  GetActiveRefundTransferModes() (TransferModeResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getActiveRefundTransferModesResponse TransferModeResponse
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
            return TransferModeResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getActiveRefundTransferModesResponse)
        if err != nil {
            return TransferModeResponse{}, common.NewFDKError(err.Error())
        }
         return getActiveRefundTransferModesResponse, nil
        
    }
          
    
    
    
  
    
    
    // EnableOrDisableRefundTransferMode Enable/Disable a mode for transferring a refund
    func (pa *Payment)  EnableOrDisableRefundTransferMode(body  UpdateRefundTransferModeRequest) (UpdateRefundTransferModeResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             enableOrDisableRefundTransferModeResponse UpdateRefundTransferModeResponse
	    )

        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return UpdateRefundTransferModeResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return UpdateRefundTransferModeResponse{}, common.NewFDKError(err.Error())
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
            return UpdateRefundTransferModeResponse{}, err
	    }
        
        err = json.Unmarshal(response, &enableOrDisableRefundTransferModeResponse)
        if err != nil {
            return UpdateRefundTransferModeResponse{}, common.NewFDKError(err.Error())
        }
         return enableOrDisableRefundTransferModeResponse, nil
        
    }
          
    
    
    
  
    
    
    //PaymentGetUserBeneficiariesDetailXQuery holds query params
    type PaymentGetUserBeneficiariesDetailXQuery struct { 
        OrderID string  `url:"order_id,omitempty"`  
    }
    
    // GetUserBeneficiariesDetail Lists the beneficiary of a refund
    func (pa *Payment)  GetUserBeneficiariesDetail(xQuery PaymentGetUserBeneficiariesDetailXQuery) (OrderBeneficiaryResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getUserBeneficiariesDetailResponse OrderBeneficiaryResponse
	    )

        

        
            
                
            
        

        
    
         
        
        
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
            return OrderBeneficiaryResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getUserBeneficiariesDetailResponse)
        if err != nil {
            return OrderBeneficiaryResponse{}, common.NewFDKError(err.Error())
        }
         return getUserBeneficiariesDetailResponse, nil
        
    }
          
    
    
    
  
    
    
    //PaymentVerifyIfscCodeXQuery holds query params
    type PaymentVerifyIfscCodeXQuery struct { 
        IfscCode string  `url:"ifsc_code,omitempty"`  
    }
    
    // VerifyIfscCode Verify IFSC Code
    func (pa *Payment)  VerifyIfscCode(xQuery PaymentVerifyIfscCodeXQuery) (IfscCodeResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             verifyIfscCodeResponse IfscCodeResponse
	    )

        

        
            
                
            
        

        
    
         
        
        
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
            return IfscCodeResponse{}, err
	    }
        
        err = json.Unmarshal(response, &verifyIfscCodeResponse)
        if err != nil {
            return IfscCodeResponse{}, common.NewFDKError(err.Error())
        }
         return verifyIfscCodeResponse, nil
        
    }
          
    
    
    
  
    
    
    //PaymentGetOrderBeneficiariesDetailXQuery holds query params
    type PaymentGetOrderBeneficiariesDetailXQuery struct { 
        OrderID string  `url:"order_id,omitempty"`  
    }
    
    // GetOrderBeneficiariesDetail Lists the beneficiary of a refund
    func (pa *Payment)  GetOrderBeneficiariesDetail(xQuery PaymentGetOrderBeneficiariesDetailXQuery) (OrderBeneficiaryResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getOrderBeneficiariesDetailResponse OrderBeneficiaryResponse
	    )

        

        
            
                
            
        

        
    
         
        
        
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
            return OrderBeneficiaryResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getOrderBeneficiariesDetailResponse)
        if err != nil {
            return OrderBeneficiaryResponse{}, common.NewFDKError(err.Error())
        }
         return getOrderBeneficiariesDetailResponse, nil
        
    }
          
    
    
    
  
    
    
    // VerifyOtpAndAddBeneficiaryForBank Verify the beneficiary details using OTP
    func (pa *Payment)  VerifyOtpAndAddBeneficiaryForBank(body  AddBeneficiaryViaOtpVerificationRequest) (AddBeneficiaryViaOtpVerificationResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             verifyOtpAndAddBeneficiaryForBankResponse AddBeneficiaryViaOtpVerificationResponse
	    )

        
            
        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return AddBeneficiaryViaOtpVerificationResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return AddBeneficiaryViaOtpVerificationResponse{}, common.NewFDKError(err.Error())
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
            return AddBeneficiaryViaOtpVerificationResponse{}, err
	    }
        
        err = json.Unmarshal(response, &verifyOtpAndAddBeneficiaryForBankResponse)
        if err != nil {
            return AddBeneficiaryViaOtpVerificationResponse{}, common.NewFDKError(err.Error())
        }
         return verifyOtpAndAddBeneficiaryForBankResponse, nil
        
    }
          
    
    
    
  
    
    
    // AddBeneficiaryDetails Save bank details for cancelled/returned order
    func (pa *Payment)  AddBeneficiaryDetails(body  AddBeneficiaryDetailsRequest) (RefundAccountResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             addBeneficiaryDetailsResponse RefundAccountResponse
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return RefundAccountResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return RefundAccountResponse{}, common.NewFDKError(err.Error())
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
            return RefundAccountResponse{}, err
	    }
        
        err = json.Unmarshal(response, &addBeneficiaryDetailsResponse)
        if err != nil {
            return RefundAccountResponse{}, common.NewFDKError(err.Error())
        }
         return addBeneficiaryDetailsResponse, nil
        
    }
          
    
    
    
  
    
    
    // AddRefundBankAccountUsingOTP Save bank details for cancelled/returned order
    func (pa *Payment)  AddRefundBankAccountUsingOTP(body  AddBeneficiaryDetailsOTPRequest) (RefundAccountResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             addRefundBankAccountUsingOTPResponse RefundAccountResponse
	    )

        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return RefundAccountResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return RefundAccountResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "post",
            "/service/application/payment/v1.0/refund/account/otp",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return RefundAccountResponse{}, err
	    }
        
        err = json.Unmarshal(response, &addRefundBankAccountUsingOTPResponse)
        if err != nil {
            return RefundAccountResponse{}, common.NewFDKError(err.Error())
        }
         return addRefundBankAccountUsingOTPResponse, nil
        
    }
          
    
    
    
  
    
    
    // VerifyOtpAndAddBeneficiaryForWallet Send OTP on adding a wallet beneficiary
    func (pa *Payment)  VerifyOtpAndAddBeneficiaryForWallet(body  WalletOtpRequest) (WalletOtpResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             verifyOtpAndAddBeneficiaryForWalletResponse WalletOtpResponse
	    )

        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return WalletOtpResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return WalletOtpResponse{}, common.NewFDKError(err.Error())
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
            return WalletOtpResponse{}, err
	    }
        
        err = json.Unmarshal(response, &verifyOtpAndAddBeneficiaryForWalletResponse)
        if err != nil {
            return WalletOtpResponse{}, common.NewFDKError(err.Error())
        }
         return verifyOtpAndAddBeneficiaryForWalletResponse, nil
        
    }
          
    
    
    
  
    
    
    // UpdateDefaultBeneficiary Set a default beneficiary for a refund
    func (pa *Payment)  UpdateDefaultBeneficiary(body  SetDefaultBeneficiaryRequest) (SetDefaultBeneficiaryResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             updateDefaultBeneficiaryResponse SetDefaultBeneficiaryResponse
	    )

        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return SetDefaultBeneficiaryResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return SetDefaultBeneficiaryResponse{}, common.NewFDKError(err.Error())
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
            return SetDefaultBeneficiaryResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateDefaultBeneficiaryResponse)
        if err != nil {
            return SetDefaultBeneficiaryResponse{}, common.NewFDKError(err.Error())
        }
         return updateDefaultBeneficiaryResponse, nil
        
    }
          
    
    
    
  
    
    
    //PaymentCustomerCreditSummaryXQuery holds query params
    type PaymentCustomerCreditSummaryXQuery struct { 
        Aggregator string  `url:"aggregator,omitempty"`  
    }
    
    // CustomerCreditSummary API to fetch the customer credit summary
    func (pa *Payment)  CustomerCreditSummary(xQuery PaymentCustomerCreditSummaryXQuery) (CustomerCreditSummaryResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             customerCreditSummaryResponse CustomerCreditSummaryResponse
	    )

        

        
            
                
            
        

        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "get",
            "/service/application/payment/v1.0/payment/credit-summary/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return CustomerCreditSummaryResponse{}, err
	    }
        
        err = json.Unmarshal(response, &customerCreditSummaryResponse)
        if err != nil {
            return CustomerCreditSummaryResponse{}, common.NewFDKError(err.Error())
        }
         return customerCreditSummaryResponse, nil
        
    }
          
    
    
    
  
    
    
    //PaymentRedirectToAggregatorXQuery holds query params
    type PaymentRedirectToAggregatorXQuery struct { 
        Source string  `url:"source,omitempty"` 
        Aggregator string  `url:"aggregator,omitempty"`  
    }
    
    // RedirectToAggregator API to get the redirect url to redirect the user to aggregator's page
    func (pa *Payment)  RedirectToAggregator(xQuery PaymentRedirectToAggregatorXQuery) (RedirectToAggregatorResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             redirectToAggregatorResponse RedirectToAggregatorResponse
	    )

        

        
            
                
            
                
            
        

        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "get",
            "/service/application/payment/v1.0/payment/redirect-to-aggregator/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return RedirectToAggregatorResponse{}, err
	    }
        
        err = json.Unmarshal(response, &redirectToAggregatorResponse)
        if err != nil {
            return RedirectToAggregatorResponse{}, common.NewFDKError(err.Error())
        }
         return redirectToAggregatorResponse, nil
        
    }
          
    
    
    
  
    
    
    //PaymentCheckCreditXQuery holds query params
    type PaymentCheckCreditXQuery struct { 
        Aggregator string  `url:"aggregator,omitempty"`  
    }
    
    // CheckCredit API to fetch the customer credit summary
    func (pa *Payment)  CheckCredit(xQuery PaymentCheckCreditXQuery) (CheckCreditResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             checkCreditResponse CheckCreditResponse
	    )

        

        
            
                
            
        

        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "get",
            "/service/application/payment/v1.0/check-credits/",
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return CheckCreditResponse{}, err
	    }
        
        err = json.Unmarshal(response, &checkCreditResponse)
        if err != nil {
            return CheckCreditResponse{}, common.NewFDKError(err.Error())
        }
         return checkCreditResponse, nil
        
    }
          
    
    
    
  
    
    
    // CustomerOnboard API to fetch the customer credit summary
    func (pa *Payment)  CustomerOnboard(body  CustomerOnboardingRequest) (CustomerOnboardingResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             customerOnboardResponse CustomerOnboardingResponse
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return CustomerOnboardingResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return CustomerOnboardingResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            pa.config,
            "post",
            "/service/application/payment/v1.0/credit-onboard/",
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return CustomerOnboardingResponse{}, err
	    }
        
        err = json.Unmarshal(response, &customerOnboardResponse)
        if err != nil {
            return CustomerOnboardingResponse{}, common.NewFDKError(err.Error())
        }
         return customerOnboardResponse, nil
        
    }
          
    

    // Order ...
    type Order struct {
        config *AppConfig
    }
    // NewOrder ...
    func NewOrder(config *AppConfig) *Order {
        return &Order{config}
    }
    
    
    
  
    
    
    //OrderGetOrdersXQuery holds query params
    type OrderGetOrdersXQuery struct { 
        PageNo float64  `url:"page_no,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"` 
        FromDate string  `url:"from_date,omitempty"` 
        ToDate string  `url:"to_date,omitempty"` 
        Status float64  `url:"status,omitempty"`  
    }
    
    // GetOrders Get all orders
    func (or *Order)  GetOrders(xQuery OrderGetOrdersXQuery) (OrderList, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getOrdersResponse OrderList
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return OrderList{}, err
	    }
        
        err = json.Unmarshal(response, &getOrdersResponse)
        if err != nil {
            return OrderList{}, common.NewFDKError(err.Error())
        }
         return getOrdersResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetOrderById Get details of an order
    func (or *Order)  GetOrderById(OrderID string) (OrderById, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getOrderByIdResponse OrderById
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
            return OrderById{}, err
	    }
        
        err = json.Unmarshal(response, &getOrderByIdResponse)
        if err != nil {
            return OrderById{}, common.NewFDKError(err.Error())
        }
         return getOrderByIdResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetShipmentById Get details of a shipment
    func (or *Order)  GetShipmentById(ShipmentID string) (ShipmentById, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getShipmentByIdResponse ShipmentById
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
            return ShipmentById{}, err
	    }
        
        err = json.Unmarshal(response, &getShipmentByIdResponse)
        if err != nil {
            return ShipmentById{}, common.NewFDKError(err.Error())
        }
         return getShipmentByIdResponse, nil
        
    }
          
    
    
    
  
    
    
    //OrderGetShipmentReasonsXQuery holds query params
    type OrderGetShipmentReasonsXQuery struct { 
        BagID float64  `url:"bag_id,omitempty"`  
    }
    
    // GetShipmentReasons Get reasons behind full or partial cancellation of a shipment
    func (or *Order)  GetShipmentReasons(ShipmentID string, xQuery OrderGetShipmentReasonsXQuery) (ShipmentReasons, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getShipmentReasonsResponse ShipmentReasons
	    )

        

        
            
                
            
        

        
        
        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            or.config,
            "get",
            fmt.Sprintf("/service/application/order/v1.0/orders/shipments/%s/reasons",ShipmentID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ShipmentReasons{}, err
	    }
        
        err = json.Unmarshal(response, &getShipmentReasonsResponse)
        if err != nil {
            return ShipmentReasons{}, common.NewFDKError(err.Error())
        }
         return getShipmentReasonsResponse, nil
        
    }
          
    
    
    
  
    
    
    // UpdateShipmentStatus Update the shipment status
    func (or *Order)  UpdateShipmentStatus(ShipmentID string, body  ShipmentStatusUpdateBody) (ShipmentStatusUpdate, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             updateShipmentStatusResponse ShipmentStatusUpdate
	    )

        
            
        
            
        
            
        

        

        
        
        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return ShipmentStatusUpdate{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return ShipmentStatusUpdate{}, common.NewFDKError(err.Error())
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
            return ShipmentStatusUpdate{}, err
	    }
        
        err = json.Unmarshal(response, &updateShipmentStatusResponse)
        if err != nil {
            return ShipmentStatusUpdate{}, common.NewFDKError(err.Error())
        }
         return updateShipmentStatusResponse, nil
        
    }
          
    
    
    
  
    
    
    // TrackShipment Track shipment
    func (or *Order)  TrackShipment(ShipmentID string) (ShipmentTrack, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             trackShipmentResponse ShipmentTrack
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
            return ShipmentTrack{}, err
	    }
        
        err = json.Unmarshal(response, &trackShipmentResponse)
        if err != nil {
            return ShipmentTrack{}, common.NewFDKError(err.Error())
        }
         return trackShipmentResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetPosOrderById Get POS Order
    func (or *Order)  GetPosOrderById(OrderID string) (PosOrderById, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getPosOrderByIdResponse PosOrderById
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
            return PosOrderById{}, err
	    }
        
        err = json.Unmarshal(response, &getPosOrderByIdResponse)
        if err != nil {
            return PosOrderById{}, common.NewFDKError(err.Error())
        }
         return getPosOrderByIdResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetCustomerDetailsByShipmentId Get Customer Details by Shipment Id
    func (or *Order)  GetCustomerDetailsByShipmentId(OrderID string, ShipmentID string) (CustomerDetailsByShipmentId, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getCustomerDetailsByShipmentIdResponse CustomerDetailsByShipmentId
	    )

        

        

        
        
        
        
        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            or.config,
            "get",
            fmt.Sprintf("/service/application/order/v1.0/orders/%s/shipments/%s/customer-details",OrderID,ShipmentID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return CustomerDetailsByShipmentId{}, err
	    }
        
        err = json.Unmarshal(response, &getCustomerDetailsByShipmentIdResponse)
        if err != nil {
            return CustomerDetailsByShipmentId{}, common.NewFDKError(err.Error())
        }
         return getCustomerDetailsByShipmentIdResponse, nil
        
    }
          
    
    
    
  
    
    
    // SendOtpToShipmentCustomer Send and Resend Otp code to Order-Shipment customer
    func (or *Order)  SendOtpToShipmentCustomer(OrderID string, ShipmentID string) (sendOTPApplicationResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             sendOtpToShipmentCustomerResponse sendOTPApplicationResponse
	    )

        

        

        
        
        
        
        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            or.config,
            "post",
            fmt.Sprintf("/service/application/order/v1.0/orders/%s/shipments/%s/otp/send/",OrderID,ShipmentID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return sendOTPApplicationResponse{}, err
	    }
        
        err = json.Unmarshal(response, &sendOtpToShipmentCustomerResponse)
        if err != nil {
            return sendOTPApplicationResponse{}, common.NewFDKError(err.Error())
        }
         return sendOtpToShipmentCustomerResponse, nil
        
    }
          
    
    
    
  
    
    
    // VerifyOtpShipmentCustomer Verify Otp code
    func (or *Order)  VerifyOtpShipmentCustomer(OrderID string, ShipmentID string, body  ReqBodyVerifyOTPShipment) (ResponseVerifyOTPShipment, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             verifyOtpShipmentCustomerResponse ResponseVerifyOTPShipment
	    )

        
            
        
            
        

        

        
        
        
        
        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return ResponseVerifyOTPShipment{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return ResponseVerifyOTPShipment{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            or.config,
            "post",
            fmt.Sprintf("/service/application/order/v1.0/orders/%s/shipments/%s/otp/verify",OrderID,ShipmentID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return ResponseVerifyOTPShipment{}, err
	    }
        
        err = json.Unmarshal(response, &verifyOtpShipmentCustomerResponse)
        if err != nil {
            return ResponseVerifyOTPShipment{}, common.NewFDKError(err.Error())
        }
         return verifyOtpShipmentCustomerResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetInvoiceByShipmentId Get Invoice URL
    func (or *Order)  GetInvoiceByShipmentId(ShipmentID string) (ResponseGetInvoiceShipment, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getInvoiceByShipmentIdResponse ResponseGetInvoiceShipment
	    )

        

        

        
        
        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            or.config,
            "get",
            fmt.Sprintf("/service/application/order/v1.0/orders/shipments/%s/invoice",ShipmentID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ResponseGetInvoiceShipment{}, err
	    }
        
        err = json.Unmarshal(response, &getInvoiceByShipmentIdResponse)
        if err != nil {
            return ResponseGetInvoiceShipment{}, common.NewFDKError(err.Error())
        }
         return getInvoiceByShipmentIdResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetCreditNoteByShipmentId Get Credit Note URL
    func (or *Order)  GetCreditNoteByShipmentId(ShipmentID string) (ResponseGetCreditNoteShipment, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getCreditNoteByShipmentIdResponse ResponseGetCreditNoteShipment
	    )

        

        

        
        
        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            or.config,
            "get",
            fmt.Sprintf("/service/application/order/v1.0/orders/shipments/%s/credit-note/",ShipmentID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return ResponseGetCreditNoteShipment{}, err
	    }
        
        err = json.Unmarshal(response, &getCreditNoteByShipmentIdResponse)
        if err != nil {
            return ResponseGetCreditNoteShipment{}, common.NewFDKError(err.Error())
        }
         return getCreditNoteByShipmentIdResponse, nil
        
    }
          
    

    // Rewards ...
    type Rewards struct {
        config *AppConfig
    }
    // NewRewards ...
    func NewRewards(config *AppConfig) *Rewards {
        return &Rewards{config}
    }
    
    
    
  
    
    
    // GetPointsOnProduct Get the eligibility of reward points on a product
    func (re *Rewards)  GetPointsOnProduct(body  CatalogueOrderRequest) (CatalogueOrderResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getPointsOnProductResponse CatalogueOrderResponse
	    )

        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return CatalogueOrderResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return CatalogueOrderResponse{}, common.NewFDKError(err.Error())
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
            return CatalogueOrderResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getPointsOnProductResponse)
        if err != nil {
            return CatalogueOrderResponse{}, common.NewFDKError(err.Error())
        }
         return getPointsOnProductResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetOfferByName Get offer by name
    func (re *Rewards)  GetOfferByName(Name string) (Offer, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getOfferByNameResponse Offer
	    )

        

        

        
        
        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            re.config,
            "get",
            fmt.Sprintf("/service/application/rewards/v1.0/offers/%s/",Name),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return Offer{}, err
	    }
        
        err = json.Unmarshal(response, &getOfferByNameResponse)
        if err != nil {
            return Offer{}, common.NewFDKError(err.Error())
        }
         return getOfferByNameResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetOrderDiscount Calculates the discount on order-amount
    func (re *Rewards)  GetOrderDiscount(body  OrderDiscountRequest) (OrderDiscountResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getOrderDiscountResponse OrderDiscountResponse
	    )

        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return OrderDiscountResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return OrderDiscountResponse{}, common.NewFDKError(err.Error())
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
            return OrderDiscountResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getOrderDiscountResponse)
        if err != nil {
            return OrderDiscountResponse{}, common.NewFDKError(err.Error())
        }
         return getOrderDiscountResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetUserPoints Get reward points available with a user
    func (re *Rewards)  GetUserPoints() (PointsResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getUserPointsResponse PointsResponse
	    )

        

        

        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            re.config,
            "get",
            "/service/application/rewards/v1.0/user/points/",
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return PointsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getUserPointsResponse)
        if err != nil {
            return PointsResponse{}, common.NewFDKError(err.Error())
        }
         return getUserPointsResponse, nil
        
    }
          
    
    
    
  
    
    
    //RewardsGetUserPointsHistoryXQuery holds query params
    type RewardsGetUserPointsHistoryXQuery struct { 
        PageID string  `url:"page_id,omitempty"` 
        PageSize float64  `url:"page_size,omitempty"`  
    }
    
    // GetUserPointsHistory Get all transactions of reward points
    func (re *Rewards)  GetUserPointsHistory(xQuery RewardsGetUserPointsHistoryXQuery) (PointsHistoryResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getUserPointsHistoryResponse PointsHistoryResponse
	    )

        

        
            
                
            
                
            
        

        
    
         
        
        
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
            return PointsHistoryResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getUserPointsHistoryResponse)
        if err != nil {
            return PointsHistoryResponse{}, common.NewFDKError(err.Error())
        }
         return getUserPointsHistoryResponse, nil
        
    }
          
            
            
            
            
                
                    
                    
                    
                    
                        
                    
                
                    
                    
                    
                        
                    
                    
                
            
            // GetUserPointsHistoryPaginator Get all transactions of reward points  
            func (re *Rewards)  GetUserPointsHistoryPaginator( xQuery RewardsGetUserPointsHistoryXQuery ) *common.Paginator {
                paginator := common.NewPaginator("cursor")
                 
                 
                 
                 xQuery.PageID = paginator.NextID
                 
                 
                 
                 
                 
                 
                 
                paginator.Next = func() (interface{}, error) {
                    response, err := re.GetUserPointsHistory(xQuery)
                    if response.Page.HasNext {
                        paginator.SetPaginator(response.Page.HasNext, int(response.Page.Current+1), response.Page.NextID)
                    }
                    return response, err
                }
                return paginator
            }
       
    
    
    
  
    
    
    // GetUserReferralDetails Get referral details of a user
    func (re *Rewards)  GetUserReferralDetails() (ReferralDetailsResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getUserReferralDetailsResponse ReferralDetailsResponse
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
            return ReferralDetailsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getUserReferralDetailsResponse)
        if err != nil {
            return ReferralDetailsResponse{}, common.NewFDKError(err.Error())
        }
         return getUserReferralDetailsResponse, nil
        
    }
          
    
    
    
  
    
    
    // RedeemReferralCode Redeems a referral code and credits reward points to users
    func (re *Rewards)  RedeemReferralCode(body  RedeemReferralCodeRequest) (RedeemReferralCodeResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             redeemReferralCodeResponse RedeemReferralCodeResponse
	    )

        
            
        
            
        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return RedeemReferralCodeResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return RedeemReferralCodeResponse{}, common.NewFDKError(err.Error())
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
            return RedeemReferralCodeResponse{}, err
	    }
        
        err = json.Unmarshal(response, &redeemReferralCodeResponse)
        if err != nil {
            return RedeemReferralCodeResponse{}, common.NewFDKError(err.Error())
        }
         return redeemReferralCodeResponse, nil
        
    }
          
    

    // PosCart ...
    type PosCart struct {
        config *AppConfig
    }
    // NewPosCart ...
    func NewPosCart(config *AppConfig) *PosCart {
        return &PosCart{config}
    }
    
    
    
  
    
    
    //PosCartGetCartXQuery holds query params
    type PosCartGetCartXQuery struct { 
        ID string  `url:"id,omitempty"` 
        I bool  `url:"i,omitempty"` 
        B bool  `url:"b,omitempty"` 
        AssignCardID float64  `url:"assign_card_id,omitempty"` 
        AreaCode string  `url:"area_code,omitempty"`  
    }
    
    // GetCart Fetch all items added to the cart
    func (po *PosCart)  GetCart(xQuery PosCartGetCartXQuery) (CartDetailResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getCartResponse CartDetailResponse
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return CartDetailResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCartResponse)
        if err != nil {
            return CartDetailResponse{}, common.NewFDKError(err.Error())
        }
         return getCartResponse, nil
        
    }
          
    
    
    
  
    
    
    //PosCartGetCartLastModifiedXQuery holds query params
    type PosCartGetCartLastModifiedXQuery struct { 
        ID string  `url:"id,omitempty"`  
    }
    
    // GetCartLastModified Fetch last-modified timestamp
    func (po *PosCart)  GetCartLastModified(xQuery PosCartGetCartLastModifiedXQuery) (interface{}, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             
	    )

        

        
            
                
            
        

        
    
         
        
        
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
            return []byte{}, err
	    }
        
        return response, nil
        
    }
          
    
    
    
  
    
    
    //PosCartAddItemsXQuery holds query params
    type PosCartAddItemsXQuery struct { 
        I bool  `url:"i,omitempty"` 
        B bool  `url:"b,omitempty"` 
        AreaCode string  `url:"area_code,omitempty"`  
    }
    
    // AddItems Add items to cart
    func (po *PosCart)  AddItems(xQuery PosCartAddItemsXQuery, body  AddCartRequest) (AddCartDetailResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             addItemsResponse AddCartDetailResponse
	    )

        

        
            
                
            
                
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return AddCartDetailResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return AddCartDetailResponse{}, common.NewFDKError(err.Error())
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
            return AddCartDetailResponse{}, err
	    }
        
        err = json.Unmarshal(response, &addItemsResponse)
        if err != nil {
            return AddCartDetailResponse{}, common.NewFDKError(err.Error())
        }
         return addItemsResponse, nil
        
    }
          
    
    
    
  
    
    
    //PosCartUpdateCartXQuery holds query params
    type PosCartUpdateCartXQuery struct { 
        ID string  `url:"id,omitempty"` 
        I bool  `url:"i,omitempty"` 
        B bool  `url:"b,omitempty"` 
        AreaCode string  `url:"area_code,omitempty"`  
    }
    
    // UpdateCart Update items in the cart
    func (po *PosCart)  UpdateCart(xQuery PosCartUpdateCartXQuery, body  UpdateCartRequest) (UpdateCartDetailResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             updateCartResponse UpdateCartDetailResponse
	    )

        

        
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return UpdateCartDetailResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return UpdateCartDetailResponse{}, common.NewFDKError(err.Error())
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
            return UpdateCartDetailResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateCartResponse)
        if err != nil {
            return UpdateCartDetailResponse{}, common.NewFDKError(err.Error())
        }
         return updateCartResponse, nil
        
    }
          
    
    
    
  
    
    
    //PosCartGetItemCountXQuery holds query params
    type PosCartGetItemCountXQuery struct { 
        ID string  `url:"id,omitempty"`  
    }
    
    // GetItemCount Count items in the cart
    func (po *PosCart)  GetItemCount(xQuery PosCartGetItemCountXQuery) (CartItemCountResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getItemCountResponse CartItemCountResponse
	    )

        

        
            
                
            
        

        
    
         
        
        
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
            return CartItemCountResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getItemCountResponse)
        if err != nil {
            return CartItemCountResponse{}, common.NewFDKError(err.Error())
        }
         return getItemCountResponse, nil
        
    }
          
    
    
    
  
    
    
    //PosCartGetCouponsXQuery holds query params
    type PosCartGetCouponsXQuery struct { 
        ID string  `url:"id,omitempty"`  
    }
    
    // GetCoupons Fetch Coupon
    func (po *PosCart)  GetCoupons(xQuery PosCartGetCouponsXQuery) (GetCouponResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getCouponsResponse GetCouponResponse
	    )

        

        
            
                
            
        

        
    
         
        
        
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
            return GetCouponResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCouponsResponse)
        if err != nil {
            return GetCouponResponse{}, common.NewFDKError(err.Error())
        }
         return getCouponsResponse, nil
        
    }
          
    
    
    
  
    
    
    //PosCartApplyCouponXQuery holds query params
    type PosCartApplyCouponXQuery struct { 
        I bool  `url:"i,omitempty"` 
        B bool  `url:"b,omitempty"` 
        P bool  `url:"p,omitempty"` 
        ID string  `url:"id,omitempty"`  
    }
    
    // ApplyCoupon Apply Coupon
    func (po *PosCart)  ApplyCoupon(xQuery PosCartApplyCouponXQuery, body  ApplyCouponRequest) (CartDetailResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             applyCouponResponse CartDetailResponse
	    )

        

        
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return CartDetailResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return CartDetailResponse{}, common.NewFDKError(err.Error())
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
            return CartDetailResponse{}, err
	    }
        
        err = json.Unmarshal(response, &applyCouponResponse)
        if err != nil {
            return CartDetailResponse{}, common.NewFDKError(err.Error())
        }
         return applyCouponResponse, nil
        
    }
          
    
    
    
  
    
    
    //PosCartRemoveCouponXQuery holds query params
    type PosCartRemoveCouponXQuery struct { 
        ID string  `url:"id,omitempty"`  
    }
    
    // RemoveCoupon Remove Coupon Applied
    func (po *PosCart)  RemoveCoupon(xQuery PosCartRemoveCouponXQuery) (CartDetailResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             removeCouponResponse CartDetailResponse
	    )

        

        
            
                
            
        

        
    
         
        
        
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
            return CartDetailResponse{}, err
	    }
        
        err = json.Unmarshal(response, &removeCouponResponse)
        if err != nil {
            return CartDetailResponse{}, common.NewFDKError(err.Error())
        }
         return removeCouponResponse, nil
        
    }
          
    
    
    
  
    
    
    //PosCartGetBulkDiscountOffersXQuery holds query params
    type PosCartGetBulkDiscountOffersXQuery struct { 
        ItemID float64  `url:"item_id,omitempty"` 
        ArticleID string  `url:"article_id,omitempty"` 
        UID float64  `url:"uid,omitempty"` 
        Slug string  `url:"slug,omitempty"`  
    }
    
    // GetBulkDiscountOffers Get discount offers based on quantity
    func (po *PosCart)  GetBulkDiscountOffers(xQuery PosCartGetBulkDiscountOffersXQuery) (BulkPriceResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getBulkDiscountOffersResponse BulkPriceResponse
	    )

        

        
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return BulkPriceResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getBulkDiscountOffersResponse)
        if err != nil {
            return BulkPriceResponse{}, common.NewFDKError(err.Error())
        }
         return getBulkDiscountOffersResponse, nil
        
    }
          
    
    
    
  
    
    
    //PosCartApplyRewardPointsXQuery holds query params
    type PosCartApplyRewardPointsXQuery struct { 
        ID string  `url:"id,omitempty"` 
        I bool  `url:"i,omitempty"` 
        B bool  `url:"b,omitempty"`  
    }
    
    // ApplyRewardPoints Apply reward points at cart
    func (po *PosCart)  ApplyRewardPoints(xQuery PosCartApplyRewardPointsXQuery, body  RewardPointRequest) (CartDetailResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             applyRewardPointsResponse CartDetailResponse
	    )

        

        
            
                
            
                
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return CartDetailResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return CartDetailResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "post",
            "/service/application/pos/cart/v1.0/redeem/points/",
            nil,
            xQuery,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return CartDetailResponse{}, err
	    }
        
        err = json.Unmarshal(response, &applyRewardPointsResponse)
        if err != nil {
            return CartDetailResponse{}, common.NewFDKError(err.Error())
        }
         return applyRewardPointsResponse, nil
        
    }
          
    
    
    
  
    
    
    //PosCartGetAddressesXQuery holds query params
    type PosCartGetAddressesXQuery struct { 
        CartID string  `url:"cart_id,omitempty"` 
        MobileNo string  `url:"mobile_no,omitempty"` 
        CheckoutMode string  `url:"checkout_mode,omitempty"` 
        Tags string  `url:"tags,omitempty"` 
        IsDefault bool  `url:"is_default,omitempty"`  
    }
    
    // GetAddresses Fetch address
    func (po *PosCart)  GetAddresses(xQuery PosCartGetAddressesXQuery) (GetAddressesResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getAddressesResponse GetAddressesResponse
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return GetAddressesResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getAddressesResponse)
        if err != nil {
            return GetAddressesResponse{}, common.NewFDKError(err.Error())
        }
         return getAddressesResponse, nil
        
    }
          
    
    
    
  
    
    
    // AddAddress Add address to an account
    func (po *PosCart)  AddAddress(body  Address) (SaveAddressResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             addAddressResponse SaveAddressResponse
	    )

        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return SaveAddressResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return SaveAddressResponse{}, common.NewFDKError(err.Error())
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
            return SaveAddressResponse{}, err
	    }
        
        err = json.Unmarshal(response, &addAddressResponse)
        if err != nil {
            return SaveAddressResponse{}, common.NewFDKError(err.Error())
        }
         return addAddressResponse, nil
        
    }
          
    
    
    
  
    
    
    //PosCartGetAddressByIdXQuery holds query params
    type PosCartGetAddressByIdXQuery struct { 
        CartID string  `url:"cart_id,omitempty"` 
        MobileNo string  `url:"mobile_no,omitempty"` 
        CheckoutMode string  `url:"checkout_mode,omitempty"` 
        Tags string  `url:"tags,omitempty"` 
        IsDefault bool  `url:"is_default,omitempty"`  
    }
    
    // GetAddressById Fetch a single address by its ID
    func (po *PosCart)  GetAddressById(ID string, xQuery PosCartGetAddressByIdXQuery) (Address, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getAddressByIdResponse Address
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
        

        
        
        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "get",
            fmt.Sprintf("/service/application/pos/cart/v1.0/address/%s",ID),
            nil,
            xQuery,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return Address{}, err
	    }
        
        err = json.Unmarshal(response, &getAddressByIdResponse)
        if err != nil {
            return Address{}, common.NewFDKError(err.Error())
        }
         return getAddressByIdResponse, nil
        
    }
          
    
    
    
  
    
    
    // UpdateAddress Update address added to an account
    func (po *PosCart)  UpdateAddress(ID string, body  Address) (UpdateAddressResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             updateAddressResponse UpdateAddressResponse
	    )

        

        

        
        
        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return UpdateAddressResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return UpdateAddressResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "put",
            fmt.Sprintf("/service/application/pos/cart/v1.0/address/%s",ID),
            nil,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return UpdateAddressResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateAddressResponse)
        if err != nil {
            return UpdateAddressResponse{}, common.NewFDKError(err.Error())
        }
         return updateAddressResponse, nil
        
    }
          
    
    
    
  
    
    
    // RemoveAddress Remove address associated with an account
    func (po *PosCart)  RemoveAddress(ID string) (DeleteAddressResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             removeAddressResponse DeleteAddressResponse
	    )

        

        

        
        
        
    
         
        
        
        //API call
        rawRequest = NewRequest(
            po.config,
            "delete",
            fmt.Sprintf("/service/application/pos/cart/v1.0/address/%s",ID),
            nil,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return DeleteAddressResponse{}, err
	    }
        
        err = json.Unmarshal(response, &removeAddressResponse)
        if err != nil {
            return DeleteAddressResponse{}, common.NewFDKError(err.Error())
        }
         return removeAddressResponse, nil
        
    }
          
    
    
    
  
    
    
    //PosCartSelectAddressXQuery holds query params
    type PosCartSelectAddressXQuery struct { 
        CartID string  `url:"cart_id,omitempty"` 
        I bool  `url:"i,omitempty"` 
        B bool  `url:"b,omitempty"`  
    }
    
    // SelectAddress Select an address from available addresses
    func (po *PosCart)  SelectAddress(xQuery PosCartSelectAddressXQuery, body  SelectCartAddressRequest) (CartDetailResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             selectAddressResponse CartDetailResponse
	    )

        

        
            
                
            
                
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return CartDetailResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return CartDetailResponse{}, common.NewFDKError(err.Error())
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
            return CartDetailResponse{}, err
	    }
        
        err = json.Unmarshal(response, &selectAddressResponse)
        if err != nil {
            return CartDetailResponse{}, common.NewFDKError(err.Error())
        }
         return selectAddressResponse, nil
        
    }
          
    
    
    
  
    
    
    //PosCartSelectPaymentModeXQuery holds query params
    type PosCartSelectPaymentModeXQuery struct { 
        ID string  `url:"id,omitempty"`  
    }
    
    // SelectPaymentMode Update cart payment
    func (po *PosCart)  SelectPaymentMode(xQuery PosCartSelectPaymentModeXQuery, body  UpdateCartPaymentRequest) (CartDetailResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             selectPaymentModeResponse CartDetailResponse
	    )

        

        
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return CartDetailResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return CartDetailResponse{}, common.NewFDKError(err.Error())
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
            return CartDetailResponse{}, err
	    }
        
        err = json.Unmarshal(response, &selectPaymentModeResponse)
        if err != nil {
            return CartDetailResponse{}, common.NewFDKError(err.Error())
        }
         return selectPaymentModeResponse, nil
        
    }
          
    
    
    
  
    
    
    //PosCartValidateCouponForPaymentXQuery holds query params
    type PosCartValidateCouponForPaymentXQuery struct { 
        ID string  `url:"id,omitempty"` 
        AddressID string  `url:"address_id,omitempty"` 
        PaymentMode string  `url:"payment_mode,omitempty"` 
        PaymentIdentifier string  `url:"payment_identifier,omitempty"` 
        AggregatorName string  `url:"aggregator_name,omitempty"` 
        MerchantCode string  `url:"merchant_code,omitempty"`  
    }
    
    // ValidateCouponForPayment Verify the coupon eligibility against the payment mode
    func (po *PosCart)  ValidateCouponForPayment(xQuery PosCartValidateCouponForPaymentXQuery) (PaymentCouponValidate, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             validateCouponForPaymentResponse PaymentCouponValidate
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return PaymentCouponValidate{}, err
	    }
        
        err = json.Unmarshal(response, &validateCouponForPaymentResponse)
        if err != nil {
            return PaymentCouponValidate{}, common.NewFDKError(err.Error())
        }
         return validateCouponForPaymentResponse, nil
        
    }
          
    
    
    
  
    
    
    //PosCartGetShipmentsXQuery holds query params
    type PosCartGetShipmentsXQuery struct { 
        PickAtStoreUID float64  `url:"pick_at_store_uid,omitempty"` 
        OrderingStoreID float64  `url:"ordering_store_id,omitempty"` 
        P bool  `url:"p,omitempty"` 
        ID string  `url:"id,omitempty"` 
        AddressID string  `url:"address_id,omitempty"` 
        AreaCode string  `url:"area_code,omitempty"` 
        OrderType string  `url:"order_type,omitempty"`  
    }
    
    // GetShipments Get delivery date and options before checkout
    func (po *PosCart)  GetShipments(xQuery PosCartGetShipmentsXQuery) (CartShipmentsResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getShipmentsResponse CartShipmentsResponse
	    )

        

        
            
                
            
                
            
                
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
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
            return CartShipmentsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getShipmentsResponse)
        if err != nil {
            return CartShipmentsResponse{}, common.NewFDKError(err.Error())
        }
         return getShipmentsResponse, nil
        
    }
          
    
    
    
  
    
    
    //PosCartUpdateShipmentsXQuery holds query params
    type PosCartUpdateShipmentsXQuery struct { 
        I bool  `url:"i,omitempty"` 
        P bool  `url:"p,omitempty"` 
        ID string  `url:"id,omitempty"` 
        AddressID string  `url:"address_id,omitempty"` 
        OrderType string  `url:"order_type,omitempty"`  
    }
    
    // UpdateShipments Update shipment delivery type and quantity before checkout
    func (po *PosCart)  UpdateShipments(xQuery PosCartUpdateShipmentsXQuery, body  UpdateCartShipmentRequest) (CartShipmentsResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             updateShipmentsResponse CartShipmentsResponse
	    )

        
            
        

        
            
                
            
                
            
                
            
                
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return CartShipmentsResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return CartShipmentsResponse{}, common.NewFDKError(err.Error())
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
            return CartShipmentsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateShipmentsResponse)
        if err != nil {
            return CartShipmentsResponse{}, common.NewFDKError(err.Error())
        }
         return updateShipmentsResponse, nil
        
    }
          
    
    
    
  
    
    
    //PosCartCheckoutCartXQuery holds query params
    type PosCartCheckoutCartXQuery struct { 
        ID string  `url:"id,omitempty"`  
    }
    
    // CheckoutCart Checkout all items in the cart
    func (po *PosCart)  CheckoutCart(xQuery PosCartCheckoutCartXQuery, body  CartPosCheckoutDetailRequest) (CartCheckoutResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             checkoutCartResponse CartCheckoutResponse
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        
            
        

        
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return CartCheckoutResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return CartCheckoutResponse{}, common.NewFDKError(err.Error())
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
            return CartCheckoutResponse{}, err
	    }
        
        err = json.Unmarshal(response, &checkoutCartResponse)
        if err != nil {
            return CartCheckoutResponse{}, common.NewFDKError(err.Error())
        }
         return checkoutCartResponse, nil
        
    }
          
    
    
    
  
    
    
    //PosCartUpdateCartMetaXQuery holds query params
    type PosCartUpdateCartMetaXQuery struct { 
        ID string  `url:"id,omitempty"`  
    }
    
    // UpdateCartMeta Update the cart meta
    func (po *PosCart)  UpdateCartMeta(xQuery PosCartUpdateCartMetaXQuery, body  CartMetaRequest) (CartMetaResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             updateCartMetaResponse CartMetaResponse
	    )

        

        
            
                
            
        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return CartMetaResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return CartMetaResponse{}, common.NewFDKError(err.Error())
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
            return CartMetaResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateCartMetaResponse)
        if err != nil {
            return CartMetaResponse{}, common.NewFDKError(err.Error())
        }
         return updateCartMetaResponse, nil
        
    }
          
    
    
    
  
    
    
    //PosCartGetAvailableDeliveryModesXQuery holds query params
    type PosCartGetAvailableDeliveryModesXQuery struct { 
        AreaCode string  `url:"area_code,omitempty"` 
        ID string  `url:"id,omitempty"`  
    }
    
    // GetAvailableDeliveryModes Get available delivery modes for cart
    func (po *PosCart)  GetAvailableDeliveryModes(xQuery PosCartGetAvailableDeliveryModesXQuery) (CartDeliveryModesResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getAvailableDeliveryModesResponse CartDeliveryModesResponse
	    )

        

        
            
                
            
                
            
        

        
    
         
        
        
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
            return CartDeliveryModesResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getAvailableDeliveryModesResponse)
        if err != nil {
            return CartDeliveryModesResponse{}, common.NewFDKError(err.Error())
        }
         return getAvailableDeliveryModesResponse, nil
        
    }
          
    
    
    
  
    
    
    //PosCartGetStoreAddressByUidXQuery holds query params
    type PosCartGetStoreAddressByUidXQuery struct { 
        StoreUID float64  `url:"store_uid,omitempty"`  
    }
    
    // GetStoreAddressByUid Get list of stores for give uids
    func (po *PosCart)  GetStoreAddressByUid(xQuery PosCartGetStoreAddressByUidXQuery) (StoreDetailsResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getStoreAddressByUidResponse StoreDetailsResponse
	    )

        

        
            
                
            
        

        
    
         
        
        
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
            return StoreDetailsResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getStoreAddressByUidResponse)
        if err != nil {
            return StoreDetailsResponse{}, common.NewFDKError(err.Error())
        }
         return getStoreAddressByUidResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetCartShareLink Generate token for sharing the cart
    func (po *PosCart)  GetCartShareLink(body  GetShareCartLinkRequest) (GetShareCartLinkResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getCartShareLinkResponse GetShareCartLinkResponse
	    )

        

        

        
    
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return GetShareCartLinkResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return GetShareCartLinkResponse{}, common.NewFDKError(err.Error())
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
            return GetShareCartLinkResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCartShareLinkResponse)
        if err != nil {
            return GetShareCartLinkResponse{}, common.NewFDKError(err.Error())
        }
         return getCartShareLinkResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetCartSharedItems Get details of a shared cart
    func (po *PosCart)  GetCartSharedItems(Token string) (SharedCartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getCartSharedItemsResponse SharedCartResponse
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
            return SharedCartResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getCartSharedItemsResponse)
        if err != nil {
            return SharedCartResponse{}, common.NewFDKError(err.Error())
        }
         return getCartSharedItemsResponse, nil
        
    }
          
    
    
    
  
    
    
    // UpdateCartWithSharedItems Merge or replace existing cart
    func (po *PosCart)  UpdateCartWithSharedItems(Token string, Action string) (SharedCartResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             updateCartWithSharedItemsResponse SharedCartResponse
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
            return SharedCartResponse{}, err
	    }
        
        err = json.Unmarshal(response, &updateCartWithSharedItemsResponse)
        if err != nil {
            return SharedCartResponse{}, common.NewFDKError(err.Error())
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
    
    
    
  
    
    
    // GetPincodeCity Get Pincode API
    func (lo *Logistic)  GetPincodeCity(Pincode string, XApplicationID string) (PincodeApiResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getPincodeCityResponse PincodeApiResponse
	    )

        

        

        
        
        
    
        
        //Adding extra headers
        var xHeaders = make(map[string]string) 
        
         
         xHeaders["x-application-id"] =  XApplicationID
         
        
        
        //API call
        rawRequest = NewRequest(
            lo.config,
            "get",
            fmt.Sprintf("/service/application/logistics/v1.0/pincode/%s",Pincode),
            xHeaders,
            nil,
            nil)
        response, err = rawRequest.Execute()
        if err != nil {
            return PincodeApiResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getPincodeCityResponse)
        if err != nil {
            return PincodeApiResponse{}, common.NewFDKError(err.Error())
        }
         return getPincodeCityResponse, nil
        
    }
          
    
    
    
  
    
    
    // GetTatProduct Get TAT API
    func (lo *Logistic)  GetTatProduct(XApplicationID string, body  TATViewRequest) (TATViewResponse, error){
        var (
            rawRequest  *RawRequest
            response    []byte
            err         error
             getTatProductResponse TATViewResponse
	    )

        
            
        
            
        
            
        
            
        
            
        
            
        

        

        
    
        
        //Adding extra headers
        var xHeaders = make(map[string]string) 
        
         
         xHeaders["x-application-id"] =  XApplicationID
         
        
        
        //Parse req body to map
        var reqBody map[string]interface{}
        reqBodyJSON, err := json.Marshal(body)
        if err != nil {
          
             return TATViewResponse{}, common.NewFDKError(err.Error())
        }
        err = json.Unmarshal([]byte(reqBodyJSON), &reqBody)
        if err != nil {
             
             return TATViewResponse{}, common.NewFDKError(err.Error())
        }
        
        //API call
        rawRequest = NewRequest(
            lo.config,
            "post",
            "/service/application/logistics/v1.0/",
            xHeaders,
            nil,
            reqBody)
        response, err = rawRequest.Execute()
        if err != nil {
            return TATViewResponse{}, err
	    }
        
        err = json.Unmarshal(response, &getTatProductResponse)
        if err != nil {
            return TATViewResponse{}, common.NewFDKError(err.Error())
        }
         return getTatProductResponse, nil
        
    }
          
    

