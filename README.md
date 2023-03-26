# iceyePoker - The Best Command Line Poker Game Known To Mankind!!
</br>

### iceyePoker is a LARVIS module, that lets you play poker with LARVIS, so you will never be bored again</br></br>Now every night can be Poker Night!

#### Modules used

- `github.com/fatih/color` - this module provides a means to output colored text to the console for a better user experience.
<br><br>I did consider using a module for the CLI interactions, but given the scope of this task, I decided against it as I felt that it wasn't really necessary
and it would also mean I can write more code to showcase.

#### Code

There are a total of 5 modules (excluding the `main` module), 1 top level module (`poker`) and 4 submodules within it.

- `poker` - this module is responsible for facilitating the poker game
  - `cli` - this module is responsible  for the cli interactions and validating all user inputs
  - `hand` - this module is responsible for analyzing the cards and determining how strong the hand is and comparing it with other hands
  - `console` - this module provides colored text output to the console for a better user experience
  - `errors` - custom errors types needed for example to help determine when a actual error occurs or when a user enters an erroneous input
  
  
### How to get the party started!

#### Run via Docker
From the root folder, run: `sudo ./run.sh`
#### or if you like to live life dangerously, you can go native
1. copy the `iceyePoker` folder to `$GOPATH/src`
2. `cd $GOPATH/src/iceyePoker`
3. `go get -d -v` - to download all the dependencies
4. `go run .` - to run iceyePoker
