# mysqlinit
Initializes mysql database in a generic way so that i can reuse the code in future projects as a package

everything except the database name has a default value that can be discovered in the variable
declaration section

If you do not need to customize the configuration, just call ConnectDefault(databaseName string) with
the name of the database you wish to connect to.