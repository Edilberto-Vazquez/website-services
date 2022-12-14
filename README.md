# website-services

### Requirements

- have docker installed
- have docker compose installed

To run the server locally with docker

### Create the image with

`docker build . -t websiteapi.azurecr.io/website-api:latest`

### Run the command

`docker compose up -d`

##### Graphql

- Graphql playground: local [localhost:8000](http://localhost:8000/)
- Graphql query: [localhost:8000/query](http://localhost:8000/)

##### Queries availables for graphql

- fullProfile

##### REST-API

- Download resume English : [localhost:8000/resume-cv-en-US](http://localhost:8000/resume-cv-en)
- Download resume Spanish : [localhost:8000/resume-cv-es](http://localhost:8000/resume-cv-es)
- Get All data: [localhost:8000/api/v1/my-info/full-profile?lang=en-US](http://localhost:8000/api/v1/my-info/full-profile?lang=en-US)
- Get Profile: [localhost:8000/api/v1/my-info/profile?lang=en-US"](http://localhost:8000/api/v1/my-info/profile?lang=en-US)
- Get Projects: [localhost:8000/api/v1/my-info/projects?lang=en-US](http://localhost:8000/api/v1/my-info/projects?lang=en-US)
- Get Jobs: [localhost:8000/api/v1/my-info/jobs?lang=en-US](http://localhost:8000/api/v1/my-info/jobs?lang=en-US)

If you want a online version change `http://localhost:8000` for `https://website-api-service.azurewebsites.net`

#### Language

- If you want to change the language, change the lang parameter in the URL for rest

| First Header | Second Header |
| ------------ | ------------- |
| English      | `lang=en-US`  |
| Spanish      | `lang=es-MX`  |

- For graphql in variables section add the field

![](https://github.com/Edilberto-Vazquez/website-services/blob/main/graphql-example.png?raw=true)
