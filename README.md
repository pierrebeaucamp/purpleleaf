#Build instructions

You'll need [Python 2.7](https://www.python.org/) and [Go](https://golang.org/doc/install). Python is most likely already preinstalled on OS X.

Additionally, you have to install the [Google App Engine SDK for Go](https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go). Just unzip the package and put it anywhere you want. Ideally, you add the location to your PATH. You can find further instructions here: [https://cloud.google.com/appengine/docs/go/gettingstarted/devenvironment](https://cloud.google.com/appengine/docs/go/gettingstarted/devenvironment)

If you have installed all the dependencies, you can run ``goapp serve`` to run a local development server on [http://localhost:8080](http://localhost:8080). The development app server knows to watch for changes in your file. As you update your source, it recompiles them and relaunches your local app. There's no need to restart ``goapp serve``.

# [About](ABOUT.md)

# Creating dummy profiles:
If you haven't created some dummy profiles yet, click here: [http://localhost:8080/profile/dummies](http://localhost:8080/profile/dummies). If you want to access the local database directly, visit this page: [http://localhost:8000/datastore](http://localhost:8000/datastore)
