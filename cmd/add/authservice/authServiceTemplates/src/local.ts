import app from './server';

const port = process.env.PORT || 3000;

app.listen(port, () =>
    app
        .get("logger")
        .info("Authentication Service is running on port " + port + "!")
);
