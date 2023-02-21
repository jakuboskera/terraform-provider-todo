# Resource: todo_task

## Example Usage

```hcl
resource "todo_task" "coding" {
  text = "Create the best application ever"
  done = true
}
```

## Argument Reference

The following arguments are supported:

- **text** - (Required) The task details
- **done** - (Optional) Deciding if task is done, default `false`
