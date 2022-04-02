# Projectionist in Go

A commandline utility that interprets `.projections.json` files. These files
are used by [`vim-projectionist`][1], but as the author says: there's nothing
Vim-specific about the configuration. Once completed, this tool will allow you
to more easily integrate projections into your non-Vim text editor.

I've chosen to write this utility in Go because I've never written Go before.
If you're interested in using Projectionist *right now*, this is not the program
for you. I recommend checking out [`vim-projectionist`][1] if you're a Vim user,
or [`projectionist`][2] if you're not a Vim user.

## System requirements

Once this software has any working functionality I will provide a binary.

To build this software from source:

- Go 1.18+

To develop and run tests against the source code you'll also need:

- [`shelltestrunner` 1.9+](https://github.com/simonmichael/shelltestrunner)

## Roadmap

I have made no progress so far. I have one priority feature I'd like to
implement so I can start using this tool day-to-day:

- **Support alternates.**
  For example: If I'm in `app/models/book.rb`, my editor should be able to
  auto-navigate to `test/models/book_test.rb`.

[1]: https://github.com/tpope/vim-projectionist
[2]: https://github.com/glittershark/projectionist
