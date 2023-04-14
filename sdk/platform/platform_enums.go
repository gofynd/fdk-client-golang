package platform

import (
    "errors"
)






    //PriorityEnum used by Lead
    type PriorityEnum  string
    
    const (
            
            //LOW defines constant for the `low` 
            LOW PriorityEnum = "low"
            
            
            //MEDIUM defines constant for the `medium` 
            MEDIUM PriorityEnum = "medium"
            
            
            //HIGH defines constant for the `high` 
            HIGH PriorityEnum = "high"
            
            
            //URGENT defines constant for the `urgent` 
            URGENT PriorityEnum = "urgent"
            
            
       )

    //IsValid return error if enum is invalid
    func (pr PriorityEnum) IsValid() error {
      switch pr {
        case  LOW, MEDIUM, HIGH, URGENT : 
        return nil
      }
      return errors.New("Invalid PriorityEnum type")
    }



    //HistoryTypeEnum used by Lead
    type HistoryTypeEnum  string
    
    const (
            
            //RATING defines constant for the `rating` 
            RATING HistoryTypeEnum = "rating"
            
            
            //LOG defines constant for the `log` 
            LOG HistoryTypeEnum = "log"
            
            
            //COMMENT defines constant for the `comment` 
            COMMENT HistoryTypeEnum = "comment"
            
            
       )

    //IsValid return error if enum is invalid
    func (hi HistoryTypeEnum) IsValid() error {
      switch hi {
        case  RATING, LOG, COMMENT : 
        return nil
      }
      return errors.New("Invalid HistoryTypeEnum type")
    }



    //TicketAssetTypeEnum used by Lead
    type TicketAssetTypeEnum  string
    
    const (
            
            //IMAGE defines constant for the `image` 
            IMAGE TicketAssetTypeEnum = "image"
            
            
            //VIDEO defines constant for the `video` 
            VIDEO TicketAssetTypeEnum = "video"
            
            
            //FILE defines constant for the `file` 
            FILE TicketAssetTypeEnum = "file"
            
            
            //YOUTUBE defines constant for the `youtube` 
            YOUTUBE TicketAssetTypeEnum = "youtube"
            
            
            //PRODUCT defines constant for the `product` 
            PRODUCT TicketAssetTypeEnum = "product"
            
            
            //COLLECTION defines constant for the `collection` 
            COLLECTION TicketAssetTypeEnum = "collection"
            
            
            //BRAND defines constant for the `brand` 
            BRAND TicketAssetTypeEnum = "brand"
            
            
            //SHIPMENT defines constant for the `shipment` 
            SHIPMENT TicketAssetTypeEnum = "shipment"
            
            
            //ORDER defines constant for the `order` 
            ORDER TicketAssetTypeEnum = "order"
            
            
       )

    //IsValid return error if enum is invalid
    func (ti TicketAssetTypeEnum) IsValid() error {
      switch ti {
        case  IMAGE, VIDEO, FILE, YOUTUBE, PRODUCT, COLLECTION, BRAND, SHIPMENT, ORDER : 
        return nil
      }
      return errors.New("Invalid TicketAssetTypeEnum type")
    }



    //TicketSourceEnum used by Lead
    type TicketSourceEnum  string
    
    const (
            
            //PLATFORM_PANEL defines constant for the `platform_panel` 
            PLATFORM_PANEL TicketSourceEnum = "platform_panel"
            
            
            //SALES_CHANNEL defines constant for the `sales_channel` 
            SALES_CHANNEL TicketSourceEnum = "sales_channel"
            
            
       )

    //IsValid return error if enum is invalid
    func (ti TicketSourceEnum) IsValid() error {
      switch ti {
        case  PLATFORM_PANEL, SALES_CHANNEL : 
        return nil
      }
      return errors.New("Invalid TicketSourceEnum type")
    }



    //TicketIntegrationDetails used by Lead
    type TicketIntegrationDetails  string
    
    const (
            
            //DEFAULT defines constant for the `default` 
            DEFAULT TicketIntegrationDetails = "default"
            
            
            //FRESHDESK defines constant for the `freshdesk` 
            FRESHDESK TicketIntegrationDetails = "freshdesk"
            
            
            //KAPTURE defines constant for the `kapture` 
            KAPTURE TicketIntegrationDetails = "kapture"
            
            
       )

    //IsValid return error if enum is invalid
    func (ti TicketIntegrationDetails) IsValid() error {
      switch ti {
        case  DEFAULT, FRESHDESK, KAPTURE : 
        return nil
      }
      return errors.New("Invalid TicketIntegrationDetails type")
    }






    //GenerationEntityType used by Content
    type GenerationEntityType  string
    
    const (
            
            //TITLE defines constant for the `title` 
            TITLE GenerationEntityType = "title"
            
            
            //DESCRIPTION defines constant for the `description` 
            DESCRIPTION GenerationEntityType = "description"
            
            
       )

    //IsValid return error if enum is invalid
    func (ge GenerationEntityType) IsValid() error {
      switch ge {
        case  TITLE, DESCRIPTION : 
        return nil
      }
      return errors.New("Invalid GenerationEntityType type")
    }



    //PageType used by Content
    type PageType  string
    
    const (
            
            //ABOUT-US defines constant for the `about-us` 
            ABOUT-US PageType = "about-us"
            
            
            //ADDRESSES defines constant for the `addresses` 
            ADDRESSES PageType = "addresses"
            
            
            //BLOG defines constant for the `blog` 
            BLOG PageType = "blog"
            
            
            //BRANDS defines constant for the `brands` 
            BRANDS PageType = "brands"
            
            
            //CARDS defines constant for the `cards` 
            CARDS PageType = "cards"
            
            
            //CART defines constant for the `cart` 
            CART PageType = "cart"
            
            
            //CATEGORIES defines constant for the `categories` 
            CATEGORIES PageType = "categories"
            
            
            //BRAND defines constant for the `brand` 
            BRAND PageType = "brand"
            
            
            //CATEGORY defines constant for the `category` 
            CATEGORY PageType = "category"
            
            
            //COLLECTION defines constant for the `collection` 
            COLLECTION PageType = "collection"
            
            
            //COLLECTIONS defines constant for the `collections` 
            COLLECTIONS PageType = "collections"
            
            
            //CONTACT-US defines constant for the `contact-us` 
            CONTACT-US PageType = "contact-us"
            
            
            //EXTERNAL defines constant for the `external` 
            EXTERNAL PageType = "external"
            
            
            //FAQ defines constant for the `faq` 
            FAQ PageType = "faq"
            
            
            //FRESHCHAT defines constant for the `freshchat` 
            FRESHCHAT PageType = "freshchat"
            
            
            //HOME defines constant for the `home` 
            HOME PageType = "home"
            
            
            //NOTIFICATION-SETTINGS defines constant for the `notification-settings` 
            NOTIFICATION-SETTINGS PageType = "notification-settings"
            
            
            //ORDERS defines constant for the `orders` 
            ORDERS PageType = "orders"
            
            
            //PAGE defines constant for the `page` 
            PAGE PageType = "page"
            
            
            //POLICY defines constant for the `policy` 
            POLICY PageType = "policy"
            
            
            //PRODUCT defines constant for the `product` 
            PRODUCT PageType = "product"
            
            
            //PRODUCT-REVIEWS defines constant for the `product-reviews` 
            PRODUCT-REVIEWS PageType = "product-reviews"
            
            
            //ADD-PRODUCT-REVIEW defines constant for the `add-product-review` 
            ADD-PRODUCT-REVIEW PageType = "add-product-review"
            
            
            //PRODUCT-REQUEST defines constant for the `product-request` 
            PRODUCT-REQUEST PageType = "product-request"
            
            
            //PRODUCTS defines constant for the `products` 
            PRODUCTS PageType = "products"
            
            
            //PROFILE defines constant for the `profile` 
            PROFILE PageType = "profile"
            
            
            //PROFILE-ORDER-SHIPMENT defines constant for the `profile-order-shipment` 
            PROFILE-ORDER-SHIPMENT PageType = "profile-order-shipment"
            
            
            //PROFILE-BASIC defines constant for the `profile-basic` 
            PROFILE-BASIC PageType = "profile-basic"
            
            
            //PROFILE-COMPANY defines constant for the `profile-company` 
            PROFILE-COMPANY PageType = "profile-company"
            
            
            //PROFILE-EMAILS defines constant for the `profile-emails` 
            PROFILE-EMAILS PageType = "profile-emails"
            
            
            //PROFILE-PHONES defines constant for the `profile-phones` 
            PROFILE-PHONES PageType = "profile-phones"
            
            
            //RATE-US defines constant for the `rate-us` 
            RATE-US PageType = "rate-us"
            
            
            //REFER-EARN defines constant for the `refer-earn` 
            REFER-EARN PageType = "refer-earn"
            
            
            //SETTINGS defines constant for the `settings` 
            SETTINGS PageType = "settings"
            
            
            //SHARED-CART defines constant for the `shared-cart` 
            SHARED-CART PageType = "shared-cart"
            
            
            //TNC defines constant for the `tnc` 
            TNC PageType = "tnc"
            
            
            //TRACK-ORDER defines constant for the `track-order` 
            TRACK-ORDER PageType = "track-order"
            
            
            //WISHLIST defines constant for the `wishlist` 
            WISHLIST PageType = "wishlist"
            
            
            //SECTIONS defines constant for the `sections` 
            SECTIONS PageType = "sections"
            
            
            //FORM defines constant for the `form` 
            FORM PageType = "form"
            
            
            //CART-DELIVERY defines constant for the `cart-delivery` 
            CART-DELIVERY PageType = "cart-delivery"
            
            
            //CART-PAYMENT defines constant for the `cart-payment` 
            CART-PAYMENT PageType = "cart-payment"
            
            
            //CART-REVIEW defines constant for the `cart-review` 
            CART-REVIEW PageType = "cart-review"
            
            
            //LOGIN defines constant for the `login` 
            LOGIN PageType = "login"
            
            
            //REGISTER defines constant for the `register` 
            REGISTER PageType = "register"
            
            
            //SHIPPING-POLICY defines constant for the `shipping-policy` 
            SHIPPING-POLICY PageType = "shipping-policy"
            
            
            //RETURN-POLICY defines constant for the `return-policy` 
            RETURN-POLICY PageType = "return-policy"
            
            
       )

    //IsValid return error if enum is invalid
    func (pa PageType) IsValid() error {
      switch pa {
        case  ABOUT-US, ADDRESSES, BLOG, BRANDS, CARDS, CART, CATEGORIES, BRAND, CATEGORY, COLLECTION, COLLECTIONS, CONTACT-US, EXTERNAL, FAQ, FRESHCHAT, HOME, NOTIFICATION-SETTINGS, ORDERS, PAGE, POLICY, PRODUCT, PRODUCT-REVIEWS, ADD-PRODUCT-REVIEW, PRODUCT-REQUEST, PRODUCTS, PROFILE, PROFILE-ORDER-SHIPMENT, PROFILE-BASIC, PROFILE-COMPANY, PROFILE-EMAILS, PROFILE-PHONES, RATE-US, REFER-EARN, SETTINGS, SHARED-CART, TNC, TRACK-ORDER, WISHLIST, SECTIONS, FORM, CART-DELIVERY, CART-PAYMENT, CART-REVIEW, LOGIN, REGISTER, SHIPPING-POLICY, RETURN-POLICY : 
        return nil
      }
      return errors.New("Invalid PageType type")
    }








    //PageType used by Catalog
    type PageType  string
    
    const (
            
            //ABOUT-US defines constant for the `about-us` 
            ABOUT-US PageType = "about-us"
            
            
            //ADDRESSES defines constant for the `addresses` 
            ADDRESSES PageType = "addresses"
            
            
            //BLOG defines constant for the `blog` 
            BLOG PageType = "blog"
            
            
            //BRANDS defines constant for the `brands` 
            BRANDS PageType = "brands"
            
            
            //CARDS defines constant for the `cards` 
            CARDS PageType = "cards"
            
            
            //CART defines constant for the `cart` 
            CART PageType = "cart"
            
            
            //CATEGORIES defines constant for the `categories` 
            CATEGORIES PageType = "categories"
            
            
            //BRAND defines constant for the `brand` 
            BRAND PageType = "brand"
            
            
            //CATEGORY defines constant for the `category` 
            CATEGORY PageType = "category"
            
            
            //COLLECTION defines constant for the `collection` 
            COLLECTION PageType = "collection"
            
            
            //COLLECTIONS defines constant for the `collections` 
            COLLECTIONS PageType = "collections"
            
            
            //CONTACT-US defines constant for the `contact-us` 
            CONTACT-US PageType = "contact-us"
            
            
            //EXTERNAL defines constant for the `external` 
            EXTERNAL PageType = "external"
            
            
            //FAQ defines constant for the `faq` 
            FAQ PageType = "faq"
            
            
            //FRESHCHAT defines constant for the `freshchat` 
            FRESHCHAT PageType = "freshchat"
            
            
            //HOME defines constant for the `home` 
            HOME PageType = "home"
            
            
            //NOTIFICATION-SETTINGS defines constant for the `notification-settings` 
            NOTIFICATION-SETTINGS PageType = "notification-settings"
            
            
            //ORDERS defines constant for the `orders` 
            ORDERS PageType = "orders"
            
            
            //PAGE defines constant for the `page` 
            PAGE PageType = "page"
            
            
            //POLICY defines constant for the `policy` 
            POLICY PageType = "policy"
            
            
            //PRODUCT defines constant for the `product` 
            PRODUCT PageType = "product"
            
            
            //PRODUCT-REVIEWS defines constant for the `product-reviews` 
            PRODUCT-REVIEWS PageType = "product-reviews"
            
            
            //ADD-PRODUCT-REVIEW defines constant for the `add-product-review` 
            ADD-PRODUCT-REVIEW PageType = "add-product-review"
            
            
            //PRODUCT-REQUEST defines constant for the `product-request` 
            PRODUCT-REQUEST PageType = "product-request"
            
            
            //PRODUCTS defines constant for the `products` 
            PRODUCTS PageType = "products"
            
            
            //PROFILE defines constant for the `profile` 
            PROFILE PageType = "profile"
            
            
            //PROFILE-ORDER-SHIPMENT defines constant for the `profile-order-shipment` 
            PROFILE-ORDER-SHIPMENT PageType = "profile-order-shipment"
            
            
            //PROFILE-BASIC defines constant for the `profile-basic` 
            PROFILE-BASIC PageType = "profile-basic"
            
            
            //PROFILE-COMPANY defines constant for the `profile-company` 
            PROFILE-COMPANY PageType = "profile-company"
            
            
            //PROFILE-EMAILS defines constant for the `profile-emails` 
            PROFILE-EMAILS PageType = "profile-emails"
            
            
            //PROFILE-PHONES defines constant for the `profile-phones` 
            PROFILE-PHONES PageType = "profile-phones"
            
            
            //RATE-US defines constant for the `rate-us` 
            RATE-US PageType = "rate-us"
            
            
            //REFER-EARN defines constant for the `refer-earn` 
            REFER-EARN PageType = "refer-earn"
            
            
            //SETTINGS defines constant for the `settings` 
            SETTINGS PageType = "settings"
            
            
            //SHARED-CART defines constant for the `shared-cart` 
            SHARED-CART PageType = "shared-cart"
            
            
            //TNC defines constant for the `tnc` 
            TNC PageType = "tnc"
            
            
            //TRACK-ORDER defines constant for the `track-order` 
            TRACK-ORDER PageType = "track-order"
            
            
            //WISHLIST defines constant for the `wishlist` 
            WISHLIST PageType = "wishlist"
            
            
            //SECTIONS defines constant for the `sections` 
            SECTIONS PageType = "sections"
            
            
            //FORM defines constant for the `form` 
            FORM PageType = "form"
            
            
            //CART-DELIVERY defines constant for the `cart-delivery` 
            CART-DELIVERY PageType = "cart-delivery"
            
            
            //CART-PAYMENT defines constant for the `cart-payment` 
            CART-PAYMENT PageType = "cart-payment"
            
            
            //CART-REVIEW defines constant for the `cart-review` 
            CART-REVIEW PageType = "cart-review"
            
            
            //LOGIN defines constant for the `login` 
            LOGIN PageType = "login"
            
            
            //REGISTER defines constant for the `register` 
            REGISTER PageType = "register"
            
            
            //SHIPPING-POLICY defines constant for the `shipping-policy` 
            SHIPPING-POLICY PageType = "shipping-policy"
            
            
            //RETURN-POLICY defines constant for the `return-policy` 
            RETURN-POLICY PageType = "return-policy"
            
            
       )

    //IsValid return error if enum is invalid
    func (pa PageType) IsValid() error {
      switch pa {
        case  ABOUT-US, ADDRESSES, BLOG, BRANDS, CARDS, CART, CATEGORIES, BRAND, CATEGORY, COLLECTION, COLLECTIONS, CONTACT-US, EXTERNAL, FAQ, FRESHCHAT, HOME, NOTIFICATION-SETTINGS, ORDERS, PAGE, POLICY, PRODUCT, PRODUCT-REVIEWS, ADD-PRODUCT-REVIEW, PRODUCT-REQUEST, PRODUCTS, PROFILE, PROFILE-ORDER-SHIPMENT, PROFILE-BASIC, PROFILE-COMPANY, PROFILE-EMAILS, PROFILE-PHONES, RATE-US, REFER-EARN, SETTINGS, SHARED-CART, TNC, TRACK-ORDER, WISHLIST, SECTIONS, FORM, CART-DELIVERY, CART-PAYMENT, CART-REVIEW, LOGIN, REGISTER, SHIPPING-POLICY, RETURN-POLICY : 
        return nil
      }
      return errors.New("Invalid PageType type")
    }












    //SubscriberStatus used by Webhook
    type SubscriberStatus  string
    
    const (
            
            //ACTIVE defines constant for the `active` 
            ACTIVE SubscriberStatus = "active"
            
            
            //INACTIVE defines constant for the `inactive` 
            INACTIVE SubscriberStatus = "inactive"
            
            
            //BLOCKED defines constant for the `blocked` 
            BLOCKED SubscriberStatus = "blocked"
            
            
       )

    //IsValid return error if enum is invalid
    func (su SubscriberStatus) IsValid() error {
      switch su {
        case  ACTIVE, INACTIVE, BLOCKED : 
        return nil
      }
      return errors.New("Invalid SubscriberStatus type")
    }



