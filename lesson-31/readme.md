---
title: JSON Marshalling and Unmarshalling 
description: Working with JSON data using struct tags and encoding/json package
date: 2024-02-16
tags: ['go', 'json', 'marshalling', 'unmarshalling', 'struct-tags']
---

# JSON Marshalling and Unmarshalling

## 🔄 Marshalling vs Unmarshalling

### Marshalling (Encoding)
- **Go struct → JSON string**
- Uses `json.Marshal()` function
- Converts Go data structures to JSON format

### Unmarshalling (Decoding)  
- **JSON string → Go struct**
- Uses `json.Unmarshal()` function
- Converts JSON data to Go data structures

## 🏷️ JSON Struct Tags

JSON tags control how struct fields are serialized/deserialized to/from JSON.

### Tag Syntax
~~~go
`json:"field_name,option1,option2"`
~~~

## 📋 Common JSON Tag Options

### `omitempty`
- Omits field from JSON if value is empty
- Empty values: `0`, `false`, `""`, `nil`, empty slices/maps

### `-` (Dash)
- Completely excludes field from JSON processing
- Field is never included in marshalling/unmarshalling

### Field Name Customization
- Custom JSON field name different from struct field name

## 📤 Marshalling Example

### Go Struct to JSON
~~~go
type User struct {
    ID          int      `json:"user_id"`
    Name        string   `json:"name,omitempty"`
    Age         int      `json:"age"`
    Password    string   `json:"-"`
    Permissions []string `json:"roles"`
}
~~~

**Output JSON:**
~~~json
{
    "user_id": 1,
    "name": "John Doe",
    "age": 20,
    "roles": ["admin", "group-member"]
}
~~~

**Note:** `Password` field is excluded, field names are customized

## 📥 Unmarshalling Example

### JSON to Go Struct
~~~go
type Person struct {
    Name       string   `json:"full_name"`
    Age        int      `json:"years_old,omitempty"`
    Occupation string   `json:"-"`
    Language   []string `json:"spoken_languages"`
}
~~~

**Input JSON:**
~~~json
{
    "full_name": "Jane Doe",
    "years_old": 12,
    "spoken_languages": ["Chinese", "English"]
}
~~~

**Result:** Struct populated with mapped field names, `Occupation` remains empty

## ⚡ Key Points

### Marshalling Behavior
- Unexported struct fields are ignored
- Tags control field names and inclusion
- `omitempty` prevents zero values in output

### Unmarshalling Behavior  
- JSON fields map to struct fields via tags
- Extra JSON fields are ignored
- Missing JSON fields leave struct fields at zero values
