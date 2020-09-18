# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Use new `v2` API endpoints for builds and deployments.\
  This also changes the output for all related commands, `ccv2ctl` returns the responses of the APIs as-is

## [0.6.0] - 2019-11-05

### Added

- Support new deployment parameters "Data Migration Mode" and "Deployment Mode"

## [0.5.0] - 2019-06-04

### Added

- Get or set environment-specific properties with the `customerproperties` command (A big "thank you" to [@yehorov-sap] for implementing this feature)
- Support new deployment modes

[@yehorov-sap]: https://github.com/yehorov-sap

[Unreleased]: https://github.com/SAP-staging/commerce-ccv2ctl/compare/v0.6.0...HEAD
[0.6.0]: https://github.com/SAP-staging/commerce-ccv2ctl/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/SAP-staging/commerce-ccv2ctl/compare/409c471d165eb8fc09ad01a5c25609e684942531...v0.5.0