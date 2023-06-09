# A template for simple web services

## Generic CRUD handler and repository for structs

`datamodel.Repository` defines the interface for storage classes following the repository pattern, i.e. functions to
* read one element
* read all elements
* create a new element
* delete an existing element
* update an existing element

An example storage class using gorm is provided in `storage.GormRepository`. Additionally, default handlers for the
gin web framework provide a convenient way to connect persistent storage.

Always check the HTTP response status codes!

## Dummy example

The `datamodel.Dummy` struct is exposed at the following endpoints:
<table>
<tr>
<td>Operation</td>
<td>Endpoint</td>
<td>Verb</td>
<td>Parameter</td>
<td>Data</td>
<td>Result</td>
</tr>
<tr>
<td>List all</td>
<td>/api/dummy</td>
<td>GET</td>
<td>-</td>
<td>-</td>
<td>json:[{},{},...]</td>
</tr>
<tr>
<td>Create</td>
<td>/api/dummy</td>
<td>Post</td>
<td>-</td>
<td>json:{"field":"value"}</td>
<td>json:{"id":numeric,"field":"value"}</td>
</tr>
<tr>
<td>List one</td>
<td>/api/dummy/{id}</td>
<td>GET</td>
<td>numeric id</td>
<td>-</td>
<td>json:{"id":numeric,"field":"value"}</td>
</tr>
<tr>
<td>Delete one</td>
<td>/api/dummy/{id}</td>
<td>DELETE</td>
<td>numeric id</td>
<td>-</td>
<td>-</td>
</tr>
<tr>
<td>Update one</td>
<td>/api/dummy/{id}</td>
<td>PUT</td>
<td>numeric id</td>
<td>json:{"field":"value"}</td>
<td>json:{"id":numeric,"field":"value"}</td>
</tr>
</table>

### List all dummies

Request:

```bash
curl -s -H 'Accept: application/json' http://{socket}/api/dummy`
```

Response:
```json
[
    {
        "ID": 1,
        "field": "myvalue"
    },
    {
        "ID": 2,
        "field": "myvalue"
    }
]
```

### Create a new dummy
Request:
```bash
curl -s -X POST -H 'Content-Type: application/json' -H 'Accept: application/json' -d '{"field":"third"}' https://{hostname}/api/dummy/2`
```

Response:
```json
{
    "ID": 3,
    "field": "third"
}
```

### Get one specific dummy 
Request:
```bash
curl -s -H 'Accept: application/json' https://{hostname}/api/dummy/2`
```

Response:
```json
{
    "ID": 2,
    "field": "myvalue"
}
```

### Delete an existing dummy
Request:
```bash
curl -s -X DELETE -H 'Accept: application/json' https://{hostname}/api/dummy/2`
```

Response:
```json
{}
```

### Update an existing dummy
Request:
```bash
curl -s -X PUT -H 'Content-Type: application/json' -H 'Accept: application/json' -d '{"field":"3"}' https://{hostname}/api/dummy/3`
```

Response:
```json
{
    "ID": 3,
    "field": "3"
}
```
