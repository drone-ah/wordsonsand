# Despatches

Supports automated posting to BlueSky & Reddit.

# Running

## Env Vars

Requires the following environmental variables

- `APP_BLUESKY_USERNAME`: BlueSky Username
- `APP_BLUESKY_PASSWORD`: Settings -> Privacy & Security -> App passwords
- `APP_REDDIT_CLIENT_ID`: Client ID (https://www.reddit.com/prefs/apps/)
- `APP_REDDIT_CLIENT_SECRET`: Client Secret
- `APP_REDDIT_REFRESH_TOKEN`: OAuth2 Refresh Token (see my blog post:
  https://drone-ah.com/2025/07/01/automated-posting-to-bluesky-reddit/)

It also expects one parameter - the path to where the despatches are:

`poetry run ./despatch.py -- ../<path-to-despatches>/`

The despatches directory is expected to follow the pattern of

`YEAR/MONTH` in `yyyy/mm` format so that the script doesn't get slower over time
traversing a larger and larger directory.
