# Codegate 

Codegate is a simple and easy-to-use pre-commit tool that helps automate checks before committing your code.

## Installation

1. Download the tool.
2. Move the downloaded file to /usr/local/bin/ and rename it to codegate.


## Initialization and Usage
1. Run the following command to initialize Codegate in your project:
```
codegate --mode init
```

2. After running the command, a .codegate folder will be created in your project, containing a pre-commit file. You can edit this file to configure the checks you want to run before each commit.



3. From now on, before every commit, Git will automatically run the commands specified in the pre-commit file.
