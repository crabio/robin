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
INFO[0000] Send success auth request                    
INFO[0000] Response: &{AuthUserResponse:success:true}   
INFO[0000] Send failed auth request                     
INFO[0000] Response: &{Error:reason:"Failed to parse proto msg: proto: cannot parse invalid wire-format data"}
```