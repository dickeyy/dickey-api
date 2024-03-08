# Time Endpoints

**Base URL**

```
https://api.dickey.gg
```

---

```
GET /time/now
```

> Returns the current server time in various formats

Example: `/time/now`

Response:

```
{
    "time": {
        "unix": {
            "seconds": number,
            "ms": number
        },
        "utc": string,
        "local": string,
        "iso": string,
        "date": string,
        "time": string,
        "day" number,
        "month": number,
        "year": number,
        "hours": number,
        "minutes": number,
        "seconds": number,
        "milliseconds": number,
        "timezone": string,
        "timezoneName": string
    },
    "message": string
}
```

---

```
GET /time/unix/:t
```

> Converts a unix timestamp (seconds) to various formats

Example: `/time/unix/1709884384`

Response:

```
{
    "time": {
        "unix": {
            "seconds": number,
            "ms": number
        },
        "utc": string,
        "local": string,
        "iso": string,
        "date": string,
        "time": string,
        "day" number,
        "month": number,
        "year": number,
        "hours": number,
        "minutes": number,
        "seconds": number,
        "milliseconds": number,
        "timezone": string,
        "timezoneName": string
    },
    "message": string
}
```

---

```
GET /time/unix/ms/:t
```

> Converts a unix timestamp (milliseconds) to various formats

Example: `/time/unix/ms/1709884384193`

Response:

```
{
    "time": {
        "unix": {
            "seconds": number,
            "ms": number
        },
        "utc": string,
        "local": string,
        "iso": string,
        "date": string,
        "time": string,
        "day" number,
        "month": number,
        "year": number,
        "hours": number,
        "minutes": number,
        "seconds": number,
        "milliseconds": number,
        "timezone": string,
        "timezoneName": string
    },
    "message": string
}
```

---
