# Set Memcache

### Cookie Returned To Client

Before we write a cookie back to the client, we always call `func makeCookie`

This is a good point to then also call a func to put our data into memcache.

This way, the value stored in the cookie will be the same as the value stored in memcache.

```go
func makeCookie(mm []byte, id string, req *http.Request) *http.Cookie {
	b64 := base64.URLEncoding.EncodeToString(mm)
	code := getCode(b64)
	cookie := &http.Cookie{
		Name:  "session-id",
		Value: id + "|" + b64 + "|" + code,
		// Secure: true,
		HttpOnly: true,
	}

	// send data to be stored in memcache
	storeMemc(mm, id, req)

	// send data to be stored in a cookie
	return cookie
}
```

```go
func storeMemc(bs []byte, id string, req *http.Request) {
	ctx := appengine.NewContext(req)

	item1 := memcache.Item{
		Key:   id,
		Value: bs,
	}

	memcache.Set(ctx, &item1)
	// production code should not ignore the error
}
```

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

### Change func Model

The function `func Model` returns the a value of type model - **this has all of the user's session data**.

We wil need to change `func Model` to see if we have data in memcache and, if so, to then get the data from there.

To do this, we will need to have access to the `*http.Request` as this is necessary to access memcache.

Let's modify `func Model` to have a parameter of type `*http.Request` ...

```go
// Model returns a value of type model
func Model(c *http.Cookie, req *http.Request) model {
	xs := strings.Split(c.Value, "|")
	usrData := xs[1]

	m := unmarshalModel(usrData)

	// if data is in memcache
	// get pictures from there
	// see refactor-notes.md for explanation
	id := xs[0]
	m2 := retrieveMemc(req, id)
	if m2.Pictures != nil {
		m.Pictures = m2.Pictures
		log.Println("Picture paths returned from memcache")
	}

	return m
}

```

Wherever func Model is called, we will need to update our code to ensure a value of type `*http.Request` is also passed in. 

WebStorm has a great feature which allows us to command-click the the identifier in the declaration of a func in order to see where that function is called.

### If Data There Is Data In Memcache ...

Now, anytime `func Model` is called, it will check to see if there is data in memcache and, if so, it will use that data:
 
 ```go
func retrieveMemc(req *http.Request, id string) model {
	ctx := appengine.NewContext(req)
	item, _ := memcache.Get(ctx, id)
	var m model
	if item != nil {
		m = unmarshalModel(string(item.Value))
	}
	return m
}
 ```
 
 ```go
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

### Refactored / Abstracted Code

Modularized code in `func Model` and put it in `func unmarshalModel`  

```go
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
### Change func makeCookie signature

We will need to change `func makeCookie` to have a parameter of type `*http.Request` ... 

```go
func makeCookie(mm []byte, id string, req *http.Request)  *http.Cookie 
```

Wherever `func makeCookie` is called, we will need to update our code to ensure a value of type `*http.Request` is also passed in. 

WebStorm has a great feature which allows us to command-click the the identifier in the declaration of a func in order to see where that function is called.

# Refactor Code For Appengine

`package main` ... to ... `package mem`

I could have called `package mem` something else like, oh, I don't know, maybe `package mickeymouse`

Took code out of `func main` and put it into `func init`

Added `app.yaml` file