
go.work can be used instead of adding replace directives to work across multiple modules.

With `go work use <module_name>`, we can easily make changes to local modules and have them
used directly in the main module without needing to doing a replace of those modules that we did before
since those modules were not published.

Steps:
1. Add a new module or clone an existing module to the workspace.
2. Run the command : go work use <module_name> in the workspace directory.
3. This will add this module in the go.work file of the workspace as `use <module_name>`.
4. Because of this, now we can make changes to this module on the fly and have them used in the main module
   without the need of replacing it in the go.mod file.

To run the main package from the workspace directory - go run <main_module_name>.
This will only work if we have already attached the main module in the go.work file using `go work use <main_module_name>`.


Explore - https://go.dev/tour/methods/15
