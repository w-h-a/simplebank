## README

A simple bank configured with github workflows. The db is powered by Postgres. The API is written with Go and SQLC.

Currently there is no UI. Open up POSTMAN and try:

POST to `/users` with

```JSON
{
    "username": string,
    "password": string,
    "full_name": string,
    "email": string
}
```

POST to `/users/login` with

```JSON
{
    "username": string,
    "password": string
}
```

Grab the `access_token` (of type `Bearer Token`) that is returned in the response. You'll need it for every request type below.

POST to `/accounts` with

```JSON
{
    "owner": string,
    "currency": "USD|CAD|EUR"
}
```

Grab your account `id`.

GET `/accounts/:id`

GET `/accounts?page_id=<int>&page_size<int>`

Create another user, login, and create a new account. Then do the following (don't forget your account `id` and `access_token`).

POST to `/transfers` with

```JSON
{
    "from_account_id": <id of account that you are currently logged into>,
    "to_account_id": <id of the previous account you created>,
    "amount": int,
    "currency": "USD|CAD|EUR"
}
```

Log into the user account that you previously created and checkout your account(s)!
