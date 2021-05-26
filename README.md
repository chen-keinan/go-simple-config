# go-simple-config
> Go Simple config is an open source configuration lib for storing and accessing configuration data with minimal dependencies
>

### supported configuration files:

- yaml
- json
- properties
- Environment variable

### usage example
```
func readConfig() error{
    c := New()
    err := c.Load("config.json")
    
    if err != nil {
       return err
     }	 
     
    fmt.Print(c.GetValueString("SERVER.host"))
}
```