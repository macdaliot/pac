import {app} from './server'

const port = 3000;

app.listen(port, () => console.log("{{.serviceName}} is running on port ${port}!"))