

<div align="center">
<img width="413" alt="スクリーンショット 2022-12-28 19 57 13" src="https://user-images.githubusercontent.com/79553411/209801618-b4a5cd84-415d-411f-8c66-e54caf1d5586.png">

# RAC-TESTER

This is a test program for grism and real machine for [Robocup ssl](https://ssl.robocup.org/)

[Features](#Features) • [Usage](#Usage) • [Related](#Related) • [License](#License)

</div>

## Features

RAC-TESTER is CLI tool using [cobra](https://github.com/spf13/cobra) and [promptui](github.com/manifoldco/promptui).

You can test ssl robot with using TUI. 

<img width="543" alt="スクリーンショット 2022-12-29 4 21 16" src="https://user-images.githubusercontent.com/79553411/209862078-2f0eef2d-36f8-4318-b160-1a198d36a928.png">

## Usage

1. Download file from [release](https://github.com/lion-rion/rac-tester/releases/tag/test)

2. Use command below in the directory where the file exists.

```sh
./rac-test test 
```

if you use simulator, use --sim flag.

```sh
./rac-test test --sim
```

### Flags

`test` : Use test mode

`--sim` : Use simulator or not

More may be added in the future...


## Related

- [cobra](https://github.com/spf13/cobra) - CLI
- [promptui](github.com/manifoldco/promptui) - TUI
- [grSim](https://github.com/RoboCup-SSL/grSim) - Simulator
- [SSL](https://ssl.robocup.org/) - Robocup Small Size League


## Author

[@Lion](https://github.com/lion-rion)



## License 

[MIT license](https://en.wikipedia.org/wiki/MIT_License).
