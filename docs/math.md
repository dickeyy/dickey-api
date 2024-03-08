# Math Endpoints

**Base URL**

```
https://api.dickey.gg
```

---

```
GET /math/prime/:n
```

> Calculates if `n` is a prime.

Example: `/math/prime/20`

Response:

```
{
    "isPrime": boolean,
    "divisors": [
        number, number
    ],
    "message": string
}
```

Constraints:

-   `n` must be > 1
-   `n` must be a number

Errors:

-   `400`: Invalid parameter

---

```
GET /math/fibonacci/:n
```

> Calculates the first `n` digits in the Fibonacci Sequence

Example: `/math/fibonacci/5`

Response:

```
{
    "sequence": number[] (of size n),
    "message": string
}
```

Constraints:

-   0 < `n` < 1476
-   `n` must be a number

Errors:

-   `400`: Invalid parameter

---

```
GET /math/random-number?min=x&max=y
```

> Generates a random number between `x` and `y` (if `x` and `y` are omitted, `x=0 y=1000`)

Example: `/math/random-number?min=10&max=100`

Response:

```
{
    "number": number,
    "message": string,
    "range": [x,y]
}
```

Constraints:

-   `x (min)` < `y (max)`
-   `x` and `y` must be numbers

Errors:

-   `400`: Invalid parameter(s)

---

```
GET /math/factorial/:n
```

> Calculates the factorial of `n` (`n!`)

Example: `/math/factorial/5`

Response:

```
{
    "number": number,
    "message": string
}
```

Constraints:

-   0 < `n` < 170
-   `n` must be a number

Errors:

-   `400`: Invalid parameter

---

```
GET /math/sqrt/:n
```

> Calculates the square root of `n`

Example: `/math/sqrt/25`

Response:

```
{
    "result": number,
    "message": string
}
```

Constraints:

-   `n` > 0
-   `n` must be a number

Errors:

-   `400`: Invalid parameter

---

```
GET /math/abs/:n
```

> Calculates the absolute value of `n`

Example: `/math/abs/-5`

Response:

```
{
    "result": number,
    "message": string
}
```

Constraints:

-   `n` must be a number

Errors:

-   `400`: Invalid parameter

---

```
GET /math/round/:n
```

> Rounds `n` to the nearest integer

Example: `/math/round/5.5`

Response:

```
{
    "result": number,
    "message": string
}
```

Constraints:

-   `n` must be a number

Errors:

-   `400`: Invalid parameter

---

```
GET /math/ceil/:n
```

> Rounds `n` up to the nearest integer

Example: `/math/ceil/5.1`

Response:

```
{
    "result": number,
    "message": string
}
```

Constraints:

-   `n` must be a number

Errors:

-   `400`: Invalid parameter

---

```
GET /math/floor/:n
```

> Rounds `n` down to the nearest integer

Example: `/math/floor/5.9`

Response:

```
{
    "result": number,
    "message": string
}
```

Constraints:

-   `n` must be a number

Errors:

-   `400`: Invalid parameter

---

```
GET /math/sin/:n
```

> Calculates the sine of `n` (in radians)

Example: `/math/sin/0`

Response:

```
{
    "result": number,
    "message": string
}
```

Constraints:

-   `n` must be a number

Errors:

-   `400`: Invalid parameter

---

```
GET /math/cos/:n
```

> Calculates the cosine of `n` (in radians)

Example: `/math/cos/0`

Response:

```
{
    "result": number,
    "message": string
}
```

Constraints:

-   `n` must be a number

Errors:

-   `400`: Invalid parameter

---

```
GET /math/tan/:n
```

> Calculates the tangent of `n` (in radians)

Example: `/math/tan/0`

Response:

```
{
    "result": number,
    "message": string
}
```

Constraints:

-   `n` must be a number

Errors:

-   `400`: Invalid parameter

---

```
GET /math/log/:n
```

> Calculates the natural logarithm of `n`

Example: `/math/log/10`

Response:

```
{
    "result": number,
    "message": string
}
```

Constraints:

-   `n` > 0
-   `n` must be a number

Errors:

-   `400`: Invalid parameter

---

```
GET /math/log10/:n
```

> Calculates the base 10 logarithm of `n`

Example: `/math/log10/10`

Response:

```
{
    "result": number,
    "message": string
}
```

Constraints:

-   `n` > 0
-   `n` must be a number

Errors:

-   `400`: Invalid parameter

---

```
GET /math/log2/:n
```

> Calculates the base 2 logarithm of `n`

Example: `/math/log2/10`

Response:

```
{
    "result": number,
    "message": string
}
```

Constraints:

-   `n` > 0
-   `n` must be a number

Errors:

-   `400`: Invalid parameter

---

```
GET /math/exp/:n
```

> Calculates `e` raised to the power of `n`

Example: `/math/exp/1`

Response:

```
{
    "result": number,
    "message": string
}
```

Constraints:

-   `n` must be a number

Errors:

-   `400`: Invalid parameter

---

```
GET /math/pow/?base=x&exponent=y
```

> Calculates `x` raised to the power of `y`

Example: `/math/pow/?base=2&exponent=3`

Response:

```
{
    "result": number,
    "message": string
}
```

Constraints:

-   `x` and `y` must be numbers
-   `x` and `y` must be present

Errors:

-   `400`: Invalid parameter(s)

---
