## wm requests get

Get all requests in journal

### Synopsis

Get all requests in journal - if an id is specified, only that request is returned

```
wm requests get [flags]
```

### Examples

```
wm get requests
wm requests get --limit 5
wm requests get --id 0baca68a-0112-4f26-8529-ac12d8eb3720

```

### Options

```
  -h, --help        help for get
  -i, --id string   Specify the id of the specific request you want to return
  -l, --limit int   Limit the number of requests returned (only used if no id is specified) (default 10)
```

### Options inherited from parent commands

```
  -a, --admin-prefix string   Wiremock admin url prefix (default "__admin")
  -H, --host string           Wiremock host (default "http://localhost")
  -p, --port string           Wiremock port (default "8080")
```

### SEE ALSO

* [wm requests](wm_requests.md)	 - Handles all request commands for the Wiremock server

###### Auto generated by spf13/cobra on 27-Dec-2023