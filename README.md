# runner
> the name of this project can be changed in the future when i have a better idea

## What is runner ?

Runner is a simple project for deploying a instance on [vultr](https://www.vultr.com/) and deploy a api ([api_runner](https://github.com/tot0p/api_runner)) on it. When the instance is deployed, the api will be available at the ip of the instance and you can use it to deploy github repository on the instance (the repository must contain a `Dockerfile`).

An `cli` is included with `runner` to interact with the api.

## How to install

you need to have `go 1.21.1 or higher` installed on your machine.

```bash
git clone https:://github.com/tot0p/runner.git
cd runner
go build .
```

and you need to create a `.env` file in the root of the binary with the following content:

```bash
touch .env && echo "API_KEY=YOUR_API_KEY_OF_VULTR" >> .env
```
> you can get your api key from [vultr](https://my.vultr.com/settings/#settingsapi)

## How to use

```bash
./runner
``` 

and the program will be deployed on the instance and the api will be available at the ip of the instance.
all step was written in the terminal.
when the instance is deployed, you can use the `cli` to interact with the api.

you can enter `help` to see all the available commands.
> and there are hidden commands that name is `logs` , this command will show you the logs of the instance.

Little update coming soon, for the project to be more user-friendly and beautiful.

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Contributores

[![](https://contributors-img.web.app/image?repo=tot0p/runner)](https://github.com/tot0p/runner/graphs/contributors)