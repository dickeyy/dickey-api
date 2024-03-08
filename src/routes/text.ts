import { Hono } from "hono";

const text = new Hono();

// GET /text/reverse?text=string
text.get("/reverse", (c) => {
    const text = c.req.query("text");

    if (!text) {
        return c.json({ error: { message: "No text provided", code: 400 } }, 400);
    }

    // url decode the text
    const decoded = decodeURIComponent(text);

    return c.json({ reversed: decoded.split("").reverse().join(""), original: decoded });
});

// GET /text/length?text=string
text.get("/length", (c) => {
    const text = c.req.query("text");

    if (!text) {
        return c.json({ error: { message: "No text provided", code: 400 } }, 400);
    }

    // url decode the text
    const decoded = decodeURIComponent(text);

    return c.json({
        length: {
            characters: decoded.length,
            words: decoded.split(" ").length
        },
        text: decoded
    });
});

// GET /text/uppercase?text=string
text.get("/uppercase", (c) => {
    const text = c.req.query("text");

    if (!text) {
        return c.json({ error: { message: "No text provided", code: 400 } }, 400);
    }

    // url decode the text
    const decoded = decodeURIComponent(text);

    return c.json({ uppercase: decoded.toUpperCase(), original: decoded });
});

// GET /text/lowercase?text=string
text.get("/lowercase", (c) => {
    const text = c.req.query("text");

    if (!text) {
        return c.json({ error: { message: "No text provided", code: 400 } }, 400);
    }

    // url decode the text
    const decoded = decodeURIComponent(text);

    return c.json({ lowercase: decoded.toLowerCase(), original: decoded });
});

// GET /text/replace?text=string&search=string&replace=string
text.get("/replace", (c) => {
    const text = c.req.query("text");
    const search = c.req.query("search");
    const replace = c.req.query("replace");

    if (!text || !search || !replace) {
        return c.json({ error: { message: "Missing text, search, or replace", code: 400 } }, 400);
    }

    // url decode the text
    const decoded = decodeURIComponent(text);

    return c.json({ replaced: decoded.split(search).join(replace), original: decoded, search, replace });
});

// GET /text/lorem?length=number
text.get("/lorem", (c) => {
    const length = parseInt(c.req.query("length"));

    if (!length) {
        return c.json({ error: { message: "No length provided", code: 400 } }, 400);
    }

    if (isNaN(length)) {
        return c.json({ error: { message: "Length is not a number", code: 400 } }, 400);
    }

    if (length <= 0 || length > 2000) {
        return c.json({ error: { message: "Length is out of range", code: 400 } }, 400);
    }

    const lpText = `Lorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod tempor incididunt ut labore et dolore magna aliqua Ut enim ad minim veniam quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur Excepteur sint occaecat cupidatat non proident sunt in culpa qui officia deserunt mollit anim id est laborum`;
    const lpWords = lpText.split(" ");

    let result = "";
    for (let i = 0; i < length; i++) {
        result += lpWords[i % lpWords.length] + " ";
    }

    return c.json({ lorem: result.trim(), length });
});

export default text;
