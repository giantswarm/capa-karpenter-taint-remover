# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Add `io.giantswarm.application.audience: all` annotation to publish the app to the customer Backstage catalog.
- Migrate chart metadata annotations to `io.giantswarm.application.*` format.

## [1.0.2] - 2026-02-19

### Changed

- Migrate to App Build Suite (ABS) for building and publishing Helm charts.

## [1.0.1] - 2025-10-16

### Changed

- Use default catalog

## [1.0.0] - 2025-10-16

### Changed

- Not managing the CAPI taint anymore, it manages the `karpenter.sh/unregistered` karpenter taint that is applied by CAPI controllers when using a stale version for the patch request.

[Unreleased]: https://github.com/giantswarm/capa-karpenter-taint-remover/compare/v1.0.2...HEAD
[1.0.2]: https://github.com/giantswarm/capa-karpenter-taint-remover/compare/v1.0.1...v1.0.2
[1.0.1]: https://github.com/giantswarm/capa-karpenter-taint-remover/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/giantswarm/capa-karpenter-taint-remover/compare/v0.6.0...v1.0.0
