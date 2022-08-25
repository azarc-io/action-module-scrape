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
| Property          | Type      | Required  | Description
--------------------------------------------------------------------------------------
| package           | string    | True      | cononical package name give by azarc
| label             | string    | True      | display label
| description       | string    | True      | display description
| tags              | []string  | False     | search tags used in verathread
```

### Spark YAML
The spark config stored in __spark.yaml__ contains the properties indicated below.
```
| Property          | Type      | Required  | Description
--------------------------------------------------------------------------------------
| label             | string    | True      | display name
| description       | string    | True      | display description
| config            | object    | False     | unique config
| extensible_inputs | bool      | False     | ability to add inputs that are not specified 
| inputs            | object    | False     | see Spark Input
| outputs           | object    | False     | see Spark Output
```

#### Spark Input
The keys of inputs are the names. The value contains the remaining properties indicated below.
```
| Property          | Type      | Required  | Description
--------------------------------------------------------------------------------------
| mime_types        | []string  | True      | possible mime types
| schema            | string    | True      | relative path to the schema
| required          | bool      | False     | set if the input is required  
```

#### Spark Output
The keys of inputs are the names. The value contains the remaining properties indicated below.
```
| Property          | Type      | Required  | Description
--------------------------------------------------------------------------------------
| mime_type         | string    | True      | mime types
| schema            | string    | True      | relative path to the schema
```

### Connector YAML
The spark config stored in __connector.yaml__ contains the properties indicated below.
```
| Property          | Type      | Required  | Description
--------------------------------------------------------------------------------------
| label             | string    | True      | display name
| description       | string    | True      | display description
| config            | object    | False     | unique config
```

### Icons
An icon's HxW cannot be larger than 500 000 pixels.</br> 
The supported image types are.
 - PNG
 - JPG
 - SVG
