# FDK Application Front API Documentaion


* [Catalog](#Catalog) - Catalog API's allows you to access list of products, prices, seller details, similar features, variants and many more useful features.  
* [Cart](#Cart) - Cart APIs 
* [Common](#Common) - Application configuration apis 
* [Lead](#Lead) - Handles communication between Staff and Users 
* [Theme](#Theme) - Responsible for themes 
* [User](#User) - Authentication Service 
* [Content](#Content) - Content System 
* [Communication](#Communication) - Manages email, sms, push notifications sent to users 
* [Share](#Share) - Short link and QR Code 
* [FileStorage](#FileStorage) - File Storage 
* [Configuration](#Configuration) - Application configuration apis 
* [Payment](#Payment) - Collect payment through many payment gateway i.e Stripe, Razorpay, Juspay etc.into Fynd or Self account 
* [Order](#Order) - Handles all Application order and shipment api(s) 
* [Rewards](#Rewards) - Earn and redeem reward points 
* [PosCart](#PosCart) - Cart APIs 
* [Logistic](#Logistic) - Logistics Promise Engine APIs allows you to configure zone, pincode, TAT, logistics and many more useful features.  

----
----

### Classes and Methods


* [Catalog](#Catalog)
  * Methods
    * [getProductDetailBySlug](#getproductdetailbyslug)
    * [getProductSizesBySlug](#getproductsizesbyslug)
    * [getProductComparisonBySlugs](#getproductcomparisonbyslugs)
    * [getSimilarComparisonProductBySlug](#getsimilarcomparisonproductbyslug)
    * [getComparedFrequentlyProductBySlug](#getcomparedfrequentlyproductbyslug)
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
    * [unfollowById](#unfollowbyid)
    * [followById](#followbyid)
    * [getFollowerCountById](#getfollowercountbyid)
    * [getFollowIds](#getfollowids)
    * [getStores](#getstores)
    * [getInStockLocations](#getinstocklocations)
    * [getLocationDetailsById](#getlocationdetailsbyid)
    * [getProductBundlesBySlug](#getproductbundlesbyslug)
    * [getProductPriceBySlug](#getproductpricebyslug)
    * [getProductSellersBySlug](#getproductsellersbyslug)
    

* [Cart](#Cart)
  * Methods
    * [getCart](#getcart)
    * [getCartLastModified](#getcartlastmodified)
    * [addItems](#additems)
    * [updateCart](#updatecart)
    * [deleteCart](#deletecart)
    * [getItemCount](#getitemcount)
    * [getCoupons](#getcoupons)
    * [applyCoupon](#applycoupon)
    * [removeCoupon](#removecoupon)
    * [getBulkDiscountOffers](#getbulkdiscountoffers)
    * [applyRewardPoints](#applyrewardpoints)
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
    * [getPromotionOffers](#getpromotionoffers)
    * [getLadderOffers](#getladderoffers)
    

* [Common](#Common)
  * Methods
    * [searchApplication](#searchapplication)
    * [getLocations](#getlocations)
    

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
    * [getAllPages](#getallpages)
    * [getPage](#getpage)
    * [getAppliedTheme](#getappliedtheme)
    * [getThemeForPreview](#getthemeforpreview)
    

* [User](#User)
  * Methods
    * [loginWithFacebook](#loginwithfacebook)
    * [loginWithGoogle](#loginwithgoogle)
    * [loginWithGoogleAndroid](#loginwithgoogleandroid)
    * [loginWithGoogleIOS](#loginwithgoogleios)
    * [loginWithAppleIOS](#loginwithappleios)
    * [loginWithOTP](#loginwithotp)
    * [loginWithEmailAndPassword](#loginwithemailandpassword)
    * [sendResetPasswordEmail](#sendresetpasswordemail)
    * [sendResetPasswordMobile](#sendresetpasswordmobile)
    * [forgotPassword](#forgotpassword)
    * [sendResetToken](#sendresettoken)
    * [loginWithToken](#loginwithtoken)
    * [registerWithForm](#registerwithform)
    * [verifyEmail](#verifyemail)
    * [verifyMobile](#verifymobile)
    * [hasPassword](#haspassword)
    * [updatePassword](#updatepassword)
    * [deleteUser](#deleteuser)
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
    * [getDataLoaders](#getdataloaders)
    * [getFaqs](#getfaqs)
    * [getFaqCategories](#getfaqcategories)
    * [getFaqBySlug](#getfaqbyslug)
    * [getFaqCategoryBySlug](#getfaqcategorybyslug)
    * [getFaqsByCategorySlug](#getfaqsbycategoryslug)
    * [getLandingPage](#getlandingpage)
    * [getLegalInformation](#getlegalinformation)
    * [getNavigations](#getnavigations)
    * [getSEOConfiguration](#getseoconfiguration)
    * [getSlideshows](#getslideshows)
    * [getSlideshow](#getslideshow)
    * [getSupportInformation](#getsupportinformation)
    * [getTags](#gettags)
    * [getPage](#getpage)
    * [getPages](#getpages)
    

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
    * [signUrls](#signurls)
    

* [Configuration](#Configuration)
  * Methods
    * [getApplication](#getapplication)
    * [getOwnerInfo](#getownerinfo)
    * [getBasicDetails](#getbasicdetails)
    * [getIntegrationTokens](#getintegrationtokens)
    * [getOrderingStores](#getorderingstores)
    * [getStoreDetailById](#getstoredetailbyid)
    * [getFeatures](#getfeatures)
    * [getContactInfo](#getcontactinfo)
    * [getCurrencies](#getcurrencies)
    * [getCurrencyById](#getcurrencybyid)
    * [getAppCurrencies](#getappcurrencies)
    * [getLanguages](#getlanguages)
    * [getOrderingStoreCookie](#getorderingstorecookie)
    * [removeOrderingStoreCookie](#removeorderingstorecookie)
    * [getAppStaffList](#getappstafflist)
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
    * [getRupifiBannerDetails](#getrupifibannerdetails)
    * [getEpaylaterBannerDetails](#getepaylaterbannerdetails)
    * [resendOrCancelPayment](#resendorcancelpayment)
    * [renderHTML](#renderhtml)
    * [validateVPA](#validatevpa)
    * [getActiveRefundTransferModes](#getactiverefundtransfermodes)
    * [enableOrDisableRefundTransferMode](#enableordisablerefundtransfermode)
    * [getUserBeneficiariesDetail](#getuserbeneficiariesdetail)
    * [verifyIfscCode](#verifyifsccode)
    * [getOrderBeneficiariesDetail](#getorderbeneficiariesdetail)
    * [verifyOtpAndAddBeneficiaryForBank](#verifyotpandaddbeneficiaryforbank)
    * [addBeneficiaryDetails](#addbeneficiarydetails)
    * [addRefundBankAccountUsingOTP](#addrefundbankaccountusingotp)
    * [verifyOtpAndAddBeneficiaryForWallet](#verifyotpandaddbeneficiaryforwallet)
    * [updateDefaultBeneficiary](#updatedefaultbeneficiary)
    * [getPaymentLink](#getpaymentlink)
    * [createPaymentLink](#createpaymentlink)
    * [resendPaymentLink](#resendpaymentlink)
    * [cancelPaymentLink](#cancelpaymentlink)
    * [getPaymentModeRoutesPaymentLink](#getpaymentmoderoutespaymentlink)
    * [pollingPaymentLink](#pollingpaymentlink)
    * [createOrderHandlerPaymentLink](#createorderhandlerpaymentlink)
    * [initialisePaymentPaymentLink](#initialisepaymentpaymentlink)
    * [checkAndUpdatePaymentStatusPaymentLink](#checkandupdatepaymentstatuspaymentlink)
    * [customerCreditSummary](#customercreditsummary)
    * [redirectToAggregator](#redirecttoaggregator)
    * [checkCredit](#checkcredit)
    * [customerOnboard](#customeronboard)
    

* [Order](#Order)
  * Methods
    * [getOrders](#getorders)
    * [getOrderById](#getorderbyid)
    * [getPosOrderById](#getposorderbyid)
    * [getShipmentById](#getshipmentbyid)
    * [getInvoiceByShipmentId](#getinvoicebyshipmentid)
    * [trackShipment](#trackshipment)
    * [getCustomerDetailsByShipmentId](#getcustomerdetailsbyshipmentid)
    * [sendOtpToShipmentCustomer](#sendotptoshipmentcustomer)
    * [verifyOtpShipmentCustomer](#verifyotpshipmentcustomer)
    * [getShipmentBagReasons](#getshipmentbagreasons)
    * [getShipmentReasons](#getshipmentreasons)
    * [updateShipmentStatus](#updateshipmentstatus)
    

* [Rewards](#Rewards)
  * Methods
    * [getOfferByName](#getofferbyname)
    * [catalogueOrder](#catalogueorder)
    * [getUserPointsHistory](#getuserpointshistory)
    * [getUserPoints](#getuserpoints)
    * [getUserReferralDetails](#getuserreferraldetails)
    * [getOrderDiscount](#getorderdiscount)
    * [redeemReferralCode](#redeemreferralcode)
    

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
    * [applyRewardPoints](#applyrewardpoints)
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
    * [getPincodeCity](#getpincodecity)
    * [getTatProduct](#gettatproduct)
    * [getAllCountries](#getallcountries)
    * [getPincodeZones](#getpincodezones)
    * [getOptimalLocations](#getoptimallocations)
    


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









---


#### getCollections
List all the collections

```golang

 data, err :=  Catalog.GetCollections(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |








| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Tag`, `Q`



Collections are a great way to organize your products and can improve the ability for customers to find items quickly and efficiently.

*Success Response:*



Success. Returns a list of collections. Check the example shown below or refer `GetCollectionListingResponse` for more details.


Schema: `GetCollectionListingResponse`









---


#### getCollectionItemsBySlug
Get the items in a collection

```golang

 data, err :=  Catalog.GetCollectionItemsBySlug(Slug, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of a collection. You can get slug value from the endpoint /service/application/catalog/v1.0/collections/. | 

















| xQuery | struct | Includes properties such as `F`, `Q`, `Filters`, `SortOn`, `PageID`, `PageSize`, `PageNo`, `PageType`



Get items in a collection specified by its `slug`.

*Success Response:*



Success. Returns a list items in a given collection. Check the example shown below or refer `ProductListingResponse` for more details.


Schema: `ProductListingResponse`









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









---


#### getStores
Get store meta information.

```golang

 data, err :=  Catalog.GetStores(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |














| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Q`, `City`, `Range`, `Latitude`, `Longitude`



Use this API to get a list of stores in a specific application.

*Success Response:*



Success. Returns a list of selling locations. Check the example shown below or refer `StoreListingResponse` for more details.


Schema: `StoreListingResponse`









---


#### getInStockLocations
Get store meta information.

```golang

 data, err :=  Catalog.GetInStockLocations(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |














| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Q`, `City`, `Range`, `Latitude`, `Longitude`



Use this API to get a list of stores in a specific application.

*Success Response:*



Success. Returns a list of selling locations. Check the example shown below or refer `StoreListingResponse` for more details.


Schema: `ApplicationStoreListing`









---


#### getLocationDetailsById
Get store meta information.

```golang

 data, err :=  Catalog.GetLocationDetailsById(LocationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| LocationID | float64 | Unique Location ID. | 




Use this API to get meta details for a store.

*Success Response:*



Success. Returns a metadata object. Check the example shown below or refer `StoreDetails` for more details.


Schema: `StoreDetails`









---


#### getProductBundlesBySlug
Get product bundles

```golang

 data, err :=  Catalog.GetProductBundlesBySlug(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `Slug`, `ID`



Use this API to retrieve products bundles to the one specified by its slug.

*Success Response:*



Success. Returns a group of products bundle.


Schema: `ProductBundle`









---


#### getProductPriceBySlug
Get the price of a product size at a PIN Code

```golang

 data, err :=  Catalog.GetProductPriceBySlug(Slug, Size, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of a product. You can get slug value from the endpoint /service/application/catalog/v1.0/products/ | 


| Size | string | A string indicating the size of the product, e.g. S, M, XL. You can get slug value from the endpoint /service/application/catalog/v1.0/products/sizes | 







| xQuery | struct | Includes properties such as `StoreID`, `Pincode`, `Moq`



Prices may vary for different sizes of a product. Use this API to retrieve the price of a product size at all the selling locations near to a PIN Code.

*Success Response:*



Success. Returns a ProductSizePriceV3 object. Check the example shown below or refer `ProductSizePriceResponseV3` for more details.


Schema: `ProductSizePriceResponseV3`









---


#### getProductSellersBySlug
Get the sellers of a product size at a PIN Code

```golang

 data, err :=  Catalog.GetProductSellersBySlug(Slug, Size, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of a product. You can get slug value from the endpoint /service/application/catalog/v1.0/products/ | 


| Size | string | A string indicating the size of the product, e.g. S, M, XL. You can get slug value from the endpoint /service/application/catalog/v1.0/products/sizes | 









| xQuery | struct | Includes properties such as `Pincode`, `Strategy`, `PageNo`, `PageSize`



A product of a particular size may be sold by multiple sellers. Use this API to fetch the sellers having the stock of a particular size at a given PIN Code.

*Success Response:*



Success. Returns a ProductSizeSellerV3 object. Check the example shown below or refer `ProductSizeSellersResponseV3` for more details.


Schema: `ProductSizeSellersResponseV3`









---



---


## Cart


#### getCart
Fetch all items added to the cart

```golang

 data, err :=  Cart.GetCart(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |












| xQuery | struct | Includes properties such as `ID`, `I`, `B`, `AssignCardID`, `AreaCode`, `BuyNow`



Use this API to get details of all the items added to a cart.

*Success Response:*



Success. Returns a Cart object. Check the example shown below or refer `CartDetailResponse` for more details.


Schema: `CartDetailResponse`









---


#### getCartLastModified
Fetch last-modified timestamp

```golang

 data, err :=  Cart.GetCartLastModified(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `ID`



Use this API to fetch Last-Modified timestamp in header metadata.

*Success Response:*



Success. Receives last modifed timestamp in the header.






---


#### addItems
Add items to cart

```golang

 data, err :=  Cart.AddItems(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |








| xQuery | struct | Includes properties such as `I`, `B`, `AreaCode`, `BuyNow`

| body |  AddCartRequest | "Request body" 


Use this API to add items to the cart.

*Success Response:*



Success. Returns a cart object as shown below. Refer `AddCartDetailResponse` for more details.


Schema: `AddCartDetailResponse`


*Examples:*


Product has been added to your cart
```json
{
  "value": {
    "message": "Product has been added to your cart",
    "success": true,
    "cart": {
      "applied_promo_details": [
        {
          "application_id": "000000000000000000000001",
          "apply_all_offers": false,
          "apply_exclusive": null,
          "buy_article_dict": {
            "rule#1": {
              "633527e850dfb8e73e6de996_0": {
                "added_in_cart": false,
                "added_quantity": 0,
                "amount_paid": 0,
                "applicable_credits": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 0,
                  "floor_value": 0,
                  "source_currency_code": "FC",
                  "source_currency_symbol": "",
                  "source_value": 0,
                  "value": 0
                },
                "article_assign_status": true,
                "article_assignment": {
                  "level": "multi-companies",
                  "strategy": "optimal"
                },
                "article_error": {
                  "type": null,
                  "value": null,
                  "message": null
                },
                "article_id": "633527e850dfb8e73e6de996",
                "article_index": "0",
                "article_meta": {
                  "article_id": "633527e850dfb8e73e6de996",
                  "bulk_margin": 0,
                  "bulk_discount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "INR",
                    "source_currency_symbol": "₹",
                    "source_value": 0,
                    "value": 0
                  },
                  "bulk_coupon_code": null,
                  "country_of_origin": "India",
                  "dimension": {
                    "width": 70,
                    "unit": "cm",
                    "height": 70,
                    "is_default": true,
                    "length": 70
                  },
                  "fragile": false,
                  "manufacturer": {
                    "address": "10, PALI MALA RD, ADARSH NAGAR, BANDRA WEST,, MAHARASHTRA, MUMBAI",
                    "name": "Sabki Shop",
                    "is_default": true
                  },
                  "weight": {
                    "shipping": 250,
                    "unit": "gram",
                    "is_default": true
                  },
                  "raw_meta": {
                    "fynd_identifier": "4187_3412343256098",
                    "is_set_article": false,
                    "set_quantity": 1
                  },
                  "is_set": false,
                  "set": {},
                  "identifier": {
                    "ean": "3412343256098"
                  },
                  "hsn_code": null,
                  "hsn_code_id": "625fbd96faeed8b3df6ec3ce",
                  "raw_price": {
                    "effective": {
                      "conversion_rate": 1,
                      "currency_code": "INR",
                      "currency_symbol": "₹",
                      "floor_source_value": 3899,
                      "floor_value": 3899,
                      "source_currency_code": "INR",
                      "source_currency_symbol": "₹",
                      "source_value": 3899,
                      "value": 3899
                    },
                    "marked": {
                      "conversion_rate": 1,
                      "currency_code": "INR",
                      "currency_symbol": "₹",
                      "floor_source_value": 3899,
                      "floor_value": 3899,
                      "source_currency_code": "INR",
                      "source_currency_symbol": "₹",
                      "source_value": 3899,
                      "value": 3899
                    },
                    "transfer": {
                      "conversion_rate": 1,
                      "currency_code": "INR",
                      "currency_symbol": "₹",
                      "floor_source_value": 0,
                      "floor_value": 0,
                      "source_currency_code": "INR",
                      "source_currency_symbol": "₹",
                      "source_value": 0,
                      "value": 0
                    }
                  },
                  "seller_identifier": "3412343256098",
                  "_custom_json": {},
                  "tags": [],
                  "return_config": {
                    "time": 0,
                    "unit": "days",
                    "returnable": false
                  }
                },
                "attributes": {
                  "essential": "No",
                  "item_code": "BSFJ2-12608",
                  "gender": [
                    "Men"
                  ],
                  "denim_type": "Solid",
                  "pattern": "Solid",
                  "product_fit": "Slim",
                  "currency": "INR",
                  "media": [
                    {
                      "type": "image",
                      "meta": {
                        "brand": "nike",
                        "item_code": "BSFJ2-12608",
                        "sequence": 0
                      },
                      "url": "https://hdn-1.addsale.com/addsale/products/pictures/item/free/original/nike/BSFJ2-12608/0/jMpkpyx17w-GWQjymow3-Black-Slim-Fit-Jeans.jpeg"
                    }
                  ],
                  "name": "Nike Jeans 12608",
                  "brand_name": "Nike"
                },
                "avl_qty": 97,
                "brand_id": 18,
                "bulk_coupon_applied": false,
                "bulk_coupon_code": null,
                "bulk_discount": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 0,
                  "floor_value": 0,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 0,
                  "value": 0
                },
                "bulk_margin": 0,
                "cancellation_allowed": true,
                "cashback": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "FC",
                    "source_currency_symbol": "",
                    "source_value": 0,
                    "value": 0
                  }
                },
                "category_id": [
                  150
                ],
                "cod_charges": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "INR",
                    "source_currency_symbol": "₹",
                    "source_value": 0,
                    "value": 0
                  }
                },
                "company_id": 2,
                "company_info": {
                  "company_id": 2,
                  "c_name": "Sabki Shop",
                  "c_taxes": [
                    {
                      "rate": 15,
                      "enable": true,
                      "effective_date": "2022-06-05T18:29:23.904000"
                    }
                  ],
                  "company_cin": "U45200MH1992PTC066474"
                },
                "company_taxes": [
                  {
                    "rate": 15,
                    "enable": true,
                    "effective_date": "2022-06-05T18:29:23.904000"
                  }
                ],
                "coupon": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "INR",
                    "source_currency_symbol": "₹",
                    "source_value": 0,
                    "value": 0
                  },
                  "article_count": 0,
                  "cancellation_allowed": true,
                  "return_allowed": true
                },
                "currency_dict": {
                  "code": "INR",
                  "rate": 1
                },
                "custom_meta": {},
                "data": {},
                "delivery_charges": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "INR",
                    "source_currency_symbol": "₹",
                    "source_value": 0,
                    "value": 0
                  }
                },
                "delivery_error_msg": "",
                "delivery_promise": {
                  "timestamp": {
                    "min": 1646257339,
                    "max": 1646689339
                  },
                  "formatted": {
                    "min": "03 Mar, Thursday",
                    "max": "08 Mar, Tuesday"
                  }
                },
                "departments": [
                  99
                ],
                "discount": 0,
                "discount_applied": {},
                "dt_currency": "INR",
                "dt_currency_symbol": "₹",
                "esp_modified": false,
                "extra_meta": {},
                "float_amount_paid": 0,
                "free_gift_items_meta": {},
                "group_id": "",
                "gst_amount": 167.1,
                "gst_tax_percentage": 5,
                "hsn_code": null,
                "hsn_code_id": "625fbd96faeed8b3df6ec3ce",
                "identifiers": {
                  "identifier": "BDnwAINORr6SlTKxZN3w"
                },
                "image_nature": "standard",
                "include": true,
                "is_valid": true,
                "item_id": 75660796,
                "item_size": "34",
                "last_update_at": 1680852767,
                "meta": {},
                "net_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3509.1,
                  "floor_value": 3509.1,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3509.1,
                  "value": 3509.1
                },
                "old_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "original_price_effective": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "original_unit_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "parent_item_identifiers": {
                  "identifier": null,
                  "parent_item_size": null,
                  "parent_item_id": null
                },
                "pickup_stores": [],
                "price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "price_effective": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "price_marked": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "product_tags": [],
                "promotions": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "INR",
                    "source_currency_symbol": "₹",
                    "source_value": 0,
                    "value": 0
                  },
                  "applied": [
                    {
                      "amount": {
                        "conversion_rate": 1,
                        "currency_code": "INR",
                        "currency_symbol": "₹",
                        "floor_source_value": 389.9,
                        "floor_value": 389.9,
                        "source_currency_code": "INR",
                        "source_currency_symbol": "₹",
                        "source_value": 389.9,
                        "value": 389.9
                      },
                      "applied_discounts": [
                        {
                          "type": "percentage",
                          "value": 10
                        }
                      ],
                      "applied_free_articles": [],
                      "apply_exclusive": null,
                      "article_quantity": 1,
                      "buy_rules": [
                        {
                          "all_items": null,
                          "cart_conditions": {},
                          "item_criteria": {
                            "item_brand": [
                              18
                            ]
                          },
                          "mrp_promo": false,
                          "slug": "rule#1"
                        }
                      ],
                      "cancellation_allowed": true,
                      "discount_rules": [
                        {
                          "all_items": null,
                          "buy_condition": "( rule#1 )",
                          "item_criteria": {
                            "buy_rules": [
                              "rule#1"
                            ]
                          },
                          "matched_buy_rules": [
                            "rule#1"
                          ],
                          "offer": {
                            "discount_percentage": 10
                          },
                          "raw_offer": {
                            "item_criteria": {
                              "buy_rules": [
                                "rule#1"
                              ]
                            },
                            "buy_condition": "( rule#1 )",
                            "discount_type": "percentage",
                            "offer": {
                              "discount_percentage": 10
                            }
                          },
                          "type": "percentage"
                        }
                      ],
                      "mrp_promo": false,
                      "offer_text": "10% of on NIKE Jeans",
                      "ownership": {
                        "payable_category": "seller",
                        "payable_by": ""
                      },
                      "payable_category": "seller",
                      "promo_code": null,
                      "promo_id": "642fe329ed159b7aa743a293",
                      "promo_type": "percentage",
                      "promotion_group": "product",
                      "return_allowed": true
                    }
                  ],
                  "apply_exclusive": null,
                  "available": [],
                  "cancellation_allowed": true,
                  "exclusive_promo_applied": false,
                  "mrp_promo_applied": false,
                  "return_allowed": true
                },
                "quantity": 1,
                "quantity_assign_status": false,
                "referral_credits": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "FC",
                    "source_currency_symbol": "",
                    "source_value": 0,
                    "value": 0
                  }
                },
                "reload_data_only": true,
                "return_allowed": false,
                "sc_currency": "INR",
                "sc_currency_symbol": "₹",
                "selling_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "service_item_meta": {
                  "product_group_tags": null,
                  "products": null
                },
                "size_level_total_qty": 1,
                "sizes": [
                  "34"
                ],
                "split_article_id": "633527e850dfb8e73e6de996_0",
                "split_index": 0,
                "store_id": 4187,
                "store_info": {
                  "store_id": 4187,
                  "s_city": "HYDERABAD",
                  "store_name": "AND Inorbit Hyderabad",
                  "store_type": "mall",
                  "store_pincode": 500081,
                  "latitude": 17.4343693,
                  "longitude": 78.3866087
                },
                "strategy_validation_data": {},
                "tags": [
                  "rule#1"
                ],
                "total_gst_amount": 167.1,
                "total_value_of_good": 3342,
                "transfer_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 0,
                  "floor_value": 0,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 0,
                  "value": 0
                },
                "unit_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3509.1,
                  "floor_value": 3509.1,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3509.1,
                  "value": 3509.1
                },
                "valid_inventory": true,
                "value_of_good": 3342,
                "verify_article": false
              }
            }
          },
          "buy_rules": [
            {
              "all_items": null,
              "cart_conditions": {},
              "item_criteria": {
                "item_brand": [
                  18
                ]
              },
              "mrp_promo": false,
              "slug": "rule#1"
            }
          ],
          "buy_rules_map": {
            "rule#1": {
              "item_brand": [
                18
              ]
            }
          },
          "calculate_on": "esp",
          "cancellation_allowed": true,
          "code": null,
          "discount": 389.9,
          "discount_rules": [
            {
              "all_items": null,
              "buy_condition": "( rule#1 )",
              "item_criteria": {
                "buy_rules": [
                  "rule#1"
                ]
              },
              "matched_buy_rules": [
                "rule#1"
              ],
              "offer": {
                "discount_percentage": 10
              },
              "raw_offer": {
                "item_criteria": {
                  "buy_rules": [
                    "rule#1"
                  ]
                },
                "buy_condition": "( rule#1 )",
                "discount_type": "percentage",
                "offer": {
                  "discount_percentage": 10
                }
              },
              "type": "percentage"
            }
          ],
          "id": "642fe329ed159b7aa743a293",
          "mrp_promo": false,
          "offer_text": "10% of on NIKE Jeans",
          "ownership": {
            "payable_category": "seller",
            "payable_by": ""
          },
          "payable_category": "seller",
          "priority": 1,
          "promo_group": "product",
          "remaining_allowed_qty": null,
          "restrictions": {
            "uses": {
              "maximum": {
                "user": 0,
                "total": 0
              },
              "remaining": {
                "user": 0,
                "total": 0
              }
            },
            "user_groups": [],
            "post_order": {
              "return_allowed": true,
              "cancellation_allowed": true
            },
            "user_id": [],
            "payments": [],
            "user_registered": {
              "start": null,
              "end": null
            },
            "platforms": [
              "web",
              "android",
              "ios"
            ],
            "anonymous_users": true
          },
          "return_allowed": true,
          "stackable": true,
          "type": "percentage",
          "usage_meta": null
        }
      ],
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
          "parent_item_identifiers": {
            "identifier": "ZASFF",
            "parent_item_id": 7501190,
            "parent_item_size": "OS"
          },
          "moq": {},
          "delivery_promise": {
            "timestamp": {
              "min": 1646257339,
              "max": 1646689339
            },
            "formatted": {
              "min": "03 Mar, Thursday",
              "max": "08 Mar, Tuesday"
            }
          },
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
            "available_sizes": [
              {
                "is_available": true,
                "display": "OS",
                "value": "OS"
              }
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
          "moq": {},
          "delivery_promise": {
            "timestamp": {
              "min": 1646257339,
              "max": 1646689339
            },
            "formatted": {
              "min": "03 Mar, Thursday",
              "max": "08 Mar, Tuesday"
            }
          },
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
            "available_sizes": [
              {
                "is_available": true,
                "display": "OS",
                "value": "OS"
              }
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
          "moq": {},
          "delivery_promise": {
            "timestamp": {
              "min": 1646257339,
              "max": 1646689339
            },
            "formatted": {
              "min": "03 Mar, Thursday",
              "max": "08 Mar, Tuesday"
            }
          },
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
            "available_sizes": [
              {
                "is_available": true,
                "display": "OS",
                "value": "OS"
              }
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
          "moq": {},
          "delivery_promise": {
            "timestamp": {
              "min": 1646257339,
              "max": 1646689339
            },
            "formatted": {
              "min": "03 Mar, Thursday",
              "max": "08 Mar, Tuesday"
            }
          },
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
            "available_sizes": [
              {
                "is_available": true,
                "display": "OS",
                "value": "OS"
              }
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
      "buy_now": false,
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
          "parent_item_identifiers": {
            "identifier": "ZASFF",
            "parent_item_id": 7501190,
            "parent_item_size": "OS"
          },
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
      "buy_now": false,
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
Update items in the cart

```golang

 data, err :=  Cart.UpdateCart(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |










| xQuery | struct | Includes properties such as `ID`, `I`, `B`, `AreaCode`, `BuyNow`

| body |  UpdateCartRequest | "Request body" 


<p>Use this API to update items added to the cart with the help of a request object containing attributes like item_quantity and item_size. These attributes will be fetched from the following APIs</p> <ul> <li><font color="monochrome">operation</font> Operation for current api call. <b>update_item</b> for update items. <b>remove_item</b> for removing items.</li> <li> <font color="monochrome">item_id</font>  "/platform/content/v1/products/"</li> <li> <font color="monochrome">item_size</font>   "/platform/content/v1/products/:slug/sizes/"</li> <li> <font color="monochrome">quantity</font>  item quantity (must be greater than or equal to 1)</li> <li> <font color="monochrome">article_id</font>   "/content​/v1​/products​/:identifier​/sizes​/price​/"</li> <li> <font color="monochrome">item_index</font>  item position in the cart (must be greater than or equal to 0)</li> </ul>

*Success Response:*



Success. Updates and returns a cart object as shown below. Refer `UpdateCartDetailResponse` for more details.


Schema: `UpdateCartDetailResponse`


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
          "parent_item_identifiers": {
            "identifier": "ZASFF",
            "parent_item_id": 7501190,
            "parent_item_size": "OS"
          },
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
      "buy_now": false,
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
      "applied_promo_details": [
        {
          "application_id": "000000000000000000000001",
          "apply_all_offers": false,
          "apply_exclusive": null,
          "buy_article_dict": {
            "rule#1": {
              "633527e850dfb8e73e6de996_0": {
                "added_in_cart": false,
                "added_quantity": 0,
                "amount_paid": 0,
                "applicable_credits": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 0,
                  "floor_value": 0,
                  "source_currency_code": "FC",
                  "source_currency_symbol": "",
                  "source_value": 0,
                  "value": 0
                },
                "article_assign_status": true,
                "article_assignment": {
                  "level": "multi-companies",
                  "strategy": "optimal"
                },
                "article_error": {
                  "type": null,
                  "value": null,
                  "message": null
                },
                "article_id": "633527e850dfb8e73e6de996",
                "article_index": "0",
                "article_meta": {
                  "article_id": "633527e850dfb8e73e6de996",
                  "bulk_margin": 0,
                  "bulk_discount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "INR",
                    "source_currency_symbol": "₹",
                    "source_value": 0,
                    "value": 0
                  },
                  "bulk_coupon_code": null,
                  "country_of_origin": "India",
                  "dimension": {
                    "width": 70,
                    "unit": "cm",
                    "height": 70,
                    "is_default": true,
                    "length": 70
                  },
                  "fragile": false,
                  "manufacturer": {
                    "address": "10, PALI MALA RD, ADARSH NAGAR, BANDRA WEST,, MAHARASHTRA, MUMBAI",
                    "name": "Sabki Shop",
                    "is_default": true
                  },
                  "weight": {
                    "shipping": 250,
                    "unit": "gram",
                    "is_default": true
                  },
                  "raw_meta": {
                    "fynd_identifier": "4187_3412343256098",
                    "is_set_article": false,
                    "set_quantity": 1
                  },
                  "is_set": false,
                  "set": {},
                  "identifier": {
                    "ean": "3412343256098"
                  },
                  "hsn_code": null,
                  "hsn_code_id": "625fbd96faeed8b3df6ec3ce",
                  "raw_price": {
                    "effective": {
                      "conversion_rate": 1,
                      "currency_code": "INR",
                      "currency_symbol": "₹",
                      "floor_source_value": 3899,
                      "floor_value": 3899,
                      "source_currency_code": "INR",
                      "source_currency_symbol": "₹",
                      "source_value": 3899,
                      "value": 3899
                    },
                    "marked": {
                      "conversion_rate": 1,
                      "currency_code": "INR",
                      "currency_symbol": "₹",
                      "floor_source_value": 3899,
                      "floor_value": 3899,
                      "source_currency_code": "INR",
                      "source_currency_symbol": "₹",
                      "source_value": 3899,
                      "value": 3899
                    },
                    "transfer": {
                      "conversion_rate": 1,
                      "currency_code": "INR",
                      "currency_symbol": "₹",
                      "floor_source_value": 0,
                      "floor_value": 0,
                      "source_currency_code": "INR",
                      "source_currency_symbol": "₹",
                      "source_value": 0,
                      "value": 0
                    }
                  },
                  "seller_identifier": "3412343256098",
                  "_custom_json": {},
                  "tags": [],
                  "return_config": {
                    "time": 0,
                    "unit": "days",
                    "returnable": false
                  }
                },
                "attributes": {
                  "essential": "No",
                  "item_code": "BSFJ2-12608",
                  "gender": [
                    "Men"
                  ],
                  "denim_type": "Solid",
                  "pattern": "Solid",
                  "product_fit": "Slim",
                  "currency": "INR",
                  "media": [
                    {
                      "type": "image",
                      "meta": {
                        "brand": "nike",
                        "item_code": "BSFJ2-12608",
                        "sequence": 0
                      },
                      "url": "https://hdn-1.addsale.com/addsale/products/pictures/item/free/original/nike/BSFJ2-12608/0/jMpkpyx17w-GWQjymow3-Black-Slim-Fit-Jeans.jpeg"
                    }
                  ],
                  "name": "Nike Jeans 12608",
                  "brand_name": "Nike"
                },
                "avl_qty": 97,
                "brand_id": 18,
                "bulk_coupon_applied": false,
                "bulk_coupon_code": null,
                "bulk_discount": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 0,
                  "floor_value": 0,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 0,
                  "value": 0
                },
                "bulk_margin": 0,
                "cancellation_allowed": true,
                "cashback": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "FC",
                    "source_currency_symbol": "",
                    "source_value": 0,
                    "value": 0
                  }
                },
                "category_id": [
                  150
                ],
                "cod_charges": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "INR",
                    "source_currency_symbol": "₹",
                    "source_value": 0,
                    "value": 0
                  }
                },
                "company_id": 2,
                "company_info": {
                  "company_id": 2,
                  "c_name": "Sabki Shop",
                  "c_taxes": [
                    {
                      "rate": 15,
                      "enable": true,
                      "effective_date": "2022-06-05T18:29:23.904000"
                    }
                  ],
                  "company_cin": "U45200MH1992PTC066474"
                },
                "company_taxes": [
                  {
                    "rate": 15,
                    "enable": true,
                    "effective_date": "2022-06-05T18:29:23.904000"
                  }
                ],
                "coupon": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "INR",
                    "source_currency_symbol": "₹",
                    "source_value": 0,
                    "value": 0
                  },
                  "article_count": 0,
                  "cancellation_allowed": true,
                  "return_allowed": true
                },
                "currency_dict": {
                  "code": "INR",
                  "rate": 1
                },
                "custom_meta": {},
                "data": {},
                "delivery_charges": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "INR",
                    "source_currency_symbol": "₹",
                    "source_value": 0,
                    "value": 0
                  }
                },
                "delivery_error_msg": "",
                "delivery_promise": {
                  "timestamp": {
                    "min": 1646257339,
                    "max": 1646689339
                  },
                  "formatted": {
                    "min": "03 Mar, Thursday",
                    "max": "08 Mar, Tuesday"
                  }
                },
                "departments": [
                  99
                ],
                "discount": 0,
                "discount_applied": {},
                "dt_currency": "INR",
                "dt_currency_symbol": "₹",
                "esp_modified": false,
                "extra_meta": {},
                "float_amount_paid": 0,
                "free_gift_items_meta": {},
                "group_id": "",
                "gst_amount": 167.1,
                "gst_tax_percentage": 5,
                "hsn_code": null,
                "hsn_code_id": "625fbd96faeed8b3df6ec3ce",
                "identifiers": {
                  "identifier": "BDnwAINORr6SlTKxZN3w"
                },
                "image_nature": "standard",
                "include": true,
                "is_valid": true,
                "item_id": 75660796,
                "item_size": "34",
                "last_update_at": 1680852767,
                "meta": {},
                "net_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3509.1,
                  "floor_value": 3509.1,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3509.1,
                  "value": 3509.1
                },
                "old_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "original_price_effective": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "original_unit_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "parent_item_identifiers": {
                  "identifier": null,
                  "parent_item_size": null,
                  "parent_item_id": null
                },
                "pickup_stores": [],
                "price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "price_effective": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "price_marked": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "product_tags": [],
                "promotions": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "INR",
                    "source_currency_symbol": "₹",
                    "source_value": 0,
                    "value": 0
                  },
                  "applied": [
                    {
                      "amount": {
                        "conversion_rate": 1,
                        "currency_code": "INR",
                        "currency_symbol": "₹",
                        "floor_source_value": 389.9,
                        "floor_value": 389.9,
                        "source_currency_code": "INR",
                        "source_currency_symbol": "₹",
                        "source_value": 389.9,
                        "value": 389.9
                      },
                      "applied_discounts": [
                        {
                          "type": "percentage",
                          "value": 10
                        }
                      ],
                      "applied_free_articles": [],
                      "apply_exclusive": null,
                      "article_quantity": 1,
                      "buy_rules": [
                        {
                          "all_items": null,
                          "cart_conditions": {},
                          "item_criteria": {
                            "item_brand": [
                              18
                            ]
                          },
                          "mrp_promo": false,
                          "slug": "rule#1"
                        }
                      ],
                      "cancellation_allowed": true,
                      "discount_rules": [
                        {
                          "all_items": null,
                          "buy_condition": "( rule#1 )",
                          "item_criteria": {
                            "buy_rules": [
                              "rule#1"
                            ]
                          },
                          "matched_buy_rules": [
                            "rule#1"
                          ],
                          "offer": {
                            "discount_percentage": 10
                          },
                          "raw_offer": {
                            "item_criteria": {
                              "buy_rules": [
                                "rule#1"
                              ]
                            },
                            "buy_condition": "( rule#1 )",
                            "discount_type": "percentage",
                            "offer": {
                              "discount_percentage": 10
                            }
                          },
                          "type": "percentage"
                        }
                      ],
                      "mrp_promo": false,
                      "offer_text": "10% of on NIKE Jeans",
                      "ownership": {
                        "payable_category": "seller",
                        "payable_by": ""
                      },
                      "payable_category": "seller",
                      "promo_code": null,
                      "promo_id": "642fe329ed159b7aa743a293",
                      "promo_type": "percentage",
                      "promotion_group": "product",
                      "return_allowed": true
                    }
                  ],
                  "apply_exclusive": null,
                  "available": [],
                  "cancellation_allowed": true,
                  "exclusive_promo_applied": false,
                  "mrp_promo_applied": false,
                  "return_allowed": true
                },
                "quantity": 1,
                "quantity_assign_status": false,
                "referral_credits": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "FC",
                    "source_currency_symbol": "",
                    "source_value": 0,
                    "value": 0
                  }
                },
                "reload_data_only": true,
                "return_allowed": false,
                "sc_currency": "INR",
                "sc_currency_symbol": "₹",
                "selling_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "service_item_meta": {
                  "product_group_tags": null,
                  "products": null
                },
                "size_level_total_qty": 1,
                "sizes": [
                  "34"
                ],
                "split_article_id": "633527e850dfb8e73e6de996_0",
                "split_index": 0,
                "store_id": 4187,
                "store_info": {
                  "store_id": 4187,
                  "s_city": "HYDERABAD",
                  "store_name": "AND Inorbit Hyderabad",
                  "store_type": "mall",
                  "store_pincode": 500081,
                  "latitude": 17.4343693,
                  "longitude": 78.3866087
                },
                "strategy_validation_data": {},
                "tags": [
                  "rule#1"
                ],
                "total_gst_amount": 167.1,
                "total_value_of_good": 3342,
                "transfer_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 0,
                  "floor_value": 0,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 0,
                  "value": 0
                },
                "unit_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3509.1,
                  "floor_value": 3509.1,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3509.1,
                  "value": 3509.1
                },
                "valid_inventory": true,
                "value_of_good": 3342,
                "verify_article": false
              }
            }
          },
          "buy_rules": [
            {
              "all_items": null,
              "cart_conditions": {},
              "item_criteria": {
                "item_brand": [
                  18
                ]
              },
              "mrp_promo": false,
              "slug": "rule#1"
            }
          ],
          "buy_rules_map": {
            "rule#1": {
              "item_brand": [
                18
              ]
            }
          },
          "calculate_on": "esp",
          "cancellation_allowed": true,
          "code": null,
          "discount": 389.9,
          "discount_rules": [
            {
              "all_items": null,
              "buy_condition": "( rule#1 )",
              "item_criteria": {
                "buy_rules": [
                  "rule#1"
                ]
              },
              "matched_buy_rules": [
                "rule#1"
              ],
              "offer": {
                "discount_percentage": 10
              },
              "raw_offer": {
                "item_criteria": {
                  "buy_rules": [
                    "rule#1"
                  ]
                },
                "buy_condition": "( rule#1 )",
                "discount_type": "percentage",
                "offer": {
                  "discount_percentage": 10
                }
              },
              "type": "percentage"
            }
          ],
          "id": "642fe329ed159b7aa743a293",
          "mrp_promo": false,
          "offer_text": "10% of on NIKE Jeans",
          "ownership": {
            "payable_category": "seller",
            "payable_by": ""
          },
          "payable_category": "seller",
          "priority": 1,
          "promo_group": "product",
          "remaining_allowed_qty": null,
          "restrictions": {
            "uses": {
              "maximum": {
                "user": 0,
                "total": 0
              },
              "remaining": {
                "user": 0,
                "total": 0
              }
            },
            "user_groups": [],
            "post_order": {
              "return_allowed": true,
              "cancellation_allowed": true
            },
            "user_id": [],
            "payments": [],
            "user_registered": {
              "start": null,
              "end": null
            },
            "platforms": [
              "web",
              "android",
              "ios"
            ],
            "anonymous_users": true
          },
          "return_allowed": true,
          "stackable": true,
          "type": "percentage",
          "usage_meta": null
        }
      ],
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
          "parent_item_identifiers": {
            "identifier": "ZASFF",
            "parent_item_id": 7501190,
            "parent_item_size": "OS"
          },
          "moq": {},
          "delivery_promise": {
            "timestamp": {
              "min": 1646257339,
              "max": 1646689339
            },
            "formatted": {
              "min": "03 Mar, Thursday",
              "max": "08 Mar, Tuesday"
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
            "available_sizes": [
              {
                "is_available": true,
                "display": "OS",
                "value": "OS"
              }
            ],
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
      "buy_now": false,
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


#### deleteCart
Delete cart once user made successful checkout

```golang

 data, err :=  Cart.DeleteCart(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `ID`



Use this API to delete the cart.

*Success Response:*



Success. Returns whether the cart has been deleted or not.


Schema: `DeleteCartDetailResponse`









---


#### getItemCount
Count items in the cart

```golang

 data, err :=  Cart.GetItemCount(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `ID`, `BuyNow`



Use this API to get the total number of items present in cart.

*Success Response:*



Success. Returns the total count of items in a user's cart.


Schema: `CartItemCountResponse`









---


#### getCoupons
Fetch Coupon

```golang

 data, err :=  Cart.GetCoupons(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `ID`, `BuyNow`



Use this API to get a list of available coupons along with their details.

*Success Response:*



Success. Returns a coupon object which has a list of all the eligible coupons. Refer `GetCouponResponse` for more details.


Schema: `GetCouponResponse`









---


#### applyCoupon
Apply Coupon

```golang

 data, err :=  Cart.ApplyCoupon(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |










| xQuery | struct | Includes properties such as `I`, `B`, `P`, `ID`, `BuyNow`

| body |  ApplyCouponRequest | "Request body" 


Use this API to apply coupons on items in the cart.

*Success Response:*



Success. Returns coupons applied to the cart along with item details and price breakup. Refer `CartDetailResponse` for more details.


Schema: `CartDetailResponse`









---


#### removeCoupon
Remove Coupon Applied

```golang

 data, err :=  Cart.RemoveCoupon(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `ID`, `BuyNow`



Remove Coupon applied on the cart by passing uid in request body.

*Success Response:*



Success. Returns coupons removed from the cart along with item details and price breakup. Refer `CartDetailResponse` for more details.


Schema: `CartDetailResponse`









---


#### getBulkDiscountOffers
Get discount offers based on quantity

```golang

 data, err :=  Cart.GetBulkDiscountOffers(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |








| xQuery | struct | Includes properties such as `ItemID`, `ArticleID`, `UID`, `Slug`



Use this API to get a list of applicable offers along with current, next and best offer for given product. Either one of uid, item_id, slug should be present.

*Success Response:*



Success. Returns a data object containing the seller details and available offers (if exists) on bulk products. Refer `BulkPriceResponse` for more details.


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









---


#### applyRewardPoints
Apply reward points at cart

```golang

 data, err :=  Cart.ApplyRewardPoints(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |








| xQuery | struct | Includes properties such as `ID`, `I`, `B`, `BuyNow`

| body |  RewardPointRequest | "Request body" 


Use this API to redeem a fixed no. of reward points by applying it to the cart.

*Success Response:*



Success. Returns a Cart object. Check the example shown below or refer `CartDetailResponse` for more details.


Schema: `CartDetailResponse`









---


#### getAddresses
Fetch address

```golang

 data, err :=  Cart.GetAddresses(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |












| xQuery | struct | Includes properties such as `CartID`, `BuyNow`, `MobileNo`, `CheckoutMode`, `Tags`, `IsDefault`



Use this API to get all the addresses associated with an account. If successful, returns a Address resource in the response body specified in GetAddressesResponse.attibutes listed below are optional <ul> <li> <font color="monochrome">uid</font></li> <li> <font color="monochrome">address_id</font></li> <li> <font color="monochrome">mobile_no</font></li> <li> <font color="monochrome">checkout_mode</font></li> <li> <font color="monochrome">tags</font></li> <li> <font color="monochrome">default</font></li> </ul>

*Success Response:*



Success. Returns an Address object containing a list of address saved in the account. Refer `GetAddressesResponse` for more details.


Schema: `GetAddressesResponse`









---


#### addAddress
Add address to an account

```golang

 data, err :=  Cart.AddAddress(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  Address | "Request body" 


Use this API to add an address to an account.

*Success Response:*



Success. Returns the address ID, a flag whether the address is set as default, and a success message. Refer `SaveAddressResponse` for more details.


Schema: `SaveAddressResponse`









---


#### getAddressById
Fetch a single address by its ID

```golang

 data, err :=  Cart.GetAddressById(ID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ID | string |  | 













| xQuery | struct | Includes properties such as `CartID`, `BuyNow`, `MobileNo`, `CheckoutMode`, `Tags`, `IsDefault`



Use this API to get an addresses using its ID. If successful, returns a Address resource in the response body specified in `Address`. Attibutes listed below are optional <ul> <li> <font color="monochrome">mobile_no</font></li> <li> <font color="monochrome">checkout_mode</font></li> <li> <font color="monochrome">tags</font></li> <li> <font color="monochrome">default</font></li> </ul>

*Success Response:*



Success. Returns an Address object containing a list of address saved in the account. Refer `Address` for more details.


Schema: `Address`









---


#### updateAddress
Update address added to an account

```golang

 data, err :=  Cart.UpdateAddress(ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ID | string | ID allotted to the selected address | 


| body |  Address | "Request body" 


<p>Use this API to update an existing address in the account. Request object should contain attributes mentioned in  <font color="blue">Address </font> can be updated. These attributes are:</p> <ul> <li> <font color="monochrome">is_default_address</font></li> <li> <font color="monochrome">landmark</font></li> <li> <font color="monochrome">area</font></li> <li> <font color="monochrome">pincode</font></li> <li> <font color="monochrome">email</font></li> <li> <font color="monochrome">address_type</font></li> <li> <font color="monochrome">name</font></li> <li> <font color="monochrome">address_id</font></li> <li> <font color="monochrome">address</font></li> </ul>

*Success Response:*



Success. Returns the address ID and a message indicating a successful address updation.


Schema: `UpdateAddressResponse`









---


#### removeAddress
Remove address associated with an account

```golang

 data, err :=  Cart.RemoveAddress(ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ID | string | ID allotted to the selected address | 




Use this API to delete an address by its ID. This will returns an object that will indicate whether the address was deleted successfully or not.

*Success Response:*



Returns a Status object indicating the success or failure of address deletion.


Schema: `DeleteAddressResponse`









---


#### selectAddress
Select an address from available addresses

```golang

 data, err :=  Cart.SelectAddress(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |








| xQuery | struct | Includes properties such as `CartID`, `BuyNow`, `I`, `B`

| body |  SelectCartAddressRequest | "Request body" 


<p>Select Address from all addresses associated with the account in order to ship the cart items to that address, otherwise default address will be selected implicitly. See `SelectCartAddressRequest` in schema of request body for the list of attributes needed to select Address from account. On successful request, this API returns a Cart object. Below address attributes are required. <ul> <li> <font color="monochrome">address_id</font></li> <li> <font color="monochrome">billing_address_id</font></li> <li> <font color="monochrome">uid</font></li> </ul></p>

*Success Response:*



Success. Returns a Cart object as shown below. Refer `CartDetailResponse` for more details.  .


Schema: `CartDetailResponse`









---


#### selectPaymentMode
Update cart payment

```golang

 data, err :=  Cart.SelectPaymentMode(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `ID`, `BuyNow`

| body |  UpdateCartPaymentRequest | "Request body" 


Use this API to update cart payment.

*Success Response:*



Success. Returns a Cart object as shown below. Refer `CartDetailResponse` for more details.


Schema: `CartDetailResponse`









---


#### validateCouponForPayment
Verify the coupon eligibility against the payment mode

```golang

 data, err :=  Cart.ValidateCouponForPayment(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |














| xQuery | struct | Includes properties such as `ID`, `BuyNow`, `AddressID`, `PaymentMode`, `PaymentIdentifier`, `AggregatorName`, `MerchantCode`



Use this API to validate a coupon against the payment mode such as NetBanking, Wallet, UPI etc.

*Success Response:*



Success. Returns a success message and the coupon validity. Refer `PaymentCouponValidate` for more details.


Schema: `PaymentCouponValidate`









---


#### getShipments
Get delivery date and options before checkout

```golang

 data, err :=  Cart.GetShipments(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |










| xQuery | struct | Includes properties such as `P`, `ID`, `BuyNow`, `AddressID`, `AreaCode`



Use this API to get shipment details, expected delivery date, items and price breakup of the shipment.

*Success Response:*



Success. Returns delivery promise along with shipment details and price breakup. Refer `CartShipmentsResponse` for more details.


Schema: `CartShipmentsResponse`


*Examples:*


Shipment Generated
```json
{
  "value": {
    "items": [],
    "buy_now": false,
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
    "buy_now": false,
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









---


#### checkoutCart
Checkout all items in the cart

```golang

 data, err :=  Cart.CheckoutCart(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `BuyNow`

| body |  CartCheckoutDetailRequest | "Request body" 


Use this API to checkout all items in the cart for payment and order generation. For COD, order will be directly generated, whereas for other checkout modes, user will be redirected to a payment gateway.

*Success Response:*



Success. Returns the status of cart checkout. Refer `CartCheckoutResponseSchema` for more details.


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
Update the cart meta

```golang

 data, err :=  Cart.UpdateCartMeta(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `ID`, `BuyNow`

| body |  CartMetaRequest | "Request body" 


Use this API to update cart meta like checkout_mode and gstin.

*Success Response:*



Returns a message indicating the success of cart meta updation as shown below.


Schema: `CartMetaResponse`









---


#### getCartShareLink
Generate token for sharing the cart

```golang

 data, err :=  Cart.GetCartShareLink(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  GetShareCartLinkRequest | "Request body" 


Use this API to generate a shared cart snapshot and return a shortlink token. The link can be shared with other users for getting the same items in their cart.

*Success Response:*



Returns a URL to share and a token as shown below.


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
Get details of a shared cart

```golang

 data, err :=  Cart.GetCartSharedItems(Token);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Token | string | Token of the shared short link | 




Use this API to get the shared cart details as per the token generated using the share-cart API.

*Success Response:*



Success. Returns a Cart object as per the valid token. Refer `SharedCartResponse` for more details.


Schema: `SharedCartResponse`









---


#### updateCartWithSharedItems
Merge or replace existing cart

```golang

 data, err :=  Cart.UpdateCartWithSharedItems(Token, Action);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Token | string | Token of the shared short link | 


| Action | string | Operation to perform on the existing cart merge or replace. | 




Use this API to merge the shared cart with existing cart, or replace the existing cart with the shared cart. The `action` parameter is used to indicate the operation Merge or Replace.

*Success Response:*



Success. Returns a merged or replaced cart as per the valid token. Refer `SharedCartResponse` for more details.


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


#### getPromotionOffers
Fetch available promotions

```golang

 data, err :=  Cart.GetPromotionOffers(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |






| xQuery | struct | Includes properties such as `Slug`, `PageSize`, `PromotionGroup`



Use this API to get top 5 offers available for current product

*Success Response:*



Success. Returns a array containing the available offers (if exists) on product via promotions. Refer `PromotionOffersResponse` for more details.


Schema: `PromotionOffersResponse`









---


#### getLadderOffers
Fetch ladder price promotion

```golang

 data, err :=  Cart.GetLadderOffers(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |








| xQuery | struct | Includes properties such as `Slug`, `StoreID`, `PromotionID`, `PageSize`



Use this API to get applicable ladder price promotion for current product

*Success Response:*



Success. Returns a object containing the applicable ladder price offers (if exists) on product. Refer `PromotionOffersResponse` for more details.


Schema: `LadderPriceOffers`









---



---


## Common


#### searchApplication
Search Application

```golang

 data, err :=  Common.SearchApplication(Authorization, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Authorization | string |  | 



| xQuery | struct | Includes properties such as `Query`



Provide application name or domain url

*Success Response:*



Success


Schema: `ApplicationResponse`









---


#### getLocations
Get countries, states, cities

```golang

 data, err :=  Common.GetLocations(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `LocationType`, `ID`





*Success Response:*



Success


Schema: `Locations`









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
          "phone": "9999999999"
        }
      ],
      "firstName": "Nikhil",
      "lastName": "Manapure",
      "emails": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "email": "niktest@xyz.com"
        }
      ],
      "username": "niktest_xyz_com_38425_20500281",
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
            "phone": "9999999999"
          }
        ],
        "firstName": "Nikhil",
        "lastName": "Manapure",
        "emails": [
          {
            "active": true,
            "primary": true,
            "verified": true,
            "email": "niktest@xyz.com"
          }
        ],
        "username": "niktest_xyz_com_38425_20500281",
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
          "phone": "9999999999"
        }
      ],
      "firstName": "Nikhil",
      "lastName": "Manapure",
      "emails": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "email": "niktest@xyz.com"
        }
      ],
      "username": "niktest_xyz_com_38425_20500281",
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
            "phone": "9999999999"
          }
        ],
        "firstName": "Nikhil",
        "lastName": "Manapure",
        "emails": [
          {
            "active": true,
            "primary": true,
            "verified": true,
            "email": "niktest@xyz.com"
          }
        ],
        "username": "niktest_xyz_com_38425_20500281",
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
            "phone": "9999999999"
          }
        ],
        "firstName": "Nikhil",
        "lastName": "Manapure",
        "emails": [
          {
            "active": true,
            "primary": true,
            "verified": true,
            "email": "niktest@xyz.com"
          }
        ],
        "username": "niktest_xyz_com_38425_20500281",
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
              "phone": "9999999999"
            }
          ],
          "firstName": "Nikhil",
          "lastName": "Manapure",
          "emails": [
            {
              "active": true,
              "primary": true,
              "verified": true,
              "email": "niktest@xyz.com"
            }
          ],
          "username": "niktest_xyz_com_38425_20500281",
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


#### getAllPages
Get all pages of a theme

```golang

 data, err :=  Theme.GetAllPages(ThemeID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ThemeID | string | ID of the theme to be retrieved | 




Use this API to retrieve all the available pages of a theme by its ID.

*Success Response:*



Success. Returns an array all the pages of the theme. Refer `AllAvailablePageSchema` for more details.


Schema: `AllAvailablePageSchema`


*Examples:*


All pages
```json
{
  "value": {
    "pages": [
      {
        "path": "products",
        "type": "system",
        "seo": {
          "title": "",
          "description": "",
          "_id": "60210832d7e981639e13f6b2"
        },
        "_id": "60ab5ca6d572fed64294eb0e",
        "value": "product-listing",
        "text": "Product Listing",
        "theme": "5fb3ee4194a5181feeeba8e5",
        "sections_meta": [],
        "__v": 9
      },
      {
        "path": "collection",
        "type": "system",
        "seo": {
          "title": "",
          "description": "",
          "_id": "60210832d7e981fc5d13f6b9"
        },
        "_id": "60ab5ca6d572fed64294eaf9",
        "text": "Collection Listing",
        "value": "collection-listing",
        "theme": "5fb3ee4194a5181feeeba8e5",
        "sections_meta": [],
        "__v": 9
      },
      {
        "path": "compare",
        "type": "system",
        "seo": {
          "title": "",
          "description": "",
          "_id": "60210832d7e981cbca13f6b1"
        },
        "_id": "60ab5ca6d572fed64294eb0b",
        "value": "compare-products",
        "text": "Compare Products",
        "theme": "5fb3ee4194a5181feeeba8e5",
        "sections_meta": [],
        "__v": 9
      },
      {
        "path": "cart/bag",
        "type": "system",
        "seo": {
          "title": "",
          "description": "",
          "_id": "60210832d7e9812fdf13f6ae"
        },
        "_id": "60ab5ca6d572fed64294eb02",
        "value": "cart-landing",
        "text": "Cart Landing",
        "theme": "5fb3ee4194a5181feeeba8e5",
        "sections_meta": [],
        "__v": 9
      },
      {
        "path": "product",
        "type": "system",
        "seo": {
          "title": "",
          "description": "",
          "_id": "60210832d7e9815c9713f6ab"
        },
        "_id": "60ab5ca6d572fed64294eaf6",
        "text": "Product Description",
        "value": "product-description",
        "theme": "5fb3ee4194a5181feeeba8e5",
        "sections_meta": [],
        "__v": 9
      },
      {
        "path": "product/:slug/reviews",
        "type": "system",
        "seo": {
          "title": "",
          "description": "",
          "_id": "60ab5ca6d572fed64294eb24"
        },
        "_id": "60ab5ca6d572fed64294eb25",
        "sections_meta": [],
        "value": "product-reviews",
        "text": "Product Reviews",
        "theme": "5fb3ee4194a5181feeeba8e5",
        "__v": 9
      },
      {
        "path": "blog",
        "type": "system",
        "seo": {
          "title": "",
          "description": "",
          "_id": "60ab5ca6d572fed64294eb22"
        },
        "_id": "60ab5ca6d572fed64294eb23",
        "sections_meta": [],
        "value": "blog",
        "text": "Blog",
        "theme": "5fb3ee4194a5181feeeba8e5",
        "__v": 9
      },
      {
        "path": "sections/cookie",
        "type": "sections",
        "seo": {
          "title": "",
          "description": "",
          "_id": "60210832d7e9814fed13f6b5"
        },
        "_id": "60ab5ca6d572fed64294eb17",
        "text": "cookie",
        "value": "cookie",
        "theme": "5fb3ee4194a5181feeeba8e5",
        "sections_meta": [],
        "__v": 9
      },
      {
        "path": "sections/vivek",
        "type": "sections",
        "seo": {
          "title": "",
          "description": "",
          "_id": "60210832d7e981b32713f6b6"
        },
        "_id": "60ab5ca6d572fed64294eb1a",
        "text": "vivek",
        "value": "vivek",
        "theme": "5fb3ee4194a5181feeeba8e5",
        "sections_meta": [],
        "__v": 9
      },
      {
        "path": "about-us",
        "type": "system",
        "seo": {
          "title": "",
          "description": "",
          "_id": "60ab5ca6d572fed64294eb28"
        },
        "_id": "60ab5ca6d572fed64294eb29",
        "sections_meta": [],
        "value": "about-us",
        "text": "About Us",
        "theme": "5fb3ee4194a5181feeeba8e5",
        "__v": 9
      },
      {
        "path": "wishlist",
        "type": "system",
        "seo": {
          "title": "",
          "description": "",
          "_id": "60210832d7e981dd2d13f6b3"
        },
        "_id": "60ab5ca6d572fed64294eb11",
        "value": "wishlist",
        "text": "Wishlist",
        "theme": "5fb3ee4194a5181feeeba8e5",
        "sections_meta": [],
        "__v": 9
      },
      {
        "path": "product/:slug/add-review",
        "type": "system",
        "seo": {
          "title": "",
          "description": "",
          "_id": "60ab5ca6d572fed64294eb26"
        },
        "_id": "60ab5ca6d572fed64294eb27",
        "sections_meta": [],
        "value": "add-product-review",
        "text": "Add Product Review",
        "theme": "5fb3ee4194a5181feeeba8e5",
        "__v": 9
      },
      {
        "path": "brands",
        "type": "system",
        "seo": {
          "title": "",
          "description": "",
          "_id": "60210832d7e981161a13f6ad"
        },
        "_id": "60ab5ca6d572fed64294eaff",
        "value": "brands",
        "text": "Brands",
        "theme": "5fb3ee4194a5181feeeba8e5",
        "sections_meta": [],
        "__v": 9
      },
      {
        "path": "",
        "type": "system",
        "seo": {
          "title": "",
          "description": "",
          "_id": "60210832d7e98115b013f6ac"
        },
        "_id": "60ab5ca6d572fed64294eafc",
        "value": "home",
        "text": "Home",
        "theme": "5fb3ee4194a5181feeeba8e5",
        "sections_meta": [],
        "__v": 9
      },
      {
        "path": "collections",
        "type": "system",
        "seo": {
          "title": "",
          "description": "",
          "_id": "60210832d7e981ad0b13f6b0"
        },
        "_id": "60ab5ca6d572fed64294eb08",
        "value": "collections",
        "text": "Collections",
        "theme": "5fb3ee4194a5181feeeba8e5",
        "sections_meta": [],
        "__v": 9
      },
      {
        "path": "categories",
        "type": "system",
        "seo": {
          "title": "",
          "description": "",
          "_id": "60210832d7e981872c13f6af"
        },
        "_id": "60ab5ca6d572fed64294eb05",
        "value": "categories",
        "text": "Categories",
        "theme": "5fb3ee4194a5181feeeba8e5",
        "sections_meta": [],
        "__v": 9
      },
      {
        "path": "sections/test",
        "type": "sections",
        "seo": {
          "title": "",
          "description": "",
          "_id": "60210832d7e98177f713f6b4"
        },
        "_id": "60ab5ca6d572fed64294eb14",
        "text": "Test",
        "value": "test",
        "theme": "5fb3ee4194a5181feeeba8e5",
        "sections_meta": [],
        "__v": 9
      },
      {
        "path": "sections/vinit",
        "type": "sections",
        "seo": {
          "title": "",
          "description": "",
          "_id": "60210832d7e98170b813f6b8"
        },
        "_id": "60ab5ca6d572fed64294eb20",
        "text": "vinit",
        "value": "vinit",
        "theme": "5fb3ee4194a5181feeeba8e5",
        "sections_meta": [],
        "__v": 9
      },
      {
        "path": "sections/maggie",
        "type": "sections",
        "seo": {
          "title": "",
          "description": "",
          "_id": "60210832d7e981469613f6b7"
        },
        "_id": "60ab5ca6d572fed64294eb1d",
        "text": "maggie",
        "value": "maggie",
        "theme": "5fb3ee4194a5181feeeba8e5",
        "sections_meta": [],
        "__v": 9
      }
    ]
  }
}
```









---


#### getPage
Get page of a theme

```golang

 data, err :=  Theme.GetPage(ThemeID, PageValue);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ThemeID | string | ID of the theme to be retrieved | 


| PageValue | string | Value of the page to be retrieved | 




Use this API to retrieve a page of a theme.

*Success Response:*



Success. Returns an object of the pages.  Refer `AvailablePageSchema` for more details.


Schema: `AvailablePageSchema`


*Examples:*


Home page
```json
{
  "value": {
    "path": "",
    "type": "system",
    "seo": {
      "title": "",
      "description": "",
      "_id": "60210832d7e98115b013f6ac"
    },
    "props": [],
    "_id": "60ab5ca6d572fed64294eafc",
    "sections": [
      {
        "blocks": [],
        "predicate": {
          "screen": {
            "mobile": false,
            "desktop": false,
            "tablet": false
          },
          "user": {
            "authenticated": true,
            "anonymous": true
          },
          "route": {
            "selected": "none",
            "exactURL": "",
            "query": {
              "utm": "facebook"
            }
          }
        },
        "name": "customHtml",
        "props": {
          "code": {
            "type": "code",
            "value": "<p style=\"text-align:right;\"><a href=\"tel:+919820204442\"><img src=\"https://dabuttonfactory.com/button.png?t=For+any+queries%2C+call%3A+%2B91+98202+04442&f=Open+Sans-Bold&ts=21&tc=fff&hp=20&vp=15&c=round&bgt=unicolored&bgc=7043f7\"/> </a></p>"
          }
        }
      },
      {
        "blocks": [],
        "predicate": {
          "screen": {
            "mobile": true,
            "desktop": true,
            "tablet": true
          },
          "user": {
            "authenticated": true,
            "anonymous": true
          },
          "route": {
            "selected": "none",
            "exactURL": "",
            "query": {
              "udm": "vivek"
            }
          }
        },
        "name": "brands-listing",
        "props": {
          "title": {
            "type": "text",
            "value": "Popular rrrrr"
          },
          "header": {
            "type": "header"
          },
          "brand_type": {
            "value": "all",
            "type": "radio"
          },
          "department": {
            "type": "department"
          },
          "item_count": {
            "value": 5,
            "type": "range"
          },
          "layout": {
            "value": "horizontal",
            "type": "select"
          },
          "view_all": {
            "value": false,
            "type": "checkbox"
          }
        }
      },
      {
        "blocks": [
          {
            "type": "product",
            "name": "Product",
            "props": {
              "product": {
                "type": "product"
              }
            }
          },
          {
            "type": "product",
            "name": "Product",
            "props": {
              "product": {
                "type": "product"
              }
            }
          },
          {
            "type": "product",
            "name": "Product",
            "props": {
              "product": {
                "type": "product"
              }
            }
          },
          {
            "type": "product",
            "name": "Product",
            "props": {
              "product": {
                "type": "product"
              }
            }
          }
        ],
        "predicate": {
          "screen": {
            "mobile": true,
            "desktop": true,
            "tablet": true
          },
          "user": {
            "authenticated": false,
            "anonymous": true
          },
          "route": {
            "selected": "none",
            "exactURL": ""
          }
        },
        "name": "featuredProducts",
        "props": {
          "heading": {
            "value": "Featured Products",
            "type": "text"
          },
          "item_count": {
            "type": "range",
            "value": 4
          },
          "full_width": {
            "value": false,
            "type": "checkbox"
          }
        },
        "preset": {
          "blocks": [
            {
              "name": "Product"
            },
            {
              "name": "Product"
            },
            {
              "name": "Product"
            },
            {
              "name": "Product"
            }
          ]
        }
      },
      {
        "blocks": [],
        "predicate": {
          "screen": {
            "mobile": true,
            "desktop": true,
            "tablet": true
          },
          "user": {
            "authenticated": true,
            "anonymous": true
          },
          "route": {
            "selected": "none",
            "exactURL": "",
            "query": {
              "": ""
            }
          }
        },
        "name": "categoryListPage",
        "props": {
          "heading": {
            "type": "text",
            "value": "Explore Categories"
          },
          "layout": {
            "type": "select",
            "value": "grid"
          },
          "item_count": {
            "value": 5,
            "type": "range"
          },
          "view_all": {
            "value": false,
            "type": "checkbox"
          }
        }
      },
      {
        "blocks": [],
        "predicate": {
          "screen": {
            "mobile": true,
            "desktop": true,
            "tablet": true
          },
          "user": {
            "authenticated": true,
            "anonymous": true
          },
          "route": {
            "selected": "none",
            "exactURL": "",
            "query": {
              "": ""
            }
          }
        },
        "name": "heroBanner",
        "props": {
          "ctaLink": {
            "type": "url",
            "value": "https://uniket.hostx0.de/about-us"
          }
        }
      },
      {
        "blocks": [],
        "predicate": {
          "screen": {
            "mobile": true,
            "desktop": true,
            "tablet": true
          },
          "user": {
            "authenticated": true,
            "anonymous": true
          },
          "route": {
            "selected": "none",
            "exactURL": "",
            "query": {
              "": ""
            }
          }
        },
        "name": "imageBanner",
        "props": {
          "image": {
            "value": "",
            "type": "image_picker"
          },
          "full_width": {
            "value": false,
            "type": "checkbox"
          },
          "overlayLayout": {
            "value": "left",
            "type": "select"
          },
          "overlayImage": {
            "value": "",
            "type": "image_picker"
          },
          "text": {
            "value": "",
            "type": "text"
          },
          "text_color": {
            "value": "#000",
            "type": "color"
          },
          "ctaLink": {
            "value": "",
            "type": "url"
          },
          "ctaText": {
            "value": "",
            "type": "text"
          },
          "layout": {
            "type": "select",
            "value": "full"
          },
          "height": {
            "type": "select",
            "value": "h-auto"
          }
        }
      },
      {
        "blocks": [],
        "predicate": {
          "screen": {
            "mobile": true,
            "desktop": true,
            "tablet": true
          },
          "user": {
            "authenticated": true,
            "anonymous": true
          },
          "route": {
            "selected": "none",
            "exactURL": ""
          }
        },
        "name": "brands-listing",
        "props": {
          "title": {
            "type": "text",
            "value": "asdfasdf"
          },
          "header": {
            "type": "header"
          },
          "brand_type": {
            "value": "all",
            "type": "radio"
          },
          "department": {
            "type": "department"
          },
          "item_count": {
            "value": 5,
            "type": "range"
          },
          "layout": {
            "value": "horizontal",
            "type": "select"
          },
          "view_all": {
            "value": false,
            "type": "checkbox"
          }
        }
      }
    ],
    "value": "home",
    "text": "Home",
    "theme": "5fb3ee4194a5181feeeba8e5",
    "sections_meta": [],
    "__v": 9
  }
}
```









---


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
  "value": {
    "information": {
      "images": {
        "desktop": [
          "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/pictures/free/original/artisan-desktop.png"
        ],
        "android": [
          "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/pictures/free/original/artisan-mobile.png"
        ],
        "ios": [
          "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/pictures/free/original/artisan-mobile.png"
        ],
        "thumbnail": [
          "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/pictures/free/original/artisan-desktop.png"
        ]
      },
      "features": [
        "Responsive"
      ],
      "name": "Akash-Artisan"
    },
    "src": {
      "link": "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/sources/J-1s-N-pl-archive.zip"
    },
    "assets": {
      "css": {
        "link": "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/assets/Ia4m885Mw2-themeBundle.css"
      },
      "umd_js": {
        "link": "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/assets/nEPGyc15g-themeBundle.umd.min.js"
      },
      "common_js": {
        "link": "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/assets/5f7tOB5rpq-themeBundle.common.js"
      }
    },
    "config": {
      "preset": {
        "sections": [
          {
            "page_sections": [
              {
                "blocks": [
                  {
                    "type": "gallery_image",
                    "name": "Image",
                    "props": {
                      "image": {
                        "type": "image_picker",
                        "value": "https://hdn-1.fynd.com/company/163/applications/5d5265c8f8ad9cae6dbf18f1/theme/pictures/free/original/theme-image-1601219133573.jpeg"
                      },
                      "slide_link": {
                        "type": "url",
                        "value": "https://www.turtleonline.in/collection/men-collection-vcahuo2q"
                      }
                    }
                  },
                  {
                    "type": "gallery_image",
                    "name": "Image",
                    "props": {
                      "image": {
                        "type": "image_picker",
                        "value": "https://hdn-1.fynd.com/company/163/applications/5d5265c8f8ad9cae6dbf18f1/theme/pictures/free/original/theme-image-1601219184350.jpeg"
                      },
                      "slide_link": {
                        "type": "url",
                        "value": "https://www.turtleonline.in/collection/shirts-5e9654ad"
                      }
                    }
                  }
                ],
                "predicate": {
                  "screen": {
                    "mobile": true,
                    "desktop": true,
                    "tablet": true
                  },
                  "user": {
                    "authenticated": true,
                    "anonymous": true
                  },
                  "route": {
                    "selected": "none",
                    "query": {},
                    "exact_u_r_l": ""
                  }
                },
                "name": "image-carousel",
                "props": {
                  "slide_height": {
                    "type": "select",
                    "value": "adapt"
                  },
                  "autoplay": {
                    "type": "checkbox",
                    "value": false
                  },
                  "slide_interval": {
                    "type": "range",
                    "value": 2
                  }
                }
              },
              {
                "blocks": [
                  {
                    "type": "collection",
                    "name": "Collection",
                    "props": {
                      "collection": {
                        "type": "collection",
                        "value": "accessories-7ee89654"
                      },
                      "title": {
                        "type": "text",
                        "value": "Turtle sports club"
                      },
                      "subtitle": {
                        "type": "text",
                        "value": "Casual Collections"
                      },
                      "overlay_image": {
                        "type": "image_picker",
                        "value": "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/pictures/free/original/casual_nfadbl.jpg"
                      }
                    }
                  },
                  {
                    "type": "collection",
                    "name": "Collection",
                    "props": {
                      "collection": {
                        "type": "collection",
                        "value": "bottom-wear-fb133293"
                      },
                      "title": {
                        "type": "text",
                        "value": "Formal Collection"
                      },
                      "subtitle": {
                        "type": "text",
                        "value": "Turtle Tailor Mark"
                      },
                      "overlay_image": {
                        "type": "image_picker",
                        "value": "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/pictures/free/original/tailors_nsrrfl.jpg"
                      }
                    }
                  },
                  {
                    "type": "collection",
                    "name": "Collection",
                    "props": {
                      "collection": {
                        "type": "collection",
                        "value": "bottom-wear-fb133293"
                      },
                      "title": {
                        "type": "text",
                        "value": "Bottomwear"
                      },
                      "subtitle": {
                        "type": "text",
                        "value": "Chinos | Trousers | Pants"
                      },
                      "overlay_image": {
                        "type": "image_picker",
                        "value": "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/pictures/free/original/bottomwear_tdkhn2.jpg"
                      }
                    }
                  },
                  {
                    "type": "collection",
                    "name": "Collection",
                    "props": {
                      "collection": {
                        "type": "collection",
                        "value": "t-shirt-7ee3cbcd"
                      },
                      "title": {
                        "type": "text",
                        "value": "Shirts"
                      },
                      "subtitle": {
                        "type": "text",
                        "value": "Casual"
                      },
                      "overlay_image": {
                        "type": "image_picker",
                        "value": "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/pictures/free/original/t-shirt_re9srk.jpg"
                      }
                    }
                  }
                ],
                "predicate": {
                  "screen": {
                    "mobile": true,
                    "desktop": true,
                    "tablet": true
                  },
                  "user": {
                    "authenticated": true,
                    "anonymous": true
                  },
                  "route": {
                    "selected": "none",
                    "query": {},
                    "exact_u_r_l": ""
                  }
                },
                "name": "collectionGrid",
                "props": {
                  "title": {
                    "type": "text",
                    "value": "Collections"
                  },
                  "subtitle": {
                    "type": "text",
                    "value": "Buy from our"
                  },
                  "full_width": {
                    "type": "checkbox",
                    "value": false
                  }
                }
              }
            ],
            "page_key": "home"
          },
          {
            "page_sections": [
              {
                "blocks": [],
                "predicate": {
                  "screen": {
                    "mobile": true,
                    "desktop": true,
                    "tablet": true
                  },
                  "user": {
                    "authenticated": true,
                    "anonymous": true
                  },
                  "route": {
                    "selected": "none",
                    "query": {},
                    "exact_u_r_l": ""
                  }
                },
                "name": "brandTemplate",
                "props": {
                  "item_count": {
                    "type": "range",
                    "value": 4
                  },
                  "full_width": {
                    "type": "checkbox",
                    "value": false
                  },
                  "heading": {
                    "type": "text",
                    "value": ""
                  },
                  "brand_type": {
                    "type": "radio",
                    "value": "all"
                  },
                  "layout": {
                    "type": "select",
                    "value": "grid"
                  },
                  "department": {
                    "type": "department",
                    "value": "others"
                  }
                }
              }
            ],
            "page_key": "brands"
          },
          {
            "page_sections": [],
            "page_key": "cart-landing"
          },
          {
            "page_sections": [
              {
                "blocks": [],
                "predicate": {
                  "screen": {
                    "mobile": true,
                    "desktop": true,
                    "tablet": true
                  },
                  "user": {
                    "authenticated": true,
                    "anonymous": true
                  },
                  "route": {
                    "selected": "none",
                    "query": {},
                    "exact_u_r_l": ""
                  }
                },
                "name": "categoriesTemplate",
                "props": {
                  "item_count": {
                    "type": "range",
                    "value": 4
                  },
                  "full_width": {
                    "type": "checkbox",
                    "value": false
                  },
                  "heading": {
                    "type": "text",
                    "value": ""
                  },
                  "layout": {
                    "type": "select",
                    "value": "grid"
                  },
                  "view_all": {
                    "type": "checkbox",
                    "value": false
                  }
                }
              }
            ],
            "page_key": "categories"
          },
          {
            "page_sections": [
              {
                "blocks": [],
                "predicate": {
                  "screen": {
                    "mobile": true,
                    "desktop": true,
                    "tablet": true
                  },
                  "user": {
                    "authenticated": true,
                    "anonymous": true
                  },
                  "route": {
                    "selected": "none",
                    "query": {},
                    "exact_u_r_l": ""
                  }
                },
                "name": "collectionTemplate",
                "props": {
                  "heading": {
                    "type": "text",
                    "value": "Featured Collections"
                  },
                  "item_count": {
                    "type": "range",
                    "value": 4
                  },
                  "full_width": {
                    "type": "checkbox",
                    "value": false
                  },
                  "layout": {
                    "type": "select",
                    "value": "grid"
                  }
                }
              }
            ],
            "page_key": "collections"
          },
          {
            "page_sections": [],
            "page_key": "compare-products"
          },
          {
            "page_sections": [],
            "page_key": "product-description"
          },
          {
            "page_sections": [],
            "page_key": "product-listing"
          },
          {
            "page_sections": [],
            "page_key": "wishlist"
          }
        ]
      },
      "global_schema": {
        "props": [
          {
            "type": "color",
            "id": "header_bg_color",
            "default": "#ffffff",
            "category": "Header",
            "label": "Header background Color"
          },
          {
            "type": "color",
            "id": "footer_bg_color",
            "default": "#1b1b1b",
            "category": "Footer",
            "label": "Footer background Color"
          },
          {
            "type": "text",
            "id": "footer_text",
            "default": "",
            "category": "Footer",
            "label": "Footer Text"
          },
          {
            "type": "checkbox",
            "id": "disable_cart",
            "default": false,
            "category": "Cart",
            "label": "Disable Cart"
          }
        ]
      },
      "current": "default",
      "page_schema": [
        {
          "props": [],
          "_id": "5fe182f763d26d042fd205c4",
          "page": "add-product-review"
        },
        {
          "props": [],
          "_id": "5fe182f763d26dadc8d205c6",
          "page": "blog"
        },
        {
          "props": [],
          "_id": "5fe182f763d26d0a36d205c7",
          "page": "brands"
        },
        {
          "props": [
            {
              "type": "checkbox",
              "id": "gst",
              "label": "GST",
              "default": true,
              "info": "Show GST on cart"
            },
            {
              "type": "checkbox",
              "id": "staff_selection",
              "label": "Staff Selection",
              "default": true,
              "info": "Show Staff selection on Cart"
            },
            {
              "type": "checkbox",
              "id": "enable_customer",
              "label": "Customer",
              "default": true,
              "info": "Placing on behalf of customer"
            },
            {
              "type": "checkbox",
              "id": "enable_guest",
              "label": "Enable Guest Checkout",
              "default": true,
              "info": "Enable Continue as Guest"
            }
          ],
          "_id": "5fe182f763d26d81c5d205c8",
          "page": "cart-landing"
        },
        {
          "props": [],
          "_id": "5fe182f763d26d7e15d205c9",
          "page": "cart-review"
        },
        {
          "props": [],
          "_id": "5fe182f763d26d270ed205ca",
          "page": "categories"
        },
        {
          "props": [],
          "_id": "5fe182f763d26d9b4fd205cb",
          "page": "collection-listing"
        },
        {
          "props": [],
          "_id": "5fe182f763d26da6ecd205cc",
          "page": "collections"
        },
        {
          "props": [],
          "_id": "5fe182f763d26d7156d205cd",
          "page": "compare-products"
        },
        {
          "props": [],
          "_id": "5fe182f763d26d62bad205ce",
          "page": "home"
        },
        {
          "props": [],
          "_id": "5fe182f763d26d5afcd205cf",
          "page": "order-review"
        },
        {
          "props": [],
          "_id": "5fe182f763d26def8dd205d0",
          "page": "order-tracking-details"
        },
        {
          "props": [],
          "_id": "5fe182f763d26d381fd205d1",
          "page": "order-tracking"
        },
        {
          "props": [
            {
              "type": "text",
              "id": "shipping_link",
              "label": "Shipping Link",
              "default": "https://fynd.freshdesk.com/support/solutions/folders/33000111600",
              "info": "Link of shipping page"
            },
            {
              "type": "checkbox",
              "id": "original_image",
              "label": "Use original Image",
              "default": false,
              "info": "Use original product image."
            },
            {
              "type": "checkbox",
              "id": "wishlist",
              "label": "Wishlist",
              "default": true,
              "info": "Show Wishlist for product"
            },
            {
              "type": "checkbox",
              "id": "reviews",
              "label": "Review",
              "default": true,
              "info": "Show Reviews of product"
            },
            {
              "type": "checkbox",
              "id": "add_to_compare",
              "label": "Add to Compare",
              "default": true,
              "info": "Allow comparison of products"
            },
            {
              "type": "checkbox",
              "id": "size_guide",
              "label": "Size Guide",
              "default": true,
              "info": "Show Size Guide"
            },
            {
              "type": "checkbox",
              "id": "product_request",
              "label": "Product Request",
              "default": true,
              "info": "Show Product Request"
            },
            {
              "type": "checkbox",
              "id": "share",
              "label": "Share",
              "default": true,
              "info": "Enable Sharing product"
            },
            {
              "type": "checkbox",
              "id": "sold_by",
              "label": "Show Sold By",
              "default": true,
              "info": "Show name of the store"
            },
            {
              "type": "checkbox",
              "id": "store_selection",
              "label": "Seller Store Selection",
              "default": true,
              "info": "Allow to explicitly select stores"
            },
            {
              "type": "checkbox",
              "id": "compare_products",
              "label": "Compare Products",
              "default": true,
              "info": "Show Most Compared Products"
            },
            {
              "type": "checkbox",
              "id": "variants",
              "label": "Product Variants",
              "default": true,
              "info": "Show Product Variants"
            },
            {
              "type": "checkbox",
              "id": "ratings",
              "label": "Product Rating",
              "default": true,
              "info": "Show Product Ratings"
            },
            {
              "type": "checkbox",
              "id": "similar_products",
              "label": "Similar Products",
              "default": true,
              "info": "Show Similar Products"
            },
            {
              "type": "checkbox",
              "id": "bulk_prices",
              "label": "Bulk Prices",
              "default": true,
              "info": "Show Bulk Prices"
            },
            {
              "type": "checkbox",
              "id": "showDeliveryInfo",
              "label": "Delivery Info",
              "default": true,
              "info": "Show Delivery Date"
            }
          ],
          "_id": "5fe182f763d26d29bbd205d2",
          "page": "product-description"
        },
        {
          "props": [],
          "_id": "5fe182f763d26da5f0d205d3",
          "page": "product-listing"
        },
        {
          "props": [],
          "_id": "5fe182f763d26d3d18d205d4",
          "page": "product-reviews"
        },
        {
          "props": [],
          "_id": "5fe182f763d26d4e2dd205d5",
          "page": "wishlist"
        },
        {
          "props": [
            {
              "default": true,
              "id": "header",
              "label": "Header",
              "type": "checkbox"
            },
            {
              "default": true,
              "id": "footer",
              "label": "Footer",
              "type": "checkbox"
            }
          ],
          "_id": "5fe7166cbaae343115de8555",
          "page": "tesr"
        }
      ],
      "list": [
        {
          "name": "default",
          "global_config": {
            "static": {
              "props": {
                "colors": {
                  "primary_color": "#7043f7",
                  "secondary_color": "#02d1cb",
                  "accent_color": "#FFFFFF",
                  "link_color": "#7043f7",
                  "button_secondary_color": "#000000",
                  "bg_color": "#F8F8F8"
                }
              }
            },
            "custom": {
              "props": {
                "disable_cart": false
              }
            }
          },
          "page": [
            {
              "settings": {
                "props": {
                  "show_info_message": true
                }
              },
              "page": "cart-landing"
            }
          ],
          "_id": "5fc4bb9078e957737e7d79a3"
        }
      ]
    },
    "colors": {
      "primary_color": "#41434C",
      "secondary_color": "#41434C",
      "accent_color": "#FFFFFF",
      "link_color": "#33B1C0",
      "button_secondary_color": "#000000",
      "bg_color": "#F8F8F8"
    },
    "font": {
      "variants": {
        "light": {
          "name": "300",
          "file": ""
        },
        "regular": {
          "name": "regular",
          "file": "https://fonts.gstatic.com/s/judson/v13/FeVRS0Fbvbc14VxRD7N01bV7kg.ttf"
        },
        "medium": {
          "name": "500",
          "file": ""
        },
        "semi_bold": {
          "name": "600",
          "file": ""
        },
        "bold": {
          "name": "700",
          "file": "https://fonts.gstatic.com/s/judson/v13/FeVSS0Fbvbc14Vxps5xQ3Z5nm29Gww.ttf"
        }
      },
      "family": "Judson"
    },
    "applied": true,
    "published": false,
    "archived": false,
    "customized": true,
    "version": "1.1.19",
    "tags": [],
    "_id": "5fe17f7063d26dc54fd202b4",
    "pages": {
      "home": {
        "path": "",
        "type": "system",
        "sections": [],
        "value": "home",
        "text": "Home"
      },
      "brands": {
        "path": "brands",
        "type": "system",
        "sections": [],
        "value": "brands",
        "text": "Brands"
      },
      "cart-landing": {
        "path": "cart/bag",
        "type": "system",
        "sections": [],
        "value": "cart-landing",
        "text": "Cart Landing"
      },
      "categories": {
        "path": "categories",
        "type": "system",
        "sections": [],
        "value": "categories",
        "text": "Categories"
      },
      "collections": {
        "path": "collections",
        "type": "system",
        "sections": [],
        "value": "collections",
        "text": "Collections"
      },
      "compare-products": {
        "path": "compare",
        "type": "system",
        "sections": [],
        "value": "compare-products",
        "text": "Compare Products"
      },
      "product-description": {
        "path": "product",
        "type": "system",
        "sections": [],
        "value": "product-description",
        "text": "Product Description"
      },
      "product-listing": {
        "path": "products",
        "type": "system",
        "sections": [],
        "value": "product-listing",
        "text": "Product Listing"
      },
      "collection-listing": {
        "path": "collection",
        "type": "system",
        "sections": [],
        "value": "collection-listing",
        "text": "COllection Listing"
      },
      "wishlist": {
        "path": "wishlist",
        "type": "system",
        "sections": [],
        "value": "wishlist",
        "text": "Wishlist"
      }
    },
    "available_sections": [
      {
        "props": [
          {
            "type": "text",
            "id": "heading",
            "default": "Brands",
            "label": "Brands Heading"
          },
          {
            "type": "radio",
            "id": "brand_type",
            "default": "all",
            "options": [
              {
                "value": "all",
                "text": "All"
              },
              {
                "value": "department",
                "text": "Department"
              },
              {
                "value": "handpicked",
                "text": "Handpicked"
              }
            ]
          },
          {
            "type": "department",
            "id": "department",
            "label": "Department",
            "info": "Select a department of brands",
            "note": "Department only applies if 'department' type is selected"
          },
          {
            "type": "range",
            "id": "item_count",
            "min": 3,
            "max": 5,
            "step": 1,
            "unit": "",
            "label": "Items per row",
            "default": 4,
            "info": "Maximum items allowed per row"
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          },
          {
            "id": "layout",
            "type": "select",
            "options": [
              {
                "value": "grid",
                "text": "Grid View"
              },
              {
                "value": "horizontal",
                "text": "Horizontal View"
              }
            ],
            "default": "grid",
            "label": "Layout",
            "info": "Alignment of content"
          },
          {
            "type": "checkbox",
            "id": "view_all",
            "default": false,
            "label": "Show View All"
          }
        ],
        "blocks": [
          {
            "type": "brand-item",
            "name": "Brand Item",
            "props": [
              {
                "type": "brand",
                "id": "brand",
                "label": "Select Brand"
              }
            ]
          }
        ],
        "_id": "5feacca5bec232d59b89283a",
        "name": "brandTemplate",
        "label": "Brands List Page"
      },
      {
        "props": [
          {
            "type": "text",
            "id": "heading",
            "default": "Categories",
            "label": "Categories Heading"
          },
          {
            "id": "layout",
            "type": "select",
            "options": [
              {
                "value": "grid",
                "text": "Grid View"
              },
              {
                "value": "horizontal",
                "text": "Horizontal View"
              }
            ],
            "default": "grid",
            "label": "Layout",
            "info": "Alignment of content"
          },
          {
            "type": "range",
            "id": "item_count",
            "min": 3,
            "max": 5,
            "step": 1,
            "unit": "",
            "label": "Items per row",
            "default": 4,
            "info": "Maximum items allowed per row"
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          },
          {
            "type": "checkbox",
            "id": "view_all",
            "default": false,
            "label": "Show View All",
            "info": "Check to show View All Button"
          }
        ],
        "blocks": [],
        "_id": "5feacca5bec232404189283b",
        "name": "categoriesTemplate",
        "label": "Categories Page"
      },
      {
        "props": [
          {
            "type": "text",
            "id": "title",
            "default": "",
            "label": "Title"
          },
          {
            "type": "text",
            "id": "subtitle",
            "default": "",
            "label": "Subtitle"
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          }
        ],
        "blocks": [
          {
            "type": "collection",
            "name": "Collection",
            "props": [
              {
                "type": "collection",
                "id": "collection",
                "label": "Select a collection"
              },
              {
                "type": "image_picker",
                "id": "overlayImage",
                "default": "",
                "label": "Background image",
                "info": "Background Image"
              },
              {
                "type": "text",
                "id": "title",
                "default": "",
                "label": "Overlay title"
              },
              {
                "type": "text",
                "id": "subtitle",
                "default": "",
                "label": "Overlay subtitle"
              }
            ]
          }
        ],
        "_id": "5feacca5bec2321fd589283c",
        "name": "collectionGrid",
        "label": "Collection Grid"
      },
      {
        "props": [
          {
            "type": "text",
            "id": "heading",
            "default": "Collection List",
            "label": "Collection Heading",
            "info": "Collection Heading"
          },
          {
            "type": "range",
            "id": "item_count",
            "min": 3,
            "max": 5,
            "step": 1,
            "unit": "",
            "label": "Collections per row",
            "default": 2,
            "info": "Maximum items allowed per row"
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          }
        ],
        "blocks": [
          {
            "type": "collection",
            "name": "Collection",
            "props": [
              {
                "type": "collection",
                "id": "collection",
                "label": "Select a collection"
              }
            ]
          }
        ],
        "_id": "5feacca5bec2323bf689283d",
        "name": "collectionList",
        "label": "Collection List"
      },
      {
        "props": [
          {
            "type": "text",
            "id": "heading",
            "default": "Collections",
            "label": "Collection Heading"
          },
          {
            "type": "range",
            "id": "item_count",
            "min": 3,
            "max": 5,
            "step": 1,
            "unit": "",
            "label": "Collections per row",
            "default": 4,
            "info": "Maximum items allowed per row"
          },
          {
            "type": "radio",
            "id": "collection_type",
            "default": "all",
            "options": [
              {
                "value": "all",
                "text": "All"
              },
              {
                "value": "handpicked",
                "text": "Handpicked"
              }
            ]
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          },
          {
            "id": "layout",
            "type": "select",
            "options": [
              {
                "value": "grid",
                "text": "Grid View"
              },
              {
                "value": "horizontal",
                "text": "Horizontal View"
              }
            ],
            "default": "grid",
            "label": "Layout",
            "info": "Alignment of content"
          },
          {
            "type": "checkbox",
            "id": "view_all",
            "default": false,
            "label": "Show View All"
          }
        ],
        "blocks": [
          {
            "type": "collection-item",
            "name": "Collection Item",
            "props": [
              {
                "type": "collection",
                "id": "collection",
                "label": "Select collection"
              }
            ]
          }
        ],
        "_id": "5feacca5bec23263b489283e",
        "name": "collectionTemplate",
        "label": "Collection List Page"
      },
      {
        "props": [
          {
            "type": "code",
            "id": "code",
            "label": "Custom HTML",
            "info": "Add Your custom HTML Code below. You can also use the full screen icon to open a code editor and add your code"
          }
        ],
        "blocks": [],
        "_id": "5feacca5bec232409489283f",
        "name": "customHtml",
        "label": "Custom HTML"
      },
      {
        "props": [
          {
            "type": "text",
            "id": "heading",
            "default": "Featured Products",
            "label": "Section Heading"
          },
          {
            "type": "range",
            "id": "item_count",
            "min": 3,
            "max": 5,
            "step": 1,
            "unit": "",
            "label": "Products per row",
            "default": 4,
            "info": "Maximum items allowed per row"
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          }
        ],
        "blocks": [
          {
            "type": "product",
            "name": "Product",
            "props": [
              {
                "type": "product",
                "id": "product",
                "label": "Select a Product",
                "info": "Product Item to be displayed"
              }
            ]
          }
        ],
        "_id": "5feacca5bec2326213892840",
        "name": "featuredProducts",
        "label": "Featured Products",
        "preset": {
          "blocks": [
            {
              "name": "Product"
            },
            {
              "name": "Product"
            },
            {
              "name": "Product"
            },
            {
              "name": "Product"
            }
          ]
        }
      },
      {
        "props": [
          {
            "type": "range",
            "id": "item_count",
            "min": 1,
            "max": 4,
            "step": 1,
            "unit": "",
            "label": "No of items",
            "default": 4,
            "info": "Maximum items allowed per row"
          }
        ],
        "blocks": [
          {
            "type": "gallery_image",
            "name": "Image",
            "props": [
              {
                "type": "image_picker",
                "id": "image",
                "label": "Gallery Image",
                "default": "https://hdn-1.fynd.com/company/163/applications/5d5265c8f8ad9cae6dbf18f1/theme/pictures/free/original/theme-image-1603773049684.svg"
              },
              {
                "type": "text",
                "id": "caption",
                "label": "Image Caption",
                "default": ""
              },
              {
                "type": "url",
                "id": "link",
                "label": "Link",
                "default": "",
                "info": "Link to redirect"
              }
            ]
          }
        ],
        "_id": "5feacca5bec2321047892841",
        "name": "gallery",
        "label": "Gallery",
        "preset": {
          "props": {
            "item_count": 4
          },
          "blocks": [
            {
              "name": "Image"
            },
            {
              "name": "Image"
            },
            {
              "name": "Image"
            },
            {
              "name": "Image"
            }
          ]
        }
      },
      {
        "props": [
          {
            "type": "url",
            "id": "ctaLink",
            "default": "",
            "label": "Redirect Link"
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          }
        ],
        "blocks": [],
        "_id": "5feacca5bec232b085892842",
        "name": "heroBanner",
        "label": "Hero Banner"
      },
      {
        "props": [
          {
            "id": "image",
            "type": "image_picker",
            "label": "Hero Image",
            "default": ""
          },
          {
            "type": "text",
            "id": "heading",
            "default": "",
            "label": "Heading"
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          },
          {
            "id": "overlayLayout",
            "type": "select",
            "options": [
              {
                "value": "left",
                "text": "Align Left"
              },
              {
                "value": "center",
                "text": "Align Center"
              },
              {
                "value": "right",
                "text": "Align Right"
              }
            ],
            "default": "left",
            "label": "Overlay Layout",
            "info": "Alignment of overlay content"
          },
          {
            "type": "image_picker",
            "id": "overlayImage",
            "default": "",
            "label": "Overlay image",
            "info": "Overlay Image"
          },
          {
            "type": "text",
            "id": "text",
            "default": "",
            "label": "Overlay Text"
          },
          {
            "type": "color",
            "id": "text_color",
            "default": "#000",
            "label": "Text Color"
          },
          {
            "type": "url",
            "id": "ctaLink",
            "default": "",
            "label": "Redirect Link"
          },
          {
            "type": "text",
            "id": "ctaText",
            "default": "Shop Now",
            "label": "Button Text"
          }
        ],
        "blocks": [],
        "_id": "5feacca5bec2321e74892843",
        "name": "imageBanner",
        "label": "Hero Image"
      },
      {
        "props": [
          {
            "id": "videoUrl",
            "type": "url",
            "label": "Video URL",
            "default": ""
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          },
          {
            "id": "coverUrl",
            "type": "image_picker",
            "label": "Video Cover Image URL",
            "default": ""
          },
          {
            "type": "checkbox",
            "id": "showcontrols",
            "default": false,
            "label": "Show Controls on Video",
            "info": "Check to show controls on video"
          },
          {
            "type": "select",
            "id": "size",
            "options": [
              {
                "value": "adapt",
                "text": "Adapt"
              },
              {
                "value": "small",
                "text": "Small"
              },
              {
                "value": "medium",
                "text": "Medium"
              },
              {
                "value": "large",
                "text": "Large"
              }
            ],
            "default": "adapt",
            "label": "Video Height",
            "info": "Height of Video"
          },
          {
            "type": "text",
            "id": "heading",
            "default": "",
            "label": "Heading"
          },
          {
            "type": "color",
            "id": "heading_color",
            "default": "#000",
            "label": "Headin Text Color"
          },
          {
            "type": "text",
            "id": "subHeading",
            "default": "",
            "label": "Sub-heading"
          },
          {
            "type": "color",
            "id": "subheading_color",
            "default": "#000",
            "label": "Subheading Text Color"
          },
          {
            "type": "url",
            "id": "ctaLink",
            "default": "",
            "label": "Redirect Link"
          },
          {
            "type": "text",
            "id": "ctaText",
            "default": "ShopNow",
            "label": "Button Text"
          }
        ],
        "blocks": [],
        "_id": "5feacca5bec232bfc8892844",
        "name": "videoBanner",
        "label": "Hero Video"
      },
      {
        "props": [
          {
            "type": "select",
            "id": "slide_height",
            "options": [
              {
                "value": "adapt",
                "text": "Adapt to first image"
              },
              {
                "value": "small",
                "text": "Small"
              },
              {
                "value": "medium",
                "text": "Medium"
              },
              {
                "value": "large",
                "text": "Large"
              }
            ],
            "default": "adapt",
            "label": "Slide height",
            "info": "Size of the slide"
          },
          {
            "type": "checkbox",
            "id": "autoplay",
            "default": false,
            "label": "AutoPlay Slides",
            "info": "Check to autoplay slides"
          },
          {
            "type": "range",
            "id": "slide_interval",
            "min": 1,
            "max": 10,
            "step": 1,
            "unit": "sec",
            "label": "Change slides after every",
            "default": 2,
            "info": "Autoplay slide duration"
          }
        ],
        "blocks": [
          {
            "type": "gallery_image",
            "name": "Image",
            "props": [
              {
                "type": "image_picker",
                "id": "image",
                "label": "Gallery Image"
              },
              {
                "type": "url",
                "id": "slide_link",
                "label": "Slide Link"
              }
            ]
          }
        ],
        "_id": "5feacca5bec232a916892845",
        "name": "image-carousel",
        "label": "Image Carousel",
        "preset": {
          "blocks": [
            {
              "name": "Image"
            },
            {
              "name": "Image"
            }
          ]
        }
      },
      {
        "props": [
          {
            "type": "text",
            "id": "heading",
            "default": "Featured Collections",
            "label": "Collection Heading"
          },
          {
            "type": "range",
            "id": "item_count",
            "min": 2,
            "max": 5,
            "step": 1,
            "unit": "",
            "label": "Products per row",
            "default": 2,
            "info": "Maximum items allowed per row"
          },
          {
            "type": "collection",
            "id": "collection",
            "label": "Collection",
            "info": "Select a collection to display its products"
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          }
        ],
        "blocks": [],
        "_id": "5feacca5bec23258ec892846",
        "name": "productList",
        "label": "Product List",
        "preset": {
          "props": {
            "heading": "Featured Products",
            "item_count": 4,
            "collection": ""
          }
        }
      },
      {
        "props": [
          {
            "type": "checkbox",
            "id": "autoplay",
            "default": false,
            "label": "AutoPlay Slides"
          },
          {
            "type": "range",
            "id": "slide_interval",
            "min": 1,
            "max": 10,
            "step": 1,
            "unit": "sec",
            "label": "Change slides every",
            "default": 2
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          }
        ],
        "blocks": [
          {
            "type": "testimonial",
            "name": "Testimonial",
            "props": [
              {
                "type": "textarea",
                "id": "testimonialText",
                "label": "Text for Testimonial",
                "default": "",
                "info": "Text for testimonial",
                "placeholder": "Text"
              },
              {
                "type": "text",
                "id": "author",
                "label": "Customers name"
              }
            ]
          }
        ],
        "_id": "5feacca5bec23299e8892847",
        "name": "testimonials",
        "label": "Testimonial"
      },
      {
        "props": [
          {
            "type": "select",
            "id": "slide_height",
            "options": [
              {
                "value": "adapt",
                "text": "Adapt to first video"
              },
              {
                "value": "small",
                "text": "Small"
              },
              {
                "value": "medium",
                "text": "Medium"
              },
              {
                "value": "large",
                "text": "Large"
              }
            ],
            "default": "adapt",
            "label": "Slide height",
            "info": "Size of the slide"
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          },
          {
            "type": "checkbox",
            "id": "autoplay",
            "default": false,
            "label": "AutoPlay Slides",
            "info": "Check to autoplay slides"
          },
          {
            "type": "range",
            "id": "slide_interval",
            "min": 1,
            "max": 10,
            "step": 1,
            "unit": "sec",
            "label": "Change slides every",
            "default": 2,
            "info": "Autoplay slide duration"
          }
        ],
        "blocks": [
          {
            "type": "video_item",
            "name": "Video Slide",
            "props": [
              {
                "id": "videoUrl",
                "type": "url",
                "label": "Video URL",
                "default": ""
              },
              {
                "type": "checkbox",
                "id": "showcontrols",
                "default": false,
                "label": "Show Controls on Video",
                "info": "Check to show controls on video"
              }
            ]
          }
        ],
        "_id": "5feacca5bec232d89b892848",
        "name": "videoCarousel",
        "label": "Video Carousel"
      }
    ],
    "sections": [
      {
        "page_sections": [
          {
            "blocks": [],
            "predicate": {
              "screen": {
                "mobile": true,
                "desktop": true,
                "tablet": true
              },
              "user": {
                "authenticated": true,
                "anonymous": true
              },
              "route": {
                "selected": "none",
                "query": {},
                "exact_u_r_l": ""
              }
            },
            "name": "brandTemplate",
            "props": {
              "title": {
                "type": "text",
                "value": "Brands"
              },
              "item_count": {
                "type": "range",
                "value": 4
              },
              "full_width": {
                "type": "checkbox",
                "value": false
              },
              "heading": {
                "type": "text",
                "value": "Brands"
              },
              "brand_type": {
                "type": "radio",
                "value": "all"
              },
              "department": {
                "type": "department",
                "value": null
              },
              "layout": {
                "type": "select",
                "value": "grid"
              },
              "view_all": {
                "type": "checkbox",
                "value": null
              }
            }
          }
        ],
        "page_key": "product-description"
      },
      {
        "page_sections": [],
        "page_key": "collection-listing"
      },
      {
        "page_sections": [
          {
            "blocks": [
              {
                "type": "brand-item",
                "name": "Brand Item",
                "props": {
                  "brand": {
                    "type": "brand",
                    "value": {
                      "display": "Apple",
                      "id": "apple"
                    }
                  }
                }
              },
              {
                "type": "brand-item",
                "name": "Brand Item",
                "props": {
                  "brand": {
                    "type": "brand",
                    "value": {
                      "display": "Gionee",
                      "id": "gionee"
                    }
                  }
                }
              }
            ],
            "predicate": {
              "screen": {
                "mobile": true,
                "desktop": true,
                "tablet": true
              },
              "user": {
                "authenticated": true,
                "anonymous": true
              },
              "route": {
                "selected": "none",
                "query": {},
                "exact_u_r_l": ""
              }
            },
            "name": "brandTemplate",
            "props": {
              "heading": {
                "type": "text",
                "value": "Brandss"
              },
              "brand_type": {
                "type": "radio",
                "value": "handpicked"
              },
              "department": {
                "type": "department"
              },
              "item_count": {
                "value": 4,
                "type": "range"
              },
              "full_width": {
                "value": false,
                "type": "checkbox"
              },
              "layout": {
                "type": "select",
                "value": "grid"
              },
              "view_all": {
                "value": false,
                "type": "checkbox"
              }
            }
          }
        ],
        "page_key": "home"
      },
      {
        "page_sections": [
          {
            "blocks": [],
            "predicate": {
              "screen": {
                "mobile": true,
                "desktop": true,
                "tablet": true
              },
              "user": {
                "authenticated": true,
                "anonymous": true
              },
              "route": {
                "selected": "none",
                "query": {},
                "exact_u_r_l": ""
              }
            },
            "name": "brandTemplate",
            "props": {
              "title": {
                "type": "text",
                "value": "Brands"
              },
              "item_count": {
                "type": "range",
                "value": 4
              },
              "full_width": {
                "type": "checkbox",
                "value": false
              }
            }
          }
        ],
        "page_key": "brands"
      },
      {
        "page_sections": [],
        "page_key": "cart-landing"
      },
      {
        "page_sections": [
          {
            "blocks": [],
            "predicate": {
              "screen": {
                "mobile": true,
                "desktop": true,
                "tablet": true
              },
              "user": {
                "authenticated": true,
                "anonymous": true
              },
              "route": {
                "selected": "none",
                "query": {},
                "exact_u_r_l": ""
              }
            },
            "name": "categoriesTemplate",
            "props": {
              "item_count": {
                "type": "range",
                "value": 4
              },
              "full_width": {
                "type": "checkbox",
                "value": false
              }
            }
          }
        ],
        "page_key": "categories"
      },
      {
        "page_sections": [
          {
            "blocks": [],
            "predicate": {
              "screen": {
                "mobile": true,
                "desktop": true,
                "tablet": true
              },
              "user": {
                "authenticated": true,
                "anonymous": true
              },
              "route": {
                "selected": "none",
                "query": {},
                "exact_u_r_l": ""
              }
            },
            "name": "collectionTemplate",
            "props": {
              "heading": {
                "type": "text",
                "value": "Featured Collections"
              },
              "item_count": {
                "type": "range",
                "value": 4
              },
              "full_width": {
                "type": "checkbox",
                "value": false
              }
            }
          }
        ],
        "page_key": "collections"
      },
      {
        "page_sections": [],
        "page_key": "compare-products"
      },
      {
        "page_sections": [],
        "page_key": "product-listing"
      },
      {
        "page_sections": [],
        "page_key": "wishlist"
      },
      {
        "page_sections": [],
        "page_key": "tesr"
      }
    ],
    "application": "5e737afb97e0f586bf9d04db",
    "available_pages": [
      {
        "path": "product",
        "type": "system",
        "sections": [],
        "text": "Product Description",
        "value": "product-description"
      },
      {
        "path": "collection",
        "type": "system",
        "sections": [],
        "text": "Collection Listing",
        "value": "collection-listing"
      },
      {
        "path": "",
        "type": "system",
        "sections": [
          {
            "_id": "5feacca5bec2327927892853",
            "attributes": {
              "page": "home"
            }
          }
        ],
        "value": "home",
        "text": "Home"
      },
      {
        "path": "brands",
        "type": "system",
        "sections": [
          {
            "_id": "5feacca5bec2326422892854",
            "attributes": {
              "page": "brands"
            }
          }
        ],
        "value": "brands",
        "text": "Brands"
      },
      {
        "path": "cart/bag",
        "type": "system",
        "sections": [],
        "value": "cart-landing",
        "text": "Cart Landing"
      },
      {
        "path": "categories",
        "type": "system",
        "sections": [
          {
            "_id": "5feacca5bec232424c892855",
            "attributes": {
              "page": "categories"
            }
          }
        ],
        "value": "categories",
        "text": "Categories"
      },
      {
        "path": "collections",
        "type": "system",
        "sections": [
          {
            "_id": "5feacca5bec23281de892856",
            "attributes": {
              "page": "collections"
            }
          }
        ],
        "value": "collections",
        "text": "Collections"
      },
      {
        "path": "compare",
        "type": "system",
        "sections": [],
        "value": "compare-products",
        "text": "Compare Products"
      },
      {
        "path": "products",
        "type": "system",
        "sections": [],
        "value": "product-listing",
        "text": "Product Listing"
      },
      {
        "path": "wishlist",
        "type": "system",
        "sections": [],
        "value": "wishlist",
        "text": "Wishlist"
      },
      {
        "path": "sections/test",
        "type": "sections",
        "sections": [],
        "text": "test",
        "value": "test"
      }
    ],
    "styles": {},
    "created_at": "2020-12-22T05:09:04.720Z",
    "updated_at": "2021-01-24T11:22:41.376Z"
  }
}
```









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
  "value": {
    "information": {
      "images": {
        "desktop": [
          "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/pictures/free/original/artisan-desktop.png"
        ],
        "android": [
          "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/pictures/free/original/artisan-mobile.png"
        ],
        "ios": [
          "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/pictures/free/original/artisan-mobile.png"
        ],
        "thumbnail": [
          "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/pictures/free/original/artisan-desktop.png"
        ]
      },
      "features": [
        "Responsive"
      ],
      "name": "Akash-Artisan"
    },
    "src": {
      "link": "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/sources/J-1s-N-pl-archive.zip"
    },
    "assets": {
      "css": {
        "link": "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/assets/Ia4m885Mw2-themeBundle.css"
      },
      "umd_js": {
        "link": "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/assets/nEPGyc15g-themeBundle.umd.min.js"
      },
      "common_js": {
        "link": "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/assets/5f7tOB5rpq-themeBundle.common.js"
      }
    },
    "config": {
      "preset": {
        "sections": [
          {
            "page_sections": [
              {
                "blocks": [
                  {
                    "type": "gallery_image",
                    "name": "Image",
                    "props": {
                      "image": {
                        "type": "image_picker",
                        "value": "https://hdn-1.fynd.com/company/163/applications/5d5265c8f8ad9cae6dbf18f1/theme/pictures/free/original/theme-image-1601219133573.jpeg"
                      },
                      "slide_link": {
                        "type": "url",
                        "value": "https://www.turtleonline.in/collection/men-collection-vcahuo2q"
                      }
                    }
                  },
                  {
                    "type": "gallery_image",
                    "name": "Image",
                    "props": {
                      "image": {
                        "type": "image_picker",
                        "value": "https://hdn-1.fynd.com/company/163/applications/5d5265c8f8ad9cae6dbf18f1/theme/pictures/free/original/theme-image-1601219184350.jpeg"
                      },
                      "slide_link": {
                        "type": "url",
                        "value": "https://www.turtleonline.in/collection/shirts-5e9654ad"
                      }
                    }
                  }
                ],
                "predicate": {
                  "screen": {
                    "mobile": true,
                    "desktop": true,
                    "tablet": true
                  },
                  "user": {
                    "authenticated": true,
                    "anonymous": true
                  },
                  "route": {
                    "selected": "none",
                    "query": {},
                    "exact_u_r_l": ""
                  }
                },
                "name": "image-carousel",
                "props": {
                  "slide_height": {
                    "type": "select",
                    "value": "adapt"
                  },
                  "autoplay": {
                    "type": "checkbox",
                    "value": false
                  },
                  "slide_interval": {
                    "type": "range",
                    "value": 2
                  }
                }
              },
              {
                "blocks": [
                  {
                    "type": "collection",
                    "name": "Collection",
                    "props": {
                      "collection": {
                        "type": "collection",
                        "value": "accessories-7ee89654"
                      },
                      "title": {
                        "type": "text",
                        "value": "Turtle sports club"
                      },
                      "subtitle": {
                        "type": "text",
                        "value": "Casual Collections"
                      },
                      "overlay_image": {
                        "type": "image_picker",
                        "value": "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/pictures/free/original/casual_nfadbl.jpg"
                      }
                    }
                  },
                  {
                    "type": "collection",
                    "name": "Collection",
                    "props": {
                      "collection": {
                        "type": "collection",
                        "value": "bottom-wear-fb133293"
                      },
                      "title": {
                        "type": "text",
                        "value": "Formal Collection"
                      },
                      "subtitle": {
                        "type": "text",
                        "value": "Turtle Tailor Mark"
                      },
                      "overlay_image": {
                        "type": "image_picker",
                        "value": "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/pictures/free/original/tailors_nsrrfl.jpg"
                      }
                    }
                  },
                  {
                    "type": "collection",
                    "name": "Collection",
                    "props": {
                      "collection": {
                        "type": "collection",
                        "value": "bottom-wear-fb133293"
                      },
                      "title": {
                        "type": "text",
                        "value": "Bottomwear"
                      },
                      "subtitle": {
                        "type": "text",
                        "value": "Chinos | Trousers | Pants"
                      },
                      "overlay_image": {
                        "type": "image_picker",
                        "value": "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/pictures/free/original/bottomwear_tdkhn2.jpg"
                      }
                    }
                  },
                  {
                    "type": "collection",
                    "name": "Collection",
                    "props": {
                      "collection": {
                        "type": "collection",
                        "value": "t-shirt-7ee3cbcd"
                      },
                      "title": {
                        "type": "text",
                        "value": "Shirts"
                      },
                      "subtitle": {
                        "type": "text",
                        "value": "Casual"
                      },
                      "overlay_image": {
                        "type": "image_picker",
                        "value": "https://hdn-1.addsale.com/x0/company/1/applications/5e737afb97e0f586bf9d04db/theme/pictures/free/original/t-shirt_re9srk.jpg"
                      }
                    }
                  }
                ],
                "predicate": {
                  "screen": {
                    "mobile": true,
                    "desktop": true,
                    "tablet": true
                  },
                  "user": {
                    "authenticated": true,
                    "anonymous": true
                  },
                  "route": {
                    "selected": "none",
                    "query": {},
                    "exact_u_r_l": ""
                  }
                },
                "name": "collectionGrid",
                "props": {
                  "title": {
                    "type": "text",
                    "value": "Collections"
                  },
                  "subtitle": {
                    "type": "text",
                    "value": "Buy from our"
                  },
                  "full_width": {
                    "type": "checkbox",
                    "value": false
                  }
                }
              }
            ],
            "page_key": "home"
          },
          {
            "page_sections": [
              {
                "blocks": [],
                "predicate": {
                  "screen": {
                    "mobile": true,
                    "desktop": true,
                    "tablet": true
                  },
                  "user": {
                    "authenticated": true,
                    "anonymous": true
                  },
                  "route": {
                    "selected": "none",
                    "query": {},
                    "exact_u_r_l": ""
                  }
                },
                "name": "brandTemplate",
                "props": {
                  "item_count": {
                    "type": "range",
                    "value": 4
                  },
                  "full_width": {
                    "type": "checkbox",
                    "value": false
                  },
                  "heading": {
                    "type": "text",
                    "value": ""
                  },
                  "brand_type": {
                    "type": "radio",
                    "value": "all"
                  },
                  "layout": {
                    "type": "select",
                    "value": "grid"
                  },
                  "department": {
                    "type": "department",
                    "value": "others"
                  }
                }
              }
            ],
            "page_key": "brands"
          },
          {
            "page_sections": [],
            "page_key": "cart-landing"
          },
          {
            "page_sections": [
              {
                "blocks": [],
                "predicate": {
                  "screen": {
                    "mobile": true,
                    "desktop": true,
                    "tablet": true
                  },
                  "user": {
                    "authenticated": true,
                    "anonymous": true
                  },
                  "route": {
                    "selected": "none",
                    "query": {},
                    "exact_u_r_l": ""
                  }
                },
                "name": "categoriesTemplate",
                "props": {
                  "item_count": {
                    "type": "range",
                    "value": 4
                  },
                  "full_width": {
                    "type": "checkbox",
                    "value": false
                  },
                  "heading": {
                    "type": "text",
                    "value": ""
                  },
                  "layout": {
                    "type": "select",
                    "value": "grid"
                  },
                  "view_all": {
                    "type": "checkbox",
                    "value": false
                  }
                }
              }
            ],
            "page_key": "categories"
          },
          {
            "page_sections": [
              {
                "blocks": [],
                "predicate": {
                  "screen": {
                    "mobile": true,
                    "desktop": true,
                    "tablet": true
                  },
                  "user": {
                    "authenticated": true,
                    "anonymous": true
                  },
                  "route": {
                    "selected": "none",
                    "query": {},
                    "exact_u_r_l": ""
                  }
                },
                "name": "collectionTemplate",
                "props": {
                  "heading": {
                    "type": "text",
                    "value": "Featured Collections"
                  },
                  "item_count": {
                    "type": "range",
                    "value": 4
                  },
                  "full_width": {
                    "type": "checkbox",
                    "value": false
                  },
                  "layout": {
                    "type": "select",
                    "value": "grid"
                  }
                }
              }
            ],
            "page_key": "collections"
          },
          {
            "page_sections": [],
            "page_key": "compare-products"
          },
          {
            "page_sections": [],
            "page_key": "product-description"
          },
          {
            "page_sections": [],
            "page_key": "product-listing"
          },
          {
            "page_sections": [],
            "page_key": "wishlist"
          }
        ]
      },
      "global_schema": {
        "props": [
          {
            "type": "color",
            "id": "header_bg_color",
            "default": "#ffffff",
            "category": "Header",
            "label": "Header background Color"
          },
          {
            "type": "color",
            "id": "footer_bg_color",
            "default": "#1b1b1b",
            "category": "Footer",
            "label": "Footer background Color"
          },
          {
            "type": "text",
            "id": "footer_text",
            "default": "",
            "category": "Footer",
            "label": "Footer Text"
          },
          {
            "type": "checkbox",
            "id": "disable_cart",
            "default": false,
            "category": "Cart",
            "label": "Disable Cart"
          }
        ]
      },
      "current": "default",
      "page_schema": [
        {
          "props": [],
          "_id": "5fe182f763d26d042fd205c4",
          "page": "add-product-review"
        },
        {
          "props": [],
          "_id": "5fe182f763d26dadc8d205c6",
          "page": "blog"
        },
        {
          "props": [],
          "_id": "5fe182f763d26d0a36d205c7",
          "page": "brands"
        },
        {
          "props": [
            {
              "type": "checkbox",
              "id": "gst",
              "label": "GST",
              "default": true,
              "info": "Show GST on cart"
            },
            {
              "type": "checkbox",
              "id": "staff_selection",
              "label": "Staff Selection",
              "default": true,
              "info": "Show Staff selection on Cart"
            },
            {
              "type": "checkbox",
              "id": "enable_customer",
              "label": "Customer",
              "default": true,
              "info": "Placing on behalf of customer"
            },
            {
              "type": "checkbox",
              "id": "enable_guest",
              "label": "Enable Guest Checkout",
              "default": true,
              "info": "Enable Continue as Guest"
            }
          ],
          "_id": "5fe182f763d26d81c5d205c8",
          "page": "cart-landing"
        },
        {
          "props": [],
          "_id": "5fe182f763d26d7e15d205c9",
          "page": "cart-review"
        },
        {
          "props": [],
          "_id": "5fe182f763d26d270ed205ca",
          "page": "categories"
        },
        {
          "props": [],
          "_id": "5fe182f763d26d9b4fd205cb",
          "page": "collection-listing"
        },
        {
          "props": [],
          "_id": "5fe182f763d26da6ecd205cc",
          "page": "collections"
        },
        {
          "props": [],
          "_id": "5fe182f763d26d7156d205cd",
          "page": "compare-products"
        },
        {
          "props": [],
          "_id": "5fe182f763d26d62bad205ce",
          "page": "home"
        },
        {
          "props": [],
          "_id": "5fe182f763d26d5afcd205cf",
          "page": "order-review"
        },
        {
          "props": [],
          "_id": "5fe182f763d26def8dd205d0",
          "page": "order-tracking-details"
        },
        {
          "props": [],
          "_id": "5fe182f763d26d381fd205d1",
          "page": "order-tracking"
        },
        {
          "props": [
            {
              "type": "text",
              "id": "shipping_link",
              "label": "Shipping Link",
              "default": "https://fynd.freshdesk.com/support/solutions/folders/33000111600",
              "info": "Link of shipping page"
            },
            {
              "type": "checkbox",
              "id": "original_image",
              "label": "Use original Image",
              "default": false,
              "info": "Use original product image."
            },
            {
              "type": "checkbox",
              "id": "wishlist",
              "label": "Wishlist",
              "default": true,
              "info": "Show Wishlist for product"
            },
            {
              "type": "checkbox",
              "id": "reviews",
              "label": "Review",
              "default": true,
              "info": "Show Reviews of product"
            },
            {
              "type": "checkbox",
              "id": "add_to_compare",
              "label": "Add to Compare",
              "default": true,
              "info": "Allow comparison of products"
            },
            {
              "type": "checkbox",
              "id": "size_guide",
              "label": "Size Guide",
              "default": true,
              "info": "Show Size Guide"
            },
            {
              "type": "checkbox",
              "id": "product_request",
              "label": "Product Request",
              "default": true,
              "info": "Show Product Request"
            },
            {
              "type": "checkbox",
              "id": "share",
              "label": "Share",
              "default": true,
              "info": "Enable Sharing product"
            },
            {
              "type": "checkbox",
              "id": "sold_by",
              "label": "Show Sold By",
              "default": true,
              "info": "Show name of the store"
            },
            {
              "type": "checkbox",
              "id": "store_selection",
              "label": "Seller Store Selection",
              "default": true,
              "info": "Allow to explicitly select stores"
            },
            {
              "type": "checkbox",
              "id": "compare_products",
              "label": "Compare Products",
              "default": true,
              "info": "Show Most Compared Products"
            },
            {
              "type": "checkbox",
              "id": "variants",
              "label": "Product Variants",
              "default": true,
              "info": "Show Product Variants"
            },
            {
              "type": "checkbox",
              "id": "ratings",
              "label": "Product Rating",
              "default": true,
              "info": "Show Product Ratings"
            },
            {
              "type": "checkbox",
              "id": "similar_products",
              "label": "Similar Products",
              "default": true,
              "info": "Show Similar Products"
            },
            {
              "type": "checkbox",
              "id": "bulk_prices",
              "label": "Bulk Prices",
              "default": true,
              "info": "Show Bulk Prices"
            },
            {
              "type": "checkbox",
              "id": "showDeliveryInfo",
              "label": "Delivery Info",
              "default": true,
              "info": "Show Delivery Date"
            }
          ],
          "_id": "5fe182f763d26d29bbd205d2",
          "page": "product-description"
        },
        {
          "props": [],
          "_id": "5fe182f763d26da5f0d205d3",
          "page": "product-listing"
        },
        {
          "props": [],
          "_id": "5fe182f763d26d3d18d205d4",
          "page": "product-reviews"
        },
        {
          "props": [],
          "_id": "5fe182f763d26d4e2dd205d5",
          "page": "wishlist"
        },
        {
          "props": [
            {
              "default": true,
              "id": "header",
              "label": "Header",
              "type": "checkbox"
            },
            {
              "default": true,
              "id": "footer",
              "label": "Footer",
              "type": "checkbox"
            }
          ],
          "_id": "5fe7166cbaae343115de8555",
          "page": "tesr"
        }
      ],
      "list": [
        {
          "name": "default",
          "global_config": {
            "static": {
              "props": {
                "colors": {
                  "primary_color": "#7043f7",
                  "secondary_color": "#02d1cb",
                  "accent_color": "#FFFFFF",
                  "link_color": "#7043f7",
                  "button_secondary_color": "#000000",
                  "bg_color": "#F8F8F8"
                }
              }
            },
            "custom": {
              "props": {
                "disable_cart": false
              }
            }
          },
          "page": [
            {
              "settings": {
                "props": {
                  "show_info_message": true
                }
              },
              "page": "cart-landing"
            }
          ],
          "_id": "5fc4bb9078e957737e7d79a3"
        }
      ]
    },
    "colors": {
      "primary_color": "#41434C",
      "secondary_color": "#41434C",
      "accent_color": "#FFFFFF",
      "link_color": "#33B1C0",
      "button_secondary_color": "#000000",
      "bg_color": "#F8F8F8"
    },
    "font": {
      "variants": {
        "light": {
          "name": "300",
          "file": ""
        },
        "regular": {
          "name": "regular",
          "file": "https://fonts.gstatic.com/s/judson/v13/FeVRS0Fbvbc14VxRD7N01bV7kg.ttf"
        },
        "medium": {
          "name": "500",
          "file": ""
        },
        "semi_bold": {
          "name": "600",
          "file": ""
        },
        "bold": {
          "name": "700",
          "file": "https://fonts.gstatic.com/s/judson/v13/FeVSS0Fbvbc14Vxps5xQ3Z5nm29Gww.ttf"
        }
      },
      "family": "Judson"
    },
    "applied": true,
    "published": false,
    "archived": false,
    "customized": true,
    "version": "1.1.19",
    "tags": [],
    "_id": "5fe17f7063d26dc54fd202b4",
    "pages": {
      "home": {
        "path": "",
        "type": "system",
        "sections": [],
        "value": "home",
        "text": "Home"
      },
      "brands": {
        "path": "brands",
        "type": "system",
        "sections": [],
        "value": "brands",
        "text": "Brands"
      },
      "cart-landing": {
        "path": "cart/bag",
        "type": "system",
        "sections": [],
        "value": "cart-landing",
        "text": "Cart Landing"
      },
      "categories": {
        "path": "categories",
        "type": "system",
        "sections": [],
        "value": "categories",
        "text": "Categories"
      },
      "collections": {
        "path": "collections",
        "type": "system",
        "sections": [],
        "value": "collections",
        "text": "Collections"
      },
      "compare-products": {
        "path": "compare",
        "type": "system",
        "sections": [],
        "value": "compare-products",
        "text": "Compare Products"
      },
      "product-description": {
        "path": "product",
        "type": "system",
        "sections": [],
        "value": "product-description",
        "text": "Product Description"
      },
      "product-listing": {
        "path": "products",
        "type": "system",
        "sections": [],
        "value": "product-listing",
        "text": "Product Listing"
      },
      "collection-listing": {
        "path": "collection",
        "type": "system",
        "sections": [],
        "value": "collection-listing",
        "text": "COllection Listing"
      },
      "wishlist": {
        "path": "wishlist",
        "type": "system",
        "sections": [],
        "value": "wishlist",
        "text": "Wishlist"
      }
    },
    "available_sections": [
      {
        "props": [
          {
            "type": "text",
            "id": "heading",
            "default": "Brands",
            "label": "Brands Heading"
          },
          {
            "type": "radio",
            "id": "brand_type",
            "default": "all",
            "options": [
              {
                "value": "all",
                "text": "All"
              },
              {
                "value": "department",
                "text": "Department"
              },
              {
                "value": "handpicked",
                "text": "Handpicked"
              }
            ]
          },
          {
            "type": "department",
            "id": "department",
            "label": "Department",
            "info": "Select a department of brands",
            "note": "Department only applies if 'department' type is selected"
          },
          {
            "type": "range",
            "id": "item_count",
            "min": 3,
            "max": 5,
            "step": 1,
            "unit": "",
            "label": "Items per row",
            "default": 4,
            "info": "Maximum items allowed per row"
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          },
          {
            "id": "layout",
            "type": "select",
            "options": [
              {
                "value": "grid",
                "text": "Grid View"
              },
              {
                "value": "horizontal",
                "text": "Horizontal View"
              }
            ],
            "default": "grid",
            "label": "Layout",
            "info": "Alignment of content"
          },
          {
            "type": "checkbox",
            "id": "view_all",
            "default": false,
            "label": "Show View All"
          }
        ],
        "blocks": [
          {
            "type": "brand-item",
            "name": "Brand Item",
            "props": [
              {
                "type": "brand",
                "id": "brand",
                "label": "Select Brand"
              }
            ]
          }
        ],
        "_id": "5feacca5bec232d59b89283a",
        "name": "brandTemplate",
        "label": "Brands List Page"
      },
      {
        "props": [
          {
            "type": "text",
            "id": "heading",
            "default": "Categories",
            "label": "Categories Heading"
          },
          {
            "id": "layout",
            "type": "select",
            "options": [
              {
                "value": "grid",
                "text": "Grid View"
              },
              {
                "value": "horizontal",
                "text": "Horizontal View"
              }
            ],
            "default": "grid",
            "label": "Layout",
            "info": "Alignment of content"
          },
          {
            "type": "range",
            "id": "item_count",
            "min": 3,
            "max": 5,
            "step": 1,
            "unit": "",
            "label": "Items per row",
            "default": 4,
            "info": "Maximum items allowed per row"
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          },
          {
            "type": "checkbox",
            "id": "view_all",
            "default": false,
            "label": "Show View All",
            "info": "Check to show View All Button"
          }
        ],
        "blocks": [],
        "_id": "5feacca5bec232404189283b",
        "name": "categoriesTemplate",
        "label": "Categories Page"
      },
      {
        "props": [
          {
            "type": "text",
            "id": "title",
            "default": "",
            "label": "Title"
          },
          {
            "type": "text",
            "id": "subtitle",
            "default": "",
            "label": "Subtitle"
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          }
        ],
        "blocks": [
          {
            "type": "collection",
            "name": "Collection",
            "props": [
              {
                "type": "collection",
                "id": "collection",
                "label": "Select a collection"
              },
              {
                "type": "image_picker",
                "id": "overlayImage",
                "default": "",
                "label": "Background image",
                "info": "Background Image"
              },
              {
                "type": "text",
                "id": "title",
                "default": "",
                "label": "Overlay title"
              },
              {
                "type": "text",
                "id": "subtitle",
                "default": "",
                "label": "Overlay subtitle"
              }
            ]
          }
        ],
        "_id": "5feacca5bec2321fd589283c",
        "name": "collectionGrid",
        "label": "Collection Grid"
      },
      {
        "props": [
          {
            "type": "text",
            "id": "heading",
            "default": "Collection List",
            "label": "Collection Heading",
            "info": "Collection Heading"
          },
          {
            "type": "range",
            "id": "item_count",
            "min": 3,
            "max": 5,
            "step": 1,
            "unit": "",
            "label": "Collections per row",
            "default": 2,
            "info": "Maximum items allowed per row"
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          }
        ],
        "blocks": [
          {
            "type": "collection",
            "name": "Collection",
            "props": [
              {
                "type": "collection",
                "id": "collection",
                "label": "Select a collection"
              }
            ]
          }
        ],
        "_id": "5feacca5bec2323bf689283d",
        "name": "collectionList",
        "label": "Collection List"
      },
      {
        "props": [
          {
            "type": "text",
            "id": "heading",
            "default": "Collections",
            "label": "Collection Heading"
          },
          {
            "type": "range",
            "id": "item_count",
            "min": 3,
            "max": 5,
            "step": 1,
            "unit": "",
            "label": "Collections per row",
            "default": 4,
            "info": "Maximum items allowed per row"
          },
          {
            "type": "radio",
            "id": "collection_type",
            "default": "all",
            "options": [
              {
                "value": "all",
                "text": "All"
              },
              {
                "value": "handpicked",
                "text": "Handpicked"
              }
            ]
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          },
          {
            "id": "layout",
            "type": "select",
            "options": [
              {
                "value": "grid",
                "text": "Grid View"
              },
              {
                "value": "horizontal",
                "text": "Horizontal View"
              }
            ],
            "default": "grid",
            "label": "Layout",
            "info": "Alignment of content"
          },
          {
            "type": "checkbox",
            "id": "view_all",
            "default": false,
            "label": "Show View All"
          }
        ],
        "blocks": [
          {
            "type": "collection-item",
            "name": "Collection Item",
            "props": [
              {
                "type": "collection",
                "id": "collection",
                "label": "Select collection"
              }
            ]
          }
        ],
        "_id": "5feacca5bec23263b489283e",
        "name": "collectionTemplate",
        "label": "Collection List Page"
      },
      {
        "props": [
          {
            "type": "code",
            "id": "code",
            "label": "Custom HTML",
            "info": "Add Your custom HTML Code below. You can also use the full screen icon to open a code editor and add your code"
          }
        ],
        "blocks": [],
        "_id": "5feacca5bec232409489283f",
        "name": "customHtml",
        "label": "Custom HTML"
      },
      {
        "props": [
          {
            "type": "text",
            "id": "heading",
            "default": "Featured Products",
            "label": "Section Heading"
          },
          {
            "type": "range",
            "id": "item_count",
            "min": 3,
            "max": 5,
            "step": 1,
            "unit": "",
            "label": "Products per row",
            "default": 4,
            "info": "Maximum items allowed per row"
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          }
        ],
        "blocks": [
          {
            "type": "product",
            "name": "Product",
            "props": [
              {
                "type": "product",
                "id": "product",
                "label": "Select a Product",
                "info": "Product Item to be displayed"
              }
            ]
          }
        ],
        "_id": "5feacca5bec2326213892840",
        "name": "featuredProducts",
        "label": "Featured Products",
        "preset": {
          "blocks": [
            {
              "name": "Product"
            },
            {
              "name": "Product"
            },
            {
              "name": "Product"
            },
            {
              "name": "Product"
            }
          ]
        }
      },
      {
        "props": [
          {
            "type": "range",
            "id": "item_count",
            "min": 1,
            "max": 4,
            "step": 1,
            "unit": "",
            "label": "No of items",
            "default": 4,
            "info": "Maximum items allowed per row"
          }
        ],
        "blocks": [
          {
            "type": "gallery_image",
            "name": "Image",
            "props": [
              {
                "type": "image_picker",
                "id": "image",
                "label": "Gallery Image",
                "default": "https://hdn-1.fynd.com/company/163/applications/5d5265c8f8ad9cae6dbf18f1/theme/pictures/free/original/theme-image-1603773049684.svg"
              },
              {
                "type": "text",
                "id": "caption",
                "label": "Image Caption",
                "default": ""
              },
              {
                "type": "url",
                "id": "link",
                "label": "Link",
                "default": "",
                "info": "Link to redirect"
              }
            ]
          }
        ],
        "_id": "5feacca5bec2321047892841",
        "name": "gallery",
        "label": "Gallery",
        "preset": {
          "props": {
            "item_count": 4
          },
          "blocks": [
            {
              "name": "Image"
            },
            {
              "name": "Image"
            },
            {
              "name": "Image"
            },
            {
              "name": "Image"
            }
          ]
        }
      },
      {
        "props": [
          {
            "type": "url",
            "id": "ctaLink",
            "default": "",
            "label": "Redirect Link"
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          }
        ],
        "blocks": [],
        "_id": "5feacca5bec232b085892842",
        "name": "heroBanner",
        "label": "Hero Banner"
      },
      {
        "props": [
          {
            "id": "image",
            "type": "image_picker",
            "label": "Hero Image",
            "default": ""
          },
          {
            "type": "text",
            "id": "heading",
            "default": "",
            "label": "Heading"
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          },
          {
            "id": "overlayLayout",
            "type": "select",
            "options": [
              {
                "value": "left",
                "text": "Align Left"
              },
              {
                "value": "center",
                "text": "Align Center"
              },
              {
                "value": "right",
                "text": "Align Right"
              }
            ],
            "default": "left",
            "label": "Overlay Layout",
            "info": "Alignment of overlay content"
          },
          {
            "type": "image_picker",
            "id": "overlayImage",
            "default": "",
            "label": "Overlay image",
            "info": "Overlay Image"
          },
          {
            "type": "text",
            "id": "text",
            "default": "",
            "label": "Overlay Text"
          },
          {
            "type": "color",
            "id": "text_color",
            "default": "#000",
            "label": "Text Color"
          },
          {
            "type": "url",
            "id": "ctaLink",
            "default": "",
            "label": "Redirect Link"
          },
          {
            "type": "text",
            "id": "ctaText",
            "default": "Shop Now",
            "label": "Button Text"
          }
        ],
        "blocks": [],
        "_id": "5feacca5bec2321e74892843",
        "name": "imageBanner",
        "label": "Hero Image"
      },
      {
        "props": [
          {
            "id": "videoUrl",
            "type": "url",
            "label": "Video URL",
            "default": ""
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          },
          {
            "id": "coverUrl",
            "type": "image_picker",
            "label": "Video Cover Image URL",
            "default": ""
          },
          {
            "type": "checkbox",
            "id": "showcontrols",
            "default": false,
            "label": "Show Controls on Video",
            "info": "Check to show controls on video"
          },
          {
            "type": "select",
            "id": "size",
            "options": [
              {
                "value": "adapt",
                "text": "Adapt"
              },
              {
                "value": "small",
                "text": "Small"
              },
              {
                "value": "medium",
                "text": "Medium"
              },
              {
                "value": "large",
                "text": "Large"
              }
            ],
            "default": "adapt",
            "label": "Video Height",
            "info": "Height of Video"
          },
          {
            "type": "text",
            "id": "heading",
            "default": "",
            "label": "Heading"
          },
          {
            "type": "color",
            "id": "heading_color",
            "default": "#000",
            "label": "Headin Text Color"
          },
          {
            "type": "text",
            "id": "subHeading",
            "default": "",
            "label": "Sub-heading"
          },
          {
            "type": "color",
            "id": "subheading_color",
            "default": "#000",
            "label": "Subheading Text Color"
          },
          {
            "type": "url",
            "id": "ctaLink",
            "default": "",
            "label": "Redirect Link"
          },
          {
            "type": "text",
            "id": "ctaText",
            "default": "ShopNow",
            "label": "Button Text"
          }
        ],
        "blocks": [],
        "_id": "5feacca5bec232bfc8892844",
        "name": "videoBanner",
        "label": "Hero Video"
      },
      {
        "props": [
          {
            "type": "select",
            "id": "slide_height",
            "options": [
              {
                "value": "adapt",
                "text": "Adapt to first image"
              },
              {
                "value": "small",
                "text": "Small"
              },
              {
                "value": "medium",
                "text": "Medium"
              },
              {
                "value": "large",
                "text": "Large"
              }
            ],
            "default": "adapt",
            "label": "Slide height",
            "info": "Size of the slide"
          },
          {
            "type": "checkbox",
            "id": "autoplay",
            "default": false,
            "label": "AutoPlay Slides",
            "info": "Check to autoplay slides"
          },
          {
            "type": "range",
            "id": "slide_interval",
            "min": 1,
            "max": 10,
            "step": 1,
            "unit": "sec",
            "label": "Change slides after every",
            "default": 2,
            "info": "Autoplay slide duration"
          }
        ],
        "blocks": [
          {
            "type": "gallery_image",
            "name": "Image",
            "props": [
              {
                "type": "image_picker",
                "id": "image",
                "label": "Gallery Image"
              },
              {
                "type": "url",
                "id": "slide_link",
                "label": "Slide Link"
              }
            ]
          }
        ],
        "_id": "5feacca5bec232a916892845",
        "name": "image-carousel",
        "label": "Image Carousel",
        "preset": {
          "blocks": [
            {
              "name": "Image"
            },
            {
              "name": "Image"
            }
          ]
        }
      },
      {
        "props": [
          {
            "type": "text",
            "id": "heading",
            "default": "Featured Collections",
            "label": "Collection Heading"
          },
          {
            "type": "range",
            "id": "item_count",
            "min": 2,
            "max": 5,
            "step": 1,
            "unit": "",
            "label": "Products per row",
            "default": 2,
            "info": "Maximum items allowed per row"
          },
          {
            "type": "collection",
            "id": "collection",
            "label": "Collection",
            "info": "Select a collection to display its products"
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          }
        ],
        "blocks": [],
        "_id": "5feacca5bec23258ec892846",
        "name": "productList",
        "label": "Product List",
        "preset": {
          "props": {
            "heading": "Featured Products",
            "item_count": 4,
            "collection": ""
          }
        }
      },
      {
        "props": [
          {
            "type": "checkbox",
            "id": "autoplay",
            "default": false,
            "label": "AutoPlay Slides"
          },
          {
            "type": "range",
            "id": "slide_interval",
            "min": 1,
            "max": 10,
            "step": 1,
            "unit": "sec",
            "label": "Change slides every",
            "default": 2
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          }
        ],
        "blocks": [
          {
            "type": "testimonial",
            "name": "Testimonial",
            "props": [
              {
                "type": "textarea",
                "id": "testimonialText",
                "label": "Text for Testimonial",
                "default": "",
                "info": "Text for testimonial",
                "placeholder": "Text"
              },
              {
                "type": "text",
                "id": "author",
                "label": "Customers name"
              }
            ]
          }
        ],
        "_id": "5feacca5bec23299e8892847",
        "name": "testimonials",
        "label": "Testimonial"
      },
      {
        "props": [
          {
            "type": "select",
            "id": "slide_height",
            "options": [
              {
                "value": "adapt",
                "text": "Adapt to first video"
              },
              {
                "value": "small",
                "text": "Small"
              },
              {
                "value": "medium",
                "text": "Medium"
              },
              {
                "value": "large",
                "text": "Large"
              }
            ],
            "default": "adapt",
            "label": "Slide height",
            "info": "Size of the slide"
          },
          {
            "type": "checkbox",
            "id": "full_width",
            "default": false,
            "label": "Full width",
            "info": "Check to allow items to take entire width of the viewport"
          },
          {
            "type": "checkbox",
            "id": "autoplay",
            "default": false,
            "label": "AutoPlay Slides",
            "info": "Check to autoplay slides"
          },
          {
            "type": "range",
            "id": "slide_interval",
            "min": 1,
            "max": 10,
            "step": 1,
            "unit": "sec",
            "label": "Change slides every",
            "default": 2,
            "info": "Autoplay slide duration"
          }
        ],
        "blocks": [
          {
            "type": "video_item",
            "name": "Video Slide",
            "props": [
              {
                "id": "videoUrl",
                "type": "url",
                "label": "Video URL",
                "default": ""
              },
              {
                "type": "checkbox",
                "id": "showcontrols",
                "default": false,
                "label": "Show Controls on Video",
                "info": "Check to show controls on video"
              }
            ]
          }
        ],
        "_id": "5feacca5bec232d89b892848",
        "name": "videoCarousel",
        "label": "Video Carousel"
      }
    ],
    "sections": [
      {
        "page_sections": [
          {
            "blocks": [],
            "predicate": {
              "screen": {
                "mobile": true,
                "desktop": true,
                "tablet": true
              },
              "user": {
                "authenticated": true,
                "anonymous": true
              },
              "route": {
                "selected": "none",
                "query": {},
                "exact_u_r_l": ""
              }
            },
            "name": "brandTemplate",
            "props": {
              "title": {
                "type": "text",
                "value": "Brands"
              },
              "item_count": {
                "type": "range",
                "value": 4
              },
              "full_width": {
                "type": "checkbox",
                "value": false
              },
              "heading": {
                "type": "text",
                "value": "Brands"
              },
              "brand_type": {
                "type": "radio",
                "value": "all"
              },
              "department": {
                "type": "department",
                "value": null
              },
              "layout": {
                "type": "select",
                "value": "grid"
              },
              "view_all": {
                "type": "checkbox",
                "value": null
              }
            }
          }
        ],
        "page_key": "product-description"
      },
      {
        "page_sections": [],
        "page_key": "collection-listing"
      },
      {
        "page_sections": [
          {
            "blocks": [
              {
                "type": "brand-item",
                "name": "Brand Item",
                "props": {
                  "brand": {
                    "type": "brand",
                    "value": {
                      "display": "Apple",
                      "id": "apple"
                    }
                  }
                }
              },
              {
                "type": "brand-item",
                "name": "Brand Item",
                "props": {
                  "brand": {
                    "type": "brand",
                    "value": {
                      "display": "Gionee",
                      "id": "gionee"
                    }
                  }
                }
              }
            ],
            "predicate": {
              "screen": {
                "mobile": true,
                "desktop": true,
                "tablet": true
              },
              "user": {
                "authenticated": true,
                "anonymous": true
              },
              "route": {
                "selected": "none",
                "query": {},
                "exact_u_r_l": ""
              }
            },
            "name": "brandTemplate",
            "props": {
              "heading": {
                "type": "text",
                "value": "Brandss"
              },
              "brand_type": {
                "type": "radio",
                "value": "handpicked"
              },
              "department": {
                "type": "department"
              },
              "item_count": {
                "value": 4,
                "type": "range"
              },
              "full_width": {
                "value": false,
                "type": "checkbox"
              },
              "layout": {
                "type": "select",
                "value": "grid"
              },
              "view_all": {
                "value": false,
                "type": "checkbox"
              }
            }
          }
        ],
        "page_key": "home"
      },
      {
        "page_sections": [
          {
            "blocks": [],
            "predicate": {
              "screen": {
                "mobile": true,
                "desktop": true,
                "tablet": true
              },
              "user": {
                "authenticated": true,
                "anonymous": true
              },
              "route": {
                "selected": "none",
                "query": {},
                "exact_u_r_l": ""
              }
            },
            "name": "brandTemplate",
            "props": {
              "title": {
                "type": "text",
                "value": "Brands"
              },
              "item_count": {
                "type": "range",
                "value": 4
              },
              "full_width": {
                "type": "checkbox",
                "value": false
              }
            }
          }
        ],
        "page_key": "brands"
      },
      {
        "page_sections": [],
        "page_key": "cart-landing"
      },
      {
        "page_sections": [
          {
            "blocks": [],
            "predicate": {
              "screen": {
                "mobile": true,
                "desktop": true,
                "tablet": true
              },
              "user": {
                "authenticated": true,
                "anonymous": true
              },
              "route": {
                "selected": "none",
                "query": {},
                "exact_u_r_l": ""
              }
            },
            "name": "categoriesTemplate",
            "props": {
              "item_count": {
                "type": "range",
                "value": 4
              },
              "full_width": {
                "type": "checkbox",
                "value": false
              }
            }
          }
        ],
        "page_key": "categories"
      },
      {
        "page_sections": [
          {
            "blocks": [],
            "predicate": {
              "screen": {
                "mobile": true,
                "desktop": true,
                "tablet": true
              },
              "user": {
                "authenticated": true,
                "anonymous": true
              },
              "route": {
                "selected": "none",
                "query": {},
                "exact_u_r_l": ""
              }
            },
            "name": "collectionTemplate",
            "props": {
              "heading": {
                "type": "text",
                "value": "Featured Collections"
              },
              "item_count": {
                "type": "range",
                "value": 4
              },
              "full_width": {
                "type": "checkbox",
                "value": false
              }
            }
          }
        ],
        "page_key": "collections"
      },
      {
        "page_sections": [],
        "page_key": "compare-products"
      },
      {
        "page_sections": [],
        "page_key": "product-listing"
      },
      {
        "page_sections": [],
        "page_key": "wishlist"
      },
      {
        "page_sections": [],
        "page_key": "tesr"
      }
    ],
    "application": "5e737afb97e0f586bf9d04db",
    "available_pages": [
      {
        "path": "product",
        "type": "system",
        "sections": [],
        "text": "Product Description",
        "value": "product-description"
      },
      {
        "path": "collection",
        "type": "system",
        "sections": [],
        "text": "Collection Listing",
        "value": "collection-listing"
      },
      {
        "path": "",
        "type": "system",
        "sections": [
          {
            "_id": "5feacca5bec2327927892853",
            "attributes": {
              "page": "home"
            }
          }
        ],
        "value": "home",
        "text": "Home"
      },
      {
        "path": "brands",
        "type": "system",
        "sections": [
          {
            "_id": "5feacca5bec2326422892854",
            "attributes": {
              "page": "brands"
            }
          }
        ],
        "value": "brands",
        "text": "Brands"
      },
      {
        "path": "cart/bag",
        "type": "system",
        "sections": [],
        "value": "cart-landing",
        "text": "Cart Landing"
      },
      {
        "path": "categories",
        "type": "system",
        "sections": [
          {
            "_id": "5feacca5bec232424c892855",
            "attributes": {
              "page": "categories"
            }
          }
        ],
        "value": "categories",
        "text": "Categories"
      },
      {
        "path": "collections",
        "type": "system",
        "sections": [
          {
            "_id": "5feacca5bec23281de892856",
            "attributes": {
              "page": "collections"
            }
          }
        ],
        "value": "collections",
        "text": "Collections"
      },
      {
        "path": "compare",
        "type": "system",
        "sections": [],
        "value": "compare-products",
        "text": "Compare Products"
      },
      {
        "path": "products",
        "type": "system",
        "sections": [],
        "value": "product-listing",
        "text": "Product Listing"
      },
      {
        "path": "wishlist",
        "type": "system",
        "sections": [],
        "value": "wishlist",
        "text": "Wishlist"
      },
      {
        "path": "sections/test",
        "type": "sections",
        "sections": [],
        "text": "test",
        "value": "test"
      }
    ],
    "styles": {},
    "created_at": "2020-12-22T05:09:04.720Z",
    "updated_at": "2021-01-24T11:22:41.376Z"
  }
}
```









---



---


## User


#### loginWithFacebook
Login or Register using Facebook

```golang

 data, err :=  User.LoginWithFacebook(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Platform`

| body |  OAuthRequestSchema | "Request body" 


Use this API to login or register using Facebook credentials.

*Success Response:*



Success. Returns a JSON object with the user details. Check the example shown below or refer `AuthSuccess` for more details.


Schema: `AuthSuccess`


*Examples:*


Success
```json
{
  "value": {
    "user_exists": false,
    "user": {
      "emails": [
        {
          "email": "www.akash24@gmail.com",
          "verified": true,
          "primary": true,
          "active": true
        }
      ],
      "phone_numbers": [],
      "first_name": "Akash",
      "last_name": "Mane",
      "debug": {
        "platform": "Fynd"
      },
      "active": true
    },
    "register_token": "d960c388-e286-43d9-b688-f6d1decc632d"
  }
}
```









---


#### loginWithGoogle
Login or Register using Google

```golang

 data, err :=  User.LoginWithGoogle(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Platform`

| body |  OAuthRequestSchema | "Request body" 


Use this API to login or register using Google Account credentials.

*Success Response:*



Success. Returns a JSON object with the user details. Check the example shown below or refer `AuthSuccess` for more details.


Schema: `AuthSuccess`


*Examples:*


Success
```json
{
  "value": {
    "user_exists": false,
    "user": {
      "emails": [
        {
          "email": "www.akash24@gmail.com",
          "verified": true,
          "primary": true,
          "active": true
        }
      ],
      "phone_numbers": [],
      "first_name": "Akash",
      "last_name": "Mane",
      "debug": {
        "platform": "Fynd"
      },
      "active": true
    },
    "register_token": "d960c388-e286-43d9-b688-f6d1decc632d"
  }
}
```









---


#### loginWithGoogleAndroid
Login or Register using Google on Android

```golang

 data, err :=  User.LoginWithGoogleAndroid(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Platform`

| body |  OAuthRequestSchema | "Request body" 


Use this API to login or register in Android app using Google Account credentials.

*Success Response:*



Success. Returns a JSON object with the user details. Check the example shown below or refer `AuthSuccess` for more details.


Schema: `AuthSuccess`


*Examples:*


Success
```json
{
  "value": {
    "user_exists": false,
    "user": {
      "emails": [
        {
          "email": "www.akash24@gmail.com",
          "verified": true,
          "primary": true,
          "active": true
        }
      ],
      "phone_numbers": [],
      "first_name": "Akash",
      "last_name": "Mane",
      "debug": {
        "platform": "Fynd"
      },
      "active": true
    },
    "register_token": "d960c388-e286-43d9-b688-f6d1decc632d"
  }
}
```









---


#### loginWithGoogleIOS
Login or Register using Google on iOS

```golang

 data, err :=  User.LoginWithGoogleIOS(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Platform`

| body |  OAuthRequestSchema | "Request body" 


Use this API to login or register in iOS app using Google Account credentials.

*Success Response:*



Success. Returns a JSON object with the user details. Check the example shown below or refer `AuthSuccess` for more details.


Schema: `AuthSuccess`


*Examples:*


Success
```json
{
  "value": {
    "user_exists": false,
    "user": {
      "emails": [
        {
          "email": "www.akash24@gmail.com",
          "verified": true,
          "primary": true,
          "active": true
        }
      ],
      "phone_numbers": [],
      "first_name": "Akash",
      "last_name": "Mane",
      "debug": {
        "platform": "Fynd"
      },
      "active": true
    },
    "register_token": "d960c388-e286-43d9-b688-f6d1decc632d"
  }
}
```









---


#### loginWithAppleIOS
Login or Register using Apple on iOS

```golang

 data, err :=  User.LoginWithAppleIOS(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Platform`

| body |  OAuthRequestAppleSchema | "Request body" 


Use this API to login or register in iOS app using Apple Account credentials.

*Success Response:*



Success. Returns a JSON object with the user details. Check the example shown below or refer `AuthSuccess` for more details.


Schema: `AuthSuccess`


*Examples:*


Success
```json
{
  "value": {
    "user_exists": false,
    "user": {
      "emails": [
        {
          "email": "www.akash24@gmail.com",
          "verified": true,
          "primary": true,
          "active": true
        }
      ],
      "phone_numbers": [],
      "first_name": "Akash",
      "last_name": "Mane",
      "debug": {
        "platform": "Fynd"
      },
      "active": true
    },
    "register_token": "d960c388-e286-43d9-b688-f6d1decc632d"
  }
}
```









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
  "value": {
    "success": true,
    "request_id": "01503005aeab87cbed93d40f46cc2749",
    "message": "OTP sent",
    "mobile": "8652523958",
    "country_code": "91",
    "resend_timer": 30,
    "resendToken": "58e72ca0-66ae-11eb-98b1-77d61363826e"
  }
}
```









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
  "value": {
    "user": {
      "debug": {
        "source": "deadlock",
        "platform": "000000000000000000000001"
      },
      "gender": "male",
      "account_type": "user",
      "active": true,
      "profile_pic_url": "https://d2co8r51m5ca2d.cloudfront.net/inapp_banners/default_profile_img.png",
      "has_old_password_hash": false,
      "_id": "5e68af49cfa09bf7233022f1",
      "first_name": "Akash",
      "last_name": "Mane",
      "username": "akashmane_gofynd_com_10039",
      "phone_numbers": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "phone": "8652523958",
          "country_code": 91
        }
      ],
      "emails": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "email": "akashmane@gofynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@fynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@uniket.store"
        }
      ],
      "created_at": "2020-03-11T09:28:41.982Z",
      "updated_at": "2021-02-04T10:10:44.981Z"
    }
  }
}
```









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









---


#### sendResetPasswordMobile
Reset Password

```golang

 data, err :=  User.SendResetPasswordMobile(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Platform`

| body |  SendResetPasswordMobileRequestSchema | "Request body" 


Use this API to reset a password using the link sent on mobile.

*Success Response:*



Success. Check the example shown below or refer `ResetPasswordSuccess` for more details.


Schema: `ResetPasswordSuccess`









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
  "value": {
    "user": {
      "debug": {
        "source": "deadlock",
        "platform": "000000000000000000000001"
      },
      "gender": "male",
      "account_type": "user",
      "active": true,
      "profile_pic_url": "https://d2co8r51m5ca2d.cloudfront.net/inapp_banners/default_profile_img.png",
      "has_old_password_hash": false,
      "_id": "5e68af49cfa09bf7233022f1",
      "first_name": "Akash",
      "last_name": "Mane",
      "username": "akashmane_gofynd_com_10039",
      "phone_numbers": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "phone": "8652523958",
          "country_code": 91
        }
      ],
      "emails": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "email": "akashmane@gofynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@fynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@uniket.store"
        }
      ],
      "created_at": "2020-03-11T09:28:41.982Z",
      "updated_at": "2021-02-04T10:10:44.981Z"
    }
  }
}
```









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
  "value": {
    "user": {
      "debug": {
        "source": "deadlock",
        "platform": "000000000000000000000001"
      },
      "gender": "male",
      "account_type": "user",
      "active": true,
      "profile_pic_url": "https://d2co8r51m5ca2d.cloudfront.net/inapp_banners/default_profile_img.png",
      "has_old_password_hash": false,
      "_id": "5e68af49cfa09bf7233022f1",
      "first_name": "Akash",
      "last_name": "Mane",
      "username": "akashmane_gofynd_com_10039",
      "phone_numbers": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "phone": "8652523958",
          "country_code": 91
        }
      ],
      "emails": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "email": "akashmane@gofynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@fynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@uniket.store"
        }
      ],
      "created_at": "2020-03-11T09:28:41.982Z",
      "updated_at": "2021-02-04T10:10:44.981Z"
    }
  }
}
```









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









---


#### deleteUser
verify otp and delete user

```golang

 data, err :=  User.DeleteUser(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  DeleteApplicationUserRequestSchema | "Request body" 


verify otp and delete user

*Success Response:*



Success. Returns a success message. Refer `DeleteUserSuccess` for more details.


Schema: `DeleteUserSuccess`









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
  "value": {
    "verify_mobile_link": true,
    "user": {
      "debug": {
        "source": "deadlock",
        "platform": "000000000000000000000001"
      },
      "gender": "male",
      "account_type": "user",
      "active": true,
      "profile_pic_url": "https://d2co8r51m5ca2d.cloudfront.net/inapp_banners/default_profile_img.png",
      "has_old_password_hash": false,
      "_id": "5e68af49cfa09bf7233022f1",
      "first_name": "Akash",
      "last_name": "Mane",
      "username": "akashmane_gofynd_com_10039",
      "phone_numbers": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "phone": "8652523958",
          "country_code": 91
        }
      ],
      "emails": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "email": "akashmane@gofynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@fynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@uniket.store"
        }
      ],
      "created_at": "2020-03-11T09:28:41.982Z",
      "updated_at": "2021-02-04T10:10:44.981Z"
    }
  }
}
```









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
  "value": {
    "verify_mobile_link": true,
    "user": {
      "debug": {
        "source": "deadlock",
        "platform": "000000000000000000000001"
      },
      "gender": "male",
      "account_type": "user",
      "active": true,
      "profile_pic_url": "https://d2co8r51m5ca2d.cloudfront.net/inapp_banners/default_profile_img.png",
      "has_old_password_hash": false,
      "_id": "5e68af49cfa09bf7233022f1",
      "first_name": "Akash",
      "last_name": "Mane",
      "username": "akashmane_gofynd_com_10039",
      "phone_numbers": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "phone": "8652523958",
          "country_code": 91
        }
      ],
      "emails": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "email": "akashmane@gofynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@fynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@uniket.store"
        }
      ],
      "created_at": "2020-03-11T09:28:41.982Z",
      "updated_at": "2021-02-04T10:10:44.981Z"
    }
  }
}
```









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
  "value": {
    "user": {
      "debug": {
        "source": "deadlock",
        "platform": "000000000000000000000001"
      },
      "gender": "male",
      "account_type": "user",
      "active": true,
      "profile_pic_url": "https://d2co8r51m5ca2d.cloudfront.net/inapp_banners/default_profile_img.png",
      "has_old_password_hash": false,
      "_id": "5e68af49cfa09bf7233022f1",
      "first_name": "Akash",
      "last_name": "Mane",
      "username": "akashmane_gofynd_com_10039",
      "phone_numbers": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "phone": "8652523958",
          "country_code": 91
        }
      ],
      "emails": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "email": "akashmane@gofynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@fynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@uniket.store"
        }
      ],
      "created_at": "2020-03-11T09:28:41.982Z",
      "updated_at": "2021-02-04T10:10:44.981Z"
    }
  }
}
```









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
  "value": {
    "user": {
      "debug": {
        "source": "deadlock",
        "platform": "000000000000000000000001"
      },
      "gender": "male",
      "account_type": "user",
      "active": true,
      "profile_pic_url": "https://d2co8r51m5ca2d.cloudfront.net/inapp_banners/default_profile_img.png",
      "has_old_password_hash": false,
      "_id": "5e68af49cfa09bf7233022f1",
      "first_name": "Akash",
      "last_name": "Mane",
      "username": "akashmane_gofynd_com_10039",
      "phone_numbers": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "phone": "8652523958",
          "country_code": 91
        }
      ],
      "emails": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "email": "akashmane@gofynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@fynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@uniket.store"
        }
      ],
      "created_at": "2020-03-11T09:28:41.982Z",
      "updated_at": "2021-02-04T10:10:44.981Z"
    }
  }
}
```









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
  "value": {
    "verify_mobile_link": true,
    "user": {
      "debug": {
        "source": "deadlock",
        "platform": "000000000000000000000001"
      },
      "gender": "male",
      "account_type": "user",
      "active": true,
      "profile_pic_url": "https://d2co8r51m5ca2d.cloudfront.net/inapp_banners/default_profile_img.png",
      "has_old_password_hash": false,
      "_id": "5e68af49cfa09bf7233022f1",
      "first_name": "Akash",
      "last_name": "Mane",
      "username": "akashmane_gofynd_com_10039",
      "phone_numbers": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "phone": "8652523958",
          "country_code": 91
        }
      ],
      "emails": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "email": "akashmane@gofynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@fynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@uniket.store"
        }
      ],
      "created_at": "2020-03-11T09:28:41.982Z",
      "updated_at": "2021-02-04T10:10:44.981Z"
    }
  }
}
```









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
  "value": {
    "user": {
      "debug": {
        "source": "deadlock",
        "platform": "000000000000000000000001"
      },
      "gender": "male",
      "account_type": "user",
      "active": true,
      "profile_pic_url": "https://d2co8r51m5ca2d.cloudfront.net/inapp_banners/default_profile_img.png",
      "has_old_password_hash": false,
      "_id": "5e68af49cfa09bf7233022f1",
      "first_name": "Akash",
      "last_name": "Mane",
      "username": "akashmane_gofynd_com_10039",
      "phone_numbers": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "phone": "8652523958",
          "country_code": 91
        }
      ],
      "emails": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "email": "akashmane@gofynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@fynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@uniket.store"
        }
      ],
      "created_at": "2020-03-11T09:28:41.982Z",
      "updated_at": "2021-02-04T10:10:44.981Z"
    }
  }
}
```









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
  "value": {
    "user": {
      "debug": {
        "source": "deadlock",
        "platform": "000000000000000000000001"
      },
      "gender": "male",
      "account_type": "user",
      "active": true,
      "profile_pic_url": "https://d2co8r51m5ca2d.cloudfront.net/inapp_banners/default_profile_img.png",
      "has_old_password_hash": false,
      "_id": "5e68af49cfa09bf7233022f1",
      "first_name": "Akash",
      "last_name": "Mane",
      "username": "akashmane_gofynd_com_10039",
      "phone_numbers": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "phone": "8652523958",
          "country_code": 91
        }
      ],
      "emails": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "email": "akashmane@gofynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@fynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@uniket.store"
        }
      ],
      "created_at": "2020-03-11T09:28:41.982Z",
      "updated_at": "2021-02-04T10:10:44.981Z"
    }
  }
}
```









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
  "value": {
    "verify_mobile_link": true,
    "user": {
      "debug": {
        "source": "deadlock",
        "platform": "000000000000000000000001"
      },
      "gender": "male",
      "account_type": "user",
      "active": true,
      "profile_pic_url": "https://d2co8r51m5ca2d.cloudfront.net/inapp_banners/default_profile_img.png",
      "has_old_password_hash": false,
      "_id": "5e68af49cfa09bf7233022f1",
      "first_name": "Akash",
      "last_name": "Mane",
      "username": "akashmane_gofynd_com_10039",
      "phone_numbers": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "phone": "8652523958",
          "country_code": 91
        }
      ],
      "emails": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "email": "akashmane@gofynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@fynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@uniket.store"
        }
      ],
      "created_at": "2020-03-11T09:28:41.982Z",
      "updated_at": "2021-02-04T10:10:44.981Z"
    }
  }
}
```









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
  "value": {
    "verify_email_link": true,
    "user": {
      "debug": {
        "source": "deadlock",
        "platform": "000000000000000000000001"
      },
      "gender": "male",
      "account_type": "user",
      "active": true,
      "profile_pic_url": "https://d2co8r51m5ca2d.cloudfront.net/inapp_banners/default_profile_img.png",
      "has_old_password_hash": false,
      "_id": "5e68af49cfa09bf7233022f1",
      "first_name": "Akash",
      "last_name": "Mane",
      "username": "akashmane_gofynd_com_10039",
      "phone_numbers": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "phone": "8652523958",
          "country_code": 91
        }
      ],
      "emails": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "email": "akashmane@gofynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@fynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@uniket.store"
        }
      ],
      "created_at": "2020-03-11T09:28:41.982Z",
      "updated_at": "2021-02-04T10:10:44.981Z"
    }
  }
}
```









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
  "value": {
    "user": {
      "debug": {
        "source": "deadlock",
        "platform": "000000000000000000000001"
      },
      "gender": "male",
      "account_type": "user",
      "active": true,
      "profile_pic_url": "https://d2co8r51m5ca2d.cloudfront.net/inapp_banners/default_profile_img.png",
      "has_old_password_hash": false,
      "_id": "5e68af49cfa09bf7233022f1",
      "first_name": "Akash",
      "last_name": "Mane",
      "username": "akashmane_gofynd_com_10039",
      "phone_numbers": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "phone": "8652523958",
          "country_code": 91
        }
      ],
      "emails": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "email": "akashmane@gofynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@fynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@uniket.store"
        }
      ],
      "created_at": "2020-03-11T09:28:41.982Z",
      "updated_at": "2021-02-04T10:10:44.981Z"
    }
  }
}
```









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
  "value": {
    "user": {
      "debug": {
        "source": "deadlock",
        "platform": "000000000000000000000001"
      },
      "gender": "male",
      "account_type": "user",
      "active": true,
      "profile_pic_url": "https://d2co8r51m5ca2d.cloudfront.net/inapp_banners/default_profile_img.png",
      "has_old_password_hash": false,
      "_id": "5e68af49cfa09bf7233022f1",
      "first_name": "Akash",
      "last_name": "Mane",
      "username": "akashmane_gofynd_com_10039",
      "phone_numbers": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "phone": "8652523958",
          "country_code": 91
        }
      ],
      "emails": [
        {
          "active": true,
          "primary": true,
          "verified": true,
          "email": "akashmane@gofynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@fynd.com"
        },
        {
          "active": true,
          "primary": false,
          "verified": true,
          "email": "akashmane@uniket.store"
        }
      ],
      "created_at": "2020-03-11T09:28:41.982Z",
      "updated_at": "2021-02-04T10:10:44.981Z"
    }
  }
}
```









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
  "value": {
    "announcements": {
      "$all": [
        {
          "announcement": "<link rel=\"stylesheet\" type=\"text/css\" href=\"https://cdnjs.cloudflare.com/ajax/libs/github-markdown-css/2.9.0/github-markdown.min.css\" /> <div class=\"markdown-body\" style='padding: 12px; color: #000; background-color: #fff;'><p>test Announcement</p>\n</div>",
          "schedule": {
            "start": "2021-03-31T11:22:08.167Z"
          }
        }
      ]
    },
    "refresh_rate": 900,
    "refresh_pages": []
  }
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



Success. Returns a JSON object with blog details. Check the example shown below or refer `BlogSchema` for more details.


Schema: `BlogSchema`


*Examples:*


default
```json
{
  "value": {
    "_id": "5eaa451a21a4dd75f0fd96c5",
    "application": "5d3ebd89f540e7506b8b3548",
    "tags": [
      "abhinav"
    ],
    "title": "my first blog",
    "slug": "1st_blog",
    "feature_image": {
      "secure_url": "https://google.com"
    },
    "content": [
      {
        "type": "html",
        "value": "<p>hey there!</p>"
      }
    ],
    "_schedule": {
      "cron": "* 10 * * *",
      "start": "2021-03-31T23:30:00.000Z",
      "end": "2021-03-31T23:55:00.000Z",
      "duration": 1000,
      "next_schedule": [
        {
          "start": "2021-03-17T04:30:00.000Z",
          "end": "2021-03-17T04:46:40.000Z"
        }
      ]
    },
    "published": true,
    "author": {
      "name": "Fynd App"
    },
    "date_meta": {
      "created_on": "2021-03-14T06:49:03.945Z",
      "modified_on": "2021-03-14T06:49:03.945Z"
    }
  }
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
  "value": {
    "items": [
      {
        "date_meta": {
          "created_on": "2021-03-14T06:49:03.945Z",
          "modified_on": "2021-03-14T06:49:03.945Z"
        },
        "tags": [],
        "_id": "604db275b3ae202873964d94",
        "content": [
          {
            "type": "html",
            "value": "<p>test abhinav</p>"
          }
        ],
        "title": "1st Blog",
        "slug": "1st-blog",
        "published": true,
        "_schedule": {
          "next_schedule": [
            {}
          ],
          "start": "2021-04-08T07:15:13.000Z",
          "end": "2021-04-10T02:00:00.000Z"
        },
        "feature_image": {
          "secure_url": ""
        },
        "application": "000000000000000000000001",
        "author": {
          "name": "Fynd App"
        }
      }
    ],
    "page": {
      "type": "number",
      "current": 1,
      "size": 1,
      "item_total": 2,
      "has_next": true
    }
  }
}
```









---


#### getDataLoaders
Get the data loaders associated with an application

```golang

 data, err :=  Content.GetDataLoaders();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to get all selected data loaders of the application in the form of tags.

*Success Response:*



Success. Returns a JSON object containing all the data loaders injected in the application. Check the example shown below or refer `DataLoadersSchema` for more details.


Schema: `DataLoadersSchema`









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
  "value": {
    "faqs": [
      {
        "_id": "5eb2db750a8ebf497e315028",
        "question": "how to refer my friend",
        "answer": "1. Click on refer and earn image in fynd app\n2. Click on share the code\n3. Use any method for sharing\n4. Once the user activates the app with your code, both of you will get the refereal credits.",
        "slug": "how to refer",
        "application": "000000000000000000000001"
      }
    ]
  }
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
  "value": {
    "_id": "5eaa451a21a4dd75f0fd96c5",
    "application": "5d3ebd89f540e7506b8b3548",
    "_custom_json": null,
    "slug": "pnc-landing",
    "action": {
      "page": {
        "type": "home"
      },
      "popup": {},
      "type": "page"
    },
    "platform": [
      "web"
    ],
    "created_by": {
      "id": "000000000000000000000000"
    },
    "date_meta": {
      "created_on": "2020-04-30T03:25:14.549Z",
      "modified_on": "2020-04-30T03:25:14.549Z"
    },
    "archived": false
  }
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
  "value": {
    "tnc": "**Terms and Conditions test**",
    "policy": "**Privacy policy test**",
    "shipping": "**Shipping term and conditions**",
    "returns": "**Terms & conditions for returns **",
    "_id": "5e8b2b96abe7dc94c02c9ac9",
    "application": "000000000000000000000001",
    "faq": [
      {
        "question": "New Question",
        "answer": "New Answer"
      },
      {
        "question": "New",
        "answer": "sdfghjhg"
      },
      {
        "question": "test",
        "answer": "test"
      },
      {
        "question": "New Test",
        "answer": "New Test answer"
      },
      {
        "question": "test",
        "answer": "test"
      }
    ],
    "created_at": "2020-04-06T13:16:06.818Z",
    "updated_at": "2020-07-16T09:47:40.751Z",
    "__v": 260
  }
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
  "value": {
    "items": [
      {
        "_id": "5ffbd9b90ac98678ae0458d7",
        "application": "000000000000000000000001",
        "_custom_json": null,
        "name": "temp",
        "slug": "temp",
        "platform": "web",
        "position": "top",
        "orientation": "landscape",
        "navigation": [
          {
            "display": "Home",
            "image": "https://res.cloudinary.com/dwzm9bysq/image/upload/v1567148153/production/system/icons/mystore-tab_y0dqzt.png",
            "sort_order": 1,
            "type": "",
            "action": {
              "page": {
                "url": "/",
                "type": "home"
              },
              "popup": {},
              "type": "page"
            },
            "active": true,
            "tags": {
              "type": "array",
              "items": {
                "type": "string"
              }
            },
            "acl": [
              "all"
            ],
            "_locale_language": {
              "hi": {
                "display": ""
              },
              "ar": {
                "display": ""
              },
              "en_us": {
                "display": ""
              }
            },
            "sub_navigation": [
              {
                "display": "Brands",
                "image": "https://res.cloudinary.com/dwzm9bysq/image/upload/v1567148153/production/system/icons/brands-tab_sfinpk.png",
                "sort_order": 1,
                "type": "",
                "action": {
                  "page": {
                    "url": "/brands/",
                    "type": "brands"
                  },
                  "popup": {},
                  "type": "page"
                },
                "active": true,
                "tags": null,
                "acl": [
                  "all"
                ],
                "_locale_language": {
                  "hi": {
                    "display": ""
                  },
                  "ar": {
                    "display": ""
                  },
                  "en_us": {
                    "display": ""
                  }
                }
              }
            ]
          },
          {
            "display": "Collections",
            "image": "https://res.cloudinary.com/dwzm9bysq/image/upload/v1567148153/production/system/icons/collections-tab_a0tg9c.png",
            "sort_order": 2,
            "type": "",
            "action": {
              "page": {
                "url": "/collections/",
                "type": "collections"
              },
              "popup": {},
              "type": "page"
            },
            "active": true,
            "tags": null,
            "acl": [
              "all"
            ],
            "_locale_language": {
              "hi": {
                "display": ""
              },
              "ar": {
                "display": ""
              },
              "en_us": {
                "display": ""
              }
            },
            "sub_navigation": [
              {
                "display": "Categories",
                "image": "https://res.cloudinary.com/dwzm9bysq/image/upload/v1567148154/production/system/icons/categories-tab_ss8e0q.png",
                "sort_order": 1,
                "type": "",
                "action": {
                  "page": {
                    "url": "/categories/",
                    "type": "categories"
                  },
                  "popup": {},
                  "type": "page"
                },
                "active": true,
                "tags": null,
                "acl": [
                  "all"
                ],
                "_locale_language": {
                  "hi": {
                    "display": ""
                  },
                  "ar": {
                    "display": ""
                  },
                  "en_us": {
                    "display": ""
                  }
                }
              }
            ]
          },
          {
            "display": "Primary Menu",
            "image": "",
            "sort_order": 3,
            "type": "",
            "action": {
              "page": {
                "type": "home"
              },
              "popup": {},
              "type": "page"
            },
            "active": true,
            "tags": null,
            "acl": [
              "all"
            ],
            "_locale_language": {
              "hi": {
                "display": ""
              },
              "ar": {
                "display": ""
              },
              "en_us": {
                "display": ""
              }
            }
          }
        ],
        "created_by": {
          "id": "000000000000000000000000"
        },
        "date_meta": {
          "created_on": "2021-01-11T04:53:13.585Z",
          "modified_on": "2021-01-14T10:24:34.485Z"
        }
      }
    ],
    "page": {
      "type": "number",
      "current": 1,
      "size": 1,
      "item_total": 2,
      "has_next": true
    }
  }
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
  "value": {
    "seo": {
      "details": {
        "title": "Zyosa Zyosa",
        "description": "",
        "image_url": ""
      },
      "robots_txt": "User-agent: * \nAllow: / \nsancisciasn xwsaixjowqnxwsiwjs",
      "sitemap_enabled": false,
      "cannonical_enabled": false,
      "_id": "6009819ee463ad40de397eb2",
      "app": "000000000000000000000001",
      "created_at": "2021-01-21T13:29:02.543Z",
      "updated_at": "2021-02-05T06:36:16.048Z",
      "__v": 11,
      "custom_meta_tags": [
        {
          "name": "test 0000",
          "content": "<meta name=\"test\" content=\"0000 cn dcje dcj rejre cjrenurenc \">",
          "_id": "6017c301bde3c21dbb13b284"
        },
        {
          "name": "cwdcdc",
          "content": "<meta content=\"wdcewdewc\">",
          "_id": "6017c675bde3c22cfb13b290"
        }
      ]
    }
  }
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
  "value": {
    "items": [
      {
        "date_meta": {
          "created_on": "2021-03-14T05:27:12.319Z",
          "modified_on": "2021-03-14T05:27:12.319Z"
        },
        "archived": false,
        "_id": "604d9eb975e9d136bb1b8b83",
        "configuration": {
          "start_on_launch": false,
          "duration": 50,
          "sleep_time": 100,
          "slide_direction": "horizontal"
        },
        "slug": "ss-sfsd-updated",
        "platform": "ios",
        "media": [
          {
            "auto_decide_duration": false,
            "type": "image",
            "url": "https://res.cloudinary.com/dwzm9bysq/image/upload/v1567148153/production/system/icons/brands-tab_sfinpk.png",
            "bg_color": "#ffffff",
            "duration": 10,
            "action": {
              "type": ""
            }
          },
          {
            "auto_decide_duration": true,
            "type": "youtube",
            "url": "https://www.youtube.com/embed/9vJRopau0g0",
            "bg_color": "#ffffff",
            "duration": 909,
            "action": {
              "type": ""
            }
          }
        ],
        "application": "5cd3db5e9d692cfe5302a7bb",
        "active": true,
        "__v": 0
      }
    ],
    "page": {
      "type": "number",
      "current": 1,
      "size": 1,
      "item_total": 2,
      "has_next": true
    }
  }
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
  "value": {
    "date_meta": {
      "created_on": "2021-03-14T05:27:12.319Z",
      "modified_on": "2021-03-14T05:27:12.319Z"
    },
    "archived": false,
    "_id": "604d9eb975e9d136bb1b8b83",
    "configuration": {
      "start_on_launch": false,
      "duration": 50,
      "sleep_time": 100,
      "slide_direction": "horizontal"
    },
    "slug": "ss-sfsd-updated",
    "platform": "ios",
    "media": [
      {
        "auto_decide_duration": false,
        "type": "image",
        "url": "https://res.cloudinary.com/dwzm9bysq/image/upload/v1567148153/production/system/icons/brands-tab_sfinpk.png",
        "bg_color": "#ffffff",
        "duration": 10,
        "action": {
          "type": ""
        }
      },
      {
        "auto_decide_duration": true,
        "type": "youtube",
        "url": "https://www.youtube.com/embed/9vJRopau0g0",
        "bg_color": "#ffffff",
        "duration": 909,
        "action": {
          "type": ""
        }
      }
    ],
    "application": "5cd3db5e9d692cfe5302a7bb",
    "active": true
  }
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
  "value": {
    "_id": "5ea4980b87a7944094216193",
    "config_type": "app",
    "application": "000000000000000000000001",
    "created_at": "2020-04-25T20:05:31.300Z",
    "updated_at": "2020-12-04T10:48:12.194Z",
    "contact": {
      "phone": {
        "active": true,
        "phone": [
          {
            "key": "Jane Doe",
            "code": "91",
            "number": "9988776655"
          }
        ]
      },
      "email": {
        "active": false,
        "email": []
      }
    },
    "created": true
  }
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









---


#### getPage
Get a page

```golang

 data, err :=  Content.GetPage(Slug, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Slug | string | A short, human-readable, URL-friendly identifier of a page. You can get slug value from the endpoint /service/application/content/v2.0/pages/. | 



| xQuery | struct | Includes properties such as `RootID`



Use this API to get the details of a page using its slug. Details include the title, seo, publish status, feature image, tags, meta, etc.

*Success Response:*



Success. Returns a JSON object with page details. Check the example shown below or refer `CustomPageSchema` for more details.


Schema: `PageSchema`


*Examples:*


default
```json
{
  "value": {
    "date_meta": {
      "created_on": "2021-03-16T08:24:19.197Z",
      "modified_on": "2021-03-16T08:24:19.197Z"
    },
    "tags": [
      "my first page"
    ],
    "published": true,
    "component_ids": [],
    "archived": false,
    "_id": "60506dcad18cb33946026862",
    "title": "my first page",
    "slug": "1st_page",
    "feature_image": {
      "secure_url": "https://google.com/some-image"
    },
    "content_path": "https://hdn-1.fynd.com/company/1526/applications/61012f6a9250ccd1b9ef8a1d/pages/content/page_slug.html",
    "platform": "web",
    "description": "hey this is my first page",
    "visibility": {
      "test": true
    },
    "_schedule": {
      "start": "2021-04-23T23:50:00.000Z",
      "next_schedule": [
        {}
      ]
    },
    "seo": {
      "title": "my first page",
      "description": "hey this is my first page",
      "image": {
        "url": ""
      }
    },
    "type": "rawhtml",
    "application": "000000000000000000000001",
    "orientation": "portrait",
    "page_meta": []
  }
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



Success. Returns a list of pages along with their details. Check the example shown below or refer `PageGetStorefrontResponse` for more details.


Schema: `PageGetResponse`


*Examples:*


default
```json
{
  "value": {
    "items": [
      {
        "date_meta": {
          "created_on": "2021-03-14T06:49:03.945Z",
          "modified_on": "2021-03-14T06:49:03.945Z"
        },
        "tags": [
          "my first page"
        ],
        "_id": "604db275b3ae202873964d94",
        "content_path": "https://hdn-1.fynd.com/company/1526/applications/61012f6a9250ccd1b9ef8a1d/pages/content/page_slug.html",
        "title": "test-page",
        "slug": "test-page",
        "published": true,
        "_schedule": {
          "next_schedule": [
            {}
          ],
          "start": "2021-04-08T07:15:13.000Z",
          "end": "2021-04-10T02:00:00.000Z"
        },
        "feature_image": {
          "secure_url": "https://google.com/some-image"
        },
        "seo": {
          "title": "my first page",
          "description": "hey this is my first page",
          "image": {
            "url": ""
          }
        },
        "application": "000000000000000000000001",
        "author": {
          "name": "Abhinav Maurya"
        }
      }
    ],
    "page": {
      "type": "number",
      "current": 1,
      "size": 1,
      "item_total": 2,
      "has_next": true
    }
  }
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



Use this API to retrieve the consent provided by the user for receiving communication messages over Email/SMS/WhatsApp.

*Success Response:*



Success. Returns all available communication opt-ins along with the consent details. Check the example shown below or refer `CommunicationConsent` for more details.


Schema: `CommunicationConsent`


*Examples:*


default
```json
{
  "value": {
    "app_id": "000000000000000000000004",
    "user_id": "5e56021c4bda3ccab6d9f884",
    "channels": {
      "email": {
        "response": "yes",
        "display_name": "Email"
      },
      "sms": {
        "response": "yes",
        "display_name": "SMS"
      },
      "whatsapp": {
        "response": "yes",
        "display_name": "WhatsApp",
        "country_code": "91",
        "phone_number": "9869821300"
      }
    }
  }
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


Use this API to update and insert the consent provided by the user for receiving communication messages over Email/SMS/WhatsApp.

*Success Response:*



Success. Updates the channels for which user has consented. Check the example shown below or refer `CommunicationConsentRes` for more details.


Schema: `CommunicationConsentRes`


*Examples:*


default
```json
{
  "value": {
    "appId": "000000000000000000000004",
    "userId": "5e56021c4bda3ccab6d9f884",
    "channels": {
      "email": {
        "response": "yes",
        "displayName": "Email"
      },
      "sms": {
        "response": "yes",
        "displayName": "SMS"
      },
      "whatsapp": {
        "response": "noaction",
        "displayName": "WhatsApp"
      }
    }
  }
}
```









---


#### upsertAppPushtoken
Upsert push token of a user

```golang

 data, err :=  Communication.UpsertAppPushtoken(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  PushtokenReq | "Request body" 


Use this API to update and insert the push token of the user.

*Success Response:*



Success. Check the example shown below or refer `PushtokenRes` for more details.


Schema: `PushtokenRes`


*Examples:*


create
```json
{
  "value": {
    "_id": "601b6924d8ea9a061570a09f",
    "bundle_identifier": "000002",
    "push_token": "45",
    "unique_device_id": "3",
    "type": "fynd-platform",
    "platform": "web",
    "application_id": "000000000000000000000004",
    "user_id": "5e56021c4bda3ccab6d9f884",
    "created_at": "2021-02-04T03:25:24.765Z",
    "updated_at": "2021-02-04T03:25:51.152Z"
  }
}
```

update
```json
{
  "value": {
    "_id": "601b6924d8ea9a061570a09f",
    "bundle_identifier": "000002",
    "push_token": "45",
    "unique_device_id": "3",
    "type": "fynd-platform",
    "platform": "web",
    "application_id": "000000000000000000000004",
    "user_id": "5e56021c4bda3ccab6d9f884",
    "created_at": "2021-02-04T03:25:24.765Z",
    "updated_at": "2021-02-04T03:25:51.152Z"
  }
}
```

reset
```json
{
  "value": {
    "_id": "601b6924d8ea9a061570a09f",
    "bundle_identifier": "000002",
    "push_token": "45",
    "unique_device_id": "3",
    "type": "fynd-platform",
    "platform": "web",
    "application_id": "000000000000000000000004",
    "user_id": "5e56021c4bda3ccab6d9f884",
    "created_at": "2021-02-04T03:25:24.765Z",
    "updated_at": "2021-02-04T03:25:51.152Z",
    "expired_at": "2021-02-05T03:25:51.138Z"
  }
}
```









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









---



---


## FileStorage


#### startUpload
Initiates an upload and returns a storage link that is valid for 30 minutes. You can use the storage link to make subsequent upload request with file buffer or blob.

```golang

 data, err :=  FileStorage.StartUpload(Namespace, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Namespace | string | Segregation of different types of files(products, orders, logistics etc), Required for validating the data of the file being uploaded, decides where exactly the file will be stored inside the storage bucket. | 


| body |  StartRequest | "Request body" 


Use this API to perform the first step of uploading (i.e. **Start**) an arbitrarily sized buffer or blob.

The three major steps are:
* Start
* Upload
* Complete

### Start
Initiates the assets upload using `startUpload`.
It returns a storage link in response.

### Upload
Use the storage link to upload a file (Buffer or Blob) to the File Storage.
Make a `PUT` request on storage link received from `startUpload` API with the file (Buffer or Blob) in the request body.

### Complete
After successfully upload, call the `completeUpload` API to finish the upload process.
This operation will return the URL of the uploaded file.


*Success Response:*



Success. Next, call the `completeUpload` API and pass the response payload of this API to finish the upload process.


Schema: `StartResponse`









---


#### completeUpload
Completes the upload process. After successfully uploading a file, call this API to finish the upload process.

```golang

 data, err :=  FileStorage.CompleteUpload(Namespace, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Namespace | string | Segregation of different types of files(products, orders, logistics etc), Required for validating the data of the file being uploaded, decides where exactly the file will be stored inside the storage bucket. | 


| body |  StartResponse | "Request body" 


Use this API to perform the third step of uploading (i.e. **Complete**) an arbitrarily sized buffer or blob.

The three major steps are:
* Start
* Upload
* Complete

### Start
Initiates the assets upload using `startUpload`.
It returns a storage link in response.

### Upload
Use the storage link to upload a file (Buffer or Blob) to the File Storage.
Make a `PUT` request on storage link received from `startUpload` API with the file (Buffer or Blob) in the request body.

### Complete
After successfully upload, call the `completeUpload` API to finish the upload process.
This operation will return the URL of the uploaded file.


*Success Response:*



Success


Schema: `CompleteResponse`









---


#### signUrls
Explain here

```golang

 data, err :=  FileStorage.SignUrls(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  SignUrlRequest | "Request body" 


Describe here

*Success Response:*



Success


Schema: `SignUrlResponse`









---



---


## Configuration


#### getApplication
Get current sales channel details

```golang

 data, err :=  Configuration.GetApplication();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to get the current sales channel details which includes configurations that indicate the status of the website, domain, ID, tokens, images, etc.

*Success Response:*



Success. Check the example shown below or refer `Application` for more details.


Schema: `Application`









---


#### getOwnerInfo
Get sales channel, owner and seller information

```golang

 data, err :=  Configuration.GetOwnerInfo();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to get the current sales channel details which includes channel name, description, banner, logo, favicon, domain details, etc. This API also retrieves the seller and owner information such as address, email address, and phone number.

*Success Response:*



Success. Check the example shown below or refer `ApplicationAboutResponse` for more details.


Schema: `ApplicationAboutResponse`









---


#### getBasicDetails
Get basic details of the application

```golang

 data, err :=  Configuration.GetBasicDetails();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to retrieve only the basic details of the application which includes channel name, description, banner, logo, favicon, domain details, etc.

*Success Response:*



Success. Check the example shown below or refer `ApplicationDetail` for more details.


Schema: `ApplicationDetail`









---


#### getIntegrationTokens
Get integration tokens

```golang

 data, err :=  Configuration.GetIntegrationTokens();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to retrieve the tokens used while integrating Firebase, MoEngage, Segment, GTM, Freshchat, Safetynet, Google Map and Facebook. **Note** - Token values are encrypted with AES encryption using a secret key. Kindly reach out to the developers for obtaining the secret key.

*Success Response:*



Success. Check the example shown below or refer `AppTokenResponse` for more details.


Schema: `AppTokenResponse`









---


#### getOrderingStores
Get all deployment stores

```golang

 data, err :=  Configuration.GetOrderingStores(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |






| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Q`



Use this API to retrieve the details of all the deployment stores (the selling locations where the application will be utilized for placing orders).

*Success Response:*



Success. Check the example shown below or refer `OrderingStores` for more details.


Schema: `OrderingStores`









---


#### getStoreDetailById
Get ordering store details

```golang

 data, err :=  Configuration.GetStoreDetailById(StoreID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| StoreID | float64 | Store uid | 




Use this API to retrieve the details of given stores uid (the selling locations where the application will be utilized for placing orders).

*Success Response:*



Success. Check the example shown below or refer `OrderingStore` for more details.


Schema: `OrderingStore`









---


#### getFeatures
Get features of application

```golang

 data, err :=  Configuration.GetFeatures();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to retrieve the configuration of features such as product detail, landing page, options in the login/registration screen, communication opt-in, cart options and many more.

*Success Response:*



Success. Check the example shown below or refer `AppFeatureResponse` for more details.


Schema: `AppFeatureResponse`









---


#### getContactInfo
Get application information

```golang

 data, err :=  Configuration.GetContactInfo();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to retrieve information about the social links, address and contact information of the company/seller/brand operating the application.

*Success Response:*



Success. Check the example shown below or refer `ApplicationAboutResponse` for more details.


Schema: `ApplicationInformation`









---


#### getCurrencies
Get all currencies list

```golang

 data, err :=  Configuration.GetCurrencies();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to get a list of currencies available. Moreover, get the name, code, symbol, and the decimal digits of the currencies.

*Success Response:*



Success. Check the example shown below or refer `CurrenciesResponse` for more details.


Schema: `CurrenciesResponse`









---


#### getCurrencyById
Get currency by its ID

```golang

 data, err :=  Configuration.GetCurrencyById(ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ID | string | Object ID assigned to the currency | 




Use this API to retrieve a currency using its ID.

*Success Response:*



Success. Check the example shown below or refer `Currency` for more details.


Schema: `Currency`









---


#### getAppCurrencies
Get currencies enabled in the application

```golang

 data, err :=  Configuration.GetAppCurrencies();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to get a list of currencies allowed in the current application. Moreover, get the name, code, symbol, and the decimal digits of the currencies.

*Success Response:*



Success. Check the example shown below or refer `AppCurrencyResponse` for more details.


Schema: `AppCurrencyResponse`









---


#### getLanguages
Get list of languages

```golang

 data, err :=  Configuration.GetLanguages();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to get a list of languages supported in the application

*Success Response:*



Success. Check the example shown below or refer `LanguageResponse` for more details.


Schema: `LanguageResponse`









---


#### getOrderingStoreCookie
Get an Ordering Store signed cookie on selection of ordering store.

```golang

 data, err :=  Configuration.GetOrderingStoreCookie(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  OrderingStoreSelectRequest | "Request body" 


Use this API to get an Ordering Store signed cookie upon selecting an ordering store. This will be used by the cart service to verify a coupon against the selected ordering store in cart.

*Success Response:*



Success


Schema: `SuccessMessageResponse`









---


#### removeOrderingStoreCookie
Unset the Ordering Store signed cookie.

```golang

 data, err :=  Configuration.RemoveOrderingStoreCookie();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to unset the Ordering Store cookie upon changing the sales channel, by its domain URL, in the Universal Fynd Store app.

*Success Response:*



Success


Schema: `SuccessMessageResponse`









---


#### getAppStaffList
Get a list of staff.

```golang

 data, err :=  Configuration.GetAppStaffList(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |










| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `OrderIncent`, `OrderingStore`, `User`



Use this API to get a list of staff including the names, employee code, incentive status, assigned ordering stores, and title of each staff added to the application.

*Success Response:*



Success. Check the example shown below or refer `AppStaffListResponse` for more details.


Schema: `AppStaffListResponse`









---


#### getAppStaffs
Get a list of staff.

```golang

 data, err :=  Configuration.GetAppStaffs(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |






| xQuery | struct | Includes properties such as `OrderIncent`, `OrderingStore`, `User`



Use this API to get a list of staff including the names, employee code, incentive status, assigned ordering stores, and title of each staff added to the application.

*Success Response:*



Success. Check the example shown below or refer `AppStaffResponse` for more details.


Schema: `AppStaffResponse`









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

| XAPIToken | string | Used for basic authentication. | 



| xQuery | struct | Includes properties such as `Refresh`



Use this API to retrieve the payment gateway key, secrets, merchant, SDK/API details to complete a payment at front-end.

*Success Response:*



Success. Returns the keys of all payment gateways. Check the example shown below or refer `AggregatorsConfigDetailResponse` for more details.


Schema: `AggregatorsConfigDetailResponse`









---


#### attachCardToCustomer
Attach a saved card to customer.

```golang

 data, err :=  Payment.AttachCardToCustomer(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  AttachCardRequest | "Request body" 


Use this API to attach a customer's saved card at the payment gateway, such as Stripe, Juspay.

*Success Response:*



Success. Check the example shown below or refer `AttachCardsResponse` for more details.


Schema: `AttachCardsResponse`









---


#### getActiveCardAggregator
Fetch active payment gateway for card payments

```golang

 data, err :=  Payment.GetActiveCardAggregator(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Refresh`



Use this API to retrieve an active payment aggregator along with the Customer ID. This is applicable for cards payments only.

*Success Response:*



Success. Returns an active payment gateway. Check the example shown below or refer `ActiveCardPaymentGatewayResponse` for more details.


Schema: `ActiveCardPaymentGatewayResponse`









---


#### getActiveUserCards
Fetch the list of cards saved by the user

```golang

 data, err :=  Payment.GetActiveUserCards(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `ForceRefresh`



Use this API to retrieve a list of cards stored by user from an active payment gateway.

*Success Response:*



Success. Returns a list of cards saved by the user. Check the example shown below or refer `ListCardsResponse` for more details.


Schema: `ListCardsResponse`









---


#### deleteUserCard
Delete a card

```golang

 data, err :=  Payment.DeleteUserCard(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  DeletehCardRequest | "Request body" 


Use this API to delete a card added by a user on the payment gateway and clear the cache.

*Success Response:*



Success. Returns a success message if card is deleted.


Schema: `DeleteCardsResponse`









---


#### verifyCustomerForPayment
Validate customer for payment

```golang

 data, err :=  Payment.VerifyCustomerForPayment(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  ValidateCustomerRequest | "Request body" 


Use this API to check if the customer is eligible to use credit-line facilities such as Simpl Pay Later and Rupifi.

*Success Response:*



Success. Check the example shown below or refer `ValidateCustomerResponse` for more details.


Schema: `ValidateCustomerResponse`


*Examples:*


success is True i.e user is allowed
```json
{
  "value": {
    "success": true,
    "message": "data fetched",
    "data": {
      "api_version": 2,
      "data": {
        "approved": true,
        "button_text": "Buy Now, Pay Later",
        "first_transaction": false
      },
      "aggregator": "Simpl"
    }
  }
}
```

success is True i.e user not allowed
```json
{
  "value": {
    "success": false,
    "message": "data fetched",
    "error": {
      "api_version": 2,
      "data": {
        "approved": false,
        "button_text": "Buy Now, Pay Later",
        "first_transaction": false
      },
      "aggregator": "Simpl"
    },
    "data": {}
  }
}
```









---


#### verifyAndChargePayment
Verify and charge payment

```golang

 data, err :=  Payment.VerifyAndChargePayment(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  ChargeCustomerRequest | "Request body" 


Use this API to verify and check the status of a payment transaction (server-to-server) made through aggregators like Simpl and Mswipe.

*Success Response:*



Success. Check the example shown below or refer `ChargeCustomerResponse` for more details.


Schema: `ChargeCustomerResponse`









---


#### initialisePayment
Initialize a payment (server-to-server) for UPI and BharatQR

```golang

 data, err :=  Payment.InitialisePayment(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  PaymentInitializationRequest | "Request body" 


PUse this API to inititate payment using UPI, BharatQR, wherein the UPI requests are send to the app and QR code is displayed on the screen.

*Success Response:*



Success. Check the example shown below or refer `PaymentInitializationResponse` for more details.


Schema: `PaymentInitializationResponse`









---


#### checkAndUpdatePaymentStatus
Performs continuous polling to check status of payment on the server

```golang

 data, err :=  Payment.CheckAndUpdatePaymentStatus(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  PaymentStatusUpdateRequest | "Request body" 


Use this API to perform continuous polling at intervals to check the status of payment until timeout.

*Success Response:*



Success. Returns the status of payment. Check the example shown below or refer `PaymentStatusUpdateResponse` for more details.


Schema: `PaymentStatusUpdateResponse`









---


#### getPaymentModeRoutes
Get applicable payment options

```golang

 data, err :=  Payment.GetPaymentModeRoutes(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |














| xQuery | struct | Includes properties such as `Amount`, `CartID`, `Pincode`, `CheckoutMode`, `Refresh`, `CardReference`, `UserDetails`



Use this API to get all valid payment options for doing a payment.

*Success Response:*



Success. Returns all available options for payment. Check the example shown below or refer `PaymentModeRouteResponse` for more details.


Schema: `PaymentModeRouteResponse`









---


#### getPosPaymentModeRoutes
Get applicable payment options for Point-of-Sale (POS)

```golang

 data, err :=  Payment.GetPosPaymentModeRoutes(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |
















| xQuery | struct | Includes properties such as `Amount`, `CartID`, `Pincode`, `CheckoutMode`, `Refresh`, `CardReference`, `OrderType`, `UserDetails`



Use this API to get all valid payment options for doing a payment in POS.

*Success Response:*



Success. Returns all available options for payment. Check the example shown below or refer `PaymentModeRouteResponse` for more details.


Schema: `PaymentModeRouteResponse`









---


#### getRupifiBannerDetails
Get CreditLine Offer

```golang

 data, err :=  Payment.GetRupifiBannerDetails();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Get CreditLine Offer if user is tentatively approved by rupifi

*Success Response:*



Success. Return CreditLine Offer detail. Check the example shown below or refer `RupifiBannerResponseSchema` for more details.


Schema: `RupifiBannerResponse`









---


#### getEpaylaterBannerDetails
Get Epaylater Enabled

```golang

 data, err :=  Payment.GetEpaylaterBannerDetails();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Get Epaylater Enabled if user is tentatively approved by epaylater

*Success Response:*



Success. Return Epaylater Offer detail. Check the example shown below or refer `EpaylaterBannerResponseSchema` for more details. if `display=True`, then show banner otherwise do not show.


Schema: `EpaylaterBannerResponse`


*Examples:*


User is registered successfully
```json
{
  "value": {
    "success": true,
    "data": {
      "display": false,
      "message": "User is Active",
      "status": "ACTIVE"
    }
  }
}
```

User is not registered or KYC not done or approval pending
```json
{
  "value": {
    "success": true,
    "data": {
      "display": true,
      "message": "User is not registered",
      "status": "NOT REGISTERED"
    }
  }
}
```









---


#### resendOrCancelPayment
API to resend and cancel a payment link which was already generated.

```golang

 data, err :=  Payment.ResendOrCancelPayment(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  ResendOrCancelPaymentRequest | "Request body" 


Use this API to perform resend or cancel a payment link based on request payload.

*Success Response:*



Success. Returns the status of payment. Check the example shown below or refer `ResendOrCancelPaymentResponse` for more details.


Schema: `ResendOrCancelPaymentResponse`


*Examples:*


request_type is cancel
```json
{
  "value": {
    "success": true,
    "data": {
      "message": "Payment link Cancelled.",
      "status": true
    }
  }
}
```

request_type is resend
```json
{
  "value": {
    "success": true,
    "data": {
      "message": "Notification triggered.",
      "status": true
    }
  }
}
```









---


#### renderHTML
Convert base64 string to HTML form

```golang

 data, err :=  Payment.RenderHTML(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  renderHTMLRequest | "Request body" 


Use this API to decode base64 html form to plain HTML string.

*Success Response:*



Success and return HTML decoded text


Schema: `renderHTMLResponse`









---


#### validateVPA
API to Validate UPI ID

```golang

 data, err :=  Payment.ValidateVPA(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  ValidateVPARequest | "Request body" 


API to Validate UPI ID

*Success Response:*



Success. Returns the status of payment. Check the example shown below or refer `ValidateVPAResponseSchema` for more details.


Schema: `ValidateVPAResponse`









---


#### getActiveRefundTransferModes
Lists the mode of refund

```golang

 data, err :=  Payment.GetActiveRefundTransferModes();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to retrieve eligible refund modes (such as Netbanking) and add the beneficiary details.

*Success Response:*



Success. Shows the available refund mode to choose, e.g. Netbanking. Check the example shown below or refer `TransferModeResponse` for more details.


Schema: `TransferModeResponse`









---


#### enableOrDisableRefundTransferMode
Enable/Disable a mode for transferring a refund

```golang

 data, err :=  Payment.EnableOrDisableRefundTransferMode(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  UpdateRefundTransferModeRequest | "Request body" 


Activate or Deactivate Transfer Mode to collect Beneficiary Details for Refund

*Success Response:*



Success. Shows whether the refund mode was successfully enabled or disabled.


Schema: `UpdateRefundTransferModeResponse`









---


#### getUserBeneficiariesDetail
Lists the beneficiary of a refund

```golang

 data, err :=  Payment.GetUserBeneficiariesDetail(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `OrderID`



Use this API to get the details of all active beneficiary added by a user for refund.

*Success Response:*



Success. Returns the details of the beneficiary getting a refund. Check the example shown below or refer `OrderBeneficiaryResponse` for more details.


Schema: `OrderBeneficiaryResponse`









---


#### verifyIfscCode
Verify IFSC Code

```golang

 data, err :=  Payment.VerifyIfscCode(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `IfscCode`



Use this API to check whether the 11-digit IFSC code is valid and to fetch the bank details for refund.

*Success Response:*



Success. Shows whether the IFSC code is valid, and returns the bank details. Check the example shown below or refer `IfscCodeResponse` for more details.


Schema: `IfscCodeResponse`









---


#### getOrderBeneficiariesDetail
Lists the beneficiary of a refund

```golang

 data, err :=  Payment.GetOrderBeneficiariesDetail(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `OrderID`



Use this API to get the details of all active beneficiary added by a user for refund.

*Success Response:*



Success. Returns the details of the beneficiary getting a refund. Check the example shown below or refer `OrderBeneficiaryResponse` for more details.


Schema: `OrderBeneficiaryResponse`









---


#### verifyOtpAndAddBeneficiaryForBank
Verify the beneficiary details using OTP

```golang

 data, err :=  Payment.VerifyOtpAndAddBeneficiaryForBank(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  AddBeneficiaryViaOtpVerificationRequest | "Request body" 


Use this API to perform an OTP validation before saving the beneficiary details added for a refund.

*Success Response:*



Success. Check the example shown below or refer `AddBeneficiaryViaOtpVerificationRequest` for more details.


Schema: `AddBeneficiaryViaOtpVerificationResponse`









---


#### addBeneficiaryDetails
Save bank details for cancelled/returned order

```golang

 data, err :=  Payment.AddBeneficiaryDetails(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  AddBeneficiaryDetailsRequest | "Request body" 


Use this API to save the bank details for a returned or cancelled order to refund the amount.

*Success Response:*



Success. Shows whether the beneficiary details were saved to a returned/cancelled order or not.


Schema: `RefundAccountResponse`









---


#### addRefundBankAccountUsingOTP
Save bank details for cancelled/returned order

```golang

 data, err :=  Payment.AddRefundBankAccountUsingOTP(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  AddBeneficiaryDetailsOTPRequest | "Request body" 


Use this API to save bank details for returned/cancelled order to refund amount in his account.

*Success Response:*



Success. Shows whether the beneficiary details were saved to a returned/cancelled order or not.


Schema: `RefundAccountResponse`









---


#### verifyOtpAndAddBeneficiaryForWallet
Send OTP on adding a wallet beneficiary

```golang

 data, err :=  Payment.VerifyOtpAndAddBeneficiaryForWallet(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  WalletOtpRequest | "Request body" 


Use this API to send an OTP while adding a wallet beneficiary by mobile no. verification.

*Success Response:*



Success. Sends the OTP to the given mobile number. Check the example shown below or refer `WalletOtpResponse` for more details.


Schema: `WalletOtpResponse`









---


#### updateDefaultBeneficiary
Set a default beneficiary for a refund

```golang

 data, err :=  Payment.UpdateDefaultBeneficiary(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  SetDefaultBeneficiaryRequest | "Request body" 


Use this API to set a default beneficiary for getting a refund.

*Success Response:*



Success. Check the example shown below or refer `SetDefaultBeneficiaryResponse` for more details.


Schema: `SetDefaultBeneficiaryResponse`









---


#### getPaymentLink
Get payment link

```golang

 data, err :=  Payment.GetPaymentLink(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `PaymentLinkID`



Use this API to get a payment link

*Success Response:*



Success. Check the example shown below


Schema: `GetPaymentLinkResponse`









---


#### createPaymentLink
Create payment link

```golang

 data, err :=  Payment.CreatePaymentLink(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  CreatePaymentLinkRequest | "Request body" 


Use this API to create a payment link for the customer

*Success Response:*



Success. Check the example shown below


Schema: `CreatePaymentLinkResponse`









---


#### resendPaymentLink
Resend payment link

```golang

 data, err :=  Payment.ResendPaymentLink(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  CancelOrResendPaymentLinkRequest | "Request body" 


Use this API to resend a payment link for the customer

*Success Response:*



Success. Check the example shown below


Schema: `ResendPaymentLinkResponse`









---


#### cancelPaymentLink
Cancel payment link

```golang

 data, err :=  Payment.CancelPaymentLink(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  CancelOrResendPaymentLinkRequest | "Request body" 


Use this API to cancel a payment link for the customer

*Success Response:*



Success. Check the example shown below


Schema: `CancelPaymentLinkResponse`









---


#### getPaymentModeRoutesPaymentLink
Get applicable payment options for payment link

```golang

 data, err :=  Payment.GetPaymentModeRoutesPaymentLink(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `PaymentLinkID`



Use this API to get all valid payment options for doing a payment through payment link

*Success Response:*



Success. Returns all available options for payment. Check the example shown below or refer `PaymentModeRouteResponse` for more details.


Schema: `PaymentModeRouteResponse`









---


#### pollingPaymentLink
Used for polling if payment successful or not

```golang

 data, err :=  Payment.PollingPaymentLink(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `PaymentLinkID`



Use this API to poll if payment through payment was successful or not

*Success Response:*



Success. Check the example shown below


Schema: `PollingPaymentLinkResponse`









---


#### createOrderHandlerPaymentLink
Create Order user

```golang

 data, err :=  Payment.CreateOrderHandlerPaymentLink(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  CreateOrderUserRequest | "Request body" 


Use this API to create a order and payment on aggregator side

*Success Response:*



Success. Check the example shown below


Schema: `CreateOrderUserResponse`









---


#### initialisePaymentPaymentLink
Initialize a payment (server-to-server) for UPI and BharatQR

```golang

 data, err :=  Payment.InitialisePaymentPaymentLink(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  PaymentInitializationRequest | "Request body" 


Use this API to inititate payment using UPI, BharatQR, wherein the UPI requests are send to the app and QR code is displayed on the screen.

*Success Response:*



Success. Check the example shown below or refer `PaymentInitializationResponse` for more details.


Schema: `PaymentInitializationResponse`









---


#### checkAndUpdatePaymentStatusPaymentLink
Performs continuous polling to check status of payment on the server

```golang

 data, err :=  Payment.CheckAndUpdatePaymentStatusPaymentLink(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  PaymentStatusUpdateRequest | "Request body" 


Use this API to perform continuous polling at intervals to check the status of payment until timeout.

*Success Response:*



Success. Returns the status of payment. Check the example shown below or refer `PaymentStatusUpdateResponse` for more details.


Schema: `PaymentStatusUpdateResponse`









---


#### customerCreditSummary
API to fetch the customer credit summary

```golang

 data, err :=  Payment.CustomerCreditSummary(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Aggregator`



Use this API to fetch the customer credit summary.

*Success Response:*



Success. Returns the status of payment. Check the example shown below or refer `CustomerCreditSummaryResponseSchema` for more details.


Schema: `CustomerCreditSummaryResponse`









---


#### redirectToAggregator
API to get the redirect url to redirect the user to aggregator's page

```golang

 data, err :=  Payment.RedirectToAggregator(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `Source`, `Aggregator`



Use this API to get the redirect url to redirect the user to aggregator's page

*Success Response:*



Success. Returns the status of payment. Check the example shown below or refer `RedirectToAggregatorResponseSchema` for more details.


Schema: `RedirectToAggregatorResponse`









---


#### checkCredit
API to fetch the customer credit summary

```golang

 data, err :=  Payment.CheckCredit(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `Aggregator`



Use this API to fetch the customer credit summary.

*Success Response:*



Success. Returns the status of payment. Check the example shown below or refer `CheckCreditResponseSchema` for more details.


Schema: `CheckCreditResponse`









---


#### customerOnboard
API to fetch the customer credit summary

```golang

 data, err :=  Payment.CustomerOnboard(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  CustomerOnboardingRequest | "Request body" 


Use this API to fetch the customer credit summary.

*Success Response:*



Success. Returns the status of payment. Check the example shown below or refer `CustomerOnboardingResponseSchema` for more details.


Schema: `CustomerOnboardingResponse`









---



---


## Order


#### getOrders
Get all orders

```golang

 data, err :=  Order.GetOrders(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |












| xQuery | struct | Includes properties such as `Status`, `PageNo`, `PageSize`, `FromDate`, `ToDate`, `CustomMeta`



Use this API to retrieve all the orders.

*Success Response:*



Success. Returns all the orders. Check the example shown below or refer `OrderList` for more details.


Schema: `OrderList`









---


#### getOrderById
Get details of an order

```golang

 data, err :=  Order.GetOrderById(OrderID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| OrderID | string | A unique number used for identifying and tracking your orders. | 




Use this API to retrieve order details such as tracking details, shipment, store information using Fynd Order ID.

*Success Response:*



Success. Check the example shown below or refer `OrderById` for more details.


Schema: `OrderById`









---


#### getPosOrderById
Get POS Order

```golang

 data, err :=  Order.GetPosOrderById(OrderID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| OrderID | string | A unique number used for identifying and tracking your orders. | 




Use this API to retrieve a POS order and all its details such as tracking details, shipment, store information using Fynd Order ID.

*Success Response:*



Success. Check the example shown below or refer `PosOrderById` for more details.


Schema: `OrderById`









---


#### getShipmentById
Get details of a shipment

```golang

 data, err :=  Order.GetShipmentById(ShipmentID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ShipmentID | string | ID of the shipment. An order may contain multiple items and may get divided into one or more shipment, each having its own ID. | 




Use this API to retrieve shipment details such as price breakup, tracking details, store information, etc. using Shipment ID.

*Success Response:*



Success. Check the example shown below or refer `ShipmentById` for more details.


Schema: `ShipmentById`









---


#### getInvoiceByShipmentId
Get Invoice of a shipment

```golang

 data, err :=  Order.GetInvoiceByShipmentId(ShipmentID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ShipmentID | string | ID of the shipment. | 




Use this API to retrieve shipment invoice.

*Success Response:*



Success. Check the example shown below or refer `ShipmentById` for more details.


Schema: `ResponseGetInvoiceShipment`









---


#### trackShipment
Track shipment

```golang

 data, err :=  Order.TrackShipment(ShipmentID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ShipmentID | string | ID of the shipment. An order may contain multiple items and may get divided into one or more shipment, each having its own ID. | 




Track Shipment by shipment id, for application based on application Id

*Success Response:*



Success. Check the example shown below or refer `ShipmentTrack` for more details.


Schema: `ShipmentTrack`









---


#### getCustomerDetailsByShipmentId
Get Customer Details by Shipment Id

```golang

 data, err :=  Order.GetCustomerDetailsByShipmentId(OrderID, ShipmentID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| OrderID | string | ID of the shipment. An order may contain multiple items and may get divided into one or more shipment, each having its own ID. | 


| ShipmentID | string | A unique number used for identifying and tracking your orders. | 




Use this API to retrieve customer details such as mobileno using Shipment ID.

*Success Response:*



Success. Check the example shown below or refer `CustomerDetailsByShipmentId` for more details.


Schema: `CustomerDetailsResponse`









---


#### sendOtpToShipmentCustomer
Send and Resend Otp code to Order-Shipment customer

```golang

 data, err :=  Order.SendOtpToShipmentCustomer(OrderID, ShipmentID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| OrderID | string | A unique number used for identifying and tracking your orders. | 


| ShipmentID | string | ID of the shipment. An order may contain multiple items and may get divided into one or more shipment, each having its own ID. | 




Use this API to send OTP to the customer of the mapped Shipment.

*Success Response:*



Success to acknowledge the service was notified


Schema: `SendOtpToCustomerResponse`









---


#### verifyOtpShipmentCustomer
Verify Otp code

```golang

 data, err :=  Order.VerifyOtpShipmentCustomer(OrderID, ShipmentID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| OrderID | string | A unique number used for identifying and tracking your orders. | 


| ShipmentID | string | ID of the shipment. An order may contain multiple items and may get divided into one or more shipment, each having its own ID. | 


| body |  VerifyOtp | "Request body" 


Use this API to verify OTP and create a session token with custom payload.

*Success Response:*



Success, the code is valid and returns a session token


Schema: `VerifyOtpResponse`









---


#### getShipmentBagReasons
Get reasons behind full or partial cancellation of a shipment

```golang

 data, err :=  Order.GetShipmentBagReasons(ShipmentID, BagID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ShipmentID | string | ID of the bag. An order may contain multiple items and may get divided into one or more shipment, each having its own ID. | 


| BagID | string | ID of the bag. An order may contain multiple items and may get divided into one or more shipment, each having its own ID. | 




Use this API to retrieve the issues that led to the cancellation of bags within a shipment.

*Success Response:*



Success. Check the example shown below or refer `ShipmentBagReasons` for more details.


Schema: `ShipmentBagReasons`









---


#### getShipmentReasons
Get reasons behind full or partial cancellation of a shipment

```golang

 data, err :=  Order.GetShipmentReasons(ShipmentID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ShipmentID | string | ID of the shipment. An order may contain multiple items and may get divided into one or more shipment, each having its own ID. | 




Use this API to retrieve the issues that led to the cancellation of bags within a shipment.

*Success Response:*



Success. Check the example shown below or refer `ShipmentBagReasons` for more details.


Schema: `ShipmentReasons`









---


#### updateShipmentStatus
Update the shipment status

```golang

 data, err :=  Order.UpdateShipmentStatus(ShipmentID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ShipmentID | string | ID of the shipment. An order may contain multiple items and may get divided into one or more shipment, each having its own ID. | 


| body |  UpdateShipmentStatusRequest | "Request body" 


Use this API to update the status of a shipment using its shipment ID.

*Success Response:*



Successfully updateShipmentStatus!


Schema: `ShipmentApplicationStatusResponse`









---



---


## Rewards


#### getOfferByName
Get offer by name

```golang

 data, err :=  Rewards.GetOfferByName(Name);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Name | string | The name given to the offer. | 




Use this API to get fetch the specific offer details and configuration by the name of the offer.

*Success Response:*



Success. Check example below or refer `Offer` for more details.


Schema: `Offer`


*Examples:*


Success
```json
{
  "value": {
    "_id": "63bfb9c1195d62ac089641cd",
    "application_id": "5d5c304a4df829372e2ad6d1",
    "name": "order",
    "_schedule": {},
    "active": true,
    "banner_image": {
      "aspect_ratio": "2:1",
      "secure_url": "https://hdn-1.fynd.com/company/884/applications/000000000000000000000001/rewards/pictures/landscape-banner/original/dbY4bHh9d-reward-banner.png"
    },
    "created_at": "2023-01-12T07:41:53.356Z",
    "display": {
      "validity": 1000,
      "validity_unit": "hours"
    },
    "info_action": {
      "type": "",
      "page": {
        "type": "",
        "params": {
          "slug": null
        },
        "url": ""
      }
    },
    "rule": {
      "claimed": true,
      "value": 10,
      "value_type": "percent",
      "validity": 60000
    },
    "share": {
      "default": "Hey, join me on {{application_name}} and get exciting offers and reward points. Signup today and quickly earn Rs.{{offer_amount}}. Visit {{offer_link}} now!",
      "text": ""
    },
    "sub_text": "Purchase and get reward points",
    "text": "Order & Earn",
    "type": "earn",
    "updated_at": "2023-01-12T15:46:04.854Z",
    "updated_by": "6678589f5d0df704c9996644",
    "url": ""
  }
}
```









---


#### catalogueOrder
Get all transactions of reward points

```golang

 data, err :=  Rewards.CatalogueOrder(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  CatalogueOrderRequest | "Request body" 


Use this API to evaluate the amount of reward points that could be earned on any catalogue product.

*Success Response:*



Success. Check example below or refer `CatalogueOrderResponse` for more details.


Schema: `CatalogueOrderResponse`


*Examples:*


Success
```json
{
  "value": {
    "articles": [
      {
        "id": "qwer",
        "price": 2356,
        "points": 1000
      }
    ]
  }
}
```









---


#### getUserPointsHistory
Get all transactions of reward points

```golang

 data, err :=  Rewards.GetUserPointsHistory(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `PageID`, `PageSize`



Use this API to fetch a list of points transactions like giveaway points, signup points, referral points, order earn points, redeem points and expired points.

*Success Response:*



Success. Check example below or refer `PointsHistoryResponse` for more details.


Schema: `PointsHistoryResponse`


*Examples:*


Success
```json
{
  "value": {
    "items": [
      {
        "_id": "6464a9eb70e0debb8462667d",
        "application_id": "000000000000000000000001",
        "user_id": "000000000000000009802478",
        "points": 500,
        "remaining_points": 500,
        "txn_name": "signup_credit",
        "claimed": true,
        "expires_on": "2023-06-14T10:18:19.118Z",
        "meta": {
          "offer_id": "630e0b8e349f3f1cfbec572f"
        },
        "created_at": "2023-05-17T10:18:19.118Z",
        "updated_at": "2023-05-17T10:18:19.118Z",
        "text_1": "Signup points",
        "text_2": "Additional Points",
        "text_3": "Will expire on 3:48 PM, 14 Jun'23"
      },
      {
        "_id": "6464a9d370e0debb84626677",
        "application_id": "000000000000000000000001",
        "user_id": "000000000000000009802478",
        "points": 500,
        "remaining_points": 500,
        "txn_name": "signup_credit",
        "claimed": true,
        "expires_on": "2023-06-14T10:17:55.588Z",
        "meta": {
          "offer_id": "630e0b8e349f3f1cfbec572f"
        },
        "created_at": "2023-05-17T10:17:55.588Z",
        "updated_at": "2023-05-17T10:17:55.588Z",
        "text_1": "Signup points",
        "text_2": "Additional Points",
        "text_3": "Will expire on 3:47 PM, 14 Jun'23"
      }
    ],
    "page": {
      "current": 0,
      "item_total": 2,
      "type": "cursor",
      "size": 1,
      "has_previous": true,
      "has_next": false,
      "next_id": ""
    }
  }
}
```









---


#### getUserPoints
Get total available points of a user

```golang

 data, err :=  Rewards.GetUserPoints();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to retrieve total available points of a user for current application.

*Success Response:*



Success. Check example below or refer `PointsResponse` for more details.


Schema: `PointsResponse`


*Examples:*


Success
```json
{
  "value": {
    "points": 100
  }
}
```









---


#### getUserReferralDetails
Get referral details of a user

```golang

 data, err :=  Rewards.GetUserReferralDetails();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Use this API to retrieve the referral details like referral code of a user.

*Success Response:*



Success. Check example below or refer `ReferralDetailsResponse` for more details.


Schema: `ReferralDetailsResponse`


*Examples:*


Success
```json
{
  "value": {
    "referral": {
      "_id": "6325b07224c4f10f7ca2ef0b",
      "application_id": "000000000000000000000001",
      "name": "referral",
      "_schedule": {
        "start": "2020-01-27T20:07:27.418Z",
        "end": "2030-01-27T18:30:00.000Z"
      },
      "active": true,
      "banner_image": {
        "secure_url": "https://hdn-1.fynd.com/company/884/applications/000000000000000000000001/rewards/pictures/landscape-banner/original/05ZGi_YkI-reward-banner.png",
        "aspect_ratio": "2:1"
      },
      "created_at": "2022-09-17T11:33:06.862Z",
      "display": {
        "validity": 4,
        "validity_unit": "weeks"
      },
      "info_action": {
        "type": "",
        "page": {
          "type": "",
          "params": {
            "slug": null
          },
          "url": ""
        }
      },
      "rule": {
        "amount": 100,
        "referrer_amount": 200,
        "amount_validity": 40320,
        "threshold": 10000,
        "counter": 0
      },
      "share": {
        "default": "",
        "text": "👋 Hey there!\n\nI just gave you 100 points.\n\nShop from brands like Campus Sutra, Snitch, Freakins, Off Duty, and other streetwear and athleisure brands\n\n10 GoFynd points = ₹1 off on your next order.\n\nLet me break it down for you: Every 10 GoFynd Points you have will get you a ₹1 discount on your transaction. So, for example, if you have 1000 GoFynd Points, you can enjoy a fabulous ₹100 off during checkout.\n\nJoin me on GoFynd and signup using my code E4WZC5\n\nVisit Now: https://www.fynd.com/l/E4WZC5"
      },
      "sub_text": "When your friends sign up using your referral code.\n\nYou get ₹200 and They get ₹100",
      "text": "Refer & Earn",
      "type": "earn",
      "updated_at": "2023-05-22T13:59:28.295Z",
      "updated_by": "b360c46c3f5c939dadf8f9ab",
      "url": "",
      "updated_by_name": "jaykaria_gofynd_com_39144"
    },
    "share": {
      "default": "",
      "text": "👋 Hey there!\n\nI just gave you 100 points.\n\nShop from brands like Campus Sutra, Snitch, Freakins, Off Duty, and other streetwear and athleisure brands\n\n10 GoFynd points = ₹1 off on your next order.\n\nLet me break it down for you: Every 10 GoFynd Points you have will get you a ₹1 discount on your transaction. So, for example, if you have 1000 GoFynd Points, you can enjoy a fabulous ₹100 off during checkout.\n\nJoin me on GoFynd and signup using my code E4WZC5\n\nVisit Now: https://www.fynd.com/l/E4WZC5"
    },
    "referrer_info": "Code redeemed",
    "user": {
      "blocked": false,
      "referral_code": "E4WZC5",
      "redeemed": true
    },
    "terms_conditions_link": "https://fynd.freshdesk.com/support/solutions/folders/33000111619"
  }
}
```









---


#### getOrderDiscount
Calculates the discount on order-amount

```golang

 data, err :=  Rewards.GetOrderDiscount(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  OrderDiscountRequest | "Request body" 


Use this API to calculate the discount on the order amount, based on all the amount range configured in Order Discount offer.

*Success Response:*



Success. Check example below or refer `OrderDiscountResponse` for more details.


Schema: `OrderDiscountResponse`


*Examples:*


Success
```json
{
  "value": {
    "order_amount": 3000,
    "discount": {
      "absolute": 0,
      "currency": "INR",
      "percent": 0,
      "display_absoulte": "0.00",
      "display_percent": 0
    },
    "base_discount": {
      "absolute": 0,
      "currency": "INR",
      "percent": 0,
      "display_absoulte": "0.00",
      "display_percent": 0
    },
    "points": 0,
    "applied_rule_bucket": {
      "low": 1,
      "high": -1,
      "max": 1000,
      "value": 10,
      "value_type": "percent"
    }
  }
}
```









---


#### redeemReferralCode
Redeems a referral code and credits reward points to referee and the referrer as per the configuration

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


*Examples:*


Success
```json
{
  "value": {
    "redeemed": true,
    "message": "Successfully redeemed referral code",
    "referrer_info": "Referred by Abhinav Maurya",
    "referrer_id": "600693a01faf8695d70a15df",
    "points": 100
  }
}
```









---



---


## PosCart


#### getCart
Fetch all items added to the cart

```golang

 data, err :=  PosCart.GetCart(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |












| xQuery | struct | Includes properties such as `ID`, `I`, `B`, `AssignCardID`, `AreaCode`, `BuyNow`



Use this API to get details of all the items added to a cart.

*Success Response:*



Success. Returns a Cart object. Check the example shown below or refer `CartDetailResponse` for more details.


Schema: `CartDetailResponse`









---


#### getCartLastModified
Fetch last-modified timestamp

```golang

 data, err :=  PosCart.GetCartLastModified(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `ID`



Use this API to fetch Last-Modified timestamp in header metadata.

*Success Response:*



Success. Receives last modifed timestamp in the header.






---


#### addItems
Add items to cart

```golang

 data, err :=  PosCart.AddItems(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |








| xQuery | struct | Includes properties such as `I`, `B`, `AreaCode`, `BuyNow`

| body |  AddCartRequest | "Request body" 


Use this API to add items to the cart.

*Success Response:*



Success. Returns a cart object as shown below. Refer `AddCartDetailResponse` for more details.


Schema: `AddCartDetailResponse`


*Examples:*


Product has been added to your cart
```json
{
  "value": {
    "message": "Product has been added to your cart",
    "success": true,
    "cart": {
      "applied_promo_details": [
        {
          "application_id": "000000000000000000000001",
          "apply_all_offers": false,
          "apply_exclusive": null,
          "buy_article_dict": {
            "rule#1": {
              "633527e850dfb8e73e6de996_0": {
                "added_in_cart": false,
                "added_quantity": 0,
                "amount_paid": 0,
                "applicable_credits": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 0,
                  "floor_value": 0,
                  "source_currency_code": "FC",
                  "source_currency_symbol": "",
                  "source_value": 0,
                  "value": 0
                },
                "article_assign_status": true,
                "article_assignment": {
                  "level": "multi-companies",
                  "strategy": "optimal"
                },
                "article_error": {
                  "type": null,
                  "value": null,
                  "message": null
                },
                "article_id": "633527e850dfb8e73e6de996",
                "article_index": "0",
                "article_meta": {
                  "article_id": "633527e850dfb8e73e6de996",
                  "bulk_margin": 0,
                  "bulk_discount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "INR",
                    "source_currency_symbol": "₹",
                    "source_value": 0,
                    "value": 0
                  },
                  "bulk_coupon_code": null,
                  "country_of_origin": "India",
                  "dimension": {
                    "width": 70,
                    "unit": "cm",
                    "height": 70,
                    "is_default": true,
                    "length": 70
                  },
                  "fragile": false,
                  "manufacturer": {
                    "address": "10, PALI MALA RD, ADARSH NAGAR, BANDRA WEST,, MAHARASHTRA, MUMBAI",
                    "name": "Sabki Shop",
                    "is_default": true
                  },
                  "weight": {
                    "shipping": 250,
                    "unit": "gram",
                    "is_default": true
                  },
                  "raw_meta": {
                    "fynd_identifier": "4187_3412343256098",
                    "is_set_article": false,
                    "set_quantity": 1
                  },
                  "is_set": false,
                  "set": {},
                  "identifier": {
                    "ean": "3412343256098"
                  },
                  "hsn_code": null,
                  "hsn_code_id": "625fbd96faeed8b3df6ec3ce",
                  "raw_price": {
                    "effective": {
                      "conversion_rate": 1,
                      "currency_code": "INR",
                      "currency_symbol": "₹",
                      "floor_source_value": 3899,
                      "floor_value": 3899,
                      "source_currency_code": "INR",
                      "source_currency_symbol": "₹",
                      "source_value": 3899,
                      "value": 3899
                    },
                    "marked": {
                      "conversion_rate": 1,
                      "currency_code": "INR",
                      "currency_symbol": "₹",
                      "floor_source_value": 3899,
                      "floor_value": 3899,
                      "source_currency_code": "INR",
                      "source_currency_symbol": "₹",
                      "source_value": 3899,
                      "value": 3899
                    },
                    "transfer": {
                      "conversion_rate": 1,
                      "currency_code": "INR",
                      "currency_symbol": "₹",
                      "floor_source_value": 0,
                      "floor_value": 0,
                      "source_currency_code": "INR",
                      "source_currency_symbol": "₹",
                      "source_value": 0,
                      "value": 0
                    }
                  },
                  "seller_identifier": "3412343256098",
                  "_custom_json": {},
                  "tags": [],
                  "return_config": {
                    "time": 0,
                    "unit": "days",
                    "returnable": false
                  }
                },
                "attributes": {
                  "essential": "No",
                  "item_code": "BSFJ2-12608",
                  "gender": [
                    "Men"
                  ],
                  "denim_type": "Solid",
                  "pattern": "Solid",
                  "product_fit": "Slim",
                  "currency": "INR",
                  "media": [
                    {
                      "type": "image",
                      "meta": {
                        "brand": "nike",
                        "item_code": "BSFJ2-12608",
                        "sequence": 0
                      },
                      "url": "https://hdn-1.addsale.com/addsale/products/pictures/item/free/original/nike/BSFJ2-12608/0/jMpkpyx17w-GWQjymow3-Black-Slim-Fit-Jeans.jpeg"
                    }
                  ],
                  "name": "Nike Jeans 12608",
                  "brand_name": "Nike"
                },
                "avl_qty": 97,
                "brand_id": 18,
                "bulk_coupon_applied": false,
                "bulk_coupon_code": null,
                "bulk_discount": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 0,
                  "floor_value": 0,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 0,
                  "value": 0
                },
                "bulk_margin": 0,
                "cancellation_allowed": true,
                "cashback": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "FC",
                    "source_currency_symbol": "",
                    "source_value": 0,
                    "value": 0
                  }
                },
                "category_id": [
                  150
                ],
                "cod_charges": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "INR",
                    "source_currency_symbol": "₹",
                    "source_value": 0,
                    "value": 0
                  }
                },
                "company_id": 2,
                "company_info": {
                  "company_id": 2,
                  "c_name": "Sabki Shop",
                  "c_taxes": [
                    {
                      "rate": 15,
                      "enable": true,
                      "effective_date": "2022-06-05T18:29:23.904000"
                    }
                  ],
                  "company_cin": "U45200MH1992PTC066474"
                },
                "company_taxes": [
                  {
                    "rate": 15,
                    "enable": true,
                    "effective_date": "2022-06-05T18:29:23.904000"
                  }
                ],
                "coupon": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "INR",
                    "source_currency_symbol": "₹",
                    "source_value": 0,
                    "value": 0
                  },
                  "article_count": 0,
                  "cancellation_allowed": true,
                  "return_allowed": true
                },
                "currency_dict": {
                  "code": "INR",
                  "rate": 1
                },
                "custom_meta": {},
                "data": {},
                "delivery_charges": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "INR",
                    "source_currency_symbol": "₹",
                    "source_value": 0,
                    "value": 0
                  }
                },
                "delivery_error_msg": "",
                "delivery_promise": {
                  "timestamp": {
                    "min": 1646257339,
                    "max": 1646689339
                  },
                  "formatted": {
                    "min": "03 Mar, Thursday",
                    "max": "08 Mar, Tuesday"
                  }
                },
                "departments": [
                  99
                ],
                "discount": 0,
                "discount_applied": {},
                "dt_currency": "INR",
                "dt_currency_symbol": "₹",
                "esp_modified": false,
                "extra_meta": {},
                "float_amount_paid": 0,
                "free_gift_items_meta": {},
                "group_id": "",
                "gst_amount": 167.1,
                "gst_tax_percentage": 5,
                "hsn_code": null,
                "hsn_code_id": "625fbd96faeed8b3df6ec3ce",
                "identifiers": {
                  "identifier": "BDnwAINORr6SlTKxZN3w"
                },
                "image_nature": "standard",
                "include": true,
                "is_valid": true,
                "item_id": 75660796,
                "item_size": "34",
                "last_update_at": 1680852767,
                "meta": {},
                "net_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3509.1,
                  "floor_value": 3509.1,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3509.1,
                  "value": 3509.1
                },
                "old_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "original_price_effective": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "original_unit_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "parent_item_identifiers": {
                  "identifier": null,
                  "parent_item_size": null,
                  "parent_item_id": null
                },
                "pickup_stores": [],
                "price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "price_effective": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "price_marked": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "product_tags": [],
                "promotions": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "INR",
                    "source_currency_symbol": "₹",
                    "source_value": 0,
                    "value": 0
                  },
                  "applied": [
                    {
                      "amount": {
                        "conversion_rate": 1,
                        "currency_code": "INR",
                        "currency_symbol": "₹",
                        "floor_source_value": 389.9,
                        "floor_value": 389.9,
                        "source_currency_code": "INR",
                        "source_currency_symbol": "₹",
                        "source_value": 389.9,
                        "value": 389.9
                      },
                      "applied_discounts": [
                        {
                          "type": "percentage",
                          "value": 10
                        }
                      ],
                      "applied_free_articles": [],
                      "apply_exclusive": null,
                      "article_quantity": 1,
                      "buy_rules": [
                        {
                          "all_items": null,
                          "cart_conditions": {},
                          "item_criteria": {
                            "item_brand": [
                              18
                            ]
                          },
                          "mrp_promo": false,
                          "slug": "rule#1"
                        }
                      ],
                      "cancellation_allowed": true,
                      "discount_rules": [
                        {
                          "all_items": null,
                          "buy_condition": "( rule#1 )",
                          "item_criteria": {
                            "buy_rules": [
                              "rule#1"
                            ]
                          },
                          "matched_buy_rules": [
                            "rule#1"
                          ],
                          "offer": {
                            "discount_percentage": 10
                          },
                          "raw_offer": {
                            "item_criteria": {
                              "buy_rules": [
                                "rule#1"
                              ]
                            },
                            "buy_condition": "( rule#1 )",
                            "discount_type": "percentage",
                            "offer": {
                              "discount_percentage": 10
                            }
                          },
                          "type": "percentage"
                        }
                      ],
                      "mrp_promo": false,
                      "offer_text": "10% of on NIKE Jeans",
                      "ownership": {
                        "payable_category": "seller",
                        "payable_by": ""
                      },
                      "payable_category": "seller",
                      "promo_code": null,
                      "promo_id": "642fe329ed159b7aa743a293",
                      "promo_type": "percentage",
                      "promotion_group": "product",
                      "return_allowed": true
                    }
                  ],
                  "apply_exclusive": null,
                  "available": [],
                  "cancellation_allowed": true,
                  "exclusive_promo_applied": false,
                  "mrp_promo_applied": false,
                  "return_allowed": true
                },
                "quantity": 1,
                "quantity_assign_status": false,
                "referral_credits": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "FC",
                    "source_currency_symbol": "",
                    "source_value": 0,
                    "value": 0
                  }
                },
                "reload_data_only": true,
                "return_allowed": false,
                "sc_currency": "INR",
                "sc_currency_symbol": "₹",
                "selling_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "service_item_meta": {
                  "product_group_tags": null,
                  "products": null
                },
                "size_level_total_qty": 1,
                "sizes": [
                  "34"
                ],
                "split_article_id": "633527e850dfb8e73e6de996_0",
                "split_index": 0,
                "store_id": 4187,
                "store_info": {
                  "store_id": 4187,
                  "s_city": "HYDERABAD",
                  "store_name": "AND Inorbit Hyderabad",
                  "store_type": "mall",
                  "store_pincode": 500081,
                  "latitude": 17.4343693,
                  "longitude": 78.3866087
                },
                "strategy_validation_data": {},
                "tags": [
                  "rule#1"
                ],
                "total_gst_amount": 167.1,
                "total_value_of_good": 3342,
                "transfer_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 0,
                  "floor_value": 0,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 0,
                  "value": 0
                },
                "unit_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3509.1,
                  "floor_value": 3509.1,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3509.1,
                  "value": 3509.1
                },
                "valid_inventory": true,
                "value_of_good": 3342,
                "verify_article": false
              }
            }
          },
          "buy_rules": [
            {
              "all_items": null,
              "cart_conditions": {},
              "item_criteria": {
                "item_brand": [
                  18
                ]
              },
              "mrp_promo": false,
              "slug": "rule#1"
            }
          ],
          "buy_rules_map": {
            "rule#1": {
              "item_brand": [
                18
              ]
            }
          },
          "calculate_on": "esp",
          "cancellation_allowed": true,
          "code": null,
          "discount": 389.9,
          "discount_rules": [
            {
              "all_items": null,
              "buy_condition": "( rule#1 )",
              "item_criteria": {
                "buy_rules": [
                  "rule#1"
                ]
              },
              "matched_buy_rules": [
                "rule#1"
              ],
              "offer": {
                "discount_percentage": 10
              },
              "raw_offer": {
                "item_criteria": {
                  "buy_rules": [
                    "rule#1"
                  ]
                },
                "buy_condition": "( rule#1 )",
                "discount_type": "percentage",
                "offer": {
                  "discount_percentage": 10
                }
              },
              "type": "percentage"
            }
          ],
          "id": "642fe329ed159b7aa743a293",
          "mrp_promo": false,
          "offer_text": "10% of on NIKE Jeans",
          "ownership": {
            "payable_category": "seller",
            "payable_by": ""
          },
          "payable_category": "seller",
          "priority": 1,
          "promo_group": "product",
          "remaining_allowed_qty": null,
          "restrictions": {
            "uses": {
              "maximum": {
                "user": 0,
                "total": 0
              },
              "remaining": {
                "user": 0,
                "total": 0
              }
            },
            "user_groups": [],
            "post_order": {
              "return_allowed": true,
              "cancellation_allowed": true
            },
            "user_id": [],
            "payments": [],
            "user_registered": {
              "start": null,
              "end": null
            },
            "platforms": [
              "web",
              "android",
              "ios"
            ],
            "anonymous_users": true
          },
          "return_allowed": true,
          "stackable": true,
          "type": "percentage",
          "usage_meta": null
        }
      ],
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
          "parent_item_identifiers": {
            "identifier": "ZASFF",
            "parent_item_id": 7501190,
            "parent_item_size": "OS"
          },
          "moq": {},
          "delivery_promise": {
            "timestamp": {
              "min": 1646257339,
              "max": 1646689339
            },
            "formatted": {
              "min": "03 Mar, Thursday",
              "max": "08 Mar, Tuesday"
            }
          },
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
            "available_sizes": [
              {
                "is_available": true,
                "display": "OS",
                "value": "OS"
              }
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
          "moq": {},
          "delivery_promise": {
            "timestamp": {
              "min": 1646257339,
              "max": 1646689339
            },
            "formatted": {
              "min": "03 Mar, Thursday",
              "max": "08 Mar, Tuesday"
            }
          },
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
            "available_sizes": [
              {
                "is_available": true,
                "display": "OS",
                "value": "OS"
              }
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
          "moq": {},
          "delivery_promise": {
            "timestamp": {
              "min": 1646257339,
              "max": 1646689339
            },
            "formatted": {
              "min": "03 Mar, Thursday",
              "max": "08 Mar, Tuesday"
            }
          },
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
            "available_sizes": [
              {
                "is_available": true,
                "display": "OS",
                "value": "OS"
              }
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
          "moq": {},
          "delivery_promise": {
            "timestamp": {
              "min": 1646257339,
              "max": 1646689339
            },
            "formatted": {
              "min": "03 Mar, Thursday",
              "max": "08 Mar, Tuesday"
            }
          },
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
            "available_sizes": [
              {
                "is_available": true,
                "display": "OS",
                "value": "OS"
              }
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
      "buy_now": false,
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
          "parent_item_identifiers": {
            "identifier": "ZASFF",
            "parent_item_id": 7501190,
            "parent_item_size": "OS"
          },
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
      "buy_now": false,
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
Update items in the cart

```golang

 data, err :=  PosCart.UpdateCart(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |










| xQuery | struct | Includes properties such as `ID`, `I`, `B`, `AreaCode`, `BuyNow`

| body |  UpdateCartRequest | "Request body" 


<p>Use this API to update items added to the cart with the help of a request object containing attributes like item_quantity and item_size. These attributes will be fetched from the following APIs</p> <ul> <li><font color="monochrome">operation</font> Operation for current api call. <b>update_item</b> for update items. <b>remove_item</b> for removing items.</li> <li> <font color="monochrome">item_id</font>  "/platform/content/v1/products/"</li> <li> <font color="monochrome">item_size</font>   "/platform/content/v1/products/:slug/sizes/"</li> <li> <font color="monochrome">quantity</font>  item quantity (must be greater than or equal to 1)</li> <li> <font color="monochrome">article_id</font>   "/content​/v1​/products​/:identifier​/sizes​/price​/"</li> <li> <font color="monochrome">item_index</font>  item position in the cart (must be greater than or equal to 0)</li> </ul>

*Success Response:*



Success. Updates and returns a cart object as shown below. Refer `UpdateCartDetailResponse` for more details.


Schema: `UpdateCartDetailResponse`


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
          "parent_item_identifiers": {
            "identifier": "ZASFF",
            "parent_item_id": 7501190,
            "parent_item_size": "OS"
          },
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
      "buy_now": false,
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
      "applied_promo_details": [
        {
          "application_id": "000000000000000000000001",
          "apply_all_offers": false,
          "apply_exclusive": null,
          "buy_article_dict": {
            "rule#1": {
              "633527e850dfb8e73e6de996_0": {
                "added_in_cart": false,
                "added_quantity": 0,
                "amount_paid": 0,
                "applicable_credits": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 0,
                  "floor_value": 0,
                  "source_currency_code": "FC",
                  "source_currency_symbol": "",
                  "source_value": 0,
                  "value": 0
                },
                "article_assign_status": true,
                "article_assignment": {
                  "level": "multi-companies",
                  "strategy": "optimal"
                },
                "article_error": {
                  "type": null,
                  "value": null,
                  "message": null
                },
                "article_id": "633527e850dfb8e73e6de996",
                "article_index": "0",
                "article_meta": {
                  "article_id": "633527e850dfb8e73e6de996",
                  "bulk_margin": 0,
                  "bulk_discount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "INR",
                    "source_currency_symbol": "₹",
                    "source_value": 0,
                    "value": 0
                  },
                  "bulk_coupon_code": null,
                  "country_of_origin": "India",
                  "dimension": {
                    "width": 70,
                    "unit": "cm",
                    "height": 70,
                    "is_default": true,
                    "length": 70
                  },
                  "fragile": false,
                  "manufacturer": {
                    "address": "10, PALI MALA RD, ADARSH NAGAR, BANDRA WEST,, MAHARASHTRA, MUMBAI",
                    "name": "Sabki Shop",
                    "is_default": true
                  },
                  "weight": {
                    "shipping": 250,
                    "unit": "gram",
                    "is_default": true
                  },
                  "raw_meta": {
                    "fynd_identifier": "4187_3412343256098",
                    "is_set_article": false,
                    "set_quantity": 1
                  },
                  "is_set": false,
                  "set": {},
                  "identifier": {
                    "ean": "3412343256098"
                  },
                  "hsn_code": null,
                  "hsn_code_id": "625fbd96faeed8b3df6ec3ce",
                  "raw_price": {
                    "effective": {
                      "conversion_rate": 1,
                      "currency_code": "INR",
                      "currency_symbol": "₹",
                      "floor_source_value": 3899,
                      "floor_value": 3899,
                      "source_currency_code": "INR",
                      "source_currency_symbol": "₹",
                      "source_value": 3899,
                      "value": 3899
                    },
                    "marked": {
                      "conversion_rate": 1,
                      "currency_code": "INR",
                      "currency_symbol": "₹",
                      "floor_source_value": 3899,
                      "floor_value": 3899,
                      "source_currency_code": "INR",
                      "source_currency_symbol": "₹",
                      "source_value": 3899,
                      "value": 3899
                    },
                    "transfer": {
                      "conversion_rate": 1,
                      "currency_code": "INR",
                      "currency_symbol": "₹",
                      "floor_source_value": 0,
                      "floor_value": 0,
                      "source_currency_code": "INR",
                      "source_currency_symbol": "₹",
                      "source_value": 0,
                      "value": 0
                    }
                  },
                  "seller_identifier": "3412343256098",
                  "_custom_json": {},
                  "tags": [],
                  "return_config": {
                    "time": 0,
                    "unit": "days",
                    "returnable": false
                  }
                },
                "attributes": {
                  "essential": "No",
                  "item_code": "BSFJ2-12608",
                  "gender": [
                    "Men"
                  ],
                  "denim_type": "Solid",
                  "pattern": "Solid",
                  "product_fit": "Slim",
                  "currency": "INR",
                  "media": [
                    {
                      "type": "image",
                      "meta": {
                        "brand": "nike",
                        "item_code": "BSFJ2-12608",
                        "sequence": 0
                      },
                      "url": "https://hdn-1.addsale.com/addsale/products/pictures/item/free/original/nike/BSFJ2-12608/0/jMpkpyx17w-GWQjymow3-Black-Slim-Fit-Jeans.jpeg"
                    }
                  ],
                  "name": "Nike Jeans 12608",
                  "brand_name": "Nike"
                },
                "avl_qty": 97,
                "brand_id": 18,
                "bulk_coupon_applied": false,
                "bulk_coupon_code": null,
                "bulk_discount": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 0,
                  "floor_value": 0,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 0,
                  "value": 0
                },
                "bulk_margin": 0,
                "cancellation_allowed": true,
                "cashback": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "FC",
                    "source_currency_symbol": "",
                    "source_value": 0,
                    "value": 0
                  }
                },
                "category_id": [
                  150
                ],
                "cod_charges": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "INR",
                    "source_currency_symbol": "₹",
                    "source_value": 0,
                    "value": 0
                  }
                },
                "company_id": 2,
                "company_info": {
                  "company_id": 2,
                  "c_name": "Sabki Shop",
                  "c_taxes": [
                    {
                      "rate": 15,
                      "enable": true,
                      "effective_date": "2022-06-05T18:29:23.904000"
                    }
                  ],
                  "company_cin": "U45200MH1992PTC066474"
                },
                "company_taxes": [
                  {
                    "rate": 15,
                    "enable": true,
                    "effective_date": "2022-06-05T18:29:23.904000"
                  }
                ],
                "coupon": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "INR",
                    "source_currency_symbol": "₹",
                    "source_value": 0,
                    "value": 0
                  },
                  "article_count": 0,
                  "cancellation_allowed": true,
                  "return_allowed": true
                },
                "currency_dict": {
                  "code": "INR",
                  "rate": 1
                },
                "custom_meta": {},
                "data": {},
                "delivery_charges": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "INR",
                    "source_currency_symbol": "₹",
                    "source_value": 0,
                    "value": 0
                  }
                },
                "delivery_error_msg": "",
                "delivery_promise": {
                  "timestamp": {
                    "min": 1646257339,
                    "max": 1646689339
                  },
                  "formatted": {
                    "min": "03 Mar, Thursday",
                    "max": "08 Mar, Tuesday"
                  }
                },
                "departments": [
                  99
                ],
                "discount": 0,
                "discount_applied": {},
                "dt_currency": "INR",
                "dt_currency_symbol": "₹",
                "esp_modified": false,
                "extra_meta": {},
                "float_amount_paid": 0,
                "free_gift_items_meta": {},
                "group_id": "",
                "gst_amount": 167.1,
                "gst_tax_percentage": 5,
                "hsn_code": null,
                "hsn_code_id": "625fbd96faeed8b3df6ec3ce",
                "identifiers": {
                  "identifier": "BDnwAINORr6SlTKxZN3w"
                },
                "image_nature": "standard",
                "include": true,
                "is_valid": true,
                "item_id": 75660796,
                "item_size": "34",
                "last_update_at": 1680852767,
                "meta": {},
                "net_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3509.1,
                  "floor_value": 3509.1,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3509.1,
                  "value": 3509.1
                },
                "old_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "original_price_effective": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "original_unit_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "parent_item_identifiers": {
                  "identifier": null,
                  "parent_item_size": null,
                  "parent_item_id": null
                },
                "pickup_stores": [],
                "price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "price_effective": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "price_marked": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "product_tags": [],
                "promotions": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "INR",
                    "source_currency_symbol": "₹",
                    "source_value": 0,
                    "value": 0
                  },
                  "applied": [
                    {
                      "amount": {
                        "conversion_rate": 1,
                        "currency_code": "INR",
                        "currency_symbol": "₹",
                        "floor_source_value": 389.9,
                        "floor_value": 389.9,
                        "source_currency_code": "INR",
                        "source_currency_symbol": "₹",
                        "source_value": 389.9,
                        "value": 389.9
                      },
                      "applied_discounts": [
                        {
                          "type": "percentage",
                          "value": 10
                        }
                      ],
                      "applied_free_articles": [],
                      "apply_exclusive": null,
                      "article_quantity": 1,
                      "buy_rules": [
                        {
                          "all_items": null,
                          "cart_conditions": {},
                          "item_criteria": {
                            "item_brand": [
                              18
                            ]
                          },
                          "mrp_promo": false,
                          "slug": "rule#1"
                        }
                      ],
                      "cancellation_allowed": true,
                      "discount_rules": [
                        {
                          "all_items": null,
                          "buy_condition": "( rule#1 )",
                          "item_criteria": {
                            "buy_rules": [
                              "rule#1"
                            ]
                          },
                          "matched_buy_rules": [
                            "rule#1"
                          ],
                          "offer": {
                            "discount_percentage": 10
                          },
                          "raw_offer": {
                            "item_criteria": {
                              "buy_rules": [
                                "rule#1"
                              ]
                            },
                            "buy_condition": "( rule#1 )",
                            "discount_type": "percentage",
                            "offer": {
                              "discount_percentage": 10
                            }
                          },
                          "type": "percentage"
                        }
                      ],
                      "mrp_promo": false,
                      "offer_text": "10% of on NIKE Jeans",
                      "ownership": {
                        "payable_category": "seller",
                        "payable_by": ""
                      },
                      "payable_category": "seller",
                      "promo_code": null,
                      "promo_id": "642fe329ed159b7aa743a293",
                      "promo_type": "percentage",
                      "promotion_group": "product",
                      "return_allowed": true
                    }
                  ],
                  "apply_exclusive": null,
                  "available": [],
                  "cancellation_allowed": true,
                  "exclusive_promo_applied": false,
                  "mrp_promo_applied": false,
                  "return_allowed": true
                },
                "quantity": 1,
                "quantity_assign_status": false,
                "referral_credits": {
                  "amount": {
                    "conversion_rate": 1,
                    "currency_code": "INR",
                    "currency_symbol": "₹",
                    "floor_source_value": 0,
                    "floor_value": 0,
                    "source_currency_code": "FC",
                    "source_currency_symbol": "",
                    "source_value": 0,
                    "value": 0
                  }
                },
                "reload_data_only": true,
                "return_allowed": false,
                "sc_currency": "INR",
                "sc_currency_symbol": "₹",
                "selling_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3899,
                  "floor_value": 3899,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3899,
                  "value": 3899
                },
                "service_item_meta": {
                  "product_group_tags": null,
                  "products": null
                },
                "size_level_total_qty": 1,
                "sizes": [
                  "34"
                ],
                "split_article_id": "633527e850dfb8e73e6de996_0",
                "split_index": 0,
                "store_id": 4187,
                "store_info": {
                  "store_id": 4187,
                  "s_city": "HYDERABAD",
                  "store_name": "AND Inorbit Hyderabad",
                  "store_type": "mall",
                  "store_pincode": 500081,
                  "latitude": 17.4343693,
                  "longitude": 78.3866087
                },
                "strategy_validation_data": {},
                "tags": [
                  "rule#1"
                ],
                "total_gst_amount": 167.1,
                "total_value_of_good": 3342,
                "transfer_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 0,
                  "floor_value": 0,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 0,
                  "value": 0
                },
                "unit_price": {
                  "conversion_rate": 1,
                  "currency_code": "INR",
                  "currency_symbol": "₹",
                  "floor_source_value": 3509.1,
                  "floor_value": 3509.1,
                  "source_currency_code": "INR",
                  "source_currency_symbol": "₹",
                  "source_value": 3509.1,
                  "value": 3509.1
                },
                "valid_inventory": true,
                "value_of_good": 3342,
                "verify_article": false
              }
            }
          },
          "buy_rules": [
            {
              "all_items": null,
              "cart_conditions": {},
              "item_criteria": {
                "item_brand": [
                  18
                ]
              },
              "mrp_promo": false,
              "slug": "rule#1"
            }
          ],
          "buy_rules_map": {
            "rule#1": {
              "item_brand": [
                18
              ]
            }
          },
          "calculate_on": "esp",
          "cancellation_allowed": true,
          "code": null,
          "discount": 389.9,
          "discount_rules": [
            {
              "all_items": null,
              "buy_condition": "( rule#1 )",
              "item_criteria": {
                "buy_rules": [
                  "rule#1"
                ]
              },
              "matched_buy_rules": [
                "rule#1"
              ],
              "offer": {
                "discount_percentage": 10
              },
              "raw_offer": {
                "item_criteria": {
                  "buy_rules": [
                    "rule#1"
                  ]
                },
                "buy_condition": "( rule#1 )",
                "discount_type": "percentage",
                "offer": {
                  "discount_percentage": 10
                }
              },
              "type": "percentage"
            }
          ],
          "id": "642fe329ed159b7aa743a293",
          "mrp_promo": false,
          "offer_text": "10% of on NIKE Jeans",
          "ownership": {
            "payable_category": "seller",
            "payable_by": ""
          },
          "payable_category": "seller",
          "priority": 1,
          "promo_group": "product",
          "remaining_allowed_qty": null,
          "restrictions": {
            "uses": {
              "maximum": {
                "user": 0,
                "total": 0
              },
              "remaining": {
                "user": 0,
                "total": 0
              }
            },
            "user_groups": [],
            "post_order": {
              "return_allowed": true,
              "cancellation_allowed": true
            },
            "user_id": [],
            "payments": [],
            "user_registered": {
              "start": null,
              "end": null
            },
            "platforms": [
              "web",
              "android",
              "ios"
            ],
            "anonymous_users": true
          },
          "return_allowed": true,
          "stackable": true,
          "type": "percentage",
          "usage_meta": null
        }
      ],
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
          "parent_item_identifiers": {
            "identifier": "ZASFF",
            "parent_item_id": 7501190,
            "parent_item_size": "OS"
          },
          "moq": {},
          "delivery_promise": {
            "timestamp": {
              "min": 1646257339,
              "max": 1646689339
            },
            "formatted": {
              "min": "03 Mar, Thursday",
              "max": "08 Mar, Tuesday"
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
            "available_sizes": [
              {
                "is_available": true,
                "display": "OS",
                "value": "OS"
              }
            ],
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
      "buy_now": false,
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
Count items in the cart

```golang

 data, err :=  PosCart.GetItemCount(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `ID`, `BuyNow`



Use this API to get the total number of items present in cart.

*Success Response:*



Success. Returns the total count of items in a user's cart.


Schema: `CartItemCountResponse`









---


#### getCoupons
Fetch Coupon

```golang

 data, err :=  PosCart.GetCoupons(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `ID`, `BuyNow`



Use this API to get a list of available coupons along with their details.

*Success Response:*



Success. Returns a coupon object which has a list of all the eligible coupons. Refer `GetCouponResponse` for more details.


Schema: `GetCouponResponse`









---


#### applyCoupon
Apply Coupon

```golang

 data, err :=  PosCart.ApplyCoupon(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |










| xQuery | struct | Includes properties such as `I`, `B`, `P`, `ID`, `BuyNow`

| body |  ApplyCouponRequest | "Request body" 


Use this API to apply coupons on items in the cart.

*Success Response:*



Success. Returns coupons applied to the cart along with item details and price breakup. Refer `CartDetailResponse` for more details.


Schema: `CartDetailResponse`









---


#### removeCoupon
Remove Coupon Applied

```golang

 data, err :=  PosCart.RemoveCoupon(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `ID`, `BuyNow`



Remove Coupon applied on the cart by passing uid in request body.

*Success Response:*



Success. Returns coupons removed from the cart along with item details and price breakup. Refer `CartDetailResponse` for more details.


Schema: `CartDetailResponse`









---


#### getBulkDiscountOffers
Get discount offers based on quantity

```golang

 data, err :=  PosCart.GetBulkDiscountOffers(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |








| xQuery | struct | Includes properties such as `ItemID`, `ArticleID`, `UID`, `Slug`



Use this API to get a list of applicable offers along with current, next and best offer for given product. Either one of uid, item_id, slug should be present.

*Success Response:*



Success. Returns a data object containing the seller details and available offers (if exists) on bulk products. Refer `BulkPriceResponse` for more details.


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









---


#### applyRewardPoints
Apply reward points at cart

```golang

 data, err :=  PosCart.ApplyRewardPoints(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |








| xQuery | struct | Includes properties such as `ID`, `I`, `B`, `BuyNow`

| body |  RewardPointRequest | "Request body" 


Use this API to redeem a fixed no. of reward points by applying it to the cart.

*Success Response:*



Success. Returns a Cart object. Check the example shown below or refer `CartDetailResponse` for more details.


Schema: `CartDetailResponse`









---


#### getAddresses
Fetch address

```golang

 data, err :=  PosCart.GetAddresses(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |












| xQuery | struct | Includes properties such as `CartID`, `BuyNow`, `MobileNo`, `CheckoutMode`, `Tags`, `IsDefault`



Use this API to get all the addresses associated with an account. If successful, returns a Address resource in the response body specified in GetAddressesResponse.attibutes listed below are optional <ul> <li> <font color="monochrome">uid</font></li> <li> <font color="monochrome">address_id</font></li> <li> <font color="monochrome">mobile_no</font></li> <li> <font color="monochrome">checkout_mode</font></li> <li> <font color="monochrome">tags</font></li> <li> <font color="monochrome">default</font></li> </ul>

*Success Response:*



Success. Returns an Address object containing a list of address saved in the account. Refer `GetAddressesResponse` for more details.


Schema: `GetAddressesResponse`









---


#### addAddress
Add address to an account

```golang

 data, err :=  PosCart.AddAddress(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  Address | "Request body" 


Use this API to add an address to an account.

*Success Response:*



Success. Returns the address ID, a flag whether the address is set as default, and a success message. Refer `SaveAddressResponse` for more details.


Schema: `SaveAddressResponse`









---


#### getAddressById
Fetch a single address by its ID

```golang

 data, err :=  PosCart.GetAddressById(ID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ID | string |  | 













| xQuery | struct | Includes properties such as `CartID`, `BuyNow`, `MobileNo`, `CheckoutMode`, `Tags`, `IsDefault`



Use this API to get an addresses using its ID. If successful, returns a Address resource in the response body specified in `Address`. Attibutes listed below are optional <ul> <li> <font color="monochrome">mobile_no</font></li> <li> <font color="monochrome">checkout_mode</font></li> <li> <font color="monochrome">tags</font></li> <li> <font color="monochrome">default</font></li> </ul>

*Success Response:*



Success. Returns an Address object containing a list of address saved in the account. Refer `Address` for more details.


Schema: `Address`









---


#### updateAddress
Update address added to an account

```golang

 data, err :=  PosCart.UpdateAddress(ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ID | string | ID allotted to the selected address | 


| body |  Address | "Request body" 


<p>Use this API to update an existing address in the account. Request object should contain attributes mentioned in  <font color="blue">Address </font> can be updated. These attributes are:</p> <ul> <li> <font color="monochrome">is_default_address</font></li> <li> <font color="monochrome">landmark</font></li> <li> <font color="monochrome">area</font></li> <li> <font color="monochrome">pincode</font></li> <li> <font color="monochrome">email</font></li> <li> <font color="monochrome">address_type</font></li> <li> <font color="monochrome">name</font></li> <li> <font color="monochrome">address_id</font></li> <li> <font color="monochrome">address</font></li> </ul>

*Success Response:*



Success. Returns the address ID and a message indicating a successful address updation.


Schema: `UpdateAddressResponse`









---


#### removeAddress
Remove address associated with an account

```golang

 data, err :=  PosCart.RemoveAddress(ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| ID | string | ID allotted to the selected address | 




Use this API to delete an address by its ID. This will returns an object that will indicate whether the address was deleted successfully or not.

*Success Response:*



Returns a Status object indicating the success or failure of address deletion.


Schema: `DeleteAddressResponse`









---


#### selectAddress
Select an address from available addresses

```golang

 data, err :=  PosCart.SelectAddress(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |








| xQuery | struct | Includes properties such as `CartID`, `BuyNow`, `I`, `B`

| body |  SelectCartAddressRequest | "Request body" 


<p>Select Address from all addresses associated with the account in order to ship the cart items to that address, otherwise default address will be selected implicitly. See `SelectCartAddressRequest` in schema of request body for the list of attributes needed to select Address from account. On successful request, this API returns a Cart object. Below address attributes are required. <ul> <li> <font color="monochrome">address_id</font></li> <li> <font color="monochrome">billing_address_id</font></li> <li> <font color="monochrome">uid</font></li> </ul></p>

*Success Response:*



Success. Returns a Cart object as shown below. Refer `CartDetailResponse` for more details.  .


Schema: `CartDetailResponse`









---


#### selectPaymentMode
Update cart payment

```golang

 data, err :=  PosCart.SelectPaymentMode(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `ID`, `BuyNow`

| body |  UpdateCartPaymentRequest | "Request body" 


Use this API to update cart payment.

*Success Response:*



Success. Returns a Cart object as shown below. Refer `CartDetailResponse` for more details.


Schema: `CartDetailResponse`









---


#### validateCouponForPayment
Verify the coupon eligibility against the payment mode

```golang

 data, err :=  PosCart.ValidateCouponForPayment(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |














| xQuery | struct | Includes properties such as `ID`, `BuyNow`, `AddressID`, `PaymentMode`, `PaymentIdentifier`, `AggregatorName`, `MerchantCode`



Use this API to validate a coupon against the payment mode such as NetBanking, Wallet, UPI etc.

*Success Response:*



Success. Returns a success message and the coupon validity. Refer `PaymentCouponValidate` for more details.


Schema: `PaymentCouponValidate`









---


#### getShipments
Get delivery date and options before checkout

```golang

 data, err :=  PosCart.GetShipments(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |














| xQuery | struct | Includes properties such as `PickAtStoreUID`, `OrderingStoreID`, `P`, `ID`, `AddressID`, `AreaCode`, `OrderType`



Use this API to get shipment details, expected delivery date, items and price breakup of the shipment.

*Success Response:*



Success. Returns delivery promise along with shipment details and price breakup. Refer `CartShipmentsResponse` for more details.


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









---


#### updateShipments
Update shipment delivery type and quantity before checkout

```golang

 data, err :=  PosCart.UpdateShipments(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |










| xQuery | struct | Includes properties such as `I`, `P`, `ID`, `AddressID`, `OrderType`

| body |  UpdateCartShipmentRequest | "Request body" 


Use this API to update the delivery type and quantity as per customer's preference for either store pick-up or home-delivery.

*Success Response:*



Success. Returns delivery promise along with shipment details and price breakup. Refer `CartShipmentsResponse` for more details.


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









---


#### checkoutCart
Checkout all items in the cart

```golang

 data, err :=  PosCart.CheckoutCart(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |


| xQuery | struct | Includes properties such as `ID`

| body |  CartPosCheckoutDetailRequest | "Request body" 


Use this API to checkout all items in the cart for payment and order generation. For COD, order will be generated directly, whereas for other checkout modes, user will be redirected to a payment gateway.

*Success Response:*



Success. Returns the status of cart checkout. Refer `CartCheckoutResponse` for more details.


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
Update the cart meta

```golang

 data, err :=  PosCart.UpdateCartMeta(xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `ID`, `BuyNow`

| body |  CartMetaRequest | "Request body" 


Use this API to update cart meta like checkout_mode and gstin.

*Success Response:*



Returns a message indicating the success of cart meta updation as shown below.


Schema: `CartMetaResponse`









---


#### getAvailableDeliveryModes
Get available delivery modes for cart

```golang

 data, err :=  PosCart.GetAvailableDeliveryModes(xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |




| xQuery | struct | Includes properties such as `AreaCode`, `ID`



Use this API to get the delivery modes (home-delivery/store-pickup) along with a list of pickup stores available for a given cart at a given PIN Code. User can then view the address of a pickup store with the help of store-address API.

*Success Response:*



Success. Returns the available delivery mode available for a given PIN Code, along with the UID of all the eligible pickup stores.


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



Use this API to get the store details by entering the unique identifier of the pickup stores shown in the response of available-delivery-mode API.

*Success Response:*



Success. Returns available store information with its address as shown below.


Schema: `StoreDetailsResponse`









---


#### getCartShareLink
Generate token for sharing the cart

```golang

 data, err :=  PosCart.GetCartShareLink(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  GetShareCartLinkRequest | "Request body" 


Use this API to generate a shared cart snapshot and return a shortlink token. The link can be shared with other users for getting the same items in their cart.

*Success Response:*



Returns a URL to share and a token as shown below.


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
Get details of a shared cart

```golang

 data, err :=  PosCart.GetCartSharedItems(Token);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Token | string | Token of the shared short link | 




Use this API to get the shared cart details as per the token generated using the share-cart API.

*Success Response:*



Success. Returns a Cart object as per the valid token. Refer `SharedCartResponse` for more details.


Schema: `SharedCartResponse`









---


#### updateCartWithSharedItems
Merge or replace existing cart

```golang

 data, err :=  PosCart.UpdateCartWithSharedItems(Token, Action);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Token | string | Token of the shared short link | 


| Action | string | Operation to perform on the existing cart merge or replace. | 




Use this API to merge the shared cart with existing cart, or replace the existing cart with the shared cart. The `action` parameter is used to indicate the operation Merge or Replace.

*Success Response:*



Success. Returns a merged or replaced cart as per the valid token. Refer `SharedCartResponse` for more details.


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


#### getPincodeCity
Get Pincode API

```golang

 data, err :=  Logistic.GetPincodeCity(Pincode);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Pincode | string | A `pincode` contains a specific address of a location. | 




Get pincode data

*Success Response:*



Get pincode data


Schema: `PincodeApiResponse`


*Examples:*


Pincode data found
```json
{
  "value": {
    "data": [
      {
        "sub_type": "pincode",
        "name": "421202",
        "error": {
          "type": null,
          "value": null,
          "message": null
        },
        "uid": "pincode:INDIA|MAHARASHTRA|MUMBAI|421202",
        "display_name": "421202",
        "meta": {
          "zone": "West",
          "internal_zone_id": 4
        },
        "meta_code": {
          "country_code": "IND",
          "isd_code": "+91"
        },
        "parents": [
          {
            "sub_type": "country",
            "name": "India",
            "display_name": "India",
            "uid": "country:INDIA"
          },
          {
            "sub_type": "state",
            "name": "Maharashtra",
            "display_name": "Maharashtra",
            "uid": "state:INDIA|MAHARASHTRA"
          },
          {
            "sub_type": "city",
            "name": "Thane",
            "display_name": "Thane",
            "uid": "city:INDIA|MAHARASHTRA|MUMBAI"
          }
        ],
        "lat_long": {
          "type": "Point",
          "coordinates": [
            3.8858955,
            7.2272335
          ]
        }
      }
    ],
    "request_uuid": "fce9f431215e71c9ee0e86e792ae1dce4",
    "stormbreaker_uuid": "56cca764-9fab-41f4-adb8-6e9683532aa5",
    "error": {
      "type": null,
      "value": null,
      "message": null
    },
    "success": true
  }
}
```

Pincode not found
```json
{
  "value": {
    "data": [
      {
        "sub_type": "pincode",
        "name": "999999",
        "error": {
          "type": "DoesNotExist",
          "value": "999999",
          "message": "pincode 999999 does not exist"
        }
      }
    ],
    "request_uuid": "fce9fb9215e71c9ee0e86e792ae1dce4",
    "stormbreaker_uuid": "03b353ed-9dbd-4629-80b2-2be337859a20",
    "error": {
      "type": null,
      "value": null,
      "message": null
    },
    "success": false
  }
}
```









---


#### getTatProduct
Get TAT API

```golang

 data, err :=  Logistic.GetTatProduct(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  TATViewRequest | "Request body" 


Get TAT data

*Success Response:*



Get TAT  data


Schema: `TATViewResponse`


*Examples:*


Pincode data found
```json
{
  "value": {
    "source": "FYND-APP",
    "identifier": "PDP",
    "journey": "forward",
    "action": "get_tat",
    "to_pincode": "143001",
    "location_details": [
      {
        "fulfillment_id": 8,
        "from_pincode": "560023",
        "articles": [
          {
            "category": {
              "level": "l3",
              "id": 155
            },
            "manufacturing_time": 2,
            "manufacturing_time_unit": "days",
            "promise": {
              "timestamp": {
                "min": 1663564548,
                "max": 1663650948
              },
              "formatted": {
                "min": "19 Sep, Monday",
                "max": "20 Sep, Tuesday"
              }
            },
            "error": {
              "type": null,
              "value": null,
              "message": null
            },
            "is_cod_available": true,
            "_manufacturing_time_seconds": 172800
          }
        ]
      }
    ],
    "request_uuid": "b4adf5508e34f17971817c3581e16310",
    "stormbreaker_uuid": "4b8084d4-ea74-45af-8ddc-c38e29bf0cfb",
    "error": {
      "type": null,
      "value": null,
      "message": null
    },
    "to_city": "Amritsar",
    "payment_mode": "prepaid",
    "is_cod_available": true,
    "success": true
  }
}
```

Pincode not found
```json
{
  "value": {
    "source": "FYND-APP",
    "identifier": "PDP",
    "journey": "forward",
    "action": "get_tat",
    "to_pincode": "99999",
    "location_details": [
      {
        "fulfillment_id": 8,
        "from_pincode": "560023",
        "articles": [
          {
            "category": {
              "level": "l3",
              "id": 155
            },
            "manufacturing_time": 2,
            "manufacturing_time_unit": "days",
            "promise": null,
            "error": {
              "type": "ValueError",
              "value": "99999",
              "message": "We are not delivering to 99999"
            }
          }
        ]
      }
    ],
    "request_uuid": "4b99d15fddb2b9fc2d6ab99a1c933010",
    "stormbreaker_uuid": "6a847909-1d59-43e7-9ae0-15f5de8fc7d7",
    "error": {
      "type": "ValueError",
      "value": "99999",
      "message": "All of the items in your cart are not deliverable to 99999"
    },
    "to_city": "",
    "payment_mode": "prepaid",
    "is_cod_available": true,
    "success": false
  }
}
```









---


#### getAllCountries
Get Country List

```golang

 data, err :=  Logistic.GetAllCountries();
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



Get all countries

*Success Response:*



Get Country List


Schema: `CountryListResponse`









---


#### getPincodeZones
GET zone from the Pincode.

```golang

 data, err :=  Logistic.GetPincodeZones(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  GetZoneFromPincodeViewRequest | "Request body" 


This API returns zone from the Pincode View.

*Success Response:*



Response status_code


Schema: `GetZoneFromPincodeViewResponse`









---


#### getOptimalLocations
GET zone from the Pincode.

```golang

 data, err :=  Logistic.GetOptimalLocations(body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| body |  ReAssignStoreRequest | "Request body" 


This API returns zone from the Pincode View.

*Success Response:*



Response status_code


Schema: `ReAssignStoreResponse`









---



---


---
---
