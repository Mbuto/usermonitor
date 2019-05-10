# ff follower monitor
freefeed.net user subscribers monitor

This is a simple program to monitor changes in the follower list (subscribers)
for a freefeed.net user.

Download the source code: 
It is written in GO language and must be compiled using the following command:
  go build
Put the source code in a folder (ex: ffmon): after compiled, the executable 
will have the same name of the folder.

On the very first run, the program asks for Username and Password,
in order to access the frefeed.net account. If everything is ok,
an Authorization Token is retrieved and stored, along with Username and
Password in the same folder as the program. Files named ffmon.conf and
ffmon.auth will contain Username, Password (in clear text) and Token.

Any subsequent run will silently use those informations. If something
goes wrong, just remove ffmon.conf and try again.
