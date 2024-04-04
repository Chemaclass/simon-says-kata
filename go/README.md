# Simon Says Kata

Simon is a game of short-term memory skill. The device creates a series of tones and lights and requires a user to repeat the sequence. If the user succeeds, the series becomes progressively longer and more complex.

<p align="center">
  <img src="simon.png" width="350" title="simon game">
</p>

## Rules
- The computer generates a random color (green, red, blue or yellow) [_and a tone linked to that color_]
- If the player presses the correct combination of colors, a new color will be appended at the end of the list
- If the player selects an incorrect color, then the game is over
- The game runs until the player loses

## Output

```txt
Welcome to the game!
Simon will show you a sequence of colors, repeat the sequence by entering the colors separated by a space.
The colors are: Red [R], Blue [B], Yellow [Y] and Green [G].

Simon says: yellow
Your turn: y
Good job!

Simon says: yellow green
Your turn: y y
Good job!

Simon says: yellow green blue
Your turn: y g r
Game over!
```

## PHP with PHPUnit

1. Install [composer](https://getcomposer.org/) `curl -sS https://getcomposer.org/installer | php`
2. `composer install` (inside the `php` folder)
3. `vendor/bin/phpunit` or `composer test`

## Javascript with Jest

1. Install [Node](http://nodejs.org/)
2. `npm install` (Being in the javascript folder)
3. `npm test`

## [Typescript with Node](/typescript/README.md)

## Typescript with Deno

1. Install [Deno](https://deno.land/#installation)
2. `deno test` (Being in the typescript folder)

## Java with Junit and Mockito

1. Install dependencies and tests with Maven [mvn test].
2. Run the tests with the IDE

## Scala with Munit and Scalacheck

1. `sbt` (in _scala_ folder)
2. `~test` to run the tests on _hot reload_

### Linux/Mac
1. Install [SDKMan](https://sdkman.io/)
2. `sdk install java 11.0.12-open` install OpenJDK
3. `sdk install sbt` once SDKMan is installed.

### Windows
1. Install [OpenJDK](https://docs.microsoft.com/es-es/java/openjdk/download#openjdk-110141-lts--see-previous-releases)
2. Install [SBT](https://www.scala-sbt.org/download.html)

### Visual Studio Code
1. Download [Visual Studio Code](https://code.visualstudio.com/)
2. Install for VS Code [Metals](https://scalameta.org/metals/docs/editors/vscode)

## Kotlin with JUnit5 and MockK

1. By console: You can install dependencies and launch tests with `gradlew test`.
2. Using IDE: Simply open the project from the root of the Kotlin template.

## C# with xUnit (with Fluent Assertions) and NSubstitute

1. Install Microsoft Visual Studio Community 2022
2. Open the project and it will automatically download the required Nuguet packages

## Python with pytest

1. Install python 3.x
2. After downloading the source code into the */python/ folder, create a virtual environment:
3. `python3 -m venv env` 4.
4. Activate in virtual environment:
- windows: `.\env\Scripts\activate.bat` 5.
- linux/mac: `source env/bin/activate`.
5. `pytest` to run the tests.
