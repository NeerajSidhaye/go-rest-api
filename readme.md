## This feature branch "feature/go-http-webservice" exposing two simeple webservice using Go standard libraries.

**You will notice that I am using two "main.go" files in this project.**

The reason for this is - my own learning :-)

**/hello service is exposed using a http default handler named "DefaultServeMux"**

**/getProducts service is exposed with a custom Product handler and a data layer.**

## /hello

This is basic hello world http service, which exposing /hello end point at port 7070

## <details><summary>Click for demo </summary>


![helloservice](goservice/static/readmeimages/helloservice.gif)

</details>


 ## /getProducts

This service returns list of products using a Product handler, which then calls the data layer to return list of prodcuts to handler.
    
Prodcuts data is exposing a method ToJSON which is marshalling collection of products to JSON using go's json NewEncoder.

## <details><summary> click for Demo </summary>

![productservice](goservice/static/readmeimages/productservice.gif)

</details>


