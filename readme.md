## This feature branch "feature/go-http-webservice" exposing two simple webservice using Go standard libraries.

**You will notice that I am using two "main.go" files in this project.**

**/hello service is exposed using a http default handler, which is "DefaultServeMux"**

**/getProducts service is exposed with a custom Product handler and a data layer.**

## /hello

This is basic hello world http service, which exposing /hello end point at port 7070

<details><summary>Click for Demo </summary>

![helloservice](goservice/static/readmeimages/helloservice.gif)

</details>


 ## /getProducts

This service returns list of products using a Product handler, which then calls the data layer to return list of prodcuts to handler.
    
Prodcuts data is exposing a method ToJSON which is marshalling collection of products to JSON using go's json NewEncoder.

<details><summary> Click for Demo </summary>

![productservice](goservice/static/readmeimages/productservice.gif)

</details>


