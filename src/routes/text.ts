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

export default text;
