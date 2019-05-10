# ff follower monitor
## freefeed.net user subscribers monitor

This is a simple program to monitor changes in the follower list (subscribers)
for a *freefeed.net* user.

Download the source code: *ffmon.go*

Put the source code in a folder (ex: ffmon): after compiled, the executable 
will have the same name as the folder where it is contained.

It is written in GO language and must be compiled using the following command:
  **go build**

On the very first run, the program asks for Username and Password,
in order to access the *freefeed.net* account. If everything is ok,
an Authorization Token is retrieved and stored, along with Username and
Password in the same folder as the program. Files named "ffmon.conf" and
"ffmon.auth" will contain Username, Password (in clear text) and Token.

Any subsequent run will silently use those informations. If something
goes wrong, just remove the file "ffmon.conf" and try again.
