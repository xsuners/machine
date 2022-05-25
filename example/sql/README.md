mysqldef -u root -p 123456 machine --export > machine.sql
mysqldef -u root -p 123456 machine --dry-run < machine.sql
mysqldef -u root -p 123456 machine < machine.sql