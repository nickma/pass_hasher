from hashlib import sha1
import getpass

def hash_mysql_password(password):
    return "*" + sha1(sha1(getpass.getpass("New MySQL Password:").encode()).digest()).hexdigest()

# Example usage:
hashed_mysql_password = hash_mysql_password("my_secret_password")
print("MySQL 3.2.3 hashed password:", hashed_mysql_password)