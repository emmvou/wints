# TODO _urgent_
- change csv reader, for now it works with the old format, and only for si5. For now it concatenates the promotion and the major of a student to get their group (with no separator), but the whole group should be given in a single column as the backend treatment already supports it
- fix frontend (especially when displaying tables of students, consider using one table per group of students)

# TODO
- wipe conventions once imported
- do not remove student having a convention
- resilience to multi import with e-mail rewriting
- server-side caching of the store
	- store interface
	- cache implement interface
- test security ?
- don't change the reset token if not needed
PENDING statistics



# Year +1 todo
- library for mailing (robustness issues)
- structured logs to play with
- generate link for the major head & the tutor to ack the convention
- gin middleware ?	/ simplify a bit if possible
- Merge context inside session (with err variable made available)
- error embed http code
- use the email provided by Claudine for the student in place of the one in convention
- head par formation
- cumul head formation & major leader
- deadline final 2 weeks before the end of the internship

