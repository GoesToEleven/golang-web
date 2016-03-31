# Setup For App Engine

Make sure your environment is configured for App Engine.

If you have not worked with App Engine yet, please see "Hello, World!" for Go on App Engine

https://cloud.google.com/appengine/docs/go/

# Create A Project on Google Cloud Platform (App Engine)

If you have not already, make sure you have a project on Google Cloud Platform 

https://console.cloud.google.com/project

# Create A Default Cloud Storage Bucket

At the time I wrote this ...

You find this by clicking the top-left hamburger menu

... then choosing ...

COMPUTER / APP ENGINE / SETTINGS / APPLICATION SETINGS 	

... or by following this link ...

https://console.cloud.google.com/appengine/settings

... and then clicking on the CREATE button at the bottom ...

Create
Default Cloud Storage Bucket
A free 5GB Cloud Storage bucket for App Engine applications, doesn't require billing enabled.

# Download These Packages

go get -u golang.org/x/oauth2
go get -u google.golang.org/cloud/storage
go get -u google.golang.org/appengine/...

# Configure Your App.yaml

Make sure it looks like this:

```go
application: <your-app-id-here>
version: v1
runtime: go
api_version: go1

handlers:
- url: /.*
  script: _go_app
```