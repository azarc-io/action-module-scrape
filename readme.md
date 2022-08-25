# VTH Module Submission

The VTH module scraper is a GitHub action that should be used on tag push.<br>
The action will scrape all the information of the repo and submit on success to veratherad.<br>

## Example module
An example module with one spark and connector is available [here](https://github.com/azarc-io/vth-module-hello-word).

## Module structure
The module structure needs to have the following folder structure.
All the files and folders need to be named exactly as indicated below.
```
[root]
|   readme.md                       (optional)
|   licence.txt                     (optional)
|   icon                            (optional)
|   module.yaml                     (required)
└───sparks (optional)
|       └───[name]
|           │   readme.md           (optional)
|           |   icon                (optional)
|           │   spark.yaml          (required)
|           │   input_schema.json   (required)
|           │   output_schema.json  (required)
└───connectors (optional)
        └───[name] (optional)
            │   readme.md           (optional)
            |   icon                (optional)
            │   connector.yaml      (required)
            │   schema.json         (required)

```

### Module YAML
The module config stored in __module.yaml__ contains the properties indicated below.
```
| Property      | Type      | Required  |
-------------------------------------------
| package       | string    | True      |
| label         | string    | True      |
| description   | string    | True      |
| tags          | []string  | True      |
```

### Spark YAML
The spark config stored in __spark.yaml__ contains the properties indicated below.
```
| Property      | Type      | Required  |
-------------------------------------------
| label         | string    | True      |
| description   | string    | True      |
```

### Connector YAML
The spark config stored in __connector.yaml__ contains the properties indicated below.
```
| Property      | Type      | Required  |
-------------------------------------------
| label         | string    | True      |
| description   | string    | True      |
```

### Icons
An icon's HxW cannot be larger than 500 000 pixels.</br> 
The supported image types are.
 - PNG
 - JPG
 - SVG
