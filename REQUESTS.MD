POST /register
{
    "username":"",
    "password":"",
    "email":""
}
username, password, email decoded from JSON, password encrypted with MD5, saved to database

POST /login
{
    "username":""
    "password":""
}
{
    "email":""
    "password":""
}
Check username or email, password MD5 hash with DB

POST /get-messages
{
AuthToken:””
}

POST /get-message
{
  AuthToken:””

}

POST /message
{
Authtoken:“”
TargetId:””

}
