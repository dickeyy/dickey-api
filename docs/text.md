# Text Endpoints

**Base URL**

```
https://api.dickey.gg
```

---

```
GET /text/reverse?text=x
```

> Reverses the input text (`x`)

Example: `/text/reverse?text=hello`

Response:

```
{
    "reversed": string,
    "original": string

}
```

Errors:

-   `400`: No text provided

---

```
GET /text/length?text=x
```

> Returns the length of the input text (`x`)

Example: `/text/length?text=hello`

Response:

```
{
    "length": {
        "characters": number,
        "words": number
    },
    "text": string
}
```

Errors:

-   `400`: No text provided

---

```
GET /text/uppercase?text=x
```

> Converts the input text (`x`) to uppercase

Example: `/text/uppercase?text=hello`

Response:

```
{
    "uppercase": string,
    "original": string
}
```

Errors:

-   `400`: No text provided

---

```
GET /text/lowercase?text=x
```

> Converts the input text (`x`) to lowercase

Example: `/text/lowercase?text=HELLO`

Response:

```
{
    "lowercase": string,
    "original": string
}
```

Errors:

-   `400`: No text provided

---

```
GET /text/replace?text=x&search=y&replace=z
```

> Replaces all occurrences of `search` with `replace` in the input text (`x`)

Example: `/text/replace?text=hello&search=he&replace=no`

Response:

```
{
    "replaced": string,
    "original": string
}
```

Errors:

-   `400`: No text provided

---

```
GET /text/lorem?length=x
```

> Generates `x` words of lorem ipsum text

Example: `/text/lorem?length=10`

Response:

```
{
    "lorem": string,
    "length": number
}
```

Constraints:

-   0 < `x` < 2000

Errors:

-   `400`: Invalid parameter

---
