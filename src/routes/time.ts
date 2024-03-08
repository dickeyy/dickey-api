import { Hono } from "hono";

const time = new Hono();

// GET /time/now
time.get("/now", (c) => {
    const now = new Date();

    return c.json({
        time: {
            unix: {
                seconds: Math.floor(now.getTime() / 1000),
                ms: now.getTime()
            },
            utc: now.toUTCString(),
            local: now.toString(),
            iso: now.toISOString(),
            date: now.toDateString(),
            time: now.toTimeString(),
            day: now.getDay(),
            month: now.getMonth(),
            year: now.getFullYear(),
            hours: now.getHours(),
            minutes: now.getMinutes(),
            seconds: now.getSeconds(),
            milliseconds: now.getMilliseconds(),
            timezone: now.getTimezoneOffset(),
            timezoneName: now.toTimeString().split(" ")[1]
        },
        message: "Current server time"
    });
});

// GET /time/unix/seconds/:timestamp (parameter: timestamp (unix seconds))
time.get("/unix/:timestamp", (c) => {
    const timestamp = parseInt(c.req.param("timestamp"));

    if (isNaN(timestamp)) {
        return c.json({ error: { message: "Invalid timestamp", code: 400 } }, 400);
    }

    const date = new Date(timestamp * 1000);

    return c.json({
        time: {
            unix: {
                seconds: timestamp,
                ms: timestamp * 1000
            },
            utc: date.toUTCString(),
            local: date.toString(),
            iso: date.toISOString(),
            date: date.toDateString(),
            time: date.toTimeString(),
            day: date.getDay(),
            month: date.getMonth(),
            year: date.getFullYear(),
            hours: date.getHours(),
            minutes: date.getMinutes(),
            seconds: date.getSeconds(),
            milliseconds: date.getMilliseconds(),
            timezone: date.getTimezoneOffset(),
            timezoneName: date.toTimeString().split(" ")[1]
        },
        message: "Time from unix timestamp"
    });
});

// GET /time/unix/ms/:timestamp (parameter: timestamp (unix ms))
time.get("/unix/ms/:timestamp", (c) => {
    const timestamp = parseInt(c.req.param("timestamp"));

    if (isNaN(timestamp)) {
        return c.json({ error: { message: "Invalid timestamp", code: 400 } }, 400);
    }

    const date = new Date(timestamp);

    return c.json({
        time: {
            unix: {
                seconds: Math.floor(timestamp / 1000),
                ms: timestamp
            },
            utc: date.toUTCString(),
            local: date.toString(),
            iso: date.toISOString(),
            date: date.toDateString(),
            time: date.toTimeString(),
            day: date.getDay(),
            month: date.getMonth(),
            year: date.getFullYear(),
            hours: date.getHours(),
            minutes: date.getMinutes(),
            seconds: date.getSeconds(),
            milliseconds: date.getMilliseconds(),
            timezone: date.getTimezoneOffset(),
            timezoneName: date.toTimeString().split(" ")[1]
        },
        message: "Time from unix ms timestamp"
    });
});

export default time;
