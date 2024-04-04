## About

A simple mini-project that is a portfolio of embroidery designs.
The server part is written using the std Go http library.
MongoDB is used as the database.
Static image files are stored in one of the application directories.

Images of designs shown in the screenshot are owned by [EmbroEVCH](https://t.me/+bwtwEvm0Zr0zNjEy).
I'm not going to use them to make money. I really like their work.

<img src="https://i.imgur.com/HGwmjYB.png" alt="screenshot">


## Launch

To start, you need to start the server and client parts.
To run the server side, use the `go run .` command in the root directory of the project.
To run the client side, first go to the `web` directory.
In it, use the commands `pnpm install` and `pnpm dev`.
The commands must be executed after providing the environment variables.
Their list can be found in the corresponding directories in the `.env.example` files.