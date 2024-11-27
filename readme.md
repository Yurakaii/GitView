
# GitView
### Disclaimer: This project is currently a work in progress and isn't fully functional.
## Outline
**GitView** is a GUI for git that is inspired by the one included in **Intellij IDEA**. **GitView** is intended to go into your project directory, make sure to add the configuration file to your **.gitignore** file.  
### NOTE: The configuration feature isn't currently implemented  
The configuration stores basic information such as, your git username, git email, and the git directory.
## Build Instructions
Make sure to have [**Golang**](https://go.dev/) installed.  
To build the program into a binary executable file, run the `make` command.  
If you want to run the program without making the build files, run the `make run` command.  
### Feature Checklist
- [x] Window Format
- [ ] List Changed Files
- [ ] Stage Changes
- [ ] Unstage Changes
- [ ] Amend Funcitonality
- [ ] Commit Name Input Functionality
- [ ] Commit Button Functionality
- [ ] Commit and Push Button Functionality
- [ ] Toolbar Functionality
### Bugfix Checklist
- [ ] Resizing not working
- [ ] Toolbar buttons won't shrink