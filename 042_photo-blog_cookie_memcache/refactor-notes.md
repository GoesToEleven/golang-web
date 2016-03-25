# Set Memcache

### Cookie Returned To Client

Before we write a cookie back to the client, we always call `func makeCookie`

This is a good point to then also call a func to put our data into memcache.

This way, the value stored in the cookie will be the same as the value stored in memcache.

### Add field to model

I'm adding this field to the model:

`req 	*http.Request`

This will allow us to get the context from the model.  We will need the context to store an item in memcache.

Notice that this field is lower case and not exported when we marshal our data.

### Change func Model signature

We wil need to change `func Model` to have a parameter of type `*http.Request` ...

```go
func Model(c *http.Cookie, req *http.Request) model 
```

... this way, whenever we ask for the model, it will have the current `*http.Request` value for the user.

Wherever func Model is called, we will need to update our code to ensure a value of type `*http.Request` is also passed in. 

WebStorm has a great feature which allows us to command-click the the identifier in the declaration of a func in order to see where that function is called.

### Change func makeCookie signature

We will need to change `func makeCookie` to have a parameter of type model ... 

```go
func makeCookie(m model, mm []byte, id string) *http.Cookie 
```

Wherever `func makeCookie` is called, we will need to update our code to ensure a value of type model is also passed in. 

WebStorm has a great feature which allows us to command-click the the identifier in the declaration of a func in order to see where that function is called.

# FYI, This Is An Unrealistic Example

FYI, this is an unrealistic example though it is building block in your educational process. 

Our uuid which uniquely identifies a user is stored in the cookie. Our []string which stores the paths to a user's photos are also stored in that cookie. 

We are now also storing all of that data in memcache. 

We access that data in memcache by the user's uuid (memcache stores key:value pairs). 

Well, to have the uuid, we have to have the cookie. 

And if we have the cookie, we have the []string with paths to the user's photos. 

**So why also get the []string from memcache? We already have it!**

Well, we're getting it from memcache just to learn this process. 

Eventually we will store our data in the datastore. We will have our uuid in the cookie. We will then check memcache for the []string which stores photo paths. If it's not in memcache, we will then check the datastore for this []string.

Eventually, we will also store the user's photos in google cloud storage (our hard drive in the cloud). 

So the whole process, at the end of this will be:

1. store uuid in **cookie**
1. store user session info in **memcache**
1. store user session info and user info in **datastore**
1. store user files in **google cloud storage**
1. attempt to retrieve user session info from **memcache**
  1. if unable to retrieve user session info from **memcache**, retrieve user session info from **datastore**
    1. store this session info in **memcache**
    1. next we retrieve user session info, it's in **memcache**
1. retrieve user photos from **google cloud storage**

# Retrieve Data From Memcache

### If Data There Is Data In Memcache ...

We will add this to `func model` so that anytime our code returns a model, it will check to see if there is data in memcache and, if so, it will use that data:
 
 ```go
 	id := xs[0]
 	m2 := retrieveMemc(req, id)
 	if m2.Pictures != "" {
 		m.Pictures = m2.Pictures
 		log.Println("Picture paths returned from memcache")
 	}
 ```

### Refactored / Abstracted Code

Modularized code in `func Model` and put it in `func unmarshalModel`  

```
func unmarshalModel(s string) model {

	bs, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		log.Println("Error decoding base64", err)
	}

	var m model
	err = json.Unmarshal(bs, &m)
	if err != nil {
		fmt.Println("error unmarshalling: ", err)
	}

	return m
}
```