# Gator

Blog Aggregator

## Dependancies

- golang
- postgres

## Usage



## Commands

- login     : login a user, must be registered
- register  : register a new user and login
- reset     : drop all users, which cascade deletes all data
- users     : list all registered users
- agg       : periodically scrape posts from all rss feeds
- addfeed   : `[must be logged in]` user adds an rss feed to the database
- feeds     : list all rss feeds along with which user added that feed first
- follow    : `[must be logged in]` user follows an rss feed
- unfollow  : `[must be logged in]` user remove an rss feed from their list of following
- following : `[must be logged in]` show all rss feeds this user follows
- browse    : `[must be logged in]` display latest scraped feeds that the user follows

