
# Someday, maybe

+ think about renaming prepend (verb) to prefix (used as a verb)
+ clip function would find the path in current environment variable and remove it
+ rearrange path to force a particular subpath to the front or back
    + action sequence
        + clip then either prepend or append as desired
    + use cases
        + For PATH="/bin:/usr/bin:/usr/local/bin" modify it to have /usr/local/bin first
            + mergepath --prefix /usr/local/bin
        + For PATH="/bin:/usr/bin:/usr/local/bin" modify to have /usr/bin last
            + mergepath --append /usr/bin
