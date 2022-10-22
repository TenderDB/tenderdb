# TenderDB. Golang App Deploying without Docker

`TenderDB` is a web service for tender analysis written in Go. It's can be deployed with Makefile on your host without `Nodejs` and `Docker` installed. Please visit [tenderdb.ru](https://tenderdb.ru) to check how it works. 

## Using

* **BoltDB** embedded key-value database
* **Chi** as a router compatible with `net/http`
* **OAuth 2.0** with Google and Yandex endpoints  
* **Gorilla sessions** to login and logout
* **Envconfig** to configure
* **Vue.js** for frontend application
* **Google Charts** for visualization  
* **Makefile** & **Systemd** to deploy  

## Features

* Access to charts in unauthorized mode
* Access to filter by regions in authorized mode 
* Access to Excel `csv` downloading in authorized mode  
* Limitation of downloads `csv` per time
* Access to personal charts collection in authorized mode 
* Testing modes without OAuth by `BasicAuth` 

## Install

Need Git, Golang, Systemd to be installed.

* `git clone  github.com/tenderdb/tenderdb.git`
* `cd tenderdb`
* `make all`
* check `localhost:8000` in browser
* use `localhost:8000/testmode` to test authorization

Databases populated by not real numbers.

## License
Licensed under [MIT License](./LICENSE)

