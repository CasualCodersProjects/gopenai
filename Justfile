set dotenv-load

OPENAI_API_KEY := env_var("OPENAI_API_KEY")

getModels:
    curl https://api.openai.com/v1/models -H 'Authorization: Bearer {{OPENAI_API_KEY}}' 

test:
    go test .
