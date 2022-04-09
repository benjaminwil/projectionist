# Projectionist in Go

A commandline utility that interprets `.projections.json` files.

[![Projectionist Test Suite][gh-workflow-badge]][gh-workflow-history]

These files are used by [`vim-projectionist`][1], but as the author says:
there's nothing Vim-specific about the configuration. Once completed, this tool
will allow you to more easily integrate projections into your non-Vim text
editor.

I've chosen to write this utility in Go because I've never written Go before.
If you're interested in using Projectionist *right now*, this is not the program
for you. I recommend checking out [`vim-projectionist`][1] if you're a Vim user,
or [`projectionist`][2] if you're not a Vim user.

[1]: https://github.com/tpope/vim-projectionist
[2]: https://github.com/glittershark/projectionist

[gh-workflow-badge]: https://github.com/benjaminwil/projectionist/actions/workflows/ci.yml/badge.svg
[gh-workflow-history]: https://github.com/benjaminwil/projectionist/actions/workflows/ci.yml

## Usage

I expect the commandline interface to be used like this:

    projectionist <command> [--config <projections.json>] <current-file>

Where `--config` can be inferred based on the `$PWD`.

Here's a more concrete example:

    $ projectionist alternate app/models/book.rb
    # returns: test/models/book_test.rb

## Dependencies

Once this software has any working functionality I will provide a binary.

### Building

To build this software from source, you'll need:

- Go 1.18+
- Make\*

\* Recommended, but not required.

I recommend using Make to build from source:

    $ make build

Once run, the binary will be built to `build/projectionist`.

### Development and testing

To develop and run tests against the source code you'll also need:

- [`shelltestrunner` 1.9+](https://github.com/simonmichael/shelltestrunner)

If you have all of the dependencies installed, you can use the Make to run
tests:

    $ make test

## Roadmap

I have made no progress so far. I have one priority feature I'd like to
implement so I can start using this tool day-to-day:

- **Support alternates.**
  For example: If I'm in `app/models/book.rb`, my editor should be able to
  auto-navigate to `test/models/book_test.rb`.
