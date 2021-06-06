# Test client

With this client you can test auth service.

## Run

For running use:

```sh
go run .
```

## Example

Example of the successfull output:

```log
INFO[0000] Loaded config: {NATS:{Hostname:localhost Request:{Subject:auth-request Queue:robin}}} 
INFO[0000] Response: &{AuthUserResponse:success:true}
```