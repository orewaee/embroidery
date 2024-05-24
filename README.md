## About

A simple mini-project that is a portfolio of embroidery designs.
The server part is written using the net/http std lib.
MongoDB is used as the database.
Static image files are stored in one of the app dirs.
Nuxt is used to write the client part.
I'm not that good at writing frontend, so there may be some problems with it.

Images of the designs shown in the screenshot below are owned by [EmbroEVCH](https://t.me/+bwtwEvm0Zr0zNjEy).
I'm not going to use them to make money. I really like their work.

<img src="https://i.imgur.com/UC3oxIY.png" alt="screenshot">


## Run

For the app to fully work, you need to run the server and client parts, as well as the database.
Use the commands provided below to achieve this.

Commands must be executed after providing environment variables.
A list of them for each part of the app can be found in the `.env.example` file in the corresponding dirs.


#### Server

To run the server side, use the `go run .` command in the root dir of the project.


#### Client

To run the client side, first go to the `web` dir.
In it, use the commands `pnpm install` and `pnpm dev`.


#### Database

You can deploy database in any way convenient for you, for example, use Docker for this.
In the `deploy` dir there is a config file `docker-compose.yml`.


## Routes

At the moment, only two routes are provided to receive data:

- `GET /designs` - returns a list of visible designs along with their ids, names, and tags.
- `GET /design/{id}` - returns an image byte array with content type `image/jpeg`. 
