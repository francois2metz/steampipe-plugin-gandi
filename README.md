# Steampipe plugin for Gandi

Use SQL to query your domains and more from [Gandi][].

- **[Get started →](docs/index.md)**
- Documentation: [Table definitions & examples](docs/tables)

## Quick start

Install the plugin with [Steampipe][]:

    steampipe plugin install francois2metz/gandi

## Development

To build the plugin and install it in your `.steampipe` directory

    make

Copy the default config file:

    cp config/gandi.spc ~/.steampipe/config/gandi.spc

## License

Apache 2

[steampipe]: https://steampipe.io
[gandi]: https://www.gandi.net/
