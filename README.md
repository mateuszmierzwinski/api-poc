# API Draft #

## Mission updates ##

### Access ###

**Endpoint:** /mission/:learnerId/latest

**Method:** GET

**Notes:** For now learnerId can be any integer. I will generate random content each request.

### Contract ###

```json
{
  "updates": [
    {
        "image": "https://someimageurl/",
        "status": "green",
        "title": "Mission 1 title",
        "link": "https://math-stg.aida.pearson.com/api/v1/assets/videos/467",
    },{
        "image": "https://someimageurl/",
        "status": "yellow",
        "title": "Mission 2 title",
        "link": "https://math-stg.aida.pearson.com/api/v1/assets/videos/467",
    },
  ]
}
```

**Status:** We didn't covered that so I can assume "green", "yellow", "red" for now.

**Images:** Images will be pushed also right away. This content will be using standard HTTP/2 Push.



# Focused Practice #

### Access ###

**Endpoint:** /practice/:learnerId/focused

**Method:** GET

**Notes:** For now learnerId can be any integer. I will generate random content each request.

### Contract ###

```json
{
    "title": "Some Practice title",
    "description": "This is practice description that can be lorem ipsum",
    "icon": "icon_globe",
    "link": "https://math-stg.aida.pearson.com/api/v1/assets/videos/467"
}
```

**Link:** I do believe this widget needs to be clickable and move somewhere - let's link it for now with video that you've provided.


# Content sitemap - Question mark button #

### Access ###

**Endpoint:** /sitemap/

**Method:** GET

**Notes:** This is static content index that can be shown.

### Contract ###

```json
{
  "sitemap": {
    "about_page": {
      "title": "About information",
      "link": "https://testing-stream-service-uqekqerd5a-uc.a.run.app/static/about.html"
    }
  }
}
```

# Notifications #

### Access ###

**Endpoint:** /notification

**Method:** GET

**Notes:** This is notifications HTTP/2 stream endpoint. Notification will be sent after 10 seconds.




