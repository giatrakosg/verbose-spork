# verbose-spork
Dropbox Syncing using Go. This project aims to recreate syncing multiple files
over a network using tcp sockets.

## server
The main server accepts connections from the clients and stores a list of files
and contents.

## client
Clients connect to the main server and sync their /data folder with that of the main server. We first get a list of all the data in a directory and hash their contents.
