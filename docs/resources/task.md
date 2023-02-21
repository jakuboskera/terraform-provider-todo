# Resource: todo_task

## Example Usage

```hcl
resource "todo_task" "coding" {
  text = "Create the best application ever"
  is_done = true
}
```

## Argument Reference

The following arguments are supported:

- **text** - (Required) The task details
- **is_done** - (Optional) Deciding if task is done, default `false`
