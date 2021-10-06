# FDK Platform Front API Documentaion


* [Common](#Common) - Application configuration apis 
* [Lead](#Lead) - Handles communication between Administrator <-> Staff and Staff <-> Users 
* [Feedback](#Feedback) - User Reviews and Rating System 
* [Theme](#Theme) - Responsible for themes 
* [User](#User) - Authentication Service 
* [Content](#Content) - Content System 
* [Billing](#Billing) - Handle platform subscription 
* [Communication](#Communication) - Manages email, sms, push notifications sent to users 
* [Payment](#Payment) - Collect payment through many payment gateway i.e Stripe, Razorpay, Juspay etc.into Fynd or Self account 
* [Order](#Order) - Handles Platform websites OMS 
* [Catalog](#Catalog) - Catalog API's allows you to access list of products, prices, seller details, similar features, variants and many more useful features.  
* [CompanyProfile](#CompanyProfile) - Company Profile API's allows you to access list of products, prices, seller details, similar features, variants and many more useful features.  
* [FileStorage](#FileStorage) - File Storage 
* [Share](#Share) - Short link and QR Code 
* [Inventory](#Inventory) -  
* [Configuration](#Configuration) - Application configuration apis 
* [Cart](#Cart) - Cart APIs 
* [Rewards](#Rewards) - Rewards 
* [Analytics](#Analytics) - Perceptor analytics 
* [Discount](#Discount) - Discount 
* [Partner](#Partner) - Partner configuration apis 
* [Webhook](#Webhook) - Webhook dispatcher with retry and one event to many subscriber vice versa 

----
----

### Classes and Methods


* [Common](#Common)
  * Methods
    * [getLocations](#getlocations)
    

* [Lead](#Lead)
  * Methods
    * [getTickets](#gettickets)
    * [createTicket](#createticket)
    * [getTickets](#gettickets)
    * [getTicket](#getticket)
    * [editTicket](#editticket)
    * [getTicket](#getticket)
    * [editTicket](#editticket)
    * [createHistory](#createhistory)
    * [getTicketHistory](#gettickethistory)
    * [getFeedbacks](#getfeedbacks)
    * [submitFeedback](#submitfeedback)
    * [createHistory](#createhistory)
    * [getTicketHistory](#gettickethistory)
    * [getCustomForm](#getcustomform)
    * [editCustomForm](#editcustomform)
    * [getCustomForms](#getcustomforms)
    * [createCustomForm](#createcustomform)
    * [getTokenForVideoRoom](#gettokenforvideoroom)
    * [getTokenForVideoRoom](#gettokenforvideoroom)
    * [getVideoParticipants](#getvideoparticipants)
    * [getVideoParticipants](#getvideoparticipants)
    * [openVideoRoom](#openvideoroom)
    * [closeVideoRoom](#closevideoroom)
    

* [Feedback](#Feedback)
  * Methods
    * [getAttributes](#getattributes)
    * [getCustomerReviews](#getcustomerreviews)
    * [updateApprove](#updateapprove)
    * [getHistory](#gethistory)
    * [getApplicationTemplates](#getapplicationtemplates)
    * [createTemplate](#createtemplate)
    * [getTemplateById](#gettemplatebyid)
    * [updateTemplate](#updatetemplate)
    * [updateTemplateStatus](#updatetemplatestatus)
    

* [Theme](#Theme)
  * Methods
    * [getAllPages](#getallpages)
    * [createPage](#createpage)
    * [updateMultiplePages](#updatemultiplepages)
    * [getPage](#getpage)
    * [updatePage](#updatepage)
    * [deletePage](#deletepage)
    * [getThemeLibrary](#getthemelibrary)
    * [addToThemeLibrary](#addtothemelibrary)
    * [applyTheme](#applytheme)
    * [isUpgradable](#isupgradable)
    * [upgradeTheme](#upgradetheme)
    * [getPublicThemes](#getpublicthemes)
    * [createTheme](#createtheme)
    * [getAppliedTheme](#getappliedtheme)
    * [getFonts](#getfonts)
    * [getThemeById](#getthemebyid)
    * [updateTheme](#updatetheme)
    * [deleteTheme](#deletetheme)
    * [getThemeForPreview](#getthemeforpreview)
    * [publishTheme](#publishtheme)
    * [unpublishTheme](#unpublishtheme)
    * [archiveTheme](#archivetheme)
    * [unarchiveTheme](#unarchivetheme)
    

* [User](#User)
  * Methods
    * [getCustomers](#getcustomers)
    * [searchUsers](#searchusers)
    * [createUser](#createuser)
    * [updateUser](#updateuser)
    * [createUserSession](#createusersession)
    * [getPlatformConfig](#getplatformconfig)
    * [updatePlatformConfig](#updateplatformconfig)
    

* [Content](#Content)
  * Methods
    * [getAnnouncementsList](#getannouncementslist)
    * [createAnnouncement](#createannouncement)
    * [getAnnouncementById](#getannouncementbyid)
    * [updateAnnouncement](#updateannouncement)
    * [updateAnnouncementSchedule](#updateannouncementschedule)
    * [deleteAnnouncement](#deleteannouncement)
    * [createBlog](#createblog)
    * [getBlogs](#getblogs)
    * [updateBlog](#updateblog)
    * [deleteBlog](#deleteblog)
    * [getComponentById](#getcomponentbyid)
    * [getFaqCategories](#getfaqcategories)
    * [getFaqCategoryBySlugOrId](#getfaqcategorybyslugorid)
    * [createFaqCategory](#createfaqcategory)
    * [updateFaqCategory](#updatefaqcategory)
    * [deleteFaqCategory](#deletefaqcategory)
    * [getFaqsByCategoryIdOrSlug](#getfaqsbycategoryidorslug)
    * [addFaq](#addfaq)
    * [updateFaq](#updatefaq)
    * [deleteFaq](#deletefaq)
    * [getFaqByIdOrSlug](#getfaqbyidorslug)
    * [getLandingPages](#getlandingpages)
    * [createLandingPage](#createlandingpage)
    * [updateLandingPage](#updatelandingpage)
    * [deleteLandingPage](#deletelandingpage)
    * [getLegalInformation](#getlegalinformation)
    * [updateLegalInformation](#updatelegalinformation)
    * [getNavigations](#getnavigations)
    * [createNavigation](#createnavigation)
    * [getDefaultNavigations](#getdefaultnavigations)
    * [getNavigationBySlug](#getnavigationbyslug)
    * [updateNavigation](#updatenavigation)
    * [deleteNavigation](#deletenavigation)
    * [getPageMeta](#getpagemeta)
    * [getPageSpec](#getpagespec)
    * [createPage](#createpage)
    * [getPages](#getpages)
    * [createPagePreview](#createpagepreview)
    * [updatePagePreview](#updatepagepreview)
    * [updatePage](#updatepage)
    * [deletePage](#deletepage)
    * [getPageBySlug](#getpagebyslug)
    * [getSEOConfiguration](#getseoconfiguration)
    * [updateSEOConfiguration](#updateseoconfiguration)
    * [getSlideshows](#getslideshows)
    * [createSlideshow](#createslideshow)
    * [getSlideshowBySlug](#getslideshowbyslug)
    * [updateSlideshow](#updateslideshow)
    * [deleteSlideshow](#deleteslideshow)
    * [getSupportInformation](#getsupportinformation)
    * [updateSupportInformation](#updatesupportinformation)
    * [updateInjectableTag](#updateinjectabletag)
    * [deleteAllInjectableTags](#deleteallinjectabletags)
    * [getInjectableTags](#getinjectabletags)
    * [addInjectableTag](#addinjectabletag)
    * [removeInjectableTag](#removeinjectabletag)
    * [editInjectableTag](#editinjectabletag)
    

* [Billing](#Billing)
  * Methods
    * [createSubscriptionCharge](#createsubscriptioncharge)
    * [getSubscriptionCharge](#getsubscriptioncharge)
    * [cancelSubscriptionCharge](#cancelsubscriptioncharge)
    * [getInvoices](#getinvoices)
    * [getInvoiceById](#getinvoicebyid)
    * [getCustomerDetail](#getcustomerdetail)
    * [upsertCustomerDetail](#upsertcustomerdetail)
    * [getSubscription](#getsubscription)
    * [getFeatureLimitConfig](#getfeaturelimitconfig)
    * [activateSubscriptionPlan](#activatesubscriptionplan)
    * [cancelSubscriptionPlan](#cancelsubscriptionplan)
    

* [Communication](#Communication)
  * Methods
    * [getCampaigns](#getcampaigns)
    * [createCampaign](#createcampaign)
    * [getCampaignById](#getcampaignbyid)
    * [updateCampaignById](#updatecampaignbyid)
    * [getStatsOfCampaignById](#getstatsofcampaignbyid)
    * [getAudiences](#getaudiences)
    * [createAudience](#createaudience)
    * [getBigqueryHeaders](#getbigqueryheaders)
    * [getAudienceById](#getaudiencebyid)
    * [updateAudienceById](#updateaudiencebyid)
    * [getNSampleRecordsFromCsv](#getnsamplerecordsfromcsv)
    * [getEmailProviders](#getemailproviders)
    * [createEmailProvider](#createemailprovider)
    * [getEmailProviderById](#getemailproviderbyid)
    * [updateEmailProviderById](#updateemailproviderbyid)
    * [getEmailTemplates](#getemailtemplates)
    * [createEmailTemplate](#createemailtemplate)
    * [getSystemEmailTemplates](#getsystememailtemplates)
    * [getEmailTemplateById](#getemailtemplatebyid)
    * [updateEmailTemplateById](#updateemailtemplatebyid)
    * [deleteEmailTemplateById](#deleteemailtemplatebyid)
    * [getEventSubscriptions](#geteventsubscriptions)
    * [getJobs](#getjobs)
    * [triggerCampaignJob](#triggercampaignjob)
    * [getJobLogs](#getjoblogs)
    * [getCommunicationLogs](#getcommunicationlogs)
    * [getSystemNotifications](#getsystemnotifications)
    * [getSmsProviders](#getsmsproviders)
    * [createSmsProvider](#createsmsprovider)
    * [getSmsProviderById](#getsmsproviderbyid)
    * [updateSmsProviderById](#updatesmsproviderbyid)
    * [getSmsTemplates](#getsmstemplates)
    * [createSmsTemplate](#createsmstemplate)
    * [getSmsTemplateById](#getsmstemplatebyid)
    * [updateSmsTemplateById](#updatesmstemplatebyid)
    * [deleteSmsTemplateById](#deletesmstemplatebyid)
    * [getSystemSystemTemplates](#getsystemsystemtemplates)
    

* [Payment](#Payment)
  * Methods
    * [getBrandPaymentGatewayConfig](#getbrandpaymentgatewayconfig)
    * [saveBrandPaymentGatewayConfig](#savebrandpaymentgatewayconfig)
    * [updateBrandPaymentGatewayConfig](#updatebrandpaymentgatewayconfig)
    * [getPaymentModeRoutes](#getpaymentmoderoutes)
    * [getAllPayouts](#getallpayouts)
    * [savePayout](#savepayout)
    * [updatePayout](#updatepayout)
    * [activateAndDectivatePayout](#activateanddectivatepayout)
    * [deletePayout](#deletepayout)
    * [getSubscriptionPaymentMethod](#getsubscriptionpaymentmethod)
    * [deleteSubscriptionPaymentMethod](#deletesubscriptionpaymentmethod)
    * [getSubscriptionConfig](#getsubscriptionconfig)
    * [saveSubscriptionSetupIntent](#savesubscriptionsetupintent)
    * [addBeneficiaryDetails](#addbeneficiarydetails)
    * [verifyIfscCode](#verifyifsccode)
    * [getUserOrderBeneficiaries](#getuserorderbeneficiaries)
    * [getUserBeneficiaries](#getuserbeneficiaries)
    * [confirmPayment](#confirmpayment)
    

* [Order](#Order)
  * Methods
    * [shipmentStatusUpdate](#shipmentstatusupdate)
    * [activityStatus](#activitystatus)
    * [storeProcessShipmentUpdate](#storeprocessshipmentupdate)
    * [checkRefund](#checkrefund)
    * [ShipmentBagsCanBreak](#shipmentbagscanbreak)
    * [getOrdersByCompanyId](#getordersbycompanyid)
    * [getOrderLanesCountByCompanyId](#getorderlanescountbycompanyid)
    * [getOrderDetails](#getorderdetails)
    * [getPicklistOrdersByCompanyId](#getpicklistordersbycompanyid)
    * [trackShipmentPlatform](#trackshipmentplatform)
    * [trackOrder](#trackorder)
    * [failedOrders](#failedorders)
    * [reprocessOrder](#reprocessorder)
    * [updateShipment](#updateshipment)
    * [getPlatformShipmentReasons](#getplatformshipmentreasons)
    * [getShipmentTrackDetails](#getshipmenttrackdetails)
    * [getShipmentAddress](#getshipmentaddress)
    * [updateShipmentAddress](#updateshipmentaddress)
    * [getPing](#getping)
    * [voiceCallback](#voicecallback)
    * [voiceClickToCall](#voiceclicktocall)
    

* [Catalog](#Catalog)
  * Methods
    * [updateSearchKeywords](#updatesearchkeywords)
    * [deleteSearchKeywords](#deletesearchkeywords)
    * [getSearchKeywords](#getsearchkeywords)
    * [createCustomKeyword](#createcustomkeyword)
    * [getAllSearchKeyword](#getallsearchkeyword)
    * [updateAutocompleteKeyword](#updateautocompletekeyword)
    * [deleteAutocompleteKeyword](#deleteautocompletekeyword)
    * [getAutocompleteKeywordDetail](#getautocompletekeyworddetail)
    * [createCustomAutocompleteRule](#createcustomautocompleterule)
    * [getAutocompleteConfig](#getautocompleteconfig)
    * [createProductBundle](#createproductbundle)
    * [getProductBundle](#getproductbundle)
    * [updateProductBundle](#updateproductbundle)
    * [getProductBundleDetail](#getproductbundledetail)
    * [createSizeGuide](#createsizeguide)
    * [getSizeGuides](#getsizeguides)
    * [updateSizeGuide](#updatesizeguide)
    * [getSizeGuide](#getsizeguide)
    * [getCatalogConfiguration](#getcatalogconfiguration)
    * [createConfigurationProductListing](#createconfigurationproductlisting)
    * [getConfigurations](#getconfigurations)
    * [createConfigurationByType](#createconfigurationbytype)
    * [getConfigurationByType](#getconfigurationbytype)
    * [getQueryFilters](#getqueryfilters)
    * [createCollection](#createcollection)
    * [getAllCollections](#getallcollections)
    * [getCollectionDetail](#getcollectiondetail)
    * [updateCollection](#updatecollection)
    * [deleteCollection](#deletecollection)
    * [addCollectionItems](#addcollectionitems)
    * [getCollectionItems](#getcollectionitems)
    * [getCatalogInsights](#getcataloginsights)
    * [getSellerInsights](#getsellerinsights)
    * [createMarketplaceOptin](#createmarketplaceoptin)
    * [getMarketplaceOptinDetail](#getmarketplaceoptindetail)
    * [getCompanyDetail](#getcompanydetail)
    * [getCompanyBrandDetail](#getcompanybranddetail)
    * [getCompanyMetrics](#getcompanymetrics)
    * [getStoreDetail](#getstoredetail)
    * [getGenderAttribute](#getgenderattribute)
    * [listProductTemplateCategories](#listproducttemplatecategories)
    * [listDepartmentsData](#listdepartmentsdata)
    * [getDepartmentData](#getdepartmentdata)
    * [listProductTemplate](#listproducttemplate)
    * [validateProductTemplate](#validateproducttemplate)
    * [downloadProductTemplateViews](#downloadproducttemplateviews)
    * [downloadProductTemplateView](#downloadproducttemplateview)
    * [validateProductTemplateSchema](#validateproducttemplateschema)
    * [listHSNCodes](#listhsncodes)
    * [listProductTemplateExportDetails](#listproducttemplateexportdetails)
    * [listTemplateBrandTypeValues](#listtemplatebrandtypevalues)
    * [createCategories](#createcategories)
    * [listCategories](#listcategories)
    * [updateCategory](#updatecategory)
    * [getCategoryData](#getcategorydata)
    * [createProduct](#createproduct)
    * [getProducts](#getproducts)
    * [editProduct](#editproduct)
    * [deleteProduct](#deleteproduct)
    * [getProduct](#getproduct)
    * [getProductValidation](#getproductvalidation)
    * [getProductSize](#getproductsize)
    * [updateProductAssetsInBulk](#updateproductassetsinbulk)
    * [getProductBulkUploadHistory](#getproductbulkuploadhistory)
    * [createProductsInBulk](#createproductsinbulk)
    * [deleteProductBulkJob](#deleteproductbulkjob)
    * [getProductTags](#getproducttags)
    * [createProductAssetsInBulk](#createproductassetsinbulk)
    * [getProductAssetsInBulk](#getproductassetsinbulk)
    * [deleteSize](#deletesize)
    * [addInventory](#addinventory)
    * [getInventoryBySize](#getinventorybysize)
    * [getInventoryBySizeIdentifier](#getinventorybysizeidentifier)
    * [deleteInventory](#deleteinventory)
    * [createBulkInventoryJob](#createbulkinventoryjob)
    * [getInventoryBulkUploadHistory](#getinventorybulkuploadhistory)
    * [createBulkInventory](#createbulkinventory)
    * [deleteBulkInventoryJob](#deletebulkinventoryjob)
    * [createInventoryExportJob](#createinventoryexportjob)
    * [getInventoryExport](#getinventoryexport)
    * [exportInventoryConfig](#exportinventoryconfig)
    * [createHsnCode](#createhsncode)
    * [getAllHsnCodes](#getallhsncodes)
    * [updateHsnCode](#updatehsncode)
    * [getHsnCode](#gethsncode)
    * [bulkHsnCode](#bulkhsncode)
    * [getApplicationBrands](#getapplicationbrands)
    * [getDepartments](#getdepartments)
    * [getCategories](#getcategories)
    * [getAppicationProducts](#getappicationproducts)
    * [getProductDetailBySlug](#getproductdetailbyslug)
    * [getAppProducts](#getappproducts)
    

* [CompanyProfile](#CompanyProfile)
  * Methods
    * [cbsOnboardGet](#cbsonboardget)
    * [updateCompany](#updatecompany)
    * [getCompanyMetrics](#getcompanymetrics)
    * [getBrand](#getbrand)
    * [editBrand](#editbrand)
    * [createBrand](#createbrand)
    * [getBrands](#getbrands)
    * [createCompanyBrandMapping](#createcompanybrandmapping)
    * [getLocations](#getlocations)
    * [createLocation](#createlocation)
    * [getLocationDetail](#getlocationdetail)
    * [updateLocation](#updatelocation)
    * [createLocationBulk](#createlocationbulk)
    

* [FileStorage](#FileStorage)
  * Methods
    * [startUpload](#startupload)
    * [completeUpload](#completeupload)
    * [appStartUpload](#appstartupload)
    * [appCompleteUpload](#appcompleteupload)
    * [getSignUrls](#getsignurls)
    * [copyFiles](#copyfiles)
    * [appCopyFiles](#appcopyfiles)
    * [browse](#browse)
    * [browse](#browse)
    * [proxy](#proxy)
    

* [Share](#Share)
  * Methods
    * [createShortLink](#createshortlink)
    * [getShortLinks](#getshortlinks)
    * [getShortLinkByHash](#getshortlinkbyhash)
    * [updateShortLinkById](#updateshortlinkbyid)
    

* [Inventory](#Inventory)
  * Methods
    * [getJobsByCompany](#getjobsbycompany)
    * [updateJob](#updatejob)
    * [createJob](#createjob)
    * [getJobByCompanyAndIntegration](#getjobbycompanyandintegration)
    * [getJobConfigDefaults](#getjobconfigdefaults)
    * [getJobByCode](#getjobbycode)
    * [getJobCodeMetrics](#getjobcodemetrics)
    * [getJobCodesByCompanyAndIntegration](#getjobcodesbycompanyandintegration)
    

* [Configuration](#Configuration)
  * Methods
    * [getBuildConfig](#getbuildconfig)
    * [updateBuildConfig](#updatebuildconfig)
    * [getPreviousVersions](#getpreviousversions)
    * [getAppFeatures](#getappfeatures)
    * [updateAppFeatures](#updateappfeatures)
    * [getAppBasicDetails](#getappbasicdetails)
    * [updateAppBasicDetails](#updateappbasicdetails)
    * [getAppContactInfo](#getappcontactinfo)
    * [updateAppContactInfo](#updateappcontactinfo)
    * [getAppApiTokens](#getappapitokens)
    * [updateAppApiTokens](#updateappapitokens)
    * [getAppCompanies](#getappcompanies)
    * [getAppStores](#getappstores)
    * [getInventoryConfig](#getinventoryconfig)
    * [updateInventoryConfig](#updateinventoryconfig)
    * [partiallyUpdateInventoryConfig](#partiallyupdateinventoryconfig)
    * [getAppCurrencyConfig](#getappcurrencyconfig)
    * [updateAppCurrencyConfig](#updateappcurrencyconfig)
    * [getOrderingStoresByFilter](#getorderingstoresbyfilter)
    * [updateOrderingStoreConfig](#updateorderingstoreconfig)
    * [getDomains](#getdomains)
    * [addDomain](#adddomain)
    * [removeDomainById](#removedomainbyid)
    * [changeDomainType](#changedomaintype)
    * [getDomainStatus](#getdomainstatus)
    * [createApplication](#createapplication)
    * [getApplications](#getapplications)
    * [getApplicationById](#getapplicationbyid)
    * [getCurrencies](#getcurrencies)
    * [getDomainAvailibility](#getdomainavailibility)
    * [getIntegrationById](#getintegrationbyid)
    * [getAvailableOptIns](#getavailableoptins)
    * [getSelectedOptIns](#getselectedoptins)
    * [getIntegrationLevelConfig](#getintegrationlevelconfig)
    * [getIntegrationByLevelId](#getintegrationbylevelid)
    * [getLevelActiveIntegrations](#getlevelactiveintegrations)
    * [getBrandsByCompany](#getbrandsbycompany)
    * [getCompanyByBrands](#getcompanybybrands)
    * [getStoreByBrands](#getstorebybrands)
    * [getOtherSellerApplications](#getothersellerapplications)
    * [getOtherSellerApplicationById](#getothersellerapplicationbyid)
    * [optOutFromApplication](#optoutfromapplication)
    

* [Cart](#Cart)
  * Methods
    * [getCoupons](#getcoupons)
    * [createCoupon](#createcoupon)
    * [getCouponById](#getcouponbyid)
    * [updateCoupon](#updatecoupon)
    * [updateCouponPartially](#updatecouponpartially)
    * [fetchAndvalidateCartItems](#fetchandvalidatecartitems)
    * [checkCartServiceability](#checkcartserviceability)
    * [checkoutCart](#checkoutcart)
    

* [Rewards](#Rewards)
  * Methods
    * [getGiveaways](#getgiveaways)
    * [createGiveaway](#creategiveaway)
    * [getGiveawayByID](#getgiveawaybyid)
    * [updateGiveaway](#updategiveaway)
    * [getOffers](#getoffers)
    * [getOfferByName](#getofferbyname)
    * [updateOfferByName](#updateofferbyname)
    * [getUserAvailablePoints](#getuseravailablepoints)
    * [updateUserStatus](#updateuserstatus)
    * [getUserPointsHistory](#getuserpointshistory)
    

* [Analytics](#Analytics)
  * Methods
    * [getStatiscticsGroups](#getstatiscticsgroups)
    * [getStatiscticsGroupComponents](#getstatiscticsgroupcomponents)
    * [getComponentStatsCSV](#getcomponentstatscsv)
    * [getComponentStatsPDF](#getcomponentstatspdf)
    * [getComponentStats](#getcomponentstats)
    * [getAbandonCartList](#getabandoncartlist)
    * [getAbandonCartsCSV](#getabandoncartscsv)
    * [getAbandonCartDetail](#getabandoncartdetail)
    * [createExportJob](#createexportjob)
    * [getExportJobStatus](#getexportjobstatus)
    * [getLogsList](#getlogslist)
    * [searchLogs](#searchlogs)
    

* [Discount](#Discount)
  * Methods
    * [getDiscounts](#getdiscounts)
    * [createDiscount](#creatediscount)
    * [getDiscount](#getdiscount)
    * [updateDiscount](#updatediscount)
    * [validateDiscountFile](#validatediscountfile)
    * [downloadDiscountFile](#downloaddiscountfile)
    * [getValidationJob](#getvalidationjob)
    * [cancelValidationJob](#cancelvalidationjob)
    * [getDownloadJob](#getdownloadjob)
    * [cancelDownloadJob](#canceldownloadjob)
    

* [Partner](#Partner)
  * Methods
    * [addProxyPath](#addproxypath)
    * [removeProxyPath](#removeproxypath)
    

* [Webhook](#Webhook)
  * Methods
    * [getSubscribersByCompany](#getsubscribersbycompany)
    * [registerSubscriberToEvent](#registersubscribertoevent)
    * [updateSubscriberConfig](#updatesubscriberconfig)
    * [getSubscribersByExtensionId](#getsubscribersbyextensionid)
    * [getSubscriberById](#getsubscriberbyid)
    * [fetchAllEventConfigurations](#fetchalleventconfigurations)
    


---
---



## Common


#### getLocations
Get countries, states, cities

```golang

data, err := Common.GetLocations(xQuery);
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


#### getTickets
Gets the list of company level tickets and/or ticket filters depending on query params

```golang

data, err := Lead.GetTickets(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID for which the data will be returned | 

















| xQuery | struct | Includes properties such as `Items`, `Filters`, `Q`, `Status`, `Priority`, `Category`, `PageNo`, `PageSize`


Gets the list of company level tickets and/or ticket filters

*Success Response:*



Success


Schema: `TicketList`


*Examples:*


Without items
```json
{
  "value": {
    "filters": {
      "statuses": [
        {
          "display": "Pending",
          "color": "#eae22b",
          "key": "pending"
        },
        {
          "display": "In Progress",
          "color": "#ffa951",
          "key": "in_progress"
        },
        {
          "display": "Resolved",
          "color": "#20c3a6",
          "key": "resolved"
        },
        {
          "display": "Closed",
          "color": "#41434c",
          "key": "closed"
        }
      ],
      "priorities": [
        {
          "display": "Low",
          "color": "#fed766",
          "key": "low"
        },
        {
          "display": "Medium",
          "color": "#f37736",
          "key": "medium"
        },
        {
          "display": "High",
          "color": "#fe4a49",
          "key": "high"
        }
      ],
      "assignees": [],
      "categories": [
        {
          "form": {
            "login_required": false,
            "should_notify": false,
            "inputs": [
              {
                "type": "email",
                "showRegexInput": false,
                "enum": [],
                "regex": "\\S+@\\S+\\.\\S+",
                "display": "email",
                "required": true,
                "key": "email"
              }
            ],
            "available_assignees": [],
            "_id": "602e900a2042255c03cadaf0",
            "title": "service-test-satyen",
            "description": "testing form from service",
            "slug": "service-test-satyen",
            "header_image": "https://hdn-1.addsale.com/x0/support-ticket/files/free/original/KZL86aN5l-service-test-satyen.jpeg",
            "application_id": "000000000000000000000001",
            "created_on": {
              "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36",
              "platform": "web",
              "meta": {
                "browser": {
                  "name": "Chrome",
                  "version": "88.0.4324.150"
                },
                "os": {
                  "name": "macOS",
                  "version": "11.2.0"
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
            "created_by": "5f8147abbd1a0a870f61f1a6",
            "createdAt": "2021-02-18T16:04:26.495Z",
            "updatedAt": "2021-02-18T16:04:26.495Z",
            "__v": 0
          },
          "key": "service-test-satyen",
          "display": "service-test-satyen"
        }
      ]
    }
  }
}
```

With items
```json
{
  "value": {
    "docs": [
      {
        "_id": "602d2652ce284d0b008d5c97",
        "status": {
          "display": "Pending",
          "color": "#eae22b",
          "key": "pending"
        },
        "priority": {
          "display": "Medium",
          "color": "#f37736",
          "key": "medium"
        },
        "assigned_to": {
          "agent_id": "5e79e721768c6bf54b783146",
          "gender": "male",
          "accountType": "user",
          "active": true,
          "profilePicUrl": "https://d2co8r51m5ca2d.cloudfront.net/inapp_banners/default_profile_img.png",
          "hasOldPasswordHash": false,
          "_id": "5e79e721768c6bf54b783146",
          "emails": [
            {
              "active": true,
              "primary": true,
              "verified": true,
              "email": "nikhilmshchs@gmail.com"
            },
            {
              "active": true,
              "primary": false,
              "verified": true,
              "email": "nikhilmanapure@gofynd.com"
            },
            {
              "active": true,
              "primary": false,
              "verified": true,
              "email": "nikhilmanapure@fynd.com"
            },
            {
              "active": true,
              "primary": false,
              "verified": true,
              "email": "nikhilmanapure@uniket.store"
            }
          ],
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
          "username": "nikhilmanapure_gofynd_com_29298",
          "createdAt": "2020-03-24T10:55:29.298Z",
          "updatedAt": "2020-05-12T07:46:41.816Z",
          "uid": "5567",
          "__v": 2
        },
        "tags": [
          "asdf444"
        ],
        "context": {
          "application_id": "000000000000000000000001",
          "company_id": "1"
        },
        "created_on": {
          "user_agent": "Fynd Platform/0.0.1 (com.fynd.platform; build:3; iOS 14.2.0) Alamofire/5.0.2",
          "platform": "web",
          "meta": {
            "browser": {
              "name": "Fynd Platform",
              "version": "0.0.1"
            }
          }
        },
        "source": "sales_channel",
        "content": {
          "title": "asdf444 Response",
          "description": "",
          "attachments": []
        },
        "response_id": "602d2652ce284dee3c8d5c96",
        "category": {
          "form": {
            "login_required": false,
            "should_notify": true,
            "inputs": [
              {
                "type": "text",
                "showRegexInput": false,
                "enum": [],
                "display": "asdf",
                "key": "asdf"
              },
              {
                "type": "mobile",
                "showRegexInput": false,
                "enum": [],
                "display": "mob num",
                "regex": "[0-9]{10}$",
                "key": "mob-num"
              }
            ],
            "available_assignees": [
              "5e79e721768c6bf54b783146"
            ],
            "_id": "60124e4a4d2bc363625e1bf4",
            "title": "asdf444",
            "description": "adf",
            "slug": "asdf444",
            "application_id": "000000000000000000000001",
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
            "created_by": "5e79e721768c6bf54b783146",
            "createdAt": "2021-01-28T05:40:26.271Z",
            "updatedAt": "2021-02-18T16:02:32.086Z",
            "__v": 0,
            "poll_for_assignment": {
              "duration": 20,
              "message": "We are looking for executive to connect you",
              "success_message": "Executive found",
              "failure_message": "All our executives are busy at the moment, We have accepted your request and someone will connect with you soon!"
            }
          },
          "key": "asdf444",
          "display": "asdf444"
        },
        "ticket_id": "472",
        "createdAt": "2021-02-17T14:21:06.774Z",
        "updatedAt": "2021-02-17T14:21:06.774Z",
        "__v": 0,
        "id": "602d2652ce284d0b008d5c97"
      }
    ],
    "total": 472,
    "limit": 10,
    "page": 1,
    "pages": 48,
    "filters": {
      "statuses": [
        {
          "display": "Pending",
          "color": "#eae22b",
          "key": "pending"
        },
        {
          "display": "In Progress",
          "color": "#ffa951",
          "key": "in_progress"
        },
        {
          "display": "Resolved",
          "color": "#20c3a6",
          "key": "resolved"
        },
        {
          "display": "Closed",
          "color": "#41434c",
          "key": "closed"
        }
      ],
      "priorities": [
        {
          "display": "Low",
          "color": "#fed766",
          "key": "low"
        },
        {
          "display": "Medium",
          "color": "#f37736",
          "key": "medium"
        },
        {
          "display": "High",
          "color": "#fe4a49",
          "key": "high"
        }
      ],
      "assignees": [],
      "categories": [
        {
          "form": {
            "login_required": false,
            "should_notify": false,
            "inputs": [
              {
                "type": "email",
                "showRegexInput": false,
                "enum": [],
                "regex": "\\S+@\\S+\\.\\S+",
                "display": "email",
                "required": true,
                "key": "email"
              }
            ],
            "available_assignees": [],
            "_id": "602e900a2042255c03cadaf0",
            "title": "service-test-satyen",
            "description": "testing form from service",
            "slug": "service-test-satyen",
            "header_image": "https://hdn-1.addsale.com/x0/support-ticket/files/free/original/KZL86aN5l-service-test-satyen.jpeg",
            "application_id": "000000000000000000000001",
            "created_on": {
              "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36",
              "platform": "web",
              "meta": {
                "browser": {
                  "name": "Chrome",
                  "version": "88.0.4324.150"
                },
                "os": {
                  "name": "macOS",
                  "version": "11.2.0"
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
            "created_by": "5f8147abbd1a0a870f61f1a6",
            "createdAt": "2021-02-18T16:04:26.495Z",
            "updatedAt": "2021-02-18T16:04:26.495Z",
            "__v": 0
          },
          "key": "service-test-satyen",
          "display": "service-test-satyen"
        }
      ]
    }
  }
}
```









---


#### createTicket
Creates a company level ticket

```golang

data, err := Lead.CreateTicket(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID for which the data will be returned | 


| body |  AddTicketPayload | "Request body" 

Creates a company level ticket

*Success Response:*



Success


Schema: `Ticket`


*Examples:*


Default
```json
{
  "value": {
    "context": {
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


#### getTickets
Gets the list of Application level Tickets and/or ticket filters depending on query params

```golang

data, err := Lead.GetTickets(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID of the application | 


| ApplicationID | string | Application ID for which the data will be returned | 













| xQuery | struct | Includes properties such as `Items`, `Filters`, `Q`, `Status`, `Priority`, `Category`


Gets the list of Application level Tickets and/or ticket filters

*Success Response:*



Success


Schema: `TicketList`


*Examples:*


Without items
```json
{
  "value": {
    "filters": {
      "statuses": [
        {
          "display": "Pending",
          "color": "#eae22b",
          "key": "pending"
        },
        {
          "display": "In Progress",
          "color": "#ffa951",
          "key": "in_progress"
        },
        {
          "display": "Resolved",
          "color": "#20c3a6",
          "key": "resolved"
        },
        {
          "display": "Closed",
          "color": "#41434c",
          "key": "closed"
        }
      ],
      "priorities": [
        {
          "display": "Low",
          "color": "#fed766",
          "key": "low"
        },
        {
          "display": "Medium",
          "color": "#f37736",
          "key": "medium"
        },
        {
          "display": "High",
          "color": "#fe4a49",
          "key": "high"
        }
      ],
      "assignees": [],
      "categories": [
        {
          "form": {
            "login_required": false,
            "should_notify": false,
            "inputs": [
              {
                "type": "email",
                "showRegexInput": false,
                "enum": [],
                "regex": "\\S+@\\S+\\.\\S+",
                "display": "email",
                "required": true,
                "key": "email"
              }
            ],
            "available_assignees": [],
            "_id": "602e900a2042255c03cadaf0",
            "title": "service-test-satyen",
            "description": "testing form from service",
            "slug": "service-test-satyen",
            "header_image": "https://hdn-1.addsale.com/x0/support-ticket/files/free/original/KZL86aN5l-service-test-satyen.jpeg",
            "application_id": "000000000000000000000001",
            "created_on": {
              "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36",
              "platform": "web",
              "meta": {
                "browser": {
                  "name": "Chrome",
                  "version": "88.0.4324.150"
                },
                "os": {
                  "name": "macOS",
                  "version": "11.2.0"
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
            "created_by": "5f8147abbd1a0a870f61f1a6",
            "createdAt": "2021-02-18T16:04:26.495Z",
            "updatedAt": "2021-02-18T16:04:26.495Z",
            "__v": 0
          },
          "key": "service-test-satyen",
          "display": "service-test-satyen"
        }
      ]
    }
  }
}
```

With items
```json
{
  "value": {
    "docs": [
      {
        "_id": "602d2652ce284d0b008d5c97",
        "status": {
          "display": "Pending",
          "color": "#eae22b",
          "key": "pending"
        },
        "priority": {
          "display": "Medium",
          "color": "#f37736",
          "key": "medium"
        },
        "assigned_to": {
          "agent_id": "5e79e721768c6bf54b783146",
          "gender": "male",
          "accountType": "user",
          "active": true,
          "profilePicUrl": "https://d2co8r51m5ca2d.cloudfront.net/inapp_banners/default_profile_img.png",
          "hasOldPasswordHash": false,
          "_id": "5e79e721768c6bf54b783146",
          "emails": [
            {
              "active": true,
              "primary": true,
              "verified": true,
              "email": "nikhilmshchs@gmail.com"
            },
            {
              "active": true,
              "primary": false,
              "verified": true,
              "email": "nikhilmanapure@gofynd.com"
            },
            {
              "active": true,
              "primary": false,
              "verified": true,
              "email": "nikhilmanapure@fynd.com"
            },
            {
              "active": true,
              "primary": false,
              "verified": true,
              "email": "nikhilmanapure@uniket.store"
            }
          ],
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
          "username": "nikhilmanapure_gofynd_com_29298",
          "createdAt": "2020-03-24T10:55:29.298Z",
          "updatedAt": "2020-05-12T07:46:41.816Z",
          "uid": "5567",
          "__v": 2
        },
        "tags": [
          "asdf444"
        ],
        "context": {
          "application_id": "000000000000000000000001",
          "company_id": "1"
        },
        "created_on": {
          "user_agent": "Fynd Platform/0.0.1 (com.fynd.platform; build:3; iOS 14.2.0) Alamofire/5.0.2",
          "platform": "web",
          "meta": {
            "browser": {
              "name": "Fynd Platform",
              "version": "0.0.1"
            }
          }
        },
        "source": "sales_channel",
        "content": {
          "title": "asdf444 Response",
          "description": "",
          "attachments": []
        },
        "response_id": "602d2652ce284dee3c8d5c96",
        "category": {
          "form": {
            "login_required": false,
            "should_notify": true,
            "inputs": [
              {
                "type": "text",
                "showRegexInput": false,
                "enum": [],
                "display": "asdf",
                "key": "asdf"
              },
              {
                "type": "mobile",
                "showRegexInput": false,
                "enum": [],
                "display": "mob num",
                "regex": "[0-9]{10}$",
                "key": "mob-num"
              }
            ],
            "available_assignees": [
              "5e79e721768c6bf54b783146"
            ],
            "_id": "60124e4a4d2bc363625e1bf4",
            "title": "asdf444",
            "description": "adf",
            "slug": "asdf444",
            "application_id": "000000000000000000000001",
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
            "created_by": "5e79e721768c6bf54b783146",
            "createdAt": "2021-01-28T05:40:26.271Z",
            "updatedAt": "2021-02-18T16:02:32.086Z",
            "__v": 0,
            "poll_for_assignment": {
              "duration": 20,
              "message": "We are looking for executive to connect you",
              "success_message": "Executive found",
              "failure_message": "All our executives are busy at the moment, We have accepted your request and someone will connect with you soon!"
            }
          },
          "key": "asdf444",
          "display": "asdf444"
        },
        "ticket_id": "472",
        "createdAt": "2021-02-17T14:21:06.774Z",
        "updatedAt": "2021-02-17T14:21:06.774Z",
        "__v": 0,
        "id": "602d2652ce284d0b008d5c97"
      }
    ],
    "total": 472,
    "limit": 10,
    "page": 1,
    "pages": 48,
    "filters": {
      "statuses": [
        {
          "display": "Pending",
          "color": "#eae22b",
          "key": "pending"
        },
        {
          "display": "In Progress",
          "color": "#ffa951",
          "key": "in_progress"
        },
        {
          "display": "Resolved",
          "color": "#20c3a6",
          "key": "resolved"
        },
        {
          "display": "Closed",
          "color": "#41434c",
          "key": "closed"
        }
      ],
      "priorities": [
        {
          "display": "Low",
          "color": "#fed766",
          "key": "low"
        },
        {
          "display": "Medium",
          "color": "#f37736",
          "key": "medium"
        },
        {
          "display": "High",
          "color": "#fe4a49",
          "key": "high"
        }
      ],
      "assignees": [],
      "categories": [
        {
          "form": {
            "login_required": false,
            "should_notify": false,
            "inputs": [
              {
                "type": "email",
                "showRegexInput": false,
                "enum": [],
                "regex": "\\S+@\\S+\\.\\S+",
                "display": "email",
                "required": true,
                "key": "email"
              }
            ],
            "available_assignees": [],
            "_id": "602e900a2042255c03cadaf0",
            "title": "service-test-satyen",
            "description": "testing form from service",
            "slug": "service-test-satyen",
            "header_image": "https://hdn-1.addsale.com/x0/support-ticket/files/free/original/KZL86aN5l-service-test-satyen.jpeg",
            "application_id": "000000000000000000000001",
            "created_on": {
              "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36",
              "platform": "web",
              "meta": {
                "browser": {
                  "name": "Chrome",
                  "version": "88.0.4324.150"
                },
                "os": {
                  "name": "macOS",
                  "version": "11.2.0"
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
            "created_by": "5f8147abbd1a0a870f61f1a6",
            "createdAt": "2021-02-18T16:04:26.495Z",
            "updatedAt": "2021-02-18T16:04:26.495Z",
            "__v": 0
          },
          "key": "service-test-satyen",
          "display": "service-test-satyen"
        }
      ]
    }
  }
}
```









---


#### getTicket
Retreives ticket details of a company level ticket with ticket ID

```golang

data, err := Lead.GetTicket(CompanyID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID for which the data will be returned | 


| ID | string | Tiket ID of the ticket to be fetched | 



Retreives ticket details of a company level ticket

*Success Response:*



Success


Schema: `Ticket`


*Examples:*


Default
```json
{
  "value": {
    "context": {
      "company_id": "1"
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


#### editTicket
Edits ticket details of a company level ticket

```golang

data, err := Lead.EditTicket(CompanyID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID for ticket | 


| ID | string | Ticket ID of ticket to be edited | 


| body |  EditTicketPayload | "Request body" 

Edits ticket details of a company level ticket such as status, priority, category, tags, attachments, assigne & ticket content changes

*Success Response:*



Success


Schema: `Ticket`


*Examples:*


Default
```json
{
  "value": {
    "context": {
      "company_id": "1"
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


#### getTicket
Retreives ticket details of a application level ticket

```golang

data, err := Lead.GetTicket(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID of the application | 


| ApplicationID | string | Application ID for which the data will be returned | 


| ID | string | Tiket ID of the ticket to be fetched | 



Retreives ticket details of a application level ticket with ticket ID

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


#### editTicket
Edits ticket details of a application level ticket

```golang

data, err := Lead.EditTicket(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID of the application | 


| ApplicationID | string | Application ID for ticket | 


| ID | string | Ticket ID of ticket to be edited | 


| body |  EditTicketPayload | "Request body" 

Edits ticket details of a application level ticket such as status, priority, category, tags, attachments, assigne & ticket content changes

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
Create history for specific company level ticket

```golang

data, err := Lead.CreateHistory(CompanyID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID for ticket | 


| ID | string | Ticket ID for which history is created | 


| body |  TicketHistoryPayload | "Request body" 

Create history for specific company level ticket, this history is seen on ticket detail page, this can be comment, log or rating.

*Success Response:*



Success


Schema: `TicketHistory`


*Examples:*


Default
```json
{
  "value": {
    "_id": "601a9d52c26687d086c499ef",
    "ticket_id": "43",
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


#### getTicketHistory
Gets history list for specific company level ticket

```golang

data, err := Lead.GetTicketHistory(CompanyID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID for ticket | 


| ID | string | Ticket ID for which history is to be fetched | 



Gets history list for specific company level ticket, this history is seen on ticket detail page, this can be comment, log or rating.

*Success Response:*



Success


Schema: `TicketHistoryList`


*Examples:*


Default
```json
{
  "value": {
    "docs": [
      {
        "_id": "602e5384204225eed5cadae7",
        "ticket_id": "41",
        "type": "comment",
        "value": {
          "text": "hello service",
          "media": []
        },
        "created_on": {
          "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36",
          "platform": "web",
          "meta": {
            "browser": {
              "name": "Chrome",
              "version": "88.0.4324.150"
            },
            "os": {
              "name": "macOS",
              "version": "11.2.0"
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
        "created_by": {
          "gender": "male",
          "accountType": "user",
          "active": true,
          "profilePicUrl": "https://hdn-1.fynd.com/company/884/applications/000000000000000000000001/theme/pictures/free/original/default-profile_nxhzui.png",
          "hasOldPasswordHash": false,
          "_id": "5f8147abbd1a0a870f61f1a6",
          "phoneNumbers": [
            {
              "active": true,
              "primary": true,
              "verified": true,
              "phone": "8412805281",
              "countryCode": 91
            }
          ],
          "firstName": "Satyen",
          "lastName": "Maurya",
          "emails": [
            {
              "active": true,
              "primary": true,
              "verified": true,
              "email": "satyenmaurya95@gmail.com"
            }
          ],
          "username": "satyenmaurya95_gmail_com_11118",
          "createdAt": "2020-10-10T05:33:31.119Z",
          "updatedAt": "2020-10-10T05:33:31.119Z",
          "uid": "5678",
          "__v": 0
        },
        "createdAt": "2021-02-18T11:46:12.522Z",
        "updatedAt": "2021-02-18T11:46:12.522Z",
        "__v": 0,
        "id": "602e5384204225eed5cadae7"
      },
      {
        "_id": "60372aa78a046d4d79c46e15",
        "ticket_id": "41",
        "type": "diff",
        "value": {
          "status": [
            "pending",
            "in_progress"
          ]
        },
        "created_by": {
          "gender": "male",
          "accountType": "user",
          "active": true,
          "profilePicUrl": "https://d2co8r51m5ca2d.cloudfront.net/inapp_banners/default_profile_img.png",
          "hasOldPasswordHash": false,
          "_id": "5e79e721768c6bf54b783146",
          "emails": [
            {
              "active": true,
              "primary": true,
              "verified": true,
              "email": "nikhilmshchs@gmail.com"
            },
            {
              "active": true,
              "primary": false,
              "verified": true,
              "email": "nikhilmanapure@gofynd.com"
            },
            {
              "active": true,
              "primary": false,
              "verified": true,
              "email": "nikhilmanapure@fynd.com"
            },
            {
              "active": true,
              "primary": false,
              "verified": true,
              "email": "nikhilmanapure@uniket.store"
            }
          ],
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
          "username": "nikhilmanapure_gofynd_com_29298",
          "createdAt": "2020-03-24T10:55:29.298Z",
          "updatedAt": "2020-05-12T07:46:41.816Z",
          "uid": "5567",
          "__v": 2
        },
        "createdAt": "2021-02-25T04:42:15.225Z",
        "updatedAt": "2021-02-25T04:42:15.225Z",
        "__v": 0,
        "id": "60372aa78a046d4d79c46e15"
      }
    ],
    "total": 2,
    "limit": 100,
    "page": 1,
    "pages": 1
  }
}
```









---


#### getFeedbacks
Gets a list of feedback submitted against that ticket

```golang

data, err := Lead.GetFeedbacks(CompanyID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID for ticket | 


| ID | string | Ticket ID for which feedbacks are to be fetched | 



Gets a list of feedback submitted against that ticket

*Success Response:*



Success


Schema: `TicketFeedbackList`


*Examples:*


Default
```json
{
  "value": {
    "items": [
      {
        "_id": "60c255bf00ecabfad19e9601",
        "company_id": "1",
        "ticket_id": "6095812876d2bf17143cb3b3",
        "user": {
          "_id": "5f8147abbd1a0a870f61f1a6",
          "authenticated": true,
          "user_id": "5f8147abbd1a0a870f61f1a6"
        },
        "category": "customers",
        "response": {
          "audio": 2,
          "video": 6
        },
        "createdAt": "2021-06-10T18:11:11.349Z",
        "updatedAt": "2021-06-10T18:11:11.349Z",
        "__v": 0
      }
    ]
  }
}
```









---


#### submitFeedback
Submit a response for feeback form against that ticket

```golang

data, err := Lead.SubmitFeedback(CompanyID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID for ticket | 


| ID | string | Ticket ID for which feedback is to be submitted | 


| body |  TicketFeedbackPayload | "Request body" 

Submit a response for feeback form against that ticket

*Success Response:*



Success


Schema: `TicketFeedback`


*Examples:*


Default
```json
{
  "value": {
    "_id": "60c255bf00ecabfad19e9601",
    "company_id": "1",
    "ticket_id": "6095812876d2bf17143cb3b3",
    "user": {
      "_id": "5f8147abbd1a0a870f61f1a6",
      "authenticated": true,
      "user_id": "5f8147abbd1a0a870f61f1a6"
    },
    "category": "customers",
    "response": {
      "audio": 2,
      "video": 6
    },
    "createdAt": "2021-06-10T18:11:11.349Z",
    "updatedAt": "2021-06-10T18:11:11.349Z",
    "__v": 0
  }
}
```









---


#### createHistory
Create history for specific application level ticket

```golang

data, err := Lead.CreateHistory(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID of the application | 


| ApplicationID | string | Application ID for ticket | 


| ID | string | Ticket ID for which history is created | 


| body |  TicketHistoryPayload | "Request body" 

Create history for specific application level ticket, this history is seen on ticket detail page, this can be comment, log or rating.

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


#### getTicketHistory
Gets history list for specific application level ticket

```golang

data, err := Lead.GetTicketHistory(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID of application | 


| ApplicationID | string | Application ID for ticket | 


| ID | string | Ticket ID for which history is to be fetched | 



Gets history list for specific application level ticket, this history is seen on ticket detail page, this can be comment, log or rating.

*Success Response:*



Success


Schema: `TicketHistoryList`


*Examples:*


Default
```json
{
  "value": {
    "docs": [
      {
        "_id": "602e5384204225eed5cadae7",
        "ticket_id": "41",
        "type": "comment",
        "value": {
          "text": "hello service",
          "media": []
        },
        "created_on": {
          "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36",
          "platform": "web",
          "meta": {
            "browser": {
              "name": "Chrome",
              "version": "88.0.4324.150"
            },
            "os": {
              "name": "macOS",
              "version": "11.2.0"
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
        "created_by": {
          "gender": "male",
          "accountType": "user",
          "active": true,
          "profilePicUrl": "https://hdn-1.fynd.com/company/884/applications/000000000000000000000001/theme/pictures/free/original/default-profile_nxhzui.png",
          "hasOldPasswordHash": false,
          "_id": "5f8147abbd1a0a870f61f1a6",
          "phoneNumbers": [
            {
              "active": true,
              "primary": true,
              "verified": true,
              "phone": "8412805281",
              "countryCode": 91
            }
          ],
          "firstName": "Satyen",
          "lastName": "Maurya",
          "emails": [
            {
              "active": true,
              "primary": true,
              "verified": true,
              "email": "satyenmaurya95@gmail.com"
            }
          ],
          "username": "satyenmaurya95_gmail_com_11118",
          "createdAt": "2020-10-10T05:33:31.119Z",
          "updatedAt": "2020-10-10T05:33:31.119Z",
          "uid": "5678",
          "__v": 0
        },
        "createdAt": "2021-02-18T11:46:12.522Z",
        "updatedAt": "2021-02-18T11:46:12.522Z",
        "__v": 0,
        "id": "602e5384204225eed5cadae7"
      },
      {
        "_id": "60372aa78a046d4d79c46e15",
        "ticket_id": "41",
        "type": "diff",
        "value": {
          "status": [
            "pending",
            "in_progress"
          ]
        },
        "created_by": {
          "gender": "male",
          "accountType": "user",
          "active": true,
          "profilePicUrl": "https://d2co8r51m5ca2d.cloudfront.net/inapp_banners/default_profile_img.png",
          "hasOldPasswordHash": false,
          "_id": "5e79e721768c6bf54b783146",
          "emails": [
            {
              "active": true,
              "primary": true,
              "verified": true,
              "email": "nikhilmshchs@gmail.com"
            },
            {
              "active": true,
              "primary": false,
              "verified": true,
              "email": "nikhilmanapure@gofynd.com"
            },
            {
              "active": true,
              "primary": false,
              "verified": true,
              "email": "nikhilmanapure@fynd.com"
            },
            {
              "active": true,
              "primary": false,
              "verified": true,
              "email": "nikhilmanapure@uniket.store"
            }
          ],
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
          "username": "nikhilmanapure_gofynd_com_29298",
          "createdAt": "2020-03-24T10:55:29.298Z",
          "updatedAt": "2020-05-12T07:46:41.816Z",
          "uid": "5567",
          "__v": 2
        },
        "createdAt": "2021-02-25T04:42:15.225Z",
        "updatedAt": "2021-02-25T04:42:15.225Z",
        "__v": 0,
        "id": "60372aa78a046d4d79c46e15"
      }
    ],
    "total": 2,
    "limit": 100,
    "page": 1,
    "pages": 1
  }
}
```









---


#### getCustomForm
Get specific custom form using it's slug

```golang

data, err := Lead.GetCustomForm(CompanyID, ApplicationID, Slug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID of the application | 


| ApplicationID | string | Application ID for the form | 


| Slug | string | Slug of form whose response is getting submitted | 



Get specific custom form using it's slug, this is used to view the form.

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


#### editCustomForm
Edit the given custom form

```golang

data, err := Lead.EditCustomForm(CompanyID, ApplicationID, Slug, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID of the application | 


| ApplicationID | string | Application ID for the form | 


| Slug | string | Slug of form whose response is getting submitted | 


| body |  EditCustomFormPayload | "Request body" 

Edit the given custom form field such as adding or deleting input, assignee, title, decription, notification and polling information.

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
        "type": "email",
        "showRegexInput": true,
        "enum": [],
        "regex": "\\S+@\\S+\\.\\S+",
        "display": "email",
        "required": true,
        "key": "email"
      },
      {
        "type": "number",
        "showRegexInput": false,
        "enum": [],
        "display": "Enter your fav number",
        "placeholder": "123",
        "key": "enter-your-fav-number"
      }
    ],
    "available_assignees": [],
    "_id": "602e900a2042255c03cadaf0",
    "title": "service-test-satyen",
    "description": "testing form from service",
    "slug": "service-test-satyen",
    "header_image": "https://hdn-1.addsale.com/x0/support-ticket/files/free/original/KZL86aN5l-service-test-satyen.jpeg",
    "application_id": "000000000000000000000001",
    "created_on": {
      "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36",
      "platform": "web",
      "meta": {
        "browser": {
          "name": "Chrome",
          "version": "88.0.4324.150"
        },
        "os": {
          "name": "macOS",
          "version": "11.2.0"
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
    "created_by": "5f8147abbd1a0a870f61f1a6",
    "createdAt": "2021-02-18T16:04:26.495Z",
    "updatedAt": "2021-02-26T10:16:49.272Z",
    "__v": 0
  }
}
```









---


#### getCustomForms
Get list of custom form

```golang

data, err := Lead.GetCustomForms(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID of the application | 


| ApplicationID | string | Application ID for the form | 



Get list of custom form for given application

*Success Response:*



Success


Schema: `CustomFormList`


*Examples:*


Default
```json
{
  "value": {
    "docs": [
      {
        "_id": "602e900a2042255c03cadaf0",
        "login_required": false,
        "should_notify": false,
        "inputs": [
          {
            "type": "email",
            "showRegexInput": true,
            "enum": [],
            "regex": "\\S+@\\S+\\.\\S+",
            "display": "email",
            "required": true,
            "key": "email"
          },
          {
            "type": "number",
            "showRegexInput": false,
            "enum": [],
            "display": "Enter your fav number",
            "placeholder": "123",
            "key": "enter-your-fav-number"
          },
          {
            "type": "textarea",
            "showRegexInput": false,
            "enum": [],
            "display": "kjhgjhvjb",
            "required": true,
            "key": "kjhgjhvjb"
          }
        ],
        "available_assignees": [],
        "title": "service-test-satyen",
        "description": "testing form from service",
        "slug": "service-test-satyen",
        "header_image": "https://hdn-1.addsale.com/x0/support-ticket/files/free/original/KZL86aN5l-service-test-satyen.jpeg",
        "application_id": "000000000000000000000001",
        "created_on": {
          "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36",
          "platform": "web",
          "meta": {
            "browser": {
              "name": "Chrome",
              "version": "88.0.4324.150"
            },
            "os": {
              "name": "macOS",
              "version": "11.2.0"
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
        "created_by": "5f8147abbd1a0a870f61f1a6",
        "createdAt": "2021-02-18T16:04:26.495Z",
        "updatedAt": "2021-02-24T15:49:56.256Z",
        "__v": 0,
        "id": "602e900a2042255c03cadaf0"
      },
      {
        "_id": "60124e4a4d2bc363625e1bf4",
        "login_required": false,
        "should_notify": true,
        "inputs": [
          {
            "type": "text",
            "showRegexInput": false,
            "enum": [],
            "display": "asdf",
            "key": "asdf"
          },
          {
            "type": "mobile",
            "showRegexInput": false,
            "enum": [],
            "display": "mob num",
            "regex": "[0-9]{10}$",
            "key": "mob-num"
          }
        ],
        "available_assignees": [
          "5e79e721768c6bf54b783146"
        ],
        "title": "asdf444",
        "description": "adf",
        "slug": "asdf444",
        "application_id": "000000000000000000000001",
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
        "created_by": "5e79e721768c6bf54b783146",
        "createdAt": "2021-01-28T05:40:26.271Z",
        "updatedAt": "2021-02-18T16:02:32.086Z",
        "__v": 0,
        "poll_for_assignment": {
          "duration": 20,
          "message": "We are looking for executive to connect you",
          "success_message": "Executive found",
          "failure_message": "All our executives are busy at the moment, We have accepted your request and someone will connect with you soon!"
        },
        "id": "60124e4a4d2bc363625e1bf4"
      }
    ],
    "total": 2,
    "limit": 10,
    "page": 1,
    "pages": 1
  }
}
```









---


#### createCustomForm
Creates a new custom form

```golang

data, err := Lead.CreateCustomForm(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID of the application | 


| ApplicationID | string | Application ID for the form | 


| body |  CreateCustomFormPayload | "Request body" 

Creates a new custom form for given application

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
        "type": "email",
        "showRegexInput": true,
        "enum": [],
        "regex": "\\S+@\\S+\\.\\S+",
        "display": "email",
        "required": true,
        "key": "email"
      },
      {
        "type": "number",
        "showRegexInput": false,
        "enum": [],
        "display": "Enter your fav number",
        "placeholder": "123",
        "key": "enter-your-fav-number"
      }
    ],
    "available_assignees": [],
    "_id": "602e900a2042255c03cadaf0",
    "title": "service-test-satyen",
    "description": "testing form from service",
    "slug": "service-test-satyen",
    "header_image": "https://hdn-1.addsale.com/x0/support-ticket/files/free/original/KZL86aN5l-service-test-satyen.jpeg",
    "application_id": "000000000000000000000001",
    "created_on": {
      "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36",
      "platform": "web",
      "meta": {
        "browser": {
          "name": "Chrome",
          "version": "88.0.4324.150"
        },
        "os": {
          "name": "macOS",
          "version": "11.2.0"
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
    "created_by": "5f8147abbd1a0a870f61f1a6",
    "createdAt": "2021-02-18T16:04:26.495Z",
    "updatedAt": "2021-02-26T10:16:49.272Z",
    "__v": 0
  }
}
```









---


#### getTokenForVideoRoom
Get Token to join a specific Video Room using it's unqiue name

```golang

data, err := Lead.GetTokenForVideoRoom(CompanyID, UniqueName);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id for video room | 


| UniqueName | string | Unique name of video room | 



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


#### getTokenForVideoRoom
Get Token to join a specific Video Room using it's unqiue name

```golang

data, err := Lead.GetTokenForVideoRoom(CompanyID, ApplicationID, UniqueName);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID of the application | 


| ApplicationID | string | Application ID for video room | 


| UniqueName | string | Unique name of video room | 



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


#### getVideoParticipants
Get participants of a specific Video Room using it's unique name

```golang

data, err := Lead.GetVideoParticipants(CompanyID, UniqueName);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id for video room | 


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


#### getVideoParticipants
Get participants of a specific Video Room using it's unique name

```golang

data, err := Lead.GetVideoParticipants(CompanyID, ApplicationID, UniqueName);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID of the application | 


| ApplicationID | string | Application ID for video room | 


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


#### openVideoRoom
Open a video room.

```golang

data, err := Lead.OpenVideoRoom(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID of the application | 


| ApplicationID | string | Application ID for video room | 


| body |  CreateVideoRoomPayload | "Request body" 

Open a video room.

*Success Response:*



Success


Schema: `CreateVideoRoomResponse`


*Examples:*


Default
```json
{
  "value": {
    "unique_name": "alphanumeric123"
  }
}
```









---


#### closeVideoRoom
Close the video room and force all participants to leave.

```golang

data, err := Lead.CloseVideoRoom(CompanyID, ApplicationID, UniqueName);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID of the application | 


| ApplicationID | string | Application ID for video room | 


| UniqueName | string | Unique name of Video Room | 



Close the video room and force all participants to leave.

*Success Response:*



Success


Schema: `CloseVideoRoomResponse`


*Examples:*


Default
```json
{
  "value": {
    "success": true
  }
}
```









---



---


## Feedback


#### getAttributes
Get list of attribute data

```golang

data, err := Feedback.GetAttributes(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| ApplicationID | string | application id | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`


Provides a list of all attribute data.

*Success Response:*



ok


Schema: `FeedbackAttributes`









---


#### getCustomerReviews
Get list of customer reviews [admin]

```golang

data, err := Feedback.GetCustomerReviews(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| ApplicationID | string | application id | 































| xQuery | struct | Includes properties such as `ID`, `EntityID`, `EntityType`, `UserID`, `Media`, `Rating`, `AttributeRating`, `Facets`, `Sort`, `Next`, `Start`, `Limit`, `Count`, `PageID`, `PageSize`


The endpoint provides a list of customer reviews based on entity and provided filters

*Success Response:*



Success


Schema: `GetReviewResponse`









---


#### updateApprove
update approve details

```golang

data, err := Feedback.UpdateApprove(CompanyID, ApplicationID, ReviewID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| ApplicationID | string | application id | 


| ReviewID | string | review id | 


| body |  ApproveRequest | "Request body" 

The is used to update approve details like status and description text

*Success Response:*



ok


Schema: `UpdateResponse`









---


#### getHistory
get history details

```golang

data, err := Feedback.GetHistory(CompanyID, ApplicationID, ReviewID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| ApplicationID | string | application id | 


| ReviewID | string | review id | 



The is used to get the history details like status and description text

*Success Response:*



ok


Schema: `Array<ActivityDump>`









---


#### getApplicationTemplates
Get list of templates

```golang

data, err := Feedback.GetApplicationTemplates(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| ApplicationID | string | application id | 





| xQuery | struct | Includes properties such as `PageID`, `PageSize`


Get all templates of application

*Success Response:*



Success


Schema: `TemplateGetResponse`









---


#### createTemplate
Create a new template

```golang

data, err := Feedback.CreateTemplate(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| ApplicationID | string | application id | 


| body |  TemplateRequestList | "Request body" 

Create a new template for review with following data:
- Enable media, rating and review
- Rating - active/inactive/selected rate choices, attributes, text on rate, comment for each rate, type
- Review - header, title, description, image and video meta, enable votes

*Success Response:*



Success


Schema: `InsertResponse`









---


#### getTemplateById
Get a template by ID

```golang

data, err := Feedback.GetTemplateById(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| ApplicationID | string | application id | 


| ID | string | template id | 



Get the template for product or l3 type by ID

*Success Response:*



Success


Schema: `Template`









---


#### updateTemplate
Update a template's status

```golang

data, err := Feedback.UpdateTemplate(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| ApplicationID | string | application id | 


| ID | string | template id | 


| body |  UpdateTemplateRequest | "Request body" 

Update existing template status, active/archive

*Success Response:*



Success


Schema: `UpdateResponse`









---


#### updateTemplateStatus
Update a template's status

```golang

data, err := Feedback.UpdateTemplateStatus(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| ApplicationID | string | application id | 


| ID | string | template id | 


| body |  UpdateTemplateStatusRequest | "Request body" 

Update existing template status, active/archive

*Success Response:*



Success


Schema: `UpdateResponse`









---



---


## Theme


#### getAllPages
Get all pages of a theme

```golang

data, err := Theme.GetAllPages(CompanyID, ApplicationID, ThemeID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID | 


| ApplicationID | string | Application ID | 


| ThemeID | string | ID of the theme to be retrieved | 



Use this API to retrieve all the available pages of a theme by its ID.

*Success Response:*



Success. Returns an array all the pages of the theme. Refer `AllAvailablePageSchema` for more details.


Schema: `AllAvailablePageSchema`


*Examples:*


All pages
```json
{
  "$ref": "#/components/examples/AllAvailablePagesExample"
}
```









---


#### createPage
Create a page 

```golang

data, err := Theme.CreatePage(CompanyID, ApplicationID, ThemeID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID | 


| ApplicationID | string | Application ID | 


| ThemeID | string | ID of the theme | 


| body |  AvailablePageSchema | "Request body" 

Use this API to create a page for a theme by its ID.

*Success Response:*



Success. Returns the page of the theme. Refer `AvailablePageSchema` for more details.


Schema: `AvailablePageSchema`


*Examples:*


page
```json
{
  "$ref": "#/components/examples/AvailablePageExample"
}
```









---


#### updateMultiplePages
Update multiple pages of a theme

```golang

data, err := Theme.UpdateMultiplePages(CompanyID, ApplicationID, ThemeID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID | 


| ApplicationID | string | Application ID | 


| ThemeID | string | ID of the theme to be retrieved | 


| body |  AllAvailablePageSchema | "Request body" 

Use this API to update multiple pages of a theme by its ID.

*Success Response:*



Success. Returns an array all the pages of the theme. Refer `AllAvailablePageSchema` for more details.


Schema: `AllAvailablePageSchema`


*Examples:*


All pages
```json
{
  "$ref": "#/components/examples/AllAvailablePagesExample"
}
```









---


#### getPage
Get page of a theme

```golang

data, err := Theme.GetPage(CompanyID, ApplicationID, ThemeID, PageValue);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID | 


| ApplicationID | string | Application ID | 


| ThemeID | string | ID of the theme to be retrieved | 


| PageValue | string | Value of the page to be retrieved | 



Use this API to retrieve a page of a theme.

*Success Response:*



Success. Returns an object of the page.  Refer `AvailablePageSchema` for more details.


Schema: `AvailablePageSchema`


*Examples:*


Home page
```json
{
  "$ref": "#/components/examples/AvailablePageExample"
}
```









---


#### updatePage
Updates a page 

```golang

data, err := Theme.UpdatePage(CompanyID, ApplicationID, ThemeID, PageValue, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID | 


| ApplicationID | string | Application ID | 


| ThemeID | string | ID of the theme | 


| PageValue | string | Value of the page to be updated | 


| body |  AvailablePageSchema | "Request body" 

Use this API to update a page for a theme by its ID.

*Success Response:*



Success. Returns a the page of the theme. Refer `AvailablePageSchema` for more details.


Schema: `AvailablePageSchema`


*Examples:*


page
```json
{
  "$ref": "#/components/examples/AvailablePageExample"
}
```









---


#### deletePage
Deletes a page 

```golang

data, err := Theme.DeletePage(CompanyID, ApplicationID, ThemeID, PageValue);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID | 


| ApplicationID | string | Application ID | 


| ThemeID | string | ID of the theme | 


| PageValue | string | Value of the page to be updated | 



Use this API to delete a page for a theme by its ID and page_value.

*Success Response:*



Success. Returns a the page of the theme. Refer `AvailablePageSchema` for more details.


Schema: `AvailablePageSchema`


*Examples:*


page
```json
{
  "$ref": "#/components/examples/AvailablePageExample"
}
```









---


#### getThemeLibrary
Get a list of themes from the theme library

```golang

data, err := Theme.GetThemeLibrary(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 





| xQuery | struct | Includes properties such as `PageSize`, `PageNo`


Theme library is a personalized collection of themes that are chosen and added from the available themes. Use this API to fetch a list of themes from the library along with their configuration details. 

*Success Response:*



Success. Refer `ThemesListingResponseSchema` for more details.


Schema: `ThemesListingResponseSchema`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/ThemesListingResponse"
}
```









---


#### addToThemeLibrary
Add a theme to the theme library

```golang

data, err := Theme.AddToThemeLibrary(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 


| body |  AddThemeRequestSchema | "Request body" 

Theme library is a personalized collection of themes that are chosen and added from the available themes. Use this API to choose a theme and add it to the theme library.

*Success Response:*



Success. Refer `ThemesSchema` for more details.


Schema: `ThemesSchema`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/Themes"
}
```









---


#### applyTheme
Apply a theme

```golang

data, err := Theme.ApplyTheme(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 


| body |  AddThemeRequestSchema | "Request body" 

Use this API to apply a theme to the website.

*Success Response:*



Success. Refer `ThemesSchema` for more details.


Schema: `ThemesSchema`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/Themes"
}
```









---


#### isUpgradable
Checks if theme is upgradable

```golang

data, err := Theme.IsUpgradable(CompanyID, ApplicationID, ThemeID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 


| ThemeID | string | Theme ID | 



There's always a possibility that new features get added to a theme. Use this API to check if the applied theme has an upgrade available.

*Success Response:*



Success. If the boolean value of `upgrade` returns **true**, the theme can be upgraded. Refer `UpgradableThemeSchema` for more details.


Schema: `UpgradableThemeSchema`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/UpgradableTheme"
}
```









---


#### upgradeTheme
Upgrade a theme

```golang

data, err := Theme.UpgradeTheme(CompanyID, ApplicationID, ThemeID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 


| ThemeID | string | ID allotted to the theme. | 



Use this API to upgrade the current theme to its latest version.

*Success Response:*



Success. Upgrades the theme and shares the details of the new version in the response. Refer `ThemesSchema` for more details.


Schema: `ThemesSchema`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/Themes"
}
```









---


#### getPublicThemes
Get all public themes

```golang

data, err := Theme.GetPublicThemes(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 





| xQuery | struct | Includes properties such as `PageSize`, `PageNo`


Use this API to get a list of free themes that you can apply to your website.

*Success Response:*



Success. Refer `ThemesListingResponseSchema` for more details.


Schema: `ThemesListingResponseSchema`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/ThemesListingResponse"
}
```









---


#### createTheme
Create a new theme

```golang

data, err := Theme.CreateTheme(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 


| body |  ThemesSchema | "Request body" 

Themes improve the look and appearance of a website. Use this API to create a theme.

*Success Response:*



Theme


Schema: `ThemesSchema`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/Themes"
}
```









---


#### getAppliedTheme
Get the applied theme

```golang

data, err := Theme.GetAppliedTheme(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 



Use this API to retrieve the theme that is currently applied to the website along with its details.

*Success Response:*



Success. Refer `ThemesSchema` for more details.


Schema: `ThemesSchema`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/Themes"
}
```









---


#### getFonts
Get all the supported fonts in a theme

```golang

data, err := Theme.GetFonts(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 



Font is a collection of characters with a similar design. Use this API to retrieve a list of website fonts.

*Success Response:*



Success. Refer `FontsSchema` for more details.


Schema: `FontsSchema`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/FontsResponse"
}
```









---


#### getThemeById
Gets theme by id

```golang

data, err := Theme.GetThemeById(CompanyID, ApplicationID, ThemeID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 


| ThemeID | string | ID allotted to the theme. | 



Use this API to retrieve the details of a specific theme by using its ID.

*Success Response:*



Success. Refer `ThemesSchema` for more details.


Schema: `ThemesSchema`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/Themes"
}
```









---


#### updateTheme
Update a theme

```golang

data, err := Theme.UpdateTheme(CompanyID, ApplicationID, ThemeID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 


| ThemeID | string | ID allotted to the theme. | 


| body |  ThemesSchema | "Request body" 

Use this API to edit an existing theme. You can customize the website font, sections, images, styles, and many more.

*Success Response:*



Success. Refer `ThemesSchema` for more details.


Schema: `ThemesSchema`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/Themes"
}
```









---


#### deleteTheme
Delete a theme

```golang

data, err := Theme.DeleteTheme(CompanyID, ApplicationID, ThemeID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 


| ThemeID | string | ID allotted to the theme. | 



Use this API to delete a theme from the theme library.

*Success Response:*



Success. Refer `ThemesSchema` for more details.


Schema: `ThemesSchema`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/Themes"
}
```









---


#### getThemeForPreview
Get a theme preview

```golang

data, err := Theme.GetThemeForPreview(CompanyID, ApplicationID, ThemeID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 


| ThemeID | string | ID allotted to the theme. | 



A theme can be previewed before applying it. Use this API to retrieve the theme preview by using its ID.

*Success Response:*



Success. Refer `ThemesSchema` for more details.


Schema: `ThemesSchema`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/Themes"
}
```









---


#### publishTheme
Publish a theme

```golang

data, err := Theme.PublishTheme(CompanyID, ApplicationID, ThemeID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 


| ThemeID | string | ID allotted to the theme. | 



Use this API to publish a theme that is either newly created or edited.

*Success Response:*



Success. Refer `ThemesSchema` for more details.


Schema: `ThemesSchema`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/Themes"
}
```









---


#### unpublishTheme
Unpublish a theme

```golang

data, err := Theme.UnpublishTheme(CompanyID, ApplicationID, ThemeID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 


| ThemeID | string | ID allotted to the theme. | 



Use this API to remove an existing theme from the list of available themes.

*Success Response:*



Success. Refer `ThemesSchema` for more details.


Schema: `ThemesSchema`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/Themes"
}
```









---


#### archiveTheme
Archive a theme

```golang

data, err := Theme.ArchiveTheme(CompanyID, ApplicationID, ThemeID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 


| ThemeID | string | ID allotted to the theme. | 



Use this API to store an existing theme but not delete it so that it can be used in future if required. 

*Success Response:*



Success. Refer `ThemesSchema` for more details.


Schema: `ThemesSchema`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/Themes"
}
```









---


#### unarchiveTheme
Unarchive a theme

```golang

data, err := Theme.UnarchiveTheme(CompanyID, ApplicationID, ThemeID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 


| ThemeID | string | ID allotted to the theme. | 



Use this API to restore an archived theme and bring it back for editing or publishing. 

*Success Response:*



Success. Refer `ThemesSchema` for more details.


Schema: `ThemesSchema`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/Themes"
}
```









---



---


## User


#### getCustomers
Get a list of customers

```golang

data, err := User.GetCustomers(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 







| xQuery | struct | Includes properties such as `Q`, `PageSize`, `PageNo`


Use this API to retrieve a list of customers who have registered in the application.

*Success Response:*



Success. Refer `CustomerListResponseSchema` for more details.


Schema: `CustomerListResponseSchema`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/CustomersListResponse"
}
```









---


#### searchUsers
Search an existing user.

```golang

data, err := User.SearchUsers(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 



| xQuery | struct | Includes properties such as `Q`


Use this API to retrieve an existing user from a list.

*Success Response:*



Success. Returns first name, last name, emails, phone number and gender of the user. Refer `UserSearchResponseSchema` for more details.


Schema: `UserSearchResponseSchema`









---


#### createUser
Create user

```golang

data, err := User.CreateUser(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID | 


| ApplicationID | string | Application ID | 


| body |  CreateUserRequestSchema | "Request body" 

Create user

*Success Response:*



User create


Schema: `CreateUserResponseSchema`









---


#### updateUser
Update user

```golang

data, err := User.UpdateUser(CompanyID, ApplicationID, UserID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID | 


| ApplicationID | string | Application ID | 


| UserID | string | User ID | 


| body |  UpdateUserRequestSchema | "Request body" 

Update user

*Success Response:*



User update


Schema: `CreateUserResponseSchema`









---


#### createUserSession
Create user session

```golang

data, err := User.CreateUserSession(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company ID | 


| ApplicationID | string | Application ID | 


| body |  CreateUserSessionRequestSchema | "Request body" 

Create user session

*Success Response:*



Create user session


Schema: `CreateUserSessionResponseSchema`









---


#### getPlatformConfig
Get platform configurations

```golang

data, err := User.GetPlatformConfig(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 



Use this API to get all the platform configurations such as mobile image, desktop image, social logins, and all other text.

*Success Response:*



Success. Returns a JSON object containing the all the platform configurations. Refer `PlatformSchema` for more details.


Schema: `PlatformSchema`









---


#### updatePlatformConfig
Update platform configurations

```golang

data, err := User.UpdatePlatformConfig(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 


| body |  PlatformSchema | "Request body" 

Use this API to edit the existing platform configurations such as mobile image, desktop image, social logins, and all other text.

*Success Response:*



Success. Returns a JSON object with the updated platform configurations. Refer `PlatformSchema` for more details.


Schema: `PlatformSchema`









---



---


## Content


#### getAnnouncementsList
Get a list of announcements

```golang

data, err := Content.GetAnnouncementsList(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`


Announcements are useful to highlight a message or information on top of a webpage. Use this API to retrieve a list of announcements.	

*Success Response:*



Success. Refer `GetAnnouncementListSchema` for more details.


Schema: `GetAnnouncementListSchema`


*Examples:*


success
```json
{
  "$ref": "#/components/examples/GetAnnouncementList"
}
```









---


#### createAnnouncement
Create an announcement

```golang

data, err := Content.CreateAnnouncement(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| body |  AdminAnnouncementSchema | "Request body" 

Announcements are useful to highlight a message or information on top of a webpage. Use this API to create an announcement.

*Success Response:*



Success. Refer `CreateAnnouncementSchema` for more details.


Schema: `CreateAnnouncementSchema`


*Examples:*


success
```json
{
  "$ref": "#/components/examples/CreateAnnouncement"
}
```









---


#### getAnnouncementById
Get announcement by ID

```golang

data, err := Content.GetAnnouncementById(CompanyID, ApplicationID, AnnouncementID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| AnnouncementID | string | ID allotted to the announcement. | 



Use this API to retrieve an announcement and its details such as the target platform and pages on which it's applicable

*Success Response:*



Success. Refer `AdminAnnouncementSchema` for more details.


Schema: `AdminAnnouncementSchema`


*Examples:*


success
```json
{
  "$ref": "#/components/examples/Announcement"
}
```









---


#### updateAnnouncement
Update an announcement

```golang

data, err := Content.UpdateAnnouncement(CompanyID, ApplicationID, AnnouncementID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| AnnouncementID | string | ID allotted to the announcement. | 


| body |  AdminAnnouncementSchema | "Request body" 

Use this API to edit an existing announcement and its details such as the target platform and pages on which it's applicable

*Success Response:*



Success. Refer `CreateAnnouncementSchema` for more details.


Schema: `CreateAnnouncementSchema`


*Examples:*


success
```json
{
  "$ref": "#/components/examples/UpdateAnnouncement"
}
```









---


#### updateAnnouncementSchedule
Update the schedule and the publish status of an announcement

```golang

data, err := Content.UpdateAnnouncementSchedule(CompanyID, ApplicationID, AnnouncementID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| AnnouncementID | string | ID allotted to the announcement. | 


| body |  ScheduleSchema | "Request body" 

Use this API to edit the duration, i.e. start date-time and end date-time of an announcement. Moreover, you can enable/disable an announcement using this API.

*Success Response:*



Success. Refer `CreateAnnouncementSchema` for more details.


Schema: `CreateAnnouncementSchema`


*Examples:*


success
```json
{
  "$ref": "#/components/examples/PatchAnnouncement"
}
```









---


#### deleteAnnouncement
Delete announcement by id

```golang

data, err := Content.DeleteAnnouncement(CompanyID, ApplicationID, AnnouncementID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| AnnouncementID | string | ID allotted to the announcement. | 



Use this API to delete an existing announcement.

*Success Response:*



Success.


Schema: `CreateAnnouncementSchema`


*Examples:*


success
```json
{
  "$ref": "#/components/examples/DeleteAnnouncement"
}
```









---


#### createBlog
Create a blog

```golang

data, err := Content.CreateBlog(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| body |  BlogRequest | "Request body" 

Use this API to create a blog.

*Success Response:*



Success. Refer `BlogSchema` for more details.


Schema: `BlogSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/BlogResponse"
}
```









---


#### getBlogs
Get blogs

```golang

data, err := Content.GetBlogs(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`


Use this API to get a list of blogs along with their details, such as the title, reading time, publish status, feature image, tags, author, etc.

*Success Response:*



Success. Refer `BlogGetResponse` for more details.


Schema: `BlogGetResponse`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/BlogGetResponse"
}
```









---


#### updateBlog
Update a blog

```golang

data, err := Content.UpdateBlog(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| ID | string | ID allotted to the blog. | 


| body |  BlogRequest | "Request body" 

Use this API to update the details of an existing blog which includes title, feature image, content, SEO details, expiry, etc.

*Success Response:*



Success.


Schema: `BlogSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/BlogResponse"
}
```









---


#### deleteBlog
Delete blogs

```golang

data, err := Content.DeleteBlog(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| ID | string | ID allotted to the blog. | 



Use this API to delete a blog.

*Success Response:*



Success.


Schema: `BlogSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/BlogResponse"
}
```









---


#### getComponentById
Get components of a blog

```golang

data, err := Content.GetComponentById(CompanyID, ApplicationID, Slug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| Slug | string | A short, human-readable, URL-friendly identifier of a blog page. You can get slug value of a blog from `getBlogs` API. | 



Use this API to retrieve the components of a blog, such as title, slug, feature image, content, schedule, publish status, author, etc.

*Success Response:*



Success. Returns a a JSON object with components. Refer `BlogSchema` for more details.


Schema: `BlogSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/BlogResponse"
}
```









---


#### getFaqCategories
Get a list of FAQ categories

```golang

data, err := Content.GetFaqCategories(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 



FAQs can be divided into categories. Use this API to get a list of FAQ categories.

*Success Response:*



Success. Refer `GetFaqCategoriesSchema` for more details.


Schema: `GetFaqCategoriesSchema`









---


#### getFaqCategoryBySlugOrId
Get an FAQ category by slug or id

```golang

data, err := Content.GetFaqCategoryBySlugOrId(CompanyID, ApplicationID, IDOrSlug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| IDOrSlug | string | ID or the slug allotted to an FAQ category. Slug is a short, human-readable, URL-friendly identifier of an object. You can get slug value of an FAQ category from `getFaqCategories` API. | 



FAQs can be divided into categories. Use this API to get an FAQ categories using its slug or ID.

*Success Response:*



Success. Refer `GetFaqCategoryBySlugSchema` for more details.


Schema: `GetFaqCategoryBySlugSchema`









---


#### createFaqCategory
Create an FAQ category

```golang

data, err := Content.CreateFaqCategory(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| body |  CreateFaqCategoryRequestSchema | "Request body" 

FAQs help users to solve an issue or know more about a process. FAQs can be categorized separately, for e.g. some questions can be related to payment, some could be related to purchase, shipping, navigating, etc. Use this API to create an FAQ category.

*Success Response:*



Success.


Schema: `CreateFaqCategorySchema`









---


#### updateFaqCategory
Update an FAQ category

```golang

data, err := Content.UpdateFaqCategory(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| ID | string | ID allotted to an FAQ category. | 


| body |  UpdateFaqCategoryRequestSchema | "Request body" 

Use this API to edit an existing FAQ category.

*Success Response:*



Success.


Schema: `CreateFaqCategorySchema`









---


#### deleteFaqCategory
Delete an FAQ category

```golang

data, err := Content.DeleteFaqCategory(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| ID | string | ID allotted to an FAQ category. | 



Use this API to delete an FAQ category.

*Success Response:*



Success.


Schema: `FaqSchema`









---


#### getFaqsByCategoryIdOrSlug
Get question and answers within an FAQ category

```golang

data, err := Content.GetFaqsByCategoryIdOrSlug(CompanyID, ApplicationID, IDOrSlug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| IDOrSlug | string | ID or the slug allotted to an FAQ category. Slug is a short, human-readable, URL-friendly identifier of an object. You can get slug value of an FAQ category from `getFaqCategories` API. | 



Use this API to retrieve all the commonly asked question and answers belonging to an FAQ category.

*Success Response:*



Success. Refer `GetFaqSchema` for more details.


Schema: `GetFaqSchema`









---


#### addFaq
Create an FAQ

```golang

data, err := Content.AddFaq(CompanyID, ApplicationID, CategoryID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| CategoryID | string | ID allotted to an FAQ category. | 


| body |  CreateFaqSchema | "Request body" 

FAQs help users to solve an issue or know more about a process. Use this API to create an FAQ for a given FAQ category.

*Success Response:*



Success.


Schema: `CreateFaqResponseSchema`









---


#### updateFaq
Update an FAQ

```golang

data, err := Content.UpdateFaq(CompanyID, ApplicationID, CategoryID, FaqID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| CategoryID | string | ID allotted to an FAQ category. | 


| FaqID | string | ID allotted to an FAQ. | 


| body |  CreateFaqSchema | "Request body" 

Use this API to edit an existing FAQ.

*Success Response:*



Success.


Schema: `CreateFaqResponseSchema`









---


#### deleteFaq
Delete an FAQ

```golang

data, err := Content.DeleteFaq(CompanyID, ApplicationID, CategoryID, FaqID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| CategoryID | string | ID allotted to an FAQ category. | 


| FaqID | string | ID allotted to an FAQ. | 



Use this API to delete an existing FAQ.

*Success Response:*



Success.


Schema: `CreateFaqResponseSchema`









---


#### getFaqByIdOrSlug
Get an FAQ

```golang

data, err := Content.GetFaqByIdOrSlug(CompanyID, ApplicationID, IDOrSlug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| IDOrSlug | string | ID or the slug allotted to an FAQ category. Slug is a short, human-readable, URL-friendly identifier of an object. You can get slug value of an FAQ category from `getFaqCategories` API. | 



Use this API to retrieve a specific FAQ. You will get the question and answer of that FAQ.

*Success Response:*



Success. Refer `CreateFaqResponseSchema` for more details.


Schema: `CreateFaqResponseSchema`









---


#### getLandingPages
Get landing pages

```golang

data, err := Content.GetLandingPages(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`


Landing page is the first page that a prospect lands upon while visiting a website. Use this API to fetch a list of landing pages.

*Success Response:*



Success. Refer `LandingPageGetResponse` for more details.


Schema: `LandingPageGetResponse`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/LandingPageGetResponse"
}
```









---


#### createLandingPage
Create a landing page

```golang

data, err := Content.CreateLandingPage(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| body |  LandingPageSchema | "Request body" 

Landing page is the first page that a prospect lands upon while visiting a website. Use this API to create a landing page.

*Success Response:*



Success.


Schema: `LandingPageSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/LandingPageResponse"
}
```









---


#### updateLandingPage
Update a landing page

```golang

data, err := Content.UpdateLandingPage(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| ID | string | ID allotted to a landing page. | 


| body |  LandingPageSchema | "Request body" 

Use this API to edit the details of an existing landing page.

*Success Response:*



Success.


Schema: `LandingPageSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/LandingPageResponse"
}
```









---


#### deleteLandingPage
Delete a landing page

```golang

data, err := Content.DeleteLandingPage(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| ID | string | ID allotted to a landing page. | 



Use this API to delete an existing landing page.

*Success Response:*



Success.


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
    "archived": true
  }
}
```









---


#### getLegalInformation
Get legal information

```golang

data, err := Content.GetLegalInformation(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 



Use this API to get the legal information of an application, which includes Policy, Terms and Conditions, Shipping Policy and FAQ regarding the application.

*Success Response:*



Success. Refer `ApplicationLegal` for more details.


Schema: `ApplicationLegal`


*Examples:*


Success
```json
{
  "$ref": "#/components/examples/Legal"
}
```









---


#### updateLegalInformation
Save legal information

```golang

data, err := Content.UpdateLegalInformation(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| body |  ApplicationLegal | "Request body" 

Use this API to edit, update and save the legal information of an application, which includes Policy, Terms and Conditions, Shipping Policy and FAQ regarding the application.

*Success Response:*



Success. Refer `ApplicationLegal` for more details.


Schema: `ApplicationLegal`









---


#### getNavigations
Get navigations

```golang

data, err := Content.GetNavigations(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 







| xQuery | struct | Includes properties such as `DevicePlatform`, `PageNo`, `PageSize`


Use this API to fetch the navigations details which includes the items of the navigation pane. It also shows the orientation, links, sub-navigations, etc.

*Success Response:*



Success. Refer `NavigationGetResponse` for more details.


Schema: `NavigationGetResponse`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/NavigationGetResponse"
}
```









---


#### createNavigation
Create a navigation

```golang

data, err := Content.CreateNavigation(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| body |  NavigationRequest | "Request body" 

Navigation is the arrangement of navigational items to ease the accessibility of resources for users on a website. Use this API to create a navigation.

*Success Response:*



Success.


Schema: `NavigationSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/NavigationResponse"
}
```









---


#### getDefaultNavigations
Get default navigations

```golang

data, err := Content.GetDefaultNavigations(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 



On any website (application), there are navigations that are present by default. Use this API to retrieve those default navigations.

*Success Response:*



Success. Refer `DefaultNavigationResponse` for more details.


Schema: `DefaultNavigationResponse`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/DefaultNavigationResponse"
}
```









---


#### getNavigationBySlug
Get a navigation by slug

```golang

data, err := Content.GetNavigationBySlug(CompanyID, ApplicationID, Slug, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| Slug | string | A short, human-readable, URL-friendly identifier of a navigation. You can get slug value of a navigation from `getNavigations` API. | 



| xQuery | struct | Includes properties such as `DevicePlatform`


Use this API to retrieve a navigation by its slug.

*Success Response:*



Success. Refer `NavigationSchema` for more details.


Schema: `NavigationSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/NavigationResponse"
}
```









---


#### updateNavigation
Update a navigation

```golang

data, err := Content.UpdateNavigation(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| ID | string | ID allotted to the navigation. | 


| body |  NavigationRequest | "Request body" 

Use this API to edit the details of an existing navigation.

*Success Response:*



Success.


Schema: `NavigationSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/NavigationResponse"
}
```









---


#### deleteNavigation
Delete a navigation

```golang

data, err := Content.DeleteNavigation(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| ID | string | ID allotted to the navigation. | 



Use this API to delete an existing navigation.

*Success Response:*



Success.


Schema: `NavigationSchema`


*Examples:*


default
```json
{
  "value": {
    "_id": "5ffbd9b90ac98678ae0458d7",
    "application": "000000000000000000000001",
    "_custom_json": null,
    "name": "temp",
    "slug": "temp",
    "platform": "[web]",
    "orientation": {
      "portrait": [
        "left"
      ]
    },
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
    },
    "archived": true
  }
}
```









---


#### getPageMeta
Get page meta

```golang

data, err := Content.GetPageMeta(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 



Use this API to get the meta of custom pages (blog, page) and default system pages (e.g. home/brand/category/collection).

*Success Response:*



Success. Refer `PageMetaSchema` for more details.


Schema: `PageMetaSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/PageMeta"
}
```









---


#### getPageSpec
Get page spec

```golang

data, err := Content.GetPageSpec(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 



Use this API to get the specifications of a page, such as page type, display name, params and query.

*Success Response:*



Success. Refer `PageSpec` for more details.


Schema: `PageSpec`


*Examples:*


default
```json
{
  "value": {
    "specifications": [
      {
        "page_type": "home",
        "display_name": "Home",
        "params": [],
        "query": []
      },
      {
        "page_type": "collections",
        "display_name": "Collections",
        "params": [],
        "query": []
      },
      {
        "page_type": "collection",
        "display_name": "Collection",
        "params": [
          {
            "key": "slug",
            "required": true
          }
        ],
        "query": []
      }
    ]
  }
}
```









---


#### createPage
Create a page

```golang

data, err := Content.CreatePage(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| body |  PageRequest | "Request body" 

Use this API to create a custom page using a title, seo, publish status, feature image, tags, meta, etc.

*Success Response:*



Success. Refer `PageSchema` for more details.


Schema: `PageSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/PageResponse"
}
```









---


#### getPages
Get a list of pages

```golang

data, err := Content.GetPages(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`


Use this API to retrieve a list of pages.

*Success Response:*



Success. Refer `PageGetResponse` for more details.


Schema: `PageGetResponse`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/PageGetResponse"
}
```









---


#### createPagePreview
Create a page preview

```golang

data, err := Content.CreatePagePreview(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| body |  PageRequest | "Request body" 

Use this API to create a page preview to check the appearance of a custom page.

*Success Response:*



Success. Refer `PageSchema` for more details.


Schema: `PageSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/PageResponse"
}
```









---


#### updatePagePreview
Change the publish status of a page

```golang

data, err := Content.UpdatePagePreview(CompanyID, ApplicationID, Slug, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| Slug | string | A short, human-readable, URL-friendly identifier of a page. You can get slug value of a page from `getPages` API. | 


| body |  PagePublishRequest | "Request body" 

Use this API to change the publish status of an existing page. Allows you to publish and unpublish the page.

*Success Response:*



Success.


Schema: `PageSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/PageResponse"
}
```









---


#### updatePage
Update a page

```golang

data, err := Content.UpdatePage(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| ID | string | ID allotted to the page. | 


| body |  PageSchema | "Request body" 

Use this API to edit the details of an existing page, such as its title, seo, publish status, feature image, tags, schedule, etc.

*Success Response:*



Success. Refer `PageSchema` for more details.


Schema: `PageSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/PageResponse"
}
```









---


#### deletePage
Delete a page

```golang

data, err := Content.DeletePage(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| ID | string | ID allotted to the page. | 



Use this API to delete an existing page.

*Success Response:*



Success.


Schema: `PageSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/PageResponse"
}
```









---


#### getPageBySlug
Get pages by component Id

```golang

data, err := Content.GetPageBySlug(CompanyID, ApplicationID, Slug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| Slug | string | A short, human-readable, URL-friendly identifier of a page. You can get slug value of a page from `getPages` API. | 



Use this API to retrieve the components of a page, such as its title, seo, publish status, feature image, tags, schedule, etc.

*Success Response:*



Success. Returns a JSON object of components. Refer `PageSchema` for more details.


Schema: `PageSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/PageResponse"
}
```









---


#### getSEOConfiguration
Get SEO configuration of an application

```golang

data, err := Content.GetSEOConfiguration(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 



Use this API to know how the SEO is configured in the application. This includes the sitemap, robot.txt, custom meta tags, etc.

*Success Response:*



Success. Refer `SeoComponent` for more details.


Schema: `SeoComponent`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/Seo"
}
```









---


#### updateSEOConfiguration
Update SEO of application

```golang

data, err := Content.UpdateSEOConfiguration(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| body |  SeoComponent | "Request body" 

Use this API to edit the SEO details of an application. This includes the sitemap, robot.txt, custom meta tags, etc.

*Success Response:*



Success. Refer `SeoSchema` for more details.


Schema: `SeoSchema`


*Examples:*


default
```json
{
  "value": {
    "details": {
      "title": "Zyosa Zyosa"
    },
    "robots_txt": "User-agent: * \nAllow: / \nsancisciasn xwsaixjowqnxwsiwjs",
    "sitemap_enabled": false,
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
```









---


#### getSlideshows
Get slideshows

```golang

data, err := Content.GetSlideshows(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 







| xQuery | struct | Includes properties such as `DevicePlatform`, `PageNo`, `PageSize`


A slideshow is a group of images, videos or a combination of both that are shown on the website in the form of slides. Use this API to fetch a list of slideshows.

*Success Response:*



Success. Refer `SlideshowGetResponse` for more details.


Schema: `SlideshowGetResponse`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/SlideshowGetResponse"
}
```









---


#### createSlideshow
Create a slideshow

```golang

data, err := Content.CreateSlideshow(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| body |  SlideshowRequest | "Request body" 

A slideshow is a group of images, videos or a combination of both that are shown on the website in the form of slides. Use this API to create a slideshow.

*Success Response:*



Success. Refer `SlideshowSchema` for more details.


Schema: `SlideshowSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/SlideshowResponse"
}
```









---


#### getSlideshowBySlug
Get slideshow by slug

```golang

data, err := Content.GetSlideshowBySlug(CompanyID, ApplicationID, Slug, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| Slug | string | A short, human-readable, URL-friendly identifier of a slideshow. You can get slug value of a page from `getSlideshows` API. | 



| xQuery | struct | Includes properties such as `DevicePlatform`


Use this API to retrieve the details of a slideshow by its slug.

*Success Response:*



Success. Refer `SlideshowSchema` for more details.


Schema: `SlideshowSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/SlideshowResponse"
}
```









---


#### updateSlideshow
Update a slideshow

```golang

data, err := Content.UpdateSlideshow(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| ID | string | ID allotted to the slideshow. | 


| body |  SlideshowRequest | "Request body" 

Use this API to edit the details of an existing slideshow.

*Success Response:*



Success. Refer `SlideshowSchema` for more details.


Schema: `SlideshowSchema`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/SlideshowResponse"
}
```









---


#### deleteSlideshow
Delete a slideshow

```golang

data, err := Content.DeleteSlideshow(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform | 


| ApplicationID | string | Numeric ID allotted to an application created within a business account. | 


| ID | string | ID allotted to the slideshow. | 



Use this API to delete an existing slideshow.

*Success Response:*



Success.


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
    "archived": true,
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
Get support information

```golang

data, err := Content.GetSupportInformation(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 



Use this API to get the contact details for customer support, including emails and phone numbers.

*Success Response:*



Success. Refer `Support` for more details.


Schema: `Support`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/Support"
}
```









---


#### updateSupportInformation
Update the support data of an application

```golang

data, err := Content.UpdateSupportInformation(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 


| body |  Support | "Request body" 

Use this API to edit the existing contact details for customer support, including emails and phone numbers.

*Success Response:*



Success. Refer `Support` for more details.


Schema: `Support`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/Support"
}
```









---


#### updateInjectableTag
Update a tag

```golang

data, err := Content.UpdateInjectableTag(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 


| body |  CreateTagRequestSchema | "Request body" 

Use this API to edit the details of an existing tag. This includes the tag name, tag type (css/js), url and position of the tag.

*Success Response:*



Success.


Schema: `TagsSchema`









---


#### deleteAllInjectableTags
Delete tags in application

```golang

data, err := Content.DeleteAllInjectableTags(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 



Use this API to delete all the existing tags at once.

*Success Response:*



Success.


Schema: `TagsSchema`









---


#### getInjectableTags
Get all the tags in an application

```golang

data, err := Content.GetInjectableTags(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 



Use this API to get all the CSS and JS injected in the application in the form of tags.

*Success Response:*



Success. Refer `TagsSchema` for more details.


Schema: `TagsSchema`









---


#### addInjectableTag
Add a tag

```golang

data, err := Content.AddInjectableTag(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 


| body |  CreateTagRequestSchema | "Request body" 

CSS and JS can be injected in the application (website) with the help of tags. Use this API to create such tags by entering the tag name, tag type (css/js), url and position of the tag.

*Success Response:*



Success.


Schema: `TagsSchema`









---


#### removeInjectableTag
Remove a tag

```golang

data, err := Content.RemoveInjectableTag(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 


| body |  RemoveHandpickedSchema | "Request body" 

Use this API to delete an existing tag.

*Success Response:*



Success.


Schema: `TagDeleteSuccessResponse`









---


#### editInjectableTag
Edit a tag by id

```golang

data, err := Content.EditInjectableTag(CompanyID, ApplicationID, TagID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Numeric ID allotted to a business account on Fynd Platform. | 


| ApplicationID | string | Alphanumeric ID allotted to an application created within a business account. | 


| TagID | string | ID allotted to the tag. | 


| body |  UpdateHandpickedSchema | "Request body" 

Use this API to edit the details of an existing tag by its ID.

*Success Response:*



Success.


Schema: `TagsSchema`









---



---


## Billing


#### createSubscriptionCharge
Create subscription charge

```golang

data, err := Billing.CreateSubscriptionCharge(CompanyID, ExtensionID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Customer unique id. In case of company it will be company id. | 


| ExtensionID | string | Extension _id | 


| body |  CreateSubscriptionCharge | "Request body" 

Register subscription charge for a seller of your extension.

*Success Response:*



Success


Schema: `CreateSubscriptionResponse`









---


#### getSubscriptionCharge
Get subscription charge details

```golang

data, err := Billing.GetSubscriptionCharge(CompanyID, ExtensionID, SubscriptionID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Customer unique id. In case of company it will be company id. | 


| ExtensionID | string | Extension _id | 


| SubscriptionID | string | Subscription charge _id | 



Get created subscription charge details

*Success Response:*



Success


Schema: `EntitySubscription`









---


#### cancelSubscriptionCharge
Cancel subscription charge

```golang

data, err := Billing.CancelSubscriptionCharge(CompanyID, ExtensionID, SubscriptionID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Customer unique id. In case of company it will be company id. | 


| ExtensionID | string | Extension _id | 


| SubscriptionID | string | Subscription charge _id | 



Cancel subscription and attached charges.

*Success Response:*



Success


Schema: `EntitySubscription`









---


#### getInvoices
Get invoices

```golang

data, err := Billing.GetInvoices(CompanyID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Customer unique id. In case of company it will be company id. | 



Get invoices.

*Success Response:*



Success


Schema: `Invoices`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/Invoices"
}
```









---


#### getInvoiceById
Get invoice by id

```golang

data, err := Billing.GetInvoiceById(CompanyID, InvoiceID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Customer unique id. In case of company it will be company id. | 


| InvoiceID | string | Invoice id | 



Get invoice by id.

*Success Response:*



Success


Schema: `Invoice`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/Invoice"
}
```









---


#### getCustomerDetail
Get subscription customer detail

```golang

data, err := Billing.GetCustomerDetail(CompanyID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Customer unique id. In case of company it will be company id. | 



Get subscription customer detail.

*Success Response:*



Success


Schema: `SubscriptionCustomer`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/SubscriptionCustomerRes"
}
```









---


#### upsertCustomerDetail
Upsert subscription customer detail

```golang

data, err := Billing.UpsertCustomerDetail(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Customer unique id. In case of company it will be company id. | 


| body |  SubscriptionCustomerCreate | "Request body" 

Upsert subscription customer detail.

*Success Response:*



Success


Schema: `SubscriptionCustomer`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/SubscriptionCustomerRes"
}
```









---


#### getSubscription
Get current subscription detail

```golang

data, err := Billing.GetSubscription(CompanyID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Customer unique id. In case of company it will be company id. | 



If subscription is active then it will return is_enabled true and return subscription object. If subscription is not active then is_enabled false and message.


*Success Response:*



Success


Schema: `SubscriptionStatus`


*Examples:*


Active subscription
```json
{
  "$ref": "#/components/examples/SubscriptionActiveRes"
}
```

Inactive subscription
```json
{
  "$ref": "#/components/examples/SubscriptionInavtiveRes"
}
```









---


#### getFeatureLimitConfig
Get subscription subscription limits

```golang

data, err := Billing.GetFeatureLimitConfig(CompanyID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Customer unique id. In case of company it will be company id. | 



Get subscription subscription limits.

*Success Response:*



Success


Schema: `SubscriptionLimit`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/CurrentLimitRes"
}
```









---


#### activateSubscriptionPlan
Activate subscription

```golang

data, err := Billing.ActivateSubscriptionPlan(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Customer unique id. In case of company it will be company id. | 


| body |  SubscriptionActivateReq | "Request body" 

It will activate subscription plan for customer

*Success Response:*



Success


Schema: `SubscriptionActivateRes`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/SubscriptionActivateRes"
}
```









---


#### cancelSubscriptionPlan
Cancel subscription

```golang

data, err := Billing.CancelSubscriptionPlan(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Customer unique id. In case of company it will be company id. | 


| body |  CancelSubscriptionReq | "Request body" 

It will cancel current active subscription.

*Success Response:*



Success


Schema: `CancelSubscriptionRes`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/CancelSubscriptionRes"
}
```









---



---


## Communication


#### getCampaigns
Get campaigns

```golang

data, err := Communication.GetCampaigns(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 







| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Sort`


Get campaigns

*Success Response:*



Success


Schema: `Campaigns`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/Campaigns"
}
```









---


#### createCampaign
Create campaign

```golang

data, err := Communication.CreateCampaign(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| body |  CampaignReq | "Request body" 

Create campaign

*Success Response:*



Success


Schema: `Campaign`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/Campaign"
}
```









---


#### getCampaignById
Get campaign by id

```golang

data, err := Communication.GetCampaignById(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| ID | string | Campaign id | 



Get campaign by id

*Success Response:*



Success


Schema: `Campaign`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/Campaign"
}
```









---


#### updateCampaignById
Update campaign by id

```golang

data, err := Communication.UpdateCampaignById(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| ID | string | Campaign id | 


| body |  CampaignReq | "Request body" 

Update campaign by id

*Success Response:*



Success


Schema: `Campaign`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/Campaign"
}
```









---


#### getStatsOfCampaignById
Get stats of campaign by id

```golang

data, err := Communication.GetStatsOfCampaignById(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| ID | string | Campaign id | 



Get stats of campaign by id

*Success Response:*



Success


Schema: `GetStats`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/GetStats"
}
```









---


#### getAudiences
Get audiences

```golang

data, err := Communication.GetAudiences(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 







| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Sort`


Get audiences

*Success Response:*



Success


Schema: `Audiences`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/Audiences"
}
```









---


#### createAudience
Create audience

```golang

data, err := Communication.CreateAudience(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| body |  AudienceReq | "Request body" 

Create audience

*Success Response:*



Success


Schema: `Audience`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/Audience"
}
```









---


#### getBigqueryHeaders
Get bigquery headers

```golang

data, err := Communication.GetBigqueryHeaders(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| body |  BigqueryHeadersReq | "Request body" 

Get bigquery headers

*Success Response:*



Success


Schema: `BigqueryHeadersRes`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/BigqueryHeadersRes"
}
```









---


#### getAudienceById
Get audience by id

```golang

data, err := Communication.GetAudienceById(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| ID | string | Audience id | 



Get audience by id

*Success Response:*



Success


Schema: `Audience`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/Audience"
}
```









---


#### updateAudienceById
Update audience by id

```golang

data, err := Communication.UpdateAudienceById(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| ID | string | Audience id | 


| body |  AudienceReq | "Request body" 

Update audience by id

*Success Response:*



Success


Schema: `Audience`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/Audience"
}
```









---


#### getNSampleRecordsFromCsv
Get n sample records from csv

```golang

data, err := Communication.GetNSampleRecordsFromCsv(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| body |  GetNRecordsCsvReq | "Request body" 

Get n sample records from csv

*Success Response:*



Success


Schema: `GetNRecordsCsvRes`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/GetNRecordsCsvRes"
}
```









---


#### getEmailProviders
Get email providers

```golang

data, err := Communication.GetEmailProviders(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 







| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Sort`


Get email providers

*Success Response:*



Success


Schema: `EmailProviders`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/EmailProviders"
}
```









---


#### createEmailProvider
Create email provider

```golang

data, err := Communication.CreateEmailProvider(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| body |  EmailProviderReq | "Request body" 

Create email provider

*Success Response:*



Success


Schema: `EmailProvider`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/EmailProvider"
}
```









---


#### getEmailProviderById
Get email provider by id

```golang

data, err := Communication.GetEmailProviderById(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| ID | string | Email provider id | 



Get email provider by id

*Success Response:*



Success


Schema: `EmailProvider`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/EmailProvider"
}
```









---


#### updateEmailProviderById
Update email provider by id

```golang

data, err := Communication.UpdateEmailProviderById(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| ID | string | Email provider id | 


| body |  EmailProviderReq | "Request body" 

Update email provider by id

*Success Response:*



Success


Schema: `EmailProvider`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/EmailProvider"
}
```









---


#### getEmailTemplates
Get email templates

```golang

data, err := Communication.GetEmailTemplates(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 







| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Sort`


Get email templates

*Success Response:*



Success


Schema: `EmailTemplates`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/EmailTemplates"
}
```









---


#### createEmailTemplate
Create email template

```golang

data, err := Communication.CreateEmailTemplate(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| body |  EmailTemplateReq | "Request body" 

Create email template

*Success Response:*



Success


Schema: `EmailTemplateRes`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/EmailTemplateRes"
}
```









---


#### getSystemEmailTemplates
Get system email templates

```golang

data, err := Communication.GetSystemEmailTemplates(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 







| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Sort`


Get system email templates

*Success Response:*



Success


Schema: `SystemEmailTemplates`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/SystemEmailTemplates"
}
```









---


#### getEmailTemplateById
Get email template by id

```golang

data, err := Communication.GetEmailTemplateById(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| ID | string | Email template id | 



Get email template by id

*Success Response:*



Success


Schema: `EmailTemplate`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/EmailTemplate"
}
```









---


#### updateEmailTemplateById
Update email template by id

```golang

data, err := Communication.UpdateEmailTemplateById(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| ID | string | Email template id | 


| body |  EmailTemplateReq | "Request body" 

Update email template by id

*Success Response:*



Success


Schema: `EmailTemplateRes`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/EmailTemplateRes"
}
```









---


#### deleteEmailTemplateById
Delete email template by id

```golang

data, err := Communication.DeleteEmailTemplateById(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| ID | string | Email template id | 



Delete email template by id

*Success Response:*



Success


Schema: `EmailTemplateDeleteSuccessRes`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/EmailTemplateDeleteSuccessRes"
}
```









---


#### getEventSubscriptions
Get event subscriptions

```golang

data, err := Communication.GetEventSubscriptions(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 







| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Populate`


Get event subscriptions

*Success Response:*



Success


Schema: `EventSubscriptions`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/EventSubscriptions"
}
```









---


#### getJobs
Get jobs

```golang

data, err := Communication.GetJobs(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 







| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Sort`


Get jobs

*Success Response:*



Success


Schema: `Jobs`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/Jobs"
}
```









---


#### triggerCampaignJob
Trigger campaign job

```golang

data, err := Communication.TriggerCampaignJob(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| body |  TriggerJobRequest | "Request body" 

Trigger campaign job

*Success Response:*



Success


Schema: `TriggerJobResponse`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/TriggerJobResponse"
}
```









---


#### getJobLogs
Get job logs

```golang

data, err := Communication.GetJobLogs(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 







| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Sort`


Get job logs

*Success Response:*



Success


Schema: `JobLogs`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/JobLogs"
}
```









---


#### getCommunicationLogs
Get communication logs

```golang

data, err := Communication.GetCommunicationLogs(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 









| xQuery | struct | Includes properties such as `PageID`, `PageSize`, `Sort`, `Query`


Get communication logs

*Success Response:*



Success


Schema: `Logs`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/Logs"
}
```









---


#### getSystemNotifications
Get system notifications

```golang

data, err := Communication.GetSystemNotifications(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`


Get system notifications

*Success Response:*



Success


Schema: `SystemNotifications`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/SystemNotifications"
}
```









---


#### getSmsProviders
Get sms providers

```golang

data, err := Communication.GetSmsProviders(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 







| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Sort`


Get sms providers

*Success Response:*



Success


Schema: `SmsProviders`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/SmsProviders"
}
```









---


#### createSmsProvider
Create sms provider

```golang

data, err := Communication.CreateSmsProvider(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| body |  SmsProviderReq | "Request body" 

Create sms provider

*Success Response:*



Success


Schema: `SmsProvider`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/SmsProvider"
}
```









---


#### getSmsProviderById
Get sms provider by id

```golang

data, err := Communication.GetSmsProviderById(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| ID | string | Sms provider id | 



Get sms provider by id

*Success Response:*



Success


Schema: `SmsProvider`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/SmsProvider"
}
```









---


#### updateSmsProviderById
Update sms provider by id

```golang

data, err := Communication.UpdateSmsProviderById(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| ID | string | Sms provider id | 


| body |  SmsProviderReq | "Request body" 

Update sms provider by id

*Success Response:*



Success


Schema: `SmsProvider`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/SmsProvider"
}
```









---


#### getSmsTemplates
Get sms templates

```golang

data, err := Communication.GetSmsTemplates(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 







| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Sort`


Get sms templates

*Success Response:*



Success


Schema: `SmsTemplates`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/SmsTemplates"
}
```









---


#### createSmsTemplate
Create sms template

```golang

data, err := Communication.CreateSmsTemplate(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| body |  SmsTemplateReq | "Request body" 

Create sms template

*Success Response:*



Success


Schema: `SmsTemplateRes`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/SmsTemplateRes"
}
```









---


#### getSmsTemplateById
Get sms template by id

```golang

data, err := Communication.GetSmsTemplateById(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| ID | string | Sms template id | 



Get sms template by id

*Success Response:*



Success


Schema: `SmsTemplate`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/SmsTemplate"
}
```









---


#### updateSmsTemplateById
Update sms template by id

```golang

data, err := Communication.UpdateSmsTemplateById(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| ID | string | Sms template id | 


| body |  SmsTemplateReq | "Request body" 

Update sms template by id

*Success Response:*



Success


Schema: `SmsTemplateRes`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/SmsTemplateRes"
}
```









---


#### deleteSmsTemplateById
Delete sms template by id

```golang

data, err := Communication.DeleteSmsTemplateById(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 


| ID | string | Sms template id | 



Delete sms template by id

*Success Response:*



Success


Schema: `SmsTemplateDeleteSuccessRes`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/SmsTemplateDeleteSuccessRes"
}
```









---


#### getSystemSystemTemplates
Get system sms templates

```golang

data, err := Communication.GetSystemSystemTemplates(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company id | 


| ApplicationID | string | Application id | 







| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Sort`


Get system sms templates

*Success Response:*



Success


Schema: `SystemSmsTemplates`


*Examples:*


default
```json
{
  "$ref": "#/components/examples/SystemSmsTemplates"
}
```









---



---


## Payment


#### getBrandPaymentGatewayConfig
Get All Brand Payment Gateway Config Secret

```golang

data, err := Payment.GetBrandPaymentGatewayConfig(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 


| ApplicationID | string | Application id | 



Get All Brand Payment Gateway Config Secret

*Success Response:*



Refund Transfer Mode


Schema: `PaymentGatewayConfigResponse`









---


#### saveBrandPaymentGatewayConfig
Save Config Secret For Brand Payment Gateway

```golang

data, err := Payment.SaveBrandPaymentGatewayConfig(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 


| ApplicationID | string | Application id | 


| body |  PaymentGatewayConfigRequest | "Request body" 

Save Config Secret For Brand Payment Gateway

*Success Response:*



Save Config Secret For Brand Payment Gateway Success Response.


Schema: `PaymentGatewayToBeReviewed`









---


#### updateBrandPaymentGatewayConfig
Save Config Secret For Brand Payment Gateway

```golang

data, err := Payment.UpdateBrandPaymentGatewayConfig(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 


| ApplicationID | string | Application id | 


| body |  PaymentGatewayConfigRequest | "Request body" 

Save Config Secret For Brand Payment Gateway

*Success Response:*



Save Config Secret For Brand Payment Gateway Success Response.


Schema: `PaymentGatewayToBeReviewed`









---


#### getPaymentModeRoutes
Get All Valid Payment Options

```golang

data, err := Payment.GetPaymentModeRoutes(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 


| ApplicationID | string | Application id | 





| xQuery | struct | Includes properties such as `Refresh`, `RequestType`


Use this API to get Get All Valid Payment Options for making payment

*Success Response:*



Success


Schema: `PaymentOptionsResponse`









---


#### getAllPayouts
Get All Payouts

```golang

data, err := Payment.GetAllPayouts(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 



| xQuery | struct | Includes properties such as `UniqueExternalID`


Get All Payouts

*Success Response:*



payouts response object


Schema: `PayoutsResponse`









---


#### savePayout
Save Payout

```golang

data, err := Payment.SavePayout(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 


| body |  PayoutRequest | "Request body" 

Save Payout

*Success Response:*



save payout response object


Schema: `PayoutResponse`









---


#### updatePayout
Update Payout

```golang

data, err := Payment.UpdatePayout(CompanyID, UniqueTransferNo, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 


| UniqueTransferNo | string | Unique transfer id | 


| body |  PayoutRequest | "Request body" 

Update Payout

*Success Response:*



save payout response object


Schema: `UpdatePayoutResponse`









---


#### activateAndDectivatePayout
Partial Update Payout

```golang

data, err := Payment.ActivateAndDectivatePayout(CompanyID, UniqueTransferNo, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 


| UniqueTransferNo | string | Unique transfer id | 


| body |  UpdatePayoutRequest | "Request body" 

Partial Update Payout

*Success Response:*



save payout response object


Schema: `UpdatePayoutResponse`









---


#### deletePayout
Delete Payout

```golang

data, err := Payment.DeletePayout(CompanyID, UniqueTransferNo);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 


| UniqueTransferNo | string | Unique transfer id | 



Delete Payout

*Success Response:*



delete payout response object


Schema: `DeletePayoutResponse`









---


#### getSubscriptionPaymentMethod
List Subscription Payment Method

```golang

data, err := Payment.GetSubscriptionPaymentMethod(CompanyID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 



Get all  Subscription  Payment Method

*Success Response:*



List Subscription Payment Method Response


Schema: `SubscriptionPaymentMethodResponse`









---


#### deleteSubscriptionPaymentMethod
Delete Subscription Payment Method

```golang

data, err := Payment.DeleteSubscriptionPaymentMethod(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 





| xQuery | struct | Includes properties such as `UniqueExternalID`, `PaymentMethodID`


Uses this api to Delete Subscription Payment Method

*Success Response:*



Delete Subscription Payment Method Response.


Schema: `DeleteSubscriptionPaymentMethodResponse`









---


#### getSubscriptionConfig
List Subscription Config

```golang

data, err := Payment.GetSubscriptionConfig(CompanyID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 



Get all  Subscription Config details

*Success Response:*



List Subscription Config Response


Schema: `SubscriptionConfigResponse`









---


#### saveSubscriptionSetupIntent
Save Subscription Setup Intent

```golang

data, err := Payment.SaveSubscriptionSetupIntent(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 


| body |  SaveSubscriptionSetupIntentRequest | "Request body" 

Uses this api to Save Subscription Setup Intent

*Success Response:*



Save Subscription Setup Intent Response.


Schema: `SaveSubscriptionSetupIntentResponse`









---


#### addBeneficiaryDetails
Save bank details for cancelled/returned order

```golang

data, err := Payment.AddBeneficiaryDetails(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 


| ApplicationID | string | Application id | 


| body |  AddBeneficiaryDetailsRequest | "Request body" 

Use this API to save bank details for returned/cancelled order to refund amount in his account.

*Success Response:*



Success


Schema: `RefundAccountResponse`









---


#### verifyIfscCode
Ifsc Code Verification

```golang

data, err := Payment.VerifyIfscCode(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 



| xQuery | struct | Includes properties such as `IfscCode`


Get True/False for correct IFSC Code for adding bank details for refund

*Success Response:*



Bank details on correct Ifsc Code


Schema: `IfscCodeResponse`









---


#### getUserOrderBeneficiaries
List Order Beneficiary

```golang

data, err := Payment.GetUserOrderBeneficiaries(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



| CompanyID | float64 | Company Id | 


| ApplicationID | string | Application id | 

| xQuery | struct | Includes properties such as `OrderID`


Get all active  beneficiary details added by the user for refund

*Success Response:*



List Order Beneficiary


Schema: `OrderBeneficiaryResponse`









---


#### getUserBeneficiaries
List User Beneficiary

```golang

data, err := Payment.GetUserBeneficiaries(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



| CompanyID | float64 | Company Id | 


| ApplicationID | string | Application id | 

| xQuery | struct | Includes properties such as `OrderID`


Get all active  beneficiary details added by the user for refund

*Success Response:*



List User Beneficiary


Schema: `OrderBeneficiaryResponse`









---


#### confirmPayment
Confirm payment after successful payment from payment gateway

```golang

data, err := Payment.ConfirmPayment(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 


| ApplicationID | string | Application id | 


| body |  PaymentConfirmationRequest | "Request body" 

Use this API to confirm payment after payment gateway accepted payment.

*Success Response:*



Success. Returns the status of payment. Check the example shown below or refer `PaymentConfirmationResponseSchema` for more details.


Schema: `PaymentConfirmationResponse`









---



---


## Order


#### shipmentStatusUpdate
Update status of Shipment

```golang

data, err := Order.ShipmentStatusUpdate(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| body |  UpdateShipmentStatusBody | "Request body" 

Update Shipment Status

*Success Response:*



Success


Schema: `UpdateShipmentStatusResponse`









---


#### activityStatus
Get Activity Status

```golang

data, err := Order.ActivityStatus(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 



| xQuery | struct | Includes properties such as `BagID`


Get Activity Status

*Success Response:*



Success


Schema: `GetActivityStatus`









---


#### storeProcessShipmentUpdate
Update Store Process-Shipment

```golang

data, err := Order.StoreProcessShipmentUpdate(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| body |  UpdateProcessShipmenstRequestBody | "Request body" 

Update Store Process-Shipment

*Success Response:*



Success


Schema: `UpdateProcessShipmenstRequestResponse`









---


#### checkRefund
Check Refund is available or not

```golang

data, err := Order.CheckRefund(CompanyID, ShipmentID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ShipmentID | string | Shipment Id | 



Check Refund is available or not

*Success Response:*



Success


Schema: `Object`









---


#### ShipmentBagsCanBreak
Decides if Shipment bags can break

```golang

data, err := Order.ShipmentBagsCanBreak(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| body |  CanBreakRequestBody | "Request body" 

Decides if Shipment bags can break

*Success Response:*



Success


Schema: `CanBreakResponse`









---


#### getOrdersByCompanyId
Get Orders for company based on Company Id

```golang

data, err := Order.GetOrdersByCompanyId(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 



























| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `FromDate`, `ToDate`, `Q`, `Stage`, `SalesChannels`, `OrderID`, `Stores`, `Status`, `Dp`, `ShortenUrls`, `FilterType`


Get Orders

*Success Response:*



Success


Schema: `OrderListing`









---


#### getOrderLanesCountByCompanyId
Get Order Lanes Count for company based on Company Id

```golang

data, err := Order.GetOrderLanesCountByCompanyId(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 

























| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `FromDate`, `ToDate`, `Q`, `Stage`, `SalesChannels`, `OrderID`, `Stores`, `Status`, `ShortenUrls`, `FilterType`


Get Orders Seperate Lane Count

*Success Response:*



Success


Schema: `OrderLanesCount`









---


#### getOrderDetails
Get Order Details for company based on Company Id and Order Id

```golang

data, err := Order.GetOrderDetails(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 







| xQuery | struct | Includes properties such as `OrderID`, `Next`, `Previous`


Get Orders

*Success Response:*



Success


Schema: `OrderDetails`









---


#### getPicklistOrdersByCompanyId
Get Orders for company based on Company Id

```golang

data, err := Order.GetPicklistOrdersByCompanyId(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 

























| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `FromDate`, `ToDate`, `Q`, `Stage`, `SalesChannels`, `OrderID`, `Stores`, `Status`, `ShortenUrls`, `FilterType`


Get Orders

*Success Response:*



Success


Schema: `OrderPicklistListing`









---


#### trackShipmentPlatform
Track Shipment by shipment id, for application based on application Id

```golang

data, err := Order.TrackShipmentPlatform(CompanyID, ApplicationID, ShipmentID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ApplicationID | string | Application Id | 


| ShipmentID | string | Shipment Id | 



Shipment Track

*Success Response:*



Success


Schema: `PlatformShipmentTrack`









---


#### trackOrder
Track Order by order id, for application based on application Id

```golang

data, err := Order.TrackOrder(CompanyID, ApplicationID, OrderID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ApplicationID | string | Application Id | 


| OrderID | string | Order Id | 



Order Track

*Success Response:*



Success


Schema: `PlatformOrderTrack`









---


#### failedOrders
Get all failed orders application wise

```golang

data, err := Order.FailedOrders(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ApplicationID | string | Application Id | 



Failed Orders

*Success Response:*



Success


Schema: `FailedOrders`









---


#### reprocessOrder
Reprocess order by order id

```golang

data, err := Order.ReprocessOrder(CompanyID, ApplicationID, OrderID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ApplicationID | string | Application Id | 


| OrderID | string | Order Id | 



Order Reprocess

*Success Response:*



Success


Schema: `UpdateOrderReprocessResponse`









---


#### updateShipment
Use this API to update the shipment using its shipment ID.

```golang

data, err := Order.UpdateShipment(CompanyID, ApplicationID, ShipmentID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ApplicationID | string | Application Id | 


| ShipmentID | string | ID of the shipment. An order may contain multiple items and may get divided into one or more shipment, each having its own ID. | 


| body |  ShipmentUpdateRequest | "Request body" 

Update the shipment

*Success Response:*



Success. Check the example shown below or refer `ShipmentUpdateRequest` for more details.


Schema: `ShipmentUpdateResponse`









---


#### getPlatformShipmentReasons
Use this API to retrieve the issues that led to the cancellation of bags within a shipment.

```golang

data, err := Order.GetPlatformShipmentReasons(CompanyID, ApplicationID, Action);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ApplicationID | string | Application Id | 


| Action | string | Action | 



Get reasons behind full or partial cancellation of a shipment

*Success Response:*



Success. Check the example shown below or refer `ShipmentReasonsResponse` for more details.


Schema: `ShipmentReasonsResponse`









---


#### getShipmentTrackDetails
Use this API to track a shipment using its shipment ID.

```golang

data, err := Order.GetShipmentTrackDetails(CompanyID, ApplicationID, OrderID, ShipmentID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ApplicationID | string | Application Id | 


| OrderID | string | ID of the order. | 


| ShipmentID | string | ID of the shipment. An order may contain multiple items and may get divided into one or more shipment, each having its own ID. | 



Track shipment

*Success Response:*



Success. Check the example shown below or refer `ShipmentTrackResponse` for more details.


Schema: `ShipmentTrackResponse`









---


#### getShipmentAddress
Use this API to get address of a shipment using its shipment ID and Address Category.

```golang

data, err := Order.GetShipmentAddress(CompanyID, ShipmentID, AddressCategory);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ShipmentID | string | ID of the shipment. An order may contain multiple items and may get divided into one or more shipment, each having its own ID. | 


| AddressCategory | string | Category of the address it falls into(billing or delivery). | 



Get Shipment Address

*Success Response:*



Success. Check the example shown below or refer `GetShipmentAddressResponse` for more details.


Schema: `GetShipmentAddressResponse`









---


#### updateShipmentAddress
Use this API to update address of a shipment using its shipment ID and Address Category.

```golang

data, err := Order.UpdateShipmentAddress(CompanyID, ShipmentID, AddressCategory, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ShipmentID | string | ID of the shipment. An order may contain multiple items and may get divided into one or more shipment, each having its own ID. | 


| AddressCategory | string | Category of the address it falls into(billing or delivery). | 


| body |  UpdateShipmentAddressRequest | "Request body" 

Update Shipment Address

*Success Response:*



Success. Check the example shown below or refer `UpdateShipmentAddressResponse` for more details.


Schema: `UpdateShipmentAddressResponse`









---


#### getPing
Get Ping

```golang

data, err := Order.GetPing(CompanyID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 



Get Ping

*Success Response:*



Success


Schema: `GetPingResponse`









---


#### voiceCallback
Get Voice Callback

```golang

data, err := Order.VoiceCallback(CompanyID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 



Voice Callback

*Success Response:*



Success


Schema: `GetVoiceCallbackResponse`









---


#### voiceClickToCall
Get Voice Click to Call

```golang

data, err := Order.VoiceClickToCall(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 





| xQuery | struct | Includes properties such as `Caller`, `Receiver`


Voice Click to Call

*Success Response:*



Success


Schema: `GetClickToCallResponse`









---



---


## Catalog


#### updateSearchKeywords
Update Search Keyword

```golang

data, err := Catalog.UpdateSearchKeywords(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 


| ID | string | A `id` is a unique identifier for a particular detail. Pass the `id` of the keywords which you want to delete. | 


| body |  CreateSearchKeyword | "Request body" 

Update Search Keyword by its id. On successful request, returns the updated collection

*Success Response:*



The Collection object. See example below or refer `GetSearchWordsDataSchema` for details.


Schema: `GetSearchWordsData`









---


#### deleteSearchKeywords
Delete a Search Keywords

```golang

data, err := Catalog.DeleteSearchKeywords(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 


| ID | string | A `id` is a unique identifier for a particular detail. Pass the `id` of the keywords which you want to delete. | 



Delete a keywords by it's id. Returns an object that tells whether the keywords was deleted successfully

*Success Response:*



Status object. Tells whether the operation was successful. See example below or refer `DeleteResponse`


Schema: `DeleteResponse`









---


#### getSearchKeywords
Get a Search Keywords Details

```golang

data, err := Catalog.GetSearchKeywords(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 


| ID | string | A `id` is a unique identifier for a particular detail. Pass the `id` of the keywords which you want to retrieve. | 



Get the details of a words by its `id`. If successful, returns a Collection resource in the response body specified in `GetSearchWordsDetailResponseSchema`

*Success Response:*



The Collection object. See example below or refer `GetSearchWordsDetailResponseSchema` for details


Schema: `GetSearchWordsDetailResponse`









---


#### createCustomKeyword
Add a Custom Search Keywords

```golang

data, err := Catalog.CreateCustomKeyword(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 


| body |  CreateSearchKeyword | "Request body" 

Create a Custom Search Keywords. See `CreateSearchKeywordSchema` for the list of attributes needed to create a mapping and /collections/query-options for the available options to create a rule. On successful request, returns a paginated list of collections specified in `CreateSearchKeywordSchema`

*Success Response:*



Get keyword object with id that is added. See example below or refer `GetSearchWordsDataSchema` for details


Schema: `GetSearchWordsData`









---


#### getAllSearchKeyword
List all Search Custom Keyword Listing

```golang

data, err := Catalog.GetAllSearchKeyword(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 



Custom Search Keyword allows you to map conditions with keywords to give you the ultimate results

*Success Response:*



List of custom search keywords. See example below or refer `GetSearchWordsResponseSchema` for details


Schema: `GetSearchWordsResponse`









---


#### updateAutocompleteKeyword
Create & Update Autocomplete Keyword

```golang

data, err := Catalog.UpdateAutocompleteKeyword(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 


| ID | string | A `id` is a unique identifier for a particular detail. Pass the `id` of the keywords which you want to delete. | 


| body |  CreateAutocompleteKeyword | "Request body" 

Update a mapping by it's id. On successful request, returns the updated Keyword mapping

*Success Response:*



The Mapping object. See example below or refer `GetAutocompleteWordsResponseSchema` for details.


Schema: `GetAutocompleteWordsResponse`









---


#### deleteAutocompleteKeyword
Delete a Autocomplete Keywords

```golang

data, err := Catalog.DeleteAutocompleteKeyword(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 


| ID | string | A `id` is a unique identifier for a particular detail. Pass the `id` of the keywords which you want to delete. | 



Delete a keywords by it's id. Returns an object that tells whether the keywords was deleted successfully

*Success Response:*



Status object. Tells whether the operation was successful. See example below or refer `DeleteResponse`


Schema: `DeleteResponse`









---


#### getAutocompleteKeywordDetail
Get a Autocomplete Keywords Details

```golang

data, err := Catalog.GetAutocompleteKeywordDetail(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 


| ID | string | A `id` is a unique identifier for a particular detail. Pass the `id` of the keywords which you want to retrieve. | 



Get the details of a words by its `id`. If successful, returns a keywords resource in the response body specified in `GetAutocompleteWordsResponseSchema`

*Success Response:*



The mapping object. See example below or refer `GetAutocompleteWordsResponseSchema` for details


Schema: `GetAutocompleteWordsResponse`









---


#### createCustomAutocompleteRule
Add a Custom Autocomplete Keywords

```golang

data, err := Catalog.CreateCustomAutocompleteRule(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 


| body |  CreateAutocompleteKeyword | "Request body" 

Create a Custom Autocomplete Keywords. See `CreateAutocompleteKeywordSchema` for the list of attributes needed to create a mapping and /collections/query-options for the available options to create a rule. On successful request, returns a paginated list of collections specified in `CreateAutocompleteKeywordSchema`

*Success Response:*



List of all the collections including the one you added. See example below or refer `CreateAutocompleteWordsResponseSchema` for details


Schema: `CreateAutocompleteWordsResponse`









---


#### getAutocompleteConfig
List all Autocomplete Keyword Listing

```golang

data, err := Catalog.GetAutocompleteConfig(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 



Custom Autocomplete Keyword allows you to map conditions with keywords to give you the ultimate results

*Success Response:*



List of custom autocomplete keywords. See example below or refer `GetAutocompleteWordsResponseSchema` for details


Schema: `GetAutocompleteWordsResponse`









---


#### createProductBundle
Create Product Bundle

```golang

data, err := Catalog.CreateProductBundle(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| body |  ProductBundleRequest | "Request body" 

Create Product Bundle. See `ProductBundleRequest` for the request body parameter need to create a product bundle. On successful request, returns in `ProductBundleRequest` with id

*Success Response:*



Get bundle with id that is added. See example below or refer `GetProductBundleCreateResponse` for details


Schema: `GetProductBundleCreateResponse`









---


#### getProductBundle
List all Product Bundles

```golang

data, err := Catalog.GetProductBundle(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 



| xQuery | struct | Includes properties such as `Q`


Get all product bundles for a particular company

*Success Response:*



List of bundle configured for a company. See example below or refer `GetProductBundleListingResponse` for details


Schema: `GetProductBundleListingResponse`









---


#### updateProductBundle
Update a Product Bundle

```golang

data, err := Catalog.UpdateProductBundle(CompanyID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ID | string | A `id` is a unique identifier for a particular detail. Pass the `id` of the keywords which you want to delete. | 


| body |  ProductBundleUpdateRequest | "Request body" 

Update a Product Bundle by its id. On successful request, returns the updated product bundle

*Success Response:*



The Collection object. See example below or refer `GetProductBundleCreateResponse` for details.


Schema: `GetProductBundleCreateResponse`









---


#### getProductBundleDetail
Get a particular Product Bundle details

```golang

data, err := Catalog.GetProductBundleDetail(CompanyID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ID | string | A `id` is a unique identifier for a particular detail. Pass the `id` of the keywords which you want to retrieve. | 



Get a particular Bundle details by its `id`. If successful, returns a Product bundle resource in the response body specified in `GetProductBundleResponse`

*Success Response:*



The Collection object. See example below or refer `GetProductBundleResponse` for details


Schema: `GetProductBundleResponse`









---


#### createSizeGuide
Create a size guide.

```golang

data, err := Catalog.CreateSizeGuide(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Id of the company inside which the size guide is to be created. | 


| body |  ValidateSizeGuide | "Request body" 

This API allows to create a size guide associated to a brand.

*Success Response:*



Returns a success response


Schema: `SuccessResponse`









---


#### getSizeGuides
Get list of size guides

```golang

data, err := Catalog.GetSizeGuides(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Id of the company for which the size guides are to be fetched. | 











| xQuery | struct | Includes properties such as `Active`, `Q`, `Tag`, `PageNo`, `PageSize`


This API allows to view all the size guides associated to the seller.

*Success Response:*



Size guide object. See example below or refer `ListSizeGuide` for details


Schema: `ListSizeGuide`









---


#### updateSizeGuide
Edit a size guide.

```golang

data, err := Catalog.UpdateSizeGuide(CompanyID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Id of the company. | 


| ID | string | Mongo id of the size guide to be edited | 


| body |  ValidateSizeGuide | "Request body" 

This API allows to edit a size guide.

*Success Response:*



Returns a success response


Schema: `SuccessResponse`









---


#### getSizeGuide
Get a single size guide.

```golang

data, err := Catalog.GetSizeGuide(CompanyID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Id of the company associated to size guide. | 


| ID | string | Id of the size guide to be viewed. | 



This API helps to get data associated to a size guide.

*Success Response:*



Brand object. See example below or refer `SizeGuideResponseSchema` for details


Schema: `SizeGuideResponse`









---


#### getCatalogConfiguration
Get configuration meta  details for catalog for admin panel

```golang

data, err := Catalog.GetCatalogConfiguration(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 



configuration meta  details for catalog.

*Success Response:*



configuration details for catalog. See example below or refer `GetCatalogConfigurationMetaDataSchema` for details


Schema: `GetCatalogConfigurationMetaData`









---


#### createConfigurationProductListing
Add configuration for products & listings

```golang

data, err := Catalog.CreateConfigurationProductListing(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 


| body |  AppConfiguration | "Request body" 

Add configuration for products & listing.

*Success Response:*



success flag will tell whether the operation was successful.


Schema: `GetAppCatalogConfiguration`









---


#### getConfigurations
Get configured details for catalog

```golang

data, err := Catalog.GetConfigurations(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 



configured details for catalog.

*Success Response:*



Get application level configured catalog details. See example below or refer `GetAppCatalogConfigurationSchema` for details


Schema: `GetAppCatalogConfiguration`









---


#### createConfigurationByType
Add configuration for categories and brands

```golang

data, err := Catalog.CreateConfigurationByType(CompanyID, ApplicationID, Type, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 


| Type | string | type can be brands, categories etc. | 


| body |  AppConfiguration | "Request body" 

Add configuration for categories & brands.

*Success Response:*



success flag will tell whether the operation was successful.


Schema: `GetAppCatalogConfiguration`









---


#### getConfigurationByType
Get configured details for catalog

```golang

data, err := Catalog.GetConfigurationByType(CompanyID, ApplicationID, Type);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 


| Type | string | type can be brands, categories etc. | 



configured details for catalog.

*Success Response:*



Get application level configured catalog details. See example below or refer `GetAppCatalogEntityConfigurationSchema` for details


Schema: `GetAppCatalogEntityConfiguration`









---


#### getQueryFilters
Get query filters to configure a collection

```golang

data, err := Catalog.GetQueryFilters(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 



Get query filters to configure a collection

*Success Response:*



The attached items of an collection. See example below or refer `GetCollectionQueryOptionResponse` for details


Schema: `GetCollectionQueryOptionResponse`









---


#### createCollection
Add a Collection

```golang

data, err := Catalog.CreateCollection(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 


| body |  CreateCollection | "Request body" 

Create a collection. See `CreateCollectionRequestSchema` for the list of attributes needed to create a collection and collections/query-options for the available options to create a collection. On successful request, returns a paginated list of collections specified in `CollectionCreateResponse`

*Success Response:*



List of all the collections including the one you added. See example below or refer `CollectionCreateResponse` for details


Schema: `CollectionCreateResponse`









---


#### getAllCollections
List all the collections

```golang

data, err := Catalog.GetAllCollections(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 



A Collection allows you to organize your products into hierarchical groups. For example, a dress might be in the category _Clothing_, the individual product might also be in the collection _Summer_. On successful request, returns all the collections as specified in `CollectionListingSchema`

*Success Response:*



List of collections. See example below or refer `GetCollectionListingResponse` for details


Schema: `GetCollectionListingResponse`









---


#### getCollectionDetail
Get a particular collection

```golang

data, err := Catalog.GetCollectionDetail(CompanyID, ApplicationID, Slug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 


| Slug | string | A `slug` is a human readable, URL friendly unique identifier of an object. Pass the `slug` of the collection which you want to retrieve. | 



Get the details of a collection by its `slug`. If successful, returns a Collection resource in the response body specified in `CollectionDetailResponse`

*Success Response:*



The Collection object. See example below or refer `CollectionDetailResponse` for details


Schema: `CollectionDetailResponse`









---


#### updateCollection
Update a collection

```golang

data, err := Catalog.UpdateCollection(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 


| ID | string | A `id` is a unique identifier of a collection. | 


| body |  UpdateCollection | "Request body" 

Update a collection by it's id. On successful request, returns the updated collection

*Success Response:*



The Collection object. See example below or refer `UpdateCollectionSchema` for details.


Schema: `UpdateCollection`









---


#### deleteCollection
Delete a Collection

```golang

data, err := Catalog.DeleteCollection(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 


| ID | string | A `id` is a unique identifier of a collection. | 



Delete a collection by it's id. Returns an object that tells whether the collection was deleted successfully

*Success Response:*



Status object. Tells whether the operation was successful. See example below or refer `DeleteResponse`


Schema: `DeleteResponse`









---


#### addCollectionItems
Add items to a collection

```golang

data, err := Catalog.AddCollectionItems(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 


| ID | string | A `id` is a unique identifier of a collection. | 


| body |  CollectionItemRequest | "Request body" 

Adds items to a collection specified by its `id`. See `CollectionItemRequest` for the list of attributes needed to add items to an collection.

*Success Response:*



Status object. Tells whether the operation was successful.


Schema: `UpdatedResponse`









---


#### getCollectionItems
Get the items for a collection

```golang

data, err := Catalog.GetCollectionItems(CompanyID, ApplicationID, ID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 


| ID | string | A `id` is a unique identifier of a collection. | 







| xQuery | struct | Includes properties such as `SortOn`, `PageID`, `PageSize`


Get items from a collection specified by its `id`.

*Success Response:*



The attached items of an collection. See example below or refer `GetCollectionItemsResponseSchema` for details


Schema: `GetCollectionItemsResponse`









---


#### getCatalogInsights
Analytics data of catalog and inventory.

```golang

data, err := Catalog.GetCatalogInsights(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 



| xQuery | struct | Includes properties such as `Brand`


Catalog Insights api returns the count of catalog related data like products, brands, departments and categories that have been made live as per configuration of the app.

*Success Response:*



Response Data


Schema: `CatalogInsightResponse`









---


#### getSellerInsights
Analytics data of catalog and inventory that are being cross-selled.

```golang

data, err := Catalog.GetSellerInsights(CompanyID, SellerAppID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| SellerAppID | string | Id of the seller application which is serving the invetory/catalog of the company | 



Analytics data of catalog and inventory that are being cross-selled.

*Success Response:*



Response Data


Schema: `CrossSellingResponse`









---


#### createMarketplaceOptin
Create/Update opt-in infomation.

```golang

data, err := Catalog.CreateMarketplaceOptin(CompanyID, Marketplace, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | The company id for which the detail needs to be retrieved. | 


| Marketplace | string | The marketplace for which the detail needs to be retrieved. | 


| body |  OptInPostRequest | "Request body" 

Use this API to create/update opt-in information for given platform. If successful, returns data in the response body as specified in `OptInPostResponseSchema`

*Success Response:*



See example below or refer `UpdatedResponse` for details.


Schema: `UpdatedResponse`









---


#### getMarketplaceOptinDetail
Get opt-in infomation.

```golang

data, err := Catalog.GetMarketplaceOptinDetail(CompanyID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 |  | 



Use this API to fetch opt-in information for all the platforms. If successful, returns a logs in the response body as specified in `GetOptInPlatformSchema`

*Success Response:*



See example below or refer `GetOptInPlatformSchema` for details.


Schema: `GetOptInPlatform`









---


#### getCompanyDetail
Get the Company details.

```golang

data, err := Catalog.GetCompanyDetail(CompanyID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | The company id for which the detail needs to be retrieved. | 



Get the details of the company associated with the given company_id passed.

*Success Response:*



See example below or refer `OptinCompanyDetailSchema` for details


Schema: `OptinCompanyDetail`









---


#### getCompanyBrandDetail
Get the Company Brand details of Optin.

```golang

data, err := Catalog.GetCompanyBrandDetail(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | The company id for which the detail needs to be retrieved. | 











| xQuery | struct | Includes properties such as `IsActive`, `Q`, `PageNo`, `PageSize`, `Marketplace`


Get the details of the Brands associated with the given company_id passed.

*Success Response:*



See example below or refer `OptinCompanyBrandDetailsView` for details


Schema: `OptinCompanyBrandDetailsView`









---


#### getCompanyMetrics
Get the Company metrics

```golang

data, err := Catalog.GetCompanyMetrics(CompanyID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | The company id for which the detail needs to be retrieved. | 



Get the Company metrics associated with the company ID passed.

*Success Response:*



See example below or refer `OptinCompanyMetrics` for details


Schema: `OptinCompanyMetrics`









---


#### getStoreDetail
Get the Store details.

```golang

data, err := Catalog.GetStoreDetail(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | The company id for which the detail needs to be retrieved. | 







| xQuery | struct | Includes properties such as `Q`, `PageNo`, `PageSize`


Get the details of the store associated with the company ID passed.

*Success Response:*



See example below or refer `OptinStoreDetailsSchema` for details


Schema: `OptinStoreDetails`









---


#### getGenderAttribute
Get gender attribute details

```golang

data, err := Catalog.GetGenderAttribute(CompanyID, AttributeSlug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | company for which you want to view the genders | 


| AttributeSlug | string | slug of the attribute for which you want to view the genders | 



This API allows to view the gender attribute details.

*Success Response:*



Size guide object. See example below or refer `GenderDetailSchema` for details


Schema: `GenderDetail`









---


#### listProductTemplateCategories
List Department specifiec product categories

```golang

data, err := Catalog.ListProductTemplateCategories(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 





| xQuery | struct | Includes properties such as `Departments`, `ItemType`


Allows you to list all product categories values for the departments specified

*Success Response:*



List of all categories attached to departments specified. See example below or refer `ProdcutTemplateCategoriesResponse` for details


Schema: `ProdcutTemplateCategoriesResponse`









---


#### listDepartmentsData
List all Departments

```golang

data, err := Catalog.ListDepartmentsData(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 











| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Name`, `Search`, `IsActive`


Allows you to list all departments, also can search using name and filter active and incative departments, and item type

*Success Response:*



List of departments data. See example below or refer `DepartmentsResponse` for details


Schema: `DepartmentsResponse`









---


#### getDepartmentData
Get specific departments details by passing in unique id of the department

```golang

data, err := Catalog.GetDepartmentData(CompanyID, UID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| UID | string | A `uid` is a unique identifier of a department. | 



Allows you to get department data, by uid

*Success Response:*



Departments Data. See example below or refer `DepartmentsResponse` for details


Schema: `DepartmentsResponse`









---


#### listProductTemplate
List all Templates

```golang

data, err := Catalog.ListProductTemplate(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 



| xQuery | struct | Includes properties such as `Department`


Allows you to list all product templates, also can filter by department

*Success Response:*



List of product templates. See example below or refer `TemplatesResponse` for details


Schema: `TemplatesResponse`









---


#### validateProductTemplate
Validate Product Template Schema

```golang

data, err := Catalog.ValidateProductTemplate(CompanyID, Slug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| Slug | string | A `slug` is a unique identifier for a particular template. | 



Allows you to list all product templates validation values for all the fields present in the database

*Success Response:*



List of fields and validation values fro each. See example below or refer `TemplatesValidationResponse` for details


Schema: `TemplatesValidationResponse`









---


#### downloadProductTemplateViews
Download Product Template View

```golang

data, err := Catalog.DownloadProductTemplateViews(CompanyID, Slug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| Slug | string | A `slug` is a unique identifier for a particular template. | 



Allows you to download product template data

*Success Response:*



CSV File of product template data. See example below or refer `TemplatesResponse` for details


Schema: `string`









---


#### downloadProductTemplateView
Download Product Template View

```golang

data, err := Catalog.DownloadProductTemplateView(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 



| xQuery | struct | Includes properties such as `ItemType`


Allows you to download product template data

*Success Response:*



CSV File of product template data.


Schema: `string`









---


#### validateProductTemplateSchema
Validate Product Template Schema

```golang

data, err := Catalog.ValidateProductTemplateSchema(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 



| xQuery | struct | Includes properties such as `ItemType`


Allows you to list all product templates validation values for all the fields present in the database

*Success Response:*



List of fields and validation values fro each. See example below or refer `InventoryValidationResponse` for details


Schema: `InventoryValidationResponse`









---


#### listHSNCodes
List HSN Codes

```golang

data, err := Catalog.ListHSNCodes(CompanyID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 



Allows you to list all hsn Codes

*Success Response:*



List of all HSN Codes. See example below or refer `HSNCodesResponse` for details


Schema: `HSNCodesResponse`









---


#### listProductTemplateExportDetails
Allows you to list all product templates export list details

```golang

data, err := Catalog.ListProductTemplateExportDetails(CompanyID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 



Can view details including trigger data, task id , etc.

*Success Response:*



List of Product Downloads Data. See example below or refer `ProductDownloadsResponse` for details


Schema: `ProductDownloadsResponse`









---


#### listTemplateBrandTypeValues
Allows you to list all values for Templates, Brands or Type

```golang

data, err := Catalog.ListTemplateBrandTypeValues(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 



| xQuery | struct | Includes properties such as `Filter`


The filter type query parameter defines what type of data to return. The type of query returns the valid values for the same

*Success Response:*



See example below or refer `ProductConfigurationDownloadsSchema` for details


Schema: `ProductConfigurationDownloads`









---


#### createCategories
Create product categories

```golang

data, err := Catalog.CreateCategories(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| body |  CategoryRequestBody | "Request body" 

This API lets user create product categories

*Success Response:*



Category Meta. See example below or refer `CategoryCreateResponse` for details


Schema: `CategoryCreateResponse`









---


#### listCategories
Get product categories list

```golang

data, err := Catalog.ListCategories(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 











| xQuery | struct | Includes properties such as `Level`, `Departments`, `Q`, `PageNo`, `PageSize`


This API gets meta associated to product categories.

*Success Response:*



Category Meta. See example below or refer `CategoryResponse` for details


Schema: `CategoryResponse`









---


#### updateCategory
Update product categories

```golang

data, err := Catalog.UpdateCategory(CompanyID, UID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| UID | string | Category unique id | 


| body |  CategoryRequestBody | "Request body" 

Update a product category using this apu

*Success Response:*



Category Meta. See example below or refer `CategoryUpdateResponse` for details


Schema: `CategoryUpdateResponse`









---


#### getCategoryData
Get product category by uid

```golang

data, err := Catalog.GetCategoryData(CompanyID, UID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| UID | string | Category unique id | 



This API gets meta associated to product categories.

*Success Response:*



Get Data for one category. See example below or refer `CategoryResponse` for details


Schema: `SingleCategoryResponse`









---


#### createProduct
Create a product.

```golang

data, err := Catalog.CreateProduct(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Id of the company associated to product that is to be viewed. | 


| body |  ProductCreateUpdate | "Request body" 

This API allows to create product.

*Success Response:*



Returns a success response


Schema: `SuccessResponse`









---


#### getProducts
Get product list

```golang

data, err := Catalog.GetProducts(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Get list of products filtered by company Id | 



















| xQuery | struct | Includes properties such as `BrandIds`, `CategoryIds`, `ItemIds`, `DepartmentIds`, `ItemCode`, `Q`, `Tags`, `PageNo`, `PageSize`


This API gets meta associated to products.

*Success Response:*



Product Meta. See example below for details


Schema: `ProductListingResponse`









---


#### editProduct
Edit a product.

```golang

data, err := Catalog.EditProduct(CompanyID, ItemID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Id of the company associated to product that is to be viewed. | 


| ItemID | float64 | Id of the product to be updated. | 


| body |  ProductCreateUpdate | "Request body" 

This API allows to edit product.

*Success Response:*



Returns a success response


Schema: `SuccessResponse`









---


#### deleteProduct
Delete a product.

```golang

data, err := Catalog.DeleteProduct(CompanyID, ItemID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id of the company associated to product that is to be deleted. | 


| ItemID | float64 | Id of the product to be updated. | 



This API allows to delete product.

*Success Response:*



Returns a success response


Schema: `SuccessResponse`









---


#### getProduct
Get a single product.

```golang

data, err := Catalog.GetProduct(CompanyID, ItemID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



| CompanyID | float64 | Company Id of the product. | 


| ItemID | float64 | Item Id of the product. | 





| xQuery | struct | Includes properties such as `ItemCode`, `BrandUID`, `UID`


This API helps to get data associated to a particular product.

*Success Response:*



Product object. See example below or refer `product.utils.format_product_response` for details


Schema: `Product`









---


#### getProductValidation
Validate product/size data

```golang

data, err := Catalog.GetProductValidation(CompanyID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Validates data against given company | 



This API validates product data.

*Success Response:*



Validate Meta. See example below for details


Schema: `ValidateProduct`









---


#### getProductSize
Get a single product size.

```golang

data, err := Catalog.GetProductSize(CompanyID, ItemID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



| CompanyID | float64 | Company Id of the product size. | 


| ItemID | float64 | Item Id of the product size. | 





| xQuery | struct | Includes properties such as `ItemCode`, `BrandUID`, `UID`


This API helps to get data associated to a particular product size.

*Success Response:*



Product object. See example below for details


Schema: `ProductListingResponse`









---


#### updateProductAssetsInBulk
Create a Bulk asset upload Job.

```golang

data, err := Catalog.UpdateProductAssetsInBulk(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id in which assets to be uploaded. | 


| body |  BulkJob | "Request body" 

This API helps to create a bulk asset upload job.

*Success Response:*



Returns a success response


Schema: `SuccessResponse`









---


#### getProductBulkUploadHistory
Get a list of all bulk product upload jobs.

```golang

data, err := Catalog.GetProductBulkUploadHistory(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id of of which Product Bulk Upload History to be obtained. | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`


This API helps to get bulk product upload jobs data.

*Success Response:*



List of bulk product upload jobs. See `BulkRequestGetSchema` for details


Schema: `ProductBulkRequestList`









---


#### createProductsInBulk
Create products in bulk associated with given batch Id.

```golang

data, err := Catalog.CreateProductsInBulk(CompanyID, BatchID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id in which assets to be uploaded. | 


| BatchID | string | Batch Id in which assets to be uploaded. | 


| body |  BulkProductRequest | "Request body" 

This API helps to create products in bulk push to kafka for approval/creation.

*Success Response:*



Returns a success response


Schema: `SuccessResponse`









---


#### deleteProductBulkJob
Delete Bulk product job.

```golang

data, err := Catalog.DeleteProductBulkJob(CompanyID, BatchID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id of the company associated to size that is to be deleted. | 


| BatchID | float64 | Batch Id of the bulk product job to be deleted. | 



This API allows to delete bulk product job associated with company.

*Success Response:*



Returns a success response


Schema: `SuccessResponse`









---


#### getProductTags
Get a list of all tags associated with company.

```golang

data, err := Catalog.GetProductTags(CompanyID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id for which tags to be fetched. | 



This API helps to get tags data associated to a particular company.

*Success Response:*



Tag List. See example below for details


Schema: `ProductTagsViewResponse`









---


#### createProductAssetsInBulk
Create a Bulk asset upload Job.

```golang

data, err := Catalog.CreateProductAssetsInBulk(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id in which assets to be uploaded. | 


| body |  ProductBulkAssets | "Request body" 

This API helps to create a bulk asset upload job.

*Success Response:*



Returns a success response


Schema: `SuccessResponse`









---


#### getProductAssetsInBulk
Get a list of all bulk asset jobs.

```golang

data, err := Catalog.GetProductAssetsInBulk(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id of the product size. | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`


This API helps to get bulk asset jobs data associated to a particular company.

*Success Response:*



List of bulk asset jobs List. See `BulkUtil.modify_batch_response` for details


Schema: `BulkAssetResponse`









---


#### deleteSize
Delete a Size associated with product.

```golang

data, err := Catalog.DeleteSize(CompanyID, ItemID, Size);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id of the company associated to size that is to be deleted. | 


| ItemID | float64 | Item Id of the product associated with size to be deleted. | 


| Size | float64 | size to be deleted. | 



This API allows to delete size associated with product.

*Success Response:*



Returns a success response


Schema: `ProductSizeDeleteResponse`









---


#### addInventory
Add Inventory for particular size and store.

```golang

data, err := Catalog.AddInventory(CompanyID, ItemID, Size, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Id of the company associated to product that is to be viewed. | 


| ItemID | float64 | Item code of the product of which size is to be get. | 


| Size | string | Size in which inventory is to be added. | 


| body |  InventoryRequest | "Request body" 

This API allows add Inventory for particular size and store.

*Success Response:*



Returns a success response


Schema: `SuccessResponse`









---


#### getInventoryBySize
Get Inventory for company

```golang

data, err := Catalog.GetInventoryBySize(CompanyID, ItemID, Size, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Id of the company associated to product that is to be viewed. | 


| ItemID | string | Item code of the product of which size is to be get. | 


| Size | string | Size of which inventory is to get. | 







| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Q`


This API allows get Inventory data for particular company grouped by size and store.

*Success Response:*



returns a list of all inventory grouped by size and store


Schema: `InventoryResponse`









---


#### getInventoryBySizeIdentifier
Get Inventory for company

```golang

data, err := Catalog.GetInventoryBySizeIdentifier(CompanyID, ItemID, SizeIdentifier, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Id of the company associated to product that is to be viewed. | 


| ItemID | string | Item code of the product of which size is to be get. | 


| SizeIdentifier | string | Size Identifier (Seller Identifier or Primary Identifier) of which inventory is to get. | 









| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Q`, `LocationIds`


This API allows get Inventory data for particular company grouped by size and store.

*Success Response:*



returns a list of all inventory grouped by size and store


Schema: `InventoryResponse`









---


#### deleteInventory
Delete a Inventory.

```golang

data, err := Catalog.DeleteInventory(CompanyID, Size, ItemID, LocationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id of the company associated with Inventory that is to be deleted. | 


| Size | string | size that is to be deleted. | 


| ItemID | float64 | Id of the product associated with Inventory to be deleted. | 


| LocationID | float64 | Location ID of store of which inventory is to be deleted. | 



This API allows to delete inventory of a particular product for particular company.

*Success Response:*



Returns a success response


Schema: `InventoryDelete`









---


#### createBulkInventoryJob
Create a Bulk Inventory upload Job.

```golang

data, err := Catalog.CreateBulkInventoryJob(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id in which Inventory to be uploaded. | 


| body |  BulkJob | "Request body" 

This API helps to create a bulk Inventory upload job.

*Success Response:*



Returns a success response


Schema: `CommonResponse`









---


#### getInventoryBulkUploadHistory
Get a list of all bulk Inventory upload jobs.

```golang

data, err := Catalog.GetInventoryBulkUploadHistory(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id of of which Inventory Bulk Upload History to be obtained. | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`


This API helps to get bulk Inventory upload jobs data.

*Success Response:*



List of bulk Inventory upload jobs. See `BulkInventoryGetSchema` for details


Schema: `BulkInventoryGet`









---


#### createBulkInventory
Create products in bulk associated with given batch Id.

```golang

data, err := Catalog.CreateBulkInventory(CompanyID, BatchID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id in which Inventory is to be uploaded. | 


| BatchID | string | Batch Id of the bulk create job. | 


| body |  InventoryBulkRequest | "Request body" 

This API helps to create products in bulk push to kafka for approval/creation.

*Success Response:*



Returns a success response


Schema: `SuccessResponse`









---


#### deleteBulkInventoryJob
Delete Bulk Inventory job.

```golang

data, err := Catalog.DeleteBulkInventoryJob(CompanyID, BatchID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id of the company of which bulk Inventory job is to be deleted. | 


| BatchID | string | Batch Id of the bulk delete job. | 



This API allows to delete bulk Inventory job associated with company.

*Success Response:*



Returns a success response


Schema: `SuccessResponse`









---


#### createInventoryExportJob
Create a Inventory export Job.

```golang

data, err := Catalog.CreateInventoryExportJob(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id in which assets to be uploaded. | 


| body |  InventoryExportRequest | "Request body" 

This API helps to create a Inventory export job.

*Success Response:*



Returns a success response


Schema: `SuccessResponse`









---


#### getInventoryExport
Get Inventory export history.

```golang

data, err := Catalog.GetInventoryExport(CompanyID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id in which assets to be uploaded. | 



This API helps to get Inventory export history.

*Success Response:*



Returns a list of inventory export jobs


Schema: `InventoryExportJob`









---


#### exportInventoryConfig
Get List of different filters for inventory export

```golang

data, err := Catalog.ExportInventoryConfig(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Id of the company associated to product that is to be viewed. | 



| xQuery | struct | Includes properties such as `FilterType`


This API allows get List of different filters like brand, store, and type for inventory export.

*Success Response:*



returns filters configuration for inventory export


Schema: `InventoryConfig`









---


#### createHsnCode
Create Hsn Code.

```golang

data, err := Catalog.CreateHsnCode(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| body |  HsnUpsert | "Request body" 

Create Hsn Code.

*Success Response:*



See example below for details


Schema: `HsnCode`









---


#### getAllHsnCodes
Hsn Code List.

```golang

data, err := Catalog.GetAllHsnCodes(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 







| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Q`


Hsn Code List.

*Success Response:*



List of all HSN Codes. See example below or refer `HsnCodesListingResponseSchema` for details


Schema: `HsnCodesListingResponse`









---


#### updateHsnCode
Update Hsn Code.

```golang

data, err := Catalog.UpdateHsnCode(CompanyID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| ID | string | Unique id | 


| body |  HsnUpsert | "Request body" 

Update Hsn Code.

*Success Response:*



See example below for details


Schema: `HsnCode`









---


#### getHsnCode
Fetch Hsn Code.

```golang

data, err := Catalog.GetHsnCode(CompanyID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| ID | string | Unique id | 



Fetch Hsn Code.

*Success Response:*



See example below details


Schema: `HsnCode`









---


#### bulkHsnCode
Bulk Create or Update Hsn Code.

```golang

data, err := Catalog.BulkHsnCode(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| body |  BulkHsnUpsert | "Request body" 

Bulk Create or Update Hsn Code.

*Success Response:*



See example below for details


Schema: `BulkHsnResponse`









---


#### getApplicationBrands
List all the brands

```golang

data, err := Catalog.GetApplicationBrands(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 







| xQuery | struct | Includes properties such as `Department`, `PageNo`, `PageSize`


A brand is the name under which a product is being sold. Use this API to list all the brands. You can pass optionally filter the brands by the department. If successful, returns a paginated list of brands specified in `BrandListingResponse`

*Success Response:*



List of Brands. See example below or refer `BrandListingResponse` for details


Schema: `BrandListingResponse`









---


#### getDepartments
List all the departments

```golang

data, err := Catalog.GetDepartments(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 



Departments are a way to categorise similar products. A product can lie in multiple departments. For example, a skirt can below to the 'Women's Fashion' Department while a handbag can lie in 'Women's Accessories' Department. Use this API to list all the departments. If successful, returns the list of departments specified in `DepartmentResponse`

*Success Response:*



List of Departments. See example below or refer `DepartmentResponse` for details.


Schema: `DepartmentResponse`









---


#### getCategories
List all the categories

```golang

data, err := Catalog.GetCategories(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 



| xQuery | struct | Includes properties such as `Department`


List all the categories. You can optionally pass filter the brands by the department. If successful, returns a paginated list of brands specified in `CategoryListingResponse`

*Success Response:*



List of Categories. See example below or refer `CategoryListingResponse` for details.


Schema: `CategoryListingResponse`









---


#### getAppicationProducts
List the products

```golang

data, err := Catalog.GetAppicationProducts(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 



















| xQuery | struct | Includes properties such as `Q`, `F`, `Filters`, `SortOn`, `PageID`, `PageSize`, `PageNo`, `PageType`, `ItemIds`


List all the products associated with a brand, collection or category in a requested sort order. The API additionally supports arbitrary search queries that may refer the name of any product, brand, category or collection. If successful, returns a paginated list of products specified in `ApplicationProductListingResponse`

*Success Response:*



List of Products. See example below or refer `ApplicationProductListingResponse` for details


Schema: `ApplicationProductListingResponse`









---


#### getProductDetailBySlug
Get a product

```golang

data, err := Catalog.GetProductDetailBySlug(CompanyID, ApplicationID, Slug);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 


| Slug | string | The unique identifier of a product. i.e; `slug` of a product. You can retrieve these from the APIs that list products like **v1.0/products/** | 



Products are the core resource of an application. Products can be associated by categories, collections, brands and more. This API retrieves the product specified by the given **slug**. If successful, returns a Product resource in the response body specified in `ProductDetail`

*Success Response:*



The Product object. See example below or refer `ProductDetail` for details.


Schema: `ProductDetail`









---


#### getAppProducts
Get applicationwise products

```golang

data, err := Catalog.GetAppProducts(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| ApplicationID | string | A `application_id` is a unique identifier for a particular sale channel. | 















| xQuery | struct | Includes properties such as `BrandIds`, `CategoryIds`, `DepartmentIds`, `Tags`, `PageNo`, `PageSize`, `Q`


Products are the core resource of an application. Products can be associated by categories, collections, brands and more. If successful, returns a Product resource in the response body specified in `ApplicationProductListingResponseDatabasePowered`

*Success Response:*



The Product object. See example below or refer `ApplicationProductListingResponseDatabasePowered` for details.


Schema: `ProductListingResponse`









---



---


## CompanyProfile


#### cbsOnboardGet
Get company profile

```golang

data, err := CompanyProfile.CbsOnboardGet(CompanyID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 



This API allows to view the company profile of the seller account.

*Success Response:*



Company profile object. See example below or refer `GetCompanyProfileSerializerResponse` for details


Schema: `GetCompanyProfileSerializerResponse`









---


#### updateCompany
Edit company profile

```golang

data, err := CompanyProfile.UpdateCompany(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 


| body |  UpdateCompany | "Request body" 

This API allows to edit the company profile of the seller account.

*Success Response:*



Returns a success message


Schema: `SuccessResponse`









---


#### getCompanyMetrics
Get company metrics

```golang

data, err := CompanyProfile.GetCompanyMetrics(CompanyID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | A `company_id` is a unique identifier for a particular seller account. | 



This API allows to view the company metrics, i.e. the status of its brand and stores. Also its allows to view the number of products, company documents & store documents which are verified and unverified.

*Success Response:*



Metrics response object. See example below or refer `MetricsSerializer` for details


Schema: `MetricsSerializer`









---


#### getBrand
Get a single brand.

```golang

data, err := CompanyProfile.GetBrand(CompanyID, BrandID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Id of the company associated to brand that is to be viewed. | 


| BrandID | string | Id of the brand to be viewed. | 



This API helps to get data associated to a particular brand.

*Success Response:*



Brand object. See example below or refer `GetBrandResponseSerializer` for details


Schema: `GetBrandResponseSerializer`









---


#### editBrand
Edit a brand.

```golang

data, err := CompanyProfile.EditBrand(CompanyID, BrandID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Id of the company associated to brand that is to be viewed. | 


| BrandID | string | Id of the brand to be viewed. | 


| body |  CreateUpdateBrandRequestSerializer | "Request body" 

This API allows to edit meta of a brand.

*Success Response:*



Returns a success response


Schema: `SuccessResponse`









---


#### createBrand
Create a Brand.

```golang

data, err := CompanyProfile.CreateBrand(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Id of the company. | 


| body |  CreateUpdateBrandRequestSerializer | "Request body" 

This API allows to create a brand associated to a company.

*Success Response:*



Returns a success response


Schema: `SuccessResponse`









---


#### getBrands
Get brands associated to a company

```golang

data, err := CompanyProfile.GetBrands(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Id of the company. | 







| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Q`


This API helps to get view brands associated to a particular company.

*Success Response:*



Brand object. See example below or refer `CompanyBrandListSerializer` for details


Schema: `CompanyBrandListSerializer`









---


#### createCompanyBrandMapping
Create a company brand mapping.

```golang

data, err := CompanyProfile.CreateCompanyBrandMapping(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Id of the company inside which the brand is to be mapped. | 


| body |  CompanyBrandPostRequestSerializer | "Request body" 

This API allows to create a company brand mapping, for a already existing brand in the system.

*Success Response:*



Returns a success response


Schema: `SuccessResponse`









---


#### getLocations
Get list of locations

```golang

data, err := CompanyProfile.GetLocations(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Id of the company whose locations are to fetched | 











| xQuery | struct | Includes properties such as `StoreType`, `Q`, `Stage`, `PageNo`, `PageSize`


This API allows to view all the locations asscoiated to a company.

*Success Response:*



Company profile object. See example below or refer `LocationListSerializer` for details


Schema: `LocationListSerializer`









---


#### createLocation
Create a location asscoiated to a company.

```golang

data, err := CompanyProfile.CreateLocation(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Id of the company inside which the location is to be created. | 


| body |  LocationSerializer | "Request body" 

This API allows to create a location associated to a company.

*Success Response:*



Returns a success response


Schema: `SuccessResponse`









---


#### getLocationDetail
Get details of a specific location.

```golang

data, err := CompanyProfile.GetLocationDetail(CompanyID, LocationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Id of the company inside which the location lies. | 


| LocationID | string | Id of the location which you want to view. | 



This API helps to get data associated to a specific location.

*Success Response:*



Brand object. See example below or refer `GetLocationSerializer` for details


Schema: `GetLocationSerializer`









---


#### updateLocation
Edit a location asscoiated to a company.

```golang

data, err := CompanyProfile.UpdateLocation(CompanyID, LocationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Id of the company inside which the location is to be created. | 


| LocationID | string | Id of the location which you want to edit. | 


| body |  LocationSerializer | "Request body" 

This API allows to edit a location associated to a company.

*Success Response:*



Returns a success response


Schema: `SuccessResponse`









---


#### createLocationBulk
Create a location asscoiated to a company in bulk.

```golang

data, err := CompanyProfile.CreateLocationBulk(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Id of the company inside which the location is to be created. | 


| body |  BulkLocationSerializer | "Request body" 

This API allows to create a location associated to a company.

*Success Response:*



Returns a success response


Schema: `SuccessResponse`









---



---


## FileStorage


#### startUpload
This operation initiates upload and returns storage link which is valid for 30 Minutes. You can use that storage link to make subsequent upload request with file buffer or blob.

```golang

data, err := FileStorage.StartUpload(Namespace, CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Namespace | string | bucket name | 


| CompanyID | float64 | company_id | 


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









---


#### completeUpload
This will complete the upload process. After successfully uploading file, you can call this operation to complete the upload process.

```golang

data, err := FileStorage.CompleteUpload(Namespace, CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Namespace | string | bucket name | 


| CompanyID | float64 | company_id | 


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









---


#### appStartUpload
This operation initiates upload and returns storage link which is valid for 30 Minutes. You can use that storage link to make subsequent upload request with file buffer or blob.

```golang

data, err := FileStorage.AppStartUpload(Namespace, CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Namespace | string | bucket name | 


| CompanyID | float64 | company_id | 


| ApplicationID | string | application id | 


| body |  StartRequest | "Request body" 

Uploads an arbitrarily sized buffer or blob.

It has three Major Steps:
* Start
* Upload
* Complete

### Start
Initiates the assets upload using `appStartUpload`.
It returns the storage link in response.

### Upload
Use the storage link to upload a file (Buffer or Blob) to the File Storage.
Make a `PUT` request on storage link received from `appStartUpload` api with file (Buffer or Blob) as a request body.

### Complete
After successfully upload, call `appCompleteUpload` api to complete the upload process.
This operation will return the url for the uploaded file.


*Success Response:*



Success


Schema: `StartResponse`









---


#### appCompleteUpload
This will complete the upload process. After successfully uploading file, you can call this operation to complete the upload process.

```golang

data, err := FileStorage.AppCompleteUpload(Namespace, CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Namespace | string | bucket name | 


| CompanyID | float64 | company_id | 


| ApplicationID | string | application id | 


| body |  StartResponse | "Request body" 

Uploads an arbitrarily sized buffer or blob.

It has three Major Steps:
* Start
* Upload
* Complete

### Start
Initiates the assets upload using `appStartUpload`.
It returns the storage link in response.

### Upload
Use the storage link to upload a file (Buffer or Blob) to the File Storage.
Make a `PUT` request on storage link received from `appStartUpload` api with file (Buffer or Blob) as a request body.

### Complete
After successfully upload, call `appCompleteUpload` api to complete the upload process.
This operation will return the url for the uploaded file.


*Success Response:*



Success


Schema: `CompleteResponse`









---


#### getSignUrls
Explain here

```golang

data, err := FileStorage.GetSignUrls(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | company_id | 


| body |  SignUrlRequest | "Request body" 

Describe here

*Success Response:*



Success


Schema: `SignUrlResponse`









---


#### copyFiles
Copy Files

```golang

data, err := FileStorage.CopyFiles(CompanyID, xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



| CompanyID | float64 | company_id | 

| xQuery | struct | Includes properties such as `Sync`

| body |  BulkRequest | "Request body" 

Copy Files

*Success Response:*



Success


Schema: `BulkResponse`









---


#### appCopyFiles
Copy Files

```golang

data, err := FileStorage.AppCopyFiles(CompanyID, ApplicationID, xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |



| CompanyID | float64 | company_id | 


| ApplicationID | float64 | application_id | 

| xQuery | struct | Includes properties such as `Sync`

| body |  BulkRequest | "Request body" 

Copy Files

*Success Response:*



Success


Schema: `BulkResponse`









---


#### browse
Browse Files

```golang

data, err := FileStorage.Browse(Namespace, CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Namespace | string | bucket name | 


| CompanyID | float64 | company_id | 



| xQuery | struct | Includes properties such as `PageNo`


Browse Files

*Success Response:*



Success


Schema: `BrowseResponse`









---


#### browse
Browse Files

```golang

data, err := FileStorage.Browse(Namespace, CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| Namespace | string | bucket name | 


| CompanyID | float64 | company_id | 


| ApplicationID | float64 | application_id | 



| xQuery | struct | Includes properties such as `PageNo`


Browse Files

*Success Response:*



Success


Schema: `BrowseResponse`









---


#### proxy
Proxy

```golang

data, err := FileStorage.Proxy(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | company_id | 



| xQuery | struct | Includes properties such as `URL`


Proxy

*Success Response:*



Success


Schema: `string`









---



---


## Share


#### createShortLink
Create short link

```golang

data, err := Share.CreateShortLink(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ApplicationID | string | Application Id | 


| body |  ShortLinkReq | "Request body" 

Create short link

*Success Response:*



Success


Schema: `ShortLinkRes`









---


#### getShortLinks
Get short links

```golang

data, err := Share.GetShortLinks(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ApplicationID | string | Application Id | 











| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `CreatedBy`, `Active`, `Q`


Get short links

*Success Response:*



Success


Schema: `ShortLinkList`









---


#### getShortLinkByHash
Get short link by hash

```golang

data, err := Share.GetShortLinkByHash(CompanyID, ApplicationID, Hash);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ApplicationID | string | Application Id | 


| Hash | string | Hash of short url | 



Get short link by hash

*Success Response:*



Success


Schema: `ShortLinkRes`









---


#### updateShortLinkById
Update short link by id

```golang

data, err := Share.UpdateShortLinkById(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ApplicationID | string | Application Id | 


| ID | string | Short link document identifier | 


| body |  ShortLinkReq | "Request body" 

Update short link by id

*Success Response:*



Success


Schema: `ShortLinkRes`









---



---


## Inventory


#### getJobsByCompany
Get Job Configs For A Company

```golang

data, err := Inventory.GetJobsByCompany(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`


REST Endpoint that returns all job configs for a company

*Success Response:*



Successful operation


Schema: `ResponseEnvelopeListJobConfigRawDTO`









---


#### updateJob
Updates An Existing Job Config

```golang

data, err := Inventory.UpdateJob(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 


| body |  JobConfigDTO | "Request body" 

REST Endpoint that updates a job config

*Success Response:*



Job Config Updated Successfully


Schema: `ResponseEnvelopeString`









---


#### createJob
Creates A New Job Config

```golang

data, err := Inventory.CreateJob(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 


| body |  JobConfigDTO | "Request body" 

REST Endpoint that creates a new job config

*Success Response:*



Job Config Created Successfully


Schema: `ResponseEnvelopeString`









---


#### getJobByCompanyAndIntegration
Get Job Configs By Company And Integration

```golang

data, err := Inventory.GetJobByCompanyAndIntegration(CompanyID, IntegrationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 


| IntegrationID | string | Integration Id | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`


REST Endpoint that returns all job configs by company And integration

*Success Response:*



Successful operation


Schema: `ResponseEnvelopeListJobConfigDTO`









---


#### getJobConfigDefaults
Get Job Configs Defaults

```golang

data, err := Inventory.GetJobConfigDefaults(CompanyID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 



REST Endpoint that returns default fields job configs by company And integration

*Success Response:*



Successful operation


Schema: `ResponseEnvelopeJobConfigDTO`









---


#### getJobByCode
Get Job Config By Code

```golang

data, err := Inventory.GetJobByCode(CompanyID, Code);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 


| Code | string | Job Code | 



REST Endpoint that returns job config by code

*Success Response:*



Successful operation


Schema: `ResponseEnvelopeJobConfigDTO`









---


#### getJobCodeMetrics
Get Job Metrics

```golang

data, err := Inventory.GetJobCodeMetrics(CompanyID, Code, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 


| Code | string | Code | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`


REST Endpoint that returns Inventory Run History For A Job Code

*Success Response:*



Successful operation


Schema: `ResponseEnvelopeJobMetricsDto`









---


#### getJobCodesByCompanyAndIntegration
Get Job Codes By Company And Integration

```golang

data, err := Inventory.GetJobCodesByCompanyAndIntegration(CompanyID, IntegrationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id | 


| IntegrationID | string | Integration Id | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`


REST Endpoint that returns all job codes by company And integration

*Success Response:*



Successful operation


Schema: `ResponseEnvelopeListJobConfigListDTO`









---



---


## Configuration


#### getBuildConfig
Get latest build config

```golang

data, err := Configuration.GetBuildConfig(CompanyID, ApplicationID, PlatformType);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 


| PlatformType | string | Current platform name | 



Get latest build config

*Success Response:*



Success


Schema: `MobileAppConfiguration`









---


#### updateBuildConfig
Update build config for next build

```golang

data, err := Configuration.UpdateBuildConfig(CompanyID, ApplicationID, PlatformType, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 


| PlatformType | string | Current platform name | 


| body |  MobileAppConfigRequest | "Request body" 

Update build config for next build

*Success Response:*



Success


Schema: `MobileAppConfiguration`









---


#### getPreviousVersions
Get previous build versions

```golang

data, err := Configuration.GetPreviousVersions(CompanyID, ApplicationID, PlatformType);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 


| PlatformType | string | Current platform name | 



Get previous build versions

*Success Response:*



Success


Schema: `BuildVersionHistory`









---


#### getAppFeatures
Get features of application

```golang

data, err := Configuration.GetAppFeatures(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 



Get features of application

*Success Response:*



Success


Schema: `AppFeatureResponse`









---


#### updateAppFeatures
Update features of application

```golang

data, err := Configuration.UpdateAppFeatures(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 


| body |  AppFeatureRequest | "Request body" 

Update features of application

*Success Response:*



Success


Schema: `AppFeature`









---


#### getAppBasicDetails
Get basic application details

```golang

data, err := Configuration.GetAppBasicDetails(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 



Get basic application details like name

*Success Response:*



Success


Schema: `ApplicationDetail`









---


#### updateAppBasicDetails
Add or update application's basic details

```golang

data, err := Configuration.UpdateAppBasicDetails(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 


| body |  ApplicationDetail | "Request body" 

Add or update application's basic details

*Success Response:*



Success


Schema: `ApplicationDetail`









---


#### getAppContactInfo
Get application information

```golang

data, err := Configuration.GetAppContactInfo(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 



Get Application Current Information. This includes information about social links, address and contact information of company/seller/brand of the application.

*Success Response:*



Success


Schema: `ApplicationInformation`









---


#### updateAppContactInfo
Get application information

```golang

data, err := Configuration.UpdateAppContactInfo(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 


| body |  ApplicationInformation | "Request body" 

Save Application Current Information. This includes information about social links, address and contact information of an application.

*Success Response:*



Success


Schema: `ApplicationInformation`









---


#### getAppApiTokens
Get social tokens

```golang

data, err := Configuration.GetAppApiTokens(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 



Get social tokens.

*Success Response:*



Success


Schema: `TokenResponse`









---


#### updateAppApiTokens
Add social tokens

```golang

data, err := Configuration.UpdateAppApiTokens(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 


| body |  TokenResponse | "Request body" 

Add social tokens.

*Success Response:*



Success


Schema: `TokenResponse`









---


#### getAppCompanies
Application inventory enabled companies

```golang

data, err := Configuration.GetAppCompanies(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`


Application inventory enabled companies.

*Success Response:*



Success


Schema: `CompaniesResponse`









---


#### getAppStores
Application inventory enabled stores

```golang

data, err := Configuration.GetAppStores(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`


Application inventory enabled stores.

*Success Response:*



Success


Schema: `StoresResponse`









---


#### getInventoryConfig
Get application configuration

```golang

data, err := Configuration.GetInventoryConfig(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 



Get application configuration for various features and data

*Success Response:*



Success


Schema: `ApplicationInventory`









---


#### updateInventoryConfig
Update application configuration

```golang

data, err := Configuration.UpdateInventoryConfig(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 


| body |  ApplicationInventory | "Request body" 

Update application configuration for various features and data

*Success Response:*



Success


Schema: `ApplicationInventory`









---


#### partiallyUpdateInventoryConfig
Partially update application configuration

```golang

data, err := Configuration.PartiallyUpdateInventoryConfig(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 


| body |  AppInventoryPartialUpdate | "Request body" 

Partially update application configuration for various features and data

*Success Response:*



Success


Schema: `ApplicationInventory`









---


#### getAppCurrencyConfig
Get application enabled currency list

```golang

data, err := Configuration.GetAppCurrencyConfig(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 



Get application enabled currency list

*Success Response:*



Success


Schema: `AppSupportedCurrency`









---


#### updateAppCurrencyConfig
Add initial application supported currency

```golang

data, err := Configuration.UpdateAppCurrencyConfig(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 


| body |  AppSupportedCurrency | "Request body" 

Add initial application supported currency for various features and data. Default INR will be enabled.

*Success Response:*



Success


Schema: `AppSupportedCurrency`









---


#### getOrderingStoresByFilter
Get ordering store by filter

```golang

data, err := Configuration.GetOrderingStoresByFilter(CompanyID, ApplicationID, xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`

| body |  FilterOrderingStoreRequest | "Request body" 

Get ordering store by filter

*Success Response:*



Success


Schema: `OrderingStores`









---


#### updateOrderingStoreConfig
Add/Update ordering store config

```golang

data, err := Configuration.UpdateOrderingStoreConfig(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 


| body |  OrderingStoreConfig | "Request body" 

Add/Update ordering store config.

*Success Response:*



Success


Schema: `DeploymentMeta`









---


#### getDomains
Get attached domain list

```golang

data, err := Configuration.GetDomains(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 



Get attached domain list.

*Success Response:*



Success


Schema: `DomainsResponse`









---


#### addDomain
Add new domain to application

```golang

data, err := Configuration.AddDomain(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 


| body |  DomainAddRequest | "Request body" 

Add new domain to application.

*Success Response:*



Success


Schema: `Domain`









---


#### removeDomainById
Remove attached domain

```golang

data, err := Configuration.RemoveDomainById(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 


| ID | string | Domain _id | 



Remove attached domain.

*Success Response:*



Success


Schema: `SuccessMessageResponse`









---


#### changeDomainType
Change domain type

```golang

data, err := Configuration.ChangeDomainType(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 


| body |  UpdateDomainTypeRequest | "Request body" 

Change a domain to Primary or Shortlink domain

*Success Response:*



Success


Schema: `DomainsResponse`









---


#### getDomainStatus
Get domain connected status.

```golang

data, err := Configuration.GetDomainStatus(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 


| body |  DomainStatusRequest | "Request body" 

Get domain connected status. Check if domain is live and mapped to appropriate IP to fynd servers.

*Success Response:*



Success


Schema: `DomainStatusResponse`









---


#### createApplication
Create application

```golang

data, err := Configuration.CreateApplication(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| body |  CreateApplicationRequest | "Request body" 

Create new application

*Success Response:*



Success


Schema: `CreateAppResponse`









---


#### getApplications
Get list of application under company

```golang

data, err := Configuration.GetApplications(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 







| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `Q`


Get list of application under company

*Success Response:*



Success


Schema: `ApplicationsResponse`









---


#### getApplicationById
Get application data from id

```golang

data, err := Configuration.GetApplicationById(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 



Get application data from id

*Success Response:*



Success


Schema: `Application`









---


#### getCurrencies
Get all currencies

```golang

data, err := Configuration.GetCurrencies(CompanyID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 



Get all currencies

*Success Response:*



Currencies Success response


Schema: `CurrenciesResponse`









---


#### getDomainAvailibility
Check domain availibility before linking to application

```golang

data, err := Configuration.GetDomainAvailibility(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| body |  DomainSuggestionsRequest | "Request body" 

Check domain availibility before linking to application. Also sends domain suggestions with similar to queried domain. \ Custom domain search is currently powered by GoDaddy provider.

*Success Response:*



Success


Schema: `DomainSuggestionsResponse`


*Examples:*


Suggestions for fynd domains
```json
{
  "value": {
    "domains": [
      {
        "name": "test.hostx1.de",
        "is_available": false
      },
      {
        "name": "testhive.hostx1.de",
        "is_available": true
      }
    ]
  }
}
```

Suggestions for custom domains
```json
{
  "value": {
    "domains": [
      {
        "name": "test25.in",
        "unsupported": false,
        "is_available": false
      },
      {
        "name": "try25.in",
        "unsupported": false,
        "is_available": true,
        "price": 14.99,
        "currency": "USD"
      }
    ]
  }
}
```









---


#### getIntegrationById
Get integration data

```golang

data, err := Configuration.GetIntegrationById(CompanyID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ID | float64 | Integration id | 



Get integration data

*Success Response:*



Success


Schema: `Integration`









---


#### getAvailableOptIns
Get all available integration opt-ins

```golang

data, err := Configuration.GetAvailableOptIns(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`


Get all available integration opt-ins

*Success Response:*



Success


Schema: `GetIntegrationsOptInsResponse`









---


#### getSelectedOptIns
Get company/store level integration opt-ins

```golang

data, err := Configuration.GetSelectedOptIns(CompanyID, Level, UID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| Level | string | Integration level | 


| UID | float64 | Integration level uid | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`


Get company/store level integration opt-ins

*Success Response:*



Success


Schema: `GetIntegrationsOptInsResponse`









---


#### getIntegrationLevelConfig
Get integration level config

```golang

data, err := Configuration.GetIntegrationLevelConfig(CompanyID, ID, Level, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ID | string | Integration id | 


| Level | string | Integration level | 





| xQuery | struct | Includes properties such as `Opted`, `CheckPermission`


Get integration/integration-opt-in level config

*Success Response:*



Success


Schema: `IntegrationConfigResponse`









---


#### getIntegrationByLevelId
Get level data for integration

```golang

data, err := Configuration.GetIntegrationByLevelId(CompanyID, ID, Level, UID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ID | string | Integration id | 


| Level | string | Integration level | 


| UID | float64 | Integration level uid | 



Get level data for integration

*Success Response:*



Success


Schema: `IntegrationLevel`









---


#### getLevelActiveIntegrations
Check store has active integration

```golang

data, err := Configuration.GetLevelActiveIntegrations(CompanyID, ID, Level, UID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ID | string | Integration id | 


| Level | string | Integration level | 


| UID | float64 | Integration level uid | 



API checks if a store is already opted in any other integrations

*Success Response:*



Success


Schema: `OptedStoreIntegration`









---


#### getBrandsByCompany
Get brands by company

```golang

data, err := Configuration.GetBrandsByCompany(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 



| xQuery | struct | Includes properties such as `Q`


Get brands by company

*Success Response:*



Success


Schema: `BrandsByCompanyResponse`









---


#### getCompanyByBrands
Get company by brand uids

```golang

data, err := Configuration.GetCompanyByBrands(CompanyID, xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`

| body |  CompanyByBrandsRequest | "Request body" 

Get company by brand uids

*Success Response:*



Success


Schema: `CompanyByBrandsResponse`









---


#### getStoreByBrands
Get stores by brand uids

```golang

data, err := Configuration.GetStoreByBrands(CompanyID, xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`

| body |  StoreByBrandsRequest | "Request body" 

Get stores by brand uids

*Success Response:*



Success


Schema: `StoreByBrandsResponse`









---


#### getOtherSellerApplications
Get other seller applications

```golang

data, err := Configuration.GetOtherSellerApplications(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`


Get other seller applications who has opted current company as inventory

*Success Response:*



Success


Schema: `OtherSellerApplications`









---


#### getOtherSellerApplicationById
Get other seller applications

```golang

data, err := Configuration.GetOtherSellerApplicationById(CompanyID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ID | string | Application Id | 



Get other seller application

*Success Response:*



Success


Schema: `OptedApplicationResponse`









---


#### optOutFromApplication
Opt out company or store from other seller application

```golang

data, err := Configuration.OptOutFromApplication(CompanyID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ID | string | Application Id | 


| body |  OptOutInventory | "Request body" 

Opt out company or store from other seller application

*Success Response:*



Success


Schema: `SuccessMessageResponse`









---



---


## Cart


#### getCoupons
Get with single coupon details or coupon list

```golang

data, err := Cart.GetCoupons(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current Application _id | 

















| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `IsArchived`, `Title`, `IsPublic`, `IsDisplay`, `TypeSlug`, `Code`


Get coupon list with pagination

*Success Response:*



Coupon List for sent page_size and page_no


Schema: `CouponsResponse`


*Examples:*


Coupon list for sent filter and page size
```json
{
  "value": {
    "items": [
      {
        "_id": "5e1d9bec6d6b7e000146c840",
        "display_meta": {
          "title": "percent50 title"
        },
        "_schedule": {
          "next_schedule": [
            {
              "start": "2020-01-14T10:45:03.600000+00:00",
              "end": "2020-01-16T10:45:03+00:00"
            }
          ],
          "duration": null,
          "start": "2020-01-14T10:45:03.600000+00:00",
          "end": "2020-01-16T10:45:03+00:00",
          "cron": ""
        },
        "state": {
          "is_public": true,
          "is_display": true,
          "is_archived": false
        },
        "ownership": {
          "payable_category": "seller",
          "payable_by": ""
        },
        "code": "percent50",
        "rule_definition": {
          "type": "percentage",
          "scope": [
            "category_id"
          ],
          "applicable_on": "quantity"
        }
      }
    ],
    "page": {
      "has_next": true,
      "size": 10,
      "current": 1,
      "item_total": 30
    }
  }
}
```









---


#### createCoupon
Create new coupon

```golang

data, err := Cart.CreateCoupon(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current Application _id | 


| body |  CouponAdd | "Request body" 

Create new coupon

*Success Response:*



Coupon Created successfully


Schema: `SuccessMessage`









---


#### getCouponById
Get with single coupon details or coupon list

```golang

data, err := Cart.GetCouponById(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current Application _id | 


| ID | string |  | 



Get single coupon details with `id` in path param

*Success Response:*



Coupon object for sent `id`


Schema: `CouponUpdate`









---


#### updateCoupon
Update existing coupon configuration

```golang

data, err := Cart.UpdateCoupon(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current Application _id | 


| ID | string |  | 


| body |  CouponUpdate | "Request body" 

Update coupon with id sent in `id`

*Success Response:*



Coupon updated successfully


Schema: `SuccessMessage`









---


#### updateCouponPartially
Update coupon archive state and schedule

```golang

data, err := Cart.UpdateCouponPartially(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current Application _id | 


| ID | string |  | 


| body |  CouponPartialUpdate | "Request body" 

Update archive/unarchive and change schedule for coupon

*Success Response:*



Coupon updated successfully


Schema: `SuccessMessage`


*Examples:*


Archive or Unarchive coupon
```json
{
  "value": {
    "success": true,
    "message": "Coupon Updated"
  }
}
```

Coupon schedule updated successfully
```json
{
  "value": {
    "success": true,
    "message": "Coupon schedule updated"
  }
}
```









---


#### fetchAndvalidateCartItems
Fetch Cart Details

```golang

data, err := Cart.FetchAndvalidateCartItems(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current Application _id | 


| body |  OpenapiCartDetailsRequest | "Request body" 

Get all the details of cart for a list of provided `cart_items`

*Success Response:*



Cart details with breakup


Schema: `OpenapiCartDetailsResponse`









---


#### checkCartServiceability
Check Pincode Serviceability

```golang

data, err := Cart.CheckCartServiceability(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current Application _id | 


| body |  OpenApiCartServiceabilityRequest | "Request body" 

Check Pincode serviceability for cart items provided in `cart_items` and address pincode in `shipping_address`

*Success Response:*



Cart details with pincode validity information at item level


Schema: `OpenApiCartServiceabilityResponse`


*Examples:*


Valid pincode
```json
{
  "value": {
    "items": [
      {
        "quantity": 1,
        "message": "",
        "coupon_message": "",
        "product": {
          "type": "product",
          "uid": 803140,
          "name": "Green Solid T-Shirt",
          "slug": "celio-green-solid-t-shirt-803140-dd9e2c",
          "brand": {
            "uid": 44,
            "name": "celio"
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
              "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/44_NEMIEL@GREENBRITISH/1_1548161273344.jpg",
              "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/44_NEMIEL@GREENBRITISH/1_1548161273344.jpg"
            }
          ],
          "action": {
            "type": "product",
            "url": "https://api.addsale.com/platform/content/v1/products/celio-green-solid-t-shirt-803140-dd9e2c/",
            "query": {
              "product_slug": [
                "celio-green-solid-t-shirt-803140-dd9e2c"
              ]
            }
          }
        },
        "article": {
          "type": "article",
          "uid": "25_44_A7050_NEMIEL@GREENBRITISH_S",
          "size": "S",
          "seller": {
            "uid": 25,
            "name": "CELIO FUTURE FASHION PRIVATE LIMITED"
          },
          "store": {
            "uid": 1486,
            "name": "Forum Mall"
          },
          "quantity": 1,
          "price": {
            "base": {
              "marked": 1299,
              "effective": 649.5,
              "currency_code": "INR"
            },
            "converted": {
              "marked": 1299,
              "effective": 649.5,
              "currency_code": "INR"
            }
          }
        },
        "key": "803140_S",
        "discount": "50% OFF",
        "price": {
          "base": {
            "add_on": 0,
            "marked": 1299,
            "effective": 649.5,
            "selling": 649.5,
            "currency_code": "INR"
          },
          "converted": {
            "add_on": 0,
            "marked": 1299,
            "effective": 649.5,
            "selling": 649.5,
            "currency_code": "INR"
          }
        },
        "availability": {
          "sizes": [
            "L",
            "XL",
            "M",
            "S"
          ],
          "other_store_quantity": 0,
          "out_of_stock": false,
          "deliverable": true,
          "is_valid": true,
          "delivery_promise": {
            "timestamp": {
              "min": 1605306343,
              "max": 1605468343
            },
            "formatted": {
              "min": "Sat, 14 Nov",
              "max": "Mon, 16 Nov"
            }
          },
          "available_sizes": [
            {
              "is_available": true,
              "display": "L",
              "value": "L"
            },
            {
              "is_available": true,
              "display": "XXL",
              "value": "XXL"
            },
            {
              "is_available": true,
              "display": "XL",
              "value": "XL"
            },
            {
              "is_available": true,
              "display": "M",
              "value": "M"
            },
            {
              "is_available": true,
              "display": "S",
              "value": "S"
            },
            {
              "is_available": false,
              "display": "30",
              "value": "30"
            }
          ]
        },
        "bulk_offer": {}
      },
      {
        "quantity": 1,
        "message": "Out of stock. Please remove item",
        "coupon_message": "",
        "product": {
          "type": "product",
          "uid": 803140,
          "name": "Green Solid T-Shirt",
          "slug": "celio-green-solid-t-shirt-803140-dd9e2c",
          "brand": {
            "uid": 44,
            "name": "celio"
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
              "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/44_NEMIEL@GREENBRITISH/1_1548161273344.jpg",
              "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/44_NEMIEL@GREENBRITISH/1_1548161273344.jpg"
            }
          ],
          "action": {
            "type": "product",
            "url": "https://api.addsale.com/platform/content/v1/products/celio-green-solid-t-shirt-803140-dd9e2c/",
            "query": {
              "product_slug": [
                "celio-green-solid-t-shirt-803140-dd9e2c"
              ]
            }
          }
        },
        "article": {},
        "key": "803140_S",
        "discount": "",
        "price": {
          "base": {
            "add_on": 0,
            "marked": 1299,
            "effective": 1299,
            "selling": 1299,
            "currency_code": "INR"
          },
          "converted": {
            "add_on": 0,
            "marked": 1299,
            "effective": 1299,
            "selling": 1299,
            "currency_code": "INR"
          }
        },
        "availability": {
          "sizes": [
            "L",
            "XXL",
            "XL",
            "M",
            "S"
          ],
          "other_store_quantity": 0,
          "out_of_stock": true,
          "deliverable": false,
          "is_valid": false,
          "delivery_promise": {
            "timestamp": {
              "min": 1605306343,
              "max": 1605468343
            },
            "formatted": {
              "min": "Sat, 14 Nov",
              "max": "Mon, 16 Nov"
            }
          },
          "available_sizes": [
            {
              "is_available": true,
              "display": "L",
              "value": "L"
            },
            {
              "is_available": true,
              "display": "XXL",
              "value": "XXL"
            },
            {
              "is_available": true,
              "display": "XL",
              "value": "XL"
            },
            {
              "is_available": true,
              "display": "M",
              "value": "M"
            },
            {
              "is_available": true,
              "display": "S",
              "value": "S"
            },
            {
              "is_available": false,
              "display": "30",
              "value": "30"
            }
          ]
        },
        "bulk_offer": {}
      }
    ],
    "delivery_promise": {
      "timestamp": {
        "min": 1605306343,
        "max": 1605468343
      },
      "formatted": {
        "min": "Sat, 14 Nov",
        "max": "Mon, 16 Nov"
      }
    },
    "is_valid": true
  }
}
```

Invalid pincode
```json
{
  "value": {
    "message": "All of the items in your cart are not deliverable to 800108",
    "is_valid": false,
    "items": [
      {
        "discount": "15% OFF",
        "price": {
          "base": {
            "add_on": 0,
            "marked": 2195,
            "effective": 1866,
            "selling": 1866,
            "currency_code": "INR"
          },
          "converted": {
            "add_on": 0,
            "marked": 2195,
            "effective": 1866,
            "selling": 1866,
            "currency_code": "INR"
          }
        },
        "product": {
          "type": "product",
          "uid": 876245,
          "name": "Brown Sandals",
          "slug": "red-chief-brown-sandals-876245-c92507",
          "brand": {
            "uid": 433,
            "name": ""
          },
          "categories": [
            {
              "uid": 176,
              "name": ""
            }
          ],
          "images": [
            {
              "aspect_ratio": "16:25",
              "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/433_RC330004/1_1564147181287.jpg",
              "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/433_RC330004/1_1564147181287.jpg"
            }
          ],
          "action": {
            "type": "product",
            "url": "https://api.addsale.com/platform/content/v1/products/red-chief-brown-sandals-876245-c92507/",
            "query": {
              "product_slug": [
                "red-chief-brown-sandals-876245-c92507"
              ]
            }
          },
          "item_code": "RC330004"
        },
        "bulk_offer": {},
        "key": "876245_6",
        "message": "We are not delivering to 800108",
        "delivery_promise": null,
        "coupon_message": "",
        "availability": {
          "sizes": [
            "7",
            "6",
            "10",
            "8"
          ],
          "other_store_quantity": 21,
          "out_of_stock": false,
          "deliverable": false,
          "is_valid": true,
          "available_sizes": [
            {
              "is_available": false,
              "display": "9",
              "value": "9"
            },
            {
              "is_available": true,
              "display": "10",
              "value": "10"
            },
            {
              "is_available": true,
              "display": "6",
              "value": "6"
            },
            {
              "is_available": true,
              "display": "7",
              "value": "7"
            },
            {
              "is_available": true,
              "display": "8",
              "value": "8"
            }
          ]
        },
        "quantity": 1,
        "article": {
          "type": "article",
          "uid": "304_433_LGPL30402_RC330004_6",
          "size": "6",
          "seller": {
            "uid": 304,
            "name": "LEAYAN GLOBAL PVT. LTD."
          },
          "store": {
            "uid": 9767,
            "name": "Udyog Kunj, Kanpur"
          },
          "quantity": 3,
          "price": {
            "base": {
              "marked": 2195,
              "effective": 1866,
              "currency_code": "INR"
            },
            "converted": {
              "marked": 2195,
              "effective": 1866,
              "currency_code": "INR"
            }
          }
        }
      },
      {
        "discount": "15% OFF",
        "price": {
          "base": {
            "add_on": 0,
            "marked": 2195,
            "effective": 1866,
            "selling": 1866,
            "currency_code": "INR"
          },
          "converted": {
            "add_on": 0,
            "marked": 2195,
            "effective": 1866,
            "selling": 1866,
            "currency_code": "INR"
          }
        },
        "product": {
          "type": "product",
          "uid": 876245,
          "name": "Brown Sandals",
          "slug": "red-chief-brown-sandals-876245-c92507",
          "brand": {
            "uid": 433,
            "name": ""
          },
          "categories": [
            {
              "uid": 176,
              "name": ""
            }
          ],
          "images": [
            {
              "aspect_ratio": "16:25",
              "url": "http://cdn4.gofynd.com/media/pictures/tagged_items/original/433_RC330004/1_1564147181287.jpg",
              "secure_url": "https://d2zv4gzhlr4ud6.cloudfront.net/media/pictures/tagged_items/original/433_RC330004/1_1564147181287.jpg"
            }
          ],
          "action": {
            "type": "product",
            "url": "https://api.addsale.com/platform/content/v1/products/red-chief-brown-sandals-876245-c92507/",
            "query": {
              "product_slug": [
                "red-chief-brown-sandals-876245-c92507"
              ]
            }
          },
          "item_code": "RC330004"
        },
        "bulk_offer": {},
        "key": "876245_6",
        "message": "We are not delivering to 800108",
        "coupon_message": "",
        "availability": {
          "sizes": [
            "7",
            "6",
            "10",
            "8"
          ],
          "other_store_quantity": 21,
          "out_of_stock": false,
          "deliverable": false,
          "is_valid": true,
          "available_sizes": [
            {
              "is_available": false,
              "display": "9",
              "value": "9"
            },
            {
              "is_available": true,
              "display": "10",
              "value": "10"
            },
            {
              "is_available": true,
              "display": "6",
              "value": "6"
            },
            {
              "is_available": true,
              "display": "7",
              "value": "7"
            },
            {
              "is_available": true,
              "display": "8",
              "value": "8"
            }
          ]
        },
        "quantity": 1,
        "article": {
          "type": "article",
          "uid": "304_433_LGPL30402_RC330004_6",
          "size": "6",
          "seller": {
            "uid": 304,
            "name": "LEAYAN GLOBAL PVT. LTD."
          },
          "store": {
            "uid": 9767,
            "name": "Udyog Kunj, Kanpur"
          },
          "quantity": 3,
          "price": {
            "base": {
              "marked": 2195,
              "effective": 1866,
              "currency_code": "INR"
            },
            "converted": {
              "marked": 2195,
              "effective": 1866,
              "currency_code": "INR"
            }
          }
        }
      }
    ]
  }
}
```









---


#### checkoutCart
Create Fynd order with cart details

```golang

data, err := Cart.CheckoutCart(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current Application _id | 


| body |  OpenApiPlatformCheckoutReq | "Request body" 

Generate Fynd order for cart details send with provided `cart_items`

*Success Response:*



Checkout cart and create Fynd order id


Schema: `OpenApiCheckoutResponse`









---



---


## Rewards


#### getGiveaways
List of giveaways of the current application.

```golang

data, err := Rewards.GetGiveaways(CompanyID, ApplicationID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| ApplicationID | string | application id | 





| xQuery | struct | Includes properties such as `PageID`, `PageSize`


List of giveaways of the current application.

*Success Response:*



ok


Schema: `GiveawayResponse`









---


#### createGiveaway
Adds a new giveaway.

```golang

data, err := Rewards.CreateGiveaway(CompanyID, ApplicationID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| ApplicationID | string | application id | 


| body |  Giveaway | "Request body" 

Adds a new giveaway.

*Success Response:*



ok


Schema: `Giveaway`









---


#### getGiveawayByID
Get giveaway by ID.

```golang

data, err := Rewards.GetGiveawayByID(CompanyID, ApplicationID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| ApplicationID | string | application id | 


| ID | string | Giveaway ID | 



Get giveaway by ID.

*Success Response:*



ok


Schema: `Giveaway`









---


#### updateGiveaway
Updates the giveaway by it's ID.

```golang

data, err := Rewards.UpdateGiveaway(CompanyID, ApplicationID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| ApplicationID | string | application id | 


| ID | string | Giveaway ID | 


| body |  Giveaway | "Request body" 

Updates the giveaway by it's ID.

*Success Response:*



ok


Schema: `Giveaway`









---


#### getOffers
List of offer of the current application.

```golang

data, err := Rewards.GetOffers(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| ApplicationID | string | application id | 



List of offer of the current application.

*Success Response:*



ok


Schema: `Array<Offer>`









---


#### getOfferByName
Get offer by name.

```golang

data, err := Rewards.GetOfferByName(CompanyID, ApplicationID, Cookie, Name);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| ApplicationID | string | application id | 


| Cookie | string | User's session cookie. This cookie is set in browser cookie when logged-in to fynd's authentication system i.e. `Grimlock` or by using grimlock-backend SDK for backend implementation. | 


| Name | string | Offer name | 



Get offer by name.

*Success Response:*



ok


Schema: `Offer`









---


#### updateOfferByName
Updates the offer by name.

```golang

data, err := Rewards.UpdateOfferByName(CompanyID, ApplicationID, Name, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| ApplicationID | string | application id | 


| Name | string | Offer name | 


| body |  Offer | "Request body" 

Updates the offer by name.

*Success Response:*



ok


Schema: `Offer`









---


#### getUserAvailablePoints
User's reward details.

```golang

data, err := Rewards.GetUserAvailablePoints(CompanyID, ApplicationID, UserID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| ApplicationID | string | application id | 


| UserID | string | user id | 



User's reward details.

*Success Response:*



ok


Schema: `UserRes`









---


#### updateUserStatus
Update User status

```golang

data, err := Rewards.UpdateUserStatus(CompanyID, ApplicationID, UserID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| ApplicationID | string | application id | 


| UserID | string | user id | 


| body |  AppUser | "Request body" 

Update user status, active/archive

*Success Response:*



Success


Schema: `AppUser`









---


#### getUserPointsHistory
Get list of points transactions.

```golang

data, err := Rewards.GetUserPointsHistory(CompanyID, ApplicationID, UserID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | company id | 


| ApplicationID | string | application id | 


| UserID | string | user id | 







| xQuery | struct | Includes properties such as `PageID`, `PageLimit`, `PageSize`


Get list of points transactions.
The list of points history is paginated.

*Success Response:*



ok


Schema: `HistoryRes`









---



---


## Analytics


#### getStatiscticsGroups
Get statistics groups

```golang

data, err := Analytics.GetStatiscticsGroups(CompanyID, ApplicationID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ApplicationID | string | Application Id | 



Get statistics groups

*Success Response:*



Success


Schema: `StatsGroups`









---


#### getStatiscticsGroupComponents
Get statistics group components

```golang

data, err := Analytics.GetStatiscticsGroupComponents(CompanyID, ApplicationID, GroupName);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ApplicationID | string | Application Id | 


| GroupName | string | Group name | 



Get statistics group components

*Success Response:*



Success


Schema: `StatsGroupComponents`









---


#### getComponentStatsCSV
Get component statistics csv

```golang

data, err := Analytics.GetComponentStatsCSV(CompanyID, ApplicationID, ComponentName);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ApplicationID | string | Application Id | 


| ComponentName | string | Component name | 



Get component statistics csv

*Success Response:*



Success


Schema: `string`









---


#### getComponentStatsPDF
Get component statistics pdf

```golang

data, err := Analytics.GetComponentStatsPDF(CompanyID, ApplicationID, ComponentName);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ApplicationID | string | Application Id | 


| ComponentName | string | Component name | 



Get component statistics pdf

*Success Response:*



Success


Schema: `string`









---


#### getComponentStats
Get component statistics

```golang

data, err := Analytics.GetComponentStats(CompanyID, ApplicationID, ComponentName);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ApplicationID | string | Application Id | 


| ComponentName | string | Component name | 



Get component statistics

*Success Response:*



Success


Schema: `StatsRes`









---


#### getAbandonCartList
Get abandon carts list

```golang

data, err := Analytics.GetAbandonCartList(CompanyID, ApplicationID, FromDate, ToDate, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ApplicationID | string | Application Id | 


| FromDate | string | From date | 


| ToDate | string | To date | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`


Get abandon carts list

*Success Response:*



Success


Schema: `AbandonCartsList`









---


#### getAbandonCartsCSV
Get abandon carts csv

```golang

data, err := Analytics.GetAbandonCartsCSV(CompanyID, ApplicationID, FromDate, ToDate);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ApplicationID | string | Application Id | 


| FromDate | string | From date | 


| ToDate | string | To date | 



Get abandon carts csv

*Success Response:*



Success


Schema: `string`









---


#### getAbandonCartDetail
Get abandon carts details

```golang

data, err := Analytics.GetAbandonCartDetail(CompanyID, ApplicationID, CartID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ApplicationID | string | Application Id | 


| CartID | string | Cart Id | 



Get abandon cart details

*Success Response:*



Success


Schema: `AbandonCartDetail`









---


#### createExportJob
Create data export job in required format

```golang

data, err := Analytics.CreateExportJob(CompanyID, ExportType, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ExportType | string | Export type / format | 


| body |  ExportJobReq | "Request body" 

Create data export job in required format

*Success Response:*



Success


Schema: `ExportJobRes`









---


#### getExportJobStatus
Get data export job status

```golang

data, err := Analytics.GetExportJobStatus(CompanyID, ExportType, JobID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| ExportType | string | Export type / format | 


| JobID | string | Export job id | 



Get data export job status

*Success Response:*



Success


Schema: `ExportJobStatusRes`









---


#### getLogsList
Get logs list

```golang

data, err := Analytics.GetLogsList(CompanyID, LogType, xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 


| LogType | string | Log type | 





| xQuery | struct | Includes properties such as `PageNo`, `PageSize`

| body |  GetLogsListReq | "Request body" 

Get logs list

*Success Response:*



Success


Schema: `GetLogsListRes`









---


#### searchLogs
Search logs

```golang

data, err := Analytics.SearchLogs(CompanyID, LogType, xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Company Id | 






| LogType | string | Log type | 

| xQuery | struct | Includes properties such as `PageNo`, `PageSize`

| body |  SearchLogReq | "Request body" 

Search logs

*Success Response:*



Success


Schema: `SearchLogRes`









---



---


## Discount


#### getDiscounts
Fetch discount list.

```golang

data, err := Discount.GetDiscounts(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | company_id | 



















| xQuery | struct | Includes properties such as `View`, `Q`, `PageNo`, `PageSize`, `Archived`, `Month`, `Year`, `Type`, `AppIds`


Fetch discount list.

*Success Response:*



Success


Schema: `ListOrCalender`









---


#### createDiscount
Create Discount.

```golang

data, err := Discount.CreateDiscount(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | company_id | 


| body |  CreateUpdateDiscount | "Request body" 

Create Discount.

*Success Response:*



Success


Schema: `DiscountJob`









---


#### getDiscount
Fetch discount.

```golang

data, err := Discount.GetDiscount(CompanyID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | company_id | 


| ID | string | unique id. | 



Fetch discount.

*Success Response:*



Success


Schema: `DiscountJob`









---


#### updateDiscount
Create Discount.

```golang

data, err := Discount.UpdateDiscount(CompanyID, ID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | company_id | 


| ID | string | id | 


| body |  CreateUpdateDiscount | "Request body" 

Create Discount.

*Success Response:*



Success


Schema: `DiscountJob`









---


#### validateDiscountFile
Validate File.

```golang

data, err := Discount.ValidateDiscountFile(CompanyID, xQuery, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | company_id | 



| xQuery | struct | Includes properties such as `Discount`

| body |  DiscountJob | "Request body" 

Validate File.

*Success Response:*



Success


Schema: `FileJobResponse`









---


#### downloadDiscountFile
Validate File.

```golang

data, err := Discount.DownloadDiscountFile(CompanyID, Type, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | company_id | 


| Type | string | type | 


| body |  DownloadFileJob | "Request body" 

Validate File.

*Success Response:*



Success


Schema: `FileJobResponse`









---


#### getValidationJob
Validate File Job.

```golang

data, err := Discount.GetValidationJob(CompanyID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | company_id | 


| ID | string | id | 



Validate File Job.

*Success Response:*



Success


Schema: `FileJobResponse`









---


#### cancelValidationJob
Cancel Validation Job.

```golang

data, err := Discount.CancelValidationJob(CompanyID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | company_id | 


| ID | string | id | 



Cancel Validation Job.

*Success Response:*



Success


Schema: `CancelJobResponse`









---


#### getDownloadJob
Download File Job.

```golang

data, err := Discount.GetDownloadJob(CompanyID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | company_id | 


| ID | string | id | 



Download File Job.

*Success Response:*



Success


Schema: `FileJobResponse`









---


#### cancelDownloadJob
Cancel Download Job.

```golang

data, err := Discount.CancelDownloadJob(CompanyID, ID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | company_id | 


| ID | string | id | 



Cancel Download Job.

*Success Response:*



Success


Schema: `CancelJobResponse`









---



---


## Partner


#### addProxyPath
Add proxy path for external url

```golang

data, err := Partner.AddProxyPath(CompanyID, ApplicationID, ExtensionID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 


| ExtensionID | string | Extension id | 


| body |  AddProxyReq | "Request body" 

Add proxy path for external url

*Success Response:*



Success


Schema: `AddProxyResponse`









---


#### removeProxyPath
Remove proxy path for external url

```golang

data, err := Partner.RemoveProxyPath(CompanyID, ApplicationID, ExtensionID, AttachedPath);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | string | Current company id | 


| ApplicationID | string | Current application id | 


| ExtensionID | string | Extension id | 


| AttachedPath | string | Attachaed path slug | 



Remove proxy path for external url

*Success Response:*



Success


Schema: `RemoveProxyResponse`









---



---


## Webhook


#### getSubscribersByCompany
Get Subscribers By Company ID

```golang

data, err := Webhook.GetSubscribersByCompany(CompanyID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |





| CompanyID | float64 | Company ID of the application | 



| xQuery | struct | Includes properties such as `PageNo`, `PageSize`, `ExtensionID`


Get Subscribers By CompanyId

*Success Response:*



Success


Schema: `SubscriberResponse`









---


#### registerSubscriberToEvent
Register Subscriber

```golang

data, err := Webhook.RegisterSubscriberToEvent(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company Id of the application | 


| body |  SubscriberConfig | "Request body" 

Register Subscriber

*Success Response:*



Success


Schema: `SubscriberConfig`









---


#### updateSubscriberConfig
Update Subscriber

```golang

data, err := Webhook.UpdateSubscriberConfig(CompanyID, body);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company ID of the application | 


| body |  SubscriberConfig | "Request body" 

Update Subscriber

*Success Response:*



Success


Schema: `SubscriberConfig`









---


#### getSubscribersByExtensionId
Get Subscribers By Extension ID

```golang

data, err := Webhook.GetSubscribersByExtensionId(CompanyID, ExtensionID, xQuery);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |





| CompanyID | float64 | Company ID of the application | 


| ExtensionID | string | Extension ID | 

| xQuery | struct | Includes properties such as `PageNo`, `PageSize`


Get Subscribers By ExtensionID

*Success Response:*



Success


Schema: `SubscriberResponse`









---


#### getSubscriberById
Get Subscriber By Subscriber ID

```golang

data, err := Webhook.GetSubscriberById(CompanyID, SubscriberID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company ID of the application | 


| SubscriberID | float64 | Subscriber ID | 



Get Subscriber By Subscriber ID

*Success Response:*



Success


Schema: `SubscriberResponse`









---


#### fetchAllEventConfigurations
Get All Webhook Events

```golang

data, err := Webhook.FetchAllEventConfigurations(CompanyID);
```

| Argument  |  Type  | Description |
| --------- | ----  | --- |

| CompanyID | float64 | Company ID of the application | 



Get All Webhook Events

*Success Response:*



Success


Schema: `EventConfigList`









---



---


---
---
