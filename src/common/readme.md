## aws support
Need setup aws credentials: https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/ 

### Parameter Store

Use aws parameter Store as a configuration center, the values in the store contains both plain text and kms encrypted secret

### Dynamodb

construct Dynamodb client for access Dynamodb tables


## jwt 

### NewToken

Generated customized token with userId and email claims

### ParseToken

Verify the token passed within expiry time and have valid claims

### GetToken

Get token from query (access_code) or Authorization header


## middleware

contains authorization filter to ensure request have valid token, and error filler to ensure we can catch and log errors