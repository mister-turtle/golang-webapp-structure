# Indicators of Compromise

This is the first draft of an attempt at a reasonable project layout in Go, with interfaces at layer
boundaries. I wanted an example project that wasn't the usual TODO application, so went with
something a bit closer to home. This shows a list of indicators of compromise that are entered by
the user, pretty simple.

## Usage

This application has executables in the [cmd/](cmd/) directory.

- `httpd` Runs the web application

  Args:

  - `-l=127.0.0.1` The IP/host to listen on.

  -  `-p=9000` The port for the application to run on
