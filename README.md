# Pass_hasher
Both python3 and goland version.

Sometime it is necery to validate userpassword was updated correctly by comparing the MySQL hash to hash generated in the terminal.

In go we hash the password twice using SHA-1 and we take hash to hexadecimal and prepend an asterisk (“*”)
