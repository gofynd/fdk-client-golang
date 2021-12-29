package platform

import (
    "errors"
)










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
            
            
       )

    //IsValid return error if enum is invalid
    func (pa PageType) IsValid() error {
      switch pa {
        case  ABOUT-US, ADDRESSES, BLOG, BRANDS, CARDS, CART, CATEGORIES, BRAND, CATEGORY, COLLECTION, COLLECTIONS, CONTACT-US, EXTERNAL, FAQ, FRESHCHAT, HOME, NOTIFICATION-SETTINGS, ORDERS, PAGE, POLICY, PRODUCT, PRODUCT-REVIEWS, ADD-PRODUCT-REVIEW, PRODUCT-REQUEST, PRODUCTS, PROFILE, PROFILE-BASIC, PROFILE-COMPANY, PROFILE-EMAILS, PROFILE-PHONES, RATE-US, REFER-EARN, SETTINGS, SHARED-CART, TNC, TRACK-ORDER, WISHLIST, SECTIONS, FORM, CART-DELIVERY, CART-PAYMENT, CART-REVIEW : 
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
            
            
       )

    //IsValid return error if enum is invalid
    func (su SubscriberStatus) IsValid() error {
      switch su {
        case  ACTIVE, INACTIVE : 
        return nil
      }
      return errors.New("Invalid SubscriberStatus type")
    }

