I # E-Mail Sender

This is a small console application to send mails over SMTP. I just tested it with GMail but other providers should work too.

### Command line parameters
| Argument  | Description   | Required                                              |         
| ----------|:-------------:|:-----------------------------------------------------:|          
| sender    | Sender of the mail.                                                   | yes      |
| receiver  | Receiver of the mail. Multiple receiver separated by semicolon.       | yes      |
| subject   | Subject of the mail                                                   | no       |
| message   | Message of the mail                                                   | no       |
| password  | Password to log into the SMTP server                                  | yes      |
| smtp      | Address of the SMTP server (e.g. smtp.gmail.com)                      | yes      |
| port      | Port of the SMTP server (e.g. 465 for gmail)                          | yes      |

### Build
Open the script **make.bash** with your favorite text-editor and and edit the array platforms. You find all available platforms [here](https://golang.org/doc/install/source#environment). Then just run the script. The assemblies will be in the sub-directory **assembly**.

### Usage
Example:
```
send_mail -sender sender@gmail.com -receiver "receiver1@outlook.com;receiver2@gmail.com" -subject testmessage -message "I great message" -password "my secret password" -smtp smtp.gmail.com -port 465 
```
