---
title: shelves.proto
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - javascript--nodejs: Node.JS
  - ruby: Ruby
  - python: Python
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
highlight_theme: darkula
headingLevel: 2

---

<h1 id="shelves.proto">shelves.proto version not set</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

<h1 id="shelves.proto-LibraryService">LibraryService</h1>

## UpdateBook

<a id="opIdUpdateBook"></a>

> Code samples

```shell
# You can also use wget
curl -X PATCH ///v1/{book.name} \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```http
PATCH ///v1/{book.name} HTTP/1.1
Host: null
Content-Type: application/json
Accept: application/json

```

```javascript
var headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'

};

$.ajax({
  url: '///v1/{book.name}',
  method: 'patch',

  headers: headers,
  success: function(data) {
    console.log(JSON.stringify(data));
  }
})

```

```javascript--nodejs
const request = require('node-fetch');
const inputBody = '{
  "name": "string"
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'

};

fetch('///v1/{book.name}',
{
  method: 'PATCH',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json'
}

result = RestClient.patch '///v1/{book.name}',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json'
}

r = requests.patch('///v1/{book.name}', params={

}, headers = headers)

print r.json()

```

```java
URL obj = new URL("///v1/{book.name}");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("PATCH");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("PATCH", "///v1/{book.name}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`PATCH /v1/{book.name}`

> Body parameter

```json
{
  "name": "string"
}
```

<h3 id="updatebook-parameters">Parameters</h3>

|Parameter|In|Type|Required|Description|
|---|---|---|---|---|
|book.name|path|string|true|none|
|body|body|[shelvesBook](#schemashelvesbook)|true|none|

> Example responses

> 200 Response

```json
{
  "name": "string"
}
```

<h3 id="updatebook-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[shelvesBook](#schemashelvesbook)|

<aside class="success">
This operation does not require authentication
</aside>

## GetBook

<a id="opIdGetBook"></a>

> Code samples

```shell
# You can also use wget
curl -X GET ///v1/{name} \
  -H 'Accept: application/json'

```

```http
GET ///v1/{name} HTTP/1.1
Host: null

Accept: application/json

```

```javascript
var headers = {
  'Accept':'application/json'

};

$.ajax({
  url: '///v1/{name}',
  method: 'get',

  headers: headers,
  success: function(data) {
    console.log(JSON.stringify(data));
  }
})

```

```javascript--nodejs
const request = require('node-fetch');

const headers = {
  'Accept':'application/json'

};

fetch('///v1/{name}',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json'
}

result = RestClient.get '///v1/{name}',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('///v1/{name}', params={

}, headers = headers)

print r.json()

```

```java
URL obj = new URL("///v1/{name}");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "///v1/{name}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /v1/{name}`

<h3 id="getbook-parameters">Parameters</h3>

|Parameter|In|Type|Required|Description|
|---|---|---|---|---|
|name|path|string|true|none|

> Example responses

> 200 Response

```json
{
  "name": "string"
}
```

<h3 id="getbook-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[shelvesBook](#schemashelvesbook)|

<aside class="success">
This operation does not require authentication
</aside>

## DeleteBook

<a id="opIdDeleteBook"></a>

> Code samples

```shell
# You can also use wget
curl -X DELETE ///v1/{name} \
  -H 'Accept: application/json'

```

```http
DELETE ///v1/{name} HTTP/1.1
Host: null

Accept: application/json

```

```javascript
var headers = {
  'Accept':'application/json'

};

$.ajax({
  url: '///v1/{name}',
  method: 'delete',

  headers: headers,
  success: function(data) {
    console.log(JSON.stringify(data));
  }
})

```

```javascript--nodejs
const request = require('node-fetch');

const headers = {
  'Accept':'application/json'

};

fetch('///v1/{name}',
{
  method: 'DELETE',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json'
}

result = RestClient.delete '///v1/{name}',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.delete('///v1/{name}', params={

}, headers = headers)

print r.json()

```

```java
URL obj = new URL("///v1/{name}");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("DELETE");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("DELETE", "///v1/{name}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`DELETE /v1/{name}`

<h3 id="deletebook-parameters">Parameters</h3>

|Parameter|In|Type|Required|Description|
|---|---|---|---|---|
|name|path|string|true|none|

> Example responses

> 200 Response

```json
{}
```

<h3 id="deletebook-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[protobufEmpty](#schemaprotobufempty)|

<aside class="success">
This operation does not require authentication
</aside>

## CreateBook

<a id="opIdCreateBook"></a>

> Code samples

```shell
# You can also use wget
curl -X POST ///v1/{parent}/books \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```http
POST ///v1/{parent}/books HTTP/1.1
Host: null
Content-Type: application/json
Accept: application/json

```

```javascript
var headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'

};

$.ajax({
  url: '///v1/{parent}/books',
  method: 'post',

  headers: headers,
  success: function(data) {
    console.log(JSON.stringify(data));
  }
})

```

```javascript--nodejs
const request = require('node-fetch');
const inputBody = '{
  "name": "string"
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'

};

fetch('///v1/{parent}/books',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json'
}

result = RestClient.post '///v1/{parent}/books',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json'
}

r = requests.post('///v1/{parent}/books', params={

}, headers = headers)

print r.json()

```

```java
URL obj = new URL("///v1/{parent}/books");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "///v1/{parent}/books", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /v1/{parent}/books`

> Body parameter

```json
{
  "name": "string"
}
```

<h3 id="createbook-parameters">Parameters</h3>

|Parameter|In|Type|Required|Description|
|---|---|---|---|---|
|parent|path|string|true|none|
|body|body|[shelvesBook](#schemashelvesbook)|true|none|

> Example responses

> 200 Response

```json
{
  "name": "string"
}
```

<h3 id="createbook-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[shelvesBook](#schemashelvesbook)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocSprotobufempty">protobufEmpty</h2>

<a id="schemaprotobufempty"></a>

```json
{}

```

*service Foo {
      rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);
    }

The JSON representation for `Empty` is empty JSON object `{}`.*

### Properties

*None*

<h2 id="tocSshelvesbook">shelvesBook</h2>

<a id="schemashelvesbook"></a>

```json
{
  "name": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|name|string|false|none|Resource name of the book. It must have the format of "shelves/*/books/*". For example: "shelves/shelf1/books/book2".|

<h2 id="tocSshelveslistbooksresponse">shelvesListBooksResponse</h2>

<a id="schemashelveslistbooksresponse"></a>

```json
{
  "books": [
    {
      "name": "string"
    }
  ],
  "next_page_token": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|books|[[shelvesBook](#schemashelvesbook)]|false|none|The field name should match the noun "books" in the method name.  There will be a maximum number of items returned based on the page_size field in the request.|
|next_page_token|string|false|none|Token to retrieve the next page of results, or empty if there are no more results in the list.|

