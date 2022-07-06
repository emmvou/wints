# Roles' mechanisms
Roles are used to grant permissions to users. Users can be granted multiple roles.
Here they go in increasing order of permissions :  
- Student  
- Tutor  
- Supervisor  
- Head  
- Admin  
- Root  

A Student is supposed to only be granted the student role.  
A Head is also supposed to be a Supervisor, as it is in charge of a large group but has extra permissions  
When checking if a simple permission is granted, highest role is used (as roles are ordered)  
Supervisor role usually comes with overviewed group as followed `(supervisor\-[a-zA-Z1-9]*)` when checking permissions related to overviewing specific students  
Keeping this amount of roles is strongly recommanded as frontend is generated with the roles ranks  

Definition can be found in multiple places :
- [frontend/home](./assets/js/home.js)
- [frontend/templates](./assets/js/templates.js)
- [backend/schema](./schema/user.go)  

and used widely throughout the codebase