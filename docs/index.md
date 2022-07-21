# TODO Provider

The TODO provider is used to configure an instance of [jakuboskera](https://github.com/jakuboskera)'s [TODO](https://github.com/jakuboskera/todo) application. The provider needs to be configured with the proper API key before it can be used.

## Resources

- [Resource: todo_task](resources/task.md)

## Authentication

```hcl
provider "todo" {
  url      = "https://todo.jakuboskera.dev"
  api_key  = "insert_api_key_here"
}
```

## Argument Reference

The following arguments are supported:

- **url** - (Required) The URL of the TODO application
- **api_key** - (Required) The API key to be used to access TODO application
- **insecure** - (Optional) Choose to ignore certificate errors, default `true`
