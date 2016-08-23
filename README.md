# Gate Keeper
**Permissions management and matching for Go projects**

Gate Keeper provides a powerful permissions toolkit which can easily
be integrated with most data stores and has been extensively used and
tested in production environments.

Permissions are represented as strings of the form `admin/users` or `project/:project`
with support for contextual replacements and wildcard permissions.

## Example

```go
m := gatekeeper.NewMatcher().WithPermissions([]string{"project/abc"}).WithContext(map[string]string{
    "project": "gatekeeper",
})

if m.Can("project/:project") {
    log.Println("Permission granted")
} else {
    log.Println("Permission denied)
}
```

## Usage
Gate Keeper's API is designed to make chaining for configuration as easy as possible.
Its configuration methods are prefixed with the term `With<Field>` and each returns a
new `gatekeeper.Matcher` object, allowing you to quickly fork configurations without
affecting previous instances.

It is expected that you'll create a new `gatekeeper.Matcher` for each request your
application processes, setting the permissions held by the user and the context in
which the request is executing. You can then use the `Can(permission string)` method
to quickly check whether the user has permission to conduct an action.

## Permission Design
Due to the way in which Gate Keeper's permissions are formatted, it is very simple
to use it for role based access control such as `user`, `administrator` etc.
Alternatively, you can opt for something as granular as your API, with permissions
matching your routes.

### Named Replacements
Another powerful tool is Gate Keeper's support for named replacements in permissions.
This allows you to check generic permissions and have the context hold sway over which
permissions are checked. For example, `project/:project` would replace `:project` with
the context's `project` value.

### Wildcard Permissions
You can also make use of wildcard permission, granting users access to all projects for
example, by granting them the `project/:project` permission explicitly (as opposed to
something like `project/gatekeeper`).