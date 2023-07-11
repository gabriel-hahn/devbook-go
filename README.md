# Devbook Go

A Golang project developed for studying purposes (CRUD)

## Getting Started

### Generating the authentication secret

We recommend creating a secret using the **base64** approach, similar to the following:

```
func init() {
  key := make([]byte, 64)

  if _, err := rand.Read(key); err != nil {
    log.Fatal(err)
  }

  base64Str := base64.StdEncoding.EncodeToString(key)
  fmt.Println(base64Str)
}
```

This function creates a base64 secret key that can be used in the **.env** file.
