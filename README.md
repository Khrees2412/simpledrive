[WIP]
Cloudvest - A simple google-drive-like tool 

This application safely and securely stores your file in the cloud.

Stack:

-   language: Go
-   Web framework: Fiber
-   Cloud storage: AWS S3
-   Database: PostgreSQL

# How to run

## Prerequisites

-   Go compiler (download at https://golang.org)
-   Docker

Base Endpoint: https://cloudvest.herokuapp.com/api/v1

Auth endpoints:

## Register route: /auth/register

**POST**
body: {
name : string,
email: string,
password: string
}

## Login route: /auth/login

**POST**
body: {
email: string,
password: string
}

# File endpoints: prefixed with **/drive**

## Store a file: /upload

**POST**
body: {
file: Formdata
}

## Download a file: /download/{fileName}

**GET**

## Get all files a user has stored: /files

**GET**

## Delete a file: /{fileName}

**DELETE**

Folder endpoints:

## Create a folder: /create-folder

**POST**
body: {
name: string
}

## Store a file in a folder: /upload/{folderName}

**POST**
body: {
file: Formdata
}
