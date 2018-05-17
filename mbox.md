---
title: mbox.proto
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

<h1 id="mbox.proto">mbox.proto version not set</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

<h1 id="mbox.proto-Mail">Mail</h1>

## Receive

<a id="opIdReceive"></a>

> Code samples

```shell
# You can also use wget
curl -X GET ///message \
  -H 'Accept: application/json'

```

```http
GET ///message HTTP/1.1
Host: null

Accept: application/json

```

```javascript
var headers = {
  'Accept':'application/json'

};

$.ajax({
  url: '///message',
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

fetch('///message',
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

result = RestClient.get '///message',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('///message', params={

}, headers = headers)

print r.json()

```

```java
URL obj = new URL("///message");
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
    req, err := http.NewRequest("GET", "///message", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /message`

<h3 id="receive-parameters">Parameters</h3>

|Parameter|In|Type|Required|Description|
|---|---|---|---|---|
|name|query|string|false|none|

> Example responses

> 200 Response

```json
{
  "from": "string",
  "to": [
    "string"
  ],
  "subject": "string",
  "text": "string",
  "html": "string"
}
```

<h3 id="receive-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|(streaming responses)|[mboxMessage](#schemamboxmessage)|

<aside class="success">
This operation does not require authentication
</aside>

## Send

<a id="opIdSend"></a>

> Code samples

```shell
# You can also use wget
curl -X POST ///message \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```http
POST ///message HTTP/1.1
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
  url: '///message',
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
  "from": "string",
  "to": [
    "string"
  ],
  "subject": "string",
  "text": "string",
  "html": "string"
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'

};

fetch('///message',
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

result = RestClient.post '///message',
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

r = requests.post('///message', params={

}, headers = headers)

print r.json()

```

```java
URL obj = new URL("///message");
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
    req, err := http.NewRequest("POST", "///message", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /message`

> Body parameter

```json
{
  "from": "string",
  "to": [
    "string"
  ],
  "subject": "string",
  "text": "string",
  "html": "string"
}
```

<h3 id="send-parameters">Parameters</h3>

|Parameter|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[mboxMessage](#schemamboxmessage)|true|none|

> Example responses

> 200 Response

```json
{}
```

<h3 id="send-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[mboxResult](#schemamboxresult)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocSmboxmessage">mboxMessage</h2>

<a id="schemamboxmessage"></a>

```json
{
  "from": "string",
  "to": [
    "string"
  ],
  "subject": "string",
  "text": "string",
  "html": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|from|string|false|none|none|
|to|[string]|false|none|none|
|subject|string|false|none|none|
|text|string|false|none|none|
|html|string|false|none|none|

<h2 id="tocSmboxresult">mboxResult</h2>

<a id="schemamboxresult"></a>

```json
{}

```

### Properties

*None*

