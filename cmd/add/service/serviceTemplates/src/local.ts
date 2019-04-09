import app from "./server";

const port = 3000;

app.listen(port, () =>
    app
        .get("logger")
        .info("{{.serviceName}} Service is running on port " + port + "!")
);
