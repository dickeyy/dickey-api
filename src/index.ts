import { Hono } from "hono";
import { logger } from "hono/logger";

// create a new instance of Hono
const app = new Hono();

// middleware
app.use(logger());

// routes
import math from "./routes/math";
import time from "./routes/time";
import text from "./routes/text";

// base routes
app.get("/", (c) => c.json({ message: "hi. go to https://docs.api.dickey.gg :)" }));

// sub routes
app.route("/math", math);
app.route("/time", time);
app.route("/text", text);

export default app;
