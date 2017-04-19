asolution
==========

*Work in Progress* 

Beego with Web Components: server rendered and single page apps

   * Components for Page Views & Single Page Applets


This project is an example layout to start using web components in your next web app backed in go.
Supports server side rendering and single page app(lets)


Comes with:

   * Polymer Components
   * Registration
   * Login
   * Authentication
   * Authorization
   * Lots of comments to understand Beego App developement

Notes:

   * bower_components renamed to components in static/.bowerrc
   * run all bower tasks inside "/static" where components folder is
   * The project is setup for development; it is at times best to concatenize polymer components with "vulcanize" for production which may cause url resolution to change slightly.  This concatenation feature is however a built in effect of the HTTP 2.0 standard once fully implemented.  Benchmark all options to find the best rule of thumb for your application.


Requires:

    Active Go Installation 
    Beego
    Bee Tool
    Configured for Postgres db and Redis for Sessions
    Bower for new and updated web components
    Creative Restraint
    Unstoppable Tenacious Unbending Intent


Development:

    Fork
    Enhance with Good Intentions
    Pull Request
    MIT Licensed


Basic Getting Started Outline:

    git clone to $GOPATH
    create main bego app
    create single page applets in components/apps
    serve


Essential Reading for App Development with Beego and or Go:

* [Astaxie](http://astaxie.gitbooks.io/build-web-application-with-golang/)

* [Richard Eng](https://medium.com/@richardeng/a-word-from-the-beegoist-d562ff8589d7)

* [Will Krause](http://hobbyisthacker.com/authentication-with-beego-pt-1-env-setup-and-table-generation/)

* [Beego Documentation](http://beego.me/docs/intro/)

* [William Kennedy](http://www.goinggo.net/)
