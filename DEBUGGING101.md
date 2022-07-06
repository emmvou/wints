# Finding errors in _wints_
## If the project does not launch correctly (e.g. the program crashes)
- read error message
- run project in debug mode
## If the project does launch correctly
- run backend in debug mode
- use web browser console to look for an error
- use web browser network tool to check requests made to backend's status and response
### If error comes from frontend
> frontend being automatically generated, it will be harder to debug
- check client data using the console and/or page elements
- check if previous calls to backend failed or responded with wrong/corrupted data
- use web browser debugger to look for an error
- check if frontend was correctly re-generated
- check if backend is running
#### If error still comes from frontend
- check with an old commit to compare what elements are generated and what are not
- use your intuition to guess what file to change
### If error comes from backend
- debug the called route, check if any error is `nil`
- read the error content
- when unexpected behavior is found, use stack tracing to find what caused it
- use and abuse breakpoints
#### if error comes from an sql statement
the requests can be found in `./sqlstore/` folder
- try the request in a different environment
- consider using an ORM