import { Hono } from "hono";

const math = new Hono();

// GET /math/prime/:number
math.get("/prime/:number", (c) => {
    const number = parseInt(c.req.param("number"));

    if (isNaN(number)) {
        return c.json({ error: { message: "Invalid number", code: 400 } }, 400);
    }

    if (number < 2) {
        return c.json({ error: { message: "Number must be greater than 1", code: 400 } }, 400);
    }

    for (let i = 2; i < number; i++) {
        if (number % i === 0) {
            return c.json({ isPrime: false, divisors: [i, number / i], message: `${number} is not prime` });
        }
    }

    return c.json({ isPrime: true, message: `${number} is prime` });
});

// GET /math/fibonacci/:number
math.get("/fibonacci/:number", (c) => {
    const number = parseInt(c.req.param("number"));

    if (isNaN(number)) {
        return c.json({ error: { message: "Invalid number", code: 400 } }, 400);
    }

    if (number < 0) {
        return c.json({ error: { message: "Number must be greater than 0", code: 400 } }, 400);
    }

    if (number === 0) {
        return c.json({ fibonacci: [0] });
    }

    // make sure the number isn't too big
    if (number > 1476) {
        return c.json({ error: { message: "Number is too big (sorry)", code: 400 } }, 400);
    }

    const fibonacci = [0, 1];

    for (let i = 2; i <= number; i++) {
        fibonacci[i] = fibonacci[i - 1] + fibonacci[i - 2];
    }

    return c.json({ sequence: fibonacci, message: `First ${number} numbers in the Fibonacci sequence` });
});

// GET /math/random-number (optional query parameters: min, max)
math.get("/random-number", (c) => {
    // get the min and max from the query (if they exist)
    const min = parseInt(c.req.query("min"));
    const max = parseInt(c.req.query("max"));

    // if the min or max are not numbers and they exist, return an error
    if ((!isNaN(min) && isNaN(max)) || (isNaN(min) && !isNaN(max))) {
        return c.json({ error: { message: "Invalid min or max", code: 400 } }, 400);
    }

    // if the min is greater than the max, return an error
    if (min > max) {
        return c.json({ error: { message: "Min must be less than max", code: 400 } }, 400);
    }

    // if the min and max are both numbers, return a random number between them
    if (!isNaN(min) && !isNaN(max)) {
        return c.json({
            number: Math.floor(Math.random() * (max - min + 1)) + min,
            message: `Random number between ${min} and ${max}`,
            range: [min, max]
        });
    }

    // if the min and max don't exist, return a random number between 0 and 1000
    return c.json({
        number: Math.floor(Math.random() * 1001),
        message: "Random number between 0 and 1000",
        range: [0, 1000]
    });
});

// GET /math/factorial/:number
math.get("/factorial/:number", (c) => {
    const number = parseInt(c.req.param("number"));

    if (isNaN(number)) {
        return c.json({ error: { message: "Invalid number", code: 400 } }, 400);
    }

    if (number < 0) {
        return c.json({ error: { message: "Number must be greater than 0", code: 400 } }, 400);
    }

    // make sure the number isn't too big
    if (number > 170) {
        return c.json({ error: { message: "Number is too big (sorry)", code: 400 } }, 400);
    }

    let factorial = 1;

    for (let i = 1; i <= number; i++) {
        factorial *= i;
    }

    return c.json({ factorial, message: `${number}! = ${factorial}` });
});

// GET /math/sqrt/:number
math.get("/sqrt/:number", (c) => {
    const number = parseFloat(c.req.param("number"));

    if (isNaN(number)) {
        return c.json({ error: { message: "Invalid number", code: 400 } }, 400);
    }

    if (number < 0) {
        return c.json({ error: { message: "Number must be greater than or equal to 0", code: 400 } }, 400);
    }

    return c.json({ result: Math.sqrt(number), message: `Square root of ${number}` });
});

