import app from './server';

const port = 3000;

app.listen(port, () =>
    app
        .get("logger")
        .info("Authentication Service is running on port " + port + "!")
);
