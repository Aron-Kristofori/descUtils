# descUtils
descUtils is a simple terminal tool designed to simplify manual tasks associated with editing the descriptions of interfaces on Cisco IOS devices.

## Functionality
- Appending to a description

    Using the -a="text" flag you can specify a string to be added to the end of the description of interfaces on a Cisco IOS device.
- Prepending to a description

    Usigne the -p="text" flag you can specify a string to be added to the beginning of the deescription of interfaces on Cisco IOS devices.

## How does it work?
The program parses an input plain-text file which has the output of the Cisco IOS `show interfaces description` command. The program then appends or prepends the specified string and the finished product is saved in out.txt or another .txt file if specified.

## Command Syntax
    descUtils [<flags>...] <cr>

### Flags
- -a="text" : specifies the text to be appended to the description
- -p="text" : specifies the text to be prepended to the description
- -i="[PATH]" : specify a path to the input plain-text file
- -o="[PATH]" : specify a path to the output plain-text file
- -h/--help : help screen