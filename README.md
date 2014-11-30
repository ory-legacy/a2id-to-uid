a2id-to-uid
================

A [http://golang.org/](Go lang) webservice for mapping athene2 uuids to ory uids.

## Running

## API

### Creat a new entry

Request `POST /a2id2uids` with the following data:

Key  | Value | Description
------------- | ------------- | -------------
uid  | int    | Contains the new id represented by an int. Example: `1234`
a2id | string | Contains the id's table prefix as well as the id. Example: `"uuid1313"`

### Get an entry by it's a2id

Request `GET /as2id2uids/{{uid}}` and you'll receive `{}`, if no entry has been found and if an entry has been found:

```
{
  "result": "uuid1313"
}
```

### Get an entry by it's uid

Request `GET /as2id2uids/{{a2id}}` and you'll receive `{}`, if no entry has been found:

```
{
  "result": 1234
}
```

### Errors

There could be something wrong with the services health. You'll be notified when that happens:

```
HTTP/1.1 500 Internal Server Error
Content-Type: application/json; charset=utf-8
Content-Length: length

{
  "error": 500,
  "message": "MySQL Server has gone away, please start mysql."
}
```