// GET /math/abs/:number
math.get("/abs/:number", (c) => {
    const number = parseFloat(c.req.param("number"));

    if (isNaN(number)) {
        return c.json({ error: { message: "Invalid number", code: 400 } }, 400);
    }

    return c.json({ result: Math.abs(number), message: `Absolute value of ${number}` });
});

// GET /math/round/:number
math.get("/round/:number", (c) => {
    const number = parseFloat(c.req.param("number"));

    if (isNaN(number)) {
        return c.json({ error: { message: "Invalid number", code: 400 } }, 400);
    }

    return c.json({ result: Math.round(number), message: `Rounded value of ${number}` });
});

// GET /math/ceil/:number
math.get("/ceil/:number", (c) => {
    const number = parseFloat(c.req.param("number"));

    if (isNaN(number)) {
        return c.json({ error: { message: "Invalid number", code: 400 } }, 400);
    }

    return c.json({ result: Math.ceil(number), message: `Ceiled value of ${number}` });
});

// GET /math/floor/:number
math.get("/floor/:number", (c) => {
    const number = parseFloat(c.req.param("number"));

    if (isNaN(number)) {
        return c.json({ error: { message: "Invalid number", code: 400 } }, 400);
    }

    return c.json({ result: Math.floor(number), message: `Floored value of ${number}` });
});

// GET /math/sin/:number
math.get("/sin/:number", (c) => {
    const number = parseFloat(c.req.param("number"));

    if (isNaN(number)) {
        return c.json({ error: { message: "Invalid number", code: 400 } }, 400);
    }

    return c.json({ result: Math.sin(number), message: `Sine of ${number}` });
});

// GET /math/cos/:number
math.get("/cos/:number", (c) => {
    const number = parseFloat(c.req.param("number"));

    if (isNaN(number)) {
        return c.json({ error: { message: "Invalid number", code: 400 } }, 400);
    }

    return c.json({ result: Math.cos(number), message: `Cosine of ${number}` });
});

// GET /math/tan/:number
math.get("/tan/:number", (c) => {
    const number = parseFloat(c.req.param("number"));

    if (isNaN(number)) {
        return c.json({ error: { message: "Invalid number", code: 400 } }, 400);
    }

    return c.json({ result: Math.tan(number), message: `Tangent of ${number}` });
});

// GET /math/log/:number
math.get("/log/:number", (c) => {
    const number = parseFloat(c.req.param("number"));

    if (isNaN(number)) {
        return c.json({ error: { message: "Invalid number", code: 400 } }, 400);
    }

    return c.json({ result: Math.log(number), message: `Natural logarithm of ${number}` });
});

// GET /math/log10/:number
math.get("/log10/:number", (c) => {
    const number = parseFloat(c.req.param("number"));

    if (isNaN(number)) {
        return c.json({ error: { message: "Invalid number", code: 400 } }, 400);
    }

    return c.json({ result: Math.log10(number), message: `Base 10 logarithm of ${number}` });
});

// GET /math/log2/:number
math.get("/log2/:number", (c) => {
    const number = parseFloat(c.req.param("number"));

    if (isNaN(number)) {
        return c.json({ error: { message: "Invalid number", code: 400 } }, 400);
    }

    return c.json({ result: Math.log2(number), message: `Base 2 logarithm of ${number}` });
});

// GET /math/exp/:number
math.get("/exp/:number", (c) => {
    const number = parseFloat(c.req.param("number"));

    if (isNaN(number)) {
        return c.json({ error: { message: "Invalid number", code: 400 } }, 400);
    }

    return c.json({ result: Math.exp(number), message: `Exponential of ${number}` });
});

// GET /math/pow/?base=number&exponent=number
math.get("/pow", (c) => {
    const base = parseFloat(c.req.query("base"));
    const exponent = parseFloat(c.req.query("exponent"));

    if (isNaN(base) || isNaN(exponent)) {
        return c.json({ error: { message: "Invalid base or exponent", code: 400 } }, 400);
    }

    return c.json({ result: Math.pow(base, exponent), message: `${base}^${exponent}` });
});

export default math;
