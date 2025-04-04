## v1.0.0 [2025-04-05]

_What's new?_

- **[Breaking change]** Remove support for the deprecated api key. You must use Persona Access Token.
- Ignore 404 and 403 errors to work with aggregate connections.
- Update steampipe sdk to 5.11.5.

## v0.5.0 [2025-03-01]

_What's new?_

- Update go to 1.23
- Update steampipe sdk to 5.11.3

## v0.4.0 [2024-02-17]

_What's new?_

- Add support for Gandi Personal Access Token. The api key is now deprecated.
- Update SDK to 5.8.0

## v0.3.0 [2023-11-10]

_What's new?_

- Update SDK to 5.6.3
- Update to go 1.21
- Add `autorenew_duration`, `autorenew_duration_type` and `autorenew_enabled` columns to the `gandi_mailbox` table.

## v0.2.0 [2023-10-06]

_What's new?_

- Update SDK to 5.6.2
- Add `expired_at` column to the `gandi_mailbox` table.

## v0.1.2 [2022-11-10]

_What's new?_

- Update SDK to 4.1.8
- Add `antispam` column to the `gandi_mailbox` table.

## v0.1.1 [2022-10-14]

_What's new?_

- Update SDK to 4.1.7
- Add `expiration` column to the `gandi_certificate` table, thanks @jdenoy.

## v0.1.0 [2022-09-01]

_What's new?_

- Update SDK to 4.1.5
- Update to go 1.19
- The default API key is commented

## v0.0.5 [2022-06-27]

_What's new?_

- Update go-gandi dependency

## v0.0.4 [2022-06-11]

_What's new?_

- Add gandi_livedns_record table
- Add livedns status, namerserver and dnssec status on the domain
- Update logo and color

## v0.0.3 [2022-05-17]

_What's new?_

- Updated doc
- Add error logs

## v0.0.2 [2022-05-07]

_What's new?_

- Update to go 1.18
- Update sdk to 3.1
- Add gandi_web_redirection table
- Add gandi_certificate table

## v0.0.1 [2022-01-25]

_What's new?_

- Initial release
