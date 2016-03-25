# Set Memcache

## Cookie Returned To Client

Before we write a cookie back to the client, we always call func makeCookie.

This is a good point to then also call a func to put our data into memcache.

This way, the value stored in the cookie will be the same as the value stored in memcache.

### Add field to model

I'm adding this field to the model:

req	*http.Request

This will allow us to get the context from the model.  We will need the context to store an item in memcache.

Notice that this field is lower case and not exported when we marshal our data.

### Change func Model signature

We wil need to change func Model to have a parameter of type *http.Request ...

```
func Model(c *http.Cookie, req *http.Request) model 
```

... this way, whenever we ask for the model, it will have the current *http.Request value for the user.

Wherever func Model is called, we will need to update our code to ensure a value of type *http.Request is also passed in. 

WebStorm has a great feature which allows us to command-click the the identifier in the declaration of a func in order to see where that function is called.

### Change func makeCookie signature

We will need to change func makeCookie to have a parameter of type model ... 

func makeCookie(m model, mm []byte, id string) *http.Cookie 

Wherever func Model is called, we will need to update our code to ensure a value of type model is also passed in. 

WebStorm has a great feature which allows us to command-click the the identifier in the declaration of a func in order to see where that function is called.

# Retrieve Data From Memcache Instead of Cookie

## Unrealistic Example

FYI, this is an unrealistic example though it is building block in your educational process. Our uuid which uniquely identifies a user is stored in the cookie. Our []string which stores a user's photos path information are also stored in the cookie. We are now also storing all of that data in memcache. We access that data by the user's uuid (memcache stores key:value pairs). Well, to have the uuid, we have to have the cookie. If we have the cookie, we have the []string with the user's photos. So why also get the []string of photo paths from memcache? We already have it! 

Well, we're getting it from memcache just to learn this process. 

Eventually we will store our data in the datastore. We will have our uuid in the cookie. We will then check memcache for the []string which stores photo paths. If it's not in memcache, we will then check the datastore for this []string.

Eventually, we will also store the user's photos in google cloud storage (our hard drive in the cloud). 

So the whole process, at the end of this will be:
1 store uuid in cookie
1 store user session info in memcache
1 store user session info and user info in datastore
1 store user files in google cloud storage
1 attempt to retrieve user session info from memcache
  *if unable to retrieve user session info from memcache, retrieve user session info from datastore
1 retrieve user photos from google cloud storage

### Update func Model signature

The user data we work with in our program is a value of type model.

Currently, whenever we get data from a cookie, we call func Model in order to take the data in the cookie and put it into a value of type model.

Currently func Model is declared with a parameter of type *http.Cookie. We then ask for the value of that cookie. This value is a string. The string is the marshalled data.
 
 We can change func Model to take in a string instead of a cookie.
 
 From this ...
 
 ```func Model(c *http.Cookie, req *http.Request) model```
 
 ... to this ...
 
 ```func Model(s string, req *http.Request) model```

### Update all calls of func Model

From this ...

``` m := Model(cookie, req) ```

... to this ...


``` m := Model(cookie.Value, req) ```

### Add conditional logic to all calls of func Model

From this ...

``` m := Model(cookie.Value, req) ```

... to this ...


``` 

```







Before we ask for a cookie ...


```cookie, _ := req.Cookie("session-id")```

... we first need to check mecache to see if the data is there. If the data is in memcache, there's no point in requesting the cookie.