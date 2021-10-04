# Admin Notification Service - Dummy

This repository provides a Admin Notification Dummy Service that runs on port
8080 for the HTTP URL `/api/notification` to get a POST request with the
following JSON data

```json
    {
        "level": "warning",
        "employeeAbbreviation": "mmu",
        "message": "some message"
    }
```

You can call the service with `curl` to see if is running

```curl
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{
        "level": "warning",
        "employeeAbbreviation": "mmu",
        "message": "some message"
    }' http://localhost:8080/api/notify
```

To build the service run

```sh
go build -o admin-alarm
```

Afterwards the service can be started with `./admin-alarm`.
