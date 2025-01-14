## metal device reboot

Reboots a device

### Synopsis

Example:

metal device reboot --id [device_UUID]

	  

```
metal device reboot [flags]
```

### Options

```
  -h, --help        help for reboot
  -i, --id string   --id or -i [device_UUID]
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -o, --output string     Output format (*table, json, yaml)
      --search string     Search keyword for use in 'get' actions. Search is not supported by all resources.
      --token string      Metal API Token (METAL_AUTH_TOKEN)
```

### SEE ALSO

* [metal device](metal_device.md)	 - Device operations

