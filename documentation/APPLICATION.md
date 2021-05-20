# FDK Application Front API Documentaion


* [Catalog](#Catalog) - Catalog API's allows you to access list of products, prices, seller details, similar features, variants and many more useful features.  
* [Cart](#Cart) - Cart APIs 
* [Lead](#Lead) - Handles communication between Staff and Users 
* [Theme](#Theme) - Responsible for themes 
* [User](#User) - Authentication Service 
* [Content](#Content) - Content System 
* [Communication](#Communication) - Manages email, sms, push notifications sent to users 
* [Share](#Share) - Short link and QR Code 
* [FileStorage](#FileStorage) - File Storage 
* [Configuration](#Configuration) - Application configuration apis 
* [Payment](#Payment) - Collect payment through many payment gateway i.e Stripe, Razorpay, Juspay etc.into Fynd or Self account 
* [Order](#Order) - Handles Platform websites OMS 
* [Rewards](#Rewards) - Earn and redeem reward points 
* [Feedback](#Feedback) - User Reviews and Rating System 
* [PosCart](#PosCart) - Cart APIs 
* [Logistic](#Logistic) - Handles Platform websites OMS 

----
----

### Classes and Methods


* [Catalog](#Catalog)
  * Methods
    * [getProductDetailBySlug](#getproductdetailbyslug)
    * [getProductSizesBySlug](#getproductsizesbyslug)
    * [getProductPriceBySlug](#getproductpricebyslug)
    * [getProductSellersBySlug](#getproductsellersbyslug)
    * [getProductComparisonBySlugs](#getproductcomparisonbyslugs)
    * [getSimilarComparisonProductBySlug](#getsimilarcomparisonproductbyslug)
    * [getComparedFrequentlyProductBySlug](#getcomparedfrequentlyproductbyslug)
    * [getProductSimilarByIdentifier](#getproductsimilarbyidentifier)
    * [getProductVariantsBySlug](#getproductvariantsbyslug)
    * [getProductStockByIds](#getproductstockbyids)
    * [getProductStockForTimeByIds](#getproductstockfortimebyids)
    * [getProducts](#getproducts)
    * [getBrands](#getbrands)
    * [getBrandDetailBySlug](#getbranddetailbyslug)
    * [getCategories](#getcategories)
    * [getCategoryDetailBySlug](#getcategorydetailbyslug)
    * [getHomeProducts](#gethomeproducts)
    * [getDepartments](#getdepartments)
    * [getSearchResults](#getsearchresults)
    * [getCollections](#getcollections)
    * [getCollectionItemsBySlug](#getcollectionitemsbyslug)
    * [getCollectionDetailBySlug](#getcollectiondetailbyslug)
    * [getFollowedListing](#getfollowedlisting)
    * [followById](#followbyid)
    * [unfollowById](#unfollowbyid)
    * [getFollowerCountById](#getfollowercountbyid)
    * [getFollowIds](#getfollowids)
    * [getStores](#getstores)
    

* [Cart](#Cart)
  * Methods
    * [getCart](#getcart)
    * [getCartLastModified](#getcartlastmodified)
    * [addItems](#additems)
    * [updateCart](#updatecart)
    * [getItemCount](#getitemcount)
    * [getCoupons](#getcoupons)
    * [applyCoupon](#applycoupon)
    * [removeCoupon](#removecoupon)
    * [getBulkDiscountOffers](#getbulkdiscountoffers)
    * [getAddresses](#getaddresses)
    * [addAddress](#addaddress)
    * [getAddressById](#getaddressbyid)
    * [updateAddress](#updateaddress)
    * [removeAddress](#removeaddress)
    * [selectAddress](#selectaddress)
    * [selectPaymentMode](#selectpaymentmode)
    * [validateCouponForPayment](#validatecouponforpayment)
    * [getShipments](#getshipments)
    * [checkoutCart](#checkoutcart)
    * [updateCartMeta](#updatecartmeta)
    * [getCartShareLink](#getcartsharelink)
    * [getCartSharedItems](#getcartshareditems)
    * [updateCartWithSharedItems](#updatecartwithshareditems)
    

* [Lead](#Lead)
  * Methods
    * [getTicket](#getticket)
    * [createHistory](#createhistory)
    * [createTicket](#createticket)
    * [getCustomForm](#getcustomform)
    * [submitCustomForm](#submitcustomform)
    * [getParticipantsInsideVideoRoom](#getparticipantsinsidevideoroom)
    * [getTokenForVideoRoom](#gettokenforvideoroom)
    

* [Theme](#Theme)
  * Methods
    * [getAppliedTheme](#getappliedtheme)
    * [getThemeForPreview](#getthemeforpreview)
    

* [User](#User)
  * Methods
    * [loginWithFacebook](#loginwithfacebook)
    * [loginWithGoogle](#loginwithgoogle)
    * [loginWithGoogleAndroid](#loginwithgoogleandroid)
    * [loginWithGoogleIOS](#loginwithgoogleios)
    * [loginWithOTP](#loginwithotp)
    * [loginWithEmailAndPassword](#loginwithemailandpassword)
    * [sendResetPasswordEmail](#sendresetpasswordemail)
    * [forgotPassword](#forgotpassword)
    * [sendResetToken](#sendresettoken)
    * [loginWithToken](#loginwithtoken)
    * [registerWithForm](#registerwithform)
    * [verifyEmail](#verifyemail)
    * [verifyMobile](#verifymobile)
    * [hasPassword](#haspassword)
    * [updatePassword](#updatepassword)
    * [logout](#logout)
    * [sendOTPOnMobile](#sendotponmobile)
    * [verifyMobileOTP](#verifymobileotp)
    * [sendOTPOnEmail](#sendotponemail)
    * [verifyEmailOTP](#verifyemailotp)
    * [getLoggedInUser](#getloggedinuser)
    * [getListOfActiveSessions](#getlistofactivesessions)
    * [getPlatformConfig](#getplatformconfig)
    * [updateProfile](#updateprofile)
    * [addMobileNumber](#addmobilenumber)
    * [deleteMobileNumber](#deletemobilenumber)
    * [setMobileNumberAsPrimary](#setmobilenumberasprimary)
    * [sendVerificationLinkToMobile](#sendverificationlinktomobile)
    * [addEmail](#addemail)
    * [deleteEmail](#deleteemail)
    * [setEmailAsPrimary](#setemailasprimary)
    * [sendVerificationLinkToEmail](#sendverificationlinktoemail)
    

* [Content](#Content)
  * Methods
    * [getAnnouncements](#getannouncements)
    * [getBlog](#getblog)
    * [getBlogs](#getblogs)
    * [getFaqs](#getfaqs)
    * [getFaqCategories](#getfaqcategories)
    * [getFaqBySlug](#getfaqbyslug)
    * [getFaqCategoryBySlug](#getfaqcategorybyslug)
    * [getFaqsByCategorySlug](#getfaqsbycategoryslug)
    * [getLandingPage](#getlandingpage)
    * [getLegalInformation](#getlegalinformation)
    * [getNavigations](#getnavigations)
    * [getPage](#getpage)
    * [getPages](#getpages)
    * [getSEOConfiguration](#getseoconfiguration)
    * [getSlideshows](#getslideshows)
    * [getSlideshow](#getslideshow)
    * [getSupportInformation](#getsupportinformation)
    * [getTags](#gettags)
    

* [Communication](#Communication)
  * Methods
    * [getCommunicationConsent](#getcommunicationconsent)
    * [upsertCommunicationConsent](#upsertcommunicationconsent)
    * [upsertAppPushtoken](#upsertapppushtoken)
    

* [Share](#Share)
  * Methods
    * [getApplicationQRCode](#getapplicationqrcode)
    * [getProductQRCodeBySlug](#getproductqrcodebyslug)
    * [getCollectionQRCodeBySlug](#getcollectionqrcodebyslug)
    * [getUrlQRCode](#geturlqrcode)
    * [createShortLink](#createshortlink)
    * [getShortLinkByHash](#getshortlinkbyhash)
    * [getOriginalShortLinkByHash](#getoriginalshortlinkbyhash)
    

* [FileStorage](#FileStorage)
  * Methods
    * [startUpload](#startupload)
    * [completeUpload](#completeupload)
    

* [Configuration](#Configuration)
  * Methods
    * [getApplication](#getapplication)
    * [getOwnerInfo](#getownerinfo)
    * [getBasicDetails](#getbasicdetails)
    * [getIntegrationTokens](#getintegrationtokens)
    * [getOrderingStores](#getorderingstores)
    * [getFeatures](#getfeatures)
    * [getContactInfo](#getcontactinfo)
    * [getCurrencies](#getcurrencies)
    * [getCurrencyById](#getcurrencybyid)
    * [getLanguages](#getlanguages)
    * [getOrderingStoreCookie](#getorderingstorecookie)
    * [removeOrderingStoreCookie](#removeorderingstorecookie)
    * [getAppStaffs](#getappstaffs)
    

* [Payment](#Payment)
  * Methods
    * [getAggregatorsConfig](#getaggregatorsconfig)
    * [attachCardToCustomer](#attachcardtocustomer)
    * [getActiveCardAggregator](#getactivecardaggregator)
    * [getActiveUserCards](#getactiveusercards)
    * [deleteUserCard](#deleteusercard)
    * [verifyCustomerForPayment](#verifycustomerforpayment)
    * [verifyAndChargePayment](#verifyandchargepayment)
    * [initialisePayment](#initialisepayment)
    * [checkAndUpdatePaymentStatus](#checkandupdatepaymentstatus)
    * [getPaymentModeRoutes](#getpaymentmoderoutes)
    * [getPosPaymentModeRoutes](#getpospaymentmoderoutes)
    * [getActiveRefundTransferModes](#getactiverefundtransfermodes)
    * [enableOrDisableRefundTransferMode](#enableordisablerefundtransfermode)
    * [getUserBeneficiariesDetail](#getuserbeneficiariesdetail)
    * [verifyIfscCode](#verifyifsccode)
    * [getOrderBeneficiariesDetail](#getorderbeneficiariesdetail)
    * [verifyOtpAndAddBeneficiaryForBank](#verifyotpandaddbeneficiaryforbank)
    * [addBeneficiaryDetails](#addbeneficiarydetails)
    * [verifyOtpAndAddBeneficiaryForWallet](#verifyotpandaddbeneficiaryforwallet)
    * [updateDefaultBeneficiary](#updatedefaultbeneficiary)
    

* [Order](#Order)
  * Methods
    * [getOrders](#getorders)
    * [getOrderById](#getorderbyid)
    * [getShipmentById](#getshipmentbyid)
    * [getShipmentReasons](#getshipmentreasons)
    * [updateShipmentStatus](#updateshipmentstatus)
    * [trackShipment](#trackshipment)
    * [getPosOrderById](#getposorderbyid)
    

* [Rewards](#Rewards)
  * Methods
    * [getPointsOnProduct](#getpointsonproduct)
    * [getOfferByName](#getofferbyname)
    * [getOrderDiscount](#getorderdiscount)
    * [getUserPoints](#getuserpoints)
    * [getUserPointsHistory](#getuserpointshistory)
    * [getUserReferralDetails](#getuserreferraldetails)
    * [redeemReferralCode](#redeemreferralcode)
    

* [Feedback](#Feedback)
  * Methods
    * [createAbuseReport](#createabusereport)
    * [updateAbuseReport](#updateabusereport)
    * [getAbuseReports](#getabusereports)
    * [getAttributes](#getattributes)
    * [createAttribute](#createattribute)
    * [getAttribute](#getattribute)
    * [updateAttribute](#updateattribute)
    * [createComment](#createcomment)
    * [updateComment](#updatecomment)
    * [getComments](#getcomments)
    * [checkEligibility](#checkeligibility)
    * [deleteMedia](#deletemedia)
    * [createMedia](#createmedia)
    * [updateMedia](#updatemedia)
    * [getMedias](#getmedias)
    * [getReviewSummaries](#getreviewsummaries)
    * [createReview](#createreview)
    * [updateReview](#updatereview)
    * [getReviews](#getreviews)
    * [getTemplates](#gettemplates)
    * [createQuestion](#createquestion)
    * [updateQuestion](#updatequestion)
    * [getQuestionAndAnswers](#getquestionandanswers)
    * [getVotes](#getvotes)
    * [createVote](#createvote)
    * [updateVote](#updatevote)
    

* [PosCart](#PosCart)
  * Methods
    * [getCart](#getcart)
    * [getCartLastModified](#getcartlastmodified)
    * [addItems](#additems)
    * [updateCart](#updatecart)
    * [getItemCount](#getitemcount)
    * [getCoupons](#getcoupons)
    * [applyCoupon](#applycoupon)
    * [removeCoupon](#removecoupon)
    * [getBulkDiscountOffers](#getbulkdiscountoffers)
    * [getAddresses](#getaddresses)
    * [addAddress](#addaddress)
    * [getAddressById](#getaddressbyid)
    * [updateAddress](#updateaddress)
    * [removeAddress](#removeaddress)
    * [selectAddress](#selectaddress)
    * [selectPaymentMode](#selectpaymentmode)
    * [validateCouponForPayment](#validatecouponforpayment)
    * [getShipments](#getshipments)
    * [updateShipments](#updateshipments)
    * [checkoutCart](#checkoutcart)
    * [updateCartMeta](#updatecartmeta)
    * [getAvailableDeliveryModes](#getavailabledeliverymodes)
    * [getStoreAddressByUid](#getstoreaddressbyuid)
    * [getCartShareLink](#getcartsharelink)
    * [getCartSharedItems](#getcartshareditems)
    * [updateCartWithSharedItems](#updatecartwithshareditems)
    

* [Logistic](#Logistic)
  * Methods
    * [getTatProduct](#gettatproduct)
    * [getPincodeCity](#getpincodecity)
    


---
---



## Catalog


#### getProductDetailBySlug
Get a product

```golang

 data, err :=  Catalog.GetProductDetailBySlug(Slug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of a product. You can get slug value from the endpoint /service/application/catalog/v1.0/products/ | 




Use this API to retrieve a product by its slug value.

*Success Response:*



Success. Returns a Product object. Check the example shown below or refer `ProductDetail` for more details.


Schema: `ProductDetail`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getProductSizesBySlug
Get the sizes of a product

```golang

 data, err :=  Catalog.GetProductSizesBySlug(Slug, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of a product. You can get slug value from the endpoint /service/application/catalog/v1.0/products/ | 



| xQuery | struct | Includes properties such as `StoreID`



A product can have multiple sizes. Use this API to fetch all the available sizes of a product.

*Success Response:*



Success. Returns a ProductSize object. Check the example shown below or refer `ProductSizes` for more details.


Schema: `ProductSizes`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getProductPriceBySlug
Get the price of a product size at a PIN Code

```golang

 data, err :=  Catalog.GetProductPriceBySlug(Slug, Size, Pincode, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of a product. You can get slug value from the endpoint /service/application/catalog/v1.0/products/ | 


| Size | string | A string indicating the size of the product, e.g. S, M, XL. You can get slug value from the endpoint /service/application/catalog/v1.0/products/sizes | 


| Pincode | string | The PIN Code of the area near which the selling locations should be searched, e.g. 400059 | 



| xQuery | struct | Includes properties such as `StoreID`



Prices may vary for different sizes of a product. Use this API to retrieve the price of a product size at all the selling locations near to a PIN Code.

*Success Response:*



Success. Returns a ProductSizePrice object. Check the example shown below or refer `ProductSizePriceResponse` for more details.


Schema: `ProductSizePriceResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getProductSellersBySlug
Get the sellers of a product size at a PIN Code

```golang

 data, err :=  Catalog.GetProductSellersBySlug(Slug, Size, Pincode, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of a product. You can get slug value from the endpoint /service/application/catalog/v1.0/products/ | 


| Size | string | A string indicating the size of the product, e.g. S, M, XL. You can get slug value from the endpoint /service/application/catalog/v1.0/products/sizes | 


| Pincode | string | The 6-digit PIN Code of the area near which the selling locations should be searched, e.g. 400059 | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`



A product of a particular size may be sold by multiple sellers. Use this API to fetch the sellers having the stock of a particular size at a given PIN Code.

*Success Response:*



Success. Returns a ProductSizeSeller object. Check the example shown below or refer `ProductSizeSellersResponse` for more details.


Schema: `ProductSizeSellersResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getProductComparisonBySlugs
Compare products

```golang

 data, err :=  Catalog.GetProductComparisonBySlugs(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Slug`



Use this API to compare the features of products belonging to the same category. Note that at least one slug is mandatory in the request query.

*Success Response:*



Success. Returns an array of objects containing the attributes for comparision. Check the example shown below or refer `ProductsComparisonResponse` for more details.


Schema: `ProductsComparisonResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getSimilarComparisonProductBySlug
Get comparison between similar products

```golang

 data, err :=  Catalog.GetSimilarComparisonProductBySlug(Slug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of a product. You can get slug value from the endpoint /service/application/catalog/v1.0/products/ | 




Use this API to compare a given product automatically with similar products. Only one slug is needed.

*Success Response:*



Success. Returns an array of objects containing the attributes for comparision. Check the example shown below or refer `ProductCompareResponse` for more details.


Schema: `ProductCompareResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getComparedFrequentlyProductBySlug
Get comparison between frequently compared products with the given product

```golang

 data, err :=  Catalog.GetComparedFrequentlyProductBySlug(Slug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of a product. You can get slug value from the endpoint /service/application/catalog/v1.0/products/ | 




Use this API to compare a given product automatically with products that are frequently compared with it. Only one slug is needed.

*Success Response:*



Success. Returns an array of objects containing the attributes for comparision. Check the example shown below or refer `ProductFrequentlyComparedSimilarResponse` for more details.


Schema: `ProductFrequentlyComparedSimilarResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getProductSimilarByIdentifier
Get similar products

```golang

 data, err :=  Catalog.GetProductSimilarByIdentifier(Slug, SimilarType);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of a product. You can get slug value from the endpoint /service/application/catalog/v1.0/products/ | 


| SimilarType | string | Similarity criteria such as basic, visual, price, seller, category and spec. Visual - Products having similar patterns, Price - Products in similar price range, Seller - Products sold by the same seller, Category - Products belonging to the same category, e.g. sports shoes, Spec - Products having similar specifications, e.g. phones with same memory. | 




Use this API to retrieve products similar to the one specified by its slug. You can search not only similar looking products, but also those that are sold by same seller, or those that belong to the same category, price, specifications, etc.

*Success Response:*



Success. Returns a group of similar products based on type. Check the example shown below or refer `SimilarProductByTypeResponse` for more details.


Schema: `SimilarProductByTypeResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getProductVariantsBySlug
Get variant of a particular product

```golang

 data, err :=  Catalog.GetProductVariantsBySlug(Slug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of a product. You can get slug value from the endpoint /service/application/catalog/v1.0/products/ | 




A product can have a different type of variants such as colour, shade, memory. Use this API to fetch all the available variants of a product using its slug.

*Success Response:*



Success. Returns all variants of a product. Check the example shown below or refer `ProductVariantsResponse` for more details. For `display_type:image`, `color` key will be present otherwise `value` key will be shown.


Schema: `ProductVariantsResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getProductStockByIds
Get the stock of a product

```golang

 data, err :=  Catalog.GetProductStockByIds(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |










| xQuery | struct | Includes properties such as `ItemID`, `Alu`, `SkuCode`, `Ean`, `Upc`



Retrieve the available stock of the products. Use this API to retrieve stock of multiple products (up to 50) at a time.

*Success Response:*



Success. Returns the status of the product stock.Check the example shown below or refer `ProductStockStatusResponse` for more details.


Schema: `ProductStockStatusResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getProductStockForTimeByIds
Get the stock of a product

```golang

 data, err :=  Catalog.GetProductStockForTimeByIds(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |






| xQuery | struct | Includes properties such as `Timestamp`, `PageSize`, `PageID`



Retrieve the available stock of the products. Use this API to get the stock status of products whose inventory is updated at the specified time

*Success Response:*



Success. Returns the status of the product stock.Check the example shown below or refer `ProductStockPolling` for more details.


Schema: `ProductStockPolling`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getProducts
Get all the products

```golang

 data, err :=  Catalog.GetProducts(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |
















| xQuery | struct | Includes properties such as `Q`, `F`, `Filters`, `SortOn`, `PageID`, `PageSize`, `PageNo`, `PageType`



Use this API to list all the products. You may choose a sort order or make arbitrary search queries by entering the product name, brand, category or collection.

*Success Response:*



Success. Returns a paginated list of products..Check the example shown below or refer `ProductListingResponse` for more details.


Schema: `ProductListingResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getBrands
Get all the brands

```golang

 data, err :=  Catalog.GetBrands(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |






| xQuery | struct | Includes properties such as `Department`, `PageNo`, `PageSize`



A brand is the name under which a product is sold. Use this API to list all the brands. You can also filter the brands by department.

*Success Response:*



Success. Returns a paginated list of brands. Check the example shown below or refer `BrandListingResponse` for more details.


Schema: `BrandListingResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getBrandDetailBySlug
Get metadata of a brand

```golang

 data, err :=  Catalog.GetBrandDetailBySlug(Slug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of a brand. You can get slug value from the endpoint /service/application/catalog/v1.0/brands/. | 




Fetch metadata of a brand such as name, information, logo, banner, etc.

*Success Response:*



Success. Returns a metadata object. Check the example shown below or refer `BrandDetailResponse` for more details.


Schema: `BrandDetailResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getCategories
List all the categories

```golang

 data, err :=  Catalog.GetCategories(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Department`



Use this API to list all the categories. You can also filter the categories by department.

*Success Response:*



Success. Returns a list of categories. Check the example shown below or refer `CategoryListingResponse` for more details.


Schema: `CategoryListingResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getCategoryDetailBySlug
Get metadata of a category

```golang

 data, err :=  Catalog.GetCategoryDetailBySlug(Slug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of a brand. You can get slug value from the endpoint /service/application/catalog/v1.0/brands/. | 




Fetch metadata of a category such as name, information, logo, banner, etc.

*Success Response:*



Success. Returns metadata of a category. Check the example shown below or refer `CategoryMetaResponse` for more details.


Schema: `CategoryMetaResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getHomeProducts
List the products

```golang

 data, err :=  Catalog.GetHomeProducts(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |






| xQuery | struct | Includes properties such as `SortOn`, `PageID`, `PageSize`



List all the products associated with a brand, collection or category in a random order.

*Success Response:*



Success. Returns a paginated list of products. Check the example shown below or refer `HomeListingResponse` for more details.


Schema: `HomeListingResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getDepartments
List all the departments

```golang

 data, err :=  Catalog.GetDepartments();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Departments are a way to categorise similar products. A product can lie in multiple departments. For example, a skirt can below to the 'Women's Fashion' Department while a handbag can lie in 'Women's Accessories' Department. Use this API to list all the departments. If successful, returns the list of departments specified in `DepartmentResponse`

*Success Response:*



List of Departments. See example below or refer `DepartmentResponse` for details.


Schema: `DepartmentResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getSearchResults
Get relevant suggestions for a search query

```golang

 data, err :=  Catalog.GetSearchResults(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Q`



Retrieves a list of suggestions for a given search query. Each suggestion is a valid search term that's generated on the basis of query. This is particularly useful to enhance the user experience while using the search tool.

*Success Response:*



Success. Returns a list autocomplete suggestions for the search query `q`. Check the example shown below or refer `AutoCompleteResponse` for more details.


Schema: `AutoCompleteResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getCollections
List all the collections

```golang

 data, err :=  Catalog.GetCollections(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `PageNo`, `PageSize`



Collections are a great way to organize your products and can improve the ability for customers to find items quickly and efficiently.

*Success Response:*



Success. Returns a list of collections. Check the example shown below or refer `GetCollectionListingResponse` for more details.


Schema: `GetCollectionListingResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getCollectionItemsBySlug
Get the items in a collection

```golang

 data, err :=  Catalog.GetCollectionItemsBySlug(Slug, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of a collection. You can get slug value from the endpoint /service/application/catalog/v1.0/collections/. | 











| xQuery | struct | Includes properties such as `F`, `Filters`, `SortOn`, `PageID`, `PageSize`



Get items in a collection specified by its `slug`.

*Success Response:*



Success. Returns a list items in a given collection. Check the example shown below or refer `ProductListingResponse` for more details.


Schema: `ProductListingResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getCollectionDetailBySlug
Get a particular collection

```golang

 data, err :=  Catalog.GetCollectionDetailBySlug(Slug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of a collection. You can get slug value from the endpoint /service/application/catalog/v1.0/collections/. | 




Get the details of a collection by its `slug`.

*Success Response:*



Success. Returns a Collection object. Check the example shown below or refer `CollectionDetailResponse` for more details.


Schema: `CollectionDetailResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getFollowedListing
Get a list of followed Products, Brands, Collections

```golang

 data, err :=  Catalog.GetFollowedListing(CollectionType, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CollectionType | string | Type of collection followed, i.e. products, brands, or collections. | 





| xQuery | struct | Includes properties such as `PageID`, `PageSize`



Users can follow a product they like. This API retrieves the products the user have followed.

*Success Response:*



Success. Returns a Followed resource object. Check the example shown below or refer `GetFollowListingResponse` for more details.


Schema: `GetFollowListingResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### followById
Follow an entity (product/brand/collection)

```golang

 data, err :=  Catalog.FollowById(CollectionType, CollectionID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CollectionType | string | Type of collection followed, i.e. products, brands, or collections. | 


| CollectionID | string | The ID of the collection type. | 




Follow a particular entity such as product, brand, collection specified by its ID.

*Success Response:*



Success. Returns a response object. Check the example shown below or refer `FollowPostResponse` for more details.


Schema: `FollowPostResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### unfollowById
Unfollow an entity (product/brand/collection)

```golang

 data, err :=  Catalog.UnfollowById(CollectionType, CollectionID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CollectionType | string | Type of collection followed, i.e. products, brands, or collections. | 


| CollectionID | string | The ID of the collection type. | 




You can undo a followed product, brand or collection by its ID. This action is referred as _unfollow_.

*Success Response:*



Success. Returns a response object. Check the example shown below or refer `FollowPostResponse` for more details.


Schema: `FollowPostResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getFollowerCountById
Get Follow Count

```golang

 data, err :=  Catalog.GetFollowerCountById(CollectionType, CollectionID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CollectionType | string | Type of collection, i.e. products, brands, or collections. | 


| CollectionID | string | The ID of the collection type. | 




Get the total count of followers for a given collection type and collection ID.

*Success Response:*



Success. Returns the number of followers for a given collection type. Check the example shown below or refer `FollowerCountResponse` for more details.


Schema: `FollowerCountResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getFollowIds
Get the IDs of followed products, brands and collections.

```golang

 data, err :=  Catalog.GetFollowIds(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `CollectionType`



You can get the IDs of all the followed Products, Brands and Collections. Pass collection_type as query parameter to fetch specific Ids

*Success Response:*



Success. Returns the IDs of all the Products, Brands and Collections which were followed. Check the example shown below or refer `FollowIdsResponse` for more details.


Schema: `FollowIdsResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---


#### getStores
Get store meta information.

```golang

 data, err :=  Catalog.GetStores(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |












| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Q`, `Range`, `Latitude`, `Longitude`



Use this API to get a list of stores in a specific application.

*Success Response:*



Success. Returns a list of selling locations. Check the example shown below or refer `StoreListingResponse` for more details.


Schema: `StoreListingResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `ErrorResponse`









---



---


## Cart


#### getCart
Fetch all Items Added to  Cart

```golang

 data, err :=  Cart.GetCart(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |








| xQuery | struct | Includes properties such as `UID`, `I`, `B`, `AssignCardID`



Get all the details of a items added to cart  by uid. If successful, returns a Cart resource in the response body specified in CartResponse

*Success Response:*



The Cart object. See example below or refer CartResponse for details


Schema: `CartResponse`









---


#### getCartLastModified
Fetch Last-Modified timestamp

```golang

 data, err :=  Cart.GetCartLastModified(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `UID`



Fetch Last-Modified timestamp in header metadata

*Success Response:*



Fetch Last-Modified Timestamp Response






---


#### addItems
Add Items to Cart

```golang

 data, err :=  Cart.AddItems(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `I`, `B`

| body |  AddCartRequest | "Request body" 


<p>Add Items to cart. See `AddCartRequest` in schema of request body for the list of attributes needed to add items to a cart. On successful request, returns cart response containing details of items, coupons available etc.these attributes will be fetched from the folowing api's</p>

*Success Response:*



Response of the cart object including all item details included in .the cart,coupons etc.


Schema: `AddCartResponse`


*Examples:*


Product has been added to your cart
```json
{
  "value": {
    "message": "Product has been added to your cart",
    "success": true,
    "cart": {
      "breakup_values": {
        "display": [
          {
            "display": "MRP Total",
            "key": "mrp_total",
            "value": 17486,
            "currency_code": "INR"
          },
          {
            "display": "Discount",
            "key": "discount",
            "value": -3540,
            "currency_code": "INR"
          },
          {
            "display": "Subtotal",
            "key": "subtotal",
            "value": 13946,
            "currency_code": "INR"
          },
          {
            "display": "Total",
            "key": "total",
            "value": 13946,
            "currency_code": "INR"
          }
        ],
        "raw": {
          "cod_charge": 0,
          "convenience_fee": 0,
          "coupon": 0,
          "delivery_charge": 0,
          "discount": -3540,
          "fynd_cash": 0,
          "gst_charges": 1529.96,
          "mrp_total": 17486,
          "subtotal": 13946,
          "total": 13946,
          "vog": 12416.04,
          "you_saved": 0
        },
        "loyalty_points": {
          "total": 0,
          "applicable": 0,
          "is_applied": false,
          "description": "Your cashback, referrals, and refund amount get credited to Fynd Cash which can be redeemed while placing an order."
        },
        "coupon": {
          "type": "cash",
          "code": "",
          "uid": null,
          "value": 0,
          "is_applied": false,
          "message": "Sorry! Invalid Coupon"
        }
      },
      "items": [
        {
          "key": "751083_10",
          "article": {
            "type": "article",
            "uid": "612_9_SE61201_19100302_10",
            "size": "10",
            "seller": {
              "uid": 612,
              "name": "SSR ENTERPRISE"
            },
            "store": {
              "uid": 4431,
              "name": "Motilal Nagar 1, Goregaon"
            },
            "quantity": 4,
            "price": {
              "base": {
                "marked": 3999,
                "effective": 2399,
                "currency_code": "INR"
              },
              "converted": {
                "marked": 3999,
                "effective": 2399,
                "currency_code": "INR"
              }
            }
          },
          "price": {
            "base": {
              "add_on": 4798,
              "marked": 7998,
              "effective": 4798,
              "selling": 4798,
              "currency_code": "INR"
            },
            "converted": {
              "add_on": 4798,
              "marked": 7998,
              "effective": 4798,
              "selling": 4798,
              "currency_code": "INR"
            }
          },
          "availability": {
            "sizes": [
              "10"
            ],
            "other_store_quantity": 2,
            "out_of_stock": false,
            "deliverable": true,
            "is_valid": true
          },
          "product": {
            "type": "product",
            "uid": 751083,
            "name": "Carson 2",
            "slug": "puma-carson-2-751083-6ad98d",
            "brand": {
              "uid": 9,
              "name": "Puma"
            },
            "categories": [
              {
                "uid": 165,
                "name": "Outdoor Sports Shoes"
              }
            ],
            "images": [
              {
                "aspect_ratio": "16:25",
                "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/9_19100302/1_1542807042296.jpg",
                "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/9_19100302/1_1542807042296.jpg"
              }
            ],
            "action": {
              "type": "product",
              "url": "https://api.addsale.com/platform/content/v1/products/puma-carson-2-751083-6ad98d/",
              "query": {
                "product_slug": [
                  "puma-carson-2-751083-6ad98d"
                ]
              }
            }
          },
          "coupon_message": "",
          "quantity": 2,
          "message": "",
          "bulk_offer": {},
          "discount": "41% OFF"
        },
        {
          "key": "246228_S",
          "article": {
            "type": "article",
            "uid": "46_235_TM62_M11029ONDSXNS_S",
            "size": "S",
            "seller": {
              "uid": 46,
              "name": "RELIANCE BRANDS LIMITED"
            },
            "store": {
              "uid": 4550,
              "name": "VR Mall"
            },
            "quantity": 1,
            "price": {
              "base": {
                "marked": 4490,
                "effective": 4490,
                "currency_code": "INR"
              },
              "converted": {
                "marked": 4490,
                "effective": 4490,
                "currency_code": "INR"
              }
            }
          },
          "price": {
            "base": {
              "add_on": 4490,
              "marked": 4490,
              "effective": 4490,
              "selling": 4490,
              "currency_code": "INR"
            },
            "converted": {
              "add_on": 4490,
              "marked": 4490,
              "effective": 4490,
              "selling": 4490,
              "currency_code": "INR"
            }
          },
          "availability": {
            "sizes": [
              "L",
              "M",
              "S",
              "XL",
              "XXL"
            ],
            "other_store_quantity": 0,
            "out_of_stock": false,
            "deliverable": true,
            "is_valid": true
          },
          "product": {
            "type": "product",
            "uid": 246228,
            "name": "Blue Solid T-Shirt",
            "slug": "superdry-blue-solid-t-shirt-2",
            "brand": {
              "uid": 235,
              "name": "Superdry"
            },
            "categories": [
              {
                "uid": 192,
                "name": "T-Shirts"
              }
            ],
            "images": [
              {
                "aspect_ratio": "16:25",
                "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/235_M11029ONDSXNS/1.jpg",
                "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/235_M11029ONDSXNS/1.jpg"
              }
            ],
            "action": {
              "type": "product",
              "url": "https://api.addsale.com/platform/content/v1/products/superdry-blue-solid-t-shirt-2/",
              "query": {
                "product_slug": [
                  "superdry-blue-solid-t-shirt-2"
                ]
              }
            }
          },
          "coupon_message": "",
          "quantity": 1,
          "message": "",
          "bulk_offer": {},
          "discount": ""
        },
        {
          "key": "443175_S",
          "article": {
            "type": "article",
            "uid": "162_207_1271_LJ03LBLUDN88_S",
            "size": "S",
            "seller": {
              "uid": 162,
              "name": "GO FASHION INDIA PRIVATE LIMITED"
            },
            "store": {
              "uid": 5784,
              "name": "Vega City mall"
            },
            "quantity": 3,
            "price": {
              "base": {
                "marked": 1599,
                "effective": 1599,
                "currency_code": "INR"
              },
              "converted": {
                "marked": 1599,
                "effective": 1599,
                "currency_code": "INR"
              }
            }
          },
          "price": {
            "base": {
              "add_on": 1599,
              "marked": 1599,
              "effective": 1599,
              "selling": 1599,
              "currency_code": "INR"
            },
            "converted": {
              "add_on": 1599,
              "marked": 1599,
              "effective": 1599,
              "selling": 1599,
              "currency_code": "INR"
            }
          },
          "availability": {
            "sizes": [
              "XL",
              "M",
              "L",
              "S"
            ],
            "other_store_quantity": 8,
            "out_of_stock": false,
            "deliverable": true,
            "is_valid": true
          },
          "product": {
            "type": "product",
            "uid": 443175,
            "name": "Light Blue Denim Jeggings",
            "slug": "go-colors-light-blue-denim-jeggings-443175-3c688c",
            "brand": {
              "uid": 207,
              "name": "Go Colors"
            },
            "categories": [
              {
                "uid": 267,
                "name": "Jeggings"
              }
            ],
            "images": [
              {
                "aspect_ratio": "16:25",
                "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/207_LJ03LBLUDN88/1_1512382513548.jpg",
                "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/207_LJ03LBLUDN88/1_1512382513548.jpg"
              }
            ],
            "action": {
              "type": "product",
              "url": "https://api.addsale.com/platform/content/v1/products/go-colors-light-blue-denim-jeggings-443175-3c688c/",
              "query": {
                "product_slug": [
                  "go-colors-light-blue-denim-jeggings-443175-3c688c"
                ]
              }
            }
          },
          "coupon_message": "",
          "quantity": 1,
          "message": "",
          "bulk_offer": {},
          "discount": ""
        },
        {
          "key": "778937_OS",
          "article": {
            "type": "article",
            "uid": "686_963_IC68601_PWUPC01977_OS",
            "size": "OS",
            "seller": {
              "uid": 686,
              "name": "INDUS CORPORATION"
            },
            "store": {
              "uid": 5059,
              "name": "Vidyaranyapura Main Road"
            },
            "quantity": 3,
            "price": {
              "base": {
                "marked": 3399,
                "effective": 3059,
                "currency_code": "INR"
              },
              "converted": {
                "marked": 3399,
                "effective": 3059,
                "currency_code": "INR"
              }
            }
          },
          "price": {
            "base": {
              "add_on": 3059,
              "marked": 3399,
              "effective": 3059,
              "selling": 3059,
              "currency_code": "INR"
            },
            "converted": {
              "add_on": 3059,
              "marked": 3399,
              "effective": 3059,
              "selling": 3059,
              "currency_code": "INR"
            }
          },
          "availability": {
            "sizes": [
              "OS"
            ],
            "other_store_quantity": 2,
            "out_of_stock": false,
            "deliverable": true,
            "is_valid": true
          },
          "product": {
            "type": "product",
            "uid": 778937,
            "name": "Colourful Carnival Bouncer",
            "slug": "fisher-price-colourful-carnival-bouncer-778937-fafa1f",
            "brand": {
              "uid": 963,
              "name": "Fisher-Price"
            },
            "categories": [
              {
                "uid": 576,
                "name": "Cradles"
              }
            ],
            "images": [
              {
                "aspect_ratio": "16:25",
                "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/963_PWUPC01977/1_1545308400588.jpg",
                "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/963_PWUPC01977/1_1545308400588.jpg"
              }
            ],
            "action": {
              "type": "product",
              "url": "https://api.addsale.com/platform/content/v1/products/fisher-price-colourful-carnival-bouncer-778937-fafa1f/",
              "query": {
                "product_slug": [
                  "fisher-price-colourful-carnival-bouncer-778937-fafa1f"
                ]
              }
            }
          },
          "coupon_message": "",
          "quantity": 1,
          "message": "",
          "bulk_offer": {},
          "discount": "11% OFF"
        }
      ],
      "delivery_charge_info": "",
      "coupon_text": "View all offers",
      "cart_id": 7927,
      "uid": "7927",
      "gstin": null,
      "checkout_mode": "self",
      "last_modified": "Tue, 03 Sep 2019 06:00:43 GMT",
      "restrict_checkout": false,
      "is_valid": true
    },
    "result": {}
  }
}
```

Sorry, item is out of stock
```json
{
  "value": {
    "message": "Sorry, item is out of stock",
    "success": false,
    "cart": {
      "breakup_values": {
        "raw": {
          "cod_charge": 0,
          "convenience_fee": 0,
          "coupon": 0,
          "delivery_charge": 0,
          "discount": -202000,
          "fynd_cash": 0,
          "gst_charges": 4804.71,
          "mrp_total": 302899,
          "subtotal": 100899,
          "total": 100899,
          "vog": 96094.29,
          "you_saved": 0
        },
        "coupon": {
          "type": "cash",
          "code": "",
          "uid": null,
          "value": 0,
          "is_applied": false,
          "message": "Sorry! Invalid Coupon"
        },
        "display": [
          {
            "display": "MRP Total",
            "key": "mrp_total",
            "value": 302899,
            "currency_code": "INR"
          },
          {
            "display": "Discount",
            "key": "discount",
            "value": -202000,
            "currency_code": "INR"
          },
          {
            "display": "Subtotal",
            "key": "subtotal",
            "value": 100899,
            "currency_code": "INR"
          },
          {
            "display": "Total",
            "key": "total",
            "value": 100899,
            "currency_code": "INR"
          }
        ],
        "loyalty_points": {
          "total": 0,
          "applicable": 0,
          "is_applied": false,
          "description": "Your cashback, referrals, and refund amount get credited to Fynd Cash which can be redeemed while placing an order."
        }
      },
      "items": [
        {
          "bulk_offer": {},
          "discount": "67% OFF",
          "article": {
            "type": "article",
            "uid": "604_902_SSTC60401_636BLUE_1",
            "size": "1",
            "seller": {
              "uid": 604,
              "name": "SHRI SHANTINATH TRADING COMPANY"
            },
            "store": {
              "uid": 4579,
              "name": "Gandhi Nagar"
            },
            "quantity": 108,
            "price": {
              "base": {
                "marked": 2999,
                "effective": 999,
                "currency_code": "INR"
              },
              "converted": {
                "marked": 2999,
                "effective": 999,
                "currency_code": "INR"
              }
            }
          },
          "coupon_message": "",
          "key": "707569_1",
          "availability": {
            "sizes": [
              "1",
              "8",
              "7",
              "2",
              "9",
              "5",
              "3",
              "6"
            ],
            "other_store_quantity": 7,
            "out_of_stock": false,
            "deliverable": true,
            "is_valid": true
          },
          "product": {
            "type": "product",
            "uid": 707569,
            "name": "Blue and Gold Printed Ethnic Set",
            "slug": "aj-dezines-blue-and-gold-printed-ethnic-set-707569-bff01a",
            "brand": {
              "uid": 902,
              "name": ""
            },
            "categories": [
              {
                "uid": 525,
                "name": ""
              }
            ],
            "images": [
              {
                "aspect_ratio": "16:25",
                "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/902_636BLUE/1_1540301094877.jpg",
                "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/902_636BLUE/1_1540301094877.jpg"
              }
            ],
            "action": {
              "type": "product",
              "url": "https://api.addsale.com/v1/products/aj-dezines-blue-and-gold-printed-ethnic-set-707569-bff01a/",
              "query": {
                "product_slug": [
                  "aj-dezines-blue-and-gold-printed-ethnic-set-707569-bff01a"
                ]
              }
            }
          },
          "price": {
            "base": {
              "add_on": 100899,
              "marked": 302899,
              "effective": 100899,
              "selling": 100899,
              "currency_code": "INR"
            },
            "converted": {
              "add_on": 100899,
              "marked": 302899,
              "effective": 100899,
              "selling": 100899,
              "currency_code": "INR"
            }
          },
          "message": "",
          "quantity": 101
        }
      ],
      "delivery_charge_info": "",
      "coupon_text": "View all offers",
      "cart_id": 54,
      "uid": "54",
      "gstin": null,
      "checkout_mode": "self",
      "restrict_checkout": false,
      "is_valid": false,
      "last_modified": "Tue, 03 Sep 2019 09:55:40 GMT"
    },
    "result": {}
  }
}
```









---


#### updateCart
Update Items already added to Cart

```golang

 data, err :=  Cart.UpdateCart(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |






| xQuery | struct | Includes properties such as `UID`, `I`, `B`

| body |  UpdateCartRequest | "Request body" 


Request object containing attributes like item_quantity and item_size which can be updated .these attributes will be fetched from the folowing api's</p> <ul> <li><font color="monochrome">operation</font> Operation for current api call. <b>update_item</b> for update items. <b>remove_item</b> for removing items.</li> <li> <font color="monochrome">item_id</font>  "/platform/content/v1/products/"</li> <li> <font color="monochrome">item_size</font>   "/platform/content/v1/products/{slug}/sizes/"</li> <li> <font color="monochrome">quantity</font>  item quantity (must be greater than or equal to 1)</li> <li> <font color="monochrome">article_id</font>   "/content​/v1​/products​/{identifier}​/sizes​/price​/"</li> <li> <font color="monochrome">item_index</font>  item position in the cart (must be greater than or equal to 0)</li> </ul>

*Success Response:*



Response of the cart object including all item with their updated details included in .the cart,coupons etc..


Schema: `UpdateCartResponse`


*Examples:*


Nothing updated
```json
{
  "value": {
    "cart": {
      "breakup_values": {
        "raw": {
          "cod_charge": 0,
          "convenience_fee": 0,
          "coupon": 0,
          "delivery_charge": 0,
          "discount": -202000,
          "fynd_cash": 0,
          "gst_charges": 4804.71,
          "mrp_total": 302899,
          "subtotal": 100899,
          "total": 100899,
          "vog": 96094.29,
          "you_saved": 0
        },
        "coupon": {
          "type": "cash",
          "code": "",
          "uid": null,
          "value": 0,
          "is_applied": false,
          "message": "Sorry! Invalid Coupon"
        },
        "display": [
          {
            "display": "MRP Total",
            "key": "mrp_total",
            "value": 302899,
            "currency_code": "INR"
          },
          {
            "display": "Discount",
            "key": "discount",
            "value": -202000,
            "currency_code": "INR"
          },
          {
            "display": "Subtotal",
            "key": "subtotal",
            "value": 100899,
            "currency_code": "INR"
          },
          {
            "display": "Total",
            "key": "total",
            "value": 100899,
            "currency_code": "INR"
          }
        ],
        "loyalty_points": {
          "total": 0,
          "applicable": 0,
          "is_applied": false,
          "description": "Your cashback, referrals, and refund amount get credited to Fynd Cash which can be redeemed while placing an order."
        }
      },
      "items": [
        {
          "bulk_offer": {},
          "discount": "67% OFF",
          "article": {
            "type": "article",
            "uid": "604_902_SSTC60401_636BLUE_1",
            "size": "1",
            "seller": {
              "uid": 604,
              "name": "SHRI SHANTINATH TRADING COMPANY"
            },
            "store": {
              "uid": 4579,
              "name": "Gandhi Nagar"
            },
            "quantity": 108,
            "price": {
              "base": {
                "marked": 2999,
                "effective": 999,
                "currency_code": "INR"
              },
              "converted": {
                "marked": 2999,
                "effective": 999,
                "currency_code": "INR"
              }
            }
          },
          "coupon_message": "",
          "key": "707569_1",
          "availability": {
            "sizes": [
              "1",
              "8",
              "7",
              "2",
              "9",
              "5",
              "3",
              "6"
            ],
            "other_store_quantity": 7,
            "out_of_stock": false,
            "deliverable": true,
            "is_valid": true
          },
          "product": {
            "type": "product",
            "uid": 707569,
            "name": "Blue and Gold Printed Ethnic Set",
            "slug": "aj-dezines-blue-and-gold-printed-ethnic-set-707569-bff01a",
            "brand": {
              "uid": 902,
              "name": ""
            },
            "categories": [
              {
                "uid": 525,
                "name": ""
              }
            ],
            "images": [
              {
                "aspect_ratio": "16:25",
                "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/902_636BLUE/1_1540301094877.jpg",
                "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/902_636BLUE/1_1540301094877.jpg"
              }
            ],
            "action": {
              "type": "product",
              "url": "https://api.addsale.com/v1/products/aj-dezines-blue-and-gold-printed-ethnic-set-707569-bff01a/",
              "query": {
                "product_slug": [
                  "aj-dezines-blue-and-gold-printed-ethnic-set-707569-bff01a"
                ]
              }
            }
          },
          "price": {
            "base": {
              "add_on": 100899,
              "marked": 302899,
              "effective": 100899,
              "selling": 100899,
              "currency_code": "INR"
            },
            "converted": {
              "add_on": 100899,
              "marked": 302899,
              "effective": 100899,
              "selling": 100899,
              "currency_code": "INR"
            }
          },
          "message": "",
          "quantity": 101
        }
      ],
      "delivery_charge_info": "",
      "coupon_text": "View all offers",
      "cart_id": 54,
      "uid": "54",
      "gstin": null,
      "checkout_mode": "self",
      "restrict_checkout": false,
      "is_valid": true,
      "last_modified": "Tue, 03 Sep 2019 10:19:20 GMT"
    },
    "result": {
      "707569_90": {
        "success": true,
        "message": "Nothing updated"
      }
    },
    "message": "Nothing updated",
    "success": true
  }
}
```

Item updated in the cart
```json
{
  "value": {
    "cart": {
      "breakup_values": {
        "coupon": {
          "type": "cash",
          "code": "",
          "uid": null,
          "value": 0,
          "is_applied": false,
          "message": "Sorry! Invalid Coupon"
        },
        "loyalty_points": {
          "total": 0,
          "applicable": 0,
          "is_applied": false,
          "description": "Your cashback, referrals, and refund amount get credited to Fynd Cash which can be redeemed while placing an order."
        },
        "raw": {
          "cod_charge": 0,
          "convenience_fee": 0,
          "coupon": 0,
          "delivery_charge": 0,
          "discount": 0,
          "fynd_cash": 0,
          "gst_charges": 838.83,
          "mrp_total": 5499,
          "subtotal": 5499,
          "total": 5499,
          "vog": 4660.17,
          "you_saved": 0
        },
        "display": [
          {
            "display": "MRP Total",
            "key": "mrp_total",
            "value": 5499,
            "currency_code": "INR"
          },
          {
            "display": "Subtotal",
            "key": "subtotal",
            "value": 5499,
            "currency_code": "INR"
          },
          {
            "display": "Total",
            "key": "total",
            "value": 5499,
            "currency_code": "INR"
          }
        ]
      },
      "items": [
        {
          "key": "437414_7",
          "message": "",
          "bulk_offer": {},
          "price": {
            "base": {
              "add_on": 5499,
              "marked": 5499,
              "effective": 5499,
              "selling": 5499,
              "currency_code": "INR"
            },
            "converted": {
              "add_on": 5499,
              "marked": 5499,
              "effective": 5499,
              "selling": 5499,
              "currency_code": "INR"
            }
          },
          "quantity": 1,
          "discount": "",
          "product": {
            "type": "product",
            "uid": 437414,
            "name": "Suede Classic",
            "slug": "puma-suede-classic-437414-6e6bbf",
            "brand": {
              "uid": 9,
              "name": "Puma"
            },
            "categories": [
              {
                "uid": 165,
                "name": "Outdoor Sports Shoes"
              }
            ],
            "images": [
              {
                "aspect_ratio": "16:25",
                "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/9_35656851/1_1511171811830.jpg",
                "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/9_35656851/1_1511171811830.jpg"
              }
            ],
            "action": {
              "type": "product",
              "url": "https://api.addsale.com/platform/content/v1/products/puma-suede-classic-437414-6e6bbf/",
              "query": {
                "product_slug": [
                  "puma-suede-classic-437414-6e6bbf"
                ]
              }
            }
          },
          "article": {
            "type": "article",
            "uid": "507_9_96099_35656851_7",
            "size": "7",
            "seller": {
              "uid": 507,
              "name": "PUMA SPORTS INDIA PVT LTD"
            },
            "store": {
              "uid": 3632,
              "name": "Colaba Causway"
            },
            "quantity": 5,
            "price": {
              "base": {
                "marked": 5499,
                "effective": 5499,
                "currency_code": "INR"
              },
              "converted": {
                "marked": 5499,
                "effective": 5499,
                "currency_code": "INR"
              }
            }
          },
          "coupon_message": "",
          "availability": {
            "sizes": [
              "10",
              "11",
              "6",
              "9",
              "7",
              "8"
            ],
            "other_store_quantity": 22,
            "out_of_stock": false,
            "deliverable": true,
            "is_valid": true
          }
        }
      ],
      "delivery_charge_info": "",
      "coupon_text": "View all offers",
      "cart_id": 12426,
      "uid": "12426",
      "gstin": null,
      "checkout_mode": "self",
      "last_modified": "Thu, 22 Aug 2019 04:51:42 GMT",
      "restrict_checkout": false,
      "is_valid": true
    },
    "result": {
      "437414_7": {
        "success": true,
        "message": "Item updated in the bag"
      }
    },
    "message": "Item updated in the bag",
    "success": true
  }
}
```









---


#### getItemCount
Cart item count

```golang

 data, err :=  Cart.GetItemCount(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `UID`



Get total count of item present in cart

*Success Response:*



OK


Schema: `CartItemCountResponse`









---


#### getCoupons
Fetch Coupon

```golang

 data, err :=  Cart.GetCoupons(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `UID`



Get all the details of a coupons applicable to cart  by uid. If successful, returns a Coupon resource in the response body specified in GetCouponResponse

*Success Response:*



Returns The Couppon object which has list of all available_coupon applicale for the cart. See example below or refer GetCouponResponse for details


Schema: `GetCouponResponse`









---


#### applyCoupon
Apply Coupon

```golang

 data, err :=  Cart.ApplyCoupon(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |








| xQuery | struct | Includes properties such as `I`, `B`, `P`, `UID`

| body |  ApplyCouponRequest | "Request body" 


<p>Apply Coupons on Items added to cart. On successful request, returns cart response containing details of items ,coupons applied etc.these attributes will be consumed by  api</p> <ul> <li> <font color="monochrome">coupon_code</font></li>
</ul>

*Success Response:*



Response of the Coupon object including all item details included in .the cart,coupons applied etc.


Schema: `CartResponse`









---


#### removeCoupon
Remove Coupon Applied

```golang

 data, err :=  Cart.RemoveCoupon(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `UID`



Remove Coupon applied on the cart by passing uid in request body.

*Success Response:*



Response of the Coupon object including all item details included in .the cart,coupons removed etc.


Schema: `CartResponse`









---


#### getBulkDiscountOffers
Get discount offers based on quantity

```golang

 data, err :=  Cart.GetBulkDiscountOffers(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |








| xQuery | struct | Includes properties such as `ItemID`, `ArticleID`, `UID`, `Slug`



List applicable offers along with current, next and best offer for given product. Either one of **uid**, **item_id**, **slug** should be present*

*Success Response:*



Offers found or not found with valid input


Schema: `BulkPriceResponse`


*Examples:*


Offers found
```json
{
  "value": {
    "data": [
      {
        "seller": {
          "uid": 248,
          "name": "MANYAVAR CREATIONS PRIVATE LIMITED"
        },
        "offers": [
          {
            "quantity": 1,
            "auto_applied": true,
            "margin": 10,
            "type": "bundle",
            "price": {
              "marked": 3999,
              "effective": 3999,
              "bulk_effective": 3599.1,
              "currency_code": "INR"
            },
            "total": 3599.1
          },
          {
            "quantity": 3,
            "auto_applied": true,
            "margin": 20,
            "type": "bundle",
            "price": {
              "marked": 3999,
              "effective": 3999,
              "bulk_effective": 3199.2,
              "currency_code": "INR"
            },
            "total": 9597.6
          },
          {
            "quantity": 9,
            "auto_applied": true,
            "margin": 30,
            "type": "bundle",
            "price": {
              "marked": 3999,
              "effective": 3999,
              "bulk_effective": 3443.4444444444,
              "currency_code": "INR"
            },
            "total": 30991,
            "best": true
          }
        ]
      }
    ]
  }
}
```

Offers not found
```json
{
  "value": {
    "data": []
  }
}
```








Unhandled api error


Schema: `Object`









---


#### getAddresses
Fetch Address

```golang

 data, err :=  Cart.GetAddresses(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |










| xQuery | struct | Includes properties such as `UID`, `MobileNo`, `CheckoutMode`, `Tags`, `IsDefault`



Get all the addresses associated with the account. If successful, returns a Address resource in the response body specified in GetAddressesResponse.attibutes listed below are optional <ul> <li> <font color="monochrome">uid</font></li> <li> <font color="monochrome">address_id</font></li> <li> <font color="monochrome">mobile_no</font></li> <li> <font color="monochrome">checkout_mode</font></li> <li> <font color="monochrome">tags</font></li> <li> <font color="monochrome">default</font></li> </ul>

*Success Response:*



Returns The Address object which has list of all address saved for the account. See example below or refer GetAddressesResponse for details


Schema: `GetAddressesResponse`









---


#### addAddress
Add Address to the account

```golang

 data, err :=  Cart.AddAddress(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  Address | "Request body" 


<p>Add Address to account. See `Address` in schema of request body for the list of attributes needed to add Address to account. On successful request, returns response containing address_id ,is_default_address and success message.

*Success Response:*



Return Address Id on successfull completion of the request.


Schema: `SaveAddressResponse`









---


#### getAddressById
Fetch Single Address

```golang

 data, err :=  Cart.GetAddressById(ID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ID | float64 |  | 











| xQuery | struct | Includes properties such as `UID`, `MobileNo`, `CheckoutMode`, `Tags`, `IsDefault`



Get a addresses with the given id. If successful, returns a Address resource in the response body specified in `Address`.attibutes listed below are optional <ul> <li> <font color="monochrome">mobile_no</font></li> <li> <font color="monochrome">checkout_mode</font></li> <li> <font color="monochrome">tags</font></li> <li> <font color="monochrome">default</font></li> </ul>

*Success Response:*



Returns The Address object which has list of all address saved for the account. See example below or refer Address for details


Schema: `Address`









---


#### updateAddress
Update Address alreay added to account

```golang

 data, err :=  Cart.UpdateAddress(ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ID | float64 | Address id | 


| body |  Address | "Request body" 


Request object containing attributes mentioned in  <font color="blue">Address </font> can be updated .these attributes are :</p> <ul> <li> <font color="monochrome">is_default_address</font></li> <li> <font color="monochrome">landmark</font></li> <li> <font color="monochrome">area</font></li> <li> <font color="monochrome">pincode</font></li> <li> <font color="monochrome">email</font></li> <li> <font color="monochrome">address_type</font></li> <li> <font color="monochrome">name</font></li> <li> <font color="monochrome">address_id</font></li> <li> <font color="monochrome">address</font></li> </ul>

*Success Response:*



Response of the Address object containing address_id and sucess message.


Schema: `UpdateAddressResponse`









---


#### removeAddress
Remove Address Associated to the account

```golang

 data, err :=  Cart.RemoveAddress(ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ID | float64 | Address id | 




Delete a Address by it's address_id. Returns an object that tells whether the address was deleted successfully

*Success Response:*



Status object. Tells whether the operation was successful. See example below or refer DeleteAddressResponse


Schema: `DeleteAddressResponse`









---


#### selectAddress
Select Address from All Addresses

```golang

 data, err :=  Cart.SelectAddress(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |






| xQuery | struct | Includes properties such as `UID`, `I`, `B`

| body |  SelectCartAddressRequest | "Request body" 


<p>Select Address from all addresses associated with the account in order to ship the cart items to .that address,otherwise default address will be selected implicitly. See `SelectCartAddressRequest` in schema of request body for the list of attributes needed to select Address from account. On successful request, returns Cart object response.below are the address attributes which needs to be sent. <ul> <li> <font color="monochrome">address_id</font></li> <li> <font color="monochrome">billing_address_id</font></li> <li> <font color="monochrome">uid</font></li> </ul>

*Success Response:*



Response of the Address object containing Cart Object and success message.  .


Schema: `CartResponse`








Address or Pincode Error


Schema: `Object`


*Examples:*


Address Not Found
```json
{
  "value": {
    "status": "ERROR",
    "message": "ADDRESS_NOT_FOUND"
  }
}
```

Pincode Not Serviciable
```json
{
  "value": {
    "status": "ERROR",
    "message": "PINCODE_NOT_SERVICIABLE"
  }
}
```









---


#### selectPaymentMode
Update Cart Payment

```golang

 data, err :=  Cart.SelectPaymentMode(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `UID`

| body |  UpdateCartPaymentRequest | "Request body" 


Update Cart Payment for Your Account

*Success Response:*



Cart response with payment options


Schema: `CartResponse`









---


#### validateCouponForPayment
Get Cart Payment for valid coupon

```golang

 data, err :=  Cart.ValidateCouponForPayment(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |












| xQuery | struct | Includes properties such as `UID`, `AddressID`, `PaymentMode`, `PaymentIdentifier`, `AggregatorName`, `MerchantCode`



Validate coupon for selected payment mode

*Success Response:*



Cart Affiliates.


Schema: `PaymentCouponValidate`









---


#### getShipments
Get delivery date and options before checkout

```golang

 data, err :=  Cart.GetShipments(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |








| xQuery | struct | Includes properties such as `P`, `UID`, `AddressID`, `AreaCode`



Shipment break up item wise with delivery date. Actual                      delivery will be during given dates only. Items will be                      delivered in group of shipments created.

*Success Response:*



OK


Schema: `CartShipmentsResponse`


*Examples:*


Shipment Generated
```json
{
  "value": {
    "items": [],
    "cart_id": 7501,
    "uid": "7501",
    "success": true,
    "error_message": "Note: Your order delivery will be delayed by 7-10 Days",
    "payment_options": {
      "payment_option": [
        {
          "name": "COD",
          "display_name": "Cash on Delivery",
          "display_priority": 1,
          "payment_mode_id": 11,
          "logo": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png",
          "logo_url": {
            "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png",
            "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png"
          },
          "list": []
        },
        {
          "name": "CARD",
          "display_priority": 2,
          "payment_mode_id": 2,
          "display_name": "Card",
          "list": []
        },
        {
          "name": "NB",
          "display_priority": 3,
          "payment_mode_id": 3,
          "display_name": "Net Banking",
          "list": [
            {
              "aggregator_name": "Razorpay",
              "bank_name": "ICICI Bank",
              "bank_code": "ICIC",
              "url": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png"
              },
              "merchant_code": "NB_ICICI",
              "display_priority": 1
            }
          ]
        },
        {
          "name": "WL",
          "display_priority": 4,
          "payment_mode_id": 4,
          "display_name": "Wallet",
          "list": [
            {
              "wallet_name": "Paytm",
              "wallet_code": "paytm",
              "wallet_id": 4,
              "merchant_code": "PAYTM",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/paytm_logo_small.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/paytm_logo_large.png"
              },
              "aggregator_name": "Juspay",
              "display_priority": 1
            }
          ]
        },
        {
          "name": "UPI",
          "display_priority": 9,
          "payment_mode_id": 6,
          "display_name": "UPI",
          "list": [
            {
              "aggregator_name": "UPI_Razorpay",
              "name": "UPI",
              "display_name": "BHIM UPI",
              "code": "UPI",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/upi_100x78.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/upi_150x100.png"
              },
              "merchant_code": "UPI",
              "timeout": 240,
              "retry_count": 0,
              "fynd_vpa": "shopsense.rzp@hdfcbank",
              "intent_flow": true,
              "intent_app_error_list": [
                "com.csam.icici.bank.imobile",
                "in.org.npci.upiapp",
                "com.whatsapp"
              ]
            }
          ]
        },
        {
          "name": "PL",
          "display_priority": 11,
          "payment_mode_id": 1,
          "display_name": "Pay Later",
          "list": [
            {
              "aggregator_name": "Simpl",
              "name": "Simpl",
              "code": "simpl",
              "merchant_code": "SIMPL",
              "logo": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png"
              }
            }
          ]
        }
      ],
      "payment_flows": {
        "Simpl": {
          "data": {
            "gateway": {
              "route": "simpl",
              "entity": "sdk",
              "is_customer_validation_required": true,
              "cust_validation_url": "https://api.addsale.com/gringotts/api/v1/validate-customer/",
              "sdk": {
                "config": {
                  "redirect": false,
                  "callback_url": null,
                  "action_url": "https://api.addsale.com/avis/api/v1/payments/charge-gringotts-transaction/"
                },
                "data": {
                  "user_phone": "8452996729",
                  "user_email": "paymentsdummy@gofynd.com"
                }
              },
              "return_url": null
            }
          },
          "api_link": "",
          "payment_flow": "sdk"
        },
        "Juspay": {
          "data": {},
          "api_link": "https://sandbox.juspay.in/txns",
          "payment_flow": "api"
        },
        "Razorpay": {
          "data": {},
          "api_link": "",
          "payment_flow": "sdk"
        },
        "UPI_Razorpay": {
          "data": {},
          "api_link": "https://api.addsale.com/gringotts/api/v1/external/payment-initialisation/",
          "payment_flow": "api"
        },
        "Fynd": {
          "data": {},
          "api_link": "",
          "payment_flow": "api"
        }
      },
      "default": {}
    },
    "user_type": "Store User",
    "cod_charges": 0,
    "order_id": null,
    "cod_available": true,
    "cod_message": "No additional COD charges applicable",
    "delivery_charges": 0,
    "delivery_charge_order_value": 0,
    "delivery_slots": [
      {
        "date": "Sat, 24 Aug",
        "delivery_slot": [
          {
            "delivery_slot_timing": "By 9:00 PM",
            "default": true,
            "delivery_slot_id": 1
          }
        ]
      }
    ],
    "store_code": "",
    "store_emps": [],
    "breakup_values": {
      "loyalty_points": {
        "total": 0,
        "applicable": 0,
        "is_applied": false,
        "description": "Your cashback, referrals, and refund amount get credited to Fynd Cash which can be redeemed while placing an order."
      },
      "coupon": {
        "type": "cash",
        "code": "",
        "uid": null,
        "value": 0,
        "is_applied": false,
        "message": "Sorry! Invalid Coupon"
      },
      "raw": {
        "cod_charge": 0,
        "convenience_fee": 0,
        "coupon": 0,
        "delivery_charge": 0,
        "discount": 0,
        "fynd_cash": 0,
        "gst_charges": 214.18,
        "mrp_total": 1999,
        "subtotal": 1999,
        "total": 1999,
        "vog": 1784.82,
        "you_saved": 0
      },
      "display": [
        {
          "display": "MRP Total",
          "key": "mrp_total",
          "value": 1999,
          "currency_code": "INR"
        },
        {
          "display": "Subtotal",
          "key": "subtotal",
          "value": 1999,
          "currency_code": "INR"
        },
        {
          "display": "Total",
          "key": "total",
          "value": 1999,
          "currency_code": "INR"
        }
      ]
    },
    "shipments": [
      {
        "fulfillment_id": 3009,
        "shipment_type": "single_shipment",
        "fulfillment_type": "store",
        "dp_id": "29",
        "dp_options": {
          "4": {
            "f_priority": 4,
            "r_priority": 5,
            "is_cod": true,
            "is_prepaid": true,
            "is_reverse": true
          },
          "7": {
            "f_priority": 3,
            "r_priority": 4,
            "is_cod": true,
            "is_prepaid": true,
            "is_reverse": true
          },
          "29": {
            "f_priority": 1,
            "r_priority": 2,
            "is_cod": true,
            "is_prepaid": true,
            "is_reverse": true
          }
        },
        "promise": {
          "timestamp": {
            "min": 1566678108,
            "max": 1567023708
          },
          "formatted": {
            "min": "Aug 24",
            "max": "Aug 28"
          }
        },
        "box_type": "Small Courier bag",
        "shipments": 1,
        "items": [
          {
            "quantity": 1,
            "product": {
              "type": "product",
              "uid": 820312,
              "name": "Navy Blue Melange Shorts",
              "slug": "883-police-navy-blue-melange-shorts-820312-4943a8",
              "brand": {
                "uid": 610,
                "name": "883 Police"
              },
              "categories": [
                {
                  "uid": 193,
                  "name": "Shorts"
                }
              ],
              "images": [
                {
                  "aspect_ratio": "16:25",
                  "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/610_SPIRAL19ANAVY/1_1549105947281.jpg",
                  "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/610_SPIRAL19ANAVY/1_1549105947281.jpg"
                }
              ],
              "action": {
                "type": "product",
                "url": "https://api.addsale.com/platform/content/v1/products/883-police-navy-blue-melange-shorts-820312-4943a8/",
                "query": {
                  "product_slug": [
                    "883-police-navy-blue-melange-shorts-820312-4943a8"
                  ]
                }
              }
            },
            "discount": "",
            "bulk_offer": {},
            "key": "820312_L",
            "price": {
              "base": {
                "add_on": 1999,
                "marked": 1999,
                "effective": 1999,
                "selling": 1999,
                "currency_code": "INR"
              },
              "converted": {
                "add_on": 1999,
                "marked": 1999,
                "effective": 1999,
                "selling": 1999,
                "currency_code": "INR"
              }
            },
            "article": {
              "type": "article",
              "uid": "381_610_IGPL01_SPIRAL19ANAVY_L",
              "size": "L",
              "seller": {
                "uid": 381,
                "name": "INTERSOURCE GARMENTS PVT LTD"
              },
              "store": {
                "uid": 3009,
                "name": "Kormangala"
              },
              "quantity": 2,
              "price": {
                "base": {
                  "marked": 1999,
                  "effective": 1999,
                  "currency_code": "INR"
                },
                "converted": {
                  "marked": 1999,
                  "effective": 1999,
                  "currency_code": "INR"
                }
              }
            },
            "availability": {
              "sizes": [
                "L",
                "XL",
                "XXL"
              ],
              "other_store_quantity": 1,
              "out_of_stock": false,
              "deliverable": true,
              "is_valid": true
            },
            "coupon_message": "",
            "message": ""
          }
        ]
      }
    ],
    "delivery_charge_info": "",
    "coupon_text": "View all offers",
    "gstin": null,
    "checkout_mode": "self",
    "last_modified": "Thu, 22 Aug 2019 20:21:48 GMT",
    "restrict_checkout": false,
    "is_valid": true
  }
}
```

Shipment Generation Failed
```json
{
  "value": {
    "items": [],
    "cart_id": 7501,
    "uid": "7501",
    "success": true,
    "error_message": "Note: Your order delivery will be delayed by 7-10 Days",
    "payment_options": {
      "payment_option": [
        {
          "name": "COD",
          "display_name": "Cash on Delivery",
          "display_priority": 1,
          "payment_mode_id": 11,
          "logo": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png",
          "logo_url": {
            "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png",
            "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png"
          },
          "list": []
        },
        {
          "name": "CARD",
          "display_priority": 2,
          "payment_mode_id": 2,
          "display_name": "Card",
          "list": []
        },
        {
          "name": "NB",
          "display_priority": 3,
          "payment_mode_id": 3,
          "display_name": "Net Banking",
          "list": [
            {
              "aggregator_name": "Razorpay",
              "bank_name": "ICICI Bank",
              "bank_code": "ICIC",
              "url": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png"
              },
              "merchant_code": "NB_ICICI",
              "display_priority": 1
            }
          ]
        },
        {
          "name": "WL",
          "display_priority": 4,
          "payment_mode_id": 4,
          "display_name": "Wallet",
          "list": [
            {
              "wallet_name": "Paytm",
              "wallet_code": "paytm",
              "wallet_id": 4,
              "merchant_code": "PAYTM",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/paytm_logo_small.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/paytm_logo_large.png"
              },
              "aggregator_name": "Juspay",
              "display_priority": 1
            }
          ]
        },
        {
          "name": "UPI",
          "display_priority": 9,
          "payment_mode_id": 6,
          "display_name": "UPI",
          "list": [
            {
              "aggregator_name": "UPI_Razorpay",
              "name": "UPI",
              "display_name": "BHIM UPI",
              "code": "UPI",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/upi_100x78.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/upi_150x100.png"
              },
              "merchant_code": "UPI",
              "timeout": 240,
              "retry_count": 0,
              "fynd_vpa": "shopsense.rzp@hdfcbank",
              "intent_flow": true,
              "intent_app_error_list": [
                "com.csam.icici.bank.imobile",
                "in.org.npci.upiapp",
                "com.whatsapp"
              ]
            }
          ]
        },
        {
          "name": "PL",
          "display_priority": 11,
          "payment_mode_id": 1,
          "display_name": "Pay Later",
          "list": [
            {
              "aggregator_name": "Simpl",
              "name": "Simpl",
              "code": "simpl",
              "merchant_code": "SIMPL",
              "logo": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png"
              }
            }
          ]
        }
      ],
      "payment_flows": {
        "Simpl": {
          "data": {
            "gateway": {
              "route": "simpl",
              "entity": "sdk",
              "is_customer_validation_required": true,
              "cust_validation_url": "https://api.addsale.com/gringotts/api/v1/validate-customer/",
              "sdk": {
                "config": {
                  "redirect": false,
                  "callback_url": null,
                  "action_url": "https://api.addsale.com/avis/api/v1/payments/charge-gringotts-transaction/"
                },
                "data": {
                  "user_phone": "8452996729",
                  "user_email": "paymentsdummy@gofynd.com"
                }
              },
              "return_url": null
            }
          },
          "api_link": "",
          "payment_flow": "sdk"
        },
        "Juspay": {
          "data": {},
          "api_link": "https://sandbox.juspay.in/txns",
          "payment_flow": "api"
        },
        "Razorpay": {
          "data": {},
          "api_link": "",
          "payment_flow": "sdk"
        },
        "UPI_Razorpay": {
          "data": {},
          "api_link": "https://api.addsale.com/gringotts/api/v1/external/payment-initialisation/",
          "payment_flow": "api"
        },
        "Fynd": {
          "data": {},
          "api_link": "",
          "payment_flow": "api"
        }
      },
      "default": {}
    },
    "user_type": "Store User",
    "cod_charges": 0,
    "order_id": null,
    "cod_available": true,
    "cod_message": "No additional COD charges applicable",
    "delivery_charges": 0,
    "delivery_charge_order_value": 0,
    "delivery_slots": [
      {
        "date": "Sat, 24 Aug",
        "delivery_slot": [
          {
            "delivery_slot_timing": "By 9:00 PM",
            "default": true,
            "delivery_slot_id": 1
          }
        ]
      }
    ],
    "store_code": "",
    "store_emps": [],
    "breakup_values": {
      "loyalty_points": {
        "total": 0,
        "applicable": 0,
        "is_applied": false,
        "description": "Your cashback, referrals, and refund amount get credited to Fynd Cash which can be redeemed while placing an order."
      },
      "coupon": {
        "type": "cash",
        "code": "",
        "uid": null,
        "value": 0,
        "is_applied": false,
        "message": "Sorry! Invalid Coupon"
      },
      "raw": {
        "cod_charge": 0,
        "convenience_fee": 0,
        "coupon": 0,
        "delivery_charge": 0,
        "discount": 0,
        "fynd_cash": 0,
        "gst_charges": 214.18,
        "mrp_total": 1999,
        "subtotal": 1999,
        "total": 1999,
        "vog": 1784.82,
        "you_saved": 0
      },
      "display": [
        {
          "display": "MRP Total",
          "key": "mrp_total",
          "value": 1999,
          "currency_code": "INR"
        },
        {
          "display": "Subtotal",
          "key": "subtotal",
          "value": 1999,
          "currency_code": "INR"
        },
        {
          "display": "Total",
          "key": "total",
          "value": 1999,
          "currency_code": "INR"
        }
      ]
    },
    "shipments": [],
    "message": "Shipments could not be generated. Please Try again after some time.",
    "delivery_charge_info": "",
    "coupon_text": "View all offers",
    "gstin": null,
    "checkout_mode": "self",
    "last_modified": "Thu, 22 Aug 2019 20:21:48 GMT",
    "restrict_checkout": false,
    "is_valid": false
  }
}
```








Unhandled api error


Schema: `Object`









---


#### checkoutCart
Checkout Cart

```golang

 data, err :=  Cart.CheckoutCart(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  CartCheckoutRequest | "Request body" 


Checkout all items in cart to payment and order generation.                         For COD only order will be generated while for other checkout mode                         user will be redirected to payment gateway

*Success Response:*



OK


Schema: `CartCheckoutResponse`


*Examples:*


Address id not found
```json
{
  "value": {
    "success": false,
    "message": "No address found with address id {address_id}"
  }
}
```

Missing address_id
```json
{
  "value": {
    "address_id": [
      "Missing data for required field."
    ]
  }
}
```

Successful checkout cod payment
```json
{
  "value": {
    "success": true,
    "cart": {
      "success": true,
      "error_message": "Note: Your order delivery will be delayed by 7-10 Days",
      "payment_options": {
        "payment_option": [
          {
            "name": "COD",
            "display_name": "Cash on Delivery",
            "display_priority": 1,
            "payment_mode_id": 11,
            "logo": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png",
            "logo_url": {
              "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png",
              "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png"
            },
            "list": []
          },
          {
            "name": "CARD",
            "display_priority": 2,
            "payment_mode_id": 2,
            "display_name": "Card",
            "list": []
          },
          {
            "name": "NB",
            "display_priority": 3,
            "payment_mode_id": 3,
            "display_name": "Net Banking",
            "list": [
              {
                "aggregator_name": "Razorpay",
                "bank_name": "ICICI Bank",
                "bank_code": "ICIC",
                "url": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png",
                "logo_url": {
                  "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png",
                  "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png"
                },
                "merchant_code": "NB_ICICI",
                "display_priority": 1
              }
            ]
          },
          {
            "name": "WL",
            "display_priority": 4,
            "payment_mode_id": 4,
            "display_name": "Wallet",
            "list": [
              {
                "wallet_name": "Paytm",
                "wallet_code": "paytm",
                "wallet_id": 4,
                "merchant_code": "PAYTM",
                "logo_url": {
                  "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/paytm_logo_small.png",
                  "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/paytm_logo_large.png"
                },
                "aggregator_name": "Juspay",
                "display_priority": 1
              }
            ]
          },
          {
            "name": "UPI",
            "display_priority": 9,
            "payment_mode_id": 6,
            "display_name": "UPI",
            "list": [
              {
                "aggregator_name": "UPI_Razorpay",
                "name": "UPI",
                "display_name": "BHIM UPI",
                "code": "UPI",
                "logo_url": {
                  "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/upi_100x78.png",
                  "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/upi_150x100.png"
                },
                "merchant_code": "UPI",
                "timeout": 240,
                "retry_count": 0,
                "fynd_vpa": "shopsense.rzp@hdfcbank",
                "intent_flow": true,
                "intent_app_error_list": [
                  "com.csam.icici.bank.imobile",
                  "in.org.npci.upiapp",
                  "com.whatsapp"
                ]
              }
            ]
          },
          {
            "name": "PL",
            "display_priority": 11,
            "payment_mode_id": 1,
            "display_name": "Pay Later",
            "list": [
              {
                "aggregator_name": "Simpl",
                "name": "Simpl",
                "code": "simpl",
                "merchant_code": "SIMPL",
                "logo": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png",
                "logo_url": {
                  "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png",
                  "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png"
                }
              }
            ]
          }
        ],
        "payment_flows": {
          "Simpl": {
            "data": {
              "gateway": {
                "route": "simpl",
                "entity": "sdk",
                "is_customer_validation_required": true,
                "cust_validation_url": "https://api.addsale.com/gringotts/api/v1/validate-customer/",
                "sdk": {
                  "config": {
                    "redirect": false,
                    "callback_url": null,
                    "action_url": "https://api.addsale.com/avis/api/v1/payments/charge-gringotts-transaction/"
                  },
                  "data": {
                    "user_phone": "8452996729",
                    "user_email": "paymentsdummy@gofynd.com"
                  }
                },
                "return_url": null
              }
            },
            "api_link": "",
            "payment_flow": "sdk"
          },
          "Juspay": {
            "data": {},
            "api_link": "https://sandbox.juspay.in/txns",
            "payment_flow": "api"
          },
          "Razorpay": {
            "data": {},
            "api_link": "",
            "payment_flow": "sdk"
          },
          "UPI_Razorpay": {
            "data": {},
            "api_link": "https://api.addsale.com/gringotts/api/v1/external/payment-initialisation/",
            "payment_flow": "api"
          },
          "Fynd": {
            "data": {},
            "api_link": "",
            "payment_flow": "api"
          }
        },
        "default": {}
      },
      "user_type": "Store User",
      "cod_charges": 0,
      "order_id": "FY5D5E215CF287584CE6",
      "cod_available": true,
      "cod_message": "No additional COD charges applicable",
      "delivery_charges": 0,
      "delivery_charge_order_value": 0,
      "delivery_slots": [
        {
          "date": "Sat, 24 Aug",
          "delivery_slot": [
            {
              "delivery_slot_timing": "By 9:00 PM",
              "default": true,
              "delivery_slot_id": 1
            }
          ]
        }
      ],
      "store_code": "",
      "store_emps": [],
      "breakup_values": {
        "coupon": {
          "type": "cash",
          "code": "",
          "uid": null,
          "value": 0,
          "is_applied": false,
          "message": "Sorry! Invalid Coupon"
        },
        "loyalty_points": {
          "total": 0,
          "applicable": 0,
          "is_applied": false,
          "description": "Your cashback, referrals, and refund amount get credited to Fynd Cash which can be redeemed while placing an order."
        },
        "raw": {
          "cod_charge": 0,
          "convenience_fee": 0,
          "coupon": 0,
          "delivery_charge": 0,
          "discount": 0,
          "fynd_cash": 0,
          "gst_charges": 214.18,
          "mrp_total": 1999,
          "subtotal": 1999,
          "total": 1999,
          "vog": 1784.82,
          "you_saved": 0
        },
        "display": [
          {
            "display": "MRP Total",
            "key": "mrp_total",
            "value": 1999,
            "currency_code": "INR"
          },
          {
            "display": "Subtotal",
            "key": "subtotal",
            "value": 1999,
            "currency_code": "INR"
          },
          {
            "display": "Total",
            "key": "total",
            "value": 1999,
            "currency_code": "INR"
          }
        ]
      },
      "items": [
        {
          "key": "820312_L",
          "message": "",
          "bulk_offer": {},
          "price": {
            "base": {
              "add_on": 1999,
              "marked": 1999,
              "effective": 1999,
              "selling": 1999,
              "currency_code": "INR"
            },
            "converted": {
              "add_on": 1999,
              "marked": 1999,
              "effective": 1999,
              "selling": 1999,
              "currency_code": "INR"
            }
          },
          "quantity": 1,
          "discount": "",
          "product": {
            "type": "product",
            "uid": 820312,
            "name": "Navy Blue Melange Shorts",
            "slug": "883-police-navy-blue-melange-shorts-820312-4943a8",
            "brand": {
              "uid": 610,
              "name": "883 Police"
            },
            "categories": [
              {
                "uid": 193,
                "name": "Shorts"
              }
            ],
            "images": [
              {
                "aspect_ratio": "16:25",
                "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/610_SPIRAL19ANAVY/1_1549105947281.jpg",
                "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/610_SPIRAL19ANAVY/1_1549105947281.jpg"
              }
            ],
            "action": {
              "type": "product",
              "url": "https://api.addsale.com/platform/content/v1/products/883-police-navy-blue-melange-shorts-820312-4943a8/",
              "query": {
                "product_slug": [
                  "883-police-navy-blue-melange-shorts-820312-4943a8"
                ]
              }
            }
          },
          "article": {
            "type": "article",
            "uid": "381_610_IGPL01_SPIRAL19ANAVY_L",
            "size": "L",
            "seller": {
              "uid": 381,
              "name": "INTERSOURCE GARMENTS PVT LTD"
            },
            "store": {
              "uid": 3009,
              "name": "Kormangala"
            },
            "quantity": 2,
            "price": {
              "base": {
                "marked": 1999,
                "effective": 1999,
                "currency_code": "INR"
              },
              "converted": {
                "marked": 1999,
                "effective": 1999,
                "currency_code": "INR"
              }
            }
          },
          "coupon_message": "",
          "availability": {
            "sizes": [
              "L",
              "XL",
              "XXL"
            ],
            "other_store_quantity": 1,
            "out_of_stock": false,
            "deliverable": true,
            "is_valid": true
          }
        }
      ],
      "delivery_charge_info": "",
      "coupon_text": "View all offers",
      "cart_id": 7483,
      "uid": "7483",
      "gstin": null,
      "checkout_mode": "self",
      "last_modified": "Thu, 22 Aug 2019 04:58:44 GMT",
      "restrict_checkout": false,
      "is_valid": true
    },
    "callback_url": "https://api.addsale.com/gringotts/api/v1/external/payment-callback/",
    "app_intercept_url": "http://uniket-testing.addsale.link/cart/order-status",
    "message": "",
    "data": {
      "order_id": "FY5D5E215CF287584CE6"
    },
    "order_id": "FY5D5E215CF287584CE6"
  }
}
```









---


#### updateCartMeta
Update Cart Meta

```golang

 data, err :=  Cart.UpdateCartMeta(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `UID`

| body |  CartMetaRequest | "Request body" 


Update cart meta like checkout_mode, gstin.

*Success Response:*



Cart meta updated successfully


Schema: `CartMetaResponse`








Missing required Field


Schema: `CartMetaMissingResponse`









---


#### getCartShareLink
Generate Cart sharing link token

```golang

 data, err :=  Cart.GetCartShareLink(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  GetShareCartLinkRequest | "Request body" 


Generates shared cart snapshot and returns shortlink token

*Success Response:*



Token Generated successfully


Schema: `GetShareCartLinkResponse`


*Examples:*


Token Generated
```json
{
  "value": {
    "token": "ZweG1XyX",
    "share_url": "https://uniket-testing.addsale.link/shared-cart/ZweG1XyX"
  }
}
```









---


#### getCartSharedItems
Get shared cart snapshot and cart response

```golang

 data, err :=  Cart.GetCartSharedItems(Token);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Token | string | Shared short link token. | 




Returns shared cart response for sent token with `shared_cart_details`                    containing shared cart details in response.

*Success Response:*



Cart for valid token


Schema: `SharedCartResponse`








No cart found for sent token


Schema: `SharedCartResponse`









---


#### updateCartWithSharedItems
Merge or Replace existing cart

```golang

 data, err :=  Cart.UpdateCartWithSharedItems(Token, Action);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Token | string | Shared short link token. | 


| Action | string | Operation to perform on existing cart, whether to merge or replace. | 




Merge or Replace cart based on `action` parameter with shared cart of `token`

*Success Response:*



Success of Merge or Replace of cart with `shared_cart_details`                    containing shared cart details in response


Schema: `SharedCartResponse`


*Examples:*


Cart Merged/Replaced
```json
{
  "value": {
    "cart": {
      "shared_cart_details": {
        "token": "BQ9jySQ9",
        "user": {
          "user_id": "23109086",
          "is_anonymous": false
        },
        "meta": {
          "selected_staff": "",
          "ordering_store": null
        },
        "selected_staff": "",
        "ordering_store": null,
        "source": {},
        "created_on": "2019-12-18T14:00:07.165000"
      },
      "items": [
        {
          "key": "791651_6",
          "discount": "",
          "bulk_offer": {},
          "coupon_message": "",
          "article": {
            "type": "article",
            "uid": "304_1054_9036_R1005753_6",
            "size": "6",
            "seller": {
              "uid": 304,
              "name": "LEAYAN GLOBAL PVT. LTD."
            },
            "store": {
              "uid": 5322,
              "name": "Vaisali Nagar"
            },
            "quantity": 1,
            "price": {
              "base": {
                "marked": 2095,
                "effective": 2095,
                "currency_code": "INR"
              },
              "converted": {
                "marked": 2095,
                "effective": 2095,
                "currency_code": "INR"
              }
            }
          },
          "product": {
            "type": "product",
            "uid": 791651,
            "name": "Black Running Shoes",
            "slug": "furo-black-running-shoes-791651-f8bcc3",
            "brand": {
              "uid": 1054,
              "name": "Furo"
            },
            "categories": [
              {
                "uid": 160,
                "name": "Running Shoes"
              }
            ],
            "images": [
              {
                "aspect_ratio": "16:25",
                "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/1054_R1005753/1_1546490507364.jpg",
                "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/1054_R1005753/1_1546490507364.jpg"
              }
            ],
            "action": {
              "type": "product",
              "url": "https://api.addsale.com/platform/content/v1/products/furo-black-running-shoes-791651-f8bcc3/",
              "query": {
                "product_slug": [
                  "furo-black-running-shoes-791651-f8bcc3"
                ]
              }
            }
          },
          "message": "",
          "quantity": 1,
          "availability": {
            "sizes": [
              "7",
              "8",
              "9",
              "10",
              "6"
            ],
            "other_store_quantity": 12,
            "out_of_stock": false,
            "deliverable": true,
            "is_valid": true
          },
          "price": {
            "base": {
              "add_on": 2095,
              "marked": 2095,
              "effective": 2095,
              "selling": 2095,
              "currency_code": "INR"
            },
            "converted": {
              "add_on": 2095,
              "marked": 2095,
              "effective": 2095,
              "selling": 2095,
              "currency_code": "INR"
            }
          }
        },
        {
          "key": "791651_7",
          "discount": "",
          "bulk_offer": {},
          "coupon_message": "",
          "article": {
            "type": "article",
            "uid": "304_1054_9036_R1005753_7",
            "size": "7",
            "seller": {
              "uid": 304,
              "name": "LEAYAN GLOBAL PVT. LTD."
            },
            "store": {
              "uid": 5322,
              "name": "Vaisali Nagar"
            },
            "quantity": 2,
            "price": {
              "base": {
                "marked": 2095,
                "effective": 2095,
                "currency_code": "INR"
              },
              "converted": {
                "marked": 2095,
                "effective": 2095,
                "currency_code": "INR"
              }
            }
          },
          "product": {
            "type": "product",
            "uid": 791651,
            "name": "Black Running Shoes",
            "slug": "furo-black-running-shoes-791651-f8bcc3",
            "brand": {
              "uid": 1054,
              "name": "Furo"
            },
            "categories": [
              {
                "uid": 160,
                "name": "Running Shoes"
              }
            ],
            "images": [
              {
                "aspect_ratio": "16:25",
                "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/1054_R1005753/1_1546490507364.jpg",
                "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/1054_R1005753/1_1546490507364.jpg"
              }
            ],
            "action": {
              "type": "product",
              "url": "https://api.addsale.com/platform/content/v1/products/furo-black-running-shoes-791651-f8bcc3/",
              "query": {
                "product_slug": [
                  "furo-black-running-shoes-791651-f8bcc3"
                ]
              }
            }
          },
          "message": "",
          "quantity": 2,
          "availability": {
            "sizes": [
              "7",
              "8",
              "9",
              "10",
              "6"
            ],
            "other_store_quantity": 7,
            "out_of_stock": false,
            "deliverable": true,
            "is_valid": true
          },
          "price": {
            "base": {
              "add_on": 4190,
              "marked": 4190,
              "effective": 4190,
              "selling": 4190,
              "currency_code": "INR"
            },
            "converted": {
              "add_on": 4190,
              "marked": 4190,
              "effective": 4190,
              "selling": 4190,
              "currency_code": "INR"
            }
          }
        }
      ],
      "cart_id": 13055,
      "uid": "13055",
      "breakup_values": {
        "raw": {
          "cod_charge": 0,
          "convenience_fee": 0,
          "coupon": 0,
          "delivery_charge": 0,
          "discount": 0,
          "fynd_cash": 0,
          "gst_charges": 958.73,
          "mrp_total": 6285,
          "subtotal": 6285,
          "total": 6285,
          "vog": 5326.27,
          "you_saved": 0
        },
        "loyalty_points": {
          "total": 0,
          "applicable": 0,
          "is_applied": false,
          "description": "Your cashback, referrals, and refund amount get credited to Fynd Cash which can be redeemed while placing an order."
        },
        "coupon": {
          "type": "cash",
          "code": "",
          "uid": null,
          "value": 0,
          "is_applied": false,
          "message": "Sorry! Invalid coupon"
        },
        "display": [
          {
            "display": "MRP Total",
            "key": "mrp_total",
            "value": 6285,
            "currency_code": "INR"
          },
          {
            "display": "Subtotal",
            "key": "subtotal",
            "value": 6285,
            "currency_code": "INR"
          },
          {
            "display": "Total",
            "key": "total",
            "value": 6285,
            "currency_code": "INR"
          }
        ]
      },
      "delivery_charge_info": "",
      "coupon_text": "View all offers",
      "gstin": null,
      "comment": "",
      "checkout_mode": "self",
      "payment_selection_lock": {
        "enabled": false,
        "default_options": "COD",
        "payment_identifier": null
      },
      "restrict_checkout": false,
      "is_valid": true,
      "last_modified": "Mon, 16 Dec 2019 07:02:18 GMT"
    }
  }
}
```









---



---


## Lead


#### getTicket
Get Ticket with the specific id

```golang

 data, err :=  Lead.GetTicket(ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ID | string | ID of ticket to be retrieved | 




Get Ticket with the specific id, this is used to view the ticket details

*Success Response:*



Success


Schema: `Ticket`


*Examples:*


Default
```json
{
  "value": {
    "context": {
      "application_id": "000000000000000000000003",
      "company_id": "884"
    },
    "content": {
      "title": "SOme title Response",
      "description": "<b>Single lineeee</b>: asdf<br><b>Email</b>: asdf@asdf.com<br><b>dfsdf</b>: asdf<br>",
      "attachments": []
    },
    "status": {
      "display": "In Progress",
      "color": "#ffa951",
      "key": "in_progress"
    },
    "priority": {
      "display": "Medium",
      "color": "#f37736",
      "key": "medium"
    },
    "assigned_to": {
      "agent_id": "5d1363adf599d850df93175e",
      "gender": "male",
      "accountType": "user",
      "active": true,
      "profilePicUrl": "https://platform-lookaside.fbsbx.com/platform/profilepic/?asid=2136700473091190&height=400&width=400&ext=1554542761&hash=AeS6cuWIdjDdJJ-b",
      "hasOldPasswordHash": false,
      "_id": "5d1363adf599d850df93175e",
      "phoneNumbers": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "countryCode": 91,
          "phone": "9890425946"
        }
      ],
      "firstName": "Nikhil",
      "lastName": "Manapure",
      "emails": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "email": "nikhilmshchs@gmail.com"
        }
      ],
      "username": "nikhilmshchs_gmail_com_38425_20500281",
      "createdAt": "2019-01-01T17:22:38.528Z",
      "updatedAt": "2021-01-22T10:02:42.258Z",
      "uid": "20500281",
      "__v": 56
    },
    "tags": [
      "some-title"
    ],
    "_id": "6012f38557751ee8fc162cf7",
    "created_on": {
      "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36",
      "platform": "web",
      "meta": {
        "browser": {
          "name": "Chrome",
          "version": "88.0.4324.96"
        },
        "os": {
          "name": "macOS",
          "version": "10.15.7",
          "versionName": "Catalina"
        },
        "platform": {
          "type": "desktop",
          "vendor": "Apple"
        },
        "engine": {
          "name": "Blink"
        }
      }
    },
    "source": "sales_channel",
    "created_by": {
      "id": "5d1363adf599d850df93175e",
      "user": {
        "gender": "male",
        "accountType": "user",
        "active": true,
        "profilePicUrl": "https://platform-lookaside.fbsbx.com/platform/profilepic/?asid=2136700473091190&height=400&width=400&ext=1554542761&hash=AeS6cuWIdjDdJJ-b",
        "hasOldPasswordHash": false,
        "_id": "5d1363adf599d850df93175e",
        "phoneNumbers": [
          {
            "active": true,
            "primary": true,
            "verified": true,
            "countryCode": 91,
            "phone": "9890425946"
          }
        ],
        "firstName": "Nikhil",
        "lastName": "Manapure",
        "emails": [
          {
            "active": true,
            "primary": true,
            "verified": true,
            "email": "nikhilmshchs@gmail.com"
          }
        ],
        "username": "nikhilmshchs_gmail_com_38425_20500281",
        "createdAt": "2019-01-01T17:22:38.528Z",
        "updatedAt": "2021-01-22T10:02:42.258Z",
        "uid": "20500281",
        "__v": 56
      }
    },
    "response_id": "6012f38457751e0fb8162cf6",
    "category": {
      "form": {
        "login_required": false,
        "should_notify": false,
        "inputs": [
          {
            "required": false,
            "type": "text",
            "enum": [],
            "display": "Single lineeee",
            "key": "single-lineeee",
            "showRegexInput": false
          },
          {
            "required": false,
            "type": "email",
            "enum": [],
            "display": "Email",
            "regex": "\\S+@\\S+\\.\\S+",
            "key": "email",
            "showRegexInput": true
          },
          {
            "required": false,
            "type": "text",
            "enum": [],
            "display": "dfsdf",
            "key": "dfsdf",
            "showRegexInput": false
          }
        ],
        "available_assignees": [
          "5b9b98150df588546aaea6d2",
          "5c45d78395d7504f76c2cb37"
        ],
        "_id": "5fd72db3dc250f8decfc61b2",
        "title": "SOme title",
        "description": "SOme big description",
        "slug": "some-title",
        "application_id": "000000000000000000000003",
        "created_on": {
          "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36",
          "platform": "web",
          "meta": {
            "browser": {
              "name": "Chrome",
              "version": "87.0.4280.88"
            },
            "os": {
              "name": "macOS",
              "version": "10.15.6",
              "versionName": "Catalina"
            },
            "platform": {
              "type": "desktop",
              "vendor": "Apple"
            },
            "engine": {
              "name": "Blink"
            }
          }
        },
        "created_by": "5d1363adf599d850df93175e",
        "createdAt": "2020-12-14T09:17:39.953Z",
        "updatedAt": "2021-01-28T18:48:07.717Z",
        "__v": 0
      },
      "key": "some-title",
      "display": "SOme title"
    },
    "ticket_id": "43",
    "createdAt": "2021-01-28T17:25:25.013Z",
    "updatedAt": "2021-01-28T17:25:33.396Z",
    "__v": 0,
    "video_room_id": "6012f38557751ee8fc162cf7"
  }
}
```









---


#### createHistory
Create history for specific Ticket

```golang

 data, err :=  Lead.CreateHistory(ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ID | string | Ticket ID for which history is created | 


| body |  TicketHistoryPayload | "Request body" 


Create history for specific Ticket, this history is seen on ticket detail page, this can be comment, log or rating.

*Success Response:*



Success


Schema: `TicketHistory`


*Examples:*


Default
```json
{
  "value": {
    "_id": "601a9d52c26687d086c499ef",
    "ticket_id": "41",
    "type": "comment",
    "value": {
      "text": "d",
      "media": []
    },
    "created_on": {
      "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36",
      "platform": "web",
      "meta": {
        "browser": {
          "name": "Chrome",
          "version": "88.0.4324.96"
        },
        "os": {
          "name": "macOS",
          "version": "10.15.7",
          "versionName": "Catalina"
        },
        "platform": {
          "type": "desktop",
          "vendor": "Apple"
        },
        "engine": {
          "name": "Blink"
        }
      }
    },
    "created_by": "5d1363adf599d850df93175e",
    "createdAt": "2021-02-03T12:55:46.808Z",
    "updatedAt": "2021-02-03T12:55:46.808Z",
    "__v": 0
  }
}
```









---


#### createTicket
Create Ticket

```golang

 data, err :=  Lead.CreateTicket(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  AddTicketPayload | "Request body" 


This is used to Create Ticket.

*Success Response:*



Success


Schema: `Ticket`


*Examples:*


Default
```json
{
  "value": {
    "context": {
      "application_id": "000000000000000000000003",
      "company_id": "884"
    },
    "content": {
      "title": "SOme title Response",
      "description": "<b>Single lineeee</b>: asdf<br><b>Email</b>: asdf@asdf.com<br><b>dfsdf</b>: asdf<br>",
      "attachments": []
    },
    "status": {
      "display": "In Progress",
      "color": "#ffa951",
      "key": "in_progress"
    },
    "priority": {
      "display": "Medium",
      "color": "#f37736",
      "key": "medium"
    },
    "assigned_to": {
      "agent_id": "5d1363adf599d850df93175e",
      "gender": "male",
      "accountType": "user",
      "active": true,
      "profilePicUrl": "https://platform-lookaside.fbsbx.com/platform/profilepic/?asid=2136700473091190&height=400&width=400&ext=1554542761&hash=AeS6cuWIdjDdJJ-b",
      "hasOldPasswordHash": false,
      "_id": "5d1363adf599d850df93175e",
      "phoneNumbers": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "countryCode": 91,
          "phone": "9890425946"
        }
      ],
      "firstName": "Nikhil",
      "lastName": "Manapure",
      "emails": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "email": "nikhilmshchs@gmail.com"
        }
      ],
      "username": "nikhilmshchs_gmail_com_38425_20500281",
      "createdAt": "2019-01-01T17:22:38.528Z",
      "updatedAt": "2021-01-22T10:02:42.258Z",
      "uid": "20500281",
      "__v": 56
    },
    "tags": [
      "some-title"
    ],
    "_id": "6012f38557751ee8fc162cf7",
    "created_on": {
      "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36",
      "platform": "web",
      "meta": {
        "browser": {
          "name": "Chrome",
          "version": "88.0.4324.96"
        },
        "os": {
          "name": "macOS",
          "version": "10.15.7",
          "versionName": "Catalina"
        },
        "platform": {
          "type": "desktop",
          "vendor": "Apple"
        },
        "engine": {
          "name": "Blink"
        }
      }
    },
    "source": "sales_channel",
    "created_by": {
      "id": "5d1363adf599d850df93175e",
      "user": {
        "gender": "male",
        "accountType": "user",
        "active": true,
        "profilePicUrl": "https://platform-lookaside.fbsbx.com/platform/profilepic/?asid=2136700473091190&height=400&width=400&ext=1554542761&hash=AeS6cuWIdjDdJJ-b",
        "hasOldPasswordHash": false,
        "_id": "5d1363adf599d850df93175e",
        "phoneNumbers": [
          {
            "active": true,
            "primary": true,
            "verified": true,
            "countryCode": 91,
            "phone": "9890425946"
          }
        ],
        "firstName": "Nikhil",
        "lastName": "Manapure",
        "emails": [
          {
            "active": true,
            "primary": true,
            "verified": true,
            "email": "nikhilmshchs@gmail.com"
          }
        ],
        "username": "nikhilmshchs_gmail_com_38425_20500281",
        "createdAt": "2019-01-01T17:22:38.528Z",
        "updatedAt": "2021-01-22T10:02:42.258Z",
        "uid": "20500281",
        "__v": 56
      }
    },
    "response_id": "6012f38457751e0fb8162cf6",
    "category": {
      "form": {
        "login_required": false,
        "should_notify": false,
        "inputs": [
          {
            "required": false,
            "type": "text",
            "enum": [],
            "display": "Single lineeee",
            "key": "single-lineeee",
            "showRegexInput": false
          },
          {
            "required": false,
            "type": "email",
            "enum": [],
            "display": "Email",
            "regex": "\\S+@\\S+\\.\\S+",
            "key": "email",
            "showRegexInput": true
          },
          {
            "required": false,
            "type": "text",
            "enum": [],
            "display": "dfsdf",
            "key": "dfsdf",
            "showRegexInput": false
          }
        ],
        "available_assignees": [
          "5b9b98150df588546aaea6d2",
          "5c45d78395d7504f76c2cb37"
        ],
        "_id": "5fd72db3dc250f8decfc61b2",
        "title": "SOme title",
        "description": "SOme big description",
        "slug": "some-title",
        "application_id": "000000000000000000000003",
        "created_on": {
          "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36",
          "platform": "web",
          "meta": {
            "browser": {
              "name": "Chrome",
              "version": "87.0.4280.88"
            },
            "os": {
              "name": "macOS",
              "version": "10.15.6",
              "versionName": "Catalina"
            },
            "platform": {
              "type": "desktop",
              "vendor": "Apple"
            },
            "engine": {
              "name": "Blink"
            }
          }
        },
        "created_by": "5d1363adf599d850df93175e",
        "createdAt": "2020-12-14T09:17:39.953Z",
        "updatedAt": "2021-01-28T18:48:07.717Z",
        "__v": 0
      },
      "key": "some-title",
      "display": "SOme title"
    },
    "ticket_id": "43",
    "createdAt": "2021-01-28T17:25:25.013Z",
    "updatedAt": "2021-01-28T17:25:33.396Z",
    "__v": 0,
    "video_room_id": "6012f38557751ee8fc162cf7"
  }
}
```









---


#### getCustomForm
Get specific Custom Form using it's slug

```golang

 data, err :=  Lead.GetCustomForm(Slug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | Slug of form whose response is getting submitted | 




Get specific Custom Form using it's slug, this is used to view the form.

*Success Response:*



Success


Schema: `CustomForm`


*Examples:*


Default
```json
{
  "value": {
    "login_required": false,
    "should_notify": false,
    "inputs": [
      {
        "required": false,
        "type": "text",
        "display": "Name",
        "placeholder": "Please enter your name",
        "key": "name"
      }
    ],
    "available_assignees": [],
    "_id": "5fd258a9088f957f34c288fc",
    "title": "trail form",
    "description": "Trail form description",
    "slug": "trail-form",
    "application_id": "000000000000000000000003",
    "created_on": {
      "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36",
      "platform": "web",
      "meta": {
        "browser": {
          "name": "Chrome",
          "version": "87.0.4280.88"
        },
        "os": {
          "name": "macOS",
          "version": "10.15.6",
          "versionName": "Catalina"
        },
        "platform": {
          "type": "desktop",
          "vendor": "Apple"
        },
        "engine": {
          "name": "Blink"
        }
      }
    },
    "created_by": "5d1363adf599d850df93175e",
    "createdAt": "2020-12-10T17:19:37.515Z",
    "updatedAt": "2020-12-10T17:19:43.214Z",
    "__v": 0
  }
}
```









---


#### submitCustomForm
Submit Response for a specific Custom Form using it's slug

```golang

 data, err :=  Lead.SubmitCustomForm(Slug, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | Slug of form whose response is getting submitted | 


| body |  CustomFormSubmissionPayload | "Request body" 


Submit Response for a specific Custom Form using it's slug, this response is then used to create a ticket on behalf of the user.

*Success Response:*



Success


Schema: `SubmitCustomFormResponse`


*Examples:*


Default
```json
{
  "value": {
    "ticket": {
      "context": {
        "application_id": "000000000000000000000003",
        "company_id": "884"
      },
      "content": {
        "title": "SOme title Response",
        "description": "<b>Single lineeee</b>: asdf<br><b>Email</b>: asdf@asdf.com<br><b>dfsdf</b>: asdf<br>",
        "attachments": []
      },
      "status": {
        "display": "In Progress",
        "color": "#ffa951",
        "key": "in_progress"
      },
      "priority": {
        "display": "Medium",
        "color": "#f37736",
        "key": "medium"
      },
      "assigned_to": {
        "agent_id": "5d1363adf599d850df93175e",
        "gender": "male",
        "accountType": "user",
        "active": true,
        "profilePicUrl": "https://platform-lookaside.fbsbx.com/platform/profilepic/?asid=2136700473091190&height=400&width=400&ext=1554542761&hash=AeS6cuWIdjDdJJ-b",
        "hasOldPasswordHash": false,
        "_id": "5d1363adf599d850df93175e",
        "phoneNumbers": [
          {
            "active": true,
            "primary": true,
            "verified": true,
            "countryCode": 91,
            "phone": "9890425946"
          }
        ],
        "firstName": "Nikhil",
        "lastName": "Manapure",
        "emails": [
          {
            "active": true,
            "primary": true,
            "verified": true,
            "email": "nikhilmshchs@gmail.com"
          }
        ],
        "username": "nikhilmshchs_gmail_com_38425_20500281",
        "createdAt": "2019-01-01T17:22:38.528Z",
        "updatedAt": "2021-01-22T10:02:42.258Z",
        "uid": "20500281",
        "__v": 56
      },
      "tags": [
        "some-title"
      ],
      "_id": "6012f38557751ee8fc162cf7",
      "created_on": {
        "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36",
        "platform": "web",
        "meta": {
          "browser": {
            "name": "Chrome",
            "version": "88.0.4324.96"
          },
          "os": {
            "name": "macOS",
            "version": "10.15.7",
            "versionName": "Catalina"
          },
          "platform": {
            "type": "desktop",
            "vendor": "Apple"
          },
          "engine": {
            "name": "Blink"
          }
        }
      },
      "source": "sales_channel",
      "created_by": {
        "id": "5d1363adf599d850df93175e",
        "user": {
          "gender": "male",
          "accountType": "user",
          "active": true,
          "profilePicUrl": "https://platform-lookaside.fbsbx.com/platform/profilepic/?asid=2136700473091190&height=400&width=400&ext=1554542761&hash=AeS6cuWIdjDdJJ-b",
          "hasOldPasswordHash": false,
          "_id": "5d1363adf599d850df93175e",
          "phoneNumbers": [
            {
              "active": true,
              "primary": true,
              "verified": true,
              "countryCode": 91,
              "phone": "9890425946"
            }
          ],
          "firstName": "Nikhil",
          "lastName": "Manapure",
          "emails": [
            {
              "active": true,
              "primary": true,
              "verified": true,
              "email": "nikhilmshchs@gmail.com"
            }
          ],
          "username": "nikhilmshchs_gmail_com_38425_20500281",
          "createdAt": "2019-01-01T17:22:38.528Z",
          "updatedAt": "2021-01-22T10:02:42.258Z",
          "uid": "20500281",
          "__v": 56
        }
      },
      "response_id": "6012f38457751e0fb8162cf6",
      "category": {
        "form": {
          "login_required": false,
          "should_notify": false,
          "inputs": [
            {
              "required": false,
              "type": "text",
              "enum": [],
              "display": "Single lineeee",
              "key": "single-lineeee",
              "showRegexInput": false
            },
            {
              "required": false,
              "type": "email",
              "enum": [],
              "display": "Email",
              "regex": "\\S+@\\S+\\.\\S+",
              "key": "email",
              "showRegexInput": true
            },
            {
              "required": false,
              "type": "text",
              "enum": [],
              "display": "dfsdf",
              "key": "dfsdf",
              "showRegexInput": false
            }
          ],
          "available_assignees": [
            "5b9b98150df588546aaea6d2",
            "5c45d78395d7504f76c2cb37"
          ],
          "_id": "5fd72db3dc250f8decfc61b2",
          "title": "SOme title",
          "description": "SOme big description",
          "slug": "some-title",
          "application_id": "000000000000000000000003",
          "created_on": {
            "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36",
            "platform": "web",
            "meta": {
              "browser": {
                "name": "Chrome",
                "version": "87.0.4280.88"
              },
              "os": {
                "name": "macOS",
                "version": "10.15.6",
                "versionName": "Catalina"
              },
              "platform": {
                "type": "desktop",
                "vendor": "Apple"
              },
              "engine": {
                "name": "Blink"
              }
            }
          },
          "created_by": "5d1363adf599d850df93175e",
          "createdAt": "2020-12-14T09:17:39.953Z",
          "updatedAt": "2021-01-28T18:48:07.717Z",
          "__v": 0
        },
        "key": "some-title",
        "display": "SOme title"
      },
      "ticket_id": "43",
      "createdAt": "2021-01-28T17:25:25.013Z",
      "updatedAt": "2021-01-28T17:25:33.396Z",
      "__v": 0,
      "video_room_id": "6012f38557751ee8fc162cf7"
    }
  }
}
```









---


#### getParticipantsInsideVideoRoom
Get participants of a specific Video Room using it's unique name

```golang

 data, err :=  Lead.GetParticipantsInsideVideoRoom(UniqueName);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| UniqueName | string | Unique name of Video Room | 




Get participants of a specific Video Room using it's unique name, this can be used to check if people are already there in the room and also to show their names.

*Success Response:*



Success


Schema: `GetParticipantsInsideVideoRoomResponse`


*Examples:*


Default
```json
{
  "value": {
    "participants": []
  }
}
```









---


#### getTokenForVideoRoom
Get Token to join a specific Video Room using it's unqiue name

```golang

 data, err :=  Lead.GetTokenForVideoRoom(UniqueName);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| UniqueName | string | Unique name of Video Room | 




Get Token to join a specific Video Room using it's unqiue name, this Token is your ticket to Room and also creates your identity there.

*Success Response:*



Success


Schema: `GetTokenForVideoRoomResponse`


*Examples:*


Default
```json
{
  "value": {
    "access_token": "your_token_to_the_room"
  }
}
```









---



---


## Theme


#### getAppliedTheme
Get the theme currently applied to an application

```golang

 data, err :=  Theme.GetAppliedTheme();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



An application has multiple themes, but only one theme can be applied at a time. Use this API to retrieve the theme currently applied to the application.

*Success Response:*



Success. Returns a JSON object of the theme. Check the example shown below or    refer `ThemesSchema` for more details.


Schema: `ThemesSchema`


*Examples:*


Applied Theme
```json
{
  "$ref": "#/components/examples/Themes"
}
```











Schema: `BlitzkriegApiError`











Schema: `BlitzkriegInternalServerError`









---


#### getThemeForPreview
Get a theme for a preview

```golang

 data, err :=  Theme.GetThemeForPreview(ThemeID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ThemeID | string | ID of the theme to be retrieved | 




A theme can be previewed before applying it. Use this API to retrieve the preview of a theme by its ID.

*Success Response:*



Success. Returns a JSON object of the theme. Check the example shown below or refer `ThemesSchema` for more details.


Schema: `ThemesSchema`


*Examples:*


Preview Theme
```json
{
  "$ref": "#/components/examples/Themes"
}
```











Schema: `BlitzkriegApiError`











Schema: `BlitzkriegInternalServerError`









---



---


## User


#### loginWithFacebook
Login or Register using Facebook

```golang

 data, err :=  User.LoginWithFacebook(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  OAuthRequestSchema | "Request body" 


Use this API to login or register using Facebook credentials.

*Success Response:*



Success. Returns a JSON object with the user details. Check the example shown below or refer `AuthSuccess` for more details.


Schema: `AuthSuccess`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/AuthSuccess"
}
```











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### loginWithGoogle
Login or Register using Google

```golang

 data, err :=  User.LoginWithGoogle(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  OAuthRequestSchema | "Request body" 


Use this API to login or register using Google Account credentials.

*Success Response:*



Success. Returns a JSON object with the user details. Check the example shown below or refer `AuthSuccess` for more details.


Schema: `AuthSuccess`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/AuthSuccess"
}
```











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### loginWithGoogleAndroid
Login or Register using Google on Android

```golang

 data, err :=  User.LoginWithGoogleAndroid(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  OAuthRequestSchema | "Request body" 


Use this API to login or register in Android app using Google Account credentials.

*Success Response:*



Success. Returns a JSON object with the user details. Check the example shown below or refer `AuthSuccess` for more details.


Schema: `AuthSuccess`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/AuthSuccess"
}
```











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### loginWithGoogleIOS
Login or Register using Google on iOS

```golang

 data, err :=  User.LoginWithGoogleIOS(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  OAuthRequestSchema | "Request body" 


Use this API to login or register in iOS app using Google Account credentials.

*Success Response:*



Success. Returns a JSON object with the user details. Check the example shown below or refer `AuthSuccess` for more details.


Schema: `AuthSuccess`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/AuthSuccess"
}
```











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### loginWithOTP
Login or Register with OTP

```golang

 data, err :=  User.LoginWithOTP(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Platform`

| body |  SendOtpRequestSchema | "Request body" 


Use this API to login or register with a One-time Password (OTP) sent via Email or SMS.

*Success Response:*



Success. Check the example shown below or refer `SendOtpResponse` for more details.


Schema: `SendOtpResponse`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/SendOtpResponse"
}
```











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### loginWithEmailAndPassword
Login or Register with password

```golang

 data, err :=  User.LoginWithEmailAndPassword(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  PasswordLoginRequestSchema | "Request body" 


Use this API to login or register using an email address and password.

*Success Response:*



Success. Check the example shown below or refer `LoginSuccess` for more details.


Schema: `LoginSuccess`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/UserExampleObject"
}
```











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### sendResetPasswordEmail
Reset Password

```golang

 data, err :=  User.SendResetPasswordEmail(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Platform`

| body |  SendResetPasswordEmailRequestSchema | "Request body" 


Use this API to reset a password using the link sent on email.

*Success Response:*



Success. Check the example shown below or refer `ResetPasswordSuccess` for more details.


Schema: `ResetPasswordSuccess`











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### forgotPassword
Forgot Password

```golang

 data, err :=  User.ForgotPassword(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  ForgotPasswordRequestSchema | "Request body" 


Use this API to reset a password using the code sent on email or SMS.

*Success Response:*



Success. Check the example shown below or refer `LoginSuccess` for more details.


Schema: `LoginSuccess`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/UserExampleObject"
}
```











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### sendResetToken
Reset Password using token

```golang

 data, err :=  User.SendResetToken(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  CodeRequestBodySchema | "Request body" 


Use this API to send code to reset password.

*Success Response:*



Success. Check the example shown below or refer `ResetPasswordSuccess` for more details.


Schema: `ResetPasswordSuccess`











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### loginWithToken
Login or Register with token

```golang

 data, err :=  User.LoginWithToken(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  TokenRequestBodySchema | "Request body" 


Use this API to login or register using a token for authentication.

*Success Response:*



Success. Check the example shown below or refer `LoginSuccess` for more details.


Schema: `LoginSuccess`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/UserExampleObject"
}
```











Schema: `Object`











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### registerWithForm
Registration using a form

```golang

 data, err :=  User.RegisterWithForm(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Platform`

| body |  FormRegisterRequestSchema | "Request body" 


Use this API to perform user registration by sending form data in the request body.

*Success Response:*



Success. Check the example shown below or refer `RegisterFormSuccess` for more details.


Schema: `RegisterFormSuccess`











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### verifyEmail
Verify email

```golang

 data, err :=  User.VerifyEmail(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  CodeRequestBodySchema | "Request body" 


Use this API to send a verification code to verify an email.

*Success Response:*



Success. Check the example shown below or refer `VerifyEmailSuccess` for more details.


Schema: `VerifyEmailSuccess`











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### verifyMobile
Verify mobile

```golang

 data, err :=  User.VerifyMobile(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  CodeRequestBodySchema | "Request body" 


Use this API to send a verification code to verify a mobile number.

*Success Response:*



Success. Check the example shown below or refer `VerifyEmailSuccess` for more details.


Schema: `VerifyEmailSuccess`











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### hasPassword
Check password

```golang

 data, err :=  User.HasPassword();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to check if user has created a password for login.

*Success Response:*



Success. Returns a boolean value. Check the example shown below or refer `HasPasswordSuccess` for more details.


Schema: `HasPasswordSuccess`











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### updatePassword
Update user password

```golang

 data, err :=  User.UpdatePassword(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  UpdatePasswordRequestSchema | "Request body" 


Use this API to update the password.

*Success Response:*



Success. Returns a success message. Refer `VerifyEmailSuccess` for more details.


Schema: `VerifyEmailSuccess`











Schema: `AuthenticationApiError`











Schema: `AuthenticationApiError`









---


#### logout
Logs out currently logged in user

```golang

 data, err :=  User.Logout();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to check to logout a user from the app.

*Success Response:*



Success. Returns a success message as shown below. Refer `LogoutSuccess` for more details.


Schema: `LogoutSuccess`











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### sendOTPOnMobile
Send OTP on mobile

```golang

 data, err :=  User.SendOTPOnMobile(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Platform`

| body |  SendMobileOtpRequestSchema | "Request body" 


Use this API to send an OTP to a mobile number.

*Success Response:*



Success. Returns a JSON object as shown below. Refer `OtpSuccess` for more details.


Schema: `OtpSuccess`











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### verifyMobileOTP
Verify OTP on mobile

```golang

 data, err :=  User.VerifyMobileOTP(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Platform`

| body |  VerifyOtpRequestSchema | "Request body" 


Use this API to verify the OTP received on a mobile number.

*Success Response:*



Success. Returns a JSON object as shown below. Refer `VerifyOtpSuccess` for more details.


Schema: `VerifyOtpSuccess`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/VerifyMobileOTP"
}
```











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### sendOTPOnEmail
Send OTP on email

```golang

 data, err :=  User.SendOTPOnEmail(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Platform`

| body |  SendEmailOtpRequestSchema | "Request body" 


Use this API to send an OTP to an email ID.

*Success Response:*



Success. Returns a JSON object as shown below. Refer `EmailOtpSuccess` for more details.


Schema: `EmailOtpSuccess`











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### verifyEmailOTP
Verify OTP on email

```golang

 data, err :=  User.VerifyEmailOTP(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Platform`

| body |  VerifyEmailOtpRequestSchema | "Request body" 


Use this API to verify the OTP received on an email ID.

*Success Response:*



Success. Returns a JSON object as shown below. Refer `VerifyOtpSuccess` for more details.


Schema: `VerifyOtpSuccess`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/VerifyMobileOTP"
}
```











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### getLoggedInUser
Get logged in user

```golang

 data, err :=  User.GetLoggedInUser();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API  to get the details of a logged in user.

*Success Response:*



Success. Returns a JSON object with user details. Refer `UserObjectSchema` for more details.


Schema: `UserObjectSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/UserExampleObject"
}
```











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### getListOfActiveSessions
Get list of sessions

```golang

 data, err :=  User.GetListOfActiveSessions();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to retrieve all active sessions of a user.

*Success Response:*



Success. Returns a JSON object containing an array of sessions. Refer `SessionListSuccess` for more details.


Schema: `SessionListSuccess`











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### getPlatformConfig
Get platform configurations

```golang

 data, err :=  User.GetPlatformConfig(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Name`



Use this API to get all the platform configurations such as mobile image, desktop image, social logins, and all other text.

*Success Response:*



Success. Returns a JSON object containing the all the platform configurations. Refer `PlatformSchema` for more details.


Schema: `PlatformSchema`











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### updateProfile
Edit Profile Details

```golang

 data, err :=  User.UpdateProfile(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Platform`

| body |  EditProfileRequestSchema | "Request body" 


Use this API to update details in the user profile. Details can be first name, last name, gender, email, phone number, or profile picture.

*Success Response:*



Success. Check the example shown below or refer `LoginSuccess` for more details.


Schema: `ProfileEditSuccess`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/UserExampleObject"
}
```











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### addMobileNumber
Add mobile number to profile

```golang

 data, err :=  User.AddMobileNumber(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Platform`

| body |  EditMobileRequestSchema | "Request body" 


Use this API to add a new mobile number to a profile.

*Success Response:*



Success. Check the example shown below or refer `VerifyMobileOTPSuccess` for more details.


Schema: `VerifyMobileOTPSuccess`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/VerifyMobileOTP"
}
```











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### deleteMobileNumber
Delete mobile number from profile

```golang

 data, err :=  User.DeleteMobileNumber(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |












| xQuery | struct | Includes properties such as `Platform`, `Active`, `Primary`, `Verified`, `CountryCode`, `Phone`



Use this API to delete a mobile number from a profile.

*Success Response:*



Success. Check the example shown below or refer `LoginSuccess` for more details.


Schema: `LoginSuccess`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/UserExampleObject"
}
```











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### setMobileNumberAsPrimary
Set mobile as primary

```golang

 data, err :=  User.SetMobileNumberAsPrimary(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  SendVerificationLinkMobileRequestSchema | "Request body" 


Use this API to set a mobile number as primary. Primary number is a verified number used for all future communications.

*Success Response:*



Success. Check the example shown below or refer `LoginSuccess` for more details.


Schema: `LoginSuccess`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/UserExampleObject"
}
```











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### sendVerificationLinkToMobile
Send verification link to mobile

```golang

 data, err :=  User.SendVerificationLinkToMobile(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Platform`

| body |  SendVerificationLinkMobileRequestSchema | "Request body" 


Use this API to send a verification link to a mobile number

*Success Response:*



Success. Check the example shown below or refer `SendMobileVerifyLinkSuccess` for more details.


Schema: `SendMobileVerifyLinkSuccess`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/VerifyMobileOTP"
}
```











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### addEmail
Add email to profile

```golang

 data, err :=  User.AddEmail(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Platform`

| body |  EditEmailRequestSchema | "Request body" 


Use this API to add a new email address to a profile

*Success Response:*



Success. Returns a JSON object with user details. Refer `VerifyEmailOTPSuccess` for more details.


Schema: `VerifyEmailOTPSuccess`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/VerifyEmailOTP"
}
```











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### deleteEmail
Delete email from profile

```golang

 data, err :=  User.DeleteEmail(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |










| xQuery | struct | Includes properties such as `Platform`, `Active`, `Primary`, `Verified`, `Email`



Use this API to delete an email address from a profile

*Success Response:*



Success. Returns a JSON object with user details. Refer `LoginSuccess` for more details.


Schema: `LoginSuccess`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/UserExampleObject"
}
```











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### setEmailAsPrimary
Set email as primary

```golang

 data, err :=  User.SetEmailAsPrimary(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  EditEmailRequestSchema | "Request body" 


Use this API to set an email address as primary. Primary email ID is a email address used for all future communications.

*Success Response:*



Success. Returns a JSON object with user details. Refer `LoginSuccess` for more details.


Schema: `LoginSuccess`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/UserExampleObject"
}
```











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---


#### sendVerificationLinkToEmail
Send verification link to email

```golang

 data, err :=  User.SendVerificationLinkToEmail(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Platform`

| body |  EditEmailRequestSchema | "Request body" 


Use this API to send verification link to an email address.

*Success Response:*



Request body must contain an email ID. Refer `EditEmailRequestSchema` for more details.


Schema: `SendEmailVerifyLinkSuccess`











Schema: `AuthenticationApiError`











Schema: `AuthenticationInternalServerError`









---



---


## Content


#### getAnnouncements
Get live announcements

```golang

 data, err :=  Content.GetAnnouncements();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Announcements are useful to highlight a message or information on top of a webpage. Use this API to retrieve live announcements. Get announcements on individual pages or for all pages.

*Success Response:*



Success. Returns a JSON object with the details of the announcement shown on an individual page. `$all` is a special slug to indicate that an announcement is being shown on all the pages. Check the example shown below or refer `AnnouncementsResponseSchema` for more details.


Schema: `AnnouncementsResponseSchema`


*Examples:*


Announcements enabled
```json
{
  "$ref": "#/components/examples/AnnouncementEnabledExample"
}
```

No Announcement enabled
```json
{
  "value": {
    "announcements": {},
    "refresh_rate": 900,
    "refresh_pages": []
  }
}
```








API Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/4XXAPIError"
}
```








Internal Server Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/5XXAPIError"
}
```









---


#### getBlog
Get a blog

```golang

 data, err :=  Content.GetBlog(Slug, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of a blog. You can get slug value from the endpoint /service/application/content/v1.0/blogs/. | 



| xQuery | struct | Includes properties such as `RootID`



Use this API to get the details of a blog using its slug. Details include the title, reading time, publish status, feature image, tags, author, etc.

*Success Response:*



Success. Returns a JSON object with blog details. Check the example shown below or refer `CustomBlogSchema` for more details.


Schema: `CustomBlogSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/CustomBlog"
}
```








API Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/4XXAPIError"
}
```








Internal Server Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/5XXAPIError"
}
```









---


#### getBlogs
Get a list of blogs

```golang

 data, err :=  Content.GetBlogs(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `PageNo`, `PageSize`



Use this API to get all the blogs.

*Success Response:*



Success. Check the example shown below or refer `BlogGetResponse` for more details.


Schema: `BlogGetResponse`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/BlogGetResponse"
}
```








API Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/4XXAPIError"
}
```








Internal Server Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/5XXAPIError"
}
```









---


#### getFaqs
Get a list of FAQs

```golang

 data, err :=  Content.GetFaqs();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to get a list of frequently asked questions. Users will benefit from it when facing any issue with the website.

*Success Response:*



Success. Returns a JSON object with question and answers. Check the example shown below or refer `FaqResponseSchema` for more details.


Schema: `FaqResponseSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/AppFaqs"
}
```








API Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/4XXAPIError"
}
```








Internal Server Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/5XXAPIError"
}
```









---


#### getFaqCategories
Get a list of FAQ categories

```golang

 data, err :=  Content.GetFaqCategories();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



FAQs can be divided into categories. Use this API to get a list of FAQ categories.

*Success Response:*



Success. Returns a JSON object with categories of FAQ. Check the example shown below or refer `GetFaqCategoriesSchema` for more details.


Schema: `GetFaqCategoriesSchema`








API Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/4XXAPIError"
}
```








Internal Server Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/5XXAPIError"
}
```









---


#### getFaqBySlug
Get an FAQ

```golang

 data, err :=  Content.GetFaqBySlug(Slug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of an FAQ. You can get slug value from the endpoint /service/application/content/v1.0/faq. | 




Use this API to get a particular FAQ by its slug.

*Success Response:*



Success. Returns a question and answer by its slug. Check the example shown below or refer `FaqSchema` for more details.


Schema: `FaqSchema`








API Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/4XXAPIError"
}
```








Internal Server Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/5XXAPIError"
}
```









---


#### getFaqCategoryBySlug
Get the FAQ category

```golang

 data, err :=  Content.GetFaqCategoryBySlug(Slug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of an FAQ category. You can get slug value from the endpoint /service/application/content/v1.0/faq/categories. | 




FAQs can be divided into categories. Use this API to get the category to which an FAQ belongs.

*Success Response:*



Success. Returns a FAQ category with its slug. Check the example shown below or refer `GetFaqCategoryBySlugSchema` for more details.


Schema: `GetFaqCategoryBySlugSchema`








API Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/4XXAPIError"
}
```








Internal Server Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/5XXAPIError"
}
```









---


#### getFaqsByCategorySlug
Get FAQs using the slug of FAQ category

```golang

 data, err :=  Content.GetFaqsByCategorySlug(Slug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of an FAQ category. You can get slug value from the endpoint /service/application/content/v1.0/faq/categories. | 




FAQs can be divided into categories. Use this API to get all the FAQs belonging to a category by using the category slug.

*Success Response:*



Success. Returns a categorized list of question and answers using its slug. Check the example shown below or refer `GetFaqSchema` for more details.


Schema: `GetFaqSchema`








API Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/4XXAPIError"
}
```








Internal Server Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/5XXAPIError"
}
```









---


#### getLandingPage
Get the landing page

```golang

 data, err :=  Content.GetLandingPage();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Landing page is the first page that a prospect lands upon while visiting a website. Use this API to fetch the details of a landing page.

*Success Response:*



Success. Returns the landing page details. Check the example shown below or refer `LandingPageSchema` for more details.


Schema: `LandingPageSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/LandingPageResponse"
}
```








API Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/4XXAPIError"
}
```








Internal Server Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/5XXAPIError"
}
```









---


#### getLegalInformation
Get legal information

```golang

 data, err :=  Content.GetLegalInformation();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to get the legal information of an application, which includes Privacy Policy, Terms and Conditions, Shipping Policy and FAQs regarding the usage of the application.

*Success Response:*



Success. Returns the T&C, Shipping Policy, Privacy Policy and Return Policy. Check the example shown below or refer `ApplicationLegal` for more details.


Schema: `ApplicationLegal`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/Legal"
}
```








API Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/4XXAPIError"
}
```








Internal Server Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/5XXAPIError"
}
```









---


#### getNavigations
Get the navigation

```golang

 data, err :=  Content.GetNavigations(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `PageNo`, `PageSize`



Use this API to fetch the navigations details which includes the items of the navigation pane. It also shows the links and sub-navigations.

*Success Response:*



Success. Returns a JSON object with navigation details. Check the example shown below or refer `NavigationGetResponse` for more details.


Schema: `NavigationGetResponse`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/NavigationGetResponse"
}
```








API Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/4XXAPIError"
}
```








Internal Server Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/5XXAPIError"
}
```









---


#### getPage
Get a page

```golang

 data, err :=  Content.GetPage(Slug, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of a page. You can get slug value from the endpoint /service/application/content/v1.0/pages/. | 



| xQuery | struct | Includes properties such as `RootID`



Use this API to get the details of a page using its slug. Details include the title, seo, publish status, feature image, tags, meta, etc.

*Success Response:*



Success. Returns a JSON object with page details. Check the example shown below or refer `CustomPageSchema` for more details.


Schema: `CustomPageSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/PageResponse"
}
```








API Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/4XXAPIError"
}
```








Internal Server Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/5XXAPIError"
}
```









---


#### getPages
Get all pages

```golang

 data, err :=  Content.GetPages(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `PageNo`, `PageSize`



Use this API to get a list of pages.

*Success Response:*



Success. Returns a list of pages along with their details. Check the example shown below or refer `PageGetResponse` for more details.


Schema: `PageGetResponse`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/PageGetResponse"
}
```








API Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/4XXAPIError"
}
```








Internal Server Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/5XXAPIError"
}
```









---


#### getSEOConfiguration
Get the SEO of an application

```golang

 data, err :=  Content.GetSEOConfiguration();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to get the SEO details of an application, which includes a robot.txt, meta-tags and sitemap.

*Success Response:*



Success. Returns a JSON object SEO details such as robots.txt, meta-tags, and sitemap. Check the example shown below or refer `SeoComponent` for more details.


Schema: `SeoComponent`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/Seo"
}
```








API Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/4XXAPIError"
}
```








Internal Server Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/5XXAPIError"
}
```









---


#### getSlideshows
Get the slideshows

```golang

 data, err :=  Content.GetSlideshows(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `PageNo`, `PageSize`



Use this API to get a list of slideshows along with their details.

*Success Response:*



Success. Check the example shown below or refer `SlideshowGetResponse` for more details.


Schema: `SlideshowGetResponse`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/SlideshowGetResponse"
}
```








API Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/4XXAPIError"
}
```








Internal Server Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/5XXAPIError"
}
```









---


#### getSlideshow
Get a slideshow

```golang

 data, err :=  Content.GetSlideshow(Slug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of a slideshow. You can get slug value from the endpoint /service/application/content/v1.0/slideshow/. | 




A slideshow is a group of images, videos or a combination of both that are shown on the website in the form of slides. Use this API to fetch a slideshow using its `slug`.

*Success Response:*



Success. Returns the details of how a slideshow is configured. Check the example shown below or refer `SlideshowSchema` for more details.


Schema: `SlideshowSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/SlideshowResponse"
}
```








API Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/4XXAPIError"
}
```








Internal Server Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/5XXAPIError"
}
```









---


#### getSupportInformation
Get the support information

```golang

 data, err :=  Content.GetSupportInformation();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to get contact details for customer support including emails and phone numbers.

*Success Response:*



Success. Returns all support information including email and phone number. Check the example shown below or refer `Support` for more details.


Schema: `Support`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/Support"
}
```








API Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/4XXAPIError"
}
```








Internal Server Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/5XXAPIError"
}
```









---


#### getTags
Get the tags associated with an application

```golang

 data, err :=  Content.GetTags();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to get all the CSS and JS injected in the application in the form of tags.

*Success Response:*



Success. Returns a JSON object containing all the tags injected in the application. Check the example shown below or refer `TagsSchema` for more details.


Schema: `TagsSchema`








API Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/4XXAPIError"
}
```








Internal Server Error. See the error object in the response body to know the exact reason.


Schema: `APIError`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/5XXAPIError"
}
```









---



---


## Communication


#### getCommunicationConsent
Get communication consent

```golang

 data, err :=  Communication.GetCommunicationConsent();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Get communication consent

*Success Response:*



Success


Schema: `CommunicationConsent`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/CommunicationConsent"
}
```









---


#### upsertCommunicationConsent
Upsert communication consent

```golang

 data, err :=  Communication.UpsertCommunicationConsent(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  CommunicationConsentReq | "Request body" 


Upsert communication consent

*Success Response:*



Success


Schema: `CommunicationConsentRes`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/CommunicationConsentRes"
}
```








Bad request


Schema: `BadRequest`









---


#### upsertAppPushtoken
Upsert push token of a user

```golang

 data, err :=  Communication.UpsertAppPushtoken(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  PushtokenReq | "Request body" 


Upsert push token of a user

*Success Response:*



Success


Schema: `PushtokenRes`


*Examples:*


create
```json
{
  "$ref": "#/components/examples/PushtokenResponseCreate"
}
```

update
```json
{
  "$ref": "#/components/examples/PushtokenResponseUpdate"
}
```

reset
```json
{
  "$ref": "#/components/examples/PushtokenResponseReset"
}
```








Bad request


Schema: `BadRequest`









---



---


## Share


#### getApplicationQRCode
Create QR Code of an app

```golang

 data, err :=  Share.GetApplicationQRCode();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to create a QR code of an app for sharing it with users who want to use the app.

*Success Response:*



Success. Check the example shown below or refer `QRCodeResp` for more details.


Schema: `QRCodeResp`









---


#### getProductQRCodeBySlug
Create QR Code of a product

```golang

 data, err :=  Share.GetProductQRCodeBySlug(Slug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of a product. You can get slug value from the endpoint. | 




Use this API to create a QR code of a product for sharing it with users who want to view/purchase the product.

*Success Response:*



Success. Check the example shown below or refer `QRCodeResp` for more details.


Schema: `QRCodeResp`








Error


Schema: `ErrorRes`









---


#### getCollectionQRCodeBySlug
Create QR Code of a collection

```golang

 data, err :=  Share.GetCollectionQRCodeBySlug(Slug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of a collection. You can get slug value from the endpoint. | 




Use this API to create a QR code of a collection of products for sharing it with users who want to view/purchase the collection.

*Success Response:*



Success. Check the example shown below or refer `QRCodeResp` for more details.


Schema: `QRCodeResp`








Error


Schema: `ErrorRes`









---


#### getUrlQRCode
Create QR Code of a URL

```golang

 data, err :=  Share.GetUrlQRCode(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `URL`



Use this API to create a QR code of a URL for sharing it with users who want to visit the link.

*Success Response:*



Success. Check the example shown below or refer `QRCodeResp` for more details.


Schema: `QRCodeResp`








Error


Schema: `ErrorRes`









---


#### createShortLink
Create a short link

```golang

 data, err :=  Share.CreateShortLink(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  ShortLinkReq | "Request body" 


Use this API to create a short link that is easy to write/share/read as compared to long URLs.

*Success Response:*



Success. Check the example shown below or refer `ShortLinkRes` for more details.


Schema: `ShortLinkRes`








Error


Schema: `ErrorRes`









---


#### getShortLinkByHash
Get short link by hash

```golang

 data, err :=  Share.GetShortLinkByHash(Hash);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Hash | string | A string value used for converting long URL to short URL and vice-versa. | 




Use this API to get a short link by using a hash value.

*Success Response:*



Success. Check the example shown below or refer `ShortLinkRes` for more details.


Schema: `ShortLinkRes`








Error


Schema: `ErrorRes`









---


#### getOriginalShortLinkByHash
Get original link by hash

```golang

 data, err :=  Share.GetOriginalShortLinkByHash(Hash);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Hash | string | A string value used for converting long URL to short URL and vice-versa. | 




Use this API to retrieve the original link from a short-link by using a hash value.

*Success Response:*



Success. Check the example shown below or refer `ShortLinkRes` for more details.


Schema: `ShortLinkRes`








Error


Schema: `ErrorRes`









---



---


## FileStorage


#### startUpload
This operation initiates upload and returns storage link which is valid for 30 Minutes. You can use that storage link to make subsequent upload request with file buffer or blob.

```golang

 data, err :=  FileStorage.StartUpload(Namespace, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Namespace | string | bucket name | 


| body |  StartRequest | "Request body" 


Uploads an arbitrarily sized buffer or blob.

It has three Major Steps:
* Start
* Upload
* Complete

### Start
Initiates the assets upload using `startUpload`.
It returns the storage link in response.

### Upload
Use the storage link to upload a file (Buffer or Blob) to the File Storage.
Make a `PUT` request on storage link received from `startUpload` api with file (Buffer or Blob) as a request body.

### Complete
After successfully upload, call `completeUpload` api to complete the upload process.
This operation will return the url for the uploaded file.


*Success Response:*



Success


Schema: `StartResponse`








Failed


Schema: `FailedResponse`









---


#### completeUpload
This will complete the upload process. After successfully uploading file, you can call this operation to complete the upload process.

```golang

 data, err :=  FileStorage.CompleteUpload(Namespace, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Namespace | string | bucket name | 


| body |  StartResponse | "Request body" 


Uploads an arbitrarily sized buffer or blob.

It has three Major Steps:
* Start
* Upload
* Complete

### Start
Initiates the assets upload using `startUpload`.
It returns the storage link in response.

### Upload
Use the storage link to upload a file (Buffer or Blob) to the File Storage.
Make a `PUT` request on storage link received from `startUpload` api with file (Buffer or Blob) as a request body.

### Complete
After successfully upload, call `completeUpload` api to complete the upload process.
This operation will return the url for the uploaded file.


*Success Response:*



Success


Schema: `CompleteResponse`








Failed


Schema: `FailedResponse`









---



---


## Configuration


#### getApplication
Get current application details

```golang

 data, err :=  Configuration.GetApplication();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Get current application details.

*Success Response:*



Success


Schema: `Application`








Not found


Schema: `NotFound`









---


#### getOwnerInfo
Get application, owner and seller information

```golang

 data, err :=  Configuration.GetOwnerInfo();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Get application information with owner and seller basic details

*Success Response:*



Success


Schema: `ApplicationAboutResponse`









---


#### getBasicDetails
Get basic application details

```golang

 data, err :=  Configuration.GetBasicDetails();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Get basic application details like name

*Success Response:*



Success


Schema: `ApplicationDetail`









---


#### getIntegrationTokens
Get integration tokens

```golang

 data, err :=  Configuration.GetIntegrationTokens();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Get tokens for multiple integrations like Facebook, Googlemaps, Segment, Firebase, etc. Note: token values are encrypted with AES encryption using secret key. Kindly reach to developers for secret key.

*Success Response:*



Success


Schema: `TokenResponse`









---


#### getOrderingStores
Get deployment meta stores

```golang

 data, err :=  Configuration.GetOrderingStores(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |






| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Q`



Get deployment meta stores.

*Success Response:*



Success


Schema: `OrderingStores`








Not found


Schema: `NotFound`









---


#### getFeatures
Get features of application

```golang

 data, err :=  Configuration.GetFeatures();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Get features of application

*Success Response:*



Success


Schema: `AppFeatureResponse`








Not found


Schema: `NotFound`









---


#### getContactInfo
Get application information

```golang

 data, err :=  Configuration.GetContactInfo();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Get Application Current Information. This includes information about social links, address and contact information of company/seller/brand of the application.

*Success Response:*



Success


Schema: `ApplicationInformation`









---


#### getCurrencies
Get application enabled currencies

```golang

 data, err :=  Configuration.GetCurrencies();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Get currency list for allowed currencies under current application

*Success Response:*



Currencies Success response


Schema: `CurrenciesResponse`









---


#### getCurrencyById
Get currency by id

```golang

 data, err :=  Configuration.GetCurrencyById(ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ID | string | Currency object id | 




Get currency object with symbol and name information by id.

*Success Response:*



Success response


Schema: `Currency`









---


#### getLanguages
Get list of languages

```golang

 data, err :=  Configuration.GetLanguages();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Get list of supported languages under application.

*Success Response:*



Success response


Schema: `LanguageResponse`









---


#### getOrderingStoreCookie
Get ordering store signed cookie on selection of ordering store. This will be used by cart service to verify coupon against selected ordering store in cart.

```golang

 data, err :=  Configuration.GetOrderingStoreCookie(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  OrderingStoreSelectRequest | "Request body" 


Get ordering store signed cookie on selection of ordering store.

*Success Response:*



Success


Schema: `SuccessMessageResponse`








Success


Schema: `NotFound`









---


#### removeOrderingStoreCookie
Unset ordering store signed cookie on change of sales channel selection via domain in universal fynd store app.

```golang

 data, err :=  Configuration.RemoveOrderingStoreCookie();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Unset ordering store cookie.

*Success Response:*



Success


Schema: `SuccessMessageResponse`









---


#### getAppStaffs
Get Staff List.

```golang

 data, err :=  Configuration.GetAppStaffs(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |






| xQuery | struct | Includes properties such as `OrderIncent`, `OrderingStore`, `User`



Get a staff list based on the user's session token passed in the header.

*Success Response:*



Success


Schema: `AppStaffResponse`








Request failed with internal server error.


Schema: `UnhandledError`









---



---


## Payment


#### getAggregatorsConfig
Get payment gateway keys

```golang

 data, err :=  Payment.GetAggregatorsConfig(XAPIToken, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| XAPIToken | string | api token | 



| xQuery | struct | Includes properties such as `Refresh`



Get payment gateway (key, secrets, merchant, sdk/api detail) to complete payment at front-end.

*Success Response:*



Keys of all payment gateway


Schema: `AggregatorsConfigDetailResponse`








Bad Request Error


Schema: `HttpErrorCodeAndResponse`








Internal Server Error


Schema: `HttpErrorCodeAndResponse`









---


#### attachCardToCustomer
Attach a saved card to customer.

```golang

 data, err :=  Payment.AttachCardToCustomer(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  AttachCardRequest | "Request body" 


Attach a saved card to customer at payment gateway i.e stripe and refresh card cache.

*Success Response:*



List of cards objects


Schema: `AttachCardsResponse`








Bad request error


Schema: `any`








Internal Server Error


Schema: `any`









---


#### getActiveCardAggregator
Fetch active payment gateway for card

```golang

 data, err :=  Payment.GetActiveCardAggregator(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Refresh`



Fetch active payment gateway along with customer id for cards payments.

*Success Response:*



Object of payment gateway and customer id


Schema: `ActiveCardPaymentGatewayResponse`








Bad request error


Schema: `HttpErrorCodeAndResponse`








Internal Server Error


Schema: `HttpErrorCodeAndResponse`









---


#### getActiveUserCards
Fetch the list of saved cards of user.

```golang

 data, err :=  Payment.GetActiveUserCards(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `ForceRefresh`



Fetch the list of saved cards of user from active payment gateway.

*Success Response:*



List of cards objects


Schema: `ListCardsResponse`








Bad request error


Schema: `HttpErrorCodeAndResponse`








Internal Server Error


Schema: `HttpErrorCodeAndResponse`









---


#### deleteUserCard
Delete an user card.

```golang

 data, err :=  Payment.DeleteUserCard(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  DeletehCardRequest | "Request body" 


Delete an added user card on payment gateway and remove from cache.

*Success Response:*



List of cards objects


Schema: `DeleteCardsResponse`








Bad request error


Schema: `HttpErrorCodeAndResponse`








Internal Server Error


Schema: `HttpErrorCodeAndResponse`









---


#### verifyCustomerForPayment
Validate customer for payment.

```golang

 data, err :=  Payment.VerifyCustomerForPayment(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  ValidateCustomerRequest | "Request body" 


Validate customer for payment i.e Simpl paylater, Rupifi loan.

*Success Response:*



List of cards objects


Schema: `ValidateCustomerResponse`








Bad request error


Schema: `HttpErrorCodeAndResponse`








Internal Server Error


Schema: `HttpErrorCodeAndResponse`









---


#### verifyAndChargePayment
Verify and charge payment

```golang

 data, err :=  Payment.VerifyAndChargePayment(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  ChargeCustomerRequest | "Request body" 


Verify and charge payment server to server for Simpl & Mswipe.

*Success Response:*



List of cards objects


Schema: `ChargeCustomerResponse`








Bad request error


Schema: `HttpErrorCodeAndResponse`








Internal Server Error


Schema: `HttpErrorCodeAndResponse`









---


#### initialisePayment
Payment Initialisation server to server for UPI and BharatQR.

```golang

 data, err :=  Payment.InitialisePayment(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  PaymentInitializationRequest | "Request body" 


Payment Initialisation for UPI & BharatQR code, UPI requests to app and QR code to be displayed on screen.

*Success Response:*



List of cards objects


Schema: `PaymentInitializationResponse`








Bad request error


Schema: `HttpErrorCodeAndResponse`








Internal Server Error


Schema: `HttpErrorCodeAndResponse`









---


#### checkAndUpdatePaymentStatus
Continous polling to check status of payment on server.

```golang

 data, err :=  Payment.CheckAndUpdatePaymentStatus(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  PaymentStatusUpdateRequest | "Request body" 


Continous polling on interval to check status of payment untill timeout.

*Success Response:*



List of cards objects


Schema: `PaymentStatusUpdateResponse`








Bad request error


Schema: `HttpErrorCodeAndResponse`








Internal Server Error


Schema: `HttpErrorCodeAndResponse`









---


#### getPaymentModeRoutes
Get All Valid Payment Options

```golang

 data, err :=  Payment.GetPaymentModeRoutes(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |














| xQuery | struct | Includes properties such as `Amount`, `CartID`, `Pincode`, `CheckoutMode`, `Refresh`, `AssignCardID`, `UserDetails`



Use this API to get Get All Valid Payment Options for making payment

*Success Response:*



Success


Schema: `PaymentModeRouteResponse`








Bad Request Error


Schema: `HttpErrorCodeAndResponse`








Internal Server Error


Schema: `HttpErrorCodeAndResponse`









---


#### getPosPaymentModeRoutes
Get All Valid Payment Options for POS

```golang

 data, err :=  Payment.GetPosPaymentModeRoutes(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |
















| xQuery | struct | Includes properties such as `Amount`, `CartID`, `Pincode`, `CheckoutMode`, `Refresh`, `AssignCardID`, `OrderType`, `UserDetails`



Use this API to get Get All Valid Payment Options for making payment

*Success Response:*



Success


Schema: `PaymentModeRouteResponse`








Bad Request Error


Schema: `HttpErrorCodeAndResponse`








Internal Server Error


Schema: `HttpErrorCodeAndResponse`









---


#### getActiveRefundTransferModes
List Refund Transfer Mode

```golang

 data, err :=  Payment.GetActiveRefundTransferModes();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Get all active transfer mode for adding beneficiary details

*Success Response:*



Refund Transfer Mode


Schema: `TransferModeResponse`








Internal Server Error


Schema: `HttpErrorCodeAndResponse`









---


#### enableOrDisableRefundTransferMode
Enable/Disable Refund Transfer Mode

```golang

 data, err :=  Payment.EnableOrDisableRefundTransferMode(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  UpdateRefundTransferModeRequest | "Request body" 


Activate or Deactivate Transfer Mode to collect Beneficiary Details for Refund

*Success Response:*



Update Refund Transfer Mode.


Schema: `UpdateRefundTransferModeResponse`








Internal Server Error


Schema: `HttpErrorCodeAndResponse`









---


#### getUserBeneficiariesDetail
List User Beneficiary

```golang

 data, err :=  Payment.GetUserBeneficiariesDetail(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `OrderID`



Get all active  beneficiary details added by the user for refund

*Success Response:*



List User Beneficiary


Schema: `OrderBeneficiaryResponse`








Bad Request Error


Schema: `NotFoundResourceError`








Internal Server Error


Schema: `HttpErrorCodeAndResponse`









---


#### verifyIfscCode
Ifsc Code Verification

```golang

 data, err :=  Payment.VerifyIfscCode(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `IfscCode`



Get True/False for correct IFSC Code for adding bank details for refund

*Success Response:*



Bank details on correct Ifsc Code


Schema: `IfscCodeResponse`








Bad Request Error


Schema: `NotFoundResourceError`








Internal Server Error


Schema: `ErrorCodeDescription`









---


#### getOrderBeneficiariesDetail
List Order Beneficiary

```golang

 data, err :=  Payment.GetOrderBeneficiariesDetail(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `OrderID`



Get all active  beneficiary details added by the user for refund

*Success Response:*



List Order Beneficiary


Schema: `OrderBeneficiaryResponse`








Bad Request Error


Schema: `NotFoundResourceError`








Internal Server Error


Schema: `HttpErrorCodeAndResponse`









---


#### verifyOtpAndAddBeneficiaryForBank
Save Beneficiary details on otp validation.

```golang

 data, err :=  Payment.VerifyOtpAndAddBeneficiaryForBank(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  AddBeneficiaryViaOtpVerificationRequest | "Request body" 


Save Beneficiary details on otp validation.

*Success Response:*



Success


Schema: `AddBeneficiaryViaOtpVerificationResponse`








Bad Request Error


Schema: `WrongOtpError`








Internal Server Error


Schema: `HttpErrorCodeAndResponse`









---


#### addBeneficiaryDetails
Save bank details for cancelled/returned order

```golang

 data, err :=  Payment.AddBeneficiaryDetails(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  AddBeneficiaryDetailsRequest | "Request body" 


Use this API to save bank details for returned/cancelled order to refund amount in his account.

*Success Response:*



Success


Schema: `RefundAccountResponse`








Bad Request Error


Schema: `NotFoundResourceError`








Internal Server Error


Schema: `HttpErrorCodeAndResponse`









---


#### verifyOtpAndAddBeneficiaryForWallet
Send Otp on Adding wallet beneficiary

```golang

 data, err :=  Payment.VerifyOtpAndAddBeneficiaryForWallet(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  WalletOtpRequest | "Request body" 


Send Otp on Adding wallet beneficiary for user mobile verification

*Success Response:*



WalletOtp


Schema: `WalletOtpResponse`








Bad Request Error


Schema: `NotFoundResourceError`








Internal Server Error


Schema: `HttpErrorCodeAndResponse`









---


#### updateDefaultBeneficiary
Mark Default Beneficiary For Refund

```golang

 data, err :=  Payment.UpdateDefaultBeneficiary(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  SetDefaultBeneficiaryRequest | "Request body" 


Mark Default Beneficiary ot of all Beneficiary Details for Refund

*Success Response:*



Set Default Beneficiary Response.


Schema: `SetDefaultBeneficiaryResponse`








Bad Request Error


Schema: `SetDefaultBeneficiaryResponse`








Internal Server Error


Schema: `HttpErrorCodeAndResponse`









---



---


## Order


#### getOrders
Get Orders for application based on application Id

```golang

 data, err :=  Order.GetOrders(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |










| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `FromDate`, `ToDate`, `OrderStatus`



Get Orders

*Success Response:*



Success


Schema: `OrderList`








API Error


Schema: `ApefaceApiError`








Internal Server Error


Schema: `ApefaceApiError`









---


#### getOrderById
Get Order by order id for application based on application Id

```golang

 data, err :=  Order.GetOrderById(OrderID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| OrderID | string | Order Id | 




Get Order By Fynd Order Id

*Success Response:*



Success


Schema: `OrderById`








API Error


Schema: `ApefaceApiError`








Internal Server Error


Schema: `ApefaceApiError`









---


#### getShipmentById
Get Shipment by shipment id and order id for application based on application Id

```golang

 data, err :=  Order.GetShipmentById(ShipmentID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ShipmentID | string | Shipment Id | 




Get Shipment

*Success Response:*



Success


Schema: `ShipmentById`








API Error


Schema: `ApefaceApiError`








Internal Server Error


Schema: `ApefaceApiError`









---


#### getShipmentReasons
Get Shipment reasons by shipment id and order id for application based on application Id

```golang

 data, err :=  Order.GetShipmentReasons(ShipmentID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ShipmentID | string | Shipment Id | 




Get Shipment Reasons

*Success Response:*



Success


Schema: `ShipmentReasons`








API Error


Schema: `ApefaceApiError`








Internal Server Error


Schema: `ApefaceApiError`









---


#### updateShipmentStatus
Update Shipment status by shipment id and order id for application based on application Id

```golang

 data, err :=  Order.UpdateShipmentStatus(ShipmentID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ShipmentID | string | Shipment Id | 


| body |  ShipmentStatusUpdateBody | "Request body" 


Update Shipment Status

*Success Response:*



Success


Schema: `ShipmentStatusUpdate`








API Error


Schema: `ApefaceApiError`








Internal Server Error


Schema: `ApefaceApiError`









---


#### trackShipment
Track Shipment by shipment id and order id for application based on application Id

```golang

 data, err :=  Order.TrackShipment(ShipmentID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ShipmentID | string | Shipment Id | 




Shipment Track

*Success Response:*



Success


Schema: `ShipmentTrack`








API Error


Schema: `ApefaceApiError`








Internal Server Error


Schema: `ApefaceApiError`









---


#### getPosOrderById
Get POS Order by order id for application based on application Id

```golang

 data, err :=  Order.GetPosOrderById(OrderID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| OrderID | string | Order Id | 




Get Order By Fynd Order Id

*Success Response:*



Success


Schema: `PosOrderById`








API Error


Schema: `ApefaceApiError`








Internal Server Error


Schema: `ApefaceApiError`









---



---


## Rewards


#### getPointsOnProduct
Get the eligibility of reward points on a product

```golang

 data, err :=  Rewards.GetPointsOnProduct(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  CatalogueOrderRequest | "Request body" 


Use this API to evaluate the amount of reward points that could be earned on any catalogue product.

*Success Response:*



Success. Check example below or refer `CatalogueOrderRequest` for more details.


Schema: `CatalogueOrderResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `Error`









---


#### getOfferByName
Get offer by name

```golang

 data, err :=  Rewards.GetOfferByName(Name);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Name | string | The name given to the offer. | 




Use this API to get the offer details and configuration by entering the name of the offer.

*Success Response:*



Success. Check example below or refer `Offer` for more details.


Schema: `Offer`








Bad request. See the error object in the response body to know the exact reason.


Schema: `Error`









---


#### getOrderDiscount
Calculates the discount on order-amount

```golang

 data, err :=  Rewards.GetOrderDiscount(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  OrderDiscountRequest | "Request body" 


Use this API to calculate the discount on order-amount based on all the amount range configured in order_discount.

*Success Response:*



Success. Check example below or refer `OrderDiscountResponse` for more details.


Schema: `OrderDiscountResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `Error`









---


#### getUserPoints
Get reward points available with a user

```golang

 data, err :=  Rewards.GetUserPoints();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to retrieve total available points of a user for current application

*Success Response:*



Success. Check example below or refer `PointsResponse` for more details.


Schema: `PointsResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `Error`









---


#### getUserPointsHistory
Get all transactions of reward points

```golang

 data, err :=  Rewards.GetUserPointsHistory(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `PageID`, `PageSize`



Use this API to get a list of points transactions. The list of points history is paginated.

*Success Response:*



Success. Check example below or refer `PointsHistoryResponse` for more details.


Schema: `PointsHistoryResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `Error`









---


#### getUserReferralDetails
Get referral details of a user

```golang

 data, err :=  Rewards.GetUserReferralDetails();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to retrieve the referral details a user has configured in the application.

*Success Response:*



Success. Check example below or refer `ReferralDetailsResponse` for more details.


Schema: `ReferralDetailsResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `Error`









---


#### redeemReferralCode
Redeems a referral code and credits reward points to users

```golang

 data, err :=  Rewards.RedeemReferralCode(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  RedeemReferralCodeRequest | "Request body" 


Use this API to enter a referral code following which, the configured points would be credited to a user's reward points account.

*Success Response:*



Success. Check example below or refer `RedeemReferralCodeResponse` for more details.


Schema: `RedeemReferralCodeResponse`








Bad request. See the error object in the response body to know the exact reason.


Schema: `Error`









---



---


## Feedback


#### createAbuseReport
Post a new abuse request

```golang

 data, err :=  Feedback.CreateAbuseReport(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  ReportAbuseRequest | "Request body" 


Use this API to report a specific entity (question/review/comment) for abuse.

*Success Response:*



Success. Returns an abuse ID.


Schema: `InsertResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### updateAbuseReport
Update abuse details

```golang

 data, err :=  Feedback.UpdateAbuseReport(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  UpdateAbuseStatusRequest | "Request body" 


Use this API to update the abuse details, i.e. status and description.

*Success Response:*



Success.


Schema: `UpdateResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### getAbuseReports
Get a list of abuse data

```golang

 data, err :=  Feedback.GetAbuseReports(EntityID, EntityType, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| EntityID | string | ID of the eligible entity as specified in the entity type (question ID/review ID/comment ID). | 


| EntityType | string | Type of entity, e.g. question, review or comment. | 







| xQuery | struct | Includes properties such as `ID`, `PageID`, `PageSize`



Use this API to retrieve a list of abuse data from entity type and entity ID.

*Success Response:*



Success. Check the example shown below or refer `ReportAbuseGetResponse` for more details.


Schema: `ReportAbuseGetResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### getAttributes
Get a list of attribute data

```golang

 data, err :=  Feedback.GetAttributes(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `PageNo`, `PageSize`



Use this API to retrieve a list of all attribute data, e.g. quality, material, product fitting, packaging, etc.

*Success Response:*



Success. Check the example shown below or refer `AttributeResponse` for more details.


Schema: `AttributeResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### createAttribute
Add a new attribute request

```golang

 data, err :=  Feedback.CreateAttribute(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  SaveAttributeRequest | "Request body" 


Use this API to add a new attribute (e.g. product quality/material/value for money) with its name, slug and description.

*Success Response:*



Success. Returns an attribute ID.


Schema: `InsertResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### getAttribute
Get data of a single attribute

```golang

 data, err :=  Feedback.GetAttribute(Slug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of an attribute. You can get slug value from the endpoint 'service/application/feedback/v1.0/attributes'. | 




Use this API to retrieve a single attribute data from a given slug.

*Success Response:*



Success. Check the example shown below or refer `Attribute` for more details.


Schema: `Attribute`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### updateAttribute
Update details of an attribute 

```golang

 data, err :=  Feedback.UpdateAttribute(Slug, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of an attribute. You can get slug value from the endpoint 'service/application/feedback/v1.0/attributes'. | 


| body |  UpdateAttributeRequest | "Request body" 


Use this API update the attribute's name and description.

*Success Response:*



Success.


Schema: `UpdateResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### createComment
Post a new comment

```golang

 data, err :=  Feedback.CreateComment(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  CommentRequest | "Request body" 


Use this API to add a new comment for a specific entity.

*Success Response:*



Success. Returns a comment ID.


Schema: `InsertResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### updateComment
Update the status of a comment

```golang

 data, err :=  Feedback.UpdateComment(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  UpdateCommentRequest | "Request body" 


Use this API to update the comment status (active or approve) along with new comment if any.

*Success Response:*



Success.


Schema: `UpdateResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### getComments
Get a list of comments

```golang

 data, err :=  Feedback.GetComments(EntityType, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| EntityType | string | Type of entity, e.g. question, review or comment. | 











| xQuery | struct | Includes properties such as `ID`, `EntityID`, `UserID`, `PageID`, `PageSize`



Use this API to retrieve a list of comments for a specific entity type, e.g. products.

*Success Response:*



Success. Check the example shown below or refer `CommentGetResponse` for more details.


Schema: `CommentGetResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### checkEligibility
Checks eligibility to rate and review, and shows the cloud media configuration

```golang

 data, err :=  Feedback.CheckEligibility(EntityType, EntityID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| EntityType | string | Type of entity, e.g. question, rate, review, answer, or comment. | 


| EntityID | string | ID of the eligible entity as specified in the entity type. | 




Use this API to check whether an entity is eligible to be rated and reviewed. Moreover, it shows the cloud media configuration too.

*Success Response:*



Success. Returns a Product object. Check the example shown below or refer `CheckEligibilityResponse` for more details.


Schema: `CheckEligibilityResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### deleteMedia
Delete Media

```golang

 data, err :=  Feedback.DeleteMedia();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to delete media for an entity ID.

*Success Response:*



Success.


Schema: `UpdateResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### createMedia
Add Media

```golang

 data, err :=  Feedback.CreateMedia(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  AddMediaListRequest | "Request body" 


Use this API to add media to an entity, e.g. review.

*Success Response:*



Success. Returns media IDs.


Schema: `InsertResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### updateMedia
Update Media

```golang

 data, err :=  Feedback.UpdateMedia(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  UpdateMediaListRequest | "Request body" 


Use this API to update media (archive/approve) for an entity.

*Success Response:*



Success.


Schema: `UpdateResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### getMedias
Get Media

```golang

 data, err :=  Feedback.GetMedias(EntityType, EntityID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| EntityType | string | Type of entity, e.g. question or product. | 


| EntityID | string | ID of the eligible entity as specified in the entity type(question ID/product ID). | 







| xQuery | struct | Includes properties such as `ID`, `PageID`, `PageSize`



Use this API to retrieve all media from an entity.

*Success Response:*



Success. Check the example shown below or refer `MediaGetResponse` for more details.


Schema: `MediaGetResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### getReviewSummaries
Get a review summary

```golang

 data, err :=  Feedback.GetReviewSummaries(EntityType, EntityID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| EntityType | string | Type of entity, e.g. product, delivery, seller, order placed, order delivered, application, or template. | 


| EntityID | string | ID of the eligible entity as specified in the entity type. | 







| xQuery | struct | Includes properties such as `ID`, `PageID`, `PageSize`



Review summary gives ratings and attribute metrics of a review per entity. Use this API to retrieve the following response data: review count, rating average. 'review metrics'/'attribute rating metrics' which contains name, type, average and count.

*Success Response:*



Success. Check the example shown below or refer `ReviewMetricGetResponse` for more details.


Schema: `ReviewMetricGetResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### createReview
Add customer reviews

```golang

 data, err :=  Feedback.CreateReview(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  UpdateReviewRequest | "Request body" 


Use this API to add customer reviews for a specific entity along with the following data: attributes rating, entity rating, title, description, media resources and template ID.

*Success Response:*



Success. Returns a review ID.


Schema: `UpdateResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### updateReview
Update customer reviews

```golang

 data, err :=  Feedback.UpdateReview(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  UpdateReviewRequest | "Request body" 


Use this API to update customer reviews for a specific entity along with following data: attributes rating, entity rating, title, description, media resources and template ID.

*Success Response:*



Success.


Schema: `UpdateResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### getReviews
Get list of customer reviews

```golang

 data, err :=  Feedback.GetReviews(EntityType, EntityID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| EntityType | string | Type of entity, e.g. product, delivery, seller, l3, order placed, order delivered, application, or template. | 


| EntityID | string | ID of the eligible entity as specified in the entity type. | 



















| xQuery | struct | Includes properties such as `ID`, `UserID`, `Media`, `Rating`, `AttributeRating`, `Facets`, `Sort`, `PageID`, `PageSize`



Use this API to retrieve a list of customer reviews based on entity and filters provided.

*Success Response:*



Success. Check the example shown below or refer `ReviewGetResponse` for more details.


Schema: `ReviewGetResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### getTemplates
Get the feedback templates for a product or l3

```golang

 data, err :=  Feedback.GetTemplates(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |






| xQuery | struct | Includes properties such as `TemplateID`, `EntityID`, `EntityType`



Use this API to retrieve the details of the following feedback template. order, delivered, application, seller, order, placed, product

*Success Response:*



Success. Check the example shown below or refer `TemplateGetResponse` for more details.


Schema: `TemplateGetResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### createQuestion
Create a new question

```golang

 data, err :=  Feedback.CreateQuestion(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  CreateQNARequest | "Request body" 


Use this API to create a new question with following data- tags, text, type, choices for MCQ type questions, maximum length of answer.

*Success Response:*



Success. Returns a qna ID.


Schema: `InsertResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### updateQuestion
Update a question

```golang

 data, err :=  Feedback.UpdateQuestion(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  UpdateQNARequest | "Request body" 


Use this API to update the status of a question, its tags and its choices.

*Success Response:*



Success.


Schema: `UpdateResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### getQuestionAndAnswers
Get a list of QnA

```golang

 data, err :=  Feedback.GetQuestionAndAnswers(EntityType, EntityID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| EntityType | string | Type of entity, e.g. product, l3, etc. | 


| EntityID | string | ID of the eligible entity as specified in the entity type. | 









| xQuery | struct | Includes properties such as `ID`, `ShowAnswer`, `PageID`, `PageSize`



Use this API to retrieve a list of questions and answers for a given entity.

*Success Response:*



Success. Check the example shown below or refer `QNAGetResponse` for more details.


Schema: `QNAGetResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### getVotes
Get a list of votes

```golang

 data, err :=  Feedback.GetVotes(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |








| xQuery | struct | Includes properties such as `ID`, `RefType`, `PageNo`, `PageSize`



Use this API to retrieve a list of votes of a current logged in user. Votes can be filtered using `ref_type`, i.e. review | comment.

*Success Response:*



Success. Check the example shown below or refer `VoteResponse` for more details.


Schema: `VoteResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### createVote
Create a new vote

```golang

 data, err :=  Feedback.CreateVote(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  VoteRequest | "Request body" 


Use this API to create a new vote, where the action could be an upvote or a downvote. This is useful when you want to give a vote (say upvote) to a review (ref_type) of a product (entity_type).

*Success Response:*



Success. Returns a vote ID.


Schema: `InsertResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---


#### updateVote
Update a vote

```golang

 data, err :=  Feedback.UpdateVote(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  UpdateVoteRequest | "Request body" 


Use this API to update a vote with a new action, i.e. either an upvote or a downvote.

*Success Response:*



Success.


Schema: `UpdateResponse`








Bad request. See the error object in the response body for specific reason.


Schema: `FeedbackError`









---



---


## PosCart


#### getCart
Fetch all Items Added to  Cart

```golang

 data, err :=  PosCart.GetCart(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |








| xQuery | struct | Includes properties such as `UID`, `I`, `B`, `AssignCardID`



Get all the details of a items added to cart  by uid. If successful, returns a Cart resource in the response body specified in CartResponse

*Success Response:*



The Cart object. See example below or refer CartResponse for details


Schema: `CartResponse`









---


#### getCartLastModified
Fetch Last-Modified timestamp

```golang

 data, err :=  PosCart.GetCartLastModified(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `UID`



Fetch Last-Modified timestamp in header metadata

*Success Response:*



Fetch Last-Modified Timestamp Response






---


#### addItems
Add Items to Cart

```golang

 data, err :=  PosCart.AddItems(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `I`, `B`

| body |  AddCartRequest | "Request body" 


<p>Add Items to cart. See `AddCartRequest` in schema of request body for the list of attributes needed to add items to a cart. On successful request, returns cart response containing details of items, coupons available etc.these attributes will be fetched from the folowing api's</p>

*Success Response:*



Response of the cart object including all item details included in .the cart,coupons etc.


Schema: `AddCartResponse`


*Examples:*


Product has been added to your cart
```json
{
  "value": {
    "message": "Product has been added to your cart",
    "success": true,
    "cart": {
      "breakup_values": {
        "display": [
          {
            "display": "MRP Total",
            "key": "mrp_total",
            "value": 17486,
            "currency_code": "INR"
          },
          {
            "display": "Discount",
            "key": "discount",
            "value": -3540,
            "currency_code": "INR"
          },
          {
            "display": "Subtotal",
            "key": "subtotal",
            "value": 13946,
            "currency_code": "INR"
          },
          {
            "display": "Total",
            "key": "total",
            "value": 13946,
            "currency_code": "INR"
          }
        ],
        "raw": {
          "cod_charge": 0,
          "convenience_fee": 0,
          "coupon": 0,
          "delivery_charge": 0,
          "discount": -3540,
          "fynd_cash": 0,
          "gst_charges": 1529.96,
          "mrp_total": 17486,
          "subtotal": 13946,
          "total": 13946,
          "vog": 12416.04,
          "you_saved": 0
        },
        "loyalty_points": {
          "total": 0,
          "applicable": 0,
          "is_applied": false,
          "description": "Your cashback, referrals, and refund amount get credited to Fynd Cash which can be redeemed while placing an order."
        },
        "coupon": {
          "type": "cash",
          "code": "",
          "uid": null,
          "value": 0,
          "is_applied": false,
          "message": "Sorry! Invalid Coupon"
        }
      },
      "items": [
        {
          "key": "751083_10",
          "article": {
            "type": "article",
            "uid": "612_9_SE61201_19100302_10",
            "size": "10",
            "seller": {
              "uid": 612,
              "name": "SSR ENTERPRISE"
            },
            "store": {
              "uid": 4431,
              "name": "Motilal Nagar 1, Goregaon"
            },
            "quantity": 4,
            "price": {
              "base": {
                "marked": 3999,
                "effective": 2399,
                "currency_code": "INR"
              },
              "converted": {
                "marked": 3999,
                "effective": 2399,
                "currency_code": "INR"
              }
            }
          },
          "price": {
            "base": {
              "add_on": 4798,
              "marked": 7998,
              "effective": 4798,
              "selling": 4798,
              "currency_code": "INR"
            },
            "converted": {
              "add_on": 4798,
              "marked": 7998,
              "effective": 4798,
              "selling": 4798,
              "currency_code": "INR"
            }
          },
          "availability": {
            "sizes": [
              "10"
            ],
            "other_store_quantity": 2,
            "out_of_stock": false,
            "deliverable": true,
            "is_valid": true
          },
          "product": {
            "type": "product",
            "uid": 751083,
            "name": "Carson 2",
            "slug": "puma-carson-2-751083-6ad98d",
            "brand": {
              "uid": 9,
              "name": "Puma"
            },
            "categories": [
              {
                "uid": 165,
                "name": "Outdoor Sports Shoes"
              }
            ],
            "images": [
              {
                "aspect_ratio": "16:25",
                "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/9_19100302/1_1542807042296.jpg",
                "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/9_19100302/1_1542807042296.jpg"
              }
            ],
            "action": {
              "type": "product",
              "url": "https://api.addsale.com/platform/content/v1/products/puma-carson-2-751083-6ad98d/",
              "query": {
                "product_slug": [
                  "puma-carson-2-751083-6ad98d"
                ]
              }
            }
          },
          "coupon_message": "",
          "quantity": 2,
          "message": "",
          "bulk_offer": {},
          "discount": "41% OFF"
        },
        {
          "key": "246228_S",
          "article": {
            "type": "article",
            "uid": "46_235_TM62_M11029ONDSXNS_S",
            "size": "S",
            "seller": {
              "uid": 46,
              "name": "RELIANCE BRANDS LIMITED"
            },
            "store": {
              "uid": 4550,
              "name": "VR Mall"
            },
            "quantity": 1,
            "price": {
              "base": {
                "marked": 4490,
                "effective": 4490,
                "currency_code": "INR"
              },
              "converted": {
                "marked": 4490,
                "effective": 4490,
                "currency_code": "INR"
              }
            }
          },
          "price": {
            "base": {
              "add_on": 4490,
              "marked": 4490,
              "effective": 4490,
              "selling": 4490,
              "currency_code": "INR"
            },
            "converted": {
              "add_on": 4490,
              "marked": 4490,
              "effective": 4490,
              "selling": 4490,
              "currency_code": "INR"
            }
          },
          "availability": {
            "sizes": [
              "L",
              "M",
              "S",
              "XL",
              "XXL"
            ],
            "other_store_quantity": 0,
            "out_of_stock": false,
            "deliverable": true,
            "is_valid": true
          },
          "product": {
            "type": "product",
            "uid": 246228,
            "name": "Blue Solid T-Shirt",
            "slug": "superdry-blue-solid-t-shirt-2",
            "brand": {
              "uid": 235,
              "name": "Superdry"
            },
            "categories": [
              {
                "uid": 192,
                "name": "T-Shirts"
              }
            ],
            "images": [
              {
                "aspect_ratio": "16:25",
                "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/235_M11029ONDSXNS/1.jpg",
                "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/235_M11029ONDSXNS/1.jpg"
              }
            ],
            "action": {
              "type": "product",
              "url": "https://api.addsale.com/platform/content/v1/products/superdry-blue-solid-t-shirt-2/",
              "query": {
                "product_slug": [
                  "superdry-blue-solid-t-shirt-2"
                ]
              }
            }
          },
          "coupon_message": "",
          "quantity": 1,
          "message": "",
          "bulk_offer": {},
          "discount": ""
        },
        {
          "key": "443175_S",
          "article": {
            "type": "article",
            "uid": "162_207_1271_LJ03LBLUDN88_S",
            "size": "S",
            "seller": {
              "uid": 162,
              "name": "GO FASHION INDIA PRIVATE LIMITED"
            },
            "store": {
              "uid": 5784,
              "name": "Vega City mall"
            },
            "quantity": 3,
            "price": {
              "base": {
                "marked": 1599,
                "effective": 1599,
                "currency_code": "INR"
              },
              "converted": {
                "marked": 1599,
                "effective": 1599,
                "currency_code": "INR"
              }
            }
          },
          "price": {
            "base": {
              "add_on": 1599,
              "marked": 1599,
              "effective": 1599,
              "selling": 1599,
              "currency_code": "INR"
            },
            "converted": {
              "add_on": 1599,
              "marked": 1599,
              "effective": 1599,
              "selling": 1599,
              "currency_code": "INR"
            }
          },
          "availability": {
            "sizes": [
              "XL",
              "M",
              "L",
              "S"
            ],
            "other_store_quantity": 8,
            "out_of_stock": false,
            "deliverable": true,
            "is_valid": true
          },
          "product": {
            "type": "product",
            "uid": 443175,
            "name": "Light Blue Denim Jeggings",
            "slug": "go-colors-light-blue-denim-jeggings-443175-3c688c",
            "brand": {
              "uid": 207,
              "name": "Go Colors"
            },
            "categories": [
              {
                "uid": 267,
                "name": "Jeggings"
              }
            ],
            "images": [
              {
                "aspect_ratio": "16:25",
                "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/207_LJ03LBLUDN88/1_1512382513548.jpg",
                "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/207_LJ03LBLUDN88/1_1512382513548.jpg"
              }
            ],
            "action": {
              "type": "product",
              "url": "https://api.addsale.com/platform/content/v1/products/go-colors-light-blue-denim-jeggings-443175-3c688c/",
              "query": {
                "product_slug": [
                  "go-colors-light-blue-denim-jeggings-443175-3c688c"
                ]
              }
            }
          },
          "coupon_message": "",
          "quantity": 1,
          "message": "",
          "bulk_offer": {},
          "discount": ""
        },
        {
          "key": "778937_OS",
          "article": {
            "type": "article",
            "uid": "686_963_IC68601_PWUPC01977_OS",
            "size": "OS",
            "seller": {
              "uid": 686,
              "name": "INDUS CORPORATION"
            },
            "store": {
              "uid": 5059,
              "name": "Vidyaranyapura Main Road"
            },
            "quantity": 3,
            "price": {
              "base": {
                "marked": 3399,
                "effective": 3059,
                "currency_code": "INR"
              },
              "converted": {
                "marked": 3399,
                "effective": 3059,
                "currency_code": "INR"
              }
            }
          },
          "price": {
            "base": {
              "add_on": 3059,
              "marked": 3399,
              "effective": 3059,
              "selling": 3059,
              "currency_code": "INR"
            },
            "converted": {
              "add_on": 3059,
              "marked": 3399,
              "effective": 3059,
              "selling": 3059,
              "currency_code": "INR"
            }
          },
          "availability": {
            "sizes": [
              "OS"
            ],
            "other_store_quantity": 2,
            "out_of_stock": false,
            "deliverable": true,
            "is_valid": true
          },
          "product": {
            "type": "product",
            "uid": 778937,
            "name": "Colourful Carnival Bouncer",
            "slug": "fisher-price-colourful-carnival-bouncer-778937-fafa1f",
            "brand": {
              "uid": 963,
              "name": "Fisher-Price"
            },
            "categories": [
              {
                "uid": 576,
                "name": "Cradles"
              }
            ],
            "images": [
              {
                "aspect_ratio": "16:25",
                "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/963_PWUPC01977/1_1545308400588.jpg",
                "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/963_PWUPC01977/1_1545308400588.jpg"
              }
            ],
            "action": {
              "type": "product",
              "url": "https://api.addsale.com/platform/content/v1/products/fisher-price-colourful-carnival-bouncer-778937-fafa1f/",
              "query": {
                "product_slug": [
                  "fisher-price-colourful-carnival-bouncer-778937-fafa1f"
                ]
              }
            }
          },
          "coupon_message": "",
          "quantity": 1,
          "message": "",
          "bulk_offer": {},
          "discount": "11% OFF"
        }
      ],
      "delivery_charge_info": "",
      "coupon_text": "View all offers",
      "cart_id": 7927,
      "uid": "7927",
      "gstin": null,
      "checkout_mode": "self",
      "last_modified": "Tue, 03 Sep 2019 06:00:43 GMT",
      "restrict_checkout": false,
      "is_valid": true
    },
    "result": {}
  }
}
```

Sorry, item is out of stock
```json
{
  "value": {
    "message": "Sorry, item is out of stock",
    "success": false,
    "cart": {
      "breakup_values": {
        "raw": {
          "cod_charge": 0,
          "convenience_fee": 0,
          "coupon": 0,
          "delivery_charge": 0,
          "discount": -202000,
          "fynd_cash": 0,
          "gst_charges": 4804.71,
          "mrp_total": 302899,
          "subtotal": 100899,
          "total": 100899,
          "vog": 96094.29,
          "you_saved": 0
        },
        "coupon": {
          "type": "cash",
          "code": "",
          "uid": null,
          "value": 0,
          "is_applied": false,
          "message": "Sorry! Invalid Coupon"
        },
        "display": [
          {
            "display": "MRP Total",
            "key": "mrp_total",
            "value": 302899,
            "currency_code": "INR"
          },
          {
            "display": "Discount",
            "key": "discount",
            "value": -202000,
            "currency_code": "INR"
          },
          {
            "display": "Subtotal",
            "key": "subtotal",
            "value": 100899,
            "currency_code": "INR"
          },
          {
            "display": "Total",
            "key": "total",
            "value": 100899,
            "currency_code": "INR"
          }
        ],
        "loyalty_points": {
          "total": 0,
          "applicable": 0,
          "is_applied": false,
          "description": "Your cashback, referrals, and refund amount get credited to Fynd Cash which can be redeemed while placing an order."
        }
      },
      "items": [
        {
          "bulk_offer": {},
          "discount": "67% OFF",
          "article": {
            "type": "article",
            "uid": "604_902_SSTC60401_636BLUE_1",
            "size": "1",
            "seller": {
              "uid": 604,
              "name": "SHRI SHANTINATH TRADING COMPANY"
            },
            "store": {
              "uid": 4579,
              "name": "Gandhi Nagar"
            },
            "quantity": 108,
            "price": {
              "base": {
                "marked": 2999,
                "effective": 999,
                "currency_code": "INR"
              },
              "converted": {
                "marked": 2999,
                "effective": 999,
                "currency_code": "INR"
              }
            }
          },
          "coupon_message": "",
          "key": "707569_1",
          "availability": {
            "sizes": [
              "1",
              "8",
              "7",
              "2",
              "9",
              "5",
              "3",
              "6"
            ],
            "other_store_quantity": 7,
            "out_of_stock": false,
            "deliverable": true,
            "is_valid": true
          },
          "product": {
            "type": "product",
            "uid": 707569,
            "name": "Blue and Gold Printed Ethnic Set",
            "slug": "aj-dezines-blue-and-gold-printed-ethnic-set-707569-bff01a",
            "brand": {
              "uid": 902,
              "name": ""
            },
            "categories": [
              {
                "uid": 525,
                "name": ""
              }
            ],
            "images": [
              {
                "aspect_ratio": "16:25",
                "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/902_636BLUE/1_1540301094877.jpg",
                "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/902_636BLUE/1_1540301094877.jpg"
              }
            ],
            "action": {
              "type": "product",
              "url": "https://api.addsale.com/v1/products/aj-dezines-blue-and-gold-printed-ethnic-set-707569-bff01a/",
              "query": {
                "product_slug": [
                  "aj-dezines-blue-and-gold-printed-ethnic-set-707569-bff01a"
                ]
              }
            }
          },
          "price": {
            "base": {
              "add_on": 100899,
              "marked": 302899,
              "effective": 100899,
              "selling": 100899,
              "currency_code": "INR"
            },
            "converted": {
              "add_on": 100899,
              "marked": 302899,
              "effective": 100899,
              "selling": 100899,
              "currency_code": "INR"
            }
          },
          "message": "",
          "quantity": 101
        }
      ],
      "delivery_charge_info": "",
      "coupon_text": "View all offers",
      "cart_id": 54,
      "uid": "54",
      "gstin": null,
      "checkout_mode": "self",
      "restrict_checkout": false,
      "is_valid": false,
      "last_modified": "Tue, 03 Sep 2019 09:55:40 GMT"
    },
    "result": {}
  }
}
```









---


#### updateCart
Update Items already added to Cart

```golang

 data, err :=  PosCart.UpdateCart(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |






| xQuery | struct | Includes properties such as `UID`, `I`, `B`

| body |  UpdateCartRequest | "Request body" 


Request object containing attributes like item_quantity and item_size which can be updated .these attributes will be fetched from the folowing api's</p> <ul> <li><font color="monochrome">operation</font> Operation for current api call. <b>update_item</b> for update items. <b>remove_item</b> for removing items.</li> <li> <font color="monochrome">item_id</font>  "/platform/content/v1/products/"</li> <li> <font color="monochrome">item_size</font>   "/platform/content/v1/products/{slug}/sizes/"</li> <li> <font color="monochrome">quantity</font>  item quantity (must be greater than or equal to 1)</li> <li> <font color="monochrome">article_id</font>   "/content​/v1​/products​/{identifier}​/sizes​/price​/"</li> <li> <font color="monochrome">item_index</font>  item position in the cart (must be greater than or equal to 0)</li> </ul>

*Success Response:*



Response of the cart object including all item with their updated details included in .the cart,coupons etc..


Schema: `UpdateCartResponse`


*Examples:*


Nothing updated
```json
{
  "value": {
    "cart": {
      "breakup_values": {
        "raw": {
          "cod_charge": 0,
          "convenience_fee": 0,
          "coupon": 0,
          "delivery_charge": 0,
          "discount": -202000,
          "fynd_cash": 0,
          "gst_charges": 4804.71,
          "mrp_total": 302899,
          "subtotal": 100899,
          "total": 100899,
          "vog": 96094.29,
          "you_saved": 0
        },
        "coupon": {
          "type": "cash",
          "code": "",
          "uid": null,
          "value": 0,
          "is_applied": false,
          "message": "Sorry! Invalid Coupon"
        },
        "display": [
          {
            "display": "MRP Total",
            "key": "mrp_total",
            "value": 302899,
            "currency_code": "INR"
          },
          {
            "display": "Discount",
            "key": "discount",
            "value": -202000,
            "currency_code": "INR"
          },
          {
            "display": "Subtotal",
            "key": "subtotal",
            "value": 100899,
            "currency_code": "INR"
          },
          {
            "display": "Total",
            "key": "total",
            "value": 100899,
            "currency_code": "INR"
          }
        ],
        "loyalty_points": {
          "total": 0,
          "applicable": 0,
          "is_applied": false,
          "description": "Your cashback, referrals, and refund amount get credited to Fynd Cash which can be redeemed while placing an order."
        }
      },
      "items": [
        {
          "bulk_offer": {},
          "discount": "67% OFF",
          "article": {
            "type": "article",
            "uid": "604_902_SSTC60401_636BLUE_1",
            "size": "1",
            "seller": {
              "uid": 604,
              "name": "SHRI SHANTINATH TRADING COMPANY"
            },
            "store": {
              "uid": 4579,
              "name": "Gandhi Nagar"
            },
            "quantity": 108,
            "price": {
              "base": {
                "marked": 2999,
                "effective": 999,
                "currency_code": "INR"
              },
              "converted": {
                "marked": 2999,
                "effective": 999,
                "currency_code": "INR"
              }
            }
          },
          "coupon_message": "",
          "key": "707569_1",
          "availability": {
            "sizes": [
              "1",
              "8",
              "7",
              "2",
              "9",
              "5",
              "3",
              "6"
            ],
            "other_store_quantity": 7,
            "out_of_stock": false,
            "deliverable": true,
            "is_valid": true
          },
          "product": {
            "type": "product",
            "uid": 707569,
            "name": "Blue and Gold Printed Ethnic Set",
            "slug": "aj-dezines-blue-and-gold-printed-ethnic-set-707569-bff01a",
            "brand": {
              "uid": 902,
              "name": ""
            },
            "categories": [
              {
                "uid": 525,
                "name": ""
              }
            ],
            "images": [
              {
                "aspect_ratio": "16:25",
                "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/902_636BLUE/1_1540301094877.jpg",
                "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/902_636BLUE/1_1540301094877.jpg"
              }
            ],
            "action": {
              "type": "product",
              "url": "https://api.addsale.com/v1/products/aj-dezines-blue-and-gold-printed-ethnic-set-707569-bff01a/",
              "query": {
                "product_slug": [
                  "aj-dezines-blue-and-gold-printed-ethnic-set-707569-bff01a"
                ]
              }
            }
          },
          "price": {
            "base": {
              "add_on": 100899,
              "marked": 302899,
              "effective": 100899,
              "selling": 100899,
              "currency_code": "INR"
            },
            "converted": {
              "add_on": 100899,
              "marked": 302899,
              "effective": 100899,
              "selling": 100899,
              "currency_code": "INR"
            }
          },
          "message": "",
          "quantity": 101
        }
      ],
      "delivery_charge_info": "",
      "coupon_text": "View all offers",
      "cart_id": 54,
      "uid": "54",
      "gstin": null,
      "checkout_mode": "self",
      "restrict_checkout": false,
      "is_valid": true,
      "last_modified": "Tue, 03 Sep 2019 10:19:20 GMT"
    },
    "result": {
      "707569_90": {
        "success": true,
        "message": "Nothing updated"
      }
    },
    "message": "Nothing updated",
    "success": true
  }
}
```

Item updated in the cart
```json
{
  "value": {
    "cart": {
      "breakup_values": {
        "coupon": {
          "type": "cash",
          "code": "",
          "uid": null,
          "value": 0,
          "is_applied": false,
          "message": "Sorry! Invalid Coupon"
        },
        "loyalty_points": {
          "total": 0,
          "applicable": 0,
          "is_applied": false,
          "description": "Your cashback, referrals, and refund amount get credited to Fynd Cash which can be redeemed while placing an order."
        },
        "raw": {
          "cod_charge": 0,
          "convenience_fee": 0,
          "coupon": 0,
          "delivery_charge": 0,
          "discount": 0,
          "fynd_cash": 0,
          "gst_charges": 838.83,
          "mrp_total": 5499,
          "subtotal": 5499,
          "total": 5499,
          "vog": 4660.17,
          "you_saved": 0
        },
        "display": [
          {
            "display": "MRP Total",
            "key": "mrp_total",
            "value": 5499,
            "currency_code": "INR"
          },
          {
            "display": "Subtotal",
            "key": "subtotal",
            "value": 5499,
            "currency_code": "INR"
          },
          {
            "display": "Total",
            "key": "total",
            "value": 5499,
            "currency_code": "INR"
          }
        ]
      },
      "items": [
        {
          "key": "437414_7",
          "message": "",
          "bulk_offer": {},
          "price": {
            "base": {
              "add_on": 5499,
              "marked": 5499,
              "effective": 5499,
              "selling": 5499,
              "currency_code": "INR"
            },
            "converted": {
              "add_on": 5499,
              "marked": 5499,
              "effective": 5499,
              "selling": 5499,
              "currency_code": "INR"
            }
          },
          "quantity": 1,
          "discount": "",
          "product": {
            "type": "product",
            "uid": 437414,
            "name": "Suede Classic",
            "slug": "puma-suede-classic-437414-6e6bbf",
            "brand": {
              "uid": 9,
              "name": "Puma"
            },
            "categories": [
              {
                "uid": 165,
                "name": "Outdoor Sports Shoes"
              }
            ],
            "images": [
              {
                "aspect_ratio": "16:25",
                "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/9_35656851/1_1511171811830.jpg",
                "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/9_35656851/1_1511171811830.jpg"
              }
            ],
            "action": {
              "type": "product",
              "url": "https://api.addsale.com/platform/content/v1/products/puma-suede-classic-437414-6e6bbf/",
              "query": {
                "product_slug": [
                  "puma-suede-classic-437414-6e6bbf"
                ]
              }
            }
          },
          "article": {
            "type": "article",
            "uid": "507_9_96099_35656851_7",
            "size": "7",
            "seller": {
              "uid": 507,
              "name": "PUMA SPORTS INDIA PVT LTD"
            },
            "store": {
              "uid": 3632,
              "name": "Colaba Causway"
            },
            "quantity": 5,
            "price": {
              "base": {
                "marked": 5499,
                "effective": 5499,
                "currency_code": "INR"
              },
              "converted": {
                "marked": 5499,
                "effective": 5499,
                "currency_code": "INR"
              }
            }
          },
          "coupon_message": "",
          "availability": {
            "sizes": [
              "10",
              "11",
              "6",
              "9",
              "7",
              "8"
            ],
            "other_store_quantity": 22,
            "out_of_stock": false,
            "deliverable": true,
            "is_valid": true
          }
        }
      ],
      "delivery_charge_info": "",
      "coupon_text": "View all offers",
      "cart_id": 12426,
      "uid": "12426",
      "gstin": null,
      "checkout_mode": "self",
      "last_modified": "Thu, 22 Aug 2019 04:51:42 GMT",
      "restrict_checkout": false,
      "is_valid": true
    },
    "result": {
      "437414_7": {
        "success": true,
        "message": "Item updated in the bag"
      }
    },
    "message": "Item updated in the bag",
    "success": true
  }
}
```









---


#### getItemCount
Cart item count

```golang

 data, err :=  PosCart.GetItemCount(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `UID`



Get total count of item present in cart

*Success Response:*



OK


Schema: `CartItemCountResponse`









---


#### getCoupons
Fetch Coupon

```golang

 data, err :=  PosCart.GetCoupons(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `UID`



Get all the details of a coupons applicable to cart  by uid. If successful, returns a Coupon resource in the response body specified in GetCouponResponse

*Success Response:*



Returns The Couppon object which has list of all available_coupon applicale for the cart. See example below or refer GetCouponResponse for details


Schema: `GetCouponResponse`









---


#### applyCoupon
Apply Coupon

```golang

 data, err :=  PosCart.ApplyCoupon(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |








| xQuery | struct | Includes properties such as `I`, `B`, `P`, `UID`

| body |  ApplyCouponRequest | "Request body" 


<p>Apply Coupons on Items added to cart. On successful request, returns cart response containing details of items ,coupons applied etc.these attributes will be consumed by  api</p> <ul> <li> <font color="monochrome">coupon_code</font></li>
</ul>

*Success Response:*



Response of the Coupon object including all item details included in .the cart,coupons applied etc.


Schema: `CartResponse`









---


#### removeCoupon
Remove Coupon Applied

```golang

 data, err :=  PosCart.RemoveCoupon(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `UID`



Remove Coupon applied on the cart by passing uid in request body.

*Success Response:*



Response of the Coupon object including all item details included in .the cart,coupons removed etc.


Schema: `CartResponse`









---


#### getBulkDiscountOffers
Get discount offers based on quantity

```golang

 data, err :=  PosCart.GetBulkDiscountOffers(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |








| xQuery | struct | Includes properties such as `ItemID`, `ArticleID`, `UID`, `Slug`



List applicable offers along with current, next and best offer for given product. Either one of **uid**, **item_id**, **slug** should be present*

*Success Response:*



Offers found or not found with valid input


Schema: `BulkPriceResponse`


*Examples:*


Offers found
```json
{
  "value": {
    "data": [
      {
        "seller": {
          "uid": 248,
          "name": "MANYAVAR CREATIONS PRIVATE LIMITED"
        },
        "offers": [
          {
            "quantity": 1,
            "auto_applied": true,
            "margin": 10,
            "type": "bundle",
            "price": {
              "marked": 3999,
              "effective": 3999,
              "bulk_effective": 3599.1,
              "currency_code": "INR"
            },
            "total": 3599.1
          },
          {
            "quantity": 3,
            "auto_applied": true,
            "margin": 20,
            "type": "bundle",
            "price": {
              "marked": 3999,
              "effective": 3999,
              "bulk_effective": 3199.2,
              "currency_code": "INR"
            },
            "total": 9597.6
          },
          {
            "quantity": 9,
            "auto_applied": true,
            "margin": 30,
            "type": "bundle",
            "price": {
              "marked": 3999,
              "effective": 3999,
              "bulk_effective": 3443.4444444444,
              "currency_code": "INR"
            },
            "total": 30991,
            "best": true
          }
        ]
      }
    ]
  }
}
```

Offers not found
```json
{
  "value": {
    "data": []
  }
}
```








Unhandled api error


Schema: `Object`









---


#### getAddresses
Fetch Address

```golang

 data, err :=  PosCart.GetAddresses(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |










| xQuery | struct | Includes properties such as `UID`, `MobileNo`, `CheckoutMode`, `Tags`, `IsDefault`



Get all the addresses associated with the account. If successful, returns a Address resource in the response body specified in GetAddressesResponse.attibutes listed below are optional <ul> <li> <font color="monochrome">uid</font></li> <li> <font color="monochrome">address_id</font></li> <li> <font color="monochrome">mobile_no</font></li> <li> <font color="monochrome">checkout_mode</font></li> <li> <font color="monochrome">tags</font></li> <li> <font color="monochrome">default</font></li> </ul>

*Success Response:*



Returns The Address object which has list of all address saved for the account. See example below or refer GetAddressesResponse for details


Schema: `GetAddressesResponse`









---


#### addAddress
Add Address to the account

```golang

 data, err :=  PosCart.AddAddress(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  Address | "Request body" 


<p>Add Address to account. See `Address` in schema of request body for the list of attributes needed to add Address to account. On successful request, returns response containing address_id ,is_default_address and success message.

*Success Response:*



Return Address Id on successfull completion of the request.


Schema: `SaveAddressResponse`









---


#### getAddressById
Fetch Single Address

```golang

 data, err :=  PosCart.GetAddressById(ID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ID | float64 |  | 











| xQuery | struct | Includes properties such as `UID`, `MobileNo`, `CheckoutMode`, `Tags`, `IsDefault`



Get a addresses with the given id. If successful, returns a Address resource in the response body specified in `Address`.attibutes listed below are optional <ul> <li> <font color="monochrome">mobile_no</font></li> <li> <font color="monochrome">checkout_mode</font></li> <li> <font color="monochrome">tags</font></li> <li> <font color="monochrome">default</font></li> </ul>

*Success Response:*



Returns The Address object which has list of all address saved for the account. See example below or refer Address for details


Schema: `Address`









---


#### updateAddress
Update Address alreay added to account

```golang

 data, err :=  PosCart.UpdateAddress(ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ID | float64 | Address id | 


| body |  Address | "Request body" 


Request object containing attributes mentioned in  <font color="blue">Address </font> can be updated .these attributes are :</p> <ul> <li> <font color="monochrome">is_default_address</font></li> <li> <font color="monochrome">landmark</font></li> <li> <font color="monochrome">area</font></li> <li> <font color="monochrome">pincode</font></li> <li> <font color="monochrome">email</font></li> <li> <font color="monochrome">address_type</font></li> <li> <font color="monochrome">name</font></li> <li> <font color="monochrome">address_id</font></li> <li> <font color="monochrome">address</font></li> </ul>

*Success Response:*



Response of the Address object containing address_id and sucess message.


Schema: `UpdateAddressResponse`









---


#### removeAddress
Remove Address Associated to the account

```golang

 data, err :=  PosCart.RemoveAddress(ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ID | float64 | Address id | 




Delete a Address by it's address_id. Returns an object that tells whether the address was deleted successfully

*Success Response:*



Status object. Tells whether the operation was successful. See example below or refer DeleteAddressResponse


Schema: `DeleteAddressResponse`









---


#### selectAddress
Select Address from All Addresses

```golang

 data, err :=  PosCart.SelectAddress(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |






| xQuery | struct | Includes properties such as `UID`, `I`, `B`

| body |  SelectCartAddressRequest | "Request body" 


<p>Select Address from all addresses associated with the account in order to ship the cart items to .that address,otherwise default address will be selected implicitly. See `SelectCartAddressRequest` in schema of request body for the list of attributes needed to select Address from account. On successful request, returns Cart object response.below are the address attributes which needs to be sent. <ul> <li> <font color="monochrome">address_id</font></li> <li> <font color="monochrome">billing_address_id</font></li> <li> <font color="monochrome">uid</font></li> </ul>

*Success Response:*



Response of the Address object containing Cart Object and success message.  .


Schema: `CartResponse`








Address or Pincode Error


Schema: `Object`


*Examples:*


Address Not Found
```json
{
  "value": {
    "status": "ERROR",
    "message": "ADDRESS_NOT_FOUND"
  }
}
```

Pincode Not Serviciable
```json
{
  "value": {
    "status": "ERROR",
    "message": "PINCODE_NOT_SERVICIABLE"
  }
}
```









---


#### selectPaymentMode
Update Cart Payment

```golang

 data, err :=  PosCart.SelectPaymentMode(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `UID`

| body |  UpdateCartPaymentRequest | "Request body" 


Update Cart Payment for Your Account

*Success Response:*



Cart response with payment options


Schema: `CartResponse`









---


#### validateCouponForPayment
Get Cart Payment for valid coupon

```golang

 data, err :=  PosCart.ValidateCouponForPayment(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |












| xQuery | struct | Includes properties such as `UID`, `AddressID`, `PaymentMode`, `PaymentIdentifier`, `AggregatorName`, `MerchantCode`



Validate coupon for selected payment mode

*Success Response:*



Cart Affiliates.


Schema: `PaymentCouponValidate`









---


#### getShipments
Get delivery date and options before checkout

```golang

 data, err :=  PosCart.GetShipments(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |














| xQuery | struct | Includes properties such as `PickAtStoreUID`, `OrderingStoreID`, `P`, `UID`, `AddressID`, `AreaCode`, `OrderType`



Shipment break up item wise with delivery date. Actual                      delivery will be during given dates only. Items will be                      delivered in group of shipments created.

*Success Response:*



OK


Schema: `CartShipmentsResponse`


*Examples:*


Shipment Generated
```json
{
  "value": {
    "items": [],
    "cart_id": 7501,
    "uid": "7501",
    "success": true,
    "error_message": "Note: Your order delivery will be delayed by 7-10 Days",
    "payment_options": {
      "payment_option": [
        {
          "name": "COD",
          "display_name": "Cash on Delivery",
          "display_priority": 1,
          "payment_mode_id": 11,
          "logo": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png",
          "logo_url": {
            "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png",
            "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png"
          },
          "list": []
        },
        {
          "name": "CARD",
          "display_priority": 2,
          "payment_mode_id": 2,
          "display_name": "Card",
          "list": []
        },
        {
          "name": "NB",
          "display_priority": 3,
          "payment_mode_id": 3,
          "display_name": "Net Banking",
          "list": [
            {
              "aggregator_name": "Razorpay",
              "bank_name": "ICICI Bank",
              "bank_code": "ICIC",
              "url": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png"
              },
              "merchant_code": "NB_ICICI",
              "display_priority": 1
            }
          ]
        },
        {
          "name": "WL",
          "display_priority": 4,
          "payment_mode_id": 4,
          "display_name": "Wallet",
          "list": [
            {
              "wallet_name": "Paytm",
              "wallet_code": "paytm",
              "wallet_id": 4,
              "merchant_code": "PAYTM",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/paytm_logo_small.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/paytm_logo_large.png"
              },
              "aggregator_name": "Juspay",
              "display_priority": 1
            }
          ]
        },
        {
          "name": "UPI",
          "display_priority": 9,
          "payment_mode_id": 6,
          "display_name": "UPI",
          "list": [
            {
              "aggregator_name": "UPI_Razorpay",
              "name": "UPI",
              "display_name": "BHIM UPI",
              "code": "UPI",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/upi_100x78.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/upi_150x100.png"
              },
              "merchant_code": "UPI",
              "timeout": 240,
              "retry_count": 0,
              "fynd_vpa": "shopsense.rzp@hdfcbank",
              "intent_flow": true,
              "intent_app_error_list": [
                "com.csam.icici.bank.imobile",
                "in.org.npci.upiapp",
                "com.whatsapp"
              ]
            }
          ]
        },
        {
          "name": "PL",
          "display_priority": 11,
          "payment_mode_id": 1,
          "display_name": "Pay Later",
          "list": [
            {
              "aggregator_name": "Simpl",
              "name": "Simpl",
              "code": "simpl",
              "merchant_code": "SIMPL",
              "logo": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png"
              }
            }
          ]
        }
      ],
      "payment_flows": {
        "Simpl": {
          "data": {
            "gateway": {
              "route": "simpl",
              "entity": "sdk",
              "is_customer_validation_required": true,
              "cust_validation_url": "https://api.addsale.com/gringotts/api/v1/validate-customer/",
              "sdk": {
                "config": {
                  "redirect": false,
                  "callback_url": null,
                  "action_url": "https://api.addsale.com/avis/api/v1/payments/charge-gringotts-transaction/"
                },
                "data": {
                  "user_phone": "8452996729",
                  "user_email": "paymentsdummy@gofynd.com"
                }
              },
              "return_url": null
            }
          },
          "api_link": "",
          "payment_flow": "sdk"
        },
        "Juspay": {
          "data": {},
          "api_link": "https://sandbox.juspay.in/txns",
          "payment_flow": "api"
        },
        "Razorpay": {
          "data": {},
          "api_link": "",
          "payment_flow": "sdk"
        },
        "UPI_Razorpay": {
          "data": {},
          "api_link": "https://api.addsale.com/gringotts/api/v1/external/payment-initialisation/",
          "payment_flow": "api"
        },
        "Fynd": {
          "data": {},
          "api_link": "",
          "payment_flow": "api"
        }
      },
      "default": {}
    },
    "user_type": "Store User",
    "cod_charges": 0,
    "order_id": null,
    "cod_available": true,
    "cod_message": "No additional COD charges applicable",
    "delivery_charges": 0,
    "delivery_charge_order_value": 0,
    "delivery_slots": [
      {
        "date": "Sat, 24 Aug",
        "delivery_slot": [
          {
            "delivery_slot_timing": "By 9:00 PM",
            "default": true,
            "delivery_slot_id": 1
          }
        ]
      }
    ],
    "store_code": "",
    "store_emps": [],
    "breakup_values": {
      "loyalty_points": {
        "total": 0,
        "applicable": 0,
        "is_applied": false,
        "description": "Your cashback, referrals, and refund amount get credited to Fynd Cash which can be redeemed while placing an order."
      },
      "coupon": {
        "type": "cash",
        "code": "",
        "uid": null,
        "value": 0,
        "is_applied": false,
        "message": "Sorry! Invalid Coupon"
      },
      "raw": {
        "cod_charge": 0,
        "convenience_fee": 0,
        "coupon": 0,
        "delivery_charge": 0,
        "discount": 0,
        "fynd_cash": 0,
        "gst_charges": 214.18,
        "mrp_total": 1999,
        "subtotal": 1999,
        "total": 1999,
        "vog": 1784.82,
        "you_saved": 0
      },
      "display": [
        {
          "display": "MRP Total",
          "key": "mrp_total",
          "value": 1999,
          "currency_code": "INR"
        },
        {
          "display": "Subtotal",
          "key": "subtotal",
          "value": 1999,
          "currency_code": "INR"
        },
        {
          "display": "Total",
          "key": "total",
          "value": 1999,
          "currency_code": "INR"
        }
      ]
    },
    "shipments": [
      {
        "fulfillment_id": 3009,
        "shipment_type": "single_shipment",
        "fulfillment_type": "store",
        "dp_id": "29",
        "dp_options": {
          "4": {
            "f_priority": 4,
            "r_priority": 5,
            "is_cod": true,
            "is_prepaid": true,
            "is_reverse": true
          },
          "7": {
            "f_priority": 3,
            "r_priority": 4,
            "is_cod": true,
            "is_prepaid": true,
            "is_reverse": true
          },
          "29": {
            "f_priority": 1,
            "r_priority": 2,
            "is_cod": true,
            "is_prepaid": true,
            "is_reverse": true
          }
        },
        "promise": {
          "timestamp": {
            "min": 1566678108,
            "max": 1567023708
          },
          "formatted": {
            "min": "Aug 24",
            "max": "Aug 28"
          }
        },
        "box_type": "Small Courier bag",
        "shipments": 1,
        "items": [
          {
            "quantity": 1,
            "product": {
              "type": "product",
              "uid": 820312,
              "name": "Navy Blue Melange Shorts",
              "slug": "883-police-navy-blue-melange-shorts-820312-4943a8",
              "brand": {
                "uid": 610,
                "name": "883 Police"
              },
              "categories": [
                {
                  "uid": 193,
                  "name": "Shorts"
                }
              ],
              "images": [
                {
                  "aspect_ratio": "16:25",
                  "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/610_SPIRAL19ANAVY/1_1549105947281.jpg",
                  "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/610_SPIRAL19ANAVY/1_1549105947281.jpg"
                }
              ],
              "action": {
                "type": "product",
                "url": "https://api.addsale.com/platform/content/v1/products/883-police-navy-blue-melange-shorts-820312-4943a8/",
                "query": {
                  "product_slug": [
                    "883-police-navy-blue-melange-shorts-820312-4943a8"
                  ]
                }
              }
            },
            "discount": "",
            "bulk_offer": {},
            "key": "820312_L",
            "price": {
              "base": {
                "add_on": 1999,
                "marked": 1999,
                "effective": 1999,
                "selling": 1999,
                "currency_code": "INR"
              },
              "converted": {
                "add_on": 1999,
                "marked": 1999,
                "effective": 1999,
                "selling": 1999,
                "currency_code": "INR"
              }
            },
            "article": {
              "type": "article",
              "uid": "381_610_IGPL01_SPIRAL19ANAVY_L",
              "size": "L",
              "seller": {
                "uid": 381,
                "name": "INTERSOURCE GARMENTS PVT LTD"
              },
              "store": {
                "uid": 3009,
                "name": "Kormangala"
              },
              "quantity": 2,
              "price": {
                "base": {
                  "marked": 1999,
                  "effective": 1999,
                  "currency_code": "INR"
                },
                "converted": {
                  "marked": 1999,
                  "effective": 1999,
                  "currency_code": "INR"
                }
              }
            },
            "availability": {
              "sizes": [
                "L",
                "XL",
                "XXL"
              ],
              "other_store_quantity": 1,
              "out_of_stock": false,
              "deliverable": true,
              "is_valid": true
            },
            "coupon_message": "",
            "message": ""
          }
        ]
      }
    ],
    "delivery_charge_info": "",
    "coupon_text": "View all offers",
    "gstin": null,
    "checkout_mode": "self",
    "last_modified": "Thu, 22 Aug 2019 20:21:48 GMT",
    "restrict_checkout": false,
    "is_valid": true
  }
}
```

Shipment Generation Failed
```json
{
  "value": {
    "items": [],
    "cart_id": 7501,
    "uid": "7501",
    "success": true,
    "error_message": "Note: Your order delivery will be delayed by 7-10 Days",
    "payment_options": {
      "payment_option": [
        {
          "name": "COD",
          "display_name": "Cash on Delivery",
          "display_priority": 1,
          "payment_mode_id": 11,
          "logo": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png",
          "logo_url": {
            "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png",
            "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png"
          },
          "list": []
        },
        {
          "name": "CARD",
          "display_priority": 2,
          "payment_mode_id": 2,
          "display_name": "Card",
          "list": []
        },
        {
          "name": "NB",
          "display_priority": 3,
          "payment_mode_id": 3,
          "display_name": "Net Banking",
          "list": [
            {
              "aggregator_name": "Razorpay",
              "bank_name": "ICICI Bank",
              "bank_code": "ICIC",
              "url": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png"
              },
              "merchant_code": "NB_ICICI",
              "display_priority": 1
            }
          ]
        },
        {
          "name": "WL",
          "display_priority": 4,
          "payment_mode_id": 4,
          "display_name": "Wallet",
          "list": [
            {
              "wallet_name": "Paytm",
              "wallet_code": "paytm",
              "wallet_id": 4,
              "merchant_code": "PAYTM",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/paytm_logo_small.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/paytm_logo_large.png"
              },
              "aggregator_name": "Juspay",
              "display_priority": 1
            }
          ]
        },
        {
          "name": "UPI",
          "display_priority": 9,
          "payment_mode_id": 6,
          "display_name": "UPI",
          "list": [
            {
              "aggregator_name": "UPI_Razorpay",
              "name": "UPI",
              "display_name": "BHIM UPI",
              "code": "UPI",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/upi_100x78.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/upi_150x100.png"
              },
              "merchant_code": "UPI",
              "timeout": 240,
              "retry_count": 0,
              "fynd_vpa": "shopsense.rzp@hdfcbank",
              "intent_flow": true,
              "intent_app_error_list": [
                "com.csam.icici.bank.imobile",
                "in.org.npci.upiapp",
                "com.whatsapp"
              ]
            }
          ]
        },
        {
          "name": "PL",
          "display_priority": 11,
          "payment_mode_id": 1,
          "display_name": "Pay Later",
          "list": [
            {
              "aggregator_name": "Simpl",
              "name": "Simpl",
              "code": "simpl",
              "merchant_code": "SIMPL",
              "logo": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png"
              }
            }
          ]
        }
      ],
      "payment_flows": {
        "Simpl": {
          "data": {
            "gateway": {
              "route": "simpl",
              "entity": "sdk",
              "is_customer_validation_required": true,
              "cust_validation_url": "https://api.addsale.com/gringotts/api/v1/validate-customer/",
              "sdk": {
                "config": {
                  "redirect": false,
                  "callback_url": null,
                  "action_url": "https://api.addsale.com/avis/api/v1/payments/charge-gringotts-transaction/"
                },
                "data": {
                  "user_phone": "8452996729",
                  "user_email": "paymentsdummy@gofynd.com"
                }
              },
              "return_url": null
            }
          },
          "api_link": "",
          "payment_flow": "sdk"
        },
        "Juspay": {
          "data": {},
          "api_link": "https://sandbox.juspay.in/txns",
          "payment_flow": "api"
        },
        "Razorpay": {
          "data": {},
          "api_link": "",
          "payment_flow": "sdk"
        },
        "UPI_Razorpay": {
          "data": {},
          "api_link": "https://api.addsale.com/gringotts/api/v1/external/payment-initialisation/",
          "payment_flow": "api"
        },
        "Fynd": {
          "data": {},
          "api_link": "",
          "payment_flow": "api"
        }
      },
      "default": {}
    },
    "user_type": "Store User",
    "cod_charges": 0,
    "order_id": null,
    "cod_available": true,
    "cod_message": "No additional COD charges applicable",
    "delivery_charges": 0,
    "delivery_charge_order_value": 0,
    "delivery_slots": [
      {
        "date": "Sat, 24 Aug",
        "delivery_slot": [
          {
            "delivery_slot_timing": "By 9:00 PM",
            "default": true,
            "delivery_slot_id": 1
          }
        ]
      }
    ],
    "store_code": "",
    "store_emps": [],
    "breakup_values": {
      "loyalty_points": {
        "total": 0,
        "applicable": 0,
        "is_applied": false,
        "description": "Your cashback, referrals, and refund amount get credited to Fynd Cash which can be redeemed while placing an order."
      },
      "coupon": {
        "type": "cash",
        "code": "",
        "uid": null,
        "value": 0,
        "is_applied": false,
        "message": "Sorry! Invalid Coupon"
      },
      "raw": {
        "cod_charge": 0,
        "convenience_fee": 0,
        "coupon": 0,
        "delivery_charge": 0,
        "discount": 0,
        "fynd_cash": 0,
        "gst_charges": 214.18,
        "mrp_total": 1999,
        "subtotal": 1999,
        "total": 1999,
        "vog": 1784.82,
        "you_saved": 0
      },
      "display": [
        {
          "display": "MRP Total",
          "key": "mrp_total",
          "value": 1999,
          "currency_code": "INR"
        },
        {
          "display": "Subtotal",
          "key": "subtotal",
          "value": 1999,
          "currency_code": "INR"
        },
        {
          "display": "Total",
          "key": "total",
          "value": 1999,
          "currency_code": "INR"
        }
      ]
    },
    "shipments": [],
    "message": "Shipments could not be generated. Please Try again after some time.",
    "delivery_charge_info": "",
    "coupon_text": "View all offers",
    "gstin": null,
    "checkout_mode": "self",
    "last_modified": "Thu, 22 Aug 2019 20:21:48 GMT",
    "restrict_checkout": false,
    "is_valid": false
  }
}
```








Unhandled api error


Schema: `Object`









---


#### updateShipments
Update shipment delivery type and quantity before checkout

```golang

 data, err :=  PosCart.UpdateShipments(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |










| xQuery | struct | Includes properties such as `I`, `P`, `UID`, `AddressID`, `OrderType`

| body |  UpdateCartShipmentRequest | "Request body" 


Shipment break up item wise with delivery date. Actual                      delivery will be during given dates only. Items will be                      delivered in group of shipments created. Update the shipment                      type and quantity as per customer preference for store pick up or home delivery

*Success Response:*



OK


Schema: `CartShipmentsResponse`


*Examples:*


Shipment Generated
```json
{
  "value": {
    "items": [],
    "cart_id": 7501,
    "uid": "7501",
    "success": true,
    "error_message": "Note: Your order delivery will be delayed by 7-10 Days",
    "payment_options": {
      "payment_option": [
        {
          "name": "COD",
          "display_name": "Cash on Delivery",
          "display_priority": 1,
          "payment_mode_id": 11,
          "logo": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png",
          "logo_url": {
            "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png",
            "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png"
          },
          "list": []
        },
        {
          "name": "CARD",
          "display_priority": 2,
          "payment_mode_id": 2,
          "display_name": "Card",
          "list": []
        },
        {
          "name": "NB",
          "display_priority": 3,
          "payment_mode_id": 3,
          "display_name": "Net Banking",
          "list": [
            {
              "aggregator_name": "Razorpay",
              "bank_name": "ICICI Bank",
              "bank_code": "ICIC",
              "url": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png"
              },
              "merchant_code": "NB_ICICI",
              "display_priority": 1
            }
          ]
        },
        {
          "name": "WL",
          "display_priority": 4,
          "payment_mode_id": 4,
          "display_name": "Wallet",
          "list": [
            {
              "wallet_name": "Paytm",
              "wallet_code": "paytm",
              "wallet_id": 4,
              "merchant_code": "PAYTM",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/paytm_logo_small.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/paytm_logo_large.png"
              },
              "aggregator_name": "Juspay",
              "display_priority": 1
            }
          ]
        },
        {
          "name": "UPI",
          "display_priority": 9,
          "payment_mode_id": 6,
          "display_name": "UPI",
          "list": [
            {
              "aggregator_name": "UPI_Razorpay",
              "name": "UPI",
              "display_name": "BHIM UPI",
              "code": "UPI",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/upi_100x78.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/upi_150x100.png"
              },
              "merchant_code": "UPI",
              "timeout": 240,
              "retry_count": 0,
              "fynd_vpa": "shopsense.rzp@hdfcbank",
              "intent_flow": true,
              "intent_app_error_list": [
                "com.csam.icici.bank.imobile",
                "in.org.npci.upiapp",
                "com.whatsapp"
              ]
            }
          ]
        },
        {
          "name": "PL",
          "display_priority": 11,
          "payment_mode_id": 1,
          "display_name": "Pay Later",
          "list": [
            {
              "aggregator_name": "Simpl",
              "name": "Simpl",
              "code": "simpl",
              "merchant_code": "SIMPL",
              "logo": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png"
              }
            }
          ]
        }
      ],
      "payment_flows": {
        "Simpl": {
          "data": {
            "gateway": {
              "route": "simpl",
              "entity": "sdk",
              "is_customer_validation_required": true,
              "cust_validation_url": "https://api.addsale.com/gringotts/api/v1/validate-customer/",
              "sdk": {
                "config": {
                  "redirect": false,
                  "callback_url": null,
                  "action_url": "https://api.addsale.com/avis/api/v1/payments/charge-gringotts-transaction/"
                },
                "data": {
                  "user_phone": "8452996729",
                  "user_email": "paymentsdummy@gofynd.com"
                }
              },
              "return_url": null
            }
          },
          "api_link": "",
          "payment_flow": "sdk"
        },
        "Juspay": {
          "data": {},
          "api_link": "https://sandbox.juspay.in/txns",
          "payment_flow": "api"
        },
        "Razorpay": {
          "data": {},
          "api_link": "",
          "payment_flow": "sdk"
        },
        "UPI_Razorpay": {
          "data": {},
          "api_link": "https://api.addsale.com/gringotts/api/v1/external/payment-initialisation/",
          "payment_flow": "api"
        },
        "Fynd": {
          "data": {},
          "api_link": "",
          "payment_flow": "api"
        }
      },
      "default": {}
    },
    "user_type": "Store User",
    "cod_charges": 0,
    "order_id": null,
    "cod_available": true,
    "cod_message": "No additional COD charges applicable",
    "delivery_charges": 0,
    "delivery_charge_order_value": 0,
    "delivery_slots": [
      {
        "date": "Sat, 24 Aug",
        "delivery_slot": [
          {
            "delivery_slot_timing": "By 9:00 PM",
            "default": true,
            "delivery_slot_id": 1
          }
        ]
      }
    ],
    "store_code": "",
    "store_emps": [],
    "breakup_values": {
      "loyalty_points": {
        "total": 0,
        "applicable": 0,
        "is_applied": false,
        "description": "Your cashback, referrals, and refund amount get credited to Fynd Cash which can be redeemed while placing an order."
      },
      "coupon": {
        "type": "cash",
        "code": "",
        "uid": null,
        "value": 0,
        "is_applied": false,
        "message": "Sorry! Invalid Coupon"
      },
      "raw": {
        "cod_charge": 0,
        "convenience_fee": 0,
        "coupon": 0,
        "delivery_charge": 0,
        "discount": 0,
        "fynd_cash": 0,
        "gst_charges": 214.18,
        "mrp_total": 1999,
        "subtotal": 1999,
        "total": 1999,
        "vog": 1784.82,
        "you_saved": 0
      },
      "display": [
        {
          "display": "MRP Total",
          "key": "mrp_total",
          "value": 1999,
          "currency_code": "INR"
        },
        {
          "display": "Subtotal",
          "key": "subtotal",
          "value": 1999,
          "currency_code": "INR"
        },
        {
          "display": "Total",
          "key": "total",
          "value": 1999,
          "currency_code": "INR"
        }
      ]
    },
    "shipments": [
      {
        "fulfillment_id": 3009,
        "shipment_type": "single_shipment",
        "fulfillment_type": "store",
        "dp_id": "29",
        "order_type": "PickAtStore",
        "dp_options": {
          "4": {
            "f_priority": 4,
            "r_priority": 5,
            "is_cod": true,
            "is_prepaid": true,
            "is_reverse": true
          },
          "7": {
            "f_priority": 3,
            "r_priority": 4,
            "is_cod": true,
            "is_prepaid": true,
            "is_reverse": true
          },
          "29": {
            "f_priority": 1,
            "r_priority": 2,
            "is_cod": true,
            "is_prepaid": true,
            "is_reverse": true
          }
        },
        "promise": {
          "timestamp": {
            "min": 1566678108,
            "max": 1567023708
          },
          "formatted": {
            "min": "Aug 24",
            "max": "Aug 28"
          }
        },
        "box_type": "Small Courier bag",
        "shipments": 1,
        "items": [
          {
            "quantity": 1,
            "product": {
              "type": "product",
              "uid": 820312,
              "name": "Navy Blue Melange Shorts",
              "slug": "883-police-navy-blue-melange-shorts-820312-4943a8",
              "brand": {
                "uid": 610,
                "name": "883 Police"
              },
              "categories": [
                {
                  "uid": 193,
                  "name": "Shorts"
                }
              ],
              "images": [
                {
                  "aspect_ratio": "16:25",
                  "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/610_SPIRAL19ANAVY/1_1549105947281.jpg",
                  "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/610_SPIRAL19ANAVY/1_1549105947281.jpg"
                }
              ],
              "action": {
                "type": "product",
                "url": "https://api.addsale.com/platform/content/v1/products/883-police-navy-blue-melange-shorts-820312-4943a8/",
                "query": {
                  "product_slug": [
                    "883-police-navy-blue-melange-shorts-820312-4943a8"
                  ]
                }
              }
            },
            "discount": "",
            "bulk_offer": {},
            "key": "820312_L",
            "price": {
              "base": {
                "add_on": 1999,
                "marked": 1999,
                "effective": 1999,
                "selling": 1999,
                "currency_code": "INR"
              },
              "converted": {
                "add_on": 1999,
                "marked": 1999,
                "effective": 1999,
                "selling": 1999,
                "currency_code": "INR"
              }
            },
            "article": {
              "type": "article",
              "uid": "381_610_IGPL01_SPIRAL19ANAVY_L",
              "size": "L",
              "seller": {
                "uid": 381,
                "name": "INTERSOURCE GARMENTS PVT LTD"
              },
              "store": {
                "uid": 3009,
                "name": "Kormangala"
              },
              "quantity": 2,
              "price": {
                "base": {
                  "marked": 1999,
                  "effective": 1999,
                  "currency_code": "INR"
                },
                "converted": {
                  "marked": 1999,
                  "effective": 1999,
                  "currency_code": "INR"
                }
              }
            },
            "availability": {
              "sizes": [
                "L",
                "XL",
                "XXL"
              ],
              "other_store_quantity": 1,
              "out_of_stock": false,
              "deliverable": true,
              "is_valid": true
            },
            "coupon_message": "",
            "message": ""
          }
        ]
      }
    ],
    "delivery_charge_info": "",
    "coupon_text": "View all offers",
    "gstin": null,
    "checkout_mode": "self",
    "last_modified": "Thu, 22 Aug 2019 20:21:48 GMT",
    "restrict_checkout": false,
    "is_valid": true
  }
}
```

Shipment Generation Failed
```json
{
  "value": {
    "items": [],
    "cart_id": 7501,
    "uid": "7501",
    "success": true,
    "error_message": "Note: Your order delivery will be delayed by 7-10 Days",
    "payment_options": {
      "payment_option": [
        {
          "name": "COD",
          "display_name": "Cash on Delivery",
          "display_priority": 1,
          "payment_mode_id": 11,
          "logo": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png",
          "logo_url": {
            "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png",
            "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png"
          },
          "list": []
        },
        {
          "name": "CARD",
          "display_priority": 2,
          "payment_mode_id": 2,
          "display_name": "Card",
          "list": []
        },
        {
          "name": "NB",
          "display_priority": 3,
          "payment_mode_id": 3,
          "display_name": "Net Banking",
          "list": [
            {
              "aggregator_name": "Razorpay",
              "bank_name": "ICICI Bank",
              "bank_code": "ICIC",
              "url": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png"
              },
              "merchant_code": "NB_ICICI",
              "display_priority": 1
            }
          ]
        },
        {
          "name": "WL",
          "display_priority": 4,
          "payment_mode_id": 4,
          "display_name": "Wallet",
          "list": [
            {
              "wallet_name": "Paytm",
              "wallet_code": "paytm",
              "wallet_id": 4,
              "merchant_code": "PAYTM",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/paytm_logo_small.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/paytm_logo_large.png"
              },
              "aggregator_name": "Juspay",
              "display_priority": 1
            }
          ]
        },
        {
          "name": "UPI",
          "display_priority": 9,
          "payment_mode_id": 6,
          "display_name": "UPI",
          "list": [
            {
              "aggregator_name": "UPI_Razorpay",
              "name": "UPI",
              "display_name": "BHIM UPI",
              "code": "UPI",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/upi_100x78.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/upi_150x100.png"
              },
              "merchant_code": "UPI",
              "timeout": 240,
              "retry_count": 0,
              "fynd_vpa": "shopsense.rzp@hdfcbank",
              "intent_flow": true,
              "intent_app_error_list": [
                "com.csam.icici.bank.imobile",
                "in.org.npci.upiapp",
                "com.whatsapp"
              ]
            }
          ]
        },
        {
          "name": "PL",
          "display_priority": 11,
          "payment_mode_id": 1,
          "display_name": "Pay Later",
          "list": [
            {
              "aggregator_name": "Simpl",
              "name": "Simpl",
              "code": "simpl",
              "merchant_code": "SIMPL",
              "logo": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png",
              "logo_url": {
                "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png",
                "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png"
              }
            }
          ]
        }
      ],
      "payment_flows": {
        "Simpl": {
          "data": {
            "gateway": {
              "route": "simpl",
              "entity": "sdk",
              "is_customer_validation_required": true,
              "cust_validation_url": "https://api.addsale.com/gringotts/api/v1/validate-customer/",
              "sdk": {
                "config": {
                  "redirect": false,
                  "callback_url": null,
                  "action_url": "https://api.addsale.com/avis/api/v1/payments/charge-gringotts-transaction/"
                },
                "data": {
                  "user_phone": "8452996729",
                  "user_email": "paymentsdummy@gofynd.com"
                }
              },
              "return_url": null
            }
          },
          "api_link": "",
          "payment_flow": "sdk"
        },
        "Juspay": {
          "data": {},
          "api_link": "https://sandbox.juspay.in/txns",
          "payment_flow": "api"
        },
        "Razorpay": {
          "data": {},
          "api_link": "",
          "payment_flow": "sdk"
        },
        "UPI_Razorpay": {
          "data": {},
          "api_link": "https://api.addsale.com/gringotts/api/v1/external/payment-initialisation/",
          "payment_flow": "api"
        },
        "Fynd": {
          "data": {},
          "api_link": "",
          "payment_flow": "api"
        }
      },
      "default": {}
    },
    "user_type": "Store User",
    "cod_charges": 0,
    "order_id": null,
    "cod_available": true,
    "cod_message": "No additional COD charges applicable",
    "delivery_charges": 0,
    "delivery_charge_order_value": 0,
    "delivery_slots": [
      {
        "date": "Sat, 24 Aug",
        "delivery_slot": [
          {
            "delivery_slot_timing": "By 9:00 PM",
            "default": true,
            "delivery_slot_id": 1
          }
        ]
      }
    ],
    "store_code": "",
    "store_emps": [],
    "breakup_values": {
      "loyalty_points": {
        "total": 0,
        "applicable": 0,
        "is_applied": false,
        "description": "Your cashback, referrals, and refund amount get credited to Fynd Cash which can be redeemed while placing an order."
      },
      "coupon": {
        "type": "cash",
        "code": "",
        "uid": null,
        "value": 0,
        "is_applied": false,
        "message": "Sorry! Invalid Coupon"
      },
      "raw": {
        "cod_charge": 0,
        "convenience_fee": 0,
        "coupon": 0,
        "delivery_charge": 0,
        "discount": 0,
        "fynd_cash": 0,
        "gst_charges": 214.18,
        "mrp_total": 1999,
        "subtotal": 1999,
        "total": 1999,
        "vog": 1784.82,
        "you_saved": 0
      },
      "display": [
        {
          "display": "MRP Total",
          "key": "mrp_total",
          "value": 1999,
          "currency_code": "INR"
        },
        {
          "display": "Subtotal",
          "key": "subtotal",
          "value": 1999,
          "currency_code": "INR"
        },
        {
          "display": "Total",
          "key": "total",
          "value": 1999,
          "currency_code": "INR"
        }
      ]
    },
    "shipments": [],
    "message": "Shipments could not be generated. Please Try again after some time.",
    "delivery_charge_info": "",
    "coupon_text": "View all offers",
    "gstin": null,
    "checkout_mode": "self",
    "last_modified": "Thu, 22 Aug 2019 20:21:48 GMT",
    "restrict_checkout": false,
    "is_valid": false
  }
}
```








Unhandled api error


Schema: `Object`









---


#### checkoutCart
Checkout Cart

```golang

 data, err :=  PosCart.CheckoutCart(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `UID`

| body |  CartPosCheckoutRequest | "Request body" 


Checkout all items in cart to payment and order generation.                        For COD only order will be generated while for other checkout mode                        user will be redirected to payment gateway

*Success Response:*



OK


Schema: `CartCheckoutResponse`


*Examples:*


Address id not found
```json
{
  "value": {
    "success": false,
    "message": "No address found with address id {address_id}"
  }
}
```

Missing address_id
```json
{
  "value": {
    "address_id": [
      "Missing data for required field."
    ]
  }
}
```

Successful checkout cod payment
```json
{
  "value": {
    "success": true,
    "cart": {
      "success": true,
      "error_message": "Note: Your order delivery will be delayed by 7-10 Days",
      "payment_options": {
        "payment_option": [
          {
            "name": "COD",
            "display_name": "Cash on Delivery",
            "display_priority": 1,
            "payment_mode_id": 11,
            "logo": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png",
            "logo_url": {
              "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png",
              "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/cod.png"
            },
            "list": []
          },
          {
            "name": "CARD",
            "display_priority": 2,
            "payment_mode_id": 2,
            "display_name": "Card",
            "list": []
          },
          {
            "name": "NB",
            "display_priority": 3,
            "payment_mode_id": 3,
            "display_name": "Net Banking",
            "list": [
              {
                "aggregator_name": "Razorpay",
                "bank_name": "ICICI Bank",
                "bank_code": "ICIC",
                "url": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png",
                "logo_url": {
                  "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png",
                  "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/NB_ICICI.png"
                },
                "merchant_code": "NB_ICICI",
                "display_priority": 1
              }
            ]
          },
          {
            "name": "WL",
            "display_priority": 4,
            "payment_mode_id": 4,
            "display_name": "Wallet",
            "list": [
              {
                "wallet_name": "Paytm",
                "wallet_code": "paytm",
                "wallet_id": 4,
                "merchant_code": "PAYTM",
                "logo_url": {
                  "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/paytm_logo_small.png",
                  "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/paytm_logo_large.png"
                },
                "aggregator_name": "Juspay",
                "display_priority": 1
              }
            ]
          },
          {
            "name": "UPI",
            "display_priority": 9,
            "payment_mode_id": 6,
            "display_name": "UPI",
            "list": [
              {
                "aggregator_name": "UPI_Razorpay",
                "name": "UPI",
                "display_name": "BHIM UPI",
                "code": "UPI",
                "logo_url": {
                  "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/upi_100x78.png",
                  "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/upi_150x100.png"
                },
                "merchant_code": "UPI",
                "timeout": 240,
                "retry_count": 0,
                "fynd_vpa": "shopsense.rzp@hdfcbank",
                "intent_flow": true,
                "intent_app_error_list": [
                  "com.csam.icici.bank.imobile",
                  "in.org.npci.upiapp",
                  "com.whatsapp"
                ]
              }
            ]
          },
          {
            "name": "PL",
            "display_priority": 11,
            "payment_mode_id": 1,
            "display_name": "Pay Later",
            "list": [
              {
                "aggregator_name": "Simpl",
                "name": "Simpl",
                "code": "simpl",
                "merchant_code": "SIMPL",
                "logo": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png",
                "logo_url": {
                  "small": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png",
                  "large": "https://d2co8r51m5ca2d.cloudfront.net/payments_assets/simpl_logo.png"
                }
              }
            ]
          }
        ],
        "payment_flows": {
          "Simpl": {
            "data": {
              "gateway": {
                "route": "simpl",
                "entity": "sdk",
                "is_customer_validation_required": true,
                "cust_validation_url": "https://api.addsale.com/gringotts/api/v1/validate-customer/",
                "sdk": {
                  "config": {
                    "redirect": false,
                    "callback_url": null,
                    "action_url": "https://api.addsale.com/avis/api/v1/payments/charge-gringotts-transaction/"
                  },
                  "data": {
                    "user_phone": "8452996729",
                    "user_email": "paymentsdummy@gofynd.com"
                  }
                },
                "return_url": null
              }
            },
            "api_link": "",
            "payment_flow": "sdk"
          },
          "Juspay": {
            "data": {},
            "api_link": "https://sandbox.juspay.in/txns",
            "payment_flow": "api"
          },
          "Razorpay": {
            "data": {},
            "api_link": "",
            "payment_flow": "sdk"
          },
          "UPI_Razorpay": {
            "data": {},
            "api_link": "https://api.addsale.com/gringotts/api/v1/external/payment-initialisation/",
            "payment_flow": "api"
          },
          "Fynd": {
            "data": {},
            "api_link": "",
            "payment_flow": "api"
          }
        },
        "default": {}
      },
      "user_type": "Store User",
      "cod_charges": 0,
      "order_id": "FY5D5E215CF287584CE6",
      "cod_available": true,
      "cod_message": "No additional COD charges applicable",
      "delivery_charges": 0,
      "delivery_charge_order_value": 0,
      "delivery_slots": [
        {
          "date": "Sat, 24 Aug",
          "delivery_slot": [
            {
              "delivery_slot_timing": "By 9:00 PM",
              "default": true,
              "delivery_slot_id": 1
            }
          ]
        }
      ],
      "store_code": "",
      "store_emps": [],
      "breakup_values": {
        "coupon": {
          "type": "cash",
          "code": "",
          "uid": null,
          "value": 0,
          "is_applied": false,
          "message": "Sorry! Invalid Coupon"
        },
        "loyalty_points": {
          "total": 0,
          "applicable": 0,
          "is_applied": false,
          "description": "Your cashback, referrals, and refund amount get credited to Fynd Cash which can be redeemed while placing an order."
        },
        "raw": {
          "cod_charge": 0,
          "convenience_fee": 0,
          "coupon": 0,
          "delivery_charge": 0,
          "discount": 0,
          "fynd_cash": 0,
          "gst_charges": 214.18,
          "mrp_total": 1999,
          "subtotal": 1999,
          "total": 1999,
          "vog": 1784.82,
          "you_saved": 0
        },
        "display": [
          {
            "display": "MRP Total",
            "key": "mrp_total",
            "value": 1999,
            "currency_code": "INR"
          },
          {
            "display": "Subtotal",
            "key": "subtotal",
            "value": 1999,
            "currency_code": "INR"
          },
          {
            "display": "Total",
            "key": "total",
            "value": 1999,
            "currency_code": "INR"
          }
        ]
      },
      "items": [
        {
          "key": "820312_L",
          "message": "",
          "bulk_offer": {},
          "price": {
            "base": {
              "add_on": 1999,
              "marked": 1999,
              "effective": 1999,
              "selling": 1999,
              "currency_code": "INR"
            },
            "converted": {
              "add_on": 1999,
              "marked": 1999,
              "effective": 1999,
              "selling": 1999,
              "currency_code": "INR"
            }
          },
          "quantity": 1,
          "discount": "",
          "product": {
            "type": "product",
            "uid": 820312,
            "name": "Navy Blue Melange Shorts",
            "slug": "883-police-navy-blue-melange-shorts-820312-4943a8",
            "brand": {
              "uid": 610,
              "name": "883 Police"
            },
            "categories": [
              {
                "uid": 193,
                "name": "Shorts"
              }
            ],
            "images": [
              {
                "aspect_ratio": "16:25",
                "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/610_SPIRAL19ANAVY/1_1549105947281.jpg",
                "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/610_SPIRAL19ANAVY/1_1549105947281.jpg"
              }
            ],
            "action": {
              "type": "product",
              "url": "https://api.addsale.com/platform/content/v1/products/883-police-navy-blue-melange-shorts-820312-4943a8/",
              "query": {
                "product_slug": [
                  "883-police-navy-blue-melange-shorts-820312-4943a8"
                ]
              }
            }
          },
          "article": {
            "type": "article",
            "uid": "381_610_IGPL01_SPIRAL19ANAVY_L",
            "size": "L",
            "seller": {
              "uid": 381,
              "name": "INTERSOURCE GARMENTS PVT LTD"
            },
            "store": {
              "uid": 3009,
              "name": "Kormangala"
            },
            "quantity": 2,
            "price": {
              "base": {
                "marked": 1999,
                "effective": 1999,
                "currency_code": "INR"
              },
              "converted": {
                "marked": 1999,
                "effective": 1999,
                "currency_code": "INR"
              }
            }
          },
          "coupon_message": "",
          "availability": {
            "sizes": [
              "L",
              "XL",
              "XXL"
            ],
            "other_store_quantity": 1,
            "out_of_stock": false,
            "deliverable": true,
            "is_valid": true
          }
        }
      ],
      "delivery_charge_info": "",
      "coupon_text": "View all offers",
      "cart_id": 7483,
      "uid": "7483",
      "gstin": null,
      "checkout_mode": "self",
      "last_modified": "Thu, 22 Aug 2019 04:58:44 GMT",
      "restrict_checkout": false,
      "is_valid": true
    },
    "callback_url": "https://api.addsale.com/gringotts/api/v1/external/payment-callback/",
    "app_intercept_url": "http://uniket-testing.addsale.link/cart/order-status",
    "message": "",
    "data": {
      "order_id": "FY5D5E215CF287584CE6"
    },
    "order_id": "FY5D5E215CF287584CE6"
  }
}
```









---


#### updateCartMeta
Update Cart Meta

```golang

 data, err :=  PosCart.UpdateCartMeta(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `UID`

| body |  CartMetaRequest | "Request body" 


Update cart meta like checkout_mode, gstin.

*Success Response:*



Cart meta updated successfully


Schema: `CartMetaResponse`








Missing required Field


Schema: `CartMetaMissingResponse`









---


#### getAvailableDeliveryModes
Get available delivery modes for cart

```golang

 data, err :=  PosCart.GetAvailableDeliveryModes(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `AreaCode`, `UID`



Get available delivery modes for cart and pick up store uid list. From given pick stores list user can pick up delivery. Use this uid to show store address

*Success Response:*



Returns Available delivery modes for cart and pick up available store uid for current cart items


Schema: `CartDeliveryModesResponse`









---


#### getStoreAddressByUid
Get list of stores for give uids

```golang

 data, err :=  PosCart.GetStoreAddressByUid(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `StoreUID`



Get list of stores by providing pick up available store uids.

*Success Response:*



Returns available store information with its address


Schema: `StoreDetailsResponse`









---


#### getCartShareLink
Generate Cart sharing link token

```golang

 data, err :=  PosCart.GetCartShareLink(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  GetShareCartLinkRequest | "Request body" 


Generates shared cart snapshot and returns shortlink token

*Success Response:*



Token Generated successfully


Schema: `GetShareCartLinkResponse`


*Examples:*


Token Generated
```json
{
  "value": {
    "token": "ZweG1XyX",
    "share_url": "https://uniket-testing.addsale.link/shared-cart/ZweG1XyX"
  }
}
```









---


#### getCartSharedItems
Get shared cart snapshot and cart response

```golang

 data, err :=  PosCart.GetCartSharedItems(Token);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Token | string | Shared short link token. | 




Returns shared cart response for sent token with `shared_cart_details`                    containing shared cart details in response.

*Success Response:*



Cart for valid token


Schema: `SharedCartResponse`








No cart found for sent token


Schema: `SharedCartResponse`









---


#### updateCartWithSharedItems
Merge or Replace existing cart

```golang

 data, err :=  PosCart.UpdateCartWithSharedItems(Token, Action);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Token | string | Shared short link token. | 


| Action | string | Operation to perform on existing cart, whether to merge or replace. | 




Merge or Replace cart based on `action` parameter with shared cart of `token`

*Success Response:*



Success of Merge or Replace of cart with `shared_cart_details`                    containing shared cart details in response


Schema: `SharedCartResponse`


*Examples:*


Cart Merged/Replaced
```json
{
  "value": {
    "cart": {
      "shared_cart_details": {
        "token": "BQ9jySQ9",
        "user": {
          "user_id": "23109086",
          "is_anonymous": false
        },
        "meta": {
          "selected_staff": "",
          "ordering_store": null
        },
        "selected_staff": "",
        "ordering_store": null,
        "source": {},
        "created_on": "2019-12-18T14:00:07.165000"
      },
      "items": [
        {
          "key": "791651_6",
          "discount": "",
          "bulk_offer": {},
          "coupon_message": "",
          "article": {
            "type": "article",
            "uid": "304_1054_9036_R1005753_6",
            "size": "6",
            "seller": {
              "uid": 304,
              "name": "LEAYAN GLOBAL PVT. LTD."
            },
            "store": {
              "uid": 5322,
              "name": "Vaisali Nagar"
            },
            "quantity": 1,
            "price": {
              "base": {
                "marked": 2095,
                "effective": 2095,
                "currency_code": "INR"
              },
              "converted": {
                "marked": 2095,
                "effective": 2095,
                "currency_code": "INR"
              }
            }
          },
          "product": {
            "type": "product",
            "uid": 791651,
            "name": "Black Running Shoes",
            "slug": "furo-black-running-shoes-791651-f8bcc3",
            "brand": {
              "uid": 1054,
              "name": "Furo"
            },
            "categories": [
              {
                "uid": 160,
                "name": "Running Shoes"
              }
            ],
            "images": [
              {
                "aspect_ratio": "16:25",
                "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/1054_R1005753/1_1546490507364.jpg",
                "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/1054_R1005753/1_1546490507364.jpg"
              }
            ],
            "action": {
              "type": "product",
              "url": "https://api.addsale.com/platform/content/v1/products/furo-black-running-shoes-791651-f8bcc3/",
              "query": {
                "product_slug": [
                  "furo-black-running-shoes-791651-f8bcc3"
                ]
              }
            }
          },
          "message": "",
          "quantity": 1,
          "availability": {
            "sizes": [
              "7",
              "8",
              "9",
              "10",
              "6"
            ],
            "other_store_quantity": 12,
            "out_of_stock": false,
            "deliverable": true,
            "is_valid": true
          },
          "price": {
            "base": {
              "add_on": 2095,
              "marked": 2095,
              "effective": 2095,
              "selling": 2095,
              "currency_code": "INR"
            },
            "converted": {
              "add_on": 2095,
              "marked": 2095,
              "effective": 2095,
              "selling": 2095,
              "currency_code": "INR"
            }
          }
        },
        {
          "key": "791651_7",
          "discount": "",
          "bulk_offer": {},
          "coupon_message": "",
          "article": {
            "type": "article",
            "uid": "304_1054_9036_R1005753_7",
            "size": "7",
            "seller": {
              "uid": 304,
              "name": "LEAYAN GLOBAL PVT. LTD."
            },
            "store": {
              "uid": 5322,
              "name": "Vaisali Nagar"
            },
            "quantity": 2,
            "price": {
              "base": {
                "marked": 2095,
                "effective": 2095,
                "currency_code": "INR"
              },
              "converted": {
                "marked": 2095,
                "effective": 2095,
                "currency_code": "INR"
              }
            }
          },
          "product": {
            "type": "product",
            "uid": 791651,
            "name": "Black Running Shoes",
            "slug": "furo-black-running-shoes-791651-f8bcc3",
            "brand": {
              "uid": 1054,
              "name": "Furo"
            },
            "categories": [
              {
                "uid": 160,
                "name": "Running Shoes"
              }
            ],
            "images": [
              {
                "aspect_ratio": "16:25",
                "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/1054_R1005753/1_1546490507364.jpg",
                "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/1054_R1005753/1_1546490507364.jpg"
              }
            ],
            "action": {
              "type": "product",
              "url": "https://api.addsale.com/platform/content/v1/products/furo-black-running-shoes-791651-f8bcc3/",
              "query": {
                "product_slug": [
                  "furo-black-running-shoes-791651-f8bcc3"
                ]
              }
            }
          },
          "message": "",
          "quantity": 2,
          "availability": {
            "sizes": [
              "7",
              "8",
              "9",
              "10",
              "6"
            ],
            "other_store_quantity": 7,
            "out_of_stock": false,
            "deliverable": true,
            "is_valid": true
          },
          "price": {
            "base": {
              "add_on": 4190,
              "marked": 4190,
              "effective": 4190,
              "selling": 4190,
              "currency_code": "INR"
            },
            "converted": {
              "add_on": 4190,
              "marked": 4190,
              "effective": 4190,
              "selling": 4190,
              "currency_code": "INR"
            }
          }
        }
      ],
      "cart_id": 13055,
      "uid": "13055",
      "breakup_values": {
        "raw": {
          "cod_charge": 0,
          "convenience_fee": 0,
          "coupon": 0,
          "delivery_charge": 0,
          "discount": 0,
          "fynd_cash": 0,
          "gst_charges": 958.73,
          "mrp_total": 6285,
          "subtotal": 6285,
          "total": 6285,
          "vog": 5326.27,
          "you_saved": 0
        },
        "loyalty_points": {
          "total": 0,
          "applicable": 0,
          "is_applied": false,
          "description": "Your cashback, referrals, and refund amount get credited to Fynd Cash which can be redeemed while placing an order."
        },
        "coupon": {
          "type": "cash",
          "code": "",
          "uid": null,
          "value": 0,
          "is_applied": false,
          "message": "Sorry! Invalid coupon"
        },
        "display": [
          {
            "display": "MRP Total",
            "key": "mrp_total",
            "value": 6285,
            "currency_code": "INR"
          },
          {
            "display": "Subtotal",
            "key": "subtotal",
            "value": 6285,
            "currency_code": "INR"
          },
          {
            "display": "Total",
            "key": "total",
            "value": 6285,
            "currency_code": "INR"
          }
        ]
      },
      "delivery_charge_info": "",
      "coupon_text": "View all offers",
      "gstin": null,
      "comment": "",
      "checkout_mode": "self",
      "payment_selection_lock": {
        "enabled": false,
        "default_options": "COD",
        "payment_identifier": null
      },
      "restrict_checkout": false,
      "is_valid": true,
      "last_modified": "Mon, 16 Dec 2019 07:02:18 GMT"
    }
  }
}
```









---



---


## Logistic


#### getTatProduct
Get Tat Product

```golang

 data, err :=  Logistic.GetTatProduct(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  GetTatProductReqBody | "Request body" 


Get Tat Product

*Success Response:*



Success


Schema: `GetTatProductResponse`








API Error


Schema: `ApefaceApiError`








Internal Server Error


Schema: `ApefaceApiError`









---


#### getPincodeCity
Get City from Pincode

```golang

 data, err :=  Logistic.GetPincodeCity(Pincode);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Pincode | string | Pincode | 




Get City from Pincode

*Success Response:*



Success


Schema: `GetPincodeCityResponse`








API Error


Schema: `ApefaceApiError`








Internal Server Error


Schema: `ApefaceApiError`









---



---


---
---
