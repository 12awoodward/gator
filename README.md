# Gator

A blog Aggregator, allows multiple users to keep track of feeds and view latest posts.

## Using Gator

You can use Gator in the command line using the following commands.

### register

```
gator register <name>
```
Registers a new user in the database and logs them in. The name given to the user must be unique.

### login

```
gator login <name>
```
Log in to the given user.

### users

```
gator users
```
Lists all registered users and identifies the currently logged in user.

### addfeed

```
gator addfeed <feed-name> <feed-url>
```
register a new feed under the given name at the specified url. Each feed URL can only be added once. The logged in user will automatically follow the added feed.

### feeds

```
gator feeds
```
Lists all available feeds.

### follow

```
gator follow <feed-url>
```
Currently logged in user will follow the given feed URL. The given feed must have already been registered.

### following

```
gator following
```
Lists all feeds the currently logged in user is following.

### unfollow

```
gator unfollow <feed-url>
```
Currently logged in user will unfollow the given feed URL.

### agg

```
gator agg <interval>
```
Triggers an infinitely running aggregator to retrieve the latest posts from registered feeds.

The given interval will determine how often it runs. For example `2m30s` will trigger every 2 minutes and 30 seconds, or `1h` will run every hour.

Use ctrl-c or close the terminal window to stop this from running.

### browse

```
gator browse [<amount>]
```
Will show most recent post information for all the current users followed feeds. It will show up to the amount given, or 2 if not provided.

The `gator agg` command should be run, to ensure the latest information is available.

### reset

```
gator reset
```
This will completely delete all database records. Only run this command if you are prepared to lose everything in Gator.

## Installing Gator

### Requirements

- go 1.24 or higher
- postgres

### How to set up

run `go install "https://github.com/12awoodward/gator"` to install Gator.

In your home directory create a file called `.gatorconfig.json`. Then put the following contents in the file:
```
{"db_url":"DATABASE_CONNECTION_STRING"}
```
Create a new postgres database for gator and replace `DATABASE_CONNECTION_STRING` with the connection string for the database.

Now you can [register your first user](#register) and start using Gator.
