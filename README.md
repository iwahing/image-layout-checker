# image-layout-checker
Checks if the images are what is specified in the template
Checks given folder if new image has been resize correctly

##  To Run
```bash
go run .
```

##  To build windows
```bash
fyne-cross windows -arch=amd64
```
![High Level](image-2.png)

### Template File
File must be in .csv format and should follow this file formatting
![sizing.csv](image.png)


### Foldering format
Folder must follow template to match
![Folder](image-1.png)